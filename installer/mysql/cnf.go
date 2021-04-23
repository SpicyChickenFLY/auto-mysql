package mysql

import (
	"fmt"
	"io/ioutil"
	"os"
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
)

// Initialize map dictionary

var (
	cnfDirKey = []string{
		baseDirKey,
		dataDirKey,
		tmpDirKey,
		undoLogDirKey,
	}
	cnfFileKey = []string{
		sockFileKey,
		pidFileKey,
		errLogFileKey,
		binLogFileKey,
		relayLogFileKey,
		slowLogFileKey,
		multiLogFileKey,
	}
)

// CheckRquireDirExists make sure all dir exists
func CheckRquireDirExists(
	servInstInfo *ServerInstanceInfo, srcCnfFile string) error {

	dirList := []string{}
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
				dirList = append(dirList, dirPath)
			}
		}
		for _, keyName := range cnfFileKey {
			if section.HasKey(keyName) {
				filePath := section.Key(keyName).Value()
				dirIndex := strings.LastIndex(filePath, "/")
				dirPath := filePath[:dirIndex]
				dirList = append(dirList, dirPath)
			}
		}
	}
	dirListStr := strings.Join(dirList, " ")
	return createDirForMySQL(servInstInfo.ServerInfo, dirListStr)
}

// GenerateStdCnf generate Standard my.cnf(single/multi) for server
func GenerateStdCnf(servInstInfo *ServerInstanceInfo) (string, error) {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo, linux.Rm(StdSrcCnfFileDef)); err != nil {
		return "", err
	}
	servInstInfo.BaseDir = StdBaseDir
	cfgGen, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
		// UnescapeValueDoublQueotes: true,
	}, StdSrcCnfTemplateGeneral)
	if err != nil {
		return "", err
	}

	// Judge the count of instances on this server
	if len(servInstInfo.InstInfos) == 1 { // Single instance
		cfgGen.DeleteSection(tplSectionDaemonMulti)
	}
	if err := replaceAllValueInCfg(
		cfgGen,
		tplPlaceHolderPort,
		fmt.Sprint(servInstInfo.InstInfos[0].Port)); err != nil {
		return "", err
	}

	cfgGen.SaveTo(StdSrcCnfFileDef)

	bytes, err := ioutil.ReadFile(StdSrcCnfTemplateInstance)
	if err != nil {
		return "", err
	}
	OriginalStr := string(bytes)

	for i := 0; i < len(servInstInfo.InstInfos); i++ {
		servInstInfo.InstInfos[i].DataDir = fmt.Sprintf(
			StdDataDir, servInstInfo.InstInfos[i].Port)
		servInstInfo.InstInfos[i].LogDir = fmt.Sprintf(
			StdErrorLogFileDir, servInstInfo.InstInfos[i].Port)
		servInstInfo.InstInfos[i].SockDir = fmt.Sprintf(
			StdSockFileDir, servInstInfo.InstInfos[i].Port)
		var newSecName string
		if len(servInstInfo.InstInfos) == 1 { // Single instance
			newSecName = fmt.Sprintf("[%s]",
				tplSectionDaemonSingle)
		} else {
			newSecName = fmt.Sprintf("[%s%d]",
				tplSectionDaemonSingle, servInstInfo.InstInfos[i].Port)
		}
		// fmt.Println(len(servInstInfo.InstInfos), newSecName)
		newInstStr := strings.ReplaceAll(
			OriginalStr, tplPlaceHolderInstMulti, newSecName)
		newInstStr = strings.ReplaceAll(
			newInstStr, tplPlaceHolderPort, fmt.Sprint(servInstInfo.InstInfos[i].Port))

		fw, err := os.OpenFile(StdSrcCnfFileDef, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return "", err
		}
		defer fw.Close()

		if _, err := fw.WriteString(newInstStr); err != nil {
			return "", err
		}
	}
	return StdSrcCnfFileDef, nil
}

func replaceAllValueInCfg(
	cfgGen *ini.File, placeHolder, replaceStr string) error {
	for _, secName := range cfgGen.SectionStrings() {
		sec, err := cfgGen.GetSection(secName)
		if err != nil {
			return err
		}
		for _, key := range sec.Keys() {
			key.SetValue(strings.Replace(key.Value(), placeHolder, replaceStr, -1))
		}
	}
	return nil
}
