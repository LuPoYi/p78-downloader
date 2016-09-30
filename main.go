package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv" // convert string to int

	"github.com/PuerkitoBio/goquery" // A little like that j-thing, only in Go.
)

func goDownload(filename, url string) {
	fmt.Println("Downloading " + url + " ...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create("images/" + filename + ".jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
}

func goLogger(title, url string) {
	// os.O_CREATE : find or create
  f, err := os.OpenFile("download_list.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    panic(err)
  }

  defer f.Close()

  if _, err = f.WriteString(title + " : " + url + "\n"); err != nil {
    panic(err)
  }
}

func goCrawler(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul#portfolio li").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(".picture_overlay img").Eq(0).Attr("alt")
		imgURL, _ := s.Find(".picture_overlay img").Eq(0).Attr("src")
		pageURL, _ := s.Find(".picture_overlay .overlay a").Eq(1).Attr("href")

		if imgURL[0:4] != "http" {
			imgURL = "https://porn77.info/" + imgURL
		}
		if pageURL[0:4] != "http" {
			pageURL = "https://porn77.info/" + pageURL
		}

		goLogger(title, pageURL)
		goDownload(title, imgURL)
	})
}

func main() {
	for i := 1; i < 3; i++ {
		fmt.Println("goCrawler FHD Page " + strconv.Itoa(i))
		goCrawler("https://porn77.info/video/index/search/FHD/page/" + strconv.Itoa(i))
	}
}

// http://studygolang.com/articles/8359
