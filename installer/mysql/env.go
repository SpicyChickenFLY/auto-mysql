package mysql

import "github.com/SpicyChickenFLY/auto-mysql/installer/utils/linux"

// PrepareMysqlUbuntu is a func to install all dependencies for
// MySQL in Ubuntu(Linux OS)
// dependencies list: libaio, libncurese
func PrepareMysqlUbuntu(servInstInfo *ServerInstanceInfo) error {
	if _, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		"apt-cache search libaio"); err != nil {
		return err
	}
	_, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		"apt install libaio1 libncurses5")
	return err
}

// PrepareMysqlCentos is a func to install all dependencies for
// MySQL in Centos(Linux OS)
// dependencies list: libaio, libnuma, perl, autoconf
func PrepareMysqlCentos(servInstInfo *ServerInstanceInfo) error {
	_, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		"yum install -y libaio libaio-devel numactl autoconf perl perl-devel perl-Data-Dumper")
	return err
}
