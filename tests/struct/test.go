package test

// PublicStruct
//
// Example:
//     test("PublicStruct")
type PublicStruct struct {
	// PublicStructField
	//
	// Example:
	//     test("PublicStruct.PublicStructField")
	PublicStructField string

	// privateStructField
	//
	// Example:
	//     test("PublicStruct.privateStructField")
	privateStructField string
}

// privateStruct
//
// Example:
//     test("privateStruct")
type privateStruct struct {
	// PublicStructField
	//
	// Example:
	//     test("privateStruct.PublicStructField")
	PublicStructField string

	// privateStructField
	//
	// Example:
	//     test("privateStruct.privateStructField")
	privateStructField string
}
