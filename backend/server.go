package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melbahja/goph"
	"net/http"
)

func CreateServer(c *gin.Context) {
	input := Server{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ipAdress := c.ClientIP()
	host := "159.69.55.238"
	username := input.UserSsh
	password := input.PasswordSsh
	client, err := goph.NewUnknown(username, host, goph.Password(password))
	if err != nil {
		fmt.Println("sono qui 1")
		panic(err)
	}
	distro, err := client.Run("cat /etc/os-release")
	if err != nil {
		fmt.Println("sono qui 2")
		panic(err)
	}
	dis := string(distro)              //per trasformare la distro in string
	pro, err1 := client.Run("rpm -qa") //per ottenere le destri installate
	if err1 != nil {                   //pro:per salvare il risultato del commando
		fmt.Println("sono qui 2")
		panic(err)
	}
	ls := string(pro)
	fmt.Println(ls)
	in := Server{
		Id:               input.Id,
		Ip:               ipAdress,
		UserSsh:          input.UserSsh,
		PasswordSsh:      input.PasswordSsh,
		DistroInstallata: dis,
		Programmi:        ls,
	}
	db.Create(&in)
	c.JSON(http.StatusOK, gin.H{
		"data": in,
		"id":   in.Id,
	})
}
func GetServers(c *gin.Context) {
	var out []Server
	db.Find(&out)
	c.JSON(http.StatusOK, out)
}
func GetOneServer(c *gin.Context) {
	id := c.Param("id")
	search := Server{}
	db.Table("server").Where("id=?", id).Find(&search)
	c.JSON(http.StatusOK, search)
}
func EditServer(c *gin.Context) {
	id := c.Param("id")
	input := Server{}
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db.Table("server").Where("id=?", id).Updates(Server{
		UserSsh:          input.UserSsh,
		PasswordSsh:      input.PasswordSsh,
		DistroInstallata: input.DistroInstallata,
		Programmi:        input.Programmi})
	c.JSON(http.StatusOK, gin.H{
		"edit": input,
		"id":   input.Id,
	})
}
func DeleteServer(c *gin.Context) {
	id := c.Param("id")
	sin := Server{} //il server da cancellare
	db.Table("server").Where("id=?", id).Delete(sin)
}
