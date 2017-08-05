package webhooks

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intent := vars["intent"]
	fmt.Fprintf(w, "Intent: "+intent)
}
