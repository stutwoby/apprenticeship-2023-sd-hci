package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Review struct {
	CustomerName string    `json:"customer_name"`
	TimeSent     time.Time `json:"time_sent"`
	Item         string    `json:"item"`
	ReviewText   string    `json:"review_text"`
}

type Menu struct {
	Menu []MenuItem `json:"menu"`
}

var db *sql.DB
var store = sessions.NewCookieStore([]byte("your-secret-key"))

type PageVariables struct {
	FoodMenu      []MenuItem
	DrinkMenu     []MenuItem
	FoodMenuJSON  string
	DrinkMenuJSON string
}

func main() {
	var err error
	db, err = sql.Open("mysql", "devuser:123456@tcp(127.0.0.1:3306)/kebabshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("html/static"))))

	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "user-session")
			username, ok := session.Values["username"].(string)
			if !ok || username == "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	//PAGES
	r.Handle("/home", authMiddleware(http.HandlerFunc(homeHandler))).Methods("GET")
	r.Handle("/profile", authMiddleware(http.HandlerFunc(profileHandler))).Methods("GET")
	r.Handle("/order", authMiddleware(http.HandlerFunc(orderHandler))).Methods("GET")
	r.Handle("/restaurant", authMiddleware(http.HandlerFunc(restaurantHandler))).Methods("GET")
	r.Handle("/reviews", authMiddleware(http.HandlerFunc(reviewsHandler))).Methods("GET")

	//NAV BAR SIGN OUT
	r.Handle("/signout", authMiddleware(http.HandlerFunc(signoutHandler))).Methods("GET")

	//PROFILE EDIT DETAILS
	r.Handle("/editdetails", authMiddleware(http.HandlerFunc(editDetailsHandler))).Methods("POST")
	r.Handle("/changepassword", authMiddleware(http.HandlerFunc(changePasswordHandler))).Methods("POST")
	r.Handle("/updateprofile", authMiddleware(http.HandlerFunc(updateProfileHandler))).Methods("POST")
	r.Handle("/signoutafterupdate", authMiddleware(http.HandlerFunc(signoutHandlerAfterUpdate))).Methods("GET")

	//LOGIN AND SIGN UP
	r.HandleFunc("/", loginHandler).Methods("GET")
	r.HandleFunc("/", loginPostHandler).Methods("POST")
	r.HandleFunc("/signup", signupHandler).Methods("POST")

	//PICTURES
	http.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("pics"))))

	//REVIEWS
	r.Handle("/submitReview", authMiddleware(http.HandlerFunc(submitReviewHandler))).Methods("POST")

	http.Handle("/", r)

	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("../html/" + tmpl + ".html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "home", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func signoutHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "user-session")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var userData struct {
		Username string
		Email    string
		Address  string
	}
	err := db.QueryRow("SELECT username, email, address FROM login WHERE username = ?", username).Scan(&userData.Username, &userData.Email, &userData.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "profile", userData)
}

func restaurantHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "restaurant", nil)
}

func reviewsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	foodMenu := loadFoodMenu()

	reviewsFilePath := "../html/reviews/review.json"
	reviews, err := getReviewsFromJSONFile(reviewsFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		FoodMenu []MenuItem
		Username string
		Reviews  []Review
	}{
		FoodMenu: foodMenu,
		Username: username,
		Reviews:  reviews,
	}

	renderTemplate(w, "reviews", data)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := db.QueryRow("SELECT password FROM login WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		http.Redirect(w, r, "/?error=1", http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/?error=2", http.StatusSeeOther)
		return
	}

	session, _ := store.Get(r, "user-session")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	newUsername := r.FormValue("newUsername")
	newPassword := r.FormValue("newPassword")
	email := r.FormValue("email")
	address := r.FormValue("address")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Password hashing error", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO login (username, password, email, address) VALUES (?, ?, ?, ?)", newUsername, hashedPassword, email, address)
	if err != nil {
		http.Error(w, "Database insert error", http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, "user-session")
	session.Values["username"] = newUsername
	session.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func editDetailsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	newUsername := r.FormValue("newUsername")
	newEmail := r.FormValue("newEmail")
	newAddress := r.FormValue("newAddress")

	_, err := db.Exec("UPDATE login SET username=?, email=?, address=? WHERE username=?", newUsername, newEmail, newAddress, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {

		http.Error(w, "Not authenticated", http.StatusUnauthorized)
		return
	}

	currentPassword := r.FormValue("currentPassword")
	newPassword := r.FormValue("newPassword")

	var storedPassword string
	err := db.QueryRow("SELECT password FROM login WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(currentPassword))
	if err != nil {

		http.Error(w, "Current password is incorrect", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {

		http.Error(w, "Password hashing error", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("UPDATE login SET password=? WHERE username=?", hashedPassword, username)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))
}

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		newUsername := r.FormValue("newUsername")
		newEmail := r.FormValue("newEmail")
		newAddress := r.FormValue("newAddress")

		_, err := db.Exec("UPDATE login SET username = ?, email = ?, address = ? WHERE username = ?", newUsername, newEmail, newAddress, username)
		if err != nil {

			fmt.Println("Error updating user data:", err)
			http.Error(w, "Failed to update profile", http.StatusInternalServerError)
			return
		}

		fmt.Println("User profile updated successfully")

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func signoutHandlerAfterUpdate(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func loadData() PageVariables {
	foodMenu := loadMenu("menus/food_menu.json")
	drinkMenu := loadMenu("menus/drink_menu.json")

	return PageVariables{
		FoodMenu:      foodMenu.Menu,
		DrinkMenu:     drinkMenu.Menu,
		FoodMenuJSON:  convertToJSON(foodMenu.Menu),
		DrinkMenuJSON: convertToJSON(drinkMenu.Menu),
	}
}

func loadMenu(filename string) *Menu {
	var menu Menu

	data, err := os.ReadFile(filename)
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

func convertToJSON(data []MenuItem) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	data := loadData()
	renderTemplate(w, "order", data)
}

func submitReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Review struct {
		CustomerName string `json:"customer_name"`
		Item         string `json:"item"`
		ReviewText   string `json:"review_text"`
	}

	var review Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reviewsFilePath := "../html/reviews/review.json"

	reviewsFile, err := os.OpenFile(reviewsFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer reviewsFile.Close()

	reviewData, err := json.Marshal(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = reviewsFile.WriteString(string(reviewData) + "\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Review submitted successfully"))
}

func loadFoodMenu() []MenuItem {
	foodMenu := loadMenu("menus/food_menu.json")
	if foodMenu != nil {
		return foodMenu.Menu
	}
	return nil
}

func getReviewsFromJSONFile(filePath string) ([]Review, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reviews []Review
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var review Review
		err := json.Unmarshal([]byte(scanner.Text()), &review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
