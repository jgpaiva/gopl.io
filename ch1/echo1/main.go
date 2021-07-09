// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func write(out io.Writer) {
	out.Write([]byte("hi"))
	time.Sleep(time.Duration(rand.Float64()*3000) * time.Millisecond)
	out.Write([]byte(" there\n"))
}

type Z int

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func removeDup(s []string) []string {
	if len(s) < 2 {
		return s
	}
	p := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[p] {
			p++
		}
		s[p] = s[i]
	}
	return s[:p+1]
}

type Point struct {
	X int
	Y int
}

type Point2 struct {
	X int
	Y int
}

type Zz struct {
	Point
	Point2
	A int
	X int
}

func c() {
	a := Zz{A: 10, X: 11, Point: Point{Y: 12, X: 13}}
	fmt.Printf("%#v\n", a)
	a.Point.X = 3
	b, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	a = Zz{}
	//b[3] = byte('{')
	err = json.Unmarshal(b, &a)
	if err != nil {
		log.Fatalf("unmarshaling: %v", err)
	}
	fmt.Printf("%#v\n", a)
}

func recur(i int) {
	if i%100 == 0 {
		fmt.Println(i)
	}
	recur(i + 1)
}

func retfunc() (func() int, func() int) {
	i := 0
	return func() int {
			i++
			return i
		},
		func() int {
			i--
			return i
		}
}

func variadic(a int, b ...int) int {
	for _, n := range b {
		a += n
	}
	return a
}

func panic1() {
	panic("panic1")
}

func panic2() {
	defer func() { panic("panic2") }()
	panic1()
}

func noreturn() (retval int) {
	defer func() {
		p := recover()
		if p == nil {

		} else {
			retval = 1
		}
	}()
	for i := 0; ; i++ {
		if i == 10 {
			panic("hi")
		}
	}
}

func (p Point) printPoint() {
	fmt.Println(p.X, p.Y)
}

func (p Zz) printZ() {
	fmt.Println(p.X, p.A)
}

type Stuff struct {
	Point
	X int
}

type Pointer interface {
	printPoint()
}

type Pointer2 interface {
	printPoint()
}

type Ss struct {
	i int
	j int
}
type S []Ss

func (s S) Len() int {
	return len(s)
}

func (s S) Less(i, j int) bool {
	return s[i].i < s[j].i
}

func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func stuff() (a int, ok bool) {
	return 1, true
}
func s(a int) {
	fmt.Println(a)
}

func writeTo(i int, c chan int) {
	c <- i
}

func pingPong(in <-chan int, out chan<- int, done <-chan struct{}) {
	consume := in
	var pub chan<- int
	var v int
	for {
		select {
		case v = <-consume:
			v += 1
			consume = nil
			pub = out
		case pub <- v:
			pub = nil
			consume = in
		case <-done:
			if pub != nil { //I own the token
				pub <- v
			}
			return
		}
	}
}

func main() {
	b := make([]byte, 1)
	os.Stdin.Read(b) //read single byte
	fmt.Printf("%q", b)

	/*
		a := make(chan int)
		b := make(chan int)
		done := make(chan struct{})
		go pingPong(a, b, done)
		go pingPong(b, a, done)
		a <- 1
		time.Sleep(10 * time.Second)
		close(done)
		select {
		case v := <-a:
			fmt.Println(v)
		case v := <-b:
			fmt.Println(v)
		}
		/*
			start := make(chan int)
			prev := start
			chrono := time.Now().UTC().UnixNano()
			for i := 0; i < 2000000; i++ {
				next := make(chan int)
				go func(out chan<- int, in <-chan int) {
					v := <-in
					out <- v
				}(next, prev)
				prev = next
			}
			fmt.Println((time.Now().UTC().UnixNano() - chrono) / int64(time.Second))
			fmt.Println("setup done")
			done := make(chan struct{})
			go func(out <-chan int) {
				close(done)
			}(prev)
			os.Stdin.Read(make([]byte, 1)) //read single byte
			start <- 10
			<-done
			fmt.Println((time.Now().UTC().UnixNano() - chrono) / int64(time.Second))
			fmt.Println("done")
			/*
				for i := 0; i < 1000000; i++ {
					c := make(chan int)
					go writeTo(i, c)
					if i%100000 == 0 {
						time.Sleep(1 * time.Second)
					}
				}
				time.Sleep(1 * time.Second)
				/*
					a := make(chan int)
					type T int
					a := T(5)
					//b := interface{}(a)
					fmt.Printf("%T %[1]v\n", a)
					s(int(a))

					/*
						loc, err := time.LoadLocation("US/Eastern")
						if err != nil {
							log.Fatalf("%v", err)
						}
						t := time.Now()
						fmt.Println(t)
						fmt.Println(t.In(loc))
						fmt.Println(t)
						/*
							a := 10
							if a, ok := stuff(); ok {
								fmt.Println(a)
							}
							fmt.Println(a)
							/*
								s := S([]Ss{{2, 1}, {1, 3}, {1, 2}, {5, 1}, {1, 1}, {4, 1}})
								fmt.Println(s)
								sort.Sort(s)
								fmt.Println(s)
								sort.Sort(sort.Reverse(s))
								fmt.Println(s)

								/*
									var a io.Writer = nil
									fmt.Printf("%T %v\n", a, a)
									if a == nil {
										fmt.Println("yep")
									} else {

										fmt.Println("nope")
									}
									var b *bytes.Buffer = nil
									a = b
									fmt.Printf("%T %v\n", a, a)
									if a == nil {
										fmt.Println("yep")
									} else {

										fmt.Println("nope")
									}
									/*
										var a Pointer = Stuff{Point: Point{X: 10, Y: 11}, X: 9}
										var b Pointer2 = a
										b.printPoint()

										/*
											a := Stuff{1}
											b := Stuff{2}
											fmt.Println(a + b)
											/*
												a := Zz{}
												a.Point = Point{X: 10, Y: 20}
												fmt.Println(a.X)
												b := []Point{}
												Point.printPoint(a.Point)
												fmt.Println(b)

												/*
													fmt.Println(noreturn())
													fmt.Println(variadic(0, 2, 3))
													c()
													inc, dec := retfunc()
													fmt.Println(inc())
													fmt.Println(inc())
													fmt.Println(dec())
													fmt.Println(inc())
													// recur(0)
													/*
														zz := Zz{Point: Point{1, 2}, a: 10} //, x: 11}
														fmt.Printf("%#v\n", zz)
														zz.x = 9
														fmt.Println(zz)
														fmt.Println(zz.x)
														aa := tempconv.T{A: 10, C: "hi"}
														fmt.Println(aa)
														var ag = map[string]int{}
														fmt.Println(ag)
														fmt.Println(ag == nil)
														ag["hi"] = 10
														fmt.Println(ag)
														delete(ag, "hi")
														fmt.Println(ag)

														fmt.Println(removeDup([]string{"a", "b", "b", "b", "c", "c", "d", "e"}))
														fmt.Println(removeDup([]string{"a"}))
														fmt.Println(removeDup([]string{"a", "a", "a", "a"}))
														r1 := []int{1, 2, 3, 4, 5}
														r2 := []int{1, 2, 3, 4}
														r3 := []int{}
														rotate(r1, 1)
														rotate(r2, 2)
														//reverse(r3)
														fmt.Println(r1, r2, r3)

														//var t1 time.Duration = 3002*time.Nanosecond + 1*time.Second
														t2 := make(map[int]map[int]int)
														fmt.Println(t2[1][0])

														z := "a"
														if z == "web" {
															rand.Seed(time.Now().UTC().UnixNano())
															handler := func(w http.ResponseWriter, r *http.Request) {
																write(w)
															}
															http.HandleFunc("/", handler)
															log.Fatal(http.ListenAndServe("localhost:8000", nil))
															return
														}

														/*
															for i, arg := range os.Args {
																fmt.Println(i, arg)
															}
	*/
}

//!-
