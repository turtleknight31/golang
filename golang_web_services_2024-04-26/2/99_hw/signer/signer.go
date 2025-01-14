package main

import (
	"fmt"
	"strconv"
	"sync"
)

func ExecutePipeline(jobs ...job) {

	var wg sync.WaitGroup
	in := make(chan interface{})
	for _, singleJob := range jobs {
		fmt.Println("Uyatsyz Nurlan kirdi ")
		out := make(chan interface{})
		wg.Add(1)
		go func(j job, in, out chan interface{}) {
			defer wg.Done()

			j(in, out)
			close(out)
		}(singleJob, in, out)

		in = out
	}

	wg.Wait()

}

func SingleHash(in, out chan interface{}) {

	for i := range in {
		value := i.(int)
		fmt.Println("Single Hash kabyldady: " + strconv.Itoa(value))
		o := DataSignerCrc32(strconv.Itoa(value))
		fmt.Println("Single Hash shygardy: " + o)
		out <- o
	}

}
