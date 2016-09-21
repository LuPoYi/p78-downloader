/*
圖片resp:
&{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[5607] Last-Modified:[Tue, 08 Mar 2016 05:51:31 GMT] Connection:[keep-alive] Etag:["56de6863-15e7"] Accept-Ranges:[bytes] Date:[Mon, 19 Sep 2016 11:15:40 GMT] Content-Type:[image/png]] 0xc4200146c0 5607 [] false false map[] 0xc4200ce0f0 0xc4203b5600}

網頁resp:
&{200 OK 200 HTTP/1.1 1 1 map[Content-Type:[text/html; charset=utf-8] X-Request-Id:[98dc0c16-a642-411c-93d3-6b859df5f7be] Etag:[W/"b757cf7e417b9534dea7329f3d5adeee"] Cache-Control:[max-age=0, private, must-revalidate] X-Content-Type-Options:[nosniff] Date:[Mon, 19 Sep 2016 11:16:05 GMT] Connection:[keep-alive] Vary:[Accept-Encoding] Set-Cookie:[BitoEXse2=04ad3605afd1482d57386a159b95a710; path=/; expires=Wed, 21 Sep 2016 11:16:05 -0000; HttpOnly] Status:[200 OK] X-Xss-Protection:[1; mode=block] X-Frame-Options:[SAMEORIGIN]] 0xc420438140 -1 [chunked] false true map[] 0xc4200ee0f0 0xc4202a2000}
 req.Header.Set("name", "value")
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery" // A little like that j-thing, only in Go.
)

func goDownload(url, filename string) {
	fmt.Println("Downloading " + url + " ...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename + ".jpg")
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
	fmt.Printf("Go crawler!")

	doc.Find("ul#portfolio li").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(".picture_overlay img").Eq(0).Attr("alt")
    img_url, _ := s.Find(".picture_overlay img").Eq(0).Attr("src")

    if img_url[0:3] != "http" {
      img_url = "https://porn77.info/" + img_url
    }

    goDownload(img_url, title)
	})
}

func main() {
	goCrawler("https://porn77.info/")
}
