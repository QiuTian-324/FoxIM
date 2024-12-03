// internal/handlers/chat.go
package handlers

import (
	"yug_server/global"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ChatHandler struct {
	chatService services.WsService
	logger      *zap.Logger
}

func NewChatHandler(chatService *services.WsService, logger *zap.Logger) *ChatHandler {
	return &ChatHandler{chatService: *chatService, logger: logger}
}

func (h *ChatHandler) Ws(ctx *gin.Context) {
	conn, err := global.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.logger.Error("websocket upgrade error", zap.Error(err))
		return
	}

	// userID := ctx.MustGet("id").(uint64)

	defer func() {
		conn.Close()
		// h.chatService.RemoveConnection(userID)
	}()

	// h.chatService.AddConnection(userID, conn)

	// 启动心跳检测
	// go h.chatService.Heartbeat(conn, userID)

	for {
		_, messageData, err := conn.ReadMessage()
		if err != nil {
			return
		}

		h.logger.Info("receive message", zap.String("message", string(messageData)))

		// err = h.chatService.SendMessage(conn, messageData)
		// if err != nil {
		// 	return
		// }
	}
}
