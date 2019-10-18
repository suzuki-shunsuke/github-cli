package watch

import (
	"github.com/urfave/cli"
)

func Subscribe(c *cli.Context) error {
	return setSubscription(c, false)
}
