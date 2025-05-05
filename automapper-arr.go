package main
    
import (
	"github.com/hunjixin/automapper"
	"reflect"
	"fmt"
	"time"
)

type UserDtoo1 struct {
	Name string
	Addr string
	Age  int
}

type Userr struct {
	Name  string `mapping:"Name"`
	Nick  string
	Addr  string `mapping:"Addr"`
	Birth time.Time
}

func main(){
	users := [2]*Userr{}
	users[0] = &Userr{"Hellen", "NICK", "B·J", time.Date(1992, 10, 3, 1, 0, 0, 0, time.UTC)}
	users[1] = &Userr{"Jack", "neo", "W·S", time.Date(1992, 10, 3, 1, 0, 0, 0, time.UTC)}
	result2 := automapper.MustMapper(users, reflect.TypeOf([]*UserDtoo1{}))
	fmt.Println(result2)
}