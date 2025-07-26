package consts

import (
	"encoding/json"
	// "fmt"
	"gfAdmin/internal/model"
	"io/ioutil"
	"os"

)

var Menus []model.Menus
var Menus_byte []byte

func init() {
	//read Routes from json file
	file, err := os.Open("routes.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	json.Unmarshal(content, &Menus)
	Menus_byte = content
	//fmt.Println("readfile:")
	//fmt.Println(Menus)
}
