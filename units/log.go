package units

import (
	log "github.com/sirupsen/logrus"
)

var LogInstance *logMaker = &logMaker{}

type logMaker struct {
	logMap map[string]*log.Logger
}

func (m *logMaker) GetLog(name string) *log.Logger {
	if m.logMap == nil {
		m.logMap = make(map[string]*log.Logger)
	}
	if logger, ok := m.logMap[name]; ok {
		return logger
	}
	logger := log.New()
	m.logMap[name] = logger
	log.SetFormatter(&log.JSONFormatter{})
	return logger
}
