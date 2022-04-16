package handler

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/codeyifei/proxy-router/types"
	"github.com/creasty/defaults"
)

type RootHandler struct {
	PathPrefix string
	StaticPath string
	IndexPaths []string `default:"[\"index.html\"]"`
}

func NewRootHandler(pathPrefix, staticPath string, indexPaths ...string) *RootHandler {
	if !strings.HasPrefix(pathPrefix, "/") {
		pathPrefix = "/" + pathPrefix
	}
	h := &RootHandler{
		PathPrefix: pathPrefix,
		StaticPath: staticPath,
		IndexPaths: indexPaths,
	}
	defaults.MustSet(h)
	return h
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 路径的绝对路径
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/web")
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// 如果获取失败，返回 400 响应
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 在 URL 路径前加上静态资源根目录
	path = filepath.Join(h.StaticPath, path)

	// 检查对应资源文件是否存在
	ok, err := fileExists(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		var files []string
		for _, index := range h.IndexPaths {
			files = append(files, filepath.Join(h.StaticPath, index))
		}
		var indexPath string
		indexPath, err = findFirstExistFile(files)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, indexPath)
		return
	}

	// 一切顺利，则使用 http.FileServer 处理静态资源请求
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}

var _ types.Handler = (*RootHandler)(nil)

func findFirstExistFile(files []string) (string, error) {
	var (
		ok  bool
		err error
	)
	for _, file := range files {
		ok, err = fileExists(file)
		if err != nil {
			return "", err
		}
		if ok {
			return file, nil
		}
	}
	return "", errors.New("所有文件均不存在")
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		} else if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
