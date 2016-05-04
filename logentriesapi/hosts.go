package logentriesapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Host struct {
	*APIObject
	Name            string `json:"name"`
	CreatedUnixNano int64  `json:"c"`
	DistName        string `json:"distname"`
	DistVer         string `json:"distver"`
	Hostname        string `json:"hostname"`
	Key             string `json:"key"`

	//Note that on a GetLog or CreateHost, the Log objects here encompossed are complete,
	//but for some reason under GetHost and GetHosts, everything but the Key will be empty :(
	//
	// in general you should probably GetLog each if you care about the details
	Logs []Log `json:"logs"`
}

func (h *Host) CreatedTime() time.Time {
	return time.Unix(h.CreatedUnixNano/1000, h.CreatedUnixNano%1000)
}

type getHostsResponse struct {
	*APIResponse
	*APIObject
	Hosts *[]Host `json:"list"`
}

func (c *Client) GetHosts() (hosts *[]Host, err error) {
	var resp getHostsResponse

	var data []byte
	if data, err = c.getBytes("/hosts"); err == nil {
		err = json.Unmarshal(data, &resp)
	}
	if err != nil {
		return
	}

	if err = resp.AssertOk(); err != nil {
		return
	}
	if err = resp.AssertType("hostlist"); err != nil {
		return
	}

	hosts = resp.Hosts
	return
}

type createHostResponse struct {
	*APIResponse
	Host *Host
}

func (c *Client) CreateHost(name string) (host *Host, err error) {
	params := url.Values{}
	params.Add("request", "register")
	params.Add("user_key", c.accountKey)
	params.Add("name", name)
	params.Add("system", "")
	params.Add("distver", "")
	params.Add("distname", "")

	var rawResp *http.Response
	var bytes []byte
	var resp createHostResponse
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

	host = resp.Host
	return
}

type getHostResponse struct {
	*APIResponse
	*Host
}

func (c *Client) GetHost(hostKey string) (host *Host, err error) {
	var data []byte
	var resp getHostResponse
	if data, err = c.getBytes("/hosts/" + hostKey); err == nil {
		err = json.Unmarshal(data, &resp)
	}
	if err != nil {
		return
	}

	if err = resp.AssertOk(); err != nil {
		return
	}
	if err = resp.AssertType("host"); err != nil {
		return
	}

	host = resp.Host
	return
}
