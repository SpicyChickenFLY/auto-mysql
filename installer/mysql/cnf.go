package mysql

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"
	"github.com/SpicyChickenFLY/ini"
)

const (
	// dir
	baseDirKey    = "basedir"
	dataDirKey    = "datadir"
	tmpDirKey     = "tmpdir"
	undoLogDirKey = "innodb_undo_directory"
	// file
	sockFileKey     = "socket"
	pidFileKey      = "pid-file"
	errLogFileKey   = "log_error"
	binLogFileKey   = "log-bin"
	relayLogFileKey = "relay_log"
	slowLogFileKey  = "slow_query_log_file"
	multiLogFileKey = "log"
	// value
	portKey = "port"
)

const (
	rootPath = "/"

	stdCnfTemplateGeneral  = "./static/conf/template_gen.cnf"
	stdCnfTemplateInstance = "./static/conf/template_inst.cnf"
	srcCnfFileDef          = "./static/conf/my.cnf"

	stdSectionClient = "client"

	stdSectionMysqlDaemonMulti   = "mysqld_multi"
	stdSectionMysqlServerGeneral = "mysqld"

	stdBaseDir         = "/usr/local/mysql"
	stdDataDir         = "/mysqldata/mysql%d/data"
	stdErrorLogFileDir = "/mysqldata/mysql%d/log/mysqld.log"
	stdSockFileDir     = "/mysqldata/mysql%d/mysql.sock"

	templateSectionName     = "template"
	templatePortPlaceHolder = "[port]"
)

// CheckRquireDirExists make sure all dir exists
func CheckRquireDirExists(
	servInstInfo *ServerInstanceInfo, srcCnfFile string) error {
	// Initialize map dictionary
	cnfDirKey := []string{
		baseDirKey,
		dataDirKey,
		tmpDirKey,
		undoLogDirKey,
	}
	cnfFileKey := []string{
		sockFileKey,
		pidFileKey,
		errLogFileKey,
		binLogFileKey,
		relayLogFileKey,
		slowLogFileKey,
		multiLogFileKey,
	}
	dirList := ""

	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
	}, srcCnfFile)
	if err != nil {
		return err
	}
	secNames := cfg.SectionStrings()
	for _, secName := range secNames {
		section, err := cfg.GetSection(secName)
		if err != nil {
			return err
		}
		for _, keyName := range cnfDirKey {
			if section.Haskey(keyName) {
				dirPath := section.Key(keyName).Value()
				dirList += " " + dirPath
			}
			if keyName == dataDirKey {
				dirPath := rootPath
				dirShards := strings.Split(section.Key(keyName).Value(), "/")
				for _, dirShard := range dirShards {
					if !strings.Contains(dirShard, "mysql") {
						path.Join(dirPath, dirShard)
					} else {
						dirPath = path.Join(dirPath, dirShard)
						modifyDataDir(servInstInfo.ServerInfo, dirPath, sqlFileMode)
						break
					}
				}
			}
		}
		for _, keyName := range cnfFileKey {
			if section.Haskey(keyName) {
				filePath := section.Key(keyName).Value()
				index := strings.LastIndex(filePath, "/")
				if index > 0 {
					dirPath := string([]byte(filePath)[:index])
					dirList += " " + dirPath
				}
				// if keyName == multiLogFileKey {
				// 	_, err := createFile(s *linux.ServerInfo, filePath string)
				// }
			}
		}
	}

	return createDir(servInstInfo.ServerInfo, dirList)
}

// GenerateStdCnf generate Standard my.cnf(single/multi) for server
func GenerateStdCnf(servInstInfo *ServerInstanceInfo) (string, error) {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo, linux.Rm(srcCnfFileDef)); err != nil {
		return "", err
	}
	servInstInfo.BaseDir = stdBaseDir
	cfgGen, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
		// UnescapeValueDoublQueotes: true,
	}, stdCnfTemplateGeneral)
	if err != nil {
		return "", err
	}

	if len(servInstInfo.InstInfos) == 1 { // Single instance
		cfgGen.DeleteSection(stdSectionMysqlDaemonMulti)
	}

	for _, secName := range cfgGen.SectionStrings() {
		sec, err := cfgGen.GetSection(secName)
		if err != nil {
			return "", err
		}
		replaceAllValueInSection(
			sec,
			templatePortPlaceHolder,
			fmt.Sprint(servInstInfo.InstInfos[0].Port))
	}
	cfgGen.SaveTo(srcCnfFileDef)

	bytes, err := ioutil.ReadFile(stdCnfTemplateInstance)
	if err != nil {
		return "", err
	}
	fw, err := os.OpenFile(srcCnfFileDef, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer fw.Close()
	OriginalStr := string(bytes)
	for i := 0; i < len(servInstInfo.InstInfos); i++ {
		servInstInfo.InstInfos[i].DataDir = fmt.Sprintf(
			stdDataDir, servInstInfo.InstInfos[i].Port)
		servInstInfo.InstInfos[i].LogDir = fmt.Sprintf(
			stdErrorLogFileDir, servInstInfo.InstInfos[i].Port)
		servInstInfo.InstInfos[i].SockDir = fmt.Sprintf(
			stdSockFileDir, servInstInfo.InstInfos[i].Port)
		var newSecName string
		if len(servInstInfo.InstInfos) == 1 { // Single instance
			newSecName = fmt.Sprintf("[%s]",
				stdSectionMysqlServerGeneral)
		} else {
			newSecName = fmt.Sprintf("[%s%d]",
				stdSectionMysqlServerGeneral, servInstInfo.InstInfos[i].Port)
		}
		// fmt.Println(len(servInstInfo.InstInfos), newSecName)
		newInstStr := strings.ReplaceAll(OriginalStr, "[template]", newSecName)
		newInstStr = strings.ReplaceAll(
			newInstStr,
			templatePortPlaceHolder,
			fmt.Sprint(servInstInfo.InstInfos[i].Port))

		_, err := fw.WriteString(newInstStr)
		if err != nil {
			return "", err
		}
	}
	return srcCnfFileDef, nil
}

func replaceAllValueInSection(
	sec *ini.Section, placeHolder, replaceStr string) {
	for _, key := range sec.Keys() {

		key.SetValue(
			strings.Replace(
				key.Value(), placeHolder, replaceStr, -1))

	}
}
