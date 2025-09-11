package quotientfilter

import (
	"fmt"
	"hash/fnv"
)

// Metadata bit masks
const (
	isOccupiedMask   uint64 = 1 << 0 // Least significant bit
	isContinuationMask uint64 = 1 << 1
	isShiftedMask    uint64 = 1 << 2
	remainderOffset  uint8  = 3 // Remainder starts after 3 metadata bits
)

// QuotientFilter represents a Quotient Filter data structure.
type QuotientFilter struct {
	table   []uint64 // Stores remainder and metadata bits
	q       uint8    // Number of bits for the quotient
	r       uint8    // Number of bits for the remainder
	p       uint8    // Total bits for the fingerprint (q + r)
	size    uint64   // Current number of elements
	maxSize uint64   // Maximum capacity (2^q)
}

// New creates and initializes a new QuotientFilter.
// q: number of bits for the quotient (determines table size 2^q)
// r: number of bits for the remainder
func New(q, r uint8) (*QuotientFilter, error) {
	if q == 0 || r == 0 {
		return nil, fmt.Errorf("q and r must be greater than 0")
	}
	if q+r > 64 { // Fingerprint cannot exceed uint64
		return nil, fmt.Errorf("q + r cannot exceed 64 bits")
	}
	if r+remainderOffset > 64 { // Remainder + metadata bits must fit in uint64
		return nil, fmt.Errorf("remainder bits + metadata bits cannot exceed 64 bits")
	}

	maxSize := uint64(1) << q
	table := make([]uint64, maxSize)

	return &QuotientFilter{
		table:   table,
		q:       q,
		r:       r,
		p:       q + r,
		size:    0,
		maxSize: maxSize,
	}, nil
}

// hashData hashes the input data and returns a p-bit fingerprint.
func (qf *QuotientFilter) hashData(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	fingerprint := h.Sum64()

	// Mask to get only p bits
	return fingerprint & ((1 << qf.p) - 1)
}

// getQuotient extracts the quotient from a p-bit fingerprint.
func (qf *QuotientFilter) getQuotient(fingerprint uint64) uint64 {
	return fingerprint >> qf.r
}

// getRemainder extracts the remainder from a p-bit fingerprint.
func (qf *QuotientFilter) getRemainder(fingerprint uint64) uint64 {
	return fingerprint & ((1 << qf.r) - 1)
}

// setRemainder sets the remainder in a slot.
func (qf *QuotientFilter) setRemainder(slot *uint64, remainder uint64) {
	// Clear existing remainder bits and then set new remainder
	*slot = (*slot & (isOccupiedMask | isContinuationMask | isShiftedMask)) | (remainder << remainderOffset)
}

// getRemainderFromSlot extracts the remainder from a slot.
func (qf *QuotientFilter) getRemainderFromSlot(slot uint64) uint64 {
	return (slot >> remainderOffset) & ((1 << qf.r) - 1)
}

// setOccupied sets the is_occupied bit for a slot.
func (qf *QuotientFilter) setOccupied(slot *uint64) {
	*slot |= isOccupiedMask
}

// clearOccupied clears the is_occupied bit for a slot.
func (qf *QuotientFilter) clearOccupied(slot *uint64) {
	*slot &^= isOccupiedMask
}

// isOccupied checks if the is_occupied bit is set for a slot.
func (qf *QuotientFilter) isOccupied(slot uint64) bool {
	return (slot & isOccupiedMask) != 0
}

// setContinuation sets the is_continuation bit for a slot.
func (qf *QuotientFilter) setContinuation(slot *uint64) {
	*slot |= isContinuationMask
}

// clearContinuation clears the is_continuation bit for a slot.
func (qf *QuotientFilter) clearContinuation(slot *uint64) {
	*slot &^= isContinuationMask
}

// isContinuation checks if the is_continuation bit is set for a slot.
func (qf *QuotientFilter) isContinuation(slot uint64) bool {
	return (slot & isContinuationMask) != 0
}

// setShifted sets the is_shifted bit for a slot.
func (qf *QuotientFilter) setShifted(slot *uint64) {
	*slot |= isShiftedMask
}

// clearShifted clears the is_shifted bit for a slot.
func (qf *QuotientFilter) clearShifted(slot *uint64) {
	*slot &^= isShiftedMask
}

// isShifted checks if the is_shifted bit is set for a slot.
func (qf *QuotientFilter) isShifted(slot uint64) bool {
	return (slot & isShiftedMask) != 0
}

// findRunStart finds the physical index where the run for a given quotient starts.
func (qf *QuotientFilter) findRunStart(quotient uint64) uint64 {
	// Count the number of occupied slots before the current quotient's canonical slot.
	// This gives the number of logical runs that start before or at this quotient.
	numOccupiedSlotsBefore := uint64(0)
	for i := uint64(0); i < quotient; i++ {
		if qf.isOccupied(qf.table[i]) {
			numOccupiedSlotsBefore++
		}
	}

	// Count the number of shifted slots before the current quotient's canonical slot.
	// This gives the number of elements that have been shifted past their canonical position.
	numShiftedSlotsBefore := uint64(0)
	for i := uint64(0); i < quotient; i++ {
		if qf.isShifted(qf.table[i]) {
			numShiftedSlotsBefore++
		}
	}

	// The logical index of the first slot for this quotient's run.
	// This is the quotient itself, plus any elements that were shifted into earlier slots,
	// minus the number of runs that started before this quotient.
	logicalRunStart := quotient + numShiftedSlotsBefore - numOccupiedSlotsBefore

	// Now, find the physical index of this logical run start.
	// This is the first slot from `logicalRunStart` that is not a continuation.
	physicalRunStart := logicalRunStart
	for physicalRunStart > 0 && qf.isContinuation(qf.table[physicalRunStart-1]) {
		physicalRunStart--
	}
	return physicalRunStart
}

// Insert adds an element to the QuotientFilter.
func (qf *QuotientFilter) Insert(data []byte) error {
	if qf.size == qf.maxSize {
		return fmt.Errorf("quotient filter is full")
	}

	fingerprint := qf.hashData(data)
	quotient := qf.getQuotient(fingerprint)
	remainder := qf.getRemainder(fingerprint)

	// Mark the canonical slot as occupied
	qf.setOccupied(&qf.table[quotient])

	// Find the physical index where the run for this quotient starts
	runStartIdx := qf.findRunStart(quotient)

	// Find the insertion point within the run, maintaining sorted order of remainders
	insertIdx := runStartIdx
	for insertIdx < qf.maxSize && qf.isShifted(qf.table[insertIdx]) && qf.getRemainderFromSlot(qf.table[insertIdx]) < remainder {
		insertIdx++
	}

	// If insertIdx has reached maxSize, it means there's no space to insert.
	// This should ideally be caught by the qf.size == qf.maxSize check,
	// but if the filter is almost full and a long run pushes the insertion point
	// beyond the end, this check is needed.
	if insertIdx >= qf.maxSize {
		return fmt.Errorf("quotient filter is full (no space for insertion)")
	}

	// Shift elements to the right to make space for the new remainder
	// If insertIdx is already at maxSize-1 and the table is full, this could be an issue.
	// The `qf.size == qf.maxSize` check at the beginning should prevent overflow.
	// However, if `insertIdx` is `qf.maxSize-1` and `qf.table[qf.maxSize-1]` is occupied,
	// shifting will push it out. This is a known limitation of QF if not resized.
	// For this basic implementation, we assume no resizing.
	if qf.size == qf.maxSize { // Double check for safety, though already checked
		return fmt.Errorf("quotient filter is full, cannot insert")
	}

	// Shift elements from the end of the table down to insertIdx+1
	for i := qf.maxSize - 1; i > insertIdx; i-- {
		qf.table[i] = qf.table[i-1]
		// For shifted elements, set their shifted bit
		qf.setShifted(&qf.table[i])
	}

	// Insert the new remainder at insertIdx
	qf.table[insertIdx] = 0 // Clear all bits initially
	qf.setRemainder(&qf.table[insertIdx], remainder)

	// Update metadata bits for the newly inserted element
	if insertIdx != quotient {
		qf.setShifted(&qf.table[insertIdx])
	} else {
		qf.clearShifted(&qf.table[insertIdx]) // If it's in its canonical slot, it's not shifted
	}

	if insertIdx > runStartIdx {
		qf.setContinuation(&qf.table[insertIdx])
	} else {
		qf.clearContinuation(&qf.table[insertIdx])
	}

	qf.size++
	return nil
}

// Contains checks if an element is probably in the QuotientFilter.
func (qf *QuotientFilter) Contains(data []byte) bool {
	fingerprint := qf.hashData(data)
	quotient := qf.getQuotient(fingerprint)
	remainder := qf.getRemainder(fingerprint)

	// If the canonical slot is not marked as occupied, the element is definitely not in the filter.
	if !qf.isOccupied(qf.table[quotient]) {
		return false
	}

	// Find the physical index where the run for this quotient starts
	runStartIdx := qf.findRunStart(quotient)

	// Scan the run for the remainder
	currentIdx := runStartIdx
	for currentIdx < qf.maxSize && qf.isShifted(qf.table[currentIdx]) { // Iterate while elements are shifted (part of a run)
		currentRemainder := qf.getRemainderFromSlot(qf.table[currentIdx])
		if currentRemainder == remainder {
			return true // Found the remainder, element is probably in the filter
		}
		if currentRemainder > remainder {
			// Remainders are sorted within a run, so if we've passed it, it's not here.
			return false
		}
		currentIdx++
	}
	return false // Remainder not found in the run
}

// Size returns the number of elements currently in the filter.
func (qf *QuotientFilter) Size() uint64 {
	return qf.size
}

// Capacity returns the maximum number of elements the filter can hold.
func (qf *QuotientFilter) Capacity() uint64 {
	return qf.maxSize
}
