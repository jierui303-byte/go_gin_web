package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string    `json:"app_name"`
	APPMode string    `json:"app_mode"`
	APPHost string    `json:"app_host"`
	APPPort string    `json:"app_port"`
	Sms     SmsConfig `json:"sms"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_srcret"`
}

var _cfg *Config = nil

//获取全局config对象
func GetConfig() *Config {
	return _cfg
}

//自定义的加载配置文件内容的函数
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
