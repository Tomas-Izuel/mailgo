package rest

import (
	"fmt"
	"mailgo/lib"
)

func Init() {
	createServer()
	initRouter()
	getServer().Run(fmt.Sprintf(":%d", lib.GetEnv().Port))
}
