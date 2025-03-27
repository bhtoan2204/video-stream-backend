# /bin/bash
cd ../apps/gateway/cmd/main
swag init --parseDependency -g main.go -d .,../../internal/modules/auth/handler,../../internal/modules/user/handler,../../internal/modules/comment/handler,../../internal/modules/video/handler