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

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string, headers ...map[string]string) (res []byte, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
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
		return res, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, errors.New("读取错误")
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
	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Request(method, url string, data interface{}, headers ...map[string]string) (res []byte, err error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr, _ := json.Marshal(data)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
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
