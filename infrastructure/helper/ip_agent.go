package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Duration(3 * time.Second),
}

type IpInfo struct {
	Ip        string `json:"ip"`
	CountryId string `json:"country_id"`
	Country   string `json:"country"`
	AreaId    string `json:"area_id"`
	Area      string `json:"area"`
	RegionId  string `json:"region_id"`
	Region    string `json:"region"`
	CityId    string `json:"city_id"`
	City      string `json:"city"`
	County_id string `json:"county_id"`
	County    string `json:"county"`
	IspId     string `json:"isp_id"`
	Isp       string `json:"isp"`
}
type IpInfoRes struct {
	Code int    `json:"code"`
	Data IpInfo `json:"data"`
}

func GetIpInfo(ip string) (IpInfo, error) {
	var ipInfoRes IpInfoRes
	var err error

	uri := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip

	resp, err := client.Get(uri)
	if err != nil {
		//logs.Error("get %s, error = %v", uri, err)
		return ipInfoRes.Data, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//logs.Error("get %s, status = %d", uri, resp.StatusCode)
		return ipInfoRes.Data, fmt.Errorf("status = %d", resp.StatusCode)
	}

	//logs.Info("uri = %s, done", uri)
	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, &ipInfoRes)
	return ipInfoRes.Data, err
}
