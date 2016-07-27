package logic

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Responce struct {
	State string
}


func GetDataJSON(url string, response *Responce) error{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response.State = string(body)
	return nil
}
