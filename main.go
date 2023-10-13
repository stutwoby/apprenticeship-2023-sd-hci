package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Menu struct {
	MenuItems []MenuItem `json:"menu"`
}

type OrderItem struct {
	Item   MenuItem `json:"item"`
	Sauces []string `json:"sauces"`
}

type Order struct {
	CustomerName  string      `json:"customer_name"`
	CustomerPhone string      `json:"customer_phone"`
	Address       string      `json:"address"`
	FoodOrder     []OrderItem `json:"food_order"`
	DrinkOrder    []OrderItem `json:"drink_order"`
}

func main() {
	fmt.Println("Dear Sir/Madam,")
	fmt.Println("I hope you are well, please follow the prompts below to place an order :)")

	order := Order{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Customer Name: ")
	order.CustomerName, _ = reader.ReadString('\n')
	order.CustomerName = order.CustomerName[:len(order.CustomerName)-1]

	fmt.Print("Customer Phone: ")
	order.CustomerPhone, _ = reader.ReadString('\n')
	order.CustomerPhone = order.CustomerPhone[:len(order.CustomerPhone)-1]

	fmt.Print("Delivery Address: ")
	order.Address, _ = reader.ReadString('\n')
	order.Address = order.Address[:len(order.Address)-1]

	fmt.Print("Do you want to see the food menu? (yes/no): ")
	showFoodMenu := "no"
	fmt.Scanln(&showFoodMenu)

	if strings.ToLower(showFoodMenu) == "yes" {
		foodMenu := loadMenu("menus/food_menu.json")

		fmt.Println("Food Menu:")
		for i, item := range foodMenu.MenuItems {
			fmt.Printf("%d. %s - $%.2f\n", i+1, item.Name, item.Price)
		}

		for {
			var choice int
			fmt.Print("Enter the item number you want to order (or 0 to finish ordering food): ")
			fmt.Scanln(&choice)

			if choice == 0 {
				break
			}

			if choice < 1 || choice > len(foodMenu.MenuItems) {
				fmt.Println("Invalid item number. Please try again.")
				continue
			}

			var sauces []string
			for {
				var sauce string
				fmt.Print("Enter a sauce for this item (or 'done' to finish sauces for this item): ")
				fmt.Scanln(&sauce)

				if sauce == "done" {
					break
				}

				sauces = append(sauces, sauce)
			}

			orderItem := OrderItem{
				Item:   foodMenu.MenuItems[choice-1],
				Sauces: sauces,
			}
			order.FoodOrder = append(order.FoodOrder, orderItem)
		}
	}

	fmt.Print("Do you want to see the drink menu? (yes/no): ")
	showDrinkMenu := "no"
	fmt.Scanln(&showDrinkMenu)

	if strings.ToLower(showDrinkMenu) == "yes" {
		drinkMenu := loadMenu("menus/drink_menu.json")

		fmt.Println("Drink Menu:")
		for i, item := range drinkMenu.MenuItems {
			fmt.Printf("%d. %s - $%.2f\n", i+1, item.Name, item.Price)
		}

		for {
			var choice int
			fmt.Print("Enter the item number you want to order (or 0 to finish ordering drinks): ")
			fmt.Scanln(&choice)

			if choice == 0 {
				break
			}

			if choice < 1 || choice > len(drinkMenu.MenuItems) {
				fmt.Println("Invalid item number. Please try again.")
				continue
			}

			orderItem := OrderItem{
				Item: drinkMenu.MenuItems[choice-1],
			}
			order.DrinkOrder = append(order.DrinkOrder, orderItem)
		}
	}

	fmt.Println("\nOrder Receipt:")
	fmt.Printf("Customer Name: %s\n", order.CustomerName)
	fmt.Printf("Customer Phone: %s\n", order.CustomerPhone)
	fmt.Printf("Delivery Address: %s\n", order.Address)

	if len(order.FoodOrder) > 0 {
		fmt.Println("\nOrdered Food Items:")
		totalFood := 0.0
		for i, item := range order.FoodOrder {
			fmt.Printf("%d. %s - $%.2f\n", i+1, item.Item.Name, item.Item.Price)
			fmt.Println("   Sauces:")
			for _, sauce := range item.Sauces {
				fmt.Printf("   - %s\n", sauce)
			}
			totalFood += item.Item.Price
		}
		fmt.Printf("Total Food Cost: $%.2f\n", totalFood)
	}

	if len(order.DrinkOrder) > 0 {
		fmt.Println("\nOrdered Drink Items:")
		totalDrinks := 0.0
		for i, item := range order.DrinkOrder {
			fmt.Printf("%d. %s - $%.2f\n", i+1, item.Item.Name, item.Item.Price)
			totalDrinks += item.Item.Price
		}
		fmt.Printf("Total Drink Cost: $%.2f\n", totalDrinks)
	}

	total := 0.0
	total += calculateTotal(order.FoodOrder)
	total += calculateTotal(order.DrinkOrder)
	fmt.Printf("Total: $%.2f\n", total)

	fmt.Print("Is this order correct? (yes/no): ")
	confirmation := ""
	fmt.Scanln(&confirmation)

	if strings.ToLower(confirmation) != "yes" {
		fmt.Print("Do you want to change your order? (yes/no): ")
		changeOrder := ""
		fmt.Scanln(&changeOrder)

		if strings.ToLower(changeOrder) == "yes" {
			main()
			return
		}
	}

	fmt.Println("Many thanks")
	fmt.Println("Litva Kebabs Ltd")
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

func calculateTotal(orderItems []OrderItem) float64 {
	total := 0.0
	for _, item := range orderItems {
		total += item.Item.Price
	}
	return total
}
