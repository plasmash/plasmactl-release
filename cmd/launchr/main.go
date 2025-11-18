// Release executes Launchr application.
package main

import (
	"github.com/launchrctl/launchr"

	_ "github.com/plasmash/plasmactl-release"
)

func main() {
	launchr.RunAndExit()
}
