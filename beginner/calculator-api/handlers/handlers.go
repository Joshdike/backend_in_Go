package handlers

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
)

type Request struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func parseRequest(r *http.Request) (Request, error) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return Request{}, errors.New("invalid request body")
	}
	return req, nil
}

func writeResponse(w http.ResponseWriter, result float64, err error) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Result: result}
	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(resp)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, req.X+req.Y, nil)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, req.X-req.Y, nil)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, req.X*req.Y, nil)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	if req.Y == 0 {
		err = errors.New("division by zero error")
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, req.X/req.Y, nil)
}

func NthRootHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, math.Pow(req.X, 1/req.Y), nil)
}
func ExponentHandler(w http.ResponseWriter, r *http.Request) {
	req, err := parseRequest(r)
	if err != nil {
		writeResponse(w, 0, err)
		return
	}
	writeResponse(w, math.Pow(req.X, req.Y), nil)
}
