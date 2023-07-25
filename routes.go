package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	r.GET("/api/recipes", get_recipes_with_search)
	r.GET("/api/recipe/:id", get_recipe_by_id)
	r.DELETE("/api/recipe/:id", delete_recipe_by_id)
	r.PUT("/api/recipe/:id", update_recipe_by_id)
	r.POST("/api/recipe", post_recipe)
}

/*
API endpoint: /api/recipe

Method: GET

Examples:

  - /api/recipe/1
  - /api/recipe/230
*/
func get_recipe_by_id(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal("We couldn't convert your string. Sorry :( :\n\t", err)
	}

	var recipe Recipe
	found_it := false

	for i := 0; i < len(Db); i++ {
		if Db[i].Id == uint(id) {
			recipe = Db[i]
			found_it = true
			break
		}
	}

	if found_it {
		c.JSON(http.StatusOK, gin.H{"recipe": recipe})
	} else {
		c.JSON(http.StatusOK, gin.H{"not found": "404 :("})
	}
}

/*
API endpoint: /api/recipes

Method: GET

Examples:

  - /api/recipes
  - ?name=Pão Caseiro
  - ?ingredients=ovo,leite,açúcar
*/
func get_recipes_with_search(c *gin.Context) {
	name := c.Query("name")
	ingredients := strings.Split(c.Query("ingredients"), ",")

	if name == "" && len(ingredients) <= 1 {
		recipes := initDB()

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, recipes)
		return
	}

	var results []Recipe

	for _, recipe := range Db {
		if recipe.Name == name {
			results = append(results, recipe)
		}

		for _, ingredient := range ingredients {
			if has_ing(recipe.Ingredients, ingredient) {
				results = append(results, recipe)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

/*
API endpoint: /api/recipe

Method: POST

Examples:

  - {"name": "New Recipe Name", "ingredients": [ {"what": "what to use",   "amount": "how much use it"} ], "preparation": "way of baking" }
*/
func post_recipe(c *gin.Context) {
	req := create_new_recipe()

	c.Bind(&req)

	add_recipe_to_db(req)

	c.JSON(http.StatusOK, gin.H{
		"recipe": req,
	})
}

/*
API endpoint: /api/recipe/:id

Method: DELETE

Examples:

  - /api/recipe/3
*/
func delete_recipe_by_id(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal("We couldn't convert your string. Sorry :( :\n\t", err)
	}

	delete_recipe_from_db(uint(id))
}

/*
API endpoint: /api/recipe/:id

Method: PUT

examples:

  - {"name": "new recipe name", "ingredients": [ {"what": "updated ingredients", "amount": "new amount"} ], "preparation": "new way of baking" }
*/
func update_recipe_by_id(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	fmt.Println(id)

	if err != nil {
		log.Fatal("We couldn't convert your string. Sorry :( :\n\t", err)
	}

	var recipe Recipe
	recipe.Id = uint(id)

	c.Bind(&recipe)

	update_recipe_in_db(recipe)

	c.JSON(http.StatusOK, gin.H{
		"recipe": recipe,
	})
}
