package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"unicode/utf8"

	"time"
)

func plus(a int, b int) int {
	return a + b
}

// consecutive parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func multiReturns() (int, int) {
	return 3, 7
}

func sumVariadic(nums ...int) {
	fmt.Print(nums, "  ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func intSeq() func() int {
	//closure value
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 { //base case
		return 1
	}

	return n * fact(n-1)
}

func zeroval(ival int) {
	fmt.Println("dasdasdasd", ival)
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func newPersonNoPointer(name string) person {
	return person{name: name, age: 18}
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println("Hello world")
	// fmt.Println(quote.Go())

	// variables
	var a string = "initial"
	fmt.Println("Hello world", a)
	var iB, iC int = 1, 2
	i := 1
	fmt.Println("Hello world", iB, iC)

	//constants

	const cA int64 = 500000000
	const nonType = 3423242342342
	// const eSomething int64 = 3e20
	// const divide = 3e20 / cA
	fmt.Println("cosntant", cA, nonType)

	// for loop
	for i <= 3 {
		fmt.Println("with single codition", i)
		i += 1
	}
	for j := 0; j < 10; j++ {
		fmt.Println("for of all times", j)
	}

	for i := range 3 {
		fmt.Println("for range", i)
	}

	for {
		fmt.Println("infinity loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("with continue", n)
	}

	//if - else
	if 7%2 == 0 {
		fmt.Println("it's pair")
	} else {
		fmt.Println("it's not pair")
	}

	if 8%4 == 0 {
		fmt.Println("it's pair")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch
	i = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// arrays
	var array [5]int
	fmt.Println("emp:", array)
	array[4] = 100
	fmt.Println("set:", array)
	fmt.Println("get:", array[4])
	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//Slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "metallica"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0])

	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice operator
	l := s[2:5]
	fmt.Println("sl1:", l)

	//slice up to
	l = s[:5]
	fmt.Println("sl2:", l)

	//slice up from
	l = s[2:]
	fmt.Println("sl3:", l)

	tL := []string{"g", "h", "i"}
	fmt.Println("dcl:", tL)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(tL, t2) {
		fmt.Println("t == t2")
	}

	//multidimensional
	twoDD := make([][]int, 3)
	fmt.Println("t == t2", twoDD)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoDD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// } fix then panic error for index out of range
	// fmt.Println("2d: ", twoDD)

	m := make(map[string]int)
	m["k1"] = 8
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//return zero when it doesn't exist
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	// delete by key
	delete(m, "k2")
	fmt.Println("map:", m)

	//remove all
	clear(m)
	fmt.Println("map:", m)
	// m["k2"] = 12

	//if the key is present in map
	firstMap, prs := m["k2"]
	fmt.Println("prs:", firstMap, prs)
	m["k2"] = 12
	firstMap2, prs2 := m["k2"]
	fmt.Println("prs:", firstMap2, prs2)

	//declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	//RANGE

	nums := []int{1, 2, 4}
	sum := 0
	for idx, num := range nums {
		fmt.Println("idx, num", idx, num)
		sum += num
	}
	fmt.Println("sum", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range with maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println("idx-rune-runeToString", i, c, strconv.QuoteRune(c))
	}

	//functions
	pl := plus(2, 2)
	fmt.Println("plus: ", pl)
	var ppl int = plusPlus(21, 12, 1)
	fmt.Println("plus plus: ", ppl)

	//multiple return values
	firstRet, secondRet := multiReturns() // if you want only value so you can use _ either first or second value or any value
	fmt.Println("first second", firstRet, secondRet)

	//variadic functions
	sumVariadic(1, 3, 12, 12, 1, 2)

	numss := []int{1, 2, 3, 4}
	sumVariadic(numss...)

	//CLOSURES when it acceses to variable, const or anything else out of its scope

	fmt.Println(intSeq()())
	nextInt := intSeq()    // the closure value keeps state so its value is 1
	fmt.Println(nextInt()) // kept state is 2
	fmt.Println(nextInt())
	fmt.Println(nextInt()) // ans so on
	newInts := intSeq()
	fmt.Println(newInts())
	//RECURSIVE FUNCTIONS
	fctrl := fact(3)
	fmt.Println("factorial: ", fctrl)
	// closures cab also be recursive
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

	//Pointers
	i = 1000
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	//STRINGS AND RUNES

	const r = "สวัสดี"
	fmt.Println("LEn", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%x ", r[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(r))
	for idx, runeValue := range r {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(r); i += w {
		runeValue, width := utf8.DecodeRuneInString(r[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
	//STRUCTS

	personStruct := person{name: "season", age: 50}
	fmt.Println("persona name", personStruct.name)

	sp := &personStruct
	fmt.Println("memory access for name", sp.name)

	sp.age = 324
	fmt.Println("memory access for age", sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}

	fmt.Println("dog", dog)

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))
	noPtr := newPersonNoPointer("hell fire")
	fmt.Println("attributes", noPtr.age, noPtr.name)
	//Methods
	// rc := rect{width: 10, height: 5}
	// fmt.Println("area", r.area())
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
