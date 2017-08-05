package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HandleExpense(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)

	vars := body["variablesMap"].(map[string]interface{})
	howMuch := vars["how_much"].(map[string]interface{})["value"].(string)
	who := vars["who"].(map[string]interface{})["value"].(string)

	fmt.Println("HOW MUCH ---->", howMuch)
	fmt.Println("WHO ---->", who)

	msg := fmt.Sprintf("Perfect! I have registered that %s owes %s€ to %s :)", who, howMuch, lastSpeaker)

	hmint, err := strconv.Atoi(howMuch)
	if err != nil {
		fmt.Println("err converting to int", err)
	}

	split.RegisterExpense(lastSpeaker, who, "", hmint)
	w.Write([]byte(`{"text":"` + msg + `"}`))
}
