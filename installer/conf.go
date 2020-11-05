package installer

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	// dir
	// BASE_DIR_KEY = "basedir"
	DATA_DIR_KEY = "datadir"
	TMP_DIR_KEY  = "tmpdir"
	UNDO_LOG_KEY = "innodb_undo_directory"
	// file
	SOCK_KEY      = "socket"
	PID_FILE_KEY  = "pid-file"
	ERR_LOG_KEY   = "log_error"
	BIN_LOG_KEY   = "log-bin"
	RELAY_LOG_KEY = "relay_log"
	SLOW_LOG_KEY  = "slow_query_log_file"
	// default value

)

// FIXME: for now it only support single instance conf
// FIXME: for now "basedir" & "datadir" must be written in conf
func checkCnfDir(srcCnfFile, userName, groupName string, fileMode uint32) error {
	// Initialize map dictionary
	confDirKV := map[string]string{
		DATA_DIR_KEY: "",
		TMP_DIR_KEY:  "",
		UNDO_LOG_KEY: "",
	}
	confFileKV := map[string]string{
		SOCK_KEY:      "",
		PID_FILE_KEY:  "",
		ERR_LOG_KEY:   "",
		BIN_LOG_KEY:   "",
		RELAY_LOG_KEY: "",
		SLOW_LOG_KEY:  "",
	}

	fr, err := os.Open(srcCnfFile)
	if err != nil {
		return err
	}
	defer fr.Close()

	br := bufio.NewReader(fr)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		lineStr := strings.Split(string(line), "#")[0] // remove comment in kv line
		lineStr = strings.TrimSpace(lineStr)           // remove prefix/posfix space
		// fmt.Println(lineStr)
		signIndex := strings.Index(lineStr, "=") // roughly separate kv
		if signIndex < 0 {
			continue // ignore blank line
		}
		key := strings.TrimSpace(lineStr[:signIndex]) // extract the key
		if len(key) == 0 {
			continue
		}
		// key matched dir path in configure
		if _, ok := confDirKV[key]; ok {
			cnfDir := strings.TrimSpace(lineStr[signIndex+1:])
			if len(cnfDir) != 0 {
				if err := createDirWithDetail(
					cnfDir, userName, groupName, fileMode); err != nil {
					return err
				}
			}
			continue
		}
		// key matched file path in configure
		if _, ok := confFileKV[key]; ok {
			cnfFile := strings.TrimSpace(lineStr[signIndex+1:])
			signIndex := strings.LastIndex(cnfFile, "/")
			if signIndex < 0 {
				return nil
			}
			cnfDir := cnfFile[:signIndex] // get prefix dir by split '/'
			if len(cnfDir) != 0 {
				if err := createDirWithDetail(
					cnfDir, userName, groupName, fileMode); err != nil {
					return err
				}
			}
			continue
		}
	}
}
