package flagService

import (
	"bytes"
	"flag"
	"fmt"
	"miu/project/business/flagBase"
	"miu/ui"
	"miu/utils/urandom"
)

func GenPasswordAction() {
	fs := flag.NewFlagSet("genpd", flag.ContinueOnError)

	var lower bool
	fs.BoolVar(&lower, "l", false, "")

	var upper bool
	fs.BoolVar(&upper, "u", false, "")

	var number bool
	fs.BoolVar(&number, "n", false, "")

	var special bool
	fs.BoolVar(&special, "s", false, "")

	var length int
	fs.IntVar(&length, "len", 20, "")

	var times int
	fs.IntVar(&times, "i", 1, "")

	fs.Usage = benPasswordHelper

	flagBase.FlagSetParse(fs)

	var buffer bytes.Buffer
	if lower {
		buffer.WriteString(urandom.LOWER)
	}
	if upper {
		buffer.WriteString(urandom.UPPER)
	}
	if number {
		buffer.WriteString(urandom.NUMBER)
	}
	if special {
		buffer.WriteString(urandom.SPECIAL)
	}
	if len(buffer.String()) == 0 {
		buffer.WriteString(urandom.LOWER)
		buffer.WriteString(urandom.NUMBER)
	}

	if length <= 8 {
		length = 8
	} else if length > 100 {
		length = 100
	}
	if times <= 0 || times > 100 {
		times = 1
	}

	result := urandom.GenStringTimes(buffer.String(), length, times)
	for _, v := range result {
		fmt.Printf("%s\r\n", v)
	}
}

func benPasswordHelper() {
	ui.Render([]string{
		ui.TAB,
		"",
		"Usage of md5:",
		"  miu genpd [Options]",
		"",
		"Options:",
		"  -l      Randomly generate lowercase letters.",
		"  -u      Randomly generate uppercase letters.",
		"  -n      Randomly generate uppercase letters.",
		"  -s      Randomly generate special characters. (~!@#$%^&*()_+-=[]{}|;:,.<>?)",
		"  -len    Generate a password with a length of 8~100.  [default 20]",
		"  -i      Number of generated passwords. [default 20] [max 100]",
		"",
		ui.TAB,
		"",
		"Example:",
		"",
		"  1: miu genpd",
		"    result: m3nqmfxtl9ibfmq8bh46",
		"  Generate a password consisting of lowercase letters and numbers.",
		"",
		"  2: miu genpd -l",
		"    result: qceijrxxyzvilnajnnsy",
		"",
		"  3: miu genpd -i 2 -len 8",
		"    result: q8i9onvf",
		"            194tzgu8",
		"",
		ui.TAB,
	})
}
