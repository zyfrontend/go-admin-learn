package main

import (
	"app/routes"
	"app/test"
	"app/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	test.TestGorm()
	r := routes.Routers()
	r.Run()
}
