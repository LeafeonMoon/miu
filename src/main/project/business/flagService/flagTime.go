package flagService

import (
	"flag"
	"fmt"
	"miu/project/business/flagBase"
	"miu/ui"
	"miu/utils/utime"
	"time"
)

func TimeAction() {
	fs := flag.NewFlagSet("time", flag.ContinueOnError)

	var ts string
	fs.StringVar(&ts, "ts", "ms", "")

	fs.Usage = timeHelper

	flagBase.FlagSetParse(fs)

	args := fs.Args()

	if len(args) == 0 {
		now := time.Now()
		switch ts {
		case "s":
			fmt.Printf("[%s    %d s]\r\n", utime.Full(now), utime.TimeS(now))
		case "ms":
			fmt.Printf("[%s    %d ms]\r\n", utime.Full(now), utime.TimeMS(now))
		default:
			fs.Usage()
		}
	} else {
		for _, arg := range args {
			dt, err := utime.StringToDate(arg)
			if err != nil {
				fmt.Printf("[%s    It is not a date]\r\n", arg)
			} else {
				switch ts {
				case "s":
					fmt.Printf("[%s    %d s]\r\n", arg, utime.TimeS(dt))
				case "ms":
					fmt.Printf("[%s    %d ms]\r\n", arg, utime.TimeMS(dt))
				default:
					fs.Usage()
				}
			}
		}
	}
}

func timeHelper() {
	ui.Render([]string{
		ui.TAB,
		"",
		"Usage of time:",
		"  miu time [Options]",
		"",
		"Options:",
		"  -ts    Convert time to timestamp. Values s, ms. [default ms]",
		"",
		ui.TAB,
		"",
		"Example:",
		"",
		"  1: miu time",
		"    result: [Mon Jan 02 2006 15:04:05    1136185445000 ms]",
		"",
		"  2: miu time '2006-01-02 15:04:05'",
		"    result: [2006-01-02 15:04:05    1136185445000 ms]",
		"",
		"  3: miu time '2006-01-02 15:04:05' '2006-01-02 15:04:05'",
		"    result: [2006-01-02 15:04:05    1136185445000 ms]",
		"            [2006-01-02 15:04:05    1136185445000 ms]",
		"",
		"  4: miu time '2006-01-02 aa:bb:cc'",
		"    result: [2006-01-02 aa:bb:cc    It is not a date]",
		"",
		ui.TAB,
	})
}
