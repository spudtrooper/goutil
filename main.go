// Helper binary for using github.com/spudtrooper/goutil.
package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/cli"
)

func main() {
	check.Err(cli.Main(context.Background()))
}
