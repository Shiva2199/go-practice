package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

/*
Sure! Below is a medium-level Go language problem that covers multiple important concepts for a senior engineer, including concurrency, error handling, data structures, and performance optimization. The problem also gives room for the candidate to demonstrate their ability to write clean, efficient, and maintainable code.

---

### **Problem: Concurrent Web Scraper with Rate Limiting**

#### **Problem Statement:**

You are tasked with implementing a **concurrent web scraper** in Go. The scraper should fetch the content of multiple URLs and parse them to extract and return certain data. Specifically, you need to fetch data from a list of URLs and extract the title of the page from the HTML.

### **Requirements:**

1. **Concurrency:**
  - Use goroutines to make concurrent requests for the URLs.
  - Limit the number of concurrent requests using a rate limiter (for example, no more than 5 concurrent requests at a time).

2. **Error Handling:**
  - Properly handle errors during HTTP requests (e.g., network errors, timeouts, non-2xx HTTP status codes).
  - If any URL fails, you must log the error and continue with the other URLs. Failures should not terminate the program.

3. **Rate Limiting:**
  - Implement a rate limiter that ensures no more than **5 concurrent requests** are made at a time. You can use a `sync.Semaphore` or another concurrency mechanism to ensure the limit is adhered to.

4. **Data Parsing:**
  - For each successful request, extract the **title** of the HTML page.
  - The title can be found within the `<title></title>` tag in the HTML content of the page.
  - If the title is not found, return an empty string.

5. **Efficient Resource Management:**
  - Use appropriate synchronization mechanisms (e.g., channels, `sync.WaitGroup`) to ensure proper handling of goroutines.
  - Ensure that the program terminates gracefully once all URLs are processed, even if some requests fail.

6. **Optional Bonus (Performance Considerations):**
  - Consider optimizing the program for high performance with minimal memory usage. For instance, make sure the program handles a large number of URLs efficiently.

### **Input:**

- A list of URLs (strings) to scrape, passed as an input argument to the program.

### **Output:**

- A map where the key is the URL and the value is the extracted title of the page. If a URL fails to be fetched or parsed, return an empty string for that URL.

### **Sample Input:**

```go

	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
	}

```

### **Sample Output:**

```go

	{
		"https://example.com": "Example Domain",
		"https://golang.org": "Go Programming Language",
		"https://github.com": "GitHub: Where the world builds software"
	}

```

### **Constraints:**

- The number of URLs can range from 1 to 1000.
- HTTP requests may have delays, timeouts, or occasional failures.
- The program should be able to handle up to 1000 URLs without crashing.

---

### **Hints:**

- Use Go's standard library packages, such as `net/http`, `html`, and `strings`.
- Handle errors gracefully—don't crash the program if one URL fails.
- Consider using a `sync.WaitGroup` to synchronize goroutines and a `chan struct{}` or a `sem` to limit concurrent requests.
- Use a `defer` statement to close any resources (e.g., HTTP response body) once you're done with them.

---

### **Evaluation Criteria:**
1. **Concurrency Handling:** Proper use of goroutines, channels, and synchronization mechanisms.
2. **Error Handling:** Ensuring that errors do not terminate the program and are handled properly.
3. **Rate Limiting:** Correct implementation of rate limiting to avoid exceeding the concurrent request limit.
4. **Code Cleanliness and Efficiency:** Clean, readable, and well-structured code with consideration for performance and memory efficiency.
5. **Edge Case Handling:** Consideration of edge cases such as invalid URLs, empty responses, or missing titles.

---

This problem covers several important concepts for a senior engineer, such as:

- **Concurrency**: Using goroutines and synchronizing them.
- **Error handling**: Gracefully managing errors without terminating the program.
- **Rate limiting**: Efficiently limiting the number of concurrent requests.
- **Data parsing**: Working with HTML and extracting relevant data.
- **Performance**: Ensuring the scraper is scalable and efficient.

The candidate's solution should show their ability to design a clean, efficient, and fault-tolerant scraper.
*/

func main() {
	url := []string{
		//"https://exmaple.com",
		"https://golang.org",
		"https://github.com",
	}
	result := make(map[string]string)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	concucrrencyLimit := 1
	//ch := make(chan string, concucrrencyLimit)
	sem := make(chan struct{}, concucrrencyLimit)
	for _, u := range url {
		wg.Add(1)
		go func(u string) {
			fmt.Println(u)
			defer wg.Done()
			sem <- struct{}{}
			defer func() {
				<-sem
			}()
			res, err := http.Get(u)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer res.Body.Close()
			if res.StatusCode < 200 || res.StatusCode > 300 {
				fmt.Println("Non - 2xx status")
				return
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				// ch <- ""
				return
			}
			//strings. (string(body),"<title>")
			data := string(body)
			mu.Lock()
			result[u] = extractTitle(data)
			mu.Unlock()
			// ch <- fmt.Sprintf("url: %s and Title %s", u, extractTitle(data))
			//fmt.Println(string(body))
		}(u)
	}
	// go func() {
	// 	wg.Wait()
	// 	fmt.Println("wait done")
	// 	// close(ch)
	// }()
	// for c := range result {
	// 	fmt.Println(re)
	// }
	wg.Wait()
	fmt.Print(result)

}

func extractTitle(body string) string {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return ""
	}
	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return title
}
