package categorify

import "fmt"

func ExampleLookup() {
	r, err := Lookup("ilyaglotov.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Domain)
	fmt.Println(r.Confidence)
	fmt.Println(r.Rating.Description)
	// Output:
	// ilyaglotov.com
	// low
	// Safe for all audiences.
}
