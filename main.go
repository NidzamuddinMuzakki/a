package main

import (
	"fmt"
	"math"
)

func narcissistic(bilangan int) bool {
	akhir := bilangan
	hasil := 0
	simpanJumalh := bilangan
	sisa := 0
	jumlah := 0
	for simpanJumalh > 0 {
		simpanJumalh = simpanJumalh / 10
		jumlah++
	}
	for {
		sisa = akhir % 10
		akhir = akhir / 10
		hasil += int(math.Pow(float64(sisa), float64(jumlah)))
		if sisa == 0 {
			break
		}
	}
	if hasil == bilangan {
		return true
	} else {
		return false
	}

}

func outlier(number ...int) bool {
	odd := 0
	even := 0
	for _, ss := range number {
		if ss%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if odd == 0 || even == 0 {
		return false
	} else {
		return true
	}
}

func findNeedle(aray []string, cari string) int {
	s := -1
	for in, ss := range aray {
		if ss == cari {
			s = in
		}
	}
	return s
}
func blueOcean(array []int, array2 []int) []int {
	tempArray := []int{}
	for _, s1 := range array {
		sama := 0
		for _, s2 := range array2 {
			if s1 == s2 {
				sama = 1
				break
			}
		}
		if sama != 1 {
			tempArray = append(tempArray, s1)
		}
	}
	return tempArray
}
func main() {
	fmt.Println("narcissistic", narcissistic(153))
	fmt.Println("narcissistic", narcissistic(111))

	fmt.Println("outlier", outlier(2, 4, 0, 100, 4, 11, 2602, 36))
	fmt.Println("outlier", outlier(160, 3, 1719, 19, 11, 13, -21))
	fmt.Println("outlier", outlier(11, 13, 15, 19, 9, 13, -21))

	fmt.Println("findNeedle", findNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	fmt.Println("blueOcean", blueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	fmt.Println("blueOcean", blueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))

}
