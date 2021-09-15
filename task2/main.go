package main

import "fmt"

type Person struct {
	Name string
	Surname string
	Address string
	*Employee
}

type Employee struct {
	SalaryPerHour int
	WorkingHours int
	Position string
	Phone string
}

type  Persons []Person

func main(){
	personsArray := Persons{
		Person{Name: "Pavel", Surname: "Kirillov", Address: "Minsk", Employee: &Employee{SalaryPerHour: 3,
				WorkingHours:  8,
				Position:      "Developer",
				Phone:         "+375291111111",
			},
		},
		Person{Name: "Artem", Surname: "Korenko", Address: "Kobrin", Employee: &Employee{SalaryPerHour: 4,
			WorkingHours:  6,
			Position:      "Front-end developer",
			Phone:         "+375292222222",
		},
		},
		Person{Name: "Gleb", Surname: "Andreev", Address: "Bobruisk", Employee: &Employee{SalaryPerHour: 5,
			WorkingHours:  8,
			Position:      "Janitor",
			Phone:         "+375293333333",
		},
		},
		Person{Name: "Kirill", Surname: "Kirichyuk", Address: "Vitebsk", Employee: &Employee{SalaryPerHour: 5,
			WorkingHours:  6,
			Position:      "Developer",
			Phone:         "+375294444444",
		},
		},
	}

	for index, value := range personsArray{
		fmt.Printf("%d | Name: %s, Surname: %s, Adress: %s, Salary per hour: %d, Working hours:" +
			" %d, Position: %s Phone: %s\n", index+1, value.Name, value.Surname, value.Address, value.SalaryPerHour,
			value.WorkingHours, value.Position, value.Phone)
	}
}
