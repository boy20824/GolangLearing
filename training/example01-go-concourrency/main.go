package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	numWorkers = 10
)

func main() {
	//網址列
	urls := []string{
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.youtube.com/",
		"https://www.hexschool.com/qa/how-to.html?fbclid=IwAR0DJETsmG7rYrMfjs2fnkAURxHVnXYBz1UwBEN9_Wrmv4DWilP3Nk69U7A",
	}
	//使用WaitGroup追蹤goroutine完成情況
	var wg sync.WaitGroup

	//創建固定大小的goroutine池
	semaphore := make(chan struct{}, numWorkers)

	//匿名函式
	processURL := func(url string) {
		//後釋放wg計時器
		defer wg.Done()
		//先釋放Channel通道
		defer func() { <-semaphore }()
		// 自定義 HTTP Client 與超時時間
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		res, err := client.Get(url)
		if err != nil {
			fmt.Printf("Error :%s \n", err)
			return
		}
		fmt.Printf("URL: %s, Status Code:%d \n", url, res.StatusCode)
	}

	for _, url := range urls {
		//放置一個計時器,遇到Done()就會-1,直到歸0
		wg.Add(1)
		//會塞一個空的strurc{}{} 最多塞到10個超過10個就會阻塞 (最主要拿來控制goroutine執行數量不超過10)
		semaphore <- struct{}{}
		go processURL(url)
	}

	//會等計時器為0,才會放行
	wg.Wait()
	fmt.Println("done")
}
