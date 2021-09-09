package main

import (
	"fmt"
	"log"
)

type ResponseError struct {
	code int
	msg  string
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("reponse failed, cdoe:%d msg:%s", re.code, re.msg)
}

func testError(str string) (string, error) {
	if str == "error" {
		return "error", &ResponseError{
			code: 500,
			msg:  "bad request",
		}
	}
	return "success", nil
}

func main() {
	s, err := testError("error")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(s)
}
