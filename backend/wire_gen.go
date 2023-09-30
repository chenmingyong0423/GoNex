// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/backend/internal/category/handler"
	"github.com/chenmingyong0423/fnote/backend/internal/category/repository"
	"github.com/chenmingyong0423/fnote/backend/internal/category/repository/dao"
	"github.com/chenmingyong0423/fnote/backend/internal/category/service"
	"github.com/chenmingyong0423/fnote/backend/internal/comment/hanlder"
	repository2 "github.com/chenmingyong0423/fnote/backend/internal/comment/repository"
	dao2 "github.com/chenmingyong0423/fnote/backend/internal/comment/repository/dao"
	service2 "github.com/chenmingyong0423/fnote/backend/internal/comment/service"
	handler2 "github.com/chenmingyong0423/fnote/backend/internal/config/handler"
	repository3 "github.com/chenmingyong0423/fnote/backend/internal/config/repository"
	dao3 "github.com/chenmingyong0423/fnote/backend/internal/config/repository/dao"
	service3 "github.com/chenmingyong0423/fnote/backend/internal/config/service"
	service5 "github.com/chenmingyong0423/fnote/backend/internal/email/service"
	hanlder2 "github.com/chenmingyong0423/fnote/backend/internal/friend/hanlder"
	repository6 "github.com/chenmingyong0423/fnote/backend/internal/friend/repository"
	dao6 "github.com/chenmingyong0423/fnote/backend/internal/friend/repository/dao"
	service8 "github.com/chenmingyong0423/fnote/backend/internal/friend/service"
	"github.com/chenmingyong0423/fnote/backend/internal/ioc"
	service7 "github.com/chenmingyong0423/fnote/backend/internal/message/service"
	handler5 "github.com/chenmingyong0423/fnote/backend/internal/message_template/handler"
	repository5 "github.com/chenmingyong0423/fnote/backend/internal/message_template/repository"
	dao5 "github.com/chenmingyong0423/fnote/backend/internal/message_template/repository/dao"
	service6 "github.com/chenmingyong0423/fnote/backend/internal/message_template/service"
	handler3 "github.com/chenmingyong0423/fnote/backend/internal/post/handler"
	repository4 "github.com/chenmingyong0423/fnote/backend/internal/post/repository"
	dao4 "github.com/chenmingyong0423/fnote/backend/internal/post/repository/dao"
	service4 "github.com/chenmingyong0423/fnote/backend/internal/post/service"
	handler4 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/handler"
	repository7 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/repository"
	dao7 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/repository/dao"
	service9 "github.com/chenmingyong0423/fnote/backend/internal/visit_log/service"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp(cfgPath string) (*gin.Engine, error) {
	config := ioc.InitConfig(cfgPath)
	database := ioc.NewMongoDB(config)
	categoryDao := dao.NewCategoryDao(database)
	categoryRepository := repository.NewCategoryRepository(categoryDao)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	commentDao := dao2.NewCommentDao(database)
	commentRepository := repository2.NewCommentRepository(commentDao)
	commentService := service2.NewCommentService(commentRepository)
	configDao := dao3.NewConfigDao(database)
	configRepository := repository3.NewConfigRepository(configDao)
	configService := service3.NewConfigService(configRepository)
	postDao := dao4.NewPostDao(database)
	postRepository := repository4.NewPostRepository(postDao)
	postService := service4.NewPostService(postRepository)
	emailService := service5.NewEmailService()
	msgTplDao := dao5.NewMsgTplDao(database)
	msgTplRepository := repository5.NewMsgTplRepository(msgTplDao)
	msgTplService := service6.NewMsgTplService(msgTplRepository)
	messageService := service7.NewMessageService(configService, emailService, msgTplService)
	commentHandler := hanlder.NewCommentHandler(commentService, configService, postService, messageService)
	configHandler := handler2.NewConfigHandler(configService)
	friendDao := dao6.NewFriendDao(database)
	friendRepository := repository6.NewFriendRepository(friendDao)
	friendService := service8.NewFriendService(friendRepository)
	friendHandler := hanlder2.NewFriendHandler(friendService, messageService, configService)
	postHandler := handler3.NewPostHandler(postService)
	visitLogDao := dao7.NewVisitLogDao(database)
	visitLogRepository := repository7.NewVisitLogRepository(visitLogDao)
	visitLogService := service9.NewVisitLogService(visitLogRepository)
	visitLogHandler := handler4.NewVisitLogHandler(visitLogService, configService)
	msgTplHandler := handler5.NewMsgTplHandler(msgTplService)
	writer := ioc.InitLogger(config)
	v := ioc.InitMiddlewares(config, writer)
	validators := ioc.InitGinValidators()
	engine, err := ioc.NewGinEngine(categoryHandler, commentHandler, configHandler, friendHandler, postHandler, visitLogHandler, msgTplHandler, v, validators)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
