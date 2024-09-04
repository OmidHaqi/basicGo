package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {

	response, err := http.Get("https://dummyjson3123.com/products/categories")

	if err != nil {
		println("an error has occurred on get request")
		return
	}

	println(response.Status)

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		println("an error has occurred on reading the response body")
	}

	println(string(responseBody))

}


//1

type HttpError struct {
	StatusCode int
	Message    string
}


func customErr() {

	response, err := GetRequest("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)

}

func (error HttpError) Error() string {
	return fmt.Sprintf("Http error occurred: %d %s", error.StatusCode, error.Message)
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{StatusCode: statusCode, Message: message}
}

func GetRequest(url string) (string, error) {
	if len(url) == 0 {
		return "", NewHttpError(400, "Url is empty")
	}
	response, err := http.Get(url)
	if err != nil {
		return "", NewHttpError(500, "Error occurred")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", NewHttpError(200, "Body is empty")
	}
	return string(responseBody), nil
}
//2
func createErr() {
	output, err := createErrorMethod2(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

}

func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("Number is not valid")
	}
	return number * 5, nil
}

func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("Number is not valid: %d", number)
	}
	return number * 5, nil
}