package webhooks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HandleExpense(w http.ResponseWriter, r *http.Request) {
	var body interface{}
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &body)

	fmt.Println(body)
}
