package cli

type cliMessage string

const (
	help     cliMessage = ""
	inputErr cliMessage = "Can't parse your input. Err: %s"
)
