package webhooks

import (
	"fmt"
	"net/http"
)

func HandleSettleDebt(w http.ResponseWriter, r *http.Request) {
	payments := split.RemoveDebt(lastSpeaker)
	msg := fmt.Sprintf("Okay! Here's the list of payments you must perform\n")

	for _, p := range payments {
		msg += fmt.Sprintf("* %d to %s\n", p.Quantity, p.To.Name+" "+p.To.LastName)
	}

	w.Write([]byte(`{"text":"` + msg + `"}`))
}
