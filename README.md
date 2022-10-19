# `templateManager` Gin Integration

This is the package for the [`templateManager`](https://github.com/paul-norman/go-template-manager) renderer integration for the [Gin](https://gin-gonic.com/) framework.

For all options, please see the main repository.

## Basic Usage

```go
package main

import (
	"github.com/gin-gonic/gin"

	TM "example/website/templateManagerGin"
)

func main() {
	router := gin.Default()

	engine := TM.Init("templates", ".html")
	engine.ExcludeDirectories([]string{"layouts", "partials"})
	if gin.IsDebugging() {
		engine.Reload(true)
		engine.Debug(true)
	}
	engine.Parse()

	router.HTMLRender = engine

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", gin.H{
			"Title": "Home",
		})
	})

	router.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test.html", gin.H{
			"Title": "Test",
		})
	})

	router.Run()
}
```