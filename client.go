// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dingding

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var UserAgent = "[fastwego/dingding] A fast dingding development sdk written in Golang"

/*
钉钉 api 服务器地址
*/
var DingdingServerUrl = "https://oapi.dingtalk.com"

/*
HttpClient 用于向接口发送请求
*/
type Client struct {
	Ctx *App
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {
	uri, err = client.applyAccessToken(uri)
	if err != nil {
		return
	}

	uri = DingdingServerUrl + uri
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("GET %s", uri)
	}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", UserAgent)
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	uri, err = client.applyAccessToken(uri)
	if err != nil {
		return
	}
	uri = DingdingServerUrl + uri
	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("POST %s", uri)
	}
	req, err := http.NewRequest(http.MethodPost, uri, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add("Content-Type", contentType)
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

/*
在请求地址上附加上 access_token
*/
func (client *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken, err := client.Ctx.AccessToken.GetAccessTokenHandler()
	if err != nil {
		return
	}

	parse, err := url.Parse(oldUrl)
	if err != nil {
		return
	}

	newUrl = parse.Host + parse.Path + "?access_token=" + accessToken
	if len(parse.RawQuery) > 0 {
		newUrl += "&" + parse.RawQuery
	}

	if len(parse.Fragment) > 0 {
		newUrl += "#" + parse.Fragment
	}
	return
}

/*
筛查钉钉 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if bytes.HasPrefix(resp, []byte(`{`)) { // 只针对 json
		errorResponse := struct {
			Errcode int64  `json:"errcode"`
			Errmsg  string `json:"errmsg"`
		}{}
		err = json.Unmarshal(resp, &errorResponse)
		if err != nil {
			return
		}

		if errorResponse.Errcode != 0 {
			err = errors.New(string(resp))
			return
		}
	}
	return
}
