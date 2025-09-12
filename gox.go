package gox

type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func Abs[T SignedNumber](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Batch[T any](items []T, batchSize int) <-chan []T {
	ch := make(chan []T)

	go func() {
		defer close(ch)

		for i := 0; i < len(items); i += batchSize {
			end := min(i+batchSize, len(items))
			ch <- items[i:end]
		}
	}()

	return ch
}
