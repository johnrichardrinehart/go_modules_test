package pkg

import "fmt"

const PKG = "PKG1"
const VERSION = "1"

func PackageAndVersion() string {
	return fmt.Sprintf("Package %s: Version %s", PKG, VERSION)
}