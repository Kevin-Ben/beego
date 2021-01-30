package wechat

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
)

func GetAccount() *officialaccount.OfficialAccount {
	//wechat
	AppID, _ := beego.AppConfig.String("AppID")
	AppSecret, _ := beego.AppConfig.String("AppSecret")
	Token, _ := beego.AppConfig.String("Token")
	EncodingAESKey, _ := beego.AppConfig.String("EncodingAESKey")

	wc := wechat.NewWechat()
	redis := GetRedisCache()

	cfg := &config.Config{
		AppID:          AppID,
		AppSecret:      AppSecret,
		Token:          Token,
		EncodingAESKey: EncodingAESKey,
		Cache:          redis,
	}
	return wc.GetOfficialAccount(cfg)
}

func GetRedisCache() *cache.Redis {
	redisHost, _ := beego.AppConfig.String("RedisHost")
	RedisDB, _ := beego.AppConfig.Int("RedisDB")
	return cache.NewRedis(&cache.RedisOpts{Host: redisHost, Database: RedisDB})
}
