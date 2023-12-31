package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums godoc
// @Summary         List albums
// @Description     gets all albums in the database
// @Tags            album
// @Accept          json
// @Produce         json
// @Success         200                {array}          Album
// @Router          /albums            [get]
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums godoc
// @Summary         Create a new album
// @Desciption      Creates a new album in the database.
// @Tags            album
// @Accept          json
// @Produce         json
// @Param           album               body            Album           true            "Album JSON Object"
// @Success         201                 {object}        Album           "OK"
// @Failure         400                 {string}        string          "BadRequest"
// @Router          /albums             [post]
func postAlbums(c *gin.Context) {
    var newAlbum Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID godoc
// @Summary         Get an album
// @Description     Gets a single album from the database corresponding to the id in the path. Otherwise, returns error
// @Tags            album
// @Accept          json
// @Produce         json
// @Param           id                  path            string          true            "Album Id"
// @Success         200                 {object}        Album           "OK"
// @Failure         404                 {string}        string          "NotFound"
// @Router          /albums/{id}        [get]
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

// updateAlbumByID godoc
// @Summary         Update an album
// @Description     Updates an existing album in the database corresponding to the id sent. Otherwise, returns error
// @Tags            album
// @Accept          json
// @Produce         json
// @Param           album               body            Album           true            "Album JSON Object"
// @Success         200                 {object}        Album           "OK"
// @Failure         404                 {string}        string          "NotFound"
// @Router          /albums             [put]
func updateAlbumByID(c *gin.Context){
    var album Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&album); err != nil {
        return
    }

    // Loop over the list of albums, looking for
    // an album whise ID value matches the parameter
    for i, a := range albums {
        if a.ID == album.ID {
            albums[i].Title = album.Title
            albums[i].Artist = album.Artist
            albums[i].Price = album.Price

            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}