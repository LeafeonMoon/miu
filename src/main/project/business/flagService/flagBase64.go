package flagService

import (
	"flag"
	"fmt"
	"miu/project/business/flagBase"
	"miu/ui"
	"miu/utils/ubase64"
)

func B64Action() {
	fs := flag.NewFlagSet("b64", flag.ContinueOnError)

	var decode bool
	fs.BoolVar(&decode, "d", false, "")

	fs.Usage = b64Helper

	flagBase.FlagSetParse(fs)
	args := fs.Args()

	if decode {
		for _, arg := range args {
			decode, err := ubase64.BasicDecode(arg)
			if err != nil {
				decode = "Decryption failed"
			}
			fmt.Printf("[%s    %s]\r\n", arg, decode)
		}
	} else {
		for _, arg := range args {
			fmt.Printf("[%s    %s]\r\n", arg, ubase64.BasicEncode(arg))
		}
	}
}

func b64Helper() {
	ui.Render([]string{
		ui.TAB,
		"",
		"Usage of b64:",
		"  miu b64 [Options]",
		"",
		"Options:",
		"  -d      Base64 decryption.",
		"",
		ui.TAB,
		"",
		"Example:",
		"",
		"  1: miu b64 123456",
		"    result: [123456    MTIzNDU2]",
		"",
		"  2: miu b64 123456 123456",
		"    result: [123456    MTIzNDU2]",
		"            [123456    MTIzNDU2]",
		"",
		"  3: miu b64 -d 123456",
		"    result: [MTIzNDU2    123456]",
		"",
		ui.TAB,
	})
}
