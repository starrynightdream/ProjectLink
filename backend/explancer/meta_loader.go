package explancer

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ProjectMetaInfo 项目的元信息
type ProjectMetaInfo struct {
	Name        string              `yaml:"Name" json:"Name"`               // 项目名称
	Type        string              `yaml:"Type" json:"Type"`               // 项目展示类型
	Tags        []string            `yaml:"Tags" json:"Tags"`               // 项目标签
	Description ProjectDescriptInfo `yaml:"Description" json:"Description"` // 项目描述
	FrontParams any                 `json:"FrontParams"`                    // 返回前端的相关参数
	Params      map[string]any      `yaml:"Params"`                         // 项目参数
}

// 展示相关的信息
type ProjectDescriptInfo struct {
	MDFilePath     string `yaml:"MDFilePath" json:"MDFilePath"`         // 对应的 markdown 描述文件
	PreviewImgPath string `yaml:"PreviewImgPath" json:"PreviewImgPath"` // 对应的预览图文件
}

// DecoderProjectMetaInfo 解码项目元信息
func DecoderProjectMetaInfos(ipt io.Reader) ([]ProjectMetaInfo, error) {
	pmi := []ProjectMetaInfo{}
	decoder := yaml.NewDecoder(ipt)
	err := decoder.Decode(&pmi)
	return pmi, err
}

// DecoderProjectMetaInfosFromFile 从文件解码项目元信息
func DecoderProjectMetaInfosFromFile(filename string) ([]ProjectMetaInfo, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return DecoderProjectMetaInfos(f)
}
