package main

import (
	"flag"
	"net/http"

	"github.com/gireeshramji/food-ui/config"
	"github.com/gireeshramji/food-ui/db"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	dirty bool // Need to make sure this is threadsafe but it is not for now
)

func setupRouter(config *config.Config) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.HTMLRender = createRender()

	recipeDB := db.ParseRecipeDb(config)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"sources": recipeDB,
			"dirty":   dirty,
		})
	})

	return router

}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/base.tmpl", "templates/index.tmpl")

	return r
}

func main() {

	cfgFile := flag.String("conf", "cfg.toml", "Filename of food-ui toml cfg file")
	flag.Parse()
	cfg, _ := config.ParseConfig(*cfgFile)
	r := setupRouter(cfg)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
