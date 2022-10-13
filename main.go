package main

import (
	"fmt"
	"github.com/sftfjugg/httpDownload/down"
)

func main() {
	var url string = "https://dl.google.com/go/go1.19.2.windows-amd64.msi"

	httpDownload := down.New(url, 8)

	fmt.Printf("Bool:%v\nContent:%d\n", httpDownload.AcceptRanges, httpDownload.ContentLength)

	httpDownload.Download()
}
