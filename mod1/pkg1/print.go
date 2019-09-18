package pkg

import "fmt"

const PKG = "PKG1"
const VERSION = "0.0.1"

func PackageAndVersion() string {
	return fmt.Sprintf("Package %s: Version %s", PKG, VERSION)
}
