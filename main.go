package main

import (
	"github.com/fanji-jared/FJCloudLink/Service"
)

func main() {
	// 启动 gRPC 服务
	// go func() {
	// 	// gRPC 服务代码
	// }()

	// 设置 HTTP 服务
	httpServer := Service.SetupHTTPServer()

	// 添加 gRPC handler 到 Gin
	// httpServer.Any("/file/*any", func(c *gin.Context) {
	// 	c.NextProto(func(rw http.ResponseWriter, req *http.Request) error {
	// 		return grpcServer.ServeHTTP(rw, req)
	// 	})
	// })

	// 设置可信代理
	httpServer.SetTrustedProxies([]string{"192.168.1.1"}) // 你的负载均衡器或反向代理的IP地址

	// 启动 HTTP 服务
	httpServer.Run(":6430")
}
