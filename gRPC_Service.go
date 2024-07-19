syntax = "proto3";

package file;

service FileService {
  rpc UploadFile(UploadRequest) returns (UploadResponse);
  rpc DownloadFile(DownloadRequest) returns (stream DownloadResponse);
}

message UploadRequest {
  string filename = 1;
  bytes data = 2;
}

message UploadResponse {
  string status = 1;
}

message DownloadRequest {
  string filename = 1;
}

message DownloadResponse {
  bytes data = 1;
}