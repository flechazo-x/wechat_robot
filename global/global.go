package global

var (
	cfg *SysConfig
)

func GetCfg() *SysConfig {
	return cfg
}
