package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type uploadForm struct {
	RecipeLink      string    `form:"recipeLink" binding:"required"`
	ContainsMeat    bool      `form:"containsMeat" binding:"required"`
	ContainsSeafood string    `form:"containsSeafood" binding:"required"`
	RecipeDate      time.Time `form:"recipeDate" binding:"required"`
}

func updateModel(c *gin.Context, form *uploadForm) error {
	var err error

	if err = c.Bind(form); err != nil {
		return err
	}

	return nil
}

func bindForm(form *uploadForm, h *gin.H) {

	(*h)["recipeLink"] = form.RecipeLink
	(*h)["containsMeat"] = form.ContainsMeat
	(*h)["containsSeafood"] = form.ContainsSeafood
	(*h)["recipeDate"] = form.RecipeDate

}
