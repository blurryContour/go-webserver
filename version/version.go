package version

import "fmt"

const (
	devVersion = "dev"
)

var cliVersion = devVersion

func getVersion() string { return fmt.Sprintf("Version: %v", cliVersion) }
