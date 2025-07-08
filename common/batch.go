package common

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
