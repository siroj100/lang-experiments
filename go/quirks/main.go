package main

import (
	"fmt"
	"strconv"
)

type (
	TheStruct struct {
		VarA int
		VarB string
	}
)

func main() {
	slice01()
	slice02()
	slice03()

	slicePtr01()
	slicePtr02()
	slicePtr03()

	map01()
	map02()
	map03()

	mapPtr01()
	mapPtr02()
	mapPtr03()

	vars()
}

func vars() {
	var (
		i1    int
		f1    float64
		b1    bool
		s1    string
		st1   TheStruct
		sl1   []int
		m1    map[int]int
		iptr1 *int
	)
	i0 := 1
	f0 := 1.0
	b0 := true
	s0 := "abc"
	st0 := TheStruct{
		VarA: 1,
		VarB: "1",
	}
	sl0 := []int{0, 1, 2, 3, 4, 5}
	m0 := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	iraw := 0
	iptr0 := &iraw
	i1 = i0
	f1 = f0
	b1 = b0
	s1 = s0
	st1 = st0
	sl1 = sl0
	m1 = m0
	iptr1 = iptr0
	i0 += 1
	f0 += 1
	b0 = false
	s0 = "def"
	st0.VarA += 1
	sl0[1] += 1
	m0[1] += 1
	*iptr0 += 1
	iraw += 1
	fmt.Println("int:", i0, i1)
	fmt.Println("float:", f0, f1)
	fmt.Println("bool:", b0, b1)
	fmt.Println("string:", s0, s1)
	fmt.Println("struct:", st0, st1)
	fmt.Println("slice:", sl0, sl1)
	fmt.Println("map:", m0, m1)
	fmt.Println("*int:", iraw, *iptr0, *iptr1)
}

func initSlice(size int) []TheStruct {
	slice := make([]TheStruct, 0)
	for i := 0; i < size; i++ {
		slice = append(slice, TheStruct{
			VarA: i,
			VarB: strconv.Itoa(i),
		})
	}
	fmt.Printf("slice: %+v\n", slice)
	return slice
}

func initSlicePtr(size int) []*TheStruct {
	slicePtr := make([]*TheStruct, 0)
	for i := 0; i < size; i++ {
		slicePtr = append(slicePtr, &TheStruct{
			VarA: i,
			VarB: strconv.Itoa(i),
		})
	}
	fmt.Printf("slicePtr: %+v\n", slicePtr)
	return slicePtr
}

func initMap(size int) map[int]TheStruct {
	result := make(map[int]TheStruct, size)
	for i := 0; i < size; i++ {
		result[i] = TheStruct{
			VarA: i,
			VarB: strconv.Itoa(i),
		}
	}
	fmt.Printf("map: %+v\n", result)
	return result
}

func initMapPtr(size int) map[int]*TheStruct {
	result := make(map[int]*TheStruct, size)
	for i := 0; i < size; i++ {
		result[i] = &TheStruct{
			VarA: i,
			VarB: strconv.Itoa(i),
		}
	}
	fmt.Printf("mapPtr: %+v\n", result)
	return result
}

func slice01() {
	fmt.Println("slice01 >>>>>")
	slice := initSlice(5)
	for i, v := range slice {
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slice01 <<<<<\n")
}

func slice02() {
	fmt.Println("slice02 >>>>>")
	slice := initSlice(5)
	for i := range slice {
		v := slice[i]
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slice02 <<<<<\n")
}

func slice03() {
	fmt.Println("slice03 >>>>>")
	slice := initSlice(5)
	for i := range slice {
		v := slice[i]
		slice[i].VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slice03 <<<<<\n")
}

func slicePtr01() {
	fmt.Println("slicePtr01 >>>>>")
	slice := initSlicePtr(5)
	for i, v := range slice {
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slicePtr01 <<<<<\n")
}

func slicePtr02() {
	fmt.Println("slicePtr02 >>>>>")
	slice := initSlicePtr(5)
	for i := range slice {
		v := slice[i]
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slicePtr02 <<<<<\n")
}

func slicePtr03() {
	fmt.Println("slicePtr03 >>>>>")
	slice := initSlicePtr(5)
	for i := range slice {
		v := slice[i]
		slice[i].VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, slice[i])
	}
	fmt.Printf("%+v\n", slice)
	fmt.Println("slicePtr03 <<<<<\n")
}

func map01() {
	fmt.Println("map01 >>>>>")
	theMap := initMap(5)
	for i, v := range theMap {
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("map01 <<<<<\n")
}

func map02() {
	fmt.Println("map02 >>>>>")
	theMap := initMap(5)
	for i := range theMap {
		v := theMap[i]
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("map02 <<<<<\n")
}

func map03() {
	fmt.Println("map03 >>>>>")
	theMap := initMap(5)
	for i := range theMap {
		v := theMap[i]
		// theMap[i].VarA += 1 -> isn't possible in Go, because map entries aren't addressable
		// so we do it slightly different
		v.VarA += 1
		theMap[i] = v
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("map03 <<<<<\n")
}

func mapPtr01() {
	fmt.Println("mapPtr01 >>>>>")
	theMap := initMapPtr(5)
	for i, v := range theMap {
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("mapPtr01 <<<<<\n")
}

func mapPtr02() {
	fmt.Println("mapPtr02 >>>>>")
	theMap := initMapPtr(5)
	for i := range theMap {
		v := theMap[i]
		v.VarA += 1
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("mapPtr02 <<<<<\n")
}

func mapPtr03() {
	fmt.Println("mapPtr03 >>>>>")
	theMap := initMapPtr(5)
	for i := range theMap {
		v := theMap[i]
		theMap[i].VarA += 1 //it's possible now because the value is a pointer, so it's addressable
		fmt.Printf("%d: %+v %+v\n", i, v, theMap[i])
	}
	fmt.Printf("%+v\n", theMap)
	fmt.Println("mapPtr03 <<<<<\n")
}
