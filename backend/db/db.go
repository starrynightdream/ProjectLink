// 实现数据库相关功能，数据库功能
package db

import (
	"io"
	"os"

	"github.com/starrynightdream/ProjectLink/backend/explancer"
)

// 数据库连接应该实现的功能
type DBfunction interface {
	DumpYaml(io.Writer) error                       // 将数据库输出为 Yaml 文件
	LoadYaml(io.Reader) error                       // 从 Yaml 文件中构建数据库
	CreateProject(*explancer.ProjectMetaInfo) error // 添加一个项目
	UpdateProject(*explancer.ProjectMetaInfo) error // 更新一个项目
	DeleteProject(*explancer.ProjectMetaInfo) error // 删除一个项目

	// 读取相关函数，数量较多
	ReadAllProjects() ([]*explancer.ProjectMetaInfo, error)         // 读取所有项目
	ReadTagedProjects(string) ([]*explancer.ProjectMetaInfo, error) // 读取指定tag的项目
}

// DBDumpYamlFile 包装从文件名构建输出路径的方法
func DBDumpYamlFile(dbf DBfunction, fn string) error {
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0x666)
	if err != nil {
		return err
	}
	defer f.Close()

	return dbf.DumpYaml(f)
}

// DBLoadYamlFile 包装从文件名读取输入文件构造数据库的方法
func DBLoadYamlFile(dbf DBfunction, fn string) error {
	f, err := os.OpenFile(fn, os.O_RDONLY, 0x666)
	if err != nil {
		return err
	}
	defer f.Close()

	return dbf.LoadYaml(f)
}
