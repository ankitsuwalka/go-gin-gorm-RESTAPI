package controller

import (
	"net/http"
	"strconv"

	"github.com/ankitkumarsuwalka/app1/databaseconn"
	"github.com/ankitkumarsuwalka/app1/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comp struct {
	DB *gorm.DB
}

func New() *Comp {
	Db := databaseconn.InitDB()
	Db.AutoMigrate(&models.Employee{})
	return &Comp{DB: Db}
}

func (comp *Comp) CreateEmp(c *gin.Context) {
	var emp models.Employee

	err1 := c.ShouldBindJSON(&emp)

	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't create"})
		return
	}
	if err2 := models.CreateEmp(comp.DB, &emp); err2 != nil {
		c.JSON(http.StatusInternalServerError, err2)
		return
	}

	c.JSON(http.StatusOK, emp)
}

func (comp *Comp) GetEmps(c *gin.Context) {
	var emps []models.Employee

	if err1 := models.GetEmps(comp.DB, &emps); err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get"})
		return
	}
	c.JSON(http.StatusOK, emps)
}

func (comp *Comp) GetEmp(c *gin.Context) {
	var emp models.Employee

	id, _ := strconv.Atoi(c.Param("id"))

	if err1 := models.GetEmp(comp.DB, &emp, uint(id)); err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get with given id"})
		return
	}
	c.JSON(http.StatusOK, emp)
}
func (comp *Comp) DeleteEmp(c *gin.Context) {
	var emp models.Employee

	id, _ := strconv.Atoi(c.Param("id"))

	if err1 := models.DeleteEmp(comp.DB, &emp, uint(id)); err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't create"})
		return
	}
	c.JSON(http.StatusOK, emp)
}

func (comp *Comp) UpdateEmp(c *gin.Context) {

	var updatedEmp models.Employee

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&updatedEmp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}

	var existEmp models.Employee

	if err1 := models.GetEmp(comp.DB, &existEmp, uint(id)); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	existEmp.Name = updatedEmp.Name
	existEmp.DOB = updatedEmp.DOB

	if err1 := models.UpdateEmp(comp.DB, &existEmp); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err1})
		return
	}
	c.JSON(http.StatusOK, updatedEmp)

}
