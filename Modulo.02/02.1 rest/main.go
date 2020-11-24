// Seminario GoLang - 02.1 rest
// https://youtu.be/IkOBOaVC5vg
// https://github.com/gin-gonic/gin es necesario para usar este codigo
/**
* Para correrlo es en una terminal escribir "go run main.go"
* nos genera un [GIN-debug] Listening and serving HTTP on :8080
* podemos ir a cualquier navegador y ir a http://localhost:8080/users
* y ya esta funcionando y lo estamos consumiendo si muestra
* {
* 	"status": "ok"
* }
**/

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// User is a concha de sumadre
type User struct {
	Name     string
	LastName string
}

func main() {

	r := gin.Default()
	r.GET("/users", getUsersHandler)
	r.GET("/users/:name/:ID", getUsersByNameHandler) //http://localhost:8080/users/juan/1?lastName=pololo
	r.POST("/users", postUsersHandler)
	r.POST("/form_post", postFormUsersHandler)
	r.POST("/JSON", postJSONHandler)
	r.Run()

}

/**
* Prestar suma atencion en la estructura de User tiene los atributos
* con la primer letra en mayuscula es fundamental si no no muestra nada
* cuando lo quieran imprimir
**/
func postJSONHandler(c *gin.Context) {
	requestBody := User{}
	c.Bind(&requestBody)
	user := User{requestBody.Name, requestBody.LastName}
	fmt.Println(user)
}

func postFormUsersHandler(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func postUsersHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"name":     "John",
		"lastName": "doe",
	})
}

func getUsersByNameHandler(c *gin.Context) {
	name := c.Param("name")         //para obtener el primer parametro de la URL
	ID := c.Param("ID")             //para obtener el primer parametro de la URL
	lastName := c.Query("lastName") //para obtener los campos adicionales que vienen con signos de preugnta "users/name?lastName=<valor>".
	c.JSON(200, gin.H{
		"ID":       ID,
		"name":     name,
		"lastName": lastName,
		// "name": name + " " + lastName, //otra manera de mostrar los datos
	})
}

func getUsersHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
