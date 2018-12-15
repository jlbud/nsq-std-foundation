package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	HttpResultSuccess = http.StatusOK
)

//返回的结构
type HttpResult struct {
	StateCode int
	Body      string
}

func SendGet(url string, rp *HttpResult) error {
	log.Printf("[nsq request method:]" + http.MethodGet)
	log.Printf("[nsq request url:]" + url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	return createResult(res, rp)
}

// postData     post方法要发送的数据
// url          请求的路由
func SendPost(postData interface{}, url string, rp *HttpResult) error {
	log.Printf("[nsq request method:]" + http.MethodPost)
	log.Printf("[nsq request url:]" + url)
	data := []byte{}
	if postData != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false) //取消特殊字符转义
		err := encoder.Encode(postData)
		if err != nil {
			return err
		}
		data = buffer.Bytes()
	}
	log.Printf("[nsq postData is:] %v", string(data))
	bodyReader := bytes.NewBuffer(data)
	res, err := http.Post(url, "application/json; charset=utf-8", bodyReader)
	if err != nil {
		return err
	}
	return createResult(res, rp)
}

// postData     post方法要发送的数据
// url          请求的路由
// resu         请求反馈
func SendPostForm(postForm map[string]string, postUrl string, resu *HttpResult) error {
	fmt.Println("[Request method:]" + http.MethodPost)
	fmt.Println("[Request url:]" + postUrl)
	val := url.Values{}
	for k, v := range postForm {
		val.Set(k, v)
	}
	res, err := http.PostForm(postUrl, val)
	if err != nil {
		return err
	}
	return createResult(res, resu)
}

func createResult(r *http.Response, hrt *HttpResult) error {
	r.Close = true
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Printf("[nsq http response status:] %v", r.Status)
	log.Printf("[nsq response data:] %v", string(body))

	if string(body) != "" {
		hrt.Body = string(body)
	}
	hrt.StateCode = r.StatusCode
	return nil
}
