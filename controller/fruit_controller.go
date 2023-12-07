package controller

import (
	"net/http"

	"github.com/sachin2605/cast0r/models"
	"github.com/sachin2605/cast0r/svc"

	"github.com/gin-gonic/gin"
)

// FruitController  FruitController
type FruitController struct {
	FruitSvc svc.FruitService
}

func NewFruitController() FruitController {
	return FruitController{
		FruitSvc: svc.NewFruitService(),
	}
}

// List : Return list of Fruits
// @Summary List all Fruits
// @Description List all Fruits
// @Tags Fruits
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Fruit
// @Router /fruits [get]
func (fc *FruitController) List(c *gin.Context) {
	var Fruits = fc.FruitSvc.List()
	c.JSON(200, Fruits)
}

// Get : Return Fruit by Id
// @Summary Return Fruit by Id
// @Description Return Fruit by Id
// @Tags Fruits
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Fruit
// @Failure 404 {object} error
// @Router /fruits/{id} [get]
func (fc *FruitController) Get(c *gin.Context) {
	id := c.Param("id")
	Fruit, err := fc.FruitSvc.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": string(err.Error()),
		})
	} else {
		c.JSON(http.StatusOK, Fruit)

	}
}

// Create : Creates new  Fruit.
// @Summary Create new Fruit record in Library
// @Description Create new Fruit record in Library
// @Tags Fruits
// @Accept  json
// @Produce  json
// @Param Fruit body models.Fruit true "Add Fruit"
// @Success 201 {object} models.Fruit
// @Router /fruits [post]
func (fc *FruitController) Create(c *gin.Context) {
	var Fruit models.Fruit

	if err := c.BindJSON(&Fruit); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		FruitID := fc.FruitSvc.Create(Fruit)
		c.JSON(http.StatusCreated, FruitID)
	}
}
