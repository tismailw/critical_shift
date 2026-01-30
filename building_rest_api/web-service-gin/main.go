package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

var albums = []album{
	{ID: "101", Title: "The only one", Artist: "saad jinxi", Price: 33.4},
	{ID: "102", Title: "Blue Ridge Mountains", Artist: "colby morton", Price: 991.12},
	{ID: "103", Title: "Lego night", Artist: "john piece", Price: 63.45},
}


// get Albums - get a list of all albums
/*
curl http://localhost:8080/albums
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


// post Albums - add another album
func postAlbums(c *gin.Context) {
/*
curl http://localhost:8080/albums \
--header "Content-type: application/json" \
--data '{"id":{}, "title":{}, "artist":{}, "price":{}}'
*/

	var newAlbum album

	//check if newAlbum is null or not
	err := c.BindJSON(&newAlbum)
	if err != nil { // if the error is not nothing then we return bc that means theres an error
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}


// get Albums{id} - getting the album with the id specified
/*
curl http://localhost:8080/albums/{id}
*/
func getAlbumsByID(c *gin.Context) {

	id := c.Param("id")

	for _, x := range albums {
		if x.ID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": " Album ID does not exist",
		"ID": id,
	})

}
