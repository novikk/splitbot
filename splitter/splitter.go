package splitter

import (
	"sort"
	"strings"
)

type Payment struct {
	From     Person
	To       Person
	Quantity int
	Concept  string
}

type Person struct {
	Name, LastName string
	Balance        int
	Link           string
}

type Splitter struct {
	people   []Person
	payments []Payment
}

type ByBalance []Person

func (s ByBalance) Len() int {
	return len(s)
}
func (s ByBalance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByBalance) Less(i, j int) bool {
	return s[i].Balance > s[j].Balance
}

type ByBalanceNeg []Person

func (s ByBalanceNeg) Len() int {
	return len(s)
}
func (s ByBalanceNeg) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByBalanceNeg) Less(i, j int) bool {
	return s[i].Balance < s[j].Balance
}

func (s *Splitter) AddPersonToGroup(nameAll string) {
	found := false
	info := strings.Split(nameAll, " ")
	name := info[0]

	for i, p := range s.people {
		if strings.ToLower(p.Name) == strings.ToLower(name) {
			if p.LastName == "" {
				// person found, update!
				s.people[i].Name = info[0]
				s.people[i].LastName = strings.Join(info[1:], " ")

				found = true
			} else {
				return
			}
		}
	}

	if !found {
		info := strings.Split(nameAll, " ")
		np := Person{info[0], strings.Join(info[1:], " "), 0, ""}
		s.people = append(s.people, np)
	}
}

func (s *Splitter) RegisterExpense(from, to, concept string, quantity int) {
	s.AddPersonToGroup(from)
	s.AddPersonToGroup(to)

	fi, ti := 0, 0
	infoFrom, infoTo := strings.Split(from, " "), strings.Split(to, " ")
	nameFrom, nameTo := infoFrom[0], infoTo[0]

	for i, p := range s.people {
		if strings.ToLower(p.Name) == strings.ToLower(nameFrom) {
			fi = i
		}

		if strings.ToLower(p.Name) == strings.ToLower(nameTo) {
			ti = i
		}
	}

	s.payments = append(s.payments, Payment{s.people[fi], s.people[ti], quantity, concept})
	s.people[fi].Balance -= quantity
	s.people[ti].Balance += quantity
}

func (s *Splitter) RemoveDebt(who string) []Payment {
	s.AddPersonToGroup(who)

	var res []Payment

	wi := 0
	info := strings.Split(who, " ")
	name := info[0]

	sort.Sort(ByBalanceNeg(s.people))

	for i, p := range s.people {
		if strings.ToLower(p.Name) == strings.ToLower(name) {
			wi = i
		}
	}

	if s.people[wi].Balance < 0 {
		return res
	}

	for i := range s.people {
		if s.people[i].Balance < -s.people[wi].Balance {
			res = append(res, Payment{
				s.people[wi],
				s.people[i],
				s.people[wi].Balance,
				"",
			})
			s.people[wi].Balance, s.people[i].Balance = 0, s.people[i].Balance+s.people[wi].Balance
			break
		} else {
			res = append(res, Payment{
				s.people[wi],
				s.people[i],
				-s.people[i].Balance,
				"",
			})
			s.people[wi].Balance, s.people[i].Balance = s.people[wi].Balance+s.people[i].Balance, 0
		}
	}

	return res
}

type Leaderboard struct {
	Person   string
	Quantity int
}

func (s *Splitter) GetBalanceLeaderboard() []Leaderboard {
	sort.Sort(ByBalance(s.people))
	var res []Leaderboard
	for i := range s.people {
		if s.people[i].Balance > 0 {
			res = append(res, Leaderboard{s.people[i].Name + " " + s.people[i].LastName, s.people[i].Balance})
		}
	}

	return res
}
