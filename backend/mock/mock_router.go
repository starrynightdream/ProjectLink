package mock

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starrynightdream/ProjectLink/backend/explancer"
	"github.com/starrynightdream/ProjectLink/backend/router"
)

const (
	MockYamlFile = "mock_info.yaml"
)

// MockDescriptReq 为 Mock 描述请求
type MockDescriptReq struct {
	Target string `json:"target"`
}

// MockDataReq 为 Mock 数据请求
type MockDataReq struct {
	Target string `json:"target"`
}

// MockService 实现一个桩服务，让前端可以启动测试
type MockService struct {
	MockDescriptRoot string // 描述文件根目录
	MockDataRoot     string // 数据文件根目录
	MockRoot         string // Mock 定位目录
}

var _ router.ProjectsService = (*MockService)(nil)

// MockServiceOption 为 Mock 服务提供选项
type MockServiceOption func(*MockService)

// NewMockService 创建一个 Mock 服务
func NewMockService(opts ...MockServiceOption) *MockService {
	s := &MockService{
		MockDescriptRoot: "./mock/descript",
		MockDataRoot:     "./mock/static",
		MockRoot:         "./mock",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// WithMockDescriptRoot 设置描述文件根目录
func WithMockDescriptRoot(root string) MockServiceOption {
	return func(s *MockService) {
		s.MockDescriptRoot = root
	}
}

// WithMockDataRoot 设置数据文件根目录
func WithMockDataRoot(root string) MockServiceOption {
	return func(s *MockService) {
		s.MockDataRoot = root
	}
}

// WithMockRoot 设置 Mock 定位目录
func WithMockRoot(root string) MockServiceOption {
	return func(s *MockService) {
		s.MockRoot = root
	}
}

func (s *MockService) List(c *gin.Context) {
	infos, err := explancer.DecoderProjectMetaInfosFromFile(
		filepath.Join(s.MockRoot, MockYamlFile))
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	err = explancer.Explance(infos)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// 避免内部数据泄露
	for idx := range infos {
		infos[idx].Params = nil
	}
	c.JSON(200, infos)
}

func (s *MockService) Descript(c *gin.Context) {
	mockDescriptreq := MockDescriptReq{}
	if err := c.BindJSON(&mockDescriptreq); err != nil {
		slog.Error("解析请求失败", "err", err)
		c.Status(http.StatusBadRequest)
		return
	}

	html_data, err := explancer.ConverToHTMLFromFile(
		filepath.Join(s.MockDescriptRoot, mockDescriptreq.Target))

	if err != nil {
		slog.Error("HTML转换失败", "err", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.String(200, string(html_data))
}

func (s *MockService) Data(c *gin.Context) {
	var fileName string

	// 根据请求方法处理输入
	switch c.Request.Method {
	case http.MethodGet:
		fileName = c.Query("target") // 从URL参数获取 ?target=xxx
	case http.MethodPost:
		var req MockDataReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		fileName = req.Target
	default:
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	// 验证文件名（防止路径遍历攻击）
	if fileName == "" || strings.Contains(fileName, "..") || strings.Contains(fileName, "/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filename"})
		return
	}

	// 安全拼接路径
	filePath := filepath.Join(s.MockDataRoot, filepath.Base(fileName))

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.Status(http.StatusNotFound)
		return
	}

	c.File(filePath)
}
