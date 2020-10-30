package butterfly

import "github.com/pwh19920920/butterfly/server"
import _ "github.com/pwh19920920/butterfly/logger"

func Run() {
	server.StartHttpServer()
}
