package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	Username string
	Pin      string
	Saldo    int
}

var db *sql.DB

type MainMenu struct {
	User *User
}

func (menu *MainMenu) display() {
	for {
		clearScreen()
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("       VIRTUAL ATM")
		fmt.Println(strings.Repeat("-", 40))

		fmt.Println("Menu:")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Top Up")
		fmt.Println("3. Tarik Tunai")
		fmt.Println("4. Ubah PIN")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			menu.CekSaldo()
		case 2:
			menu.TopUp()
		case 3:
			menu.TarikTunai()
		case 4:
			menu.ChangePin()
		case 5:
			fmt.Println("Terima kasih telah menggunakan layanan kami!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}

func (menu *MainMenu) CekSaldo() {
	clearScreen()
	fmt.Printf("Halo %s!\n", menu.User.Username)
	fmt.Printf("Saldo Anda: Rp%d\n", menu.User.Saldo)
	fmt.Print("\nTekan 0 untuk kembali ke menu utama: ")
	var input string
	fmt.Scan(&input)
	if input != "0" {
		fmt.Println("Input tidak valid. Menutup aplikasi.")
		os.Exit(1)
	}
}

func (menu *MainMenu) TopUp() {
	clearScreen()
	fmt.Print("Masukkan jumlah top up: ")
	var amount int
	fmt.Scan(&amount)
	menu.User.Saldo += amount

	_, err := db.Exec("UPDATE users SET saldo=$1 WHERE username=$2", menu.User.Saldo, menu.User.Username)
	if err != nil {
		fmt.Println("Error updating saldo:", err)
		return
	}

	menu.logTransaction(amount, "topup")
	fmt.Println("Top up berhasil!")
	fmt.Print("\nTekan 0 untuk kembali ke menu utama: ")
	var input string
	fmt.Scan(&input)
	if input != "0" {
		fmt.Println("Input tidak valid. Menutup aplikasi.")
		os.Exit(1)
	}
}

func (menu *MainMenu) TarikTunai() {
	clearScreen()
	fmt.Print("Masukkan jumlah tarik tunai: ")
	var amount int
	fmt.Scan(&amount)

	if amount > menu.User.Saldo {
		fmt.Println("Saldo tidak cukup!")
	} else {
		menu.User.Saldo -= amount
		_, err := db.Exec("UPDATE users SET saldo=$1 WHERE username=$2", menu.User.Saldo, menu.User.Username)
		if err != nil {
			fmt.Println("Error updating saldo:", err)
			return
		}
		menu.logTransaction(amount, "tarik")
		fmt.Println("Tarik tunai berhasil!")
	}

	fmt.Print("\nTekan 0 untuk kembali ke menu utama: ")
	var input string
	fmt.Scan(&input)
	if input != "0" {
		fmt.Println("Input tidak valid. Menutup aplikasi.")
		os.Exit(1)
	}
}

func (menu *MainMenu) ChangePin() {
	clearScreen()
	fmt.Print("Masukkan PIN baru: ")
	var newPin string
	fmt.Scan(&newPin)

	_, err := db.Exec("UPDATE users SET pin=$1 WHERE username=$2", newPin, menu.User.Username)
	if err != nil {
		fmt.Println("Error updating PIN:", err)
		return
	}

	fmt.Println("PIN berhasil diubah!")
	fmt.Print("\nTekan 0 untuk kembali ke menu utama: ")
	var input string
	fmt.Scan(&input)
	if input != "0" {
		fmt.Println("Input tidak valid. Menutup aplikasi.")
		os.Exit(1)
	}
}

func (menu *MainMenu) logTransaction(amount int, transactionType string) {
	_, err := db.Exec(
		"INSERT INTO transactions (username, amount, transaction_type) VALUES ($1, $2, $3)",
		menu.User.Username, amount, transactionType,
	)
	if err != nil {
		fmt.Println("Error logging transaction:", err)
	}
}

func authenticate(username, pin string) *User {
	var user User
	err := db.QueryRow("SELECT username, pin, saldo FROM users WHERE username=$1 AND pin=$2", username, pin).Scan(&user.Username, &user.Pin, &user.Saldo)
	if err != nil {
		return nil
	}
	return &user
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var err error
	// Load DATABASE_URL from environment
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		fmt.Println("Error: DATABASE_URL environment variable not set")
		return
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Selamat datang di Virtual ATM!")
	for {
		fmt.Print("Masukkan username: ")
		var username string
		fmt.Scan(&username)

		fmt.Print("Masukkan PIN: ")
		var pin string
		fmt.Scan(&pin)

		user := authenticate(username, pin)
		if user != nil {
			menu := MainMenu{User: user}
			menu.display()
		} else {
			fmt.Println("Username atau PIN salah. Coba lagi.")
		}
	}
}
