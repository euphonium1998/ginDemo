package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main()  {
	var rectangle Rectangle
	rand.Seed(time.Now().Unix())
	rectangle.X = rand.Intn(10)
	rectangle.Y = rand.Intn(10)
	fmt.Println(rectangle.X, "\t", rectangle.Y)
	body, err := json.Marshal(rectangle)

	resp, err := http.Post("http://localhost:8080/test", "application/json", bytes.NewBuffer(body))
	//request, err := http.NewRequest("POST", "localhost:8080/test", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println("ans below: ")
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(buf))
}

type Rectangle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Square int `json:"square"`
}