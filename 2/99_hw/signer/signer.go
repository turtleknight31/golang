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
	iter := 0
	fmt.Println("SingleHash kirdi : ")
	fmt.Println(len(in))
	for i := range in {
		fmt.Println("SingleHash znacheniya in : ")
		fmt.Println(i.(int))
		if iter == 0 {
			data += DataSignerCrc32(strconv.Itoa(i.(int)))
			data += "~"
		}
		if iter > 0 {
			data += DataSignerMd5(data)
		}

		iter++
	}

	out <- data
	fmt.Println("Outka jazdyk: " + data)
	fmt.Println("Outka jazdyk test2 : " + data)

}
