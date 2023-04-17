package main

import (
	"search/src/dao"
	"search/src/routes"
)

func main() {

	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}

	defer dao.Close()
	//dao.SqlSession.AutoMigrate(&model.Record{})

	r := routes.InitRouter()
	errRun := r.Run(":8080")
	if errRun != nil {
		return
	}

}
