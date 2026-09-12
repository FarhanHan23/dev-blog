package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DatabaseConfig struct {
	MaxConns int
}

type ServiceConfig struct {
	AppName string
	Timeout time.Duration
	DB      *DatabaseConfig
}

// challenge 2
type Post struct {
	Title       string
	IsPublished bool
}

// buat method untuk mengubah status isPublished menjadi true
func (p *Post) Publish() {
	p.IsPublished = true
}

// global config
var GlobalConfig = ServiceConfig{
	AppName: "My Blog Production",
	Timeout: 1 * time.Second,
	DB:      nil, // Skenario: DB belum di-connect / disabled
}

//	func handleCustomRequest(cfg *ServiceConfig) {
//		cfg.Timeout = 5 * time.Second
//	}
func handleCustomRequest(cfg ServiceConfig) { // tanpa pointer, data akan dicopy, sehingga perubahan tidak akan mempengaruhi data asli
	cfg.Timeout = 5 * time.Second
}

func handleCheckDB(cfg *ServiceConfig) int {
	if cfg.DB != nil {
		return cfg.DB.MaxConns
	} else {
		return 0
	}
}

// challenge 3
type PostRepoImpl struct {
	DBName string
}
type PostRepository interface {
	GetTitleByID(id int) string
}

func (repo *PostRepoImpl) GetTitleByID(id int) string {
	// implementasi query ke database untuk mendapatkan title berdasarkan id
	return fmt.Sprintf("Title for ID %d", id)
}

// challenge 4
type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// implementasi untuk membuat post baru
	var req CreatePostRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// misal kita berhasil membuat post baru
	response := APIResponse{
		Status:  "success",
		Message: "Post created successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// challenge 5
func FetchPostFromDB(id int) (string, error) {
	defer fmt.Println("DEBUG: Koneksi database berhasil ditutup")
	// implementasi untuk mengambil post dari database

	if id <= 0 {
		return "", fmt.Errorf("Invalid post ID: %d", id)
	}
	return fmt.Sprintf("Judul Artikel Artikel  Id %d", id), nil
}

func main() {
	fmt.Println("Timeout Awal Global:", GlobalConfig.Timeout)

	// test Request 1
	// handleCustomRequest(&GlobalConfig)
	// fmt.Println("Timeout Global setelah Request 1:", GlobalConfig.Timeout)

	handleCustomRequest(GlobalConfig)
	fmt.Println("Timeout Global setelah Request 1:", GlobalConfig.Timeout)

	// test Request 2
	maxConn := handleCheckDB(&GlobalConfig)
	fmt.Println("Max Conns DB:", maxConn)

	// challenge 2 test
	var Post1 = Post{
		Title:       "Belajar Golang",
		IsPublished: false,
	}
	Post1.Publish()
	fmt.Println("Post 1 Published:", Post1.IsPublished) // berubah jadi true pass by reference karena methodnya menggunakan pointer receiver

	// challenge 3 test
	var repo PostRepository = &PostRepoImpl{
		DBName: "MyBlogDB",
	}
	title := repo.GetTitleByID(1)
	fmt.Println("Title dari PostRepository:", title)

	// challenge 5 test
	FetchPostFromDB(-1) // akan memicu error dan menampilkan debug message
	FetchPostFromDB(1)  // akan berhasil dan menampilkan debug message

	// challenge 4 test
	http.HandleFunc("/posts", CreatePostHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
