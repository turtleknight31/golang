package main

import (
	"fmt"
	"strconv"
	"sync"
)

func ExecutePipeline(jobfree ...job) {
	in := make(chan interface{})
	wg := &sync.WaitGroup{}

	for _, singleJob := range jobfree {
		out := make(chan interface{})
		wg.Add(1)
		go func(jobs job, in, out chan interface{}) {
			defer wg.Done()
			jobs(in, out)
			close(out)

		}(singleJob, in, out)

		in = out
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {

	var data string
	fmt.Println("SingleHash kirdi : ")
	for i := range in {
		fmt.Println("SingleHash znacheniya in : ")
		fmt.Println(i.(int))
		data += DataSignerCrc32(strconv.Itoa(i.(int)))
		data += "~"
	}

	out <- data
	fmt.Println("Outka jazdyk: " + data)
	fmt.Println("Outka jazdyk test2 : " + data)

}
