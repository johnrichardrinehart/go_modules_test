package pkg

import "fmt"

const PKG = "PKG2"
const VERSION = "0.0.2"

func PackageAndVersion() string {
	return fmt.Sprintf("Package %s: Version %s", PKG, VERSION)
}
