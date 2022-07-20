package main

import (
	"server/common"

	"github.com/gin-gonic/gin"
)

func main() {
	db := common.InitDB()
	defer db.Statement.ReflectValue.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
