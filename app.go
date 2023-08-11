package main

import (
	"context"
	// "encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Body struct {
	body [][]string
}

type Res struct {
	body   string
	status int
}

func (a *App) Req(path, method string, bodyIn Body) [2]string {
	// Create a new HTTP client
	client := &http.Client{}

	form := url.Values{}
	for _, field := range bodyIn.body {
		form.Add(field[0], field[1])
	}
	// Create a new request object
	req, err := http.NewRequest(method, path, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	// Add any headers that the API requires
	req.Header.Add("Authorization", "Bearer <your_access_token>")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Make sure to close the response body when we're done with it
	defer resp.Body.Close()

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)

	println(string(body), resp.Status)
	if err != nil {
		panic(err)
	}

	returned := [2]string{string(body), resp.Status}
	return returned
}
