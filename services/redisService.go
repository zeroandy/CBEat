package services

import (
	"OrderApi/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {

}

func SaveItem(item interface{}) {
	fmt.Println()
	fmt.Printf("%+v", item)
	if c, ok := item.(model.Item); ok {
		fmt.Println("Save To Redis")
		jsondata, _ := json.Marshal(c)
		fmt.Println(string(jsondata))
		err := rdb.HSet(ctx, "Items", c.Name, string(jsondata)).Err()
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(ok)
	}
}
