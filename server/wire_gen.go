// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chenmingyong0423/fnote/server/internal/aggregate_post"
	handler9 "github.com/chenmingyong0423/fnote/server/internal/backup/handler"
	service12 "github.com/chenmingyong0423/fnote/server/internal/backup/service"
	handler2 "github.com/chenmingyong0423/fnote/server/internal/category/handler"
	repository2 "github.com/chenmingyong0423/fnote/server/internal/category/repository"
	dao2 "github.com/chenmingyong0423/fnote/server/internal/category/repository/dao"
	service3 "github.com/chenmingyong0423/fnote/server/internal/category/service"
	"github.com/chenmingyong0423/fnote/server/internal/comment/hanlder"
	repository4 "github.com/chenmingyong0423/fnote/server/internal/comment/repository"
	dao4 "github.com/chenmingyong0423/fnote/server/internal/comment/repository/dao"
	service4 "github.com/chenmingyong0423/fnote/server/internal/comment/service"
	handler8 "github.com/chenmingyong0423/fnote/server/internal/count_stats/handler"
	repository3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository"
	dao3 "github.com/chenmingyong0423/fnote/server/internal/count_stats/repository/dao"
	service2 "github.com/chenmingyong0423/fnote/server/internal/count_stats/service"
	handler7 "github.com/chenmingyong0423/fnote/server/internal/data_analysis/handler"
	service6 "github.com/chenmingyong0423/fnote/server/internal/email/service"
	"github.com/chenmingyong0423/fnote/server/internal/file/handler"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository"
	"github.com/chenmingyong0423/fnote/server/internal/file/repository/dao"
	"github.com/chenmingyong0423/fnote/server/internal/file/service"
	hanlder2 "github.com/chenmingyong0423/fnote/server/internal/friend/hanlder"
	repository7 "github.com/chenmingyong0423/fnote/server/internal/friend/repository"
	dao7 "github.com/chenmingyong0423/fnote/server/internal/friend/repository/dao"
	service9 "github.com/chenmingyong0423/fnote/server/internal/friend/service"
	"github.com/chenmingyong0423/fnote/server/internal/global"
	"github.com/chenmingyong0423/fnote/server/internal/ioc"
	service8 "github.com/chenmingyong0423/fnote/server/internal/message/service"
	handler5 "github.com/chenmingyong0423/fnote/server/internal/message_template/handler"
	repository6 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository"
	dao6 "github.com/chenmingyong0423/fnote/server/internal/message_template/repository/dao"
	service7 "github.com/chenmingyong0423/fnote/server/internal/message_template/service"
	handler3 "github.com/chenmingyong0423/fnote/server/internal/post/handler"
	repository5 "github.com/chenmingyong0423/fnote/server/internal/post/repository"
	dao5 "github.com/chenmingyong0423/fnote/server/internal/post/repository/dao"
	service5 "github.com/chenmingyong0423/fnote/server/internal/post/service"
	"github.com/chenmingyong0423/fnote/server/internal/post_draft"
	"github.com/chenmingyong0423/fnote/server/internal/post_index"
	"github.com/chenmingyong0423/fnote/server/internal/post_like"
	handler6 "github.com/chenmingyong0423/fnote/server/internal/tag/handler"
	repository9 "github.com/chenmingyong0423/fnote/server/internal/tag/repository"
	dao9 "github.com/chenmingyong0423/fnote/server/internal/tag/repository/dao"
	service11 "github.com/chenmingyong0423/fnote/server/internal/tag/service"
	handler4 "github.com/chenmingyong0423/fnote/server/internal/visit_log/handler"
	repository8 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository"
	dao8 "github.com/chenmingyong0423/fnote/server/internal/visit_log/repository/dao"
	service10 "github.com/chenmingyong0423/fnote/server/internal/visit_log/service"
	"github.com/chenmingyong0423/fnote/server/internal/website_config"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func initializeApp() (*gin.Engine, error) {
	database := ioc.NewMongoDB()
	fileDao := dao.NewFileDao(database)
	fileRepository := repository.NewFileRepository(fileDao)
	fileService := service.NewFileService(fileRepository)
	fileHandler := handler.NewFileHandler(fileService)
	categoryDao := dao2.NewCategoryDao(database)
	categoryRepository := repository2.NewCategoryRepository(categoryDao)
	countStatsDao := dao3.NewCountStatsDao(database)
	countStatsRepository := repository3.NewCountStatsRepository(countStatsDao)
	countStatsService := service2.NewCountStatsService(countStatsRepository)
	model := website_config.InitWebsiteConfigModule(database)
	iWebsiteConfigService := model.Svc
	categoryService := service3.NewCategoryService(categoryRepository, countStatsService, iWebsiteConfigService)
	categoryHandler := handler2.NewCategoryHandler(categoryService)
	commentDao := dao4.NewCommentDao(database)
	commentRepository := repository4.NewCommentRepository(commentDao)
	commentService := service4.NewCommentService(commentRepository)
	postDao := dao5.NewPostDao(database)
	postRepository := repository5.NewPostRepository(postDao)
	post_likeModel := post_like.InitPostLikeModule(database)
	iPostLikeService := post_likeModel.Svc
	postService := service5.NewPostService(postRepository, iWebsiteConfigService, countStatsService, fileService, iPostLikeService)
	emailService := service6.NewEmailService()
	msgTplDao := dao6.NewMsgTplDao(database)
	msgTplRepository := repository6.NewMsgTplRepository(msgTplDao)
	msgTplService := service7.NewMsgTplService(msgTplRepository)
	messageService := service8.NewMessageService(iWebsiteConfigService, emailService, msgTplService)
	commentHandler := hanlder.NewCommentHandler(commentService, iWebsiteConfigService, postService, messageService, countStatsService)
	websiteConfigHandler := model.Hdl
	friendDao := dao7.NewFriendDao(database)
	friendRepository := repository7.NewFriendRepository(friendDao)
	friendService := service9.NewFriendService(friendRepository)
	friendHandler := hanlder2.NewFriendHandler(friendService, messageService, iWebsiteConfigService)
	postHandler := handler3.NewPostHandler(postService, iWebsiteConfigService)
	visitLogDao := dao8.NewVisitLogDao(database)
	visitLogRepository := repository8.NewVisitLogRepository(visitLogDao)
	visitLogService := service10.NewVisitLogService(visitLogRepository)
	visitLogHandler := handler4.NewVisitLogHandler(visitLogService, countStatsService)
	msgTplHandler := handler5.NewMsgTplHandler(msgTplService)
	tagDao := dao9.NewTagDao(database)
	tagRepository := repository9.NewTagRepository(tagDao)
	tagService := service11.NewTagService(tagRepository, countStatsService)
	tagHandler := handler6.NewTagHandler(tagService)
	dataAnalysisHandler := handler7.NewDataAnalysisHandler(visitLogService, countStatsService)
	countStatsHandler := handler8.NewCountStatsHandler(countStatsService)
	backupService := service12.NewBackupService(database)
	backupHandler := handler9.NewBackupHandler(backupService)
	writer := ioc.InitLogger()
	v, err := global.IsWebsiteInitializedFn(database)
	if err != nil {
		return nil, err
	}
	v2 := ioc.InitMiddlewares(writer, v)
	validators := ioc.InitGinValidators()
	post_indexModel := post_index.InitPostIndexModule(model)
	postIndexHandler := post_indexModel.Hdl
	post_draftModel := post_draft.InitPostDraftModule(database)
	postDraftHandler := post_draftModel.Hdl
	aggregate_postModel := aggregate_post.InitAggregatePostModule(postService, post_draftModel)
	aggregatePostHandler := aggregate_postModel.Hdl
	postLikeHandler := post_likeModel.Hdl
	engine, err := ioc.NewGinEngine(fileHandler, categoryHandler, commentHandler, websiteConfigHandler, friendHandler, postHandler, visitLogHandler, msgTplHandler, tagHandler, dataAnalysisHandler, countStatsHandler, backupHandler, v2, validators, postIndexHandler, postDraftHandler, aggregatePostHandler, postLikeHandler)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
