package main

import (
	"fmt"

	"github.com/rlyders/go-receivers/src/model"
)

func main() {
	cat := model.CreateCat(1)
	fmt.Printf("cat=%+v", cat)
}
