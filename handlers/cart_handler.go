package handlers

import (
	"encoding/json"
	"example.com/shopping/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type cartHandler struct {
	cartService services.CartService
}

type CartHandler interface {
	AddToCart(w http.ResponseWriter, r *http.Request)
}

func (cH *cartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var itemIds []int
	userIdStr := r.Header.Get("X-User-Id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, err.Error()+"\ninvalid user id", http.StatusBadRequest)
	}
	err = json.NewDecoder(r.Body).Decode(&itemIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unAddedItems, err := cH.cartService.AddToCart(userId, itemIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(unAddedItems) > 0 {
		w.WriteHeader(http.StatusPartialContent)
		response := strings.Trim(strings.Join(strings.Split(fmt.Sprint(unAddedItems), " "), ","), "[]")
		_, err = w.Write([]byte(response))
		if err != nil {
			fmt.Print("could not write the response")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
