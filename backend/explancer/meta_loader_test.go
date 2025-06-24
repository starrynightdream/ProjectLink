package explancer

import (
	"log"
	"testing"
)

const RefMockFilePath = "../mock/mock_info.yaml"

func TestLoadFile(t *testing.T) {
	infos, err := DecoderProjectMetaInfosFromFile(RefMockFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(infos) != 2 {
		t.Fatalf("期望解析出2个项目，实际为%d", len(infos))
	}
	log.Printf("解析出的项目信息：%v", infos)
	log.Printf("参数获取情况：%v", infos[0].Params)
}
