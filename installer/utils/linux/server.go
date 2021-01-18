package linux

import (
	"strconv"
	"strings"

	"github.com/lingdor/stackerror"
)

const (
	userNameDef = "root"
	userPwdDef  = "123"
	hostDef     = "localhost"
	portDef     = 22
)

// LocalHost is a local server dont need to login
var LocalHost = &ServerInfo{
	Host: "localhost",
}

// ServerInfo store
type ServerInfo struct {
	UserName string
	UserPwd  string
	Host     string
	Port     int
}

// NewServInfo is a construction func for ServerInstanceInfo
// infomation format: `userName(root):userPwd@host:port`
func NewServInfo(infoStr string) (*ServerInfo, error) {
	servInfo := &ServerInfo{
		UserName: userNameDef,
		UserPwd:  userPwdDef,
		Host:     hostDef,
		Port:     portDef,
	}
	infoStr = strings.Replace(infoStr, " ", "", -1)
	// divide into userInfo and addrInfo by @
	infoParts := strings.Split(infoStr, "@")
	switch len(infoParts) {
	case 2:
		addrStr := infoParts[1]
		// divide addrInfo into host and ports
		addrParts := strings.Split(addrStr, ":")
		if len(addrParts) != 2 {
			return nil, stackerror.New("invalid host format")
		}
		if err := CheckHostValid(addrParts[0]); err != nil {
			return nil, err
		}
		port, err := CheckPortValid(addrParts[1], 0, 65535)
		if err != nil {
			return nil, err
		}
		servInfo.Host = addrParts[0]
		servInfo.Port = port
		fallthrough
	case 1:
		userStr := infoParts[0]
		// divide userInfo into userName and userPwd
		userParts := strings.Split(userStr, ":")
		if len(userParts) != 2 {
			return nil, stackerror.New("invalid user format")
		}
		if userParts[0] != "root" {
			return nil, stackerror.New("user must be root")
		}
		if userParts[1] == "" {
			return nil, stackerror.New("password cant be empty")
		}
		servInfo.UserName = userParts[0]
		servInfo.UserPwd = userParts[1]
	default:
		return nil, stackerror.New("invalid instance format")
	}
	return servInfo, nil
}

func checkIPv4(IP string) error {
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return stackerror.New("invalid IP format")
	}
	for _, s := range strs {
		number, err := strconv.Atoi(s)
		if err != nil {
			return stackerror.New("invalid IP format")
		}
		if number < 0 || number > 255 {
			return stackerror.New("invalid IP format")
		}
	}
	return nil
}

// CheckHostValid check if host is IP or localhost
// FIXME: Hosts who are not IP or localhost is not in consideration now
func CheckHostValid(host string) error {
	if host == "localhost" {
		return nil
	}
	return checkIPv4(host)
}

// CheckPortValid check if port is a integer and range in 1024-65535
func CheckPortValid(portStr string, min, max int) (int, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return portDef, stackerror.New("port is not a integer")
	}
	if port < min || port > max {
		return portDef, stackerror.New("port must be a integer[1024-65535]")
	}
	return port, err
}
