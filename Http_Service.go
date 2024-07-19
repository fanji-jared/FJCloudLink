package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handleUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	// 检查上传的文件数量
	if len(files) == 1 {
		// 处理单个文件上传
		file := files[0]
		err := saveFile(file, "./uploads/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "uploaded", "file": file.Filename})
	} else {
		// 处理文件夹上传
		for _, file := range files {
			err := saveFile(file, "./uploads/")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "uploadedFolder", "files": len(files)})
	}
}

func saveFile(file *multipart.FileHeader, uploadDir string) error {
	// 确保上传目录存在
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return err
	}

	// 检查文件大小限制
	if file.Size > 10*1024*1024 { // 限制文件大小为10MB
		return fmt.Errorf("文件大小超过限制")
	}

	// 检查文件类型限制
	allowedTypes := map[string]bool{
		// 文本格式
		// "text/*": true, // 所有文本文件
		"text/plain": true, // 纯文本文件
		"text/markdown": true, // Markdown文件
		"text/csv": true, // CSV文件
		"text/html": true, // HTML文件
		"text/css": true, // CSS文件
		"text/javascript": true, // JavaScript文件

		// 文档格式
		"application/pdf": true, // PDF文档
		"application/msword": true, // Word文档
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true, // Word文档（新格式）
		"application/vnd.ms-excel": true, // Excel文档
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true, // Excel文档（新格式）
		"application/vnd.ms-powerpoint": true, // PowerPoint文档
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": true, // PowerPoint文档（新格式）
		"application/x-dxf": true, // DXF文件
		"application/x-eps": true, // EPS文件
		"application/x-latex": true, // LaTeX文件

		// 统一格式
		"application/x-compressed": true, // 所有压缩文件
		"video/*": true, // 所有视频文件
		"audio/*": true, // 所有音频文件
		"image/*": true, // 所有图片文件
		"model/*": true, // 所有模型文件

		// 不允许
		"application/x-msdownload": false, // 所有可执行文件
	}
	if _, ok := allowedTypes[file.Header.Get("Content-Type")]; !ok {
		return fmt.Errorf("不允许的文件类型")
	}

	// 检查文件名长度限制
	if len(file.Filename) > 255 {
		return fmt.Errorf("文件名过长")
	}

	// 检查文件数量限制
	if len(files) > 10 {
		return fmt.Errorf("文件数量超过限制")
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(filepath.Join(uploadDir, sanitizeFilename(file.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

// sanitizeFilename 清理文件名，防止路径遍历攻击
func sanitizeFilename(filename string) string {
	return strings.ReplaceAll(filename, "..", "")
}

func setupHTTPServer() *gin.Engine {
	r := gin.Default()

	// 用户注册
	r.POST("/register", func(c *gin.Context) {
		// 处理用户注册逻辑
		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	})

	// 用户登录
	r.POST("/login", func(c *gin.Context) {
		// 处理用户登录逻辑
		c.JSON(http.StatusOK, gin.H{"status": "logged in"})
	})

	// 文件 / 文件夹 上传
	r.POST("/upload", handleUpload})

	// 文件下载
	r.GET("/download", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "downloaded"})
	})

	return r
}