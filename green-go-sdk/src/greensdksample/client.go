package greensdksample

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DefaultClient struct {
	Profile Profile
}

func (defaultClient DefaultClient) GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) ([]byte, error) {
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)
	//purl, _ := url.Parse(config.MediaConfig.ProxyHttp)
	//proxy := http.ProxyURL(purl)
	client := &http.Client{
		Timeout:   time.Second * 15,
		Transport: &http.Transport{
			//Proxy: proxy,
		},
	}
	req, err := http.NewRequest(method, host+path+"?clientInfo="+url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		// handle error
		return nil, err
	} else {
		addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, defaultClient.Profile.AccessKeyId, defaultClient.Profile.AccessKeySecret)

		response, _ := client.Do(req)

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			// handle error
			return nil, err
		} else {
			return body, err
		}
	}
}

type IAliYunClient interface {
	GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) ([]byte, error)
}
