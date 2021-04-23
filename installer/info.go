package installer

import (
	"strings"

	"github.com/SpicyChickenFLY/auto-mysql/utils/linux"
	"github.com/lingdor/stackerror"
)

const (
	userName    = "mysql"
	groupName   = "mysql"
	sqlFileMode = 775
	cnfFileMode = 644
)

// ServerInstanceInfo store
type ServerInstanceInfo struct {
	ServerInfo *linux.ServerInfo
	InstInfos  []InstanceInfo
	BaseDir    string
	HasMater   bool
}

// InstanceInfo store the information of MySQL instance
type InstanceInfo struct {
	Port    int
	DataDir string
	LogDir  string
	SockDir string
	// isMaster bool
}

// NewServInstInfo is a construction func for ServerInstanceInfo
// infomation format: `userName(root):userPwd@host:port@port_1|port_2|port_3`
func NewServInstInfo(infoStr string) (servInstInfo *ServerInstanceInfo, err error) {
	servInstInfo = &ServerInstanceInfo{
		InstInfos: []InstanceInfo{},
	}

	infoStr = strings.Replace(infoStr, " ", "", -1)
	// divide into userInfo and addrInfo by #
	infoParts := strings.Split(infoStr, "#")
	// userName(root):userPwd@host:port$port_1|port_2|port_3
	switch len(infoParts) {
	case 2:
		instStr := infoParts[1]
		// divide instance ports by |
		portParts := strings.Split(instStr, "|")
		if len(portParts) == 0 { // use default 3306
			servInstInfo.InstInfos = []InstanceInfo{{Port: 3306}}
		}
		for _, portStr := range portParts {
			port, err := linux.CheckPortValid(portStr, 1024, 65535)
			if err != nil {
				return nil, stackerror.New("port must be a integer[1024-65535]")
			}
			servInstInfo.InstInfos = append(servInstInfo.InstInfos, InstanceInfo{Port: port})
		}
	case 1:
		servInstInfo.InstInfos = []InstanceInfo{{Port: 3306}}
	default:
		return nil, stackerror.New("invalid ServerInstanceInfomation format")
	}
	servStr := infoParts[0]
	servInfo, err := linux.NewServInfo(servStr)
	if err != nil {
		return nil, err
	}
	servInstInfo.ServerInfo = servInfo
	return servInstInfo, nil
}

func (s *ServerInstanceInfo) findInstByPort(port int) (*InstanceInfo, error) {
	for i, InstInfo := range s.InstInfos {
		if InstInfo.Port == port {
			return &s.InstInfos[i], nil
		}
	}
	return nil, stackerror.New("no matched Port found in InstInfos")
}

// ParseServerStr parse the input param of address string
// and create a relationship of Master/Slaves for these Instances
// sign ";" must be insert between information
// infomation format: `userName(root):userPwd@host:port@port_1|port_2|port_3`
func ParseServerStr(
	infoStr string) (
	allServInstInfos []*ServerInstanceInfo,
	err error) {

	// preprocess the string
	infoStr = strings.Replace(infoStr, " ", "", -1)
	infoStr = strings.Replace(infoStr, "\n", "", -1)
	infoStr = strings.Replace(infoStr, "\t", "", -1)
	infoStrs := strings.Split(infoStr, ";")

	if len(infoStrs) == 0 { //empty string
		return allServInstInfos, stackerror.New("")
	}

	for _, infoStr := range infoStrs {
		serverInfo, err := NewServInstInfo(infoStr)
		if err != nil {
			return allServInstInfos, err
		}
		allServInstInfos = append(allServInstInfos, serverInfo)
	}
	return allServInstInfos, nil
}

// KillMysqlProcess kill all process of mysql
func KillMysqlProcess(
	servInstInfo *ServerInstanceInfo) error {
	_, err := linux.ExecuteCommand(
		servInstInfo.ServerInfo,
		"ps -ef | grep mysql | grep -v grep | cut -c 9-15 | xargs kill -s 9")
	return err
}
