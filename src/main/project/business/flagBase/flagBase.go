package flagBase

import (
	"flag"
	"os"
)

func FlagSetParse(fs *flag.FlagSet) error {
	return fs.Parse(os.Args[2:])
}
