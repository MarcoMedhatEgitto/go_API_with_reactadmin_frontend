package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSwarms(c *gin.Context) {
	var out []Swarm
	serverId := c.Query("server_id")
	fmt.Println(serverId)
	//serverId := c.Request.URL.Query()
	query := db
	//search := sw{}
	if serverId != "" {
		query.Where("server_id=?", serverId).Find(&out)
		c.JSON(http.StatusOK, out)
	} else {
		query.Find(&out)
		c.JSON(http.StatusOK, out)
	}
}
func GetOneSwarm(c *gin.Context) {
	id := c.Param("id")
	search := Swarm{}
	db.Table("swarm").Where("id=?", id).Find(&search)
	c.JSON(http.StatusOK, search)
}
func EditSwarm(c *gin.Context) {
	id := c.Param("id")
	sin := Swarm{}
	if err := c.BindJSON(&sin); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	search := Swarm{}
	db.Table("swarm").Where("id=?", id).Find(&search).Updates(Swarm{Uuid: sin.Uuid, Nome: sin.Nome, FileContentText: sin.FileContentText, ServerId: sin.ServerId})
	c.JSON(http.StatusOK, gin.H{
		"data": search,
		"id":   search.Id,
	})
}
func CreateSwarm(c *gin.Context) {
	sin := Swarm{}
	if err := c.BindJSON(&sin); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ne := Swarm{
		Uuid:            sin.Uuid,
		Nome:            sin.Nome,
		FileContentText: sin.FileContentText,
		ServerId:        sin.ServerId,
	}
	db.Create(&ne)
	c.JSON(http.StatusOK, gin.H{
		"data": ne,
		"id":   ne.Id,
	})
}
func DeleteSwarm(c *gin.Context) {
	id := c.Param("id")
	sin := Swarm{}
	db.Table("swarm").Where("id=?", id).Delete(sin)
}
