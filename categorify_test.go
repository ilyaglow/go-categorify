package categorify

import "fmt"

func ExampleLookup() {
	r, err := Lookup("ilyaglotov.com")
	if err != nil {
		panic(err)
	}

	_, err = Lookup("invaliddomain")

	fmt.Println(r.Domain)
	fmt.Println(r.Confidence)
	fmt.Println(r.Rating.Description)
	fmt.Println(err.Error())
	// Output:
	// ilyaglotov.com
	// low
	// Safe for all audiences.
	// result: error, reason: invalid domain
}
