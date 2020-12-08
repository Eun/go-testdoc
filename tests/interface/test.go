package test

// PublicInterface
//
// Example:
//     test("PublicInterface")
type PublicInterface interface {
	// PublicInterfaceMethod
	//
	// Example:
	//     test("PublicInterface.PublicInterfaceMethod")
	PublicInterfaceMethod()

	// privateInterfaceMethod
	//
	// Example:
	//     test("PublicInterface.privateInterfaceMethod")
	privateInterfaceMethod()
}

// privateInterface
//
// Example:
//     test("privateInterface")
type privateInterface interface {
	// PublicInterfaceMethod
	//
	// Example:
	//     test("privateInterface.PublicInterfaceMethod")
	PublicInterfaceMethod()

	// privateInterfaceMethod
	//
	// Example:
	//     test("privateInterface.privateInterfaceMethod")
	privateInterfaceMethod()
}
