package main

import "fmt"

func main() {
	b := make([][]int, 2)
	a := make([]int, 70)
	for i := 0; i < 70; i++ {
		a[i] = i
	}
	if len(a) <= 55 {
		b[0] = a
	} else if len(a) > 55 {
		for i := 0; i < len(a); i++ {
			switch {
			case a[i] >= 0 && a[i] <= 300:
				b[0] = append(b[0], a[i])
			default:
				b[1] = append(b[1], a[i])
			}
		}
	}
	for i := 0; i < len(b); i++ {
		group(b[i], i+1)
	}
}

func group(team []int, num int) {
	teamList := make([][]int, 0)
	number := len(team) / 50
	remainder := len(team) % 50
	if remainder > 30 && remainder < 50 {
		number = number + 1
		for i := 0; i < number; i++ {
			if 50*(i+1) >= len(team) {
				teamList = append(teamList, team[50*i:])
				break
			}
			teamList = append(teamList, team[50*i:50*(i+1)])
		}
	} else if remainder == 0 {
		for i := 0; i < number; i++ {
			teamList = append(teamList, team[50*i:50*(i+1)])
		}
	} else if remainder > 0 {
		fmt.Println("number: ", number, " remainder: ", remainder)
		number1 := remainder / number
		remainder1 := remainder % number
		fmt.Println("number1: ", number1, " remiander1: ", remainder1)
		if number1 > 5 {
			number = number + 1
			for i := 0; i < number; i++ {
				if 50*(i+1) >= len(team) {
					teamList = append(teamList, team[50*i:])
					break
				}
				teamList = append(teamList, team[50*i:50*(i+1)])
			}
		} else {
			for i := 0; i < number; i++ {
				teamList = append(teamList, team[(50+number1)*i:(50+number1)*(i+1)])
			}
			for i := 0; i < remainder1; i++ {
				teamList[i] = append(teamList[i], team[(50+number1)*number+i])
			}
		}
	}
	for i := 0; i < len(teamList); i++ {
		fmt.Printf("team=%d, list=%d, value: %+v\n", num, i+1, teamList[i])
	}
	fmt.Printf("team=%d分组完成\n", num)
}
