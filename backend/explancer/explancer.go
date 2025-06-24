// Description: 解释器，用于解释配置参数，保留修改信息的能力
package explancer

import (
	"encoding/json"
	"fmt"
)

// ExpplanceRegistry 解释器注册表
var ExpplanceRegistry = map[string]Explancer{}

// RegisterExplancer 注册解释器
func RegisterExplancer(explancer Explancer) {
	ExpplanceRegistry[explancer.Type()] = explancer
}

func init() {
	// 注册解释器
	RegisterExplancer(&URLExplance{})
	RegisterExplancer(&VedioExplance{})
}

// Explance 通过解释器类型解释参数
func Explance(infos []ProjectMetaInfo) error {
	for idx, info := range infos {
		explancer, ok := ExpplanceRegistry[info.Type]
		if !ok {
			return fmt.Errorf("项目 %s 未找到对应的解释器:%s", info.Name, info.Type)
		}
		err := explancer.Explain(&infos[idx])
		if err != nil {
			return fmt.Errorf("项目 %s 通过解释器 %s 解释参数失败:%w",
				info.Name, info.Type, err)
		}
	}
	return nil
}

// Explancer 定义了一个解释器的接口
type Explancer interface {
	Type() string                   // 对应的解释器类型
	Explain(*ProjectMetaInfo) error // 解析参数
}

// URLExplance 重定向解释器
type URLExplance struct{}

var _ Explancer = (*URLExplance)(nil)

// URLParams 重定向解释器参数
type URLParams struct {
	URLs []string `json:"URLs"`
}

// Type 返回解释器类型
func (rde *URLExplance) Type() string {
	return "URL"
}

// Explain 加载并解释参数
func (rde *URLExplance) Explain(n *ProjectMetaInfo) error {
	params := URLParams{}
	// 借助 json 实现替换
	jsonData, err := json.Marshal(n.Params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &params)
	n.FrontParams = params
	return err
}

// VedioExplance 视频解释器
type VedioExplance struct{}

var _ Explancer = (*VedioExplance)(nil)

// VedioParams 视频解释器参数
type VedioParams struct {
	Paths []string `json:"Paths"`
}

// Type 返回解释器类型
func (rde *VedioExplance) Type() string {
	return "Vedio"
}

// Explain 加载并解释参数
func (rde *VedioExplance) Explain(n *ProjectMetaInfo) error {
	params := VedioParams{}
	jsonData, err := json.Marshal(n.Params)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &params)
	n.FrontParams = params
	return err
}
