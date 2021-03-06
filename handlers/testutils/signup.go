package testutils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber"
)

type TSignupOutput struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type TSignInputs struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var TSignupInputsVal TSignInputs

func init() {
	TSignupInputsVal = TSignInputs{
		Email:    "kiran@gmail.com",
		Password: "password",
		UserName: "kiran",
	}
}

func TSignup(app *fiber.App, in ...TSignInputs) (resp *http.Response, inputs TSignInputs, data TSignupOutput) {
	inputs = TSignupInputsVal

	if len(in) > 0 {
		inputs = in[0]
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&inputs)

	if err != nil {
		log.Fatal(err)
	}

	// http.Request
	req := httptest.NewRequest(
		"POST",
		"/api/v1/signup",
		buf,
	)
	req.Header.Set("Content-Type", "application/json")

	resp, _ = app.Test(req, -1)

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		panic(err)
	}

	return resp, inputs, data
}
