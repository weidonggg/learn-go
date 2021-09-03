package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(p, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}

var (
	h  = "$2y$10$iHAKymQtPTJtxRoF5ozWC.BHM5lwy6P6/0T4ZliPkM1jUxwfTH4ay"
	h2 = "$2y$10$tv.VSx1zb/7OEwf5PD1.IuAhhmCd0DKh43KzQ58YHEIPwTcin97p6"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, r.ContentLength)

	r.Body.Read(buf)

	fmt.Println(r.Header)
	fmt.Println(string(buf))

	io.WriteString(w, "ok")

}

func main() {

	// h := hashWithSalted("5mMFL8Ug")
	// fmt.Println(h)
	// pass := "1kDix6qg"

	// // p, _ := PasswordHash(pass)
	// // fmt.Println(p)

	// match := PasswordVerify(pass, h2)
	// fmt.Println("验证：", match)
	http.HandleFunc("/send", webhook)

	http.ListenAndServe(":8888", nil)

}
