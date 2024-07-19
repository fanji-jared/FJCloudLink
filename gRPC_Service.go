// 定义协议版本
syntax = "proto3";

// 定义包名
package file;

// 定义服务
service FileService {
  // 上传文件
  rpc UploadFile(UploadRequest) returns (UploadResponse);
  // 下载文件
  rpc DownloadFile(DownloadRequest) returns (stream DownloadResponse);
}

// 定义上传请求消息
message UploadRequest {
  // 文件名
  string filename = 1;
  // 文件数据
  bytes data = 2;
}

// 定义上传响应消息
message UploadResponse {
  // 状态
  string status = 1;
}

// 定义下载请求消息
message DownloadRequest {
  // 文件名
  string filename = 1;
}

// 定义下载响应消息
message DownloadResponse {
  // 文件数据
  bytes data = 1;
}