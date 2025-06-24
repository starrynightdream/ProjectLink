package mock

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/starrynightdream/ProjectLink/backend/explancer"
	"github.com/starrynightdream/ProjectLink/backend/router"
)

func TestMockRouter(t *testing.T) {
	r := router.SetupRouter(NewMockService(
		WithMockRoot("."),
		WithMockDescriptRoot("descript"),
		WithMockDataRoot("static")))
	w := httptest.NewRecorder()

	// List 测试
	req, err := http.NewRequest("POST", "/projects/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("期望状态码200，实际为%d", w.Code)
	}

	result := []explancer.ProjectMetaInfo{}
	err = json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("解析失败 %v", w.Body.String())
	}
	if len(result) != 2 {
		t.Fatalf("期望解析出2个项目，实际为%d", len(result))
	}

	if result[0].FrontParams == nil {
		t.Fatalf("解释器解释出的参数为空 %v", result[0])
	}

	// slog.Info("测试通过", "w.body", w.Body.String())
	// slog.Info("测试通过", "result", result)

	// Descript 测试
	descriptReq := MockDescriptReq{
		Target: "mock_info.md",
	}
	descJson, err := json.Marshal(descriptReq)
	if err != nil {
		t.Fatalf("序列化失败 %v", descriptReq)
	}
	req, err = http.NewRequest("POST", "/projects/descript", bytes.NewReader(descJson))
	if err != nil {
		t.Fatalf("创建请求失败 %v", descriptReq)
	}
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("期望状态码200，实际为%d", w.Code)
	}
	// slog.Info("测试通过", "w.body=", w.Body.String())

	// Data 测试
	dataReq := MockDataReq{
		Target: "never.gif",
	}
	dataJson, err := json.Marshal(dataReq)
	if err != nil {
		t.Fatalf("序列化失败 %v", dataReq)
	}
	req, err = http.NewRequest("POST", "/projects/data", bytes.NewReader(dataJson))
	if err != nil {
		t.Fatalf("创建请求失败 %v", dataReq)
	}
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("期望状态码200，实际为%d", w.Code)
	}
}
