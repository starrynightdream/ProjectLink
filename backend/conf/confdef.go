// Server:
//   Port: 10086
//   HOST: localhost
//   FourceMock: false

package conf

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// EnvConfig 配置文件定义
type EnvConfig struct {
	ServerConf ServerConfig `yaml:"Server"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port       int    `yaml:"Port"`       // 端口
	Host       string `yaml:"Host"`       // 主机名
	FourceMock bool   `yaml:"FourceMock"` // 强制mock
}

// DecodeEnvConfig 从输入源读取数据
func DecodeEnvConfig(reader io.Reader) (*EnvConfig, error) {
	ec := &EnvConfig{}
	dec := yaml.NewDecoder(reader)
	err := dec.Decode(ec)
	return ec, err
}

// DecodeEnvConfigFromFile 从文件读取配置
func DecodeEnvConfigFromFile(p string) (*EnvConfig, error) {
	f, err := os.OpenFile(p, os.O_RDONLY, 0x444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return DecodeEnvConfig(f)
}
