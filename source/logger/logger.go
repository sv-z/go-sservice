package logger

import (
	"log"
	"strings"
	"time"
)

const (
	LEVEL_DEBUG     = 100
	LEVEL_INFO      = 200
	LEVEL_NOTICE    = 250
	LEVEL_WARNING   = 300
	LEVEL_ERROR     = 400
	LEVEL_CRITICAL  = 500
	LEVEL_ALERT     = 550
	LEVEL_EMERGENCY = 60
)

type Logger struct {
	logger *log.Logger
	chanel string
	format string
}

func New(chanel string) (*Logger, error) {
	var logger Logger
	logger.chanel = chanel
	//logger.format = "[%timestamp%] %chanel%:%level% %message% \n"
	logger.format = "%chanel%:%level% %message% \n"

	return &logger, nil
}

func (this *Logger) Debug(m string) {
	this.log(LEVEL_DEBUG, m)
}

func (this *Logger) Info(m string) {
	this.log(LEVEL_INFO, m)
}

func (this *Logger) Notice(m string) {
	this.log(LEVEL_NOTICE, m)
}

func (this *Logger) Warning(m string) {
	this.log(LEVEL_WARNING, m)
}

func (this *Logger) Err(m string) {
	this.log(LEVEL_ERROR, m)
}

func (this *Logger) Critical(m string) {
	this.log(LEVEL_CRITICAL, m)
}

func (this *Logger) Alert(m string) {
	this.log(LEVEL_ALERT, m)
}

func (this *Logger) Emerg(m string) {
	this.log(LEVEL_EMERGENCY, m)
}

func (this *Logger) log(level int, message string) {
	format := strings.Replace(this.format, "%timestamp%", time.Now().UTC().Format(time.RFC3339), -1)
	format = strings.Replace(format, "%chanel%", this.chanel, -1)
	format = strings.Replace(format, "%level%", this.strLevel(level), -1)
	format = strings.Replace(format, "%message%", message, -1)

	log.Print(format)
}

func (this *Logger) strLevel(level int) string {

	str := "NON"

	switch level {
	case LEVEL_DEBUG:
		str = "DBG"
	case LEVEL_NOTICE:
		str = "NOT"
	case LEVEL_INFO:
		str = "INF"
	case LEVEL_WARNING:
		str = "WRN"
	case LEVEL_ERROR:
		str = "ERR"
	case LEVEL_CRITICAL:
		str = "CRT"
	case LEVEL_ALERT:
		str = "ALR"
	case LEVEL_EMERGENCY:
		str = "EMR"
	}

	return str
}
