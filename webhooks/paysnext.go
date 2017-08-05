package webhooks

import (
	"fmt"
	"net/http"
)

func HandlePaysNext(w http.ResponseWriter, r *http.Request) {
	whoP := split.GetBalanceLeaderboard()[0]
	go sc.ShowContacts()
	msg := fmt.Sprintf("The next one to pay is... %s!", whoP.Person)

	w.Write([]byte(`{"text":"` + msg + `"}`))
}
