# gin_gonic_tutorial
In this tutorial, you will build a RESTful API server with two endpoints. Your example project will be a repository of data about vintage jazz records.

## Prerequisites
An installation of Go 1.16 or later.

## API endpoints
/albums<br>

GET – Get a list of all albums, returned as JSON.<br>
POST – Add a new album from request data sent as JSON.<br><br>
PUT - Updates an existing album by its ID from request data sent as JSON.<br><br>

/albums/:id<br>

GET – Get an album by its ID, returning the album data as JSON.<br>

# References
* https://go.dev/doc/tutorial/web-service-gin