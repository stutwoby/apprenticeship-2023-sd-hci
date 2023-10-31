package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Menu struct {
	MenuItems []MenuItem `json:"menu"`
}

type Order struct {
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	Address       string `json:"address"`
	FoodOrder     []string
	DrinkOrder    []string
}

type PageVariables struct {
	FoodMenu      []MenuItem
	DrinkMenu     []MenuItem
	FoodMenuJSON  string
	DrinkMenuJSON string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := loadData() // Load data from JSON files
			tmpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, data)
		}
	})

	http.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("pics"))))

	http.ListenAndServe(":8080", nil)
}

func loadData() PageVariables {
	// Load menu items from JSON files
	foodMenu := loadMenu("menus/food_menu.json")
	drinkMenu := loadMenu("menus/drink_menu.json")

	return PageVariables{
		FoodMenu:      foodMenu.MenuItems,
		DrinkMenu:     drinkMenu.MenuItems,
		FoodMenuJSON:  convertToJSON(foodMenu.MenuItems),
		DrinkMenuJSON: convertToJSON(drinkMenu.MenuItems),
	}
}

func loadMenu(filename string) *Menu {
	var menu Menu

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading menu file: %v\n", err)
		return nil
	}

	err = json.Unmarshal(data, &menu)
	if err != nil {
		fmt.Printf("Error parsing menu JSON: %v\n", err)
		return nil
	}

	return &menu
}

func convertToJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}
