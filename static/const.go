package static

const (
	ImageURL       = "https://img.qianxiaoduan.com/"                                      //图片生成网址
	ImageSourceURL = "https://imgegg.qianxiaoduan.com/wallpaper?offset=%d&limit=3&type=1" //图片生成请求网址
	Redis          = "Resis"
	TianApi        = "https://apis.tianapi.com/%s/index?key=%s"
)

type Message string

const (
	HotSearch Message = "networkhot" // 励志古言
	PyqWenAn  Message = "pyqwenan"   // 朋友圈文案
	CaiHongPi Message = "caihongpi"  //彩虹屁
	ZaoAn     Message = "zaoan"      //早安日记
	DuJiTang  Message = "dujitang"   //舔狗语句
	LiShi     Message = "lishi"      //舔狗语句
	Everyday  Message = "everyday"   //每日英语
)
