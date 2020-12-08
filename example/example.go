// example.go file
package example

// IsLoggedIn returns true if the current user has access to the service.
//
// Example:
//     if IsLoggedIn() {
//         fmt.Println("You are logged in")
//     }
func IsLoggedIn() bool {
	return true
}

// CurrentUser returns the current username of the logged in user.
//
// Examples:
//     fmt.Println(CurrentUser())
//
//     if IsLoggedIn() {
//         fmt.Println(CurrentUser())
//     }
func CurrentUser() string {
	return "Joe Doe"
}
