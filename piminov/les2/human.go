package main

import "fmt"

type Human struct {
	Name  string
	Money int
	Bonus int
}

func (s *Human) HumanInShop(shop Shop, tovars ...Tovar) {
	var temp = make([]Tovar, 0)
	var money = 0

	for _, el := range tovars {
		for _, j := range shop.Tovars {
			if el == j {
				temp = append(temp, j)
				money += j.Price
			} else {
				fmt.Println("Товара " + fmt.Sprint(el.Name) + " нет в магазине")
			}
		}
	}

	for _, el := range temp {
		if s.Money < money {
			fmt.Println("У вас не хватает денег, пополните баланс")
			return
		} else {
			s.Money -= el.Price
		}
	}
}
