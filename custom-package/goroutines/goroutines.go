package goroutines

import (
	"fmt"
)

type Student struct {
	name string
	old  int
}

func StructType() {
	// named
	st1 := Student{
		name: "brian",
		old:  25,
	}
	fmt.Println(st1)
	fmt.Println(st1.name)
	fmt.Println(st1.old)
	//
	st2 := Student{"brian2", 252}
	fmt.Println(st2)

	//
	var st3 Student = struct {
		name string
		old  int
	}{
		name: "brian",
		old:  25,
	}
	fmt.Println(st3)

	// anonymous struct
	var st4 = struct {
		name string
		old  int
	}{
		name: "brian",
		old:  25,
	}
	fmt.Println(st4)

	// pointer struct

	st5 := &Student{
		name: "brian",
		old:  25,
	}
	fmt.Println(&st5)
	fmt.Println(st5.name)
	fmt.Println((*st5).name)
	fmt.Println(st5.old)
	fmt.Println((*st5).old)

	// anonymous field

	type NoName struct {
		string
		int
	}
	n := NoName{"eqweqwe", 323}
	fmt.Println(n)

	// struct lòng struct
	type Info struct {
		class string
		name  string
	}

	type School struct {
		id   string
		info Info
	}

	x := School{
		id: "fsdfsd",
		info: Info{
			class: "vvvvvv",
			name:  "aaaaa",
		},
	}
	fmt.Println(x)

	// compare struct khi kiểu dữ liệu có thể so sánh dc (int, string, ...)
	type student1 struct {
		id   int
		name string
	}

	ss1 := student1{
		id:   111,
		name: "aaa",
	}

	ss2 := student1{
		id:   111,
		name: "aaa",
	}

	if ss1 == ss2 {
		fmt.Println("ok bằng")
	} else {
		fmt.Println("ko bằng")
	}

	//
	// type student2 struct {
	// 	id   int
	// 	name string
	// 	maps map[int]int
	// }

	// ss3 := student2{
	// 	id:   111,
	// 	name: "aaa",
	// 	maps: map[int]int{
	// 		1: 1,
	// 		2: 2,
	// 	},
	// }

	// ss4 := student2{
	// 	id:   111,
	// 	name: "aaa",
	// 	maps: map[int]int{
	// 		1: 1,
	// 		2: 2,
	// 	},
	// }
	// error
	// if ss3 == ss4 {
	// 	fmt.Println("ok bằng")
	// } else {
	// 	fmt.Println("ko bằng")
	// }

	// zero value theo kiểu của struct
}

func Method() {
	fmt.Println("Method")
}

func Goroutine1() {
	fmt.Println("Goroutine 111111")
}

func Goroutine2() {
	fmt.Println("Goroutine 222222")
}

func Variables() {
	fmt.Println("Variables")
	var (
		name string
		old  int
	)

	var (
		name1 string = "Brian2"
		old1  int    = 26
	)

	const (
		address1 string = "aaaa"
		address2 string = "bbbb"
	)

	name = "brian"
	old = 25

	var name2, old2 = "nguyen", 27

	fmt.Println("Channel", name, old, name1, old1, name2, old2)
}

func MapType() {
	var myMap = make(map[string]int) // dùng make đã dc khởi tạo vào bộ nhớ | giá trị zero value không phải nil
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("myMap == nil")
	} else {
		fmt.Println("myMap != nil")
	}

	var myMap1 map[string]int // chưa dc khởi tạo trong bô nhớ | giá trị zero value là nil
	if myMap1 == nil {
		fmt.Println("myMap1 == nil")
	}
	myMap1["key1"] = 1
	myMap1["key2"] = 2
	myMap1["key3"] = 3
	myMap1["key4"] = 4
	// delete
	delete(myMap1, "key1")
	// map là reference type

	value, found := myMap1["key5"]
	if found {
		fmt.Println("tim thay key", value)
	} else {
		fmt.Println("khong tim thay key")

	}

	// trong map ko có toán tử ==
}

func DataType() {
	// byte alias for uint8 (0->255)
	// rune alias for int32

	// var text rune = "A"
	// fmt.Println(text)
	// fmt.Println("DataType")
}

func PointerType() {
	a := 1000
	var pointer *int
	pointer = &a
	fmt.Println(pointer)

	b := 1000
	pointer2 := new(int)
	pointer2 = &b
	fmt.Println(pointer2)

	// zero value
	var pointerZoroValue *int // nil
	fmt.Println(pointerZoroValue)
	pointerZoroValue2 := new(int) // not nil nó sẽ khởi tạo 1 vị trí trong bộ nhớ
	fmt.Println(pointerZoroValue2)

	//
	var p *int
	s := 100
	p = &s
	fmt.Println(p)
	*p = 999 // gán giá trị mới cho pointer p thì value làm thay đổi value s vì đang trỏ tới vi trí của s
	fmt.Println(p)
	fmt.Println(s)

	// pointer tới array
	array := [3]int{1, 2, 3}
	var pArray *[3]int
	pArray = &array
	fmt.Println(pArray)

	c := 33
	var pointer3 *int = &c
	ApplyPointer(pointer3)
	fmt.Println(c) // biến c bị thay đổi thành 777

}

func ApplyPointer(pointer *int) {
	*pointer = 777
	fmt.Println(pointer)
}

func ArraySliceType() int {
	// mang 2 chieu
	// array2 := [4][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// 	{7, 8},
	// }

	// array => value type
	// slice => reference type (lam thay doi vì la reference type (con tro))
	var array1 = [4]int{1, 2, 3, 4}
	slice1 := array1[:] // type slice va cach tao slice
	slice1[0] = 123
	// slice2 := array1 // type array
	fmt.Println("slice1", slice1)
	// fmt.Printf("slice %T", slice1)
	fmt.Println("array", array1)

	// (capacity) cap bộ nhớ mở rộng để chuẩn bị sẵn cho việc thêm phần tử vào slice
	// (slice) đoạn code này sao cap(slice 13) lại ra 4 ko
	// slice13 := make([]int, 2)
	// fmt.Println(slice13)
	// slice13 = append(slice13, 100)
	// fmt.Println(slice13)
	// fmt.Println(len(slice13))
	// fmt.Println(cap(slice13))
	// because
	// Câu hỏi hay lắm bạn, bạn hiểu rằng cap (capacity) là con số giúp tối ưu việc reslice khi hàm append được gọi, giả sử mình khai báo 1 slice có cap là 1 (tức là sức chứa của slice sẽ là 1 phần tử)  thì khi mình thêm 1 phần thử mới thì slice nó sẽ phải tạo ra 1 array mới có 2 phân tử rồi nó copy phần tử đẩu tiên qua rồi add phần tử mới bạn muốn thêm vào. Bạn hãy tưởng tượng nếu bạn gọi append cả ngàn phần tử thì chương trình sẽ cấp phát các array mới liên tục điều này ko tốt chút nào, do đó slice nó sinh ra cái cơ chế  grow slice và con số capacity nó thể hiện điều này. Khi bạn khai báo slice13 := make([]int, 2)
	// thì lúc nào cap=2, khi bạn  append(slice13, 100)
	// tức là lúc này len(slice) = 3 trong khi cap=2 nó ko đủ chỗ chứa phần tử 100,  thì lúc này slice nó tự động grow (mở rộng) lên là 4 (double cái con số 2 lên) và khi số 100 nó được thêm vào tức là len(slice) = 3 như vậy slice nó ko cần tạo ra một array mới khi thêm 1 phần tử ở tương lai nữa vì nó có thể chứa tới 4 phần tử. khi mình tiếp tục gọi slice13 = append(slice13, 101) thì lúc này len(slice) = 4 và cái do capacity = 4 vẫn đủ chưa nên nó ko cần phải tạo ra một array mới nữa, tuy nhiên nếu bạn tiếp tục gọi gọi
	// slice13 = append(slice13, 102) thì lúc này len(slice)=5 vượt quá capacity=4 ban đầu thì lúc này cơ chế grow slice sẽ lại xảy ra nó lại tiếp tục double cái capacity ban đầu tức là lúc này new_capacity=4*2=8 cứ thế quá trình này sẽ tiếp diễn nếu bạn còn gọi hàm append . Trường hợp nếu cap(slice) lớn hơn 1024 thì lúc này con số grow slice được tính theo công thức khác bạn tham khảo cái hàm growSlice ở đây nhé https://golang.org/src/runtime/slice.go; ====>newcap += newcap / 4 = con số này có thể khác nhau giữa các máy tính của bạn vì nó bị phụ thuộc vào bộ nhớ .

	return 1
}

func Channel() {
	fmt.Println("Channel")
}

func Item() {
	// AddItem(1, 100, 101, 102, 103)
	var list = []int{100, 101, 102, 1003}
	// not use array [4]int{100, 101, 102, 1003} because is (list ...int) is slice
	AddItem(1, list...)
}

func AddItem(item int, list ...int) {
	list = append(list, item)
	fmt.Println("list", list)
}
