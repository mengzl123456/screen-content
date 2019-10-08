package main

import (
	"fmt"
	"log"
	"net/http"
	"screen-content/screenContent"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	//静态文件
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("html/static"))))
	//http.Handle("/",http.FileServer(http.Dir("html/index.html")))
	http.HandleFunc("/screenContentByOneSeparator", screenContent.ScreenContentBySeparator)
	http.Handle("/", http.FileServer(http.Dir("html/")))
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println()
		log.Fatal(err)
	}
}
