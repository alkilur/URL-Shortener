package handlers

import "net/http"

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"`
}

func Ok(alias string) Response {
	return Response{
		Status: http.StatusOK,
		Alias:  alias,
	}
}

func Error(msg string) Response {
	return Response{
		Status: http.StatusBadRequest,
		Error:  msg,
	}
}
