package global

import (
	"github.com/flechazo-x/go_common/db/redisx/redispool"
	"github.com/flechazo-x/go_common/logx"
	"github.com/flechazo-x/go_common/util/filex"
)

// SysConfig 系统配置
type SysConfig struct {
	Redis        *redispool.RedisConfig `json:"Redis,omitempty"`        //redis配置
	LogConfigSrc string                 `json:"LogConfigSrc,omitempty"` //日志目录
	App          *App                   `json:"app" yaml:"app"`         //app对应环境
	Keys         *Keys
}

func NweSysConfig() *SysConfig {
	return new(SysConfig)
}

// InitSysConfig 初始化系统配置
func InitSysConfig() {
	cfg = NweSysConfig()
	if err := filex.Deserialize("./config/config.json", cfg); err != nil {
		logx.LFatal("[InitSysConfig] Deserialize Error: ", err)
		return
	}
}

// App .
type App struct {
	Env string `json:"Env,omitempty"`
}

// Keys api
type Keys struct {
	ChristmasHatURL    string `json:"ChristmasHatURL"`
	BotName            string `json:"BotName"`       //机器人名称
	WeatherKey         string `json:"WeatherKey"`    //高德天气 api key
	TianapiKey         string `json:"TianapiKey"`    //天行 api key
	TianapiKey1        string `json:"TianapiKey1"`   //天行 api key
	HoneyLove          string `json:"HoneyLove"`     //和风天气 api key
	LoverChName        string `json:"LoverChName"`   //恋人名字
	LoverBirthday      string `json:"LoverBirthday"` //恋人生日
	MasterAccount      string `json:"MasterAccount"`
	HouchangcunFans    string `json:"HouchangcunFans"`
	BanzhuanGroup      string `json:"BanzhuanGroup"`
	BubeiGroup         string `json:"BubeiGroup"`
	QweatherKey        string `json:"QweatherKey"`
	BubeiStartDate     string `json:"BubeiStartDate"`
	WuZhuangShiMembers string `json:"WuZhuangShiMembers"`
	RemindMsg          string `json:"RemindMsg"`
}
