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

type Persons []Person

type PrintPerson interface {
	SalaryCalculation() int
	PrintInfo() string
}

type PrintPersons interface {
	GetAllPersons() []Person
}

func (p* Persons)GetAllPersons() []Person {
	return *p
}

func PrintAllPersons(p PrintPersons){
	for index, value := range p.GetAllPersons(){
		fmt.Printf("%d | Name: %s, Surname: %s, Adress: %s, Salary per hour: %d, Working hours:" +
			" %d, Position: %sPhone: %s\n", index+1, value.Name, value.Surname, value.Address, value.SalaryPerHour,
			value.WorkingHours, value.Position, value.Phone)
	}
}

func (p *Person)PrintInfo() string{
	return fmt.Sprintf("Information about the person\nName: %s\nSurname: %s\nAdress: %s\n" +
		"Salary per hour: %d\nWorking hours: %d\nPosition: %s\nPhone: %s",
		p.Name, p.Surname, p.Address, p.SalaryPerHour, p.WorkingHours, p.Position, p.Phone)
}


func (p *Person)SalaryCalculation() int{
	return p.SalaryPerHour*p.WorkingHours*20
}


func GetSalaryCalculation(p PrintPerson){
	fmt.Printf("Monthly salary is %d$\n\n", p.SalaryCalculation())
}

func GetInfo(p PrintPerson){
	fmt.Println(p.PrintInfo())
}

func main(){

	p1 := &Person{Name: "Pavel", Surname: "Kirillov", Address: "Minsk", Employee: &Employee{SalaryPerHour: 3,
		WorkingHours:  8,
		Position:      "Developer",
		Phone:         "+375291111111",
	},
	}
	p2 := &Person{Name: "Artem", Surname: "Korenko", Address: "Kobrin", Employee: &Employee{SalaryPerHour: 4,
		WorkingHours:  6,
		Position:      "Front-end developer",
		Phone:         "+375292222222",
	},
	}
	p3 := &Person{Name: "Kirill", Surname: "Kirichyuk", Address: "Vitebsk", Employee: &Employee{SalaryPerHour: 5,
		WorkingHours:  6,
		Position:      "Developer",
		Phone:         "+375294444444",
	},
	}
	p4 := &Person{Name: "Gleb", Surname: "Andreev", Address: "Bobruisk", Employee: &Employee{SalaryPerHour: 5,
		WorkingHours:  8,
		Position:      "Janitor",
		Phone:         "+375293333333",
	},
	}
	GetInfo(p1)
	GetSalaryCalculation(p1)
	GetInfo(p2)
	GetSalaryCalculation(p2)
	GetInfo(p3)
	GetSalaryCalculation(p3)
	GetInfo(p4)
	GetSalaryCalculation(p4)
	PrintAllPersons(&Persons{*p1,*p2,*p3,*p4})
}
