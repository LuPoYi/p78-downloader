package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

type downloadTask struct {
	filename string
	url      string
}


func truncateString(str string, maxLength int) string {
	if len(str) > maxLength {
			return str[:maxLength]
	}
	return str
}


func goDownload(task downloadTask, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Downloading " + task.url + " ...")
	resp, err := http.Get(task.url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create("images/" + task.filename + ".jpg")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		log.Println(err)
	}
}

func goLogger(title, url string) {
	f, err := os.OpenFile("download_list.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	if _, err = f.WriteString(title + " : " + url + "\n"); err != nil {
		log.Println(err)
	}
}

func goCrawler(url string, page int, wg *sync.WaitGroup, tasks chan downloadTask) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err)
		return
	}

	doc.Find("ul#portfolio li").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(".picture_overlay img").Eq(0).Attr("alt")
		imgURL, _ := s.Find(".picture_overlay img").Eq(0).Attr("src")
		pageURL, _ := s.Find(".picture_overlay .overlay a").Eq(0).Attr("href")
		datetime, _ := s.Find("time").Eq(0).Attr("datetime")

		fileName := strconv.Itoa(page) + "_" + strings.Replace(datetime, "/", "", -1) + "_" + title

		goLogger(fileName, pageURL)
		tasks <- downloadTask{filename: fileName, url: imgURL}
	})
	wg.Done()
}

func main() {
	godotenv.Load()
	url := os.Getenv("URL")

	var wg sync.WaitGroup
	tasks := make(chan downloadTask)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		
    
		go goCrawler(url+strconv.Itoa(i), i, &wg, tasks)
	}

	go func() {
		wg.Wait()
		close(tasks)
	}()

	for task := range tasks {
		wg.Add(1)
		go goDownload(task, &wg)
	}

	wg.Wait()
}
