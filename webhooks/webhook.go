package webhooks

import (
	"github.com/novikk/splitbot/sage"
	"github.com/novikk/splitbot/splitter"
)

var lastSpeaker string
var split splitter.Splitter
var sc sage.SageClient

func init() {
	split = splitter.Splitter{}
	sc = sage.SageClient{}

	sc.RefreshToken = "4109e944807d9f7cda0c345fed136564a4a26501"
	sc.AccessToken = "9b872d0717eff2e296557dd09f0db4d2076a369a"
	sc.ResourceOwnerID = "cd955c24cdd52c60fa835a1ff54ffb4d"
	sc.ExpirationDate = 1501951648
}

func SetLastSpeaker(speaker string) {
	lastSpeaker = speaker
}
