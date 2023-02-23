package main

import (
	"fmt"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := make(map[string]Man)

	people["Джонни"] = Man{
		Name:     "Джони",
		LastName: "Крапленный",
		Age:      33,
		Gender:   "м",
		Crimes:   10,
	}

	people["Аль"] = Man{
		Name:     "Аль",
		LastName: "Капоне",
		Age:      44,
		Gender:   "м",
		Crimes:   1,
	}

	people["Усама"] = Man{
		Name:     "Усама",
		LastName: "бен Ладен",
		Age:      65,
		Gender:   "м",
		Crimes:   3,
	}

	people["Бонни"] = Man{
		Name:     "Бонни",
		LastName: "Паркер",
		Age:      24,
		Gender:   "ж",
		Crimes:   8,
	}

	people["Клайд"] = Man{
		Name:     "Клайд",
		LastName: "Бэрроу",
		Age:      25,
		Gender:   "м",
		Crimes:   9,
	}

	people["Зинаида"] = Man{
		Name:     "Зинаида",
		LastName: "Тютюкина",
		Age:      30,
		Gender:   "ж",
		Crimes:   34,
	}

	var suspects []string

	suspects = append(suspects, "Зинаида", "Клайд", "Бонни", "Усама", "Аль", "Джонни", "Борис Бритва", "Иван")

	var theCriminal Man

	for i, Name := range suspects {
		criminal, ok := people[Name]
		if ok == false {
			fmt.Printf("Преступник %s отсутсвует в базе данных.\n", suspects[i])
			recover()
		}
		if criminal.Crimes > theCriminal.Crimes {
			theCriminal = criminal
		}

	}

	fmt.Printf("Самый опасный преступник: %s %s.\nУ него(нее) %d преступлений.\n", theCriminal.Name, theCriminal.LastName, theCriminal.Crimes)

}
