package flagService

import "miu/ui"

func MainAction() {
	mainHelper()
}

func mainHelper() {
	ui.Render([]string{
		ui.TAB,
		"",
		"Usage of miu:",
		"  miu [Commands] [Options]",
		"",
		"Commands:",
		"  b64      Base64 encryption.",
		"  genpd    Password generator.",
		"  md5      MD5 encryption.",
		"  time     Time tool.",
		"",
		"Run 'miu COMMAND -h' for more information on a command.",
		"",
		ui.TAB,
	})
}
