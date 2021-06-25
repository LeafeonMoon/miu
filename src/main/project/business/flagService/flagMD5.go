package flagService

import (
	"flag"
	"fmt"
	"miu/project/business/flagBase"
	"miu/ui"
	"miu/utils/ubase64"
	"miu/utils/umd5"
	"strings"
)

func MD5Action() {
	fs := flag.NewFlagSet("md5", flag.ContinueOnError)

	var upper bool
	fs.BoolVar(&upper, "u", false, "")

	var b64 bool
	fs.BoolVar(&b64, "b64", false, "")

	fs.Usage = md5Hepler

	flagBase.FlagSetParse(fs)
	args := fs.Args()

	if b64 {
		for _, arg := range args {
			md5Str := umd5.MD5(arg)
			if upper {
				fmt.Printf("[%s    %s]\r\n", arg, ubase64.BasicEncode(strings.ToUpper(md5Str)))
			} else {
				fmt.Printf("[%s    %s]\r\n", arg, ubase64.BasicEncode(md5Str))
			}
		}
	} else {
		for _, arg := range args {
			if upper {
				fmt.Printf("[%s    %s]\r\n", arg, strings.ToUpper(umd5.MD5(arg)))
			} else {
				fmt.Printf("[%s    %s]\r\n", arg, umd5.MD5(arg))
			}
		}
	}
}

func md5Hepler() {
	ui.Render([]string{
		ui.TAB,
		"",
		"Usage of md5:",
		"  miu md5 [Options]",
		"",
		"Options:",
		"  -u      MD5 encryption converted to uppercase.",
		"  -b64    Base64 encryption after MD5 encryption.",
		"",
		ui.TAB,
		"",
		"Example:",
		"",
		"  1: miu md5 123456",
		"    result: [123456    e10adc3949ba59abbe56e057f20f883e]",
		"",
		"  2: miu md5 123456 123456",
		"    result: [123456    e10adc3949ba59abbe56e057f20f883e]",
		"            [123456    e10adc3949ba59abbe56e057f20f883e]",
		"",
		"  3: miu md5 -b64 123456",
		"    result: [123456    ZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2U=]",
		"",
		ui.TAB,
	})
}
