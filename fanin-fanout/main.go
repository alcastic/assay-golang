package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var workerNumber int
	flag.IntVar(&workerNumber, "w", 1, "number of worker for fanout data processing")
	flag.Parse()
	startTime := time.Now()

	fmt.Println("#### Fanin-Fanout demo start")
	dataStream := streamFileLines("rock_bands")

	// One data stream processed by multiples workers
	//                  __worker1
	//    dataStream __|__worker3
	//                 |__  ...
	//                 |__workerN

	workerStreams := Fanout(dataStream, toUpperCase, workerNumber)

	// Multiple data streams reconciliated to one data stream
	//    dataStream1__
	//    dataStream2__|__reconciliationDataStream
	//       ...     __|
	//    dataStreamN__|
	reconciliationDataStream := Fanin(workerStreams...)

	for d := range reconciliationDataStream {
		fmt.Println(d)
	}

	fmt.Printf("#### Elapse: %f\n", time.Since(startTime).Seconds())
	fmt.Println("#### Fanin-Fanout demo Finish")
}

func Fanin(streams ...<-chan string) <-chan string {
	chout := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(streams))
		for _, stream := range streams {
			go func(s <-chan string, wg *sync.WaitGroup) {
				for cad := range s {
					chout <- cad
				}
				wg.Done()
			}(stream, &wg)
		}
		wg.Wait()
		close(chout)
	}()
	return chout
}

func Fanout(dataStream <-chan string, pipeStream func(<-chan string) <-chan string, pipeStreamNumber int) []<-chan string {
	out := make([]<-chan string, pipeStreamNumber)
	for i := 0; i < pipeStreamNumber; i++ {
		out[i] = pipeStream(dataStream)
	}
	return out

}

func toUpperCase(chin <-chan string) <-chan string {
	chout := make(chan string)

	go func() {
		for d := range chin {
			chout <- strings.ToUpper(d)
		}
		close(chout)
	}()

	return chout
}

func streamFileLines(fileName string) <-chan string {
	ch := make(chan string)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		file.Close()
		close(ch)
	}()

	return ch
}
