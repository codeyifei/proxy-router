package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

type RootHandler struct {
	StaticPath string
	IndexPath  string
}

func NewRootHandler(staticPath, indexPath string) *RootHandler {
	return &RootHandler{
		StaticPath: staticPath,
		IndexPath:  indexPath,
	}
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 路径的绝对路径
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// 如果获取失败，返回 400 响应
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 在 URL 路径前加上静态资源根目录
	path = filepath.Join(h.StaticPath, path)

	// 检查对应资源文件是否存在
	switch _, err = os.Stat(path); {
	case os.IsNotExist(err):
		// 文件不存在返回入口 HTML 文档内容作为响应
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	case err != nil:
		// 如果期间报错，返回 500 响应
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 一切顺利，则使用 http.FileServer 处理静态资源请求
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

var _ http.Handler = (*RootHandler)(nil)
