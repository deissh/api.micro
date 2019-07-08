package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-messages/common"
	service "github.com/nekko-ru/api/service-messages/handlers"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/messages.addChatUser")
		g.GET("/messages.createChat")
		g.GET("/messages.delete")
		g.GET("/messages.deleteChatPhoto")
		g.GET("/messages.setChatPhoto")
		g.GET("/messages.deleteConversation")
		g.GET("/messages.edit")
		g.GET("/messages.editChat")
		g.GET("/messages.getById")
		g.GET("/messages.getChat")
		g.GET("/messages.getConversationMembers")
		g.GET("/messages.getConversations")
		g.GET("/messages.getConversationsById")
		g.GET("/messages.getHistory")
		g.GET("/messages.getLastActivity")
		g.GET("/messages.getInviteLink")
		g.GET("/messages.joinChatByInviteLink")
		g.GET("/messages.markAsRead")
		g.GET("/messages.removeChatUser")
		g.GET("/messages.send")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
