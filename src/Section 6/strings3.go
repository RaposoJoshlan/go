package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love Go Programming", "lovex")

	p(result)

	result = strings.ContainsAny("success", "sy")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PYTHON TO JAVA"))
	p(strings.ToUpper("GO python TO JAVA"))

	p("go" == "go")
	p("GO" == "go")
	p(strings.ToLower("GO") == strings.ToLower("go"))

	p(strings.EqualFold("Go", "go"))

	stringsPart2()

}

func stringsPart2() {

	p := fmt.Println

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.138.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("192.138.0.1", ".", "/")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)
	s = strings.Split("Go for go", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "golang"}
	myStr = strings.Join(s, "XXXX")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t Goodbye windows, welcome linus! \n")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("....Hello!!!!!!?", ".!?")
	p(s2)
}
