//天气

package qweather

import (
	"encoding/json"
	"fmt"
	"github.com/flechazo-x/go_common/middleware/xhttp"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"wechat_robot/global"

	"github.com/pkg/errors"
)

const QWeatherHOST = "https://geoapi.qweather.com"

type QWeatherResp struct {
	Code     string `json:"Code,omitempty"`
	Location []struct {
		Name      string `json:"Name,omitempty"`
		Id        string `json:"Id,omitempty"`
		Lat       string `json:"Lat,omitempty"`
		Lon       string `json:"Lon,omitempty"`
		Adm2      string `json:"Adm2,omitempty"`
		Adm1      string `json:"Adm1,omitempty"`
		Country   string `json:"Country,omitempty"`
		Tz        string `json:"Tz,omitempty"`
		UtcOffset string `json:"UtcOffset,omitempty"`
		IsDst     string `json:"IsDst,omitempty"`
		Type      string `json:"Type,omitempty"`
		Rank      string `json:"Rank,omitempty"`
		FxLink    string `json:"FxLink,omitempty"`
	} `json:"Location,omitempty"`
	Refer struct {
		Sources []string `json:"Sources,omitempty"`
		License []string `json:"License,omitempty"`
	} `json:"Refer"`
}

// GetLocationID 根据位置获取地理id
func GetLocationID(cityName string) (id string, err error) {
	var (
		surl         string
		respBody     []byte
		params       = url.Values{}
		qweatherResp = new(QWeatherResp)
	)

	params.Set("location", cityName)
	params.Set("key", global.GetCfg().Keys.QweatherKey)
	params.Set("range", "cn")
	//params.Set("adm", cityName)

	surl = fmt.Sprintf("%s/v2/city/lookup?%s", QWeatherHOST, params.Encode())
	respBody, err = xhttp.Ask(nil, &xhttp.Option{Method: http.MethodGet, URL: surl, ContentType: xhttp.ContentTypeIsForm, Timeout: 30 * time.Second})
	if err != nil {
		return "", errors.Wrapf(err, "GetLocationID HTTP error")
	}
	err = json.Unmarshal(respBody, &qweatherResp)
	if err != nil {
		return "", errors.Wrapf(err, "GetLocationID json.Unmarshal error")
	}

	if qweatherResp.Code != strconv.Itoa(http.StatusOK) {
		return "", fmt.Errorf("GetLocationID qweatherResp.Code error: %s", qweatherResp.Code)
	}

	if len(qweatherResp.Location) == 0 {
		return "", fmt.Errorf("GetLocationID qweatherResp.Location empyt")
	}

	// 匹配
	for _, v := range qweatherResp.Location {
		if v.Name == cityName {
			return v.Id, nil
		}
		if v.Adm2 == cityName {
			return v.Id, nil
		}
		if strings.Contains(v.Adm1, cityName) {
			return v.Id, nil
		}
	}

	return "", fmt.Errorf("GetLocationID not found")
}

// QWeatherDetailResp 天气返回
type QWeatherDetailResp struct {
	Code       string `json:"Code,omitempty"`
	UpdateTime string `json:"UpdateTime,omitempty"`
	FxLink     string `json:"FxLink,omitempty"`
	Now        struct {
		ObsTime   string `json:"ObsTime,omitempty"`
		Temp      string `json:"Temp,omitempty"`
		FeelsLike string `json:"FeelsLike,omitempty"`
		Icon      string `json:"Icon,omitempty"`
		Text      string `json:"Text,omitempty"`
		Wind360   string `json:"Wind360,omitempty"`
		WindDir   string `json:"WindDir,omitempty"`
		WindScale string `json:"WindScale,omitempty"`
		WindSpeed string `json:"WindSpeed,omitempty"`
		Humidity  string `json:"Humidity,omitempty"`
		Precip    string `json:"Precip,omitempty"`
		Pressure  string `json:"Pressure,omitempty"`
		Vis       string `json:"Vis,omitempty"`
		Cloud     string `json:"Cloud,omitempty"`
		Dew       string `json:"Dew,omitempty"`
	} `json:"Now"`
	Refer struct {
		Sources []string `json:"Sources,omitempty"`
		License []string `json:"License,omitempty"`
	} `json:"Refer"`
}

// GetQWeatherDetail 获取天气详情
func GetQWeatherDetail(cityID, cityName string) (detail string, err error) {
	var (
		surl               string
		respBody           []byte
		params             = url.Values{}
		qweatherDetailResp = new(QWeatherDetailResp)
	)

	params.Set("location", cityID)
	params.Set("key", global.GetCfg().Keys.QweatherKey)

	surl = fmt.Sprintf("https://devapi.qweather.com/v7/weather/now?%s", params.Encode())
	respBody, err = xhttp.Ask(nil, &xhttp.Option{Method: http.MethodGet, URL: surl, ContentType: xhttp.ContentTypeIsForm, Timeout: 30 * time.Second})

	if err != nil {
		return "", errors.Wrapf(err, "GetQWeatherDetail HTTP error")
	}

	err = json.Unmarshal(respBody, &qweatherDetailResp)
	if err != nil {
		return "", errors.Wrapf(err, "GetQWeatherDetail .Unmarshal error")
	}

	if qweatherDetailResp.Code != strconv.Itoa(http.StatusOK) {
		return "", fmt.Errorf("GetQWeatherDetail qweatherResp.Code error: %s", qweatherDetailResp.Code)
	}

	detail = fmt.Sprintf(`%s今天天气,温度 %s 度,%s,%s %s 级,相对湿度 %s`, cityName,
		qweatherDetailResp.Now.Temp, qweatherDetailResp.Now.Text,
		qweatherDetailResp.Now.WindDir, qweatherDetailResp.Now.WindScale,
		qweatherDetailResp.Now.Humidity,
	)

	return detail, nil
}
