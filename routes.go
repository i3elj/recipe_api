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
	r.POST("/api/recipe", post_recipe)
	r.DELETE("/api/recipe/:id", delete_recipe_by_id)
	r.PUT("/api/recipe/:id", update_recipe_by_id)
}

/*
API endpoint: /api/recipes

Method: GET

Examples:

	/api/recipes
	?name=Pão Caseiro
	?ingredients=ovo,leite,açúcar
*/
func get_recipes_with_search(c *gin.Context) {
	name := c.Query("name")
	ingredients := c.Query("ingredients")

	name_is_not_set := name == ""
	ingredients_is_not_set := ingredients == ""

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	recipes := initDB()

	if name_is_not_set && ingredients_is_not_set {
		c.JSON(http.StatusOK, recipes)
		return
	}

	var results []Recipe

	for _, recipe := range recipes {
		if recipe.Name == name {
			results = append(results, recipe)
		}

		for _, ingredient := range strings.Split(ingredients, ",") {
			if has_ing(recipe.Ingredients, ingredient) {
				results = append(results, recipe)
			}
		}
	}

	c.JSON(http.StatusOK, results)
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

	recipes := initDB()
	var recipe Recipe
	found_it := false

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == uint(id) {
			recipe = recipes[i]
			found_it = true
			break
		}
	}

	if found_it {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, recipe)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"not found": "404 :("})
	}
}

/*
API endpoint: /api/recipe/img/:url

Method: GET

Examples:

  - /api/recipe/img/image_number_one.png
*/
// func get_recipes_image(c *gin.Context) {
// 	image_name := c.Query("name")
// 	gin.Default().Static("./assets/"+image_name, "./assets/")
// }

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
