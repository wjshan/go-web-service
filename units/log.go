package units

import log "github.com/sirupsen/logrus"

var LogMap map[string]*log.Logger

func GetLog(name string) *log.Logger {
	if logger, ok := LogMap[name]; ok {
		return logger
	}
	logger := log.New()
	LogMap[name] = logger
	log.SetFormatter(&log.JSONFormatter{})
	return logger
}
