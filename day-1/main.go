package main

import "fmt"

type AppConfig struct {
	AppName string
	Port    int
}

// Mengubah nilai dengan Pass by Value (Data dicopy)
func updateConfigByValue(cfg AppConfig) {
	cfg.Port = 9090
	fmt.Printf("[By Value] Alamat memori di dalam fungsi: %p\n", &cfg)
}

// Mengubah nilai dengan Pass by Reference/Pointer (Mengakses memori asli)
func updateConfigByPointer(cfg *AppConfig) {
	cfg.Port = 8080
	fmt.Printf("[By Pointer] Alamat memori di dalam fungsi: %p\n", cfg)
}

func main() {
	// Inisialisasi struct config awal
	config := AppConfig{
		AppName: "My Blog API",
		Port:    3000,
	}

	fmt.Printf("[Main] Alamat memori asli: %p\n", &config)
	fmt.Println("Port Awal:", config.Port)
	fmt.Println("--------------------------------------------")

	// Tes 1: Pass by Value
	updateConfigByValue(config)
	fmt.Println("Port setelah updateByValue:", config.Port) // Nilai TIDAK berubah
	fmt.Println("--------------------------------------------")

	// Tes 2: Pass by Pointer (Gunakan tanda & untuk ambil alamat memori)
	updateConfigByPointer(&config)
	fmt.Println("Port setelah updateByPointer:", config.Port) // Nilai BERUBAH
}
