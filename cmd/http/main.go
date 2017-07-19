package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
)

func Get(url string) {
	//url+="?sort=desc"
	//url+="&room=xxx"
	//url+="&offset=0"
	//url+="&max=10"
	fmt.Println(url)

	http.Header{}.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 wechatdevtools/6.5.6 MicroMessenger/{{version}} Language/zh_CN webview/0")

	resp, err :=http.Get(url)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func MessageFetch(url string) {
	var param = &struct {
		Room   string `json:"room"`
		Offset int64  `json:"offset"`
		Limit  int    `json:"max"`
		Sort   string `json:"sort"`
	}{}
	param.Room=""
	param.Sort="desc"
	param.Offset=0
	param.Limit=10
	body, _ := json.Marshal(param)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	if err!=nil{
		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(resBody))
}

func main() {
	var url=""
	Get(url)
	fmt.Println("test")
}

type Message struct {
	Room          string `json:"room"`
	Type          int16  `json:"type"`
	Msg           string `json:"msg"`
	CorrelationID string `json:"correlation_id,omitempty"`
	Extra         string `json:"extra"`
}

func Post(at,url string) {
	msg:=&Message{
		Room:"",
		Type:257,
		Msg:"对的，是这样的",
		Extra:"{\"name\":\"xxx 024\"}",
	}
	body, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", at)
	if err!=nil{
		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println(string(resBody))
}

func PostWechat(url string) {
	msg:=&Message{
		Room:"",
		Type:257,
		Msg:"对的，是这样的",
		Extra:"{\"name\":\"xxx 024\"}",
	}
	body, _ := json.Marshal(msg)
	req, err := http.NewRequest("GET", url, strings.NewReader(string(body)))
	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; U; Android 2.3.6; zh-cn; GT-S5660 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MicroMessenger/4.5.255")
	//req.Header.Add("Referer", "https://mp.weixin.qq.com")
	if err != nil{

		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	fmt.Println("complete")

	fmt.Println(string(resBody))
}
