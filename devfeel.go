// package main

// import (
// 	"fmt"

// 	"github.com/alexsem80/go-mapper/mapper"

// 	dc "github.com/fluidtruck/deepcopy"
// )

// func main() {
// 	src := Source{
// 		ID:        1,
// 		FirstName: "Name",
// 		Arr:      []work{{base: "hi"}, {base: "erw"}},
// 	}

// 	dest := &Destination{Namee: "hu"}

	
// 	mapper := mapper.NewMapper()

// 	// mapper.CreateMap((*Source)(nil), (*Destination)(nil))

// 	// mapper.Init()

// 	// mapper.Map(src, dest)

// 	// fmt.Println(src)
// 	// fmt.Println(dest)
// 	// fmt.Println(src.Work)
// 	// fmt.Println(&dest.Arr)

// 	// mapper = mapper.NewMapper()
	
// 	mapper.CreateMap((*work)(nil), (*arr)(nil)) 
// 	mapper.Init()

// 	// for i := range src.Work {
// 	// 	mapper.Map(src.Work[i], &dest.Arr[i])
// 	// }
// 	dc.DeepCopy(src.Arr, &dest.Arr)


// 	for i := range src.Arr {
// 		mapper.Map(src.Arr[i], dest.Arr[i])
// 		// dest.Arr = append(dest.Arr, destArr)
// 	}


// 		// destArr := arr{}
// 		// mapper.Map(src.Arr[i], &destArr) 
// 		// dest.Arr = append(dest.Arr, destArr)

// 	fmt.Println(src)
// 	fmt.Println(dest)
// }

// type Source struct {
// 	ID        int    `mapper:"Id"`
// 	FirstName string `mapper:"Name"`
// 	Arr      []work `mapper:"Arr"`
// }

// type Destination struct {
// 	Id    int
// 	Namee string
// 	Arr   []arr
// }

// type arr struct {
// 	base string
// }

// type work struct {
// 	base string `mapper:"base"`
// }
