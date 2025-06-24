package router

import (
	"github.com/gin-gonic/gin"
)

const (
	ProjectsPath = "/projects" // 项目路径
	ListPath     = "/list"     // 项目列表路径
	DescriptPath = "/descript" // 项目描述路径
	DataPath     = "/data"     // 项目数据路径
)

// ProjectsService 定义了网站的基本服务接口
type ProjectsService interface {
	List(*gin.Context)     // 获取项目列表
	Descript(*gin.Context) // 获取特定项目的详细信息
	Data(*gin.Context)     // 获取特定项目的数据
}

// BindProjects 将服务绑定至路由
func BindProjects(r *gin.Engine, sve ProjectsService) {
	v1 := r.Group(ProjectsPath)
	// 跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	{
		v1.POST(ListPath, sve.List)
		v1.POST(DescriptPath, sve.Descript)
		v1.POST(DataPath, sve.Data)
		v1.GET(DataPath, sve.Data)
	}
}

// SetupRouter 设置路由
func SetupRouter(sve ProjectsService) *gin.Engine {
	r := gin.Default()
	BindProjects(r, sve)
	return r
}
