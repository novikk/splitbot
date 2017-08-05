package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle webhook")

	var body map[string]interface{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)

	vars := body["variablesMap"].(map[string]interface{})
	howMuch := vars["how_much"].(map[string]interface{})["value"].(string)
	who := vars["who"].(map[string]interface{})["value"].(string)

	fmt.Println("HOW MUCH ---->", howMuch)
	fmt.Println("WHO ---->", who)

	msg := fmt.Sprintf("Perfect! I have registered that %s owes %s euros to %s :)", who, howMuch, lastSpeaker)

	w.Write([]byte(`{"text":"` + msg + `"}`))
	//fmt.Println(string(b))
}
