package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mathcore-go/domain"
	"os"
	"sort"
	"strconv"
	"time"
)

const totalPoints = 100
const pointsPerQuestion = 100

var id uint64 = 1

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func main() {
	fmt.Println("Вітаю у грі MathCore \\m/")

	var users []domain.User

	for {
		menu()

		choice := ""
		fmt.Scan(&choice)

		switch choice {
		case "1":
			u := play()
			users = append(users, u)
			sortAndSave(users)
		case "2":
			for _, u := range users {
				fmt.Printf("Id: %v Name: %s Time: %v\n",
					u.Id, u.Name, u.Time)
			}
		case "3":
			return
		default:
		}
	}
}

func play() domain.User {
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

	fmt.Println("Введіть ім'я: ")
	name := ""

	fmt.Scan(&name)

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Time < users[j].Time
	})

	file, err := os.OpenFile(
		"users.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755,
	)
	if err != nil {
		log.Printf("sortAndSave(os.OpenFile): %s", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("sortAndSave(file.Close): %s", err)
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		log.Printf("sortAndSave(encoder.Encode): %s", err)
		return
	}
}

func getUsers() []domain.User {

}
