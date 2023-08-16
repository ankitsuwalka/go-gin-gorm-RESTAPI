package main

import (
	"github.com/ankitkumarsuwalka/app1/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	Comp := controller.New()

	r.GET("/api/emps", Comp.GetEmps)
	r.GET("/api/emps/:id", Comp.GetEmp)
	r.POST("/api/emps", Comp.CreateEmp)
	r.DELETE("/api/emps/:id", Comp.DeleteEmp)
	r.PUT("/api/emps/:id", Comp.UpdateEmp)

	r.Run(":8080")
}
