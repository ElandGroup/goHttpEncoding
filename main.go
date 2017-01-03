package main

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/axgle/mahonia"
	"github.com/labstack/echo"
	"github.com/mozillazg/request"
)

func test1(url string) (bt []byte, err error) {
	testResp, err := http.Get(url)
	if err != nil {
		return
	}
	bt, err = ioutil.ReadAll(testResp.Body)
	return
}

func test2(url string) (bt []byte, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	bt, err = ioutil.ReadAll(resp.Body)
	return
}

func test3(url string) (bt []byte, err error) {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get(url)
	if err != nil {
		return
	}
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
	} else {
		reader = resp.Body
	}

	bt, err = ioutil.ReadAll(reader)
	return
}

var (
	e = echo.New()
)

func main() {
	e.GET("/", func(c echo.Context) error {
		url := "http://www.baidu.com"
		b3, _ := test3(url)
		b4 := string(b3)
		//fmt.Println(b4)

		//fmt.Println()
		//b5 := mahonia.NewDecoder("UTF-8").ConvertString(b4)
		return c.HTML(http.StatusOK, b4)
	})
	e.Start(":1002")

}

func EncodingAlUrl(response string) string {
	c := mahonia.NewDecoder("gbk").ConvertString(response)
	return mahonia.NewEncoder("UTF-8").ConvertString(c)
}
