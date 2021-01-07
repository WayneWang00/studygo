package main

import (
	"fmt"
)

func main() {
	//pointList := make([][]int, 9)
	//points := []int{1, 301, 601, 901, 1201, 1501, 2101, 3001, 4001}
	//for i := 0; i < len(points); i++ {
	//	switch {
	//	case points[i] >= 0 && points[i] <= 300:
	//		pointList[0] = append(pointList[0], points[i])
	//	case points[i] > 300 && points[i] <= 600:
	//		pointList[1] = append(pointList[1], points[i])
	//	case points[i] > 600 && points[i] <= 900:
	//		pointList[2] = append(pointList[2], points[i])
	//	case points[i] > 900 && points[i] <= 1200:
	//		pointList[3] = append(pointList[3], points[i])
	//	case points[i] > 1200 && points[i] <= 1500:
	//		pointList[4] = append(pointList[4], points[i])
	//	case points[i] > 1500 && points[i] <= 2100:
	//		pointList[5] = append(pointList[5], points[i])
	//	case points[i] > 2100 && points[i] <= 3000:
	//		pointList[6] = append(pointList[6], points[i])
	//	case points[i] > 3000 && points[i] <= 4000:
	//		pointList[7] = append(pointList[7], points[i])
	//	default:
	//		pointList[8] = append(pointList[8], points[i])
	//	}
	//}
	//fmt.Println("pointList Values: ", pointList)

	n := 201
	teamList := make([]int, n)
	for i := 0; i < n; i++ {
		teamList[i] = i
	}
	newGrouping(teamList)
}

func newGrouping(team []int) {
	teamList := make([][]int, 0)
	number := len(team) / 50
	remainder := len(team) % 50
	switch {
	case number == 0:
		//teamList[1] = team
		teamList = append(teamList, team)
	case remainder > 30 && remainder < 50:
		number = number + 1
		for i := 0; i < number; i++ {
			if 50*(i+1) >= len(team) {
				teamList = append(teamList, team[50*i:])
				//teamList[i+1] = team[50*i:]
				//TODO: 添加机器人使每个战区达到100人
				break
			}
			teamList = append(teamList, team[50*i:50*(i+1)])
			//teamList[i+1] = team[50*i : 50*(i+1)]
			//TODO: 添加机器人使每个战区达到100人
		}
	case remainder == 0:
		for i := 0; i < number; i++ {
			teamList = append(teamList, team[50*i:50*(i+1)])
			//teamList[i+1] = team[50*i : 50*(i+1)]
			//TODO: 添加机器人使每个战区达到100人
		}
	case remainder > 0 && remainder <= 30:
		fmt.Println("test start")
		number1 := remainder / number
		remainder1 := remainder % number
		fmt.Println("number1:", number1, " remainder1:", remainder1)
		if number1 > 5 {
			fmt.Println("if")
			number = number + 1
			for i := 0; i < number; i++ {
				if 50*(i+1) >= len(team) {
					teamList = append(teamList, team[50*i:])
					//teamList[i+1] = team[50*i:]
					//TODO: 添加机器人是每个战区达到100人
					break
				}
				teamList = append(teamList, team[50*i:50*(i+1)])
				//teamList[i+1] = team[50*i : 50*(i+1)]
				//TODO: 添加机器人使每个战区达到100人
			}
		} else {
			fmt.Println("else")
			for i := 0; i < number; i++ {
				//list := team[(50+number1)*i : (50+number1)*(i+1) : 2*(50+number1)]
				//fmt.Println("number")
				//fmt.Println(i, ":", list)
				//if i < remainder1 {
				//	list = append(list, team[(50+number1)*number+i])
				//	fmt.Println("remainder1")
				//	fmt.Println(i, ":", list)
				//}
				//teamList = append(teamList, list)
				//teamList = append(teamList, team[(50+number1)*i:(50+number1)*(i+1)])
				newList := make([]int, 50+number1)
				for z := i * 50; z < (i+1)*50; z++ {
					if z < len(team) {
						newList[z-i*50] = team[z]
					}
				}
				//newList := team[(50+number1)*i : (50+number1)*(i+1)]
				fmt.Println("number")
				fmt.Println(i, ":", newList)
				if i < remainder1 {
					//teamList[i] = append(teamList[i], team[(50+number1)*number+i])
					newList = append(newList, team[(50+number1)*number+i])
					//teamList[i+1] = append(teamList[i+1], team[(50+number1)*number+i])
					fmt.Println("remainder1")
					fmt.Println(i, ":", newList)
					//}
				}
				//teamList[i+1] = newList
				teamList = append(teamList, newList)
				//for i := 0; i < remainder1; i++ {
				//	teamList[i] = append(teamList[i], team[(50+number1)*number+i])
				//	fmt.Println("remainder1")
				//	fmt.Println(i, ":", teamList[i])
				//}
				//TODO: 添加机器人使每个战区达到100人
			}
		}
		fmt.Println("end")
		for i := 0; i < len(teamList); i++ {
			fmt.Println(i, ":", teamList[i])
		}
		//for k, v := range teamList {
		//	fmt.Println(k, ":", v)
		//}
	}
}
