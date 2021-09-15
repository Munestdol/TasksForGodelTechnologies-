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

func (p *Person)PrintInfo(){
	fmt.Printf("Information about the person\nName: %s\nSurname: %s\nAdress: %s\n" +
		"Salary per hour: %d\nWorking hours: %d\nPosition: %s\nPhone: %s\n",
		p.Name, p.Surname, p.Address, p.SalaryPerHour, p.WorkingHours, p.Position, p.Phone)
}

func SalaryCalculation(p* Person) int {
	return p.SalaryPerHour*p.WorkingHours*20
}

func main(){
	p := &Person{Name: "Pavel",
		Surname: "Kirillov",
		Address: "Minsk",
		Employee: &Employee{
			SalaryPerHour: 3,
			WorkingHours: 8,
			Position: "Developer",
			Phone: "+375291111111",
		},
	}
	p.PrintInfo();
	sc := SalaryCalculation(p)
	if sc == 0 {
		fmt.Println("Error: SalaryPerHour or WorkingHours field are empty")
	} else {
		fmt.Printf("%s %s get's monthly salary is %d$", p.Name, p.Surname, sc)
	}
}