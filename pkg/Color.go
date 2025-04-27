package pkg

type Color struct {
	Reset           string
	Red             string
	Orange          string
	Green           string
	Yellow          string
	Blue            string
	White           string
	Black           string
	Magenta         string
	Cyan            string
	Gray            string
	LightRed        string
	LightGreen      string
	LightYellow     string
	LightBlue       string
	LightMagenta    string
	LightCyan       string
	LightGray       string
	LightWhite      string
	LightBlack      string
	LightWhiteBold  string
	LightRedBold    string
	LightGreenBold  string
	LightYellowBold string
	LightBlueBold   string
}

var ColorCodes = Color{
	Reset:           "\033[0m",
	Red:             "\033[31m",
	Orange:          "\033[38;5;214m",
	Green:           "\033[32m",
	Yellow:          "\033[33m",
	Blue:            "\033[34m",
	White:           "\033[0;37m",
	Black:           "\033[0;30m",
	Magenta:         "\033[0;35m",
	Cyan:            "\033[0;36m",
	Gray:            "\033[0;37m",
	LightRed:        "\033[1;31m",
	LightGreen:      "\033[1;32m",
	LightYellow:     "\033[1;33m",
	LightBlue:       "\033[1;34m",
	LightMagenta:    "\033[1;35m",
	LightCyan:       "\033[1;36m",
	LightGray:       "\033[1;37m",
	LightWhite:      "\033[1;37m",
	LightBlack:      "\033[1;30m",
	LightWhiteBold:  "\033[1;37;1m",
	LightRedBold:    "\033[1;31;1m",
	LightGreenBold:  "\033[1;32;1m",
	LightYellowBold: "\033[1;33;1m",
	LightBlueBold:   "\033[1;34;1m",
}
