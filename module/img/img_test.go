package img

import (
	"fmt"
	"github.com/flechazo-x/go_common/middleware/xhttp"
	"github.com/flechazo-x/go_common/util/timex"
	"net/http"
	"testing"
	"time"
	"wechat_robot/static"
)

func TestImg(t *testing.T) {
	surl := fmt.Sprintf(static.ImageSourceURL, timex.GenRandNum(10000))
	o := &xhttp.Option{
		Method:      http.MethodGet,
		URL:         surl,
		ContentType: xhttp.ContentTypeIsForm,
		Body:        nil,
		Timeout:     30 * time.Second,
	}
	b, _ := xhttp.Ask(nil, o)
	t.Log(string(b))
}

func Test(t *testing.T) {
	t.Log(GetImage())
}
