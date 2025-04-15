package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	listOfFiles := []string{"file1.txt", "file2.txt", "file3.txt"}
	ch := make(chan struct {
		content string
		file    string
	})
	n := 10
	wg := &sync.WaitGroup{}

	for _, file := range listOfFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			b, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
			}
			ch <- struct {
				content string
				file    string
			}{file: file, content: string(b)}
		}(file)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	// content := make(map[string]string)
	var content string
	for i := range ch {
		content = content + i.content
	}
	//for _, v := range content {
	freq := make(map[string]int)
	c := strings.Split(content, " ")
	for _, word := range c {
		freq[word]++
	}
	wordCount := make([]struct {
		word  string
		count int
	}, len(freq))
	for k, v := range freq {
		wordCount = append(wordCount, struct {
			word  string
			count int
		}{word: k, count: v})
	}
	sort.Slice(wordCount, func(i, j int) bool {
		return wordCount[i].count > wordCount[j].count
	})
	if len(wordCount) < n {
		n = len(wordCount)
	}
	fmt.Println(wordCount[:n])
	//}

}
