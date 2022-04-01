package route

import (
	"log"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func init() {
	handlers = append(handlers,
		Handler{
			"GET",
			"/ws",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return wsHandler
			},
		},
	)
}

var upgrader = websocket.Upgrader{} // use default options

func wsHandler(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	c.WriteJSON(map[string]string{"status": "ok"})
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
