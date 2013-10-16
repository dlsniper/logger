// Copyright 2013 motain GmbH. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"github.com/motain/amqp"
	"log"
	"log/syslog"
)

// Log to classic log way
const GOLOG = 1

// Log to syslog
const SYSLOG = 2

// Log to AMQP
const AMQP = 4

// Log to all
const ALL = 7

// Logger is the basic structure for the logger
type Logger struct {
	prefix      string
	goLog       bool
	sysLog      *syslog.Writer
	amqpChannel *amqp.Channel
}

// Init the logger prefix
func (logger *Logger) Init(prefix string) {
	logger.prefix = prefix
}

// EnableGoLog should be called if you want to enable or not the GOlang logger
func (logger *Logger) EnableGoLog(golog bool) {
	logger.goLog = golog
}

// SetSyslog should be called if you want to use the syslog logging facility
func (logger *Logger) SetSyslog(sysLog *syslog.Writer) {
	logger.sysLog = sysLog
}

// SetAmqp should be called if you want to use the amqp logging facility
func (logger *Logger) SetAmqp(amqpChannel *amqp.Channel) {
	logger.amqpChannel = amqpChannel
}

// log is our internal function where the logging takes place
func (logger *Logger) log(loggerType uint, logPriority syslog.Priority, message string) error {
	var err error

	var priorityString string

	switch logPriority {
	case syslog.LOG_DEBUG:
		priorityString = "DEBUG"
	case syslog.LOG_INFO:
		priorityString = "INFO"
	case syslog.LOG_NOTICE:
		priorityString = "NOTICE"
	case syslog.LOG_WARNING:
		priorityString = "WARNING"
	case syslog.LOG_ERR:
		priorityString = "ERR"
	case syslog.LOG_CRIT:
		priorityString = "CRIT"
	case syslog.LOG_ALERT:
		priorityString = "ALERT"
	case syslog.LOG_EMERG:
		priorityString = "EMERG"
	}

	if (loggerType&GOLOG) != 0 && logger.goLog {
		log.Printf("[%s] [%s] %s", logger.prefix, priorityString, message)
	}

	if (loggerType&SYSLOG) != 0 && logger.sysLog != nil {
		switch logPriority {
		case syslog.LOG_DEBUG:
			err = logger.sysLog.Debug(message)
		case syslog.LOG_INFO:
			err = logger.sysLog.Info(message)
		case syslog.LOG_NOTICE:
			err = logger.sysLog.Notice(message)
		case syslog.LOG_WARNING:
			err = logger.sysLog.Warning(message)
		case syslog.LOG_ERR:
			err = logger.sysLog.Err(message)
		case syslog.LOG_CRIT:
			err = logger.sysLog.Crit(message)
		case syslog.LOG_ALERT:
			err = logger.sysLog.Alert(message)
		case syslog.LOG_EMERG:
			err = logger.sysLog.Emerg(message)
		}

	}

	if (loggerType&AMQP) != 0 && logger.amqpChannel != nil {
		err = logger.amqpChannel.Publish(
			"logs",
			priorityString,
			false,
			false,
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "",
				Body:            []byte(message),
				DeliveryMode:    amqp.Transient,
				Priority:        0,
			},
		)
	}

	return err
}

// Debug messages should be logged with this function
func (logger *Logger) Debug(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_DEBUG, message)
}

// Debugf is just like a wrapper over Debug with a fmt.Sprintf()
func (logger *Logger) Debugf(loggerType uint, format string, a ...interface{}) error {
	return logger.Debug(loggerType, fmt.Sprintf(format, a...))
}

// Info messages should be logged with this function
func (logger *Logger) Info(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_INFO, message)
}

// Infof is just like a wrapper over Info with a fmt.Sprintf()
func (logger *Logger) Infof(loggerType uint, format string, a ...interface{}) error {
	return logger.Info(loggerType, fmt.Sprintf(format, a...))
}

// Notice messages should be logged with this function
func (logger *Logger) Notice(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_NOTICE, message)
}

// Noticef is just like a wrapper over Notice with a fmt.Sprintf()
func (logger *Logger) Noticef(loggerType uint, format string, a ...interface{}) error {
	return logger.Notice(loggerType, fmt.Sprintf(format, a...))
}

// Warning messages should be logged with this function
func (logger *Logger) Warning(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_WARNING, message)
}

// Warningf is just like a wrapper over Warning with a fmt.Sprintf()
func (logger *Logger) Warningf(loggerType uint, format string, a ...interface{}) error {
	return logger.Warning(loggerType, fmt.Sprintf(format, a...))
}

// Err messages should be logged with this function
func (logger *Logger) Err(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_ERR, message)
}

// Errf is just like a wrapper over Err with a fmt.Sprintf()
func (logger *Logger) Errf(loggerType uint, format string, a ...interface{}) error {
	return logger.Err(loggerType, fmt.Sprintf(format, a...))
}

// Crit messages should be logged with this function
func (logger *Logger) Crit(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_CRIT, message)
}

// Critf is just like a wrapper over Crit with a fmt.Sprintf()
func (logger *Logger) Critf(loggerType uint, format string, a ...interface{}) error {
	return logger.Crit(loggerType, fmt.Sprintf(format, a...))
}

// Alert messages should be logged with this function
func (logger *Logger) Alert(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_ALERT, message)
}

// Alertf is just like a wrapper over Alert with a fmt.Sprintf()
func (logger *Logger) Alertf(loggerType uint, format string, a ...interface{}) error {
	return logger.Alert(loggerType, fmt.Sprintf(format, a...))
}

// Emerg messages should be logged with this function
func (logger *Logger) Emerg(loggerType uint, message string) error {
	return logger.log(loggerType, syslog.LOG_EMERG, message)
}

// Emergf is just like a wrapper over Emerg with a fmt.Sprintf()
func (logger *Logger) Emergf(loggerType uint, format string, a ...interface{}) error {
	return logger.Emerg(loggerType, fmt.Sprintf(format, a...))
}
