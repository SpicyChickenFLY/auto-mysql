package main

import (
	"flag"
	"fmt"

	"github.com/romberli/log"
	"spicychicken.top/auto-mysql/installer"
)

const (
	LOG_FILE_NAME = "/home/chow/run.log"
	// LOG_LEVEL         = "info"
	// LOG_FORMAT        = "TEXT"
	// LOG_FILE_MAX_SIZE = 100 // unit:MB
	// LOG_EXPIRED_DAY   = 7
	// LOG_MAX_BACKUPS   = 5

	SRC_SQL_FILE = "./src/mysql/mysql.tar.gz"
	DST_SQL_PATH = "/usr/local/mysql"
	SRC_CNF_FILE = "./src/conf/my.cnf"
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
	fmt.Println("============================")
	fmt.Println("MySQL Automatic Installation")
	fmt.Print("============================\n\n")

	// Custom parameters
	mode := flag.String("m", "install", "install/remove/test")
	srcSqlFile := flag.String(
		"s", SRC_SQL_FILE, "postion of mysql-binary file")
	dstSqlPath := flag.String(
		"d", DST_SQL_PATH, "position for installation")
	srcCnfFile := flag.String(
		"c", SRC_CNF_FILE, "postion of you configure file")
	startPos := flag.Int(
		"p", 0, "startPhase")

	// Fixed parameters
	dstCnfFile := DST_CNF_FILE
	flag.Parse()

	log.Info("Custom parameters:")
	log.Info(fmt.Sprintf("srcSqlFile: %s", *srcSqlFile))
	log.Info(fmt.Sprintf("dstSqlPath: %s", *dstSqlPath))
	log.Info(fmt.Sprintf("srcCnfFile: %s", *srcCnfFile))
	log.Info(fmt.Sprintf("startPos: %d", *startPos))

	fmt.Println("Please check your input parameter:")
	fmt.Printf("srcSqlFile: %s\n", *srcSqlFile)
	fmt.Printf("dstSqlPath: %s\n", *dstSqlPath)
	fmt.Printf("srcCnfFile: %s\n", *srcCnfFile)
	fmt.Printf("startPos: %d\n", *startPos)

	// Analyze the running mode
	if *mode == "test" {
		log.Info("MySQL automatic installation test started")
		if err := installer.TestInstall(
			*startPos, *srcSqlFile, *dstSqlPath,
			*srcCnfFile, dstCnfFile); err != nil {
			fmt.Println(err)
			log.Info("MySQL automatic installation test failed")
		}
		log.Info("MySQL automatic installation test succeed")
	} else if *mode == "install" {
		log.Info("MySQL automatic installation started")
		if err := installer.Install(
			*srcSqlFile, *dstSqlPath,
			*srcCnfFile, dstCnfFile); err != nil {
			fmt.Println(err)
			log.Info("MySQL automatic installation failed")
			return
		}
		log.Info("MySQL automatic installation succeed")
	} else {
		installer.Remove()
	}

}
