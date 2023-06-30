package logext

//
//import (
//	"errors"
//	"go.elastic.co/apm/module/apmlogrus/v2"
//	"io"
//	"os"
//
//	"github.com/sirupsen/logrus"
//	"gitlab.podro.com/monorepo/gojeh/pkg/configext"
//	"gitlab.podro.com/monorepo/gojeh/pkg/envext"
//	"gopkg.in/natefinch/lumberjack.v2"
//)
//
//// depricated
//type Configs struct {
//	Mode     envext.Mode  `env:"MODE,required"`
//	LogLevel logrus.Level `env:"LOG_LEVEL,required"`
//	LogFile  string       `env:"LOG_FILE,required"`
//	LogSize  int          `env:"LOG_SIZE,required"` // megabytes
//	LogAge   int          `env:"LOG_AGE,required"`  // days
//}
//
//type Logger struct {
//	*logrus.Logger
//}
//
//// depricated
//func init() {
//	configs := new(Configs)
//	if err := envext.Load(configs); err != nil {
//		logrus.WithError(err).Fatal("can load log configs")
//	}
//	if configs.Mode != envext.ModeLocal {
//		logrus.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
//			Filename: configs.LogFile,
//			MaxSize:  configs.LogSize,
//			MaxAge:   configs.LogAge,
//		}))
//		logrus.SetFormatter(new(logrus.JSONFormatter))
//	}
//	logrus.SetLevel(configs.LogLevel)
//	logrus.SetReportCaller(true)
//}
//
//func NewLog(conf configext.SectionLog) *Logger {
//	logger, err := stdoutInit(conf.Level)
//	logger.SetReportCaller(true)
//	if err != nil {
//		logrus.Panic(err)
//	}
//	return &Logger{logger}
//}
//
//func stdoutInit(lvl string) (*logrus.Logger, error) {
//	var err error
//	logger := logrus.New()
//	logger.AddHook(&apmlogrus.Hook{})
//
//	logger.SetFormatter(&logrus.TextFormatter{
//		DisableColors: false,
//		FullTimestamp: true,
//	})
//	level, err := logrus.ParseLevel(lvl)
//	if err != nil {
//		err = errors.New("failed to parse level")
//		return nil, err
//	}
//	logger.Level = level
//	var logWriter io.Writer = os.Stdout
//	logger.SetOutput(logWriter)
//
//	return logger, err
//}
