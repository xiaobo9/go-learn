package http

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var myUrl = "http://www.baidu.com" //全局变量链接
func Http() {
	HttpGet()
}

func HttpGet() {
	response, err := http.Get(myUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	log.Println(string(body))
}

func HttpPost() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				// 连接超时
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				// 发送接受数据超时
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	reqest, err := http.NewRequest("POST", myUrl, strings.NewReader("name=PostName"))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}

	response, err := client.Do(reqest)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
	status := response.StatusCode
	body, _ := ioutil.ReadAll(response.Body)
	log.Println(string(body))
	log.Println(status)

}

func HttpPost2() {
	req, err := http.Post(myUrl, "application/x-www-form-urlencoded", strings.NewReader("name=myname"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(body))
}

func PostForm() {
	req, err := http.PostForm(myUrl, url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(body))
}
