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
	fmt.Println("SingleHash kirdi : ")
	fmt.Println(len(in))
	for i := range in {
		data := DataSignerCrc32(strconv.Itoa(i.(int))) + "~" + DataSignerCrc32(DataSignerMd5(strconv.Itoa(i.(int))))
		fmt.Println("SingleHash kirdi : " + data)
		out <- data
	}

}

func MultiHash(in, out chan interface{}) {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("MultiHash kirdi : ")
	for i := range in {
		var data string
		for _, j := range s {
			data += DataSignerCrc32(strconv.Itoa((j)) + i.(string))
		}
		fmt.Println("MultiHash kirdi : " + data)
		out <- data
	}
}

func CombineResults(in, out chan interface{}) {
	var data string
	j := 0
	for i := range in {
		data += i.(string)
		if j == 0 {
			data += "_"
		}
		j++
	}

	out <- data
}
