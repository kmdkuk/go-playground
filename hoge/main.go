package main

import "fmt"

func main() {
	nonss := []string{
		"nonss0",
		"nonss1",
		"nonss2",
		"nonss3",
		"nonss4",
		"nonss5",
		"nonss6",
		"nonss7",
		"nonss8",
		"nonss9",
		"nonss10",
		"nonss11",
		"nonss12",
		"nonss13",
		"nonss14",
		"nonss15",
		"nonss16",
		"nonss17",
		"nonss18",
	}
	ss := []string{
		"ss0",
		"ss1",
		"ss2",
		"ss3",
		"ss4",
		"ss5",
		"ss6",
		"ss7",
		"ss8",
		"ss9",
		"ss10",
	}
	nonssIndex := 0
	ssIndex := 0
	for i := 0; i < len(ss)+len(nonss); i++ {
		var address string
		if nonssIndex <= ssIndex {
			address = nonss[nonssIndex]
			nonssIndex++
			fmt.Println("pick nonss")
		} else {
			address = ss[ssIndex]
			ssIndex++
			fmt.Println("pick ss")
		}
		fmt.Println("i:", i)
		fmt.Println("nonssIndex:", nonssIndex)
		fmt.Println("ssIndex:", ssIndex)
		fmt.Println("address:", address)
		fmt.Println()
	}
}
