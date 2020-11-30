package chat

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//HTTPService is ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//NewHTTPTransport is...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/producto",
		function: getAll(s),
	}, &endpoint{
		method:   "POST",
		path:     "/producto",
		function: addProducto(s),
	}, &endpoint{
		method:   "GET",
		path:     "/producto/:id",
		function: getByID(s),
	}, &endpoint{
		method:   "PUT",
		path:     "/producto",
		function: updateProducto(s),
	}, &endpoint{
		method:   "DELETE",
		path:     "/producto/:id",
		function: delProducto(s),
	})

	return list
}

func updateProducto(s Service) gin.HandlerFunc {
	var p Producto
	return func(c *gin.Context) {
		c.BindJSON(&p)
		result, err := s.UpdateProducto(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"producto": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"producto": result,
			})
		}
	}
}

func getByID(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.FindByID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"producto": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"producto": result,
			})
		}
	}
}

func delProducto(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, err := s.RemoveByID(ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"producto": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"producto": result,
			})
		}
	}
}

func addProducto(s Service) gin.HandlerFunc {
	var p Producto
	return func(c *gin.Context) {
		c.BindJSON(&p)
		result, err := s.AddProducto(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"producto": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"producto": result,
			})
		}
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"producto": s.FindAll(),
		})
	}
}

func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
