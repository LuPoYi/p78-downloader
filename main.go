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

func goDownload(url, filename string) {
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

func goCrawler(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul#portfolio li").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(".picture_overlay img").Eq(0).Attr("alt")
    img_url, _ := s.Find(".picture_overlay img").Eq(0).Attr("src")

    if img_url[0:4] != "http" {
      img_url = "https://porn77.info/" + img_url
    }

    goDownload(img_url, title)
	})
}

func main() {
	for i := 1 ; i < 3 ; i++ {
		fmt.Println("goCrawler FHD Page " + strconv.Itoa(i))
		goCrawler("https://porn77.info/video/index/search/FHD/page/" + strconv.Itoa(i))
	}
}


// http://studygolang.com/articles/8359