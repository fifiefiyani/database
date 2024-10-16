package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Struct untuk User
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"` // Admin atau Editor
}

// Struct untuk Todo
type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`  // Pending, Completed, etc.
	UserID      uint   `json:"user_id"` // Foreign key ke tabel User
}

var DB *gorm.DB

// Inisialisasi koneksi ke database
func InitDB() {
	// Ubah sesuai konfigurasi MariaDB Anda
	dsn := "root:password123@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MariaDB:", err)
	}
	log.Println("Connected to MariaDB successfully")

	// Migrasi tabel User dan Todo
	DB.AutoMigrate(&User{}, &Todo{})
}

// Handler untuk mengambil semua user
func getUsers(c echo.Context) error {
	var users []User
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// Handler untuk membuat user baru
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

// Handler untuk mengambil semua todo
func getTodos(c echo.Context) error {
	var todos []Todo
	DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

// Handler untuk membuat todo baru
func createTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func main() {
	// Inisialisasi koneksi database
	InitDB()

	// Inisialisasi Echo
	e := echo.New()

	// Definisikan route untuk CRUD User
	e.GET("/users", getUsers)
	e.POST("/users", createUser)

	// Definisikan route untuk CRUD Todo
	e.GET("/todos", getTodos)
	e.POST("/todos", createTodo)

	// Jalankan server di port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
