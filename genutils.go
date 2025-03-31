// dsagoutils is a package containing general purpose functions.
package genutilsgo

import (
	"fmt"
	"log/syslog"
	"os"

	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"

	yaml "gopkg.in/yaml.v2"
)

var (
	ctxlog *log.Entry
)

func init() {
	appName := "-"
	moduleName := "goutils"
	log.SetFormatter(&log.JSONFormatter{})

	hook, err := logrus_syslog.NewSyslogHook("tcp", "localhost:514", syslog.LOG_INFO, "")

	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
	standardFields := log.Fields{
		"app":        appName,
		"moduleName": moduleName,
	}
	ctxlog = log.WithFields(standardFields)
}

// ReadYaml reads a yaml formatted file into the specified structure
func ReadYaml(fn string, st interface{}) error {
	ymlFile, err := os.ReadFile(fn)
	if err != nil {
		emsg := fmt.Sprintf("Error reading %s\n", fn)
		ctxlog.Error(emsg)
	} else {
		err = yaml.UnmarshalStrict(ymlFile, st)
		if err != nil {
			emsg := fmt.Sprintf("Error unmarshalling %s\n", fn)
			ctxlog.Error(emsg)
			ctxlog.Error(err)
		}
	}
	return err
}

// WriteYaml writes a structure to a yaml formatted file
func WriteYaml(fn string, st interface{}) error {
	yaml, err := yaml.Marshal(st)
	if err != nil {
		emsg := fmt.Sprintf("Error Marshaling data you YAML %v\n", st)
		ctxlog.Error(emsg)
	} else {
		err = os.WriteFile(fn, yaml, 0644)
		if err != nil {
			emsg := fmt.Sprintf("Error writeing %s\n", fn)
			ctxlog.Error(emsg)
		}
	}
	return err
}
