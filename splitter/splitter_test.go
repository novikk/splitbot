package splitter

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	s := Splitter{}
	s.AddPersonToGroup("Ivan de la Rubia")
	s.AddPersonToGroup("cristian")
	s.AddPersonToGroup("ricard")
	s.AddPersonToGroup("Jan Carbonell")

	s.RegisterExpense("ivan", "Cristian Jara", "test", 100)
	s.RegisterExpense("Ricard Tapias", "Cristian Jara", "test", 100)
	s.RegisterExpense("Ricard Tapias", "Jan", "test", 33)

	fmt.Println(s.GetBalanceLeaderboard())

	fmt.Println(s.RemoveDebt("cristian"))
	fmt.Println(s.GetBalanceLeaderboard())
}
