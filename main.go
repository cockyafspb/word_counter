package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
)

// структура для хранения результатов каждого запроса
// включая кол-во вхождений строки "Go" в теле ответа
type result struct {
	url   string
	count int
	err   error
}

// word_counter -f path/to/file.txt -p workers_count -w `Word`
func main() {
	var word string
	flag.StringVar(&word, "w", "", "flag for word to find")
	var maxWorkers int
	flag.IntVar(&maxWorkers, "p", 0, "flag for workers count")
	var filepath string
	flag.StringVar(&filepath, "f", "", "flag for file's with urls path")
	flag.Parse()

	resultChannel := make(chan *result)
	done := make(chan bool)

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	go runConcurrentTasks(file, word, done, resultChannel, maxWorkers)

	go func() {
		<-done
		close(resultChannel)
	}()

	printResults(resultChannel)
}

func runConcurrentTasks(file *os.File, word string, done chan bool, resultChan chan *result, workersCount int) {
	var wg sync.WaitGroup
	c := make(chan struct{}, workersCount)

	defer func() {
		close(c)
	}()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		c <- struct{}{}
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			res := getWordCount(url, word)
			resultChan <- &res

			<-c

		}(scanner.Text())
	}
	wg.Wait()
	done <- true
}

func getWordCount(url, word string) result {
	response, err := http.Get(url)
	if err != nil {
		return result{url, 0, err}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return result{url, 0, err}
	}
	// wordCount := countWords([]rune(string(body)), word)
	regex := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
	wordCount := len(regex.FindAll(body, -1))

	return result{url, wordCount, err}
}

func printResults(resultChan chan *result) {
	totalCount := 0

	for res := range resultChan {
		if res.err != nil {
			continue
		}
		fmt.Printf("Count for %s: %d\n", res.url, res.count)

		totalCount += res.count
	}
	fmt.Printf("Total: %v\n", totalCount)
}
