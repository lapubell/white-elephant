package server

import "sort"

func (s *server) availableGiftList() []string {
	output := []string{}

	// build list of all presents brought
	for _, p := range s.clients {
		gift := p.Gift
		giftIsChosen := false

		for _, p := range s.clients {
			if p.ChosenGift == gift {
				giftIsChosen = true
				break
			}
		}

		if !giftIsChosen {
			output = append(output, p.Gift)
		}
	}

	sort.Strings(output)
	return output
}
