package installer

import "github.com/SpicyChickenFLY/auto-mysql/utils/progress"

// CreateMGRRelation create Master/Slaves for instances
//	Procedure
//	1. Start instance [Master]
//	2. Setup Master instance [Master]
//	3. Copy data from Master to Slave [Master/Slave]
//	4. Start instance [Master/Slave]
//	5. Setup Slave instances [Slave]
//	6. Test MGRtion
func CreateMGRRelation(
	allServInstInfos []*ServerInstanceInfo,
	newPwd string) error {
	masterServInst, slaveServInsts := sperateMasterSlaveInstance(allServInstInfos)

	// Setup Master instance
	if err := progress.Check("Setup Master instance",
		setupMasterInst(masterServInst, newPwd)); err != nil {
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
	// Test MGRtion
	// return progress.Check("Test MGRtion",
	// 	testMGRtion(masterServInst, slaveServInsts, newPwd))
	return nil
}
