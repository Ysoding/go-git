package internal

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Init(cCtx *cli.Context) error {
	fmt.Println("init test")
	return nil
}
