package explancer

import (
	"io"
	"log/slog"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// CoventToHTML 将 markdown 转换为 HTML
func ConverToHTML(md io.Reader) ([]byte, error) {
	mddata, err := io.ReadAll(md)
	if err != nil {
		return nil, err
	}

	// Markdown parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(mddata)

	// HTML renderer
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer), nil
}

// ConverToHTMLFromFile 从文件中读取 markdown 并转换为 HTML
func ConverToHTMLFromFile(filename string) ([]byte, error) {
	slog.Error("filename", "filename", filename)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ConverToHTML(f)
}
