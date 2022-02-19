package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// User представляет данные
type User struct {
	ID        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}

// albums slice для заполнения данных User
var albums = []User{
	{ID: "1", Firstname: "Naruto", Lastname: "Uzumake", Email: "konoha-uzumake@mail.com", Age: 18, Created: time.Now()},
	{ID: "2", Firstname: "Saske", Lastname: "Uchiha", Email: "clan-uchiha@mail.com", Age: 19, Created: time.Now()},
	{ID: "3", Firstname: "Pain", Lastname: "Uzu", Email: "uzu-pain@mail.com", Age: 25, Created: time.Now()},
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumID)
	router.PUT("/albums/:id", updateAlbumID)
	return router
}

func main() {
	router := getRouter()
	router.Run("localhost:8080")
}

// getAlbums ответ всех альбомов в формате JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID находит альбом, значение ID которого совпадает с параметром id
// параметру, отправленному клиентом, а затем возвращает этот альбом в качестве ответа.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Перебираем список альбомов, ищем
	// альбом, значение ID которого совпадает с параметром.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "album not found"})
}

// postAlbums добавляет User из JSON, полученного в теле запроса
func postAlbums(c *gin.Context) {
	var newAlbum User
	// вызов BindJSON для привязки полученного JSON к
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func deleteAlbumID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"massage": "not_found"})
}

func updateAlbumID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if albums[i].ID == id {
			c.BindJSON(&albums[i])
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"massage": "not_found"})
}
