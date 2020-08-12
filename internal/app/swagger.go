/*
Package app 生成swagger文档

文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format

使用方式：

	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./internal/app/swagger.go --output ./internal/app/swagger */
package app

// @title gin-admin-template
// @version 6.4.2
// @description RBAC scaffolding based on GIN + GORM/MONGO + CASBIN + WIRE.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @contact.name
// @contact.email
