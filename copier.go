package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type Use struct {
	Name        string
	Role        string
	Age         int32
	Salary   int
}

type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	Name      string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	Age       int32  `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	Salary    int   `copier:"must"`
	Salary12  int  `copier:"-"`

}


func main() {
	var (
		user      = Use{Name: "Jinzhu", Age: 18, Salary: 200000}
		users     = []Use{{Name: "Jinzhu", Age: 18, Salary: 100000}, {Name: "jinzhu 2", Age: 30, Salary: 60000}}
		employee  = Employee{Salary: 1111111111111, Salary12: 1}
		employees = []Employee{}
	)

	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    Salary:150000,            // Copying explicitly ignored
	// }

	// Copy struct to slice
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeeId: 0, SuperRole: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	employees = append(employees, Employee{Name:"Jinzhu", Age:18, Salary:1111111111111, Salary12:12345})
	employees = append(employees, Employee{Name:"JJJ", Age:118, Salary:100, Salary12:1})
	fmt.Println(employees)
	copier.Copy(&employees, &users)

	fmt.Printf("%#v \n", employees)
	// fmt.Printf("%#v \n", users)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeeId: 0, SuperRole: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, Salary:0, DoubleAge: 60, EmployeeId: 0, SuperRole: "Super Dev"},
	// }

	// copier.CopyWithOption(&employee, &user, copier.Option{IgnoreEmpty: true, DeepCopy: false})
	// fmt.Printf("%#v \n", employee)

}