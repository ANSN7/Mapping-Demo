package main

import (
	"context"
	"fmt"

	"github.com/rentiansheng/mapper"
)

type testCopyStructInline struct {
	A align `json:"a"`
	B string `json:"b" copy:"abc"`
}
type align struct {
	A1 string `json:"a1" copy:"abc"`
}
type testCopyStructInline1 struct {
	// A1 string `json:"a1"`
	B1 string `json:"abc"`
}

type testCopyStructSrc struct {
	Int           int
	AliasCopy     string `json:"alias_copy1" copy:"alias_copy"`
	privateString string
	Strings       []testCopyStructInline `json:"arr" copy:"arr1"`
	// testCopyStructInline
	// Newfield align `json:"abc"`
}

type testCopyStructDst struct {
	Int           int
	Copy          string `json:"alias_copy"`
	privateString string
	// testCopyStructInline  `json:"obj"`
	Stringss       []testCopyStructInline1 `json:"arr1"`
	// Newfield string `json:"abc"`
	
}

func main() {
	ctx := context.Background()

	src := testCopyStructSrc{
		Int:           100,
		AliasCopy:     "alias copy",
		privateString: "private",
		Strings:       []testCopyStructInline{{A: align{A1: "weeeeeeeeeeee"}, B: "we"}, 
		{A: align{A1: "beee"}, B: "we"}},
		// Newfield: align{A1: "weeeeeeeeeeeeeeeeee"},
		// testCopyStructInline: testCopyStructInline{
		// 	A: "inline a",
		// 	B: "inline b private",
		// },
	}
	dst := testCopyStructDst{}
	// err := mapper.AllMapper(context.TODO(), src, &dst)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err := mapper.Mapper(ctx, src, &dst)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dst)

}