package main
    
import (
	"github.com/hunjixin/automapper"
	"reflect"
	"fmt"
	"time"
)

type UserDto struct {
	Name string
	Addr string
	Age  int
}

type User1 struct {
	Name  string
	Nick  string
	Addr  string
	Birth time.Time
}

func main(){
	users := [2]*User1{}
	u := UserDto{}
	users[0] = &User1{"Hellen", "NICK", "B·J", time.Date(1992, 10, 3, 1, 0, 0, 0, time.UTC)}
	users[1] = &User1{"Jack", "neo", "W·S", time.Date(1992, 10, 3, 1, 0, 0, 0, time.UTC)}
	// result2 := automapper.MustMapper(users[0], reflect.TypeOf([]*UserDto{}))

	for i := range users {
		result2 := automapper.MustMapper(users[i], reflect.TypeOf((*UserDto)(nil)))
		fmt.Println(result2)
		resultSlice := result2.(*UserDto)
		// fmt.Println(resultSlice)
		u=*resultSlice
	}
	fmt.Print(u)

	// resultSlice := result2.([]*UserDto)

	// for i := range resultSlice {
	// 	fmt.Println(resultSlice[i])
	// }
}