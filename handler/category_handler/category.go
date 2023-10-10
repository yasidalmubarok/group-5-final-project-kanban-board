package category_handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/service/category_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category_service.CategoryService
}

func (ch *categoryHandler) Create(ctx *gin.Context) {
	var newCategoryRequest = &dto.NewCategoryRequest{}

	if err := ctx.ShouldBindJSON(&newCategoryRequest); err != nil{
		errBindJson := errs.NewUnprocessibleEntityError("invalid request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := ch.categoryService.Create(newCategoryRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}