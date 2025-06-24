package conf

import (
	"log"
	"testing"
)

const confFilepath = "env.yaml"

func TestLoadFile(t *testing.T) {
	envConfig, err := DecodeEnvConfigFromFile(confFilepath)
	if err != nil {
		t.Fatalf("加载配置文件失败:%v", err)
	}
	log.Println(envConfig) // 可能修改内容，所以不做内容判断
}
