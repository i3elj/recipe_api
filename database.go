package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Ingredient struct {
	What   string
	Amount string
}

type Recipe struct {
	Id          uint
	Name        string
	Ingredients []Ingredient
	Image_url   string
	Preparation string
}

type DB struct {
	recipes []Recipe
}

func create_new_recipe() Recipe {
	return Recipe{
		Id: db_get_min_id(),
	}
}

func db_get_min_id() uint {
	recipes := initDB()

	var min uint = 0

	for i := 1; i < len(recipes)+1; i++ {
		if recipes[i].Id != uint(i) {
			min = uint(i)
			break
		}
	}

	return min
}

func initDB() []Recipe {
	content, err := ioutil.ReadFile("./db.json")

	if err != nil {
		log.Fatal("For some reason we couldn't read the json file:\n\t", err)
	}

	var json_data []Recipe
	err = json.Unmarshal(content, &json_data)

	if err != nil {
		log.Fatal("Unable to unmarsh the json file:\n\t", err)
	}

	return json_data
}

func add_recipe_to_db(recipe Recipe) {
	recipes := initDB()
	recipes = append(recipes, recipe)

	updated_content, err := json.Marshal(recipes)

	if err != nil {
		log.Fatal("Something is wrong with your JSON:\n\t", err)
	}

	ioutil.WriteFile("./db.json", updated_content, os.ModePerm)
}

func delete_recipe_from_db(id uint) {
	recipes := initDB()

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == id {
			recipes = append(recipes[:i], recipes[i+1:]...)
			break
		}
	}

	updated_content, err := json.Marshal(recipes)

	if err != nil {
		log.Fatal("Something is wrong with your JSON:\n\t", err)
	}

	ioutil.WriteFile("./db.json", updated_content, os.ModePerm)
}

func update_recipe_in_db(recipe Recipe) {
	recipes := initDB()

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == recipe.Id {
			recipes[i] = recipe
			break
		}
	}

	updated_content, err := json.Marshal(recipes)

	if err != nil {
		log.Fatal("Something is wrong with your JSON:\n\t", err)
	}

	ioutil.WriteFile("./db.json", updated_content, os.ModePerm)
}
