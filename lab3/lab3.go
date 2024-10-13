package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	path := r.URL.Path
	info := strings.Split(path, "/")
	if len(info) != 4 {
		w.Write([]byte("Error!"))
		return
	}
	num1, err3 := strconv.Atoi(info[2])
	num2, err2 := strconv.Atoi(info[3])
	if err3 != nil || err2 != nil {
		w.Write([]byte("Error!"))
	} else {
		switch info[1] {
		case "add":
			w.Write([]byte(fmt.Sprintf("%d + %d = %d", num1, num2, num1+num2)))
		case "sub":
			w.Write([]byte(fmt.Sprintf("%d - %d = %d", num1, num2, num1-num2)))
		case "mul":
			w.Write([]byte(fmt.Sprintf("%d * %d = %d", num1, num2, num1*num2)))
		case "div":
			if num2 == 0 {
				w.Write([]byte("Error!"))
				return
			}
			w.Write([]byte(fmt.Sprintf("%d / %d = %d, remainder %d", num1, num2, num1/num2, num1%num2)))
		default:
			w.Write([]byte("Error!"))

		}
	}

}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
