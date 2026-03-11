package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const totalPoints = 100
const pointsPerQuestion = 20

func main() {
	fmt.Println("Вітаю у грі MathCore \\m/")

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Хай щастить!")
	myPoints := 0
	start := time.Now()

	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Не правильно!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Printf("Правильно! Кількість балів: %v\n", myPoints)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}

	end := time.Now()
	timeSpent := end.Sub(start)
	fmt.Printf("Вітаю! Ти впорався(лась) за %v!", timeSpent)
}
