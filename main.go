package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

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

func getToken() {
	getTokenUrl := "https://oapi.dingtalk.com/gettoken?appkey=dingvshzpnithdjocvz6&appsecret=qRf0Zclbiq6Q3sqNcu8wqFkpJJeMiqgmuQtpa08uEVRggH6k2lvu-lqjG_TLyVJC"

	// getUser := "https://oapi.dingtalk.com/user/get?userid=m0isr85?access_token=951cf5c4f5ba357084d33ccedd368296"
	// // m0isr85
	resp, err := http.Get(getTokenUrl)
	if err != nil {
		fmt.Println("request:", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read", err)
		return
	}
	fmt.Println(string(data))
}

func main() {

	// h := hashWithSalted("5mMFL8Ug")
	// fmt.Println(h)
	// pass := "1kDix6qg"

	// // p, _ := PasswordHash(pass)
	// // fmt.Println(p)

	// match := PasswordVerify(pass, h2)
	// fmt.Println("验证：", match)
	// http.HandleFunc("/send", webhook)
	// http.HandleFunc("/robots", webhook)
	// http.ListenAndServe(":8080", nil)

	// GET https://oapi.dingtalk.com/gettoken?appkey=appkey&appsecret=appsecret

	// getToken()

	// body := `{
	// 	"robot_code": "sJNNs9ixxxx",
	// 	"target_open_conversation_id": "cideGsbg4nGkkrhtpTqCLhH0A==",
	// 	"msg_template_id": "offical_template_test_action_card",
	// 	"at": {
	// 		"atMobiles": [
	// 			"17781698985"
	// 		],
	// 		"isAtAll": false
	// 	},
	// 	"text": {
	// 		"content": "我就是我, @180xxxxxx 是不一样的烟火"
	// 	},
	// 	"msgtype": "text"
	// }`

	// resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token=367bfcf9377527af8cee7c7547cf6a823407e57c8742fcb4b1d705bfb583e4db", "application/json", strings.NewReader(body))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))
	// m1 := [3]int{1, 2, 3}
	// mapValue(m1)
	// fmt.Println(m1)
	// revertV2("Hello world")
	// f1()

	// a := adder()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(a(i))
	// }

	// a := adder2(0)
	// for i := 0; i < 10; i++ {
	// 	var s int
	// 	s, a = a(i)
	// 	fmt.Println(s)

	// }

	deferDemo()

}

func deferDemo() {
	file, err := os.Open("test.txt")
	defer fmt.Println("Call defer")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = ioutil.ReadAll(file)

}

func adder() func(int) int {
	sum := 0 // 自由变量
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func mapValue(m [3]int) {
	m[1] = 100
}

func revert(s string) {
	var result string
	for _, v := range s {
		result = string(v) + result

	}
	fmt.Println(result)
}

func revertV2(s string) {
	// var result string
	a := []rune(s)

	// i 为字符串长度
	// 从高位开始减
	//
	for i := len(s) - 1; i > -1; i-- {
		fmt.Printf("%c", a[i])
	}
	fmt.Println()
}

func f1() {
	i := 10
	j := f2(&i)
	fmt.Printf("i: %d\nj:= %d\n", i, *j)
}

func f2(i *int) *int {
	defer func() {

		*i = 19
	}()

	return i

}
