package installer

import (
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/utils/linux"
	"github.com/SpicyChickenFLY/auto-mysql/utils/progress"
)

const (
	processName = "mysql"
)

// prepare for mySQL installation
//	Procedure:

//	1. Create group/user
// Kill exists MySQL process
//	2. Decompress the archive
//	3. Move configure file
//	4. Create data directory
func prepare(
	useSrcFile bool,
	servInstInfo *ServerInstanceInfo,
	srcSQLFile, srcCnfFile string) error {
	//
	if err := progress.Check("Install Dependencies",
		PrepareMysqlCentos(servInstInfo)); err != nil {
		return err
	}
	// Create group/user
	if err := progress.Check("Create group/user",
		linux.CreateUserWithGroup(servInstInfo.ServerInfo, userName, groupName)); err != nil {
		return err
	}
	// Kill exists MySQL process
	if err := progress.Check("Kill exists MySQL process",
		linux.KillProcess(servInstInfo.ServerInfo, processName)); err != nil {
		return err
	}
	// Move configure file
	if err := progress.Check("Move configure file",
		MoveCnfFile(servInstInfo.ServerInfo, srcCnfFile)); err != nil {
		return err
	}
	// Create data directory
	if err := progress.Check("Create data directory",
		CheckRquireDirExists(servInstInfo, srcCnfFile)); err != nil {
		return err
	}
	// Decompress the archive if srcSQLFile exists
	if useSrcFile {
		return progress.Check("Decompress the archive",
			ExtractSoftware(
				servInstInfo.ServerInfo,
				srcSQLFile,
				servInstInfo.BaseDir))
	}
	return nil
}

// InstallCustomInstance install MySQL single instance
// on  server
//	Procedure:
//	1. Prepare for environment for Installation
//	2. Initialize MySQL(without password)
func InstallCustomInstance(
	useSrcFile bool,
	servInstInfo *ServerInstanceInfo,
	srcSQLFile, srcCnfFile string) error {
	// Prepare for environment
	if err := prepare(
		useSrcFile, servInstInfo, srcSQLFile, srcCnfFile); err != nil {
		return err
	}
	// Initialize MySQL(without password)
	if err := progress.Check(
		"Initialize MySQL(without password)",
		InitInstance(servInstInfo)); err != nil {
		return err
	}
	// Installation compelete
	return progress.Check("Install Compelete", nil)
}

// InstallStandardMGRInstance is a func
//	Procedure:
//	1. Parse the server parameter
//	2. Generate Cnf for each server
//	2. Install instance on each server by parameter
//	3. Create Master/Slave relationship
func InstallStandardMGRInstance(
	useSrcFile bool, srcSQLFile, infoStr, mysqlPwd string) error {

	// Parse the server parameter
	allServInstInfos, err := ParseServerStr(infoStr)
	if progress.Check("Parse the server parameter", err) != nil {
		return err
	}

	for _, servInstInfo := range allServInstInfos {
		fmt.Printf("--- Server: %s:%d ---\n",
			servInstInfo.ServerInfo.Host, servInstInfo.ServerInfo.Port)
		// Generate Cnf for Server
		srcCnfFile, err := GeneratestdCnf(servInstInfo)
		if progress.Check("Generate Cnf for Server", err) != nil {
			return err
		}
		// Install instance on each server by parameter
		if err := InstallCustomInstance(
			useSrcFile, servInstInfo, srcSQLFile, srcCnfFile); err != nil {
			return err
		}
	}

	fmt.Println("--- Creating MGR Relationship ---")
	// Create MGR relationship
	if len(allServInstInfos) > 1 || len(allServInstInfos[0].InstInfos) > 1 {
		if err := progress.Check("Create MGR relationship",
			CreateMGRRelation(
				allServInstInfos, mysqlPwd)); err != nil {
			return err
		}
	}

	return nil
}

// InstallStandardReplicaIntance is a func
//	Procedure:
//	1. Parse the server parameter
//	2. Generate Cnf for each server
//	2. Install instance on each server by parameter
//	3. Create Master/Slave relationship
func InstallStandardReplicaIntance(
	useSrcFile bool, srcSQLFile, infoStr, mysqlPwd string) error {

	// Parse the server parameter
	allServInstInfos, err := ParseServerStr(infoStr)
	if progress.Check("Parse the server parameter", err) != nil {
		return err
	}

	for _, servInstInfo := range allServInstInfos {
		fmt.Printf("--- Server: %s:%d ---\n",
			servInstInfo.ServerInfo.Host, servInstInfo.ServerInfo.Port)
		// Generate Cnf for Server
		srcCnfFile, err := GeneratestdCnf(servInstInfo)
		if progress.Check("Generate Cnf for Server", err) != nil {
			return err
		}
		// Install instance on each server by parameter
		if err := InstallCustomInstance(
			useSrcFile, servInstInfo, srcSQLFile, srcCnfFile); err != nil {
			return err
		}
	}

	fmt.Println("--- Creating Master/Slave Relationship ---")
	// Create Master/Slave relationship
	if len(allServInstInfos) > 1 || len(allServInstInfos[0].InstInfos) > 1 {
		if err := progress.Check("Create Master/Slave relationship",
			CreateReplicaRelation(
				allServInstInfos, mysqlPwd)); err != nil {
			return err
		}
	}

	return nil
}
