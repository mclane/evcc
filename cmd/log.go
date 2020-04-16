package cmd

import (
	"fmt"
	"os"

	"github.com/andig/evcc/api"
	"github.com/spf13/cobra"
)

type LogConfig struct {
	fileName string
	file     *os.File
}

func (l *LogConfig) Configure(cmd *cobra.Command) {
	if level, err := cmd.PersistentFlags().GetString("log"); err == nil {
		fmt.Printf("level: %s\n", level)
		l.Level(level)
	} else {
		log.FATAL.Fatal(err)
	}

	if logfile, err := cmd.PersistentFlags().GetString("logfile"); err == nil {
		fmt.Printf("logfile: %s\n", logfile)
		l.File(logfile)
	} else {
		log.FATAL.Fatal(err)
	}
}

func (l *LogConfig) Level(level string) {
	api.OutThreshold = api.LogLevelToThreshold(level)
	api.LogThreshold = api.OutThreshold

	api.Loggers(func(name string, logger *api.Logger) {
		logger.SetStdoutThreshold(api.OutThreshold)
		logger.SetLogThreshold(api.LogThreshold)
	})
}

func (l *LogConfig) File(logfile string) {
	if l.file != nil {
		println("closing")
		if err := l.file.Close(); err != nil {
			log.ERROR.Printf("error closing log file: %v", err)
		}
		l.file = nil
	}

	if logfile != "" {
		println("reopening")

		f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.FATAL.Fatalf("error opening log file: %v", err)
		}
		l.file = f
	}

	l.fileName = logfile
	log.SetLogOutput(l.file)
}
