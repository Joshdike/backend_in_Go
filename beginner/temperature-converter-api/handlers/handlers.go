package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Joshdike/backend_in_Go/beginner/temperature-converter-api/models"
)

func Convert(w http.ResponseWriter, r *http.Request){
	var req models.ConversionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	var celsius float64
	if strings.HasPrefix(req.From, "C"){
		celsius = req.Val
	}else if strings.HasPrefix(req.From, "K"){
		celsius = req.Val - 273.15
	}else if strings.HasPrefix(req.From, "F"){
		celsius = (req.Val - 32) * 5/9
	}else{
		fmt.Fprint(w, "Invalid request")
	}

	var result float64
	if strings.HasPrefix(req.To, "C"){
		result = celsius
	}else if strings.HasPrefix(req.To, "K"){
		result = celsius + 273.15
	}else if strings.HasPrefix(req.To, "F"){
		result = (celsius * 9/5) + 32
	}else{
		fmt.Fprint(w, "Invalid request")
	}

	fmt.Fprintf(w, "converting %v from %v to %v is equal to %v", req.Val, req.From, req.To, result)
}