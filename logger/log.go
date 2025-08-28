package logger

import "fmt"

type backGroundcolor string
type color string

const (
	Reset        color = "\033[0m"
	Red          color = "\033[31m"
	Green        color = "\033[32m"
	Yellow       color = "\033[33m"
	Blue         color = "\033[34m"
	Black        color = "\033[30m"
	Magentacolor color = "\033[35m"
	Cyan         color = "\033[36m"
	White        color = "\033[37m"

	BackGroundReset   backGroundcolor = "\u001b[0m"
	BackgroundBlack   backGroundcolor = "\u001b[40m"
	BackgroundRed     backGroundcolor = "\u001b[41m"
	BackgroundGreen   backGroundcolor = "\u001b[42m"
	BackgroundYellow  backGroundcolor = "\u001b[43m"
	BackgroundBlue    backGroundcolor = "\u001b[44m"
	BackgroundMagenta backGroundcolor = "\u001b[45m"
	BackgroundCyan    backGroundcolor = "\u001b[46m"
	BackgroundWhite   backGroundcolor = "\u001b[47m"
)

type Logger struct{}

func (l *Logger) PrintColor(color color, lg string) {
	lg = fmt.Sprintf("%s %s%s", color, lg, Reset)
	fmt.Printf("%s", lg)
}
func (l *Logger) PrintColorf(color color, lg string, params ...interface{}) {
	str := formatString(lg, params...)
	l.PrintColor(color, str)
}

func (l *Logger) PrintBackground(bgColor backGroundcolor, lg string) {
	lg = fmt.Sprintf("%s %s%s", bgColor, lg, BackGroundReset)
	fmt.Printf("%s", lg)
}
func (l *Logger) PrintBackgroundf(bgColor backGroundcolor, lg string, params ...interface{}) {
	str := formatString(lg, params...)
	l.PrintBackground(bgColor, str)
}

func (l *Logger) Info(lg string) string {
	lg = fmt.Sprintf("%s %s%s", Blue, lg, Reset)
	fmt.Printf("%s", lg)
	return lg
}

func (l *Logger) Infof(lg string, params ...interface{}) string {
	str := formatString(lg, params...)
	return l.Info(str)
}

func (l *Logger) Warn(lg string) string {
	lg = fmt.Sprintf("%s %s%s", Yellow, lg, Reset)
	fmt.Printf("%s", lg)
	return lg
}

func (l *Logger) Warnf(lg string, params ...interface{}) string {
	str := formatString(lg, params...)
	return l.Warn(str)
}

func (l *Logger) Error(lg string) string {
	lg = fmt.Sprintf("%s✖️ %s %s", Red, lg, Reset)
	fmt.Printf("%s", lg)
	return lg
}
func (l *Logger) Errorf(lg string, params ...interface{}) string {
	str := formatString(lg, params...)
	return l.Error(str)
}

func (l *Logger) Success(lg string) string {
	lg = fmt.Sprintf("%s✔ %s %s", Green, lg, Reset)
	fmt.Printf("%s", lg)
	return lg
}

func (l *Logger) Successf(lg string, params ...interface{}) string {
	str := formatString(lg, params...)
	return l.Success(str)
}
func (l *Logger) Print(lg string) string {
	lg = fmt.Sprintf("%s %s %s", Reset, lg, Reset)
	fmt.Printf("%s", lg)
	return lg
}
func (l *Logger) Printf(lg string, params ...interface{}) string {
	str := formatString(lg, params...)
	return l.Print(str)
}

func formatString(lg string, params ...interface{}) string {
	str := fmt.Sprintf(lg, params...)
	return str
}
