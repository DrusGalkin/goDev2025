package main

type Shop struct {
	Name        string
	Tovars      []Tovar
	MapCategory map[category][]Tovar
}

type BonusManager struct {
	Shop Shop
	Bonus Bonus
}

//func (bm *BonusManager) isBonus() int{
//
//}
//
//func (s *Shop) AddCategoryToBonus(c category) {
//	var category, trueIs = s.MapCategory[c]
//	if trueIs {
//		for _, el := range category {
//			el.Bonus = &BonusManager{Bonus: }.isBonus
//		}
//	}
//}
