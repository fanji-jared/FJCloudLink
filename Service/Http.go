package Service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	// 最大文件大小为10MB
	maxFileSize = 1024 * 1024 * 10
	// 最大文件名长度为255个字符
	maxFileNameLength = 255
	// 最大文件数量为10个
	maxFiles = 10
)

// 定义允许的文件类型
var allowedTypes = []string{
	"text/plain", "text/markdown", "text/csv", "text/html", "text/css",
	"text/javascript", "application/pdf", "application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.ms-powerpoint", "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"application/x-dxf", "application/x-eps", "application/x-latex",
	"application/x-compressed", "video/*", "audio/*", "model/*",
}

func handleUpload(c *gin.Context) {
	// 获取上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["files"]

	// 检查是否有文件上传
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有选择任何文件"})
		return
	}

	// 检查文件总数是否超过限制
	if len(files) > maxFiles {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": fmt.Sprintf("上传的文件数量过多，最多允许%d个文件", maxFiles)})
		return
	}

	// 遍历所有文件并保存
	for _, fileHeader := range files {
		if err := saveFile(fileHeader, "./uploads/"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// 发送成功响应
	if len(files) == 1 {
		c.JSON(http.StatusOK, gin.H{"status": "uploaded", "file": files[0].Filename})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "uploadedFolder", "files": len(files)})
	}
}

func saveFile(file *multipart.FileHeader, uploadDir string) error {
	// 确保上传目录存在
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}

	// 检查文件大小限制
	if file.Size > maxFileSize {
		return fmt.Errorf("文件大小超过限制")
	}

	// 检查文件类型限制
	if !isAllowedType(file.Header.Get("Content-Type")) {
		return fmt.Errorf("不允许的文件类型")
	}

	// 检查文件名长度限制
	if len(file.Filename) > maxFileNameLength {
		return fmt.Errorf("文件名过长")
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件
	safeFilename := sanitizeFilename(file.Filename)
	dst, err := os.Create(filepath.Join(uploadDir, safeFilename))
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

// 判断文件类型是否被允许
func isAllowedType(mimeType string) bool {

	// 遍历允许的文件类型
	for _, t := range allowedTypes {
		// 如果文件类型等于允许的类型或者以允许的类型开头，则返回true
		if mimeType == t || strings.HasPrefix(mimeType, t+";") {
			return true
		}
	}

	// 如果文件类型不在允许的类型中，则返回false
	return false
}

func sanitizeFilename(filename string) string {
	// 使用更安全的清理方式
	safeName := filepath.Base(filename)
	return strings.ReplaceAll(safeName, "..", "")
}

func SetupHTTPServer() *gin.Engine {
	r := gin.Default()

	// 允许跨域请求
	r.Use(cors.Default())

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
	r.POST("/upload", handleUpload)

	// 文件下载
	r.GET("/download", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "downloaded"})
	})

	return r
}
