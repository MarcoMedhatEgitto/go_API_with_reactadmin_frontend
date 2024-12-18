package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Server struct {
	Id               int    `json:"id"`
	Ip               string `json:"ip"`
	UserSsh          string `json:"user_ssh"`
	PasswordSsh      string `json:"password_ssh"`
	DistroInstallata string `json:"distro_installata"`
	Programmi        string `json:"lista_programmi"`
}
type Swarm struct {
	Id              int    `json:"id"`
	Uuid            string `json:"uuid"`
	Nome            string `json:"nome"`
	FileContentText string `json:"file_content_text"`
	ServerId        int    `json:"server_id"`
}

func (Swarm) TableName() string {
	return "swarm"
}

func (Server) TableName() string {
	return "server"
}

var db *gorm.DB

func main() {
	prog := "user=postgres host=localhost password='' dbname=chat port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(prog), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.POST("/server", CreateServer)
	router.GET("/server", GetServers)
	router.GET("/server/:id", GetOneServer)
	router.PUT("/server/:id", EditServer)
	router.DELETE("/server/:id", DeleteServer)
	//router.GET("/swarm?server_id=:id", func(c *gin.Context) {
	//	var out []sw
	//	serverId := c.Param("id")
	//	db.Table("swarmuno").Where("server_id=?", serverId).Find(&out)
	//	c.JSON(http.StatusOK, out)
	//})
	router.GET("/swarm", GetSwarms)
	router.GET("/swarm/:id", GetOneSwarm)
	router.PUT("/swarm/:id", EditSwarm)
	router.POST("/swarm", CreateSwarm)
	router.DELETE("/swarm/:id", DeleteSwarm)
	//router.GET("/swarn?serverId=1", func(c *gin.Context) {
	//	id := c.Param("id")
	//	var sin []swstruct int
	//	db.Table("swarmuno").Where("server_id=?", id).Find(&sin)
	//	c.JSON(http.StatusOK, sin)
	//})
	/*
		-- INPUT BDY
		{
			"nome": "Extark Server",
			"ip": "...",
			"username": "stage1",
			"password": "..."
		}
		-- OUTPUT
		{"response": "ok"}
	*/
	// suggerimento: con le credenziali appena ricevute fate la connessione ssh e prendete la stringa contenente la lista dei programmi,
	// successivamente con tutte le informazioni che avete ottenuto aggiungete la riga nella tabella dei server
	//	router.POST("/server", func(c *gin.Context) {
	/*var dimostra []OutPut
	ipAdress := c.ClientIP()
	db.Find(&dimostra)
	input := OutPut{
		Ip: ipAdress,
	}
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"ip adress": ipAdress,
	})*/
	//	})

	/*
		-- INPUT BODY
		vuoto
		-- OUTPUT
		[
			{
				"nome": "Extark Server",
				"ip": "...",
				"username": "stage1",
				"programs": "..."
			},
			{
				"nome": "NASA Server",
				"ip": "...",
				"username": "pippo",
				"programs": "..."
			},
			...
		]
	*/
	//	router.GET("/server", func(c *gin.Context) {
	/*var dimostra []OutPut
	ipAdress := c.ClientIP()
	db.Find(&dimostra)
	input := OutPut{
		Ip: ipAdress,
	}
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"ip adress": ipAdress,
	})*/
	//	})

	// QUESTO DEVE ESSERE TRASFORMATO NELLA API POST:server che è presente sopra
	/*router.GET("/dimostra/i/dati", func(c *gin.Context) { //per provare la funzionalita' del programma
		var dimostra []OutPut
		ipAdress := c.ClientIP()
		db.Find(&dimostra)
		input := OutPut{
			Ip: ipAdress,
		}
		db.Create(&input)
		c.JSON(http.StatusOK, gin.H{
			"ip adress": ipAdress,
		})
	})*/

	// QUESTO NON DEVE ESISTERE
	/*router.GET("/dimostra/i/pachetti", func(c *gin.Context) {
		cmd := exec.Command("rpm", "-qa")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		lines := strings.Split(string(out), "\n")
		fmt.Println("installed packages:")
		for _, line := range lines {
			c.JSON(http.StatusOK, gin.H{
				"": line,
			})
		}
	})*/
	router.Run("0.0.0.0:8080")
	//bjnk
	//lansciamo macchine vituali sul nostro sito
}
