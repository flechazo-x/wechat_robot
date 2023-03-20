package initapp

import (
	_redisx "github.com/flechazo-x/go_common/db/redisx"
	"github.com/flechazo-x/go_common/db/redisx/redispool"
	"github.com/flechazo-x/go_common/logx"
	"wechat_robot/global"
	"wechat_robot/static"
)

// InitRedis 初始化redis
func InitRedis() {
	err := redispool.GetRedisPoolManager().InitPoolsWithConfigs(map[string]*redispool.RedisConfig{static.Redis: global.GetCfg().Redis})
	if err != nil {
		logx.LFatalf("[InitApp] InitPoolsWithConfigs err :%s", err.Error())
		return
	}

	_redisx.RDB, err = redispool.GetRedisPoolManager().GetConn(static.Redis)

	if err != nil {
		logx.LFatalf("[InitApp] GetRedisPoolManager err :%s", err.Error())
		return
	}

	if err = _redisx.RDB.Ping(); err != nil {
		logx.LFatalf("[InitApp] Ping err :%s", err.Error())
		return
	}
	return
}
