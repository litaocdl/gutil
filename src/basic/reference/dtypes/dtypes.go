package dtypes

import (
	"fmt"
	"os"
)

func arrayType() [5]int {
	var array = [5]int{1, 2, 3, 4, 5}
	return array
}

func structType() {
	type nested struct {
		name   string
		age    int
		school struct {
			name string
			city string
		}
	}

	tao := nested{
		name: "tao",
	}

	mai := nested{
		name: "mai",
		age:  10,
	}
	wei := nested{
		name: "wei",
		age:  30,
	}
	wei.school.city = "beijing"
	wei.school.name = "BUAA"

	fmt.Println(tao)
	fmt.Println(mai)
	fmt.Println(wei)

	/**
	* annonymous fields, 字段名和类型相同，如果引入其他包的类型作为annonymous fields, 则字段名不包含包名
	* 接口类型和多级指针 不可以作为匿名类型
	 */
	type A struct {
		name string
	}
	type annonTypeA struct {
		location string
		A        //A这里是annonymous字段，类型为A，名称也是A
	}
	type annonTypeB struct {
		location string
		os.File  //这里os.File为annonmous字段，类型为os.File, 名称不包含包名，为File
	}

	annonTypeA1 := annonTypeA{
		location: "test1",
		A: A{
			name: "ASample",
		},
	}

	annonTypeB1 := annonTypeB{
		location: "test2",
		File:     os.File{},
	}
	fmt.Printf("annonTypeA name is %s \n", annonTypeA1.name) //可直接引用name

	fmt.Printf("annonTypeB  is %#v \n", annonTypeB1) //可直接引用name
}

func interf() {
	type data int
	var b data = 5
	var x interface{} = b
	switch v := x.(type) {
	case nil:
		println("nil")
	default:
		fmt.Println(v)
	}
	fmt.Println("down")
}
