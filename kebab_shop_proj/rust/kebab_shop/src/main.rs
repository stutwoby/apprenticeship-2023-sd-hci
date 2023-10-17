use std::io;
use std::fs::File;
use std::io::Read;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone)]
struct MenuItem {
    name: String,
    price: f64,
}

#[derive(Serialize, Deserialize)]
struct Menu {
    menu: Vec<MenuItem>,
}

#[derive(Serialize, Deserialize)]
struct OrderItem {
    item: MenuItem,
    sauces: Vec<String>,
}

#[derive(Serialize, Deserialize)]
struct Order {
    customer_name: String,
    customer_phone: String,
    address: String,
    food_order: Vec<OrderItem>,
    drink_order: Vec<OrderItem>,
}

fn main() {
    println!("Dear Sir/Madam,");
    println!("please follow the prompts below to place an order :)");

    let mut order = Order {
        customer_name: String::new(),
        customer_phone: String::new(),
        address: String::new(),
        food_order: Vec::new(),
        drink_order: Vec::new(),
    };

    let mut input = String::new();
    let mut show_food_menu = String::new();

    print!("Customer Name: ");
    io::stdin().read_line(&mut input).expect("Failed to read line");
    order.customer_name = input.trim().to_string();
    input.clear();

    print!("Customer Phone: ");
    io::stdin().read_line(&mut input).expect("Failed to read line");
    order.customer_phone = input.trim().to_string();
    input.clear();

    print!("Delivery Address: ");
    io::stdin().read_line(&mut input).expect("Failed to read line");
    order.address = input.trim().to_string();
    input.clear();

    print!("Do you want to see the food menu? (yes/no): ");
    io::stdin().read_line(&mut show_food_menu).expect("Failed to read line");
    show_food_menu = show_food_menu.trim().to_string();

    if show_food_menu.trim().to_lowercase() == "yes" {
        let food_menu = load_menu("menus/food_menu.json");

        println!("Food Menu:");
        for (i, item) in food_menu.menu.iter().enumerate() {
            println!("{}. {} - ${:.2}", i + 1, item.name, item.price);
        }

        loop {
            let mut choice = String::new();
            print!("Enter the item number you want to order (or 0 to finish ordering food): ");
            io::stdin().read_line(&mut choice).expect("Failed to read line");
            let choice = choice.trim().parse::<usize>().expect("Invalid input");

            if choice == 0 {
                break;
            }

            if choice < 1 || choice > food_menu.menu.len() {
                println!("Invalid item number. Please try again.");
                continue;
            }

            let mut sauces = Vec::new();
            loop {
                let mut sauce = String::new();
                print!("Enter a sauce for this item (or 'done' to finish sauces for this item): ");
                io::stdin().read_line(&mut sauce).expect("Failed to read line");
                let sauce = sauce.trim().to_string();

                if sauce == "done" {
                    break;
                }

                sauces.push(sauce);
            }

            let order_item = OrderItem {
                item: food_menu.menu[choice - 1].clone(),
                sauces,
            };
            order.food_order.push(order_item);
        }
    }

    let mut show_drink_menu = String::new();
    print!("Do you want to see the drink menu? (yes/no): ");
    io::stdin().read_line(&mut show_drink_menu).expect("Failed to read line");
    show_drink_menu = show_drink_menu.trim().to_string();

    if show_drink_menu.trim().to_lowercase() == "yes" {
        let drink_menu = load_menu("menus/drink_menu.json");

        println!("Drink Menu:");
        for (i, item) in drink_menu.menu.iter().enumerate() {
            println!("{}. {} - ${:.2}", i + 1, item.name, item.price);
        }

        loop {
            let mut choice = String::new();
            print!("Enter the item number you want to order (or 0 to finish ordering drinks): ");
            io::stdin().read_line(&mut choice).expect("Failed to read line");
            let choice = choice.trim().parse::<usize>().expect("Invalid input");

            if choice == 0 {
                break;
            }

            if choice < 1 || choice > drink_menu.menu.len() {
                println!("Invalid item number. Please try again.");
                continue;
            }

            let order_item = OrderItem {
                item: drink_menu.menu[choice - 1].clone(),
                sauces: Vec::new(),
            };
            order.drink_order.push(order_item);
        }
    }

    println!("\nOrder Receipt:");
    println!("Customer Name: {}", order.customer_name);
    println!("Customer Phone: {}", order.customer_phone);
    println!("Delivery Address: {}", order.address);

    if !order.food_order.is_empty() {
        println!("\nOrdered Food Items:");
        let mut total_food = 0.0;
        for (i, item) in order.food_order.iter().enumerate() {
            println!("{}. {} - ${:.2}", i + 1, item.item.name, item.item.price);
            println!("   Sauces:");
            for sauce in &item.sauces {
                println!("   - {}", sauce);
            }
            total_food += item.item.price;
        }
        println!("Total Food Cost: ${:.2}", total_food);
    }

    if !order.drink_order.is_empty() {
        println!("\nOrdered Drink Items:");
        let mut total_drinks = 0.0;
        for (i, item) in order.drink_order.iter().enumerate() {
            println!("{}. {} - ${:.2}", i + 1, item.item.name, item.item.price);
            total_drinks += item.item.price;
        }
        println!("Total Drink Cost: ${:.2}", total_drinks);
    }

    let mut total = 0.0;
    total += calculate_total(&order.food_order);
    total += calculate_total(&order.drink_order);
    println!("Total: {:.2}", total);

    print!("Is this order correct? (yes/no): ");
    let mut confirmation = String::new();
    io::stdin().read_line(&mut confirmation).expect("Failed to read line");

    if confirmation.trim().to_lowercase() != "yes" {
        loop {
            print!("Do you want to change your order? (yes/no): ");
            let mut change_order = String::new();
            io::stdin().read_line(&mut change_order).expect("Failed to read line");

            if change_order.trim().to_lowercase() == "no" {
                break;
            }
        }
    }

    println!("Many thanks");
    println!("Litva Kebabs Ltd");
}

fn load_menu(filename: &str) -> Menu {
    let mut file = File::open(filename).expect("Failed to open file");
    let mut menu_json = String::new();
    file.read_to_string(&mut menu_json).expect("Failed to read file");

    serde_json::from_str(&menu_json).expect("Error parsing menu JSON")
}

fn calculate_total(order_items: &[OrderItem]) -> f64 {
    order_items.iter().map(|item| item.item.price).sum()
}
