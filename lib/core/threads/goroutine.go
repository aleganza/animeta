package threads

type ParallelFetchFunc[T any] struct {
	Name string
	Func func() (T, error)
}

type ParallelFetchResult[T any] struct {
	Name string
	Body T
	Err  error
}

func RunParallel[T any](funcs ...ParallelFetchFunc[T]) []ParallelFetchResult[T] {
	ch := make(chan ParallelFetchResult[T], len(funcs))

	for _, fn := range funcs {
		go func(fn ParallelFetchFunc[T]) {
			data, err := fn.Func()

			ch <- ParallelFetchResult[T]{
				Name: fn.Name,
				Body: data,
				Err:  err,
			}
		}(fn)
	}

	results := make([]ParallelFetchResult[T], 0, len(funcs))

	for range funcs {
		results = append(results, <-ch)
	}

	return results
}