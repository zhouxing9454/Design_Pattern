package Builder

func ExampleBuilder() {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	// Output:
	// part1
	// part2
	// part3
}
