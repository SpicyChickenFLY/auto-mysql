package mysql

import (
	"fmt"
	"time"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"
	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/progress"
)

const (
	mysqlInitUserPwd  = ""
	mysqlGenUserPwd   = "123456"
	allowRemoteAccess = `mysql -uroot -S %s -e "grant all privileges  on *.* to root@'%%' identified by '%s' WITH GRANT OPTION;flush privileges;"`
)

// CreateReplicaRelation create Master/Slaves for instances
//	Procedure
//	1. Start instance [Master]
//	2. Setup Master instance [Master]
//	3. Copy data from Master to Slave [Master/Slave]
//	4. Start instance [Master/Slave]
//	5. Setup Slave instances [Slave]
//	6. Test Replication
func CreateReplicaRelation(
	allServInstInfos []*ServerInstanceInfo,
	newPwd string) error {
	masterServInst, slaveServInsts := sperateMasterSlaveInstance(allServInstInfos)

	// Setup Master instance
	if err := progress.Check("Setup Master instance",
		setupMasterInst(masterServInst, newPwd)); err != nil {
		return err
	}
	// Copy data from masterServInst to slave
	if err := progress.Check("Copy data from masterServInst to slave",
		CopyDataDir(masterServInst, allServInstInfos)); err != nil {
		return err
	}
	// Start instance [Master/Slave]
	for _, servInstInfo := range allServInstInfos {
		if err := StartMultiInst(servInstInfo); err != nil {
			return progress.Check("Start instance [Master/Slaves]", err)
		}
	}
	progress.Check("Start instance [Master/Slaves]", nil)
	// Setup Slave instances
	if err := progress.Check("Setup Slave instances",
		setupSlaveInst(masterServInst, slaveServInsts, newPwd)); err != nil {
		return err
	}
	// Test Replication
	return progress.Check("Test Replication",
		testReplication(masterServInst, slaveServInsts, newPwd))
	// return nil
}

// testReplication test replica between Master/Slaves
func testReplication(
	masterServInstInfo *ServerInstanceInfo,
	slaveServInstInfos []*ServerInstanceInfo,
	newPwd string) error {
	// Connect to Instance [Master]
	db, err := CreateConn(
		masterServInstInfo.ServerInfo.Host,
		masterServInstInfo.InstInfos[0].Port,
		mysqlGenUserPwd)
	if err != nil {
		return err
	}
	defer db.Close()
	// Create test data
	if err := CreateTestData(db); err != nil {
		return err
	}

	time.Sleep(time.Duration(5) * time.Second)

	for _, slaveServInstInfo := range slaveServInstInfos {
		for _, slaveInstInfo := range slaveServInstInfo.InstInfos {
			// Connect to Instance [Slave]
			db, err := CreateConn(
				slaveServInstInfo.ServerInfo.Host,
				slaveInstInfo.Port,
				mysqlGenUserPwd)
			if err != nil {
				return err
			}
			defer db.Close()
			// Retrive test data
			if err := RetriveTestData(db); err != nil {
				return err
			}
		}
	}
	return nil
}

// setupMasterInst is a func to setup for Master
//	Procedure:
//	1. Connect to Instance [Master]
//	2. Grant admin user [Master]
//	3. Grant replica user [Master]
//	4. Create test data [Master]
//	5. Close Connection to instance [Master]
func setupMasterInst(
	masterServInst *ServerInstanceInfo, generalPwd string) error {
	// Start instance [Master]
	if err := progress.Check("Start instance [Master]",
		StartMultiInst(masterServInst)); err != nil {
		return err
	}
	// Allow TCP connection to server and Alter root password [Master]
	_, err := linux.ExecuteCommand(
		masterServInst.ServerInfo,
		fmt.Sprintf(allowRemoteAccess,
			masterServInst.InstInfos[0].SockDir,
			mysqlGenUserPwd))
	if err != nil {
		progress.Check("Allow TCP connection to server and Alter root password",
			err)
	}
	// Restart instance [Master]
	if err := progress.Check("Stop instance [Master]",
		StopMultiInst(masterServInst)); err != nil {
		return err
	}
	if err := progress.Check("Start instance [Master]",
		StartMultiInst(masterServInst)); err != nil {
		return err
	}
	// Connect to Instance [Master]
	db, err := CreateConn(
		masterServInst.ServerInfo.Host,
		masterServInst.InstInfos[0].Port,
		generalPwd)
	if err != nil {
		return err
	}
	if err := progress.Check("Connect to Instance [Master]", err); err != nil {
		return err
	}
	// Grant admin user [Master]
	if err := progress.Check("Grant admin user [Master]",
		CreateAdminUser(db, generalPwd)); err != nil {
		return err
	}
	// Grant replica user [Master]
	if err := progress.Check("Grant replica user [Master]",
		CreateReplicaUser(db, generalPwd)); err != nil {
		return err
	}
	// Create test data [Master]
	if err := progress.Check("Create test data [Master]",
		CreateTestEnv(db)); err != nil {
		return err
	}
	// Close Connection to instance [Master]
	return progress.Check("Close Connection to instance [Master]",
		db.Close())
}

// setupSlaveInst is a func to setup for Slave
// to create Master-Slave Relationship [Slave]
//	Procedure:
//	1. Connect to Slave instance
//	2. Change Master for slave instance
//	3. Start Slave Replication
//	4. Close Connection to instance
func setupSlaveInst(
	masterServInst *ServerInstanceInfo,
	slaveServInsts []*ServerInstanceInfo,
	generalPwd string) error {
	for _, slaveServInst := range slaveServInsts {
		for _, slaveInst := range slaveServInst.InstInfos {
			// Start instance [Slave]
			if err := progress.Check("Start instance [Slave]",
				StartMultiInst(slaveServInst)); err != nil {
				return err
			}
			// Allow TCP connection to server and Alter root password [Slave]
			_, err := linux.ExecuteCommand(
				slaveServInst.ServerInfo,
				fmt.Sprintf(allowRemoteAccess,
					slaveInst.SockDir,
					mysqlGenUserPwd))
			if err != nil {
				progress.Check("Allow TCP connection to server and Alter root password",
					err)
			}
			// Restart instance [Slave]
			if err := progress.Check("Stop instance [Slave]",
				StopMultiInst(slaveServInst)); err != nil {
				return err
			}
			if err := progress.Check("Start instance [Slave]",
				StartMultiInst(slaveServInst)); err != nil {
				return err
			}
			// Connect to Slave instance
			db, err := CreateConn(
				slaveServInst.ServerInfo.Host, slaveInst.Port, mysqlGenUserPwd)
			if err := progress.Check("Connect to Slave Instance", err); err != nil {
				return err
			}
			// Change Master for slave instance
			if err := progress.Check("Change Master for slave instance",
				ChangeMaster(
					db,
					masterServInst.ServerInfo.Host,
					masterServInst.InstInfos[0].Port,
					generalPwd)); err != nil {
				return err
			}
			// Start Slave Replication
			if err := progress.Check("Start Slave Replication",
				StartSlave(db)); err != nil {
				return err
			}
			// Close Connection to instance [Slave]
			if err := progress.Check("Close Connection to instance [Slave]",
				db.Close()); err != nil {
				return err
			}
		}
	}
	return nil
}

func sperateMasterSlaveInstance(
	all []*ServerInstanceInfo) (
	masterServInst *ServerInstanceInfo,
	slaveServInsts []*ServerInstanceInfo) {
	// fmt.Println(all[0].InstInfos)
	masterServInst = &ServerInstanceInfo{
		ServerInfo: all[0].ServerInfo,
		InstInfos: []InstanceInfo{
			all[0].InstInfos[0]},
		BaseDir:  all[0].BaseDir,
		HasMater: true}
	for _, servInstInfo := range all {
		slaveServInsts = append(slaveServInsts, servInstInfo)
	}
	if len(all[0].InstInfos) >= 2 {
		slaveServInsts[0] = &ServerInstanceInfo{
			ServerInfo: all[0].ServerInfo,
			InstInfos:  all[0].InstInfos[1:],
			BaseDir:    all[0].BaseDir,
			HasMater:   false}
	}
	return masterServInst, slaveServInsts
}
