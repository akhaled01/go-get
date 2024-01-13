package funcs

import "flag"

var (
	BgMode       = flag.Bool("B", false, "Enables Silent mode")
	NameOfOutput = flag.String("O", "", "save under a different Name")
	PathVar      = flag.String("P", ".", "save under a different Path")
	RateLimit    = flag.String("rate-limit", "", "specify a rate at which the file is to be downloaded")
	InputFile    = flag.String("i", "", "specify a file with paths, and the program will download them async style :)")
	Mirror       = flag.Bool("mirror", false, "Mirror a website's frontend by parsing html")
	HelpMessage  = flag.Bool("help", false, "prints help message")
)
