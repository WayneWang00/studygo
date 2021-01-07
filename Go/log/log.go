package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 如果要创建的文件已存在，会将文件清空
	logFile, err := os.Create("test.log")
	if err != nil {
		log.Fatalln("create file failed:", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "[DEUBG]", log.LstdFlags)

	logInput()
	loggerInput(logger)
	logModify(logger)
	loggerModify(logger)
	logSetOutput(logFile)
}

// 使用log.Func方法，不会输入数据到test.log文件
func logInput() {
	log.Println("a log.println message is here")
	err := log.Output(2, "a log.output message is here")
	if err != nil {
		log.Fatalln("log.output failed:", err)
	}
}

// 使用新创建的logger.Func方法，可以输入数据到test.log文件
func loggerInput(logger *log.Logger) {
	logger.Println("a logger.println message is here")
	err := logger.Output(2, "a logger.output message is here")
	if err != nil {
		logger.Println("logger.output failed:", err)
	}
}

// 使用log.setFunc，不会修改test.log输入格式
func logModify(logger *log.Logger) {
	log.SetPrefix("[INFO]")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("a log.println message is here")
	err := log.Output(2, "a log.output message is here")
	if err != nil {
		log.Fatalln("log.output failed:", err)
	}

	logger.Println("a logger.println message is here")
	err = logger.Output(2, "a logger.output message is here")
	if err != nil {
		logger.Fatalln("logger.output failed:", err)
	}
}

// 使用新创建的logger.setFunc，可以修改test.log输入的格式
func loggerModify(logger *log.Logger) {
	logger.SetPrefix("[INFO]")
	logger.SetFlags(log.LstdFlags | log.Llongfile)
	logger.Println("a logger.println message is here")
	err := logger.Output(2, "a logger.output message is here")
	if err != nil {
		log.Println("logger.output failed:", err)
	}
}

// 使用log.setOutput，可以修改test.log输入的格式，也可以输入数据到test.log文件
func logSetOutput(logFile io.Writer) {
	log.SetOutput(logFile)
	log.SetPrefix("[INFO]")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("a log.println message is here")
	err := log.Output(2, "a log.output message is here")
	if err != nil {
		log.Fatalln("log.output failed:", err)
	}
}
