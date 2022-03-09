package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// idname is a structure that we use to pass data.
type idname struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// idnames is an array of structure idname elements.
var idnames []idname

// idnameFile is the name of the file, where idnames are stored.
var idnameFile string = "idnames.csv"

// check is a function to check, if err is nil or not.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// readIdnames is a function, that reads idnames from a file
func readIdnames() {
	f, err := os.Open(idnameFile)
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	defer func() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(line[0])
		check(err)
		idnames = append(idnames, idname{ID: id, Name: line[1]})
	}
}

// getIdnamesAPI is a function that returns idnames in a json format.
func getIdnamesAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, idnames)
}

// getIdnamesWP is a function that parses a webpage, where idnames are shown
func getIdnamesWP(c *gin.Context) {
	c.HTML(http.StatusOK, "idnames.html", gin.H{
		"title":   "idnames page",
		"idnames": idnames,
	})
}

// mainPage is a function to parse and show main webpage
func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "main page",
	})
}

func main() {
	// read in the idnames
	readIdnames()
	// setup router
	router := gin.Default()
	//// at the moment do not set up any trusted proxies
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("templates/*.html")
	// GET and POST
	// router.GET("/", func(c *gin.Context) {
	// c.String(http.StatusOK, "Hi there\n")
	// })
	router.GET("/", mainPage)
	router.GET("/idapi", getIdnamesAPI)
	router.GET("/idwp", getIdnamesWP)
	// router.POST("/newid", postIdname)
	// router.GET("/newid/:id/:name", postIdnameParams)
	// router.GET("/newid/idparams", postIdnameParams2)

	// Run
	// router.Run("localhost:8080")
	router.Run(":8080")

}
