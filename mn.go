package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var url string

func main() {
	var r *http.Response
	var err error
	var mode = "aheaders"
	var contentType string
	var b = bytes.NewBuffer(nil)

	if len(os.Args) == 1 {
		fmt.Printf("%v\n", help)
		os.Exit(0)
	}
	url = os.Args[1]
	if r, err = http.Get(url); err != nil {
		fmt.Printf("getrequester: err=%v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		switch os.Args[2] {
		case "all":
			mode = "all"
		case "head":
			mode = "head"
		}
	}

	contentType = r.Header.Get("Content-Type")
	if contentType == "" {
		fmt.Printf("getrequester: Content-Type is empty!\n")
		mode = "aheaders"
	}

	switch mode {
	case "all":
		r.Write(b)
		fmt.Printf("%v\n", string(b.Bytes()))
	case "aheaders":
		printHeaders(r)
	case "head":
		printHeaders(r)
		printHeadTag(r)
	}

}

//if error of reading the body it prinnts a fragment (perhaps empty) of the body instead of head tag
//if is not a head tag it prints only error
func printHeadTag(r *http.Response) {
	var head string
	var err error
	var b = bytes.NewBuffer(nil)
	var n int64
	var openInd, closeInd int
	//if len(r.Body) == 0 {
	//	fmt.Printf("getrequester.getHeadTag: no body\n")
	//	return
	//}
	if n, err = b.ReadFrom(r.Body); err != nil {
		fmt.Printf("getrequester.getHeadTag err=%v\n", err.Error())
		head = b.String()
		if head == "" {
			fmt.Printf("getrequester.getHeadTag: nothing was read from the body\n")
		} else {
			fmt.Printf("%v\n", head)
		}
		return
	}
	if n == 0 {
		fmt.Printf("getrequester.getHeadTag: no body\n")
		return
	}
	r.Body.Close()
	openInd = strings.Index(b.String(), "<head>")
	closeInd = strings.Index(b.String(), "</head>")
	if openInd == -1 || closeInd == -1 {
		fmt.Printf("getrequester.getHeadTag no head tag (%v;%v)", openInd, closeInd)
		return
	}
	head = b.String()
	head = head[openInd : closeInd+6]
	fmt.Printf("%v\n", head)
}

func printHeaders(r *http.Response) {
	fmt.Printf("There are %v headers\n", len(r.Header))
	for k, v := range r.Header {
		fmt.Printf("%v:%v\n", k, v)
	}

}
