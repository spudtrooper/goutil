package parallel

import (
	"log"
	"sync"
)

func ExecAndDrain(collection chan interface{}, threads int, fn func(interface{}) (interface{}, error)) {
	res, errs := Exec(collection, threads, fn)
	EmptyDrain(res, errs)
}

func DoTimes(threads int, fn func()) {
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}

func Exec(collection chan interface{}, threads int, fn func(interface{}) (interface{}, error)) (chan interface{}, chan error) {
	results := make(chan interface{})
	errors := make(chan error)
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for x := range collection {
					res, err := fn(x)
					if err != nil {
						errors <- err
					} else {
						results <- res
					}
				}
			}()
		}
		wg.Wait()
		close(results)
		close(errors)
	}()

	return results, errors
}

func LazyDrain(results chan interface{}, errors chan error) {
	WaitFor(func() {
		for r := range results {
			log.Printf("result: %v", r)
		}
	}, func() {
		for e := range errors {
			log.Printf("error: %v", e)
		}
	})
}

func EmptyDrain(results chan interface{}, errors chan error) {
	WaitFor(func() {
		for range results {
		}
	}, func() {
		for range errors {
		}
	})
}

func WaitFor(fns ...func()) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		fn := fn
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}
