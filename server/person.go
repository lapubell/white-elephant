package server

import "sort"

type Person struct {
	Name       string `json:"name"`
	Gift       string `json:"gift"`
	ChosenGift string `json:"chosenGift"`
}

func (s *server) peopleList() []Person {
	output := []Person{}
	for _, p := range s.clients {
		output = append(output, p)
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Name < output[j].Name
	})

	return output
}
