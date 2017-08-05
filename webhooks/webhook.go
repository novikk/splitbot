package webhooks

import "github.com/novikk/splitbot/splitter"

var lastSpeaker string
var split splitter.Splitter

func init() {
	split = splitter.Splitter{}
}

func SetLastSpeaker(speaker string) {
	lastSpeaker = speaker
}
