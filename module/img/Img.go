package img

import (
	"encoding/json"
	"fmt"
	"github.com/flechazo-x/go_common/middleware/xhttp"
	"github.com/flechazo-x/go_common/util/timex"
	"net/http"
	"time"
	"wechat_robot/static"
)

// SourceResp 图片源响应
type SourceResp struct {
	Success bool     `json:"Success,omitempty"`
	Code    int      `json:"Code,omitempty"`
	Message string   `json:"Message,omitempty"`
	Result  struct { //结果
		Count int `json:"Count,omitempty"`
		Rows  []struct {
			Id        string      `json:"Id,omitempty"`
			Thumbnail string      `json:"Thumbnail,omitempty"`
			Url       string      `json:"Url,omitempty"`
			Hot       int         `json:"Hot,omitempty"`
			Width     int         `json:"Width,omitempty"`
			Height    int         `json:"Height,omitempty"`
			Name      interface{} `json:"Name,omitempty"`
			Type      string      `json:"Type,omitempty"`
			Scale     string      `json:"Scale,omitempty"`
			Tag       string      `json:"Tag,omitempty"`
			CreatedAt string      `json:"CreatedAt,omitempty"`
			UpdatedAt string      `json:"UpdatedAt,omitempty"`
		} `json:"Rows,omitempty"`
	} `json:"Result"`
}

// GetImage 获取图片
func GetImage() (string, error) {
	var (
		imgSourceResp = new(SourceResp)
		imgURL        string
		err           error
	)
	surl := fmt.Sprintf(static.ImageSourceURL, timex.GenRandNum(10000))
	o := &xhttp.Option{
		Method:      http.MethodGet,
		URL:         surl,
		ContentType: xhttp.ContentTypeIsForm,
		Timeout:     30 * time.Second,
	}
	respBody, err := xhttp.Ask(nil, o)
	if err != nil {
		return imgURL, err
	}
	err = json.Unmarshal(respBody, &imgSourceResp)
	if err != nil {
		return imgURL, err
	}

	if len(imgSourceResp.Result.Rows) == 0 {
		return imgURL, fmt.Errorf("GetImage Result.Rows len is 0 error:%s, ", surl)
	}
	return fmt.Sprintf("%s%s", static.ImageURL, imgSourceResp.Result.Rows[0].Url), nil
}
