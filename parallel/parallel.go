package parallel

import (
	"log"
	"sync"
)

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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for r := range results {
			log.Printf("result: %v", r)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range errors {
			log.Printf("error: %v", e)
		}
	}()
	wg.Wait()
}
