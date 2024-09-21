package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// Initialize the MySQL database connection
func initDB() {
	dsn := "root:siddhant@tcp(127.0.0.1:3306)/Shop_data"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	fmt.Println("Database Connected")
}

// Middleware to allow CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}

// Handle user signup and save user data to the database
func signupHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle OPTIONS preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")

		// Hash the password before saving it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password:", err)
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Insert user into the database
		_, err = db.Exec("INSERT INTO users (username, phone_number, password, role) VALUES (?, ?, ?, 'user')", username, phone, hashedPassword)
		if err != nil {
			log.Println("Error inserting user into database:", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Respond with JSON success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Handle user login and authenticate based on hashed password
func loginHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Handle GET request to serve the login page
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./templates/index.html") // Serve login.html on GET
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Retrieve hashed password from the database
		var hashedPassword string
		err = db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&hashedPassword)
		if err != nil {
			log.Println("Error retrieving user from database:", err)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Compare the hashed password with the entered password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			log.Println("Password mismatch:", err)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Redirect to the shop page after successful login
		http.Redirect(w, r, "/shop.html", http.StatusSeeOther) // Redirect to shop page
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Serve shop page after login
func shopHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// Check if user is logged in (use session management or token authentication)
	// Placeholder for session check logic

	// Serve shop page
	http.ServeFile(w, r, "./templates/shop.html")
}

// Serve profile page after successful login
func profileHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	http.ServeFile(w, r, "./templates/profile.html") // Serve profile page
}

// Serve cart page
func cartHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	http.ServeFile(w, r, "./templates/cart.html") // Serve cart page
}

// Handle logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Logic to clear session or authentication token
	http.Redirect(w, r, "/index.html", http.StatusSeeOther) // Redirect to login page after logout
}

// Handle missing favicon requests
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico") // Replace with your actual favicon path
}

func main() {
	// Initialize database connection
	initDB()

	// Routes for handling signup, login, and profile
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/index", loginHandler)    // Serve login page
	http.HandleFunc("/shop.html", shopHandler) // Serve shop page
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/cart", cartHandler)      // Serve cart page
	http.HandleFunc("/logout", logoutHandler)  // Logout and redirect to login page
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Serve static files (CSS/JS)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))) // Correctly serve CSS files
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))    // Correctly serve JS files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))) // Serve image files

	// Serve HTML files
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	// Default route to serve login page when the server starts
	http.HandleFunc("/", loginHandler)

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

