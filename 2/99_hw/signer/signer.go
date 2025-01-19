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
	for i := range in {
		data = DataSignerCrc32(strconv.Itoa(i.(int)))
		data += "~"
		data += DataSignerCrc32(DataSignerMd5(strconv.Itoa(i.(int))))
		fmt.Println("Result singlehash: " + data)
		out <- data

	}
	/*
	   var s = []int{}

	   	for i := range in {
	   		s = append(s, i.(int))
	   	}

	   	for _, j := range s {
	   		data := DataSignerCrc32(strconv.Itoa(j) + "~" + DataSignerMd5(strconv.Itoa(j)))
	   		out <- data
	   	}
	*/
}

func MultiHash(in, out chan interface{}) {

	for i := range in {
		mu := &sync.Mutex{}
		wg := &sync.WaitGroup{}
		/* for _, j := range s {
			data += DataSignerCrc32(strconv.Itoa(j) + i.(string))
		}
		*/

		wg.Add(1)

		go func(i string) {
			defer wg.Done()
			var data string
			s := []int{0, 1, 2, 3, 4, 5}
			for _, j := range s {
				mu.Lock()
				data += DataSignerCrc32(strconv.Itoa(j) + i)
				mu.Unlock()
			}
			fmt.Println("Result multihash: " + data)
			out <- data

		}(i.(string))

		wg.Wait()

		/*
			data = DataSignerCrc32(ReturnText(i.(string), s))

					fmt.Println("Result multihash: " + data)

				fmt.Println("Result multihash: " + data)
				out <- data
		*/
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
	fmt.Println(data)
	out <- data
}

func ReturnText(s string, slice []int) string {
	var data string

	for _, j := range slice {
		data += strconv.Itoa(j) + s
	}

	return data
}
