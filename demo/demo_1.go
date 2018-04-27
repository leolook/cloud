package main

import (
	"cloud/common/db"
	"time"
)

func main() {
	client := db.GetClient()
	client.HSet("demo_test", "demo", time.Now().Unix())
}
