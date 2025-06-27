package ulog

import (
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func FileLog(logPath string, logName string, logConfig *FileConfig) *File {
	f := &File{}
	f.Init(logPath, logName, logConfig)
	return f
}

////////////////////////////////////////////////////////////////////////

type FileConfig struct {
	MaxSize int `yaml:"max_file_size"` //日志大小MB
	MaxAge  int `yaml:"max_save_day"`  //保留天数
}

type File struct {
	wirter *lumberjack.Logger
}

func (f *File) Init(logPath string, logName string, logConfig *FileConfig) {
	os.MkdirAll(logPath, 0755)

	logFilePath := filepath.Join(logPath, logName)

	f.wirter = &lumberjack.Logger{
		Filename: logFilePath,
		MaxSize:  logConfig.MaxSize,
		MaxAge:   logConfig.MaxAge,
		Compress: true,
	}
}

func (f *File) Write(data []byte) (int, error) {
	return f.wirter.Write(data)
}
