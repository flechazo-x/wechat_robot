package message

import (
	"encoding/json"
	"fmt"
	"github.com/flechazo-x/go_common/middleware/xhttp"
	"github.com/pkg/errors"
	"net/http"
	"time"
	"wechat_robot/global"
	"wechat_robot/static"
)

type (
	Info struct {
		Code   int    `json:"Code,omitempty"`
		Msg    string `json:"Msg,omitempty"`
		Result struct {
			Curpage int `json:"Curpage,omitempty"`
			Allnum  int `json:"Allnum,omitempty"`
			List    []struct {
				Id          int    `json:"Id,omitempty"`
				Ctime       string `json:"Ctime,omitempty"`
				Title       string `json:"Title,omitempty"`
				Description string `json:"Description,omitempty"`
				PicUrl      string `json:"PicUrl,omitempty"`
				Url         string `json:"Url,omitempty"`
				Type        int    `json:"Type,omitempty"`
				DetaType    string `json:"DetaType,omitempty"`
				Digest      string `json:"Digest,omitempty"`
				Hotnum      int    `json:"Hotnum,omitempty"`
				Lsdate      string `json:"Lsdate,omitempty"`
			} `json:"List,omitempty"`
		}
	}
	OneInfo struct {
		Code   int    `json:"Code,omitempty"`
		Msg    string `json:"Msg,omitempty"`
		Result struct {
			Id       int32  `json:"Id,omitempty"`
			Note     string `json:"Note,omitempty"`    //释义
			Reply    string `json:"Reply,omitempty"`   //	回答内容
			Content  string `json:"Content,omitempty"` //内容
			Source   string `json:"Source,omitempty"`  //来源出处
			Datatype string `json:"Datatype,omitempty"`
			Date     string `json:"Date,omitempty"`
		}
	}
)

func GetMessage(stype static.Message, word ...string) (string, error) {
	var (
		surl     string
		respBody []byte
		info     = new(Info)
		oneInfo  = new(OneInfo)
		err      error
		message  string
		code     int
	)

	surl = fmt.Sprintf(static.TianApi, stype, global.GetCfg().Keys.TianapiKey)
	o := &xhttp.Option{
		Method:      http.MethodGet,
		URL:         surl,
		ContentType: xhttp.ContentTypeIsForm,
		Timeout:     30 * time.Second,
	}
	respBody, err = xhttp.Ask(nil, o)
	if err != nil {
		return "", fmt.Errorf("GetMessage http Ask Err:%s", err.Error())
	}
	if checkOneInfo(stype) {
		if err = json.Unmarshal(respBody, oneInfo); err != nil {
			return "", errors.Wrapf(err, "GetMessage http [oneInfo] Unmarshal err:%s", err.Error())
		}
		code = oneInfo.Code
	} else {
		if err = json.Unmarshal(respBody, info); err != nil {
			return "", errors.Wrapf(err, "GetMessage http [Info] Unmarshal err:%s", err.Error())
		}
		code = info.Code
	}

	if code != http.StatusOK {
		return message, errors.New("GetMessage http Code not is  200")
	}

	if stype == static.HotSearch {
		if len(info.Result.List) > 0 {
			for i, v := range info.Result.List {
				message += fmt.Sprintf("热搜榜指数:%d\n标题:%s\n内容：%s\n", info.Result.List[i].Hotnum, info.Result.List[i].Title, v.Digest)
			}
		}
		return message, nil
	}
	if stype == static.LiShi {
		if len(info.Result.List) > 0 {
			for i, _ := range info.Result.List {
				message += fmt.Sprintf("日期:%s\n内容：%s\n", info.Result.List[i].Lsdate, info.Result.List[i].Title)
			}
		}
		return message, nil
	}

	if stype == static.Everyday {
		return fmt.Sprintf("今日时间:%s\n句子内容:%s\n句子解释:%s\n", oneInfo.Result.Date, oneInfo.Result.Content, oneInfo.Result.Note), nil
	}
	//单个结果集时返回
	return oneInfo.Result.Content, nil

	//return message, fmt.Errorf("GetMessage http content empty")
}

// 检测是否是单个
func checkOneInfo(stype static.Message) bool {
	return stype == static.PyqWenAn || stype == static.CaiHongPi || stype == static.ZaoAn || stype == static.DuJiTang || stype == static.Everyday
}
