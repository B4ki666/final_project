package handler

import (
	"encoding/json"
	"net/http"
)

type SignInRequest struct {
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var request SignInRequest
	var response SignInResponse

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.HandleError(err, w)
		return
	}

	token, err := h.Service.SignIn(request.Password)
	if err != nil {
		h.HandleError(err, w)
		return
	}

	response.Token = token

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Printf("Error encoding: %v", err)
	}
}
