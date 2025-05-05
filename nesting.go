// package main

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/devfeel/mapper"
// )

// // type RefLevel3 struct {
// // 	What string `json:"finally"`
// // }
// // type Level2 struct {
// // 	*RefLevel3 `json:"ref_level3"`
// // }
// // type Level1 struct {
// // 	Level2 `json:"level2"`
// // }
// // type TopLayer struct {
// // 	Level1 `json:"level1"`
// // }
// // type MadNest struct {
// // 	TopLayer `json:"top"`
// // }

// // var madnestStruct MadNest = MadNest{
// // 	TopLayer: TopLayer{
// // 		Level1: Level1{
// // 			Level2: Level2{
// // 				RefLevel3: &RefLevel3{
// // 					What: "matryoska",
// // 				},
// // 			},
// // 		},
// // 	},
// // }

// // func main() {
// // 	// since we're targeting the same MadNest, both of functions will yield
// // 	// same result hence this unified example/test.
// // 	var madnestObj = MadNest{}
// // 	var err error

// // 		madnestMap := smapping.MapTags(madnestStruct, "json")
// // 		err = smapping.FillStructByTags(&madnestObj, madnestMap, "json")
// // 		fmt.Println(madnestObj)

// // 		madnestMap = smapping.MapFields(madnestStruct)
// // 		err = smapping.FillStruct(&madnestObj, madnestMap)
// // 		fmt.Println(madnestObj)

// // 	if err != nil {
// // 		fmt.Printf("%s", err.Error())
// // 		return
// // 	}
// // 	// the result should yield as intented value.
// // 	if madnestObj.TopLayer.Level1.Level2.RefLevel3.What != "matryoska" {
// // 		fmt.Printf("Error: expected \"matroska\" got \"%s\"", madnestObj.Level1.Level2.RefLevel3.What)
// // 	}
// // }



  
//   type (
// 	codeco struct {
// 		Name1  string `json:"name1" mapper:"mm"`
// 		Name12  string `json:"name12"`
// 		Mine mine `json:"mine" mapper:"mmm"`
// 	}

// 	mine struct {
// 		Base string `json:"base"`
// 	}

// 	codecoS struct {
// 		Name134  string `json:"name14" mapper:"mm"`
// 	}

// 	User struct {
// 	  Name  string `json:"name"`
// 	  Class int 
// 	  Age   int    `json:"age"`
// 	  Codeco codeco `json:"codeco" mapper:"m"`
// 	}
  
// 	Student struct {
// 	  Name  string `json:"name"`
// 	  Class int 
// 	  Age   int  `json:"age"`
// 	  Codeco1 codecoS `json:"codecos" mapper:"m"`
// 	}
//   )
  
  
//   func main() {
// 	user := User{Name: "abc", Class: 90, Age: 10, Codeco: codeco{Name1: "helo", Name12: "ff", Mine: mine{Base: "goo"}}}
// 	student := Student{Name: "ABC", Class: 234, Age: 1000000000, Codeco1: codecoS{Name134: "hh"}}

// 	// fmt.Println(user)
  
// 	// create mapper object
// 	m := mapper.NewMapper()
  
// 	// in the version < v0.7.8, we will use field name as key when mapping structs
// 	// we keep it as default behavior in this version
// 	// m.SetEnableIgnoreFieldTag(true)
  
// 	// student.Age = 3
  
// 	// disable the json tag
// 	// m.SetEnabledJsonTag(false)
  
// 	// student::age should be 1
// 	m.Mapper(&user.Codeco, &student.Codeco1)
  
// 	fmt.Println("user:")
// 	fmt.Println(&user)
// 	fmt.Println("student:")
// 	fmt.Println(&student)



	
// 	PrintFields(user)

	
	
//   }

//   func PrintFields(user User) {
// 	val := reflect.ValueOf(user)
// 	for i := 0; i < val.Type().NumField(); i++ {
// 	   fmt.Println(val.Type().Field(i).Tag.Get("mapper"))
// 	}
//   }