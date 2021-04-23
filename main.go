package main

import (
	"flag"
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer"
	"github.com/SpicyChickenFLY/auto-mysql/installer/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/romberli/log"
)

const (
	logFileName = "/tmp/run.log"
	// LOG_LEVEL         = "info"
	// LOG_FORMAT        = "TEXT"
	// LOG_FILE_MAX_SIZE = 100 // unit:MB
	// LOG_EXPIRED_DAY   = 7
	// LOG_MAX_BACKUPS   = 5

	srcSQLFileDef   = "./static/mysql/mysql.tar.gz"
	dstSQLPathDef   = "/home/chow/Softs/mysql"
	srcCnfFileDef   = mysql.StdSrcCnfFileDef
	servInstInfoDef = "root:123@localhost:3306|3307"
	mysqlPwdDef     = "123456"
)

func main() {
	// 初始化全局变量
	_, _, err := log.InitLoggerWithDefaultConfig(logFileName)
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
	runMode := flag.String("m", "mgr-s", "single/multi/standard/remove/test")
	useSrcFile := flag.Bool("f", false, "true if need copy file")
	srcSQLFile := flag.String(
		"s", srcSQLFileDef, "postion of mysql-binary file")
	dstSQLPath := flag.String(
		"d", dstSQLPathDef, "position for installation")
	srcCnfFile := flag.String(
		"c", mysql.StdSrcCnfFileDef, "postion of you configure file")
	servInstInfo := flag.String(
		"i", servInstInfoDef,
		"information of instance - userName:userPwd@host:port#port1|port2|port3;userName:...")
	mysqlPwd := flag.String(
		"p", mysqlPwdDef, "password for mysql user root")

	flag.Parse()

	log.Info("Custom parameters:")
	log.Info(fmt.Sprintf("useSrcFile: %t", *useSrcFile))
	log.Info(fmt.Sprintf("srcSQLFile: %s", *srcSQLFile))
	log.Info(fmt.Sprintf("dstSQLPath: %s", *dstSQLPath))
	log.Info(fmt.Sprintf("srcCnfFile: %s", *srcCnfFile))
	log.Info(fmt.Sprintf("mysqlPwd: %s", *mysqlPwd))
	log.Info(fmt.Sprintf("RunMode: %s", *runMode))

	fmt.Println("Please check your input parameter:")
	fmt.Printf("useSrcFile: %t\n", *useSrcFile)
	fmt.Printf("srcSQLFile: %s\n", *srcSQLFile)
	fmt.Printf("dstSQLPath: %s\n", *dstSQLPath)
	fmt.Printf("srcCnfFile: %s\n", *srcCnfFile)
	fmt.Printf("mysqlPwd: %s\n\n", *mysqlPwd)
	fmt.Printf("RunMode: %s\n\n", *runMode)

	// Analyze the installMode
	switch *runMode {
	// case "single":
	// 	installer.InstallCustomSingleInstance(
	// 		*srcSQLFile, *dstSQLPath, *srcCnfFile, *mysqlPwd)
	// case "multi":
	// 	installer.InstallCustomMultiInstance(
	// 		*srcSQLFile, *dstSQLPath, *srcCnfFile, *mysqlPwd)
	// case "standard":
	// 	installer.InstallStandardMultiInstanceOnMultiServer(*useSrcFile, *srcSQLFile, *servInstInfo, *mysqlPwd)
	case "mgr-s":
		installer.InstallStandardMGRInstance(*useSrcFile, *srcSQLFile, *servInstInfo, *mysqlPwd)
	}
	fmt.Print("============================\n\n")
}
