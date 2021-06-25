package main

import (
	"miu/project/business/flagService"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		flagService.MainAction()
	} else {
		switch args[1] {
		case "time":
			flagService.TimeAction()
		case "md5":
			flagService.MD5Action()
		case "b64":
			flagService.B64Action()
		case "genpd":
			flagService.GenPasswordAction()
		default:
			flagService.MainAction()
		}
	}
}
