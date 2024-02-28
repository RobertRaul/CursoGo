package vinylapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	Id     int     `json:"id"`
	Tittle string  `json:"tittle"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumnes = []Album{
	{Id: 1, Tittle: "Hora Loca", Artist: "Quien sabe", Price: 10.99},
	{Id: 2, Tittle: "Loca", Artist: "Quien sabe", Price: 12.99},
	{Id: 3, Tittle: "Que hacen", Artist: "Quien sabe", Price: 13.99},
}

func Ini() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:9099")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumnes)
}
func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albumnes = append(albumnes, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
