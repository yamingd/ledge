package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getFileName(url string) (result string) {
	h := md5.New()
	io.WriteString(h, url)
	result = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func getUrl(url string) int {
	r, err := http.Get(url)
	if err != nil {
		log.Printf("Error:%v\n", err)
		return 0
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error:%v\n", err)
		return 0
	}

	fname := fmt.Sprintf("%v.html", getFileName(url))
	f, err := os.Create(fname)
	if err != nil {
		log.Printf("Error:%v\n", err)
		return 0
	}

	w := bufio.NewWriter(f)
	defer w.Flush()

	ret, err := w.WriteString(string(b))
	if err != nil {
		log.Printf("Error:%v\n", err)
		return 0
	}
	return ret
}

func start(total int, ch chan string, quit chan int) {
	num := 0
	for {
		url := <-ch
		size := getUrl(url)
		num++
		fmt.Printf("get %v, size=%v\n", url, size)
		if num == total {
			quit <- 0
			break
		} else {
			quit <- 1
		}
	}
}

//
// Read datasource
//
func feed(urls []string, ch chan string, quit chan int) {
	for i := 0; i < len(urls); i++ {
		ch <- urls[i]
	}

	for {
		flag := <-quit
		if flag == 0 {
			fmt.Println("Quit.")
			break
		}
	}
}

func main() {
	urls := []string{
		"http://bing.com/robots.txt",
		"http://baidu.com/robots.txt",
		"http://cn.bing.com/search?q=godoc&go=&qs=bs&intlF=&upl=zh-chs&first=11&FORM=PERE"}

	ch := make(chan string, 2)
	quit := make(chan int)

	go start(len(urls), ch, quit)

	feed(urls, ch, quit)
}
