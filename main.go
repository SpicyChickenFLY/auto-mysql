package main

import (
	"flag"
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer"
	"github.com/romberli/log"
)

const (
	LOG_FILE_NAME = "/tmp/run.log"
	// LOG_LEVEL         = "info"
	// LOG_FORMAT        = "TEXT"
	// LOG_FILE_MAX_SIZE = 100 // unit:MB
	// LOG_EXPIRED_DAY   = 7
	// LOG_MAX_BACKUPS   = 5

	SRC_SQL_FILE = "./static/mysql/mysql.tar.gz"
	DST_SQL_PATH = "/home/chow/Softs/mysql"
	SRC_CNF_FILE = "./static/conf/my.cnf"
	DST_CNF_FILE = "/etc/my.cnf"
)

func main() {
	// 初始化全局变量
	_, _, err := log.InitLoggerWithDefaultConfig(LOG_FILE_NAME)
	if err != nil {
		fmt.Printf("Init logger failed: %s\n", err.Error())
		panic(err)
	}
	fmt.Println("Init logger succeed")

	log.Info("=============================")
	log.Info("Program Started")
	fmt.Println("\n============================")
	fmt.Println("MySQL Automatic Installation")
	fmt.Print("============================\n\n")

	// Custom parameters
	runMode := flag.String("m", "multi", "single/multi/remove/test")
	srcSQLFile := flag.String(
		"s", SRC_SQL_FILE, "postion of mysql-binary file")
	dstSQLPath := flag.String(
		"d", DST_SQL_PATH, "position for installation")
	srcCnfFile := flag.String(
		"c", SRC_CNF_FILE, "postion of you configure file")

	// Fixed parameters
	dstCnfFile := DST_CNF_FILE
	flag.Parse()

	log.Info("Custom parameters:")
	log.Info(fmt.Sprintf("srcSQLFile: %s", *srcSQLFile))
	log.Info(fmt.Sprintf("dstSQLPath: %s", *dstSQLPath))
	log.Info(fmt.Sprintf("srcCnfFile: %s", *srcCnfFile))
	log.Info(fmt.Sprintf("RunMode: %s", *runMode))

	fmt.Println("Please check your input parameter:")
	fmt.Printf("srcSQLFile: %s\n", *srcSQLFile)
	fmt.Printf("dstSQLPath: %s\n", *dstSQLPath)
	fmt.Printf("srcCnfFile: %s\n\n", *srcCnfFile)
	fmt.Printf("RunMode: %s\n\n", *runMode)

	// Analyze the installMode
	switch *runMode {
	case "single":
		installer.InstallSingleInstance(
			*srcSQLFile, *dstSQLPath,
			*srcCnfFile, dstCnfFile)
	case "multi":
		installer.InstallMultiInstance(
			*srcSQLFile, *dstSQLPath,
			*srcCnfFile, dstCnfFile)
	case "remove":
		installer.Remove()
	case "test":
		// TestCreateConnBySock()
	}
	fmt.Print("============================\n\n")
}
