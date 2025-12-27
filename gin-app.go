package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    
    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func index(){}

func main(){
	router := gin.Default()
    router.LoadHTMLGlob("web/*") //加载文件夹
    router.GET("/", func(c *gin.Context) {
        // 渲染index.html（参数2为模板文件名，需和目录中一致）
        c.HTML(200, "index.html", nil)
    })//无法嵌入文件 
    router.GET("/api/albums", getAlbums)
    router.GET("/api/albums/:id", getAlbumByID)
    router.POST("/api/albums", postAlbums)
    //router.Run("localhost:8080"
    router.StaticFS("/dl",gin.Dir("files",true))//true 显示文件夹内容
    router.Run()

}