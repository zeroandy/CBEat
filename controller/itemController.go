package controller

import (
	"OrderApi/model"
	services "OrderApi/services"
	"encoding/json"
	"fmt"
	"net/http"
)

var Items []model.Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Items")
	Items = []model.Item{
		model.Item{Name: "金牌愛玉檸檬", Desc: "就愛玉檸檬 在那邊", Price: 500},
		model.Item{Name: "鳴光蜜金吉", Desc: "金吉檸檬?", Price: 500},
	}
	json.NewEncoder(w).Encode(Items)
}

func SaveItems(w http.ResponseWriter, r *http.Request) {
	services.SaveItem(model.Item{Name: "金牌愛玉檸檬", Desc: "就愛玉檸檬 在那邊", Price: 500})
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test")
	fmt.Fprintf(w, "ApiWroked")
}
