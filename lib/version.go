package lib

import "fmt"

var (
	// Logo
	Logo = `
Glance V%s
https://github.com/TimothyYe/glance
`
)

func DisplayVersion(version string) {
	fmt.Printf(Logo, version)
}
