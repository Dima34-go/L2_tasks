package wgetUtil

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)



func MainWget() {

	URL := "https://seasonkrasoty.ru/product/legkiy_omolazhivayushchiy_krem"

	segments := strings.Split(URL, "/")
	fileName := segments[len(segments)-1]
	fileName += ".html"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//Http request
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}