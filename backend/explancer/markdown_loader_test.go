package explancer

import (
	"log/slog"
	"testing"
)

func TestLoadMarkdown(t *testing.T) {
	filename := "../mock/descript/mock_info.md"
	infos, err := ConverToHTMLFromFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("解析的HTML：", "html=\n", infos)
}
