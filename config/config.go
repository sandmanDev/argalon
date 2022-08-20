package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger

func init() {

	logger, errLoger := InitZapLogger()
	if errLoger != nil {
		log.Panic(errLoger)
	}
	defer logger.Sync()
}

//InitZapLogger - init zap logger from JSON file and set as global zap
func InitZapLogger() (*zap.Logger, error) {
	jsonFile, errConfig := os.Open("zapCongif.json")
	if errConfig != nil {
		panic(errConfig)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cfg zap.Config
	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	var err2 error
	ZapLogger, err2 = cfg.Build()
	if err2 != nil {
		panic(err2)
	}
	ZapLogger.Info("Zap logger construction succeeded")
	zap.ReplaceGlobals(ZapLogger)
	return ZapLogger, err2

}

//SyslogTimeEncoder - time stamp format for zap logger
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 01, 2006  15:04:05"))
}
