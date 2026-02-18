package duration

import "fmt"

type duration int

// Import
func FromSekunde(t int) duration {
	return duration(t)
}

func FromMinute(t int) duration {
	return duration(t * 60)
}

func FromStunde(t int) duration {
	return duration(t * 3600)
}

func FromTag(t int) duration {
	return duration(t * 86400)
}

func FromJahr(t int) duration {
	return duration(t * 31536000)
}

// Export
func (l duration) Sekunde() int {
	return int(l)
}

func (l duration) Minute() int {
	return l.Sekunde() / 60
}

func (l duration) Stunde() int {
	return l.Sekunde() / 3600
}

func (l duration) Tag() int {
	return l.Sekunde() / 86400
}

func (l *duration) Scale(factor int) {
	//
	*l = duration(l.Sekunde() * factor)

}
func Example() {

	a := FromSekunde(3000000)
	b := FromMinute(50000)

	fmt.Println(a.Sekunde())
	//fmt.Println("s")
	fmt.Println(a.Minute())
	//fmt.Println("m")
	fmt.Println(a.Stunde())
	//fmt.Println("h")
	fmt.Println(a.Tag())
	//fmt.Println("d")

	fmt.Println(b.Sekunde())
	fmt.Println(b.Minute())
	fmt.Println(b.Stunde())
	fmt.Println(b.Tag())

	//Output:
}
