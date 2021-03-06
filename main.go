package main

import (
	"bufio"
	"errors"
	"fmt"
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

// idnameFileExist is a function that checks if idname file exists.
func idnameFileExist() bool {
	if _, err := os.Stat(idnameFile); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
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

// persistNewIdname is a function that saves new idname to the idnameFile
func persistNewIdname() {
	f, err := os.OpenFile(idnameFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	defer f.Close()
	check(err)
	for _, elem := range idnames {
		_, err = f.WriteString(fmt.Sprintf("%v,%v\n", elem.ID, elem.Name))
		check(err)
	}
}

// postIdnameAPI is a function that takes json data from POST request and adds the info to idnames file (if the data is correct)
// example:
// curl localhost:8080/newid \
// 	--request "POST" \
// 	--header "Content-Type: application/json" \
// 	--data '{"id": 5,"name": "test5"}'
func postIdnameAPI(c *gin.Context) {
	var newIdname idname
	if err := c.BindJSON(&newIdname); err != nil {
		return
	}
	idnames = append(idnames, newIdname)
	c.IndentedJSON(http.StatusCreated, idnames)
	persistNewIdname()
}

// postIdnameWP is a function that creates new idname from the params given in url like /newid/:id/:name
func postIdnameWP(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	check(err)
	name := c.Param("name")
	idnames = append(idnames, idname{ID: id, Name: name})
	c.IndentedJSON(http.StatusCreated, idnames)
	persistNewIdname()
}

// postIdnameWPParams is a function that creates new idname from the named params given in url like /newid?id=<id>&name=<name>
func postIdnameWPParams(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		return
	}
	id, err := strconv.Atoi(idStr)
	check(err)
	name, ok := c.GetQuery("name")
	if !ok {
		return
	}
	idnames = append(idnames, idname{ID: id, Name: name})
	c.IndentedJSON(http.StatusCreated, idnames)
	persistNewIdname()
}

func main() {
	// read in the idnames
	if idnameFileExist() {
		readIdnames()
	}
	// setup router
	router := gin.Default()
	//// at the moment do not set up any trusted proxies
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("templates/*.html")
	// GET and POST
	router.GET("/hi", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi there\n")
	})
	router.GET("/", mainPage)
	router.GET("/idapi", getIdnamesAPI)
	router.GET("/idwp", getIdnamesWP)
	router.POST("/newid", postIdnameAPI)
	router.GET("/newid/:id/:name", postIdnameWP)
	router.GET("/newid/w_params", postIdnameWPParams)

	// Run
	// router.Run("localhost:8080")
	router.Run(":8080")

}
