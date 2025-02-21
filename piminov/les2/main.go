package main

import "fmt"

func chet(a []int) []int {
	var temp = make([]int, 0)
	for i, el := range a {
		if el%2 != 0 && i >= 3 {
			temp = append(temp, el)
		}
	}
	return temp
}

func superZ() func() int {
	var f int
	return func() int {
		f++
		return f * 2
	}
}

// уникальные символы
func uniqueSymbols(str string) string {
	var temp = make(map[rune]string)
	var res string

	for _, el := range str {
		temp[el] = string(el)
	}

	for _, el := range temp {
		res += el
	}

	return res
}

func bubbleSort(sl ...int) []int {
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl)-1-i; j++ {
			if sl[j] > sl[j+1] {
				temp := sl[j]
				sl[j] = sl[j+1]
				sl[j+1] = temp
			}
		}
	}
	return sl
}

func twoInOne(sl1, sl2 []int) []int {
	var temp = make([]int, len(sl1)+len(sl2))
	temp = append(sl1, sl2...)
	return bubbleSort(temp...)
}

func main() {
	//уникальные
	fmt.Println(uniqueSymbols("ййййццццц"))
	//дерево
	tree := &Tree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(14)
	tree.Insert(1)
	tree.Insert(654)
	tree.Insert(1000)

	tree.Print()

	//2 слайза в 1
	var s1 = []int{3, 5, 2, 5, 5, 5}
	var s2 = []int{1, 3, 5, 4}

	res := twoInOne(s1, s2)
	fmt.Println(res)

}
