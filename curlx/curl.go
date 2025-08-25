/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2024-04-21 16:12:12
 * @Desc: 数据请求
 */
package curlx

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/goccy/go-json"
)

// AsyncResult 用于封装异步请求的结果
type AsyncResult struct {
	Data []byte
	Err  error
}

// PostAsync 异步发送 POST 请求，并通过返回的 channel 获取最终结果
func PostAsync(url string, data interface{}, headers ...map[string]string) <-chan AsyncResult {
	resultChan := make(chan AsyncResult, 1) // 使用缓冲1，防止 goroutine 阻塞
	go func() {
		// 超时时间：10秒
		client := &http.Client{Timeout: 10 * time.Second}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			resultChan <- AsyncResult{nil, err}
			return
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			resultChan <- AsyncResult{nil, err}
			return
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		// 添加额外 header
		for _, v := range headers {
			for kk, vv := range v {
				if kk != "" && vv != "" {
					req.Header.Set(kk, vv)
				}
			}
		}
		resp, err := client.Do(req)
		if err != nil {
			resultChan <- AsyncResult{nil, err}
			return
		}
		defer resp.Body.Close()
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			resultChan <- AsyncResult{nil, err}
			return
		}
		resultChan <- AsyncResult{result, nil}
	}()
	return resultChan
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string, headers ...map[string]string) (res []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, errors.New("Get request new error:" + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		for kk, vv := range v {
			if kk != "" && vv != "" {
				req.Header.Set(kk, vv)
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return res, errors.New("Get request server error:" + err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, errors.New("Get request read error:" + err.Error())
	}
	return result, nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, headers ...map[string]string) (res []byte, err error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Post json marshal error:" + err.Error())
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.New("Post new request error:" + err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		for kk, vv := range v {
			if kk != "" && vv != "" {
				req.Header.Set(kk, vv)
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Post request server error:" + err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Post request read error:" + err.Error())
	}
	return result, nil
}

func Request(method, url string, data interface{}, headers ...map[string]string) (res []byte, err error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Request json marshal error:" + err.Error())
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, errors.New("Request new request error:" + err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	for _, v := range headers {
		for kk, vv := range v {
			if kk != "" && vv != "" {
				req.Header.Set(kk, vv)
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Request server error:" + err.Error())
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Request read error:" + err.Error())
	}
	return result, nil
}

func PostForm(urls string, data map[string]string, headers ...map[string]string) (res []byte, err error) {
	client := &http.Client{Timeout: 10 * time.Second}

	postData := url.Values{}
	for i, v := range data {
		postData.Add(i, v)
	}

	req, err := http.NewRequest("POST", urls, strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, errors.New("request error")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range headers {
		for kk, vv := range v {
			if kk != "" && vv != "" {
				req.Header.Set(kk, vv)
			}
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("request server error")
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("request read error")
	}
	return result, nil
}
