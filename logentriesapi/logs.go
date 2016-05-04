package logentriesapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Log struct {
	*APIObject
	CreatedUnixNano int64 `json:"created"`
	Name            string
	Filename        string
	Follow          string
	Key             string
	Token           string
	Retention       int
	LogType         string `json:"type"`
	IsComplete      bool
}

func (l *Log) CreatedTime() time.Time {
	return time.Unix(l.CreatedUnixNano/1000, l.CreatedUnixNano%1000)
}

type createLogResponse struct {
	*APIResponse
	Log *Log
}

func (c *Client) CreateLog(hostKey string, name string) (log *Log, err error) {
	params := url.Values{}
	params.Add("request", "new_log")
	params.Add("user_key", c.accountKey)
	params.Add("host_key", hostKey)
	params.Add("name", name)
	params.Add("retention", "-1")
	params.Add("type", "")
	params.Add("source", "token")

	var rawResp *http.Response
	var bytes []byte
	var resp createLogResponse
	if rawResp, err = http.PostForm(c.baseURL.String(), params); err == nil {
		if bytes, err = ioutil.ReadAll(rawResp.Body); err == nil {
			err = json.Unmarshal(bytes, &resp)
		}
	}
	if err != nil {
		return
	}

	if err = resp.AssertOk(); err != nil {
		return
	}

	log = resp.Log
	return
}

type getLogResponse struct {
	*APIResponse
	*Log
}

func (c *Client) GetLog(logKey string) (log *Log, err error) {
	var resp getLogResponse

	var data []byte
	if data, err = c.getBytes("/logs/" + logKey); err == nil {
		err = json.Unmarshal(data, &resp)
	}
	if err != nil {
		return
	}

	if err = resp.AssertOk(); err != nil {
		return
	}
	if err = resp.AssertType("log"); err != nil {
		return
	}

	log = resp.Log
	return
}
