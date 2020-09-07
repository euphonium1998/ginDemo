package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rectangle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Square int `json:"square"`
}

func main()  {
	http.HandleFunc("/test", myFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("begin http server error[%v]\n", err)
	}
}

func myFunc(w http.ResponseWriter, r *http.Request) {
	var rectangle Rectangle
	defer func() {
		wData, _ := json.Marshal(rectangle)
		w.Write(wData)
	}()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("err[%v]\n", err)
		return
	}
	err = json.Unmarshal(buf, &rectangle)
	if err != nil {
		fmt.Printf("json unmarshal err[%v]\n", err)
	}
	rectangle.Square = rectangle.X * rectangle.Y
	fmt.Println(rectangle)
}
