package installer

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// TODO: complete this function to make module_test feasible
func compareFile() bool {
	return true
}

func moveFile(srcFile, dstFile string) error {
	if err := os.Rename(srcFile, dstFile); err != nil {
		return err
	}
	return nil
}

func modifyDirOnOwner(dirPath, userName, groupName string) error {
	// chown -R userName:groupName dirPath
	cmdChown := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo chown -R %s:%s %s", userName, groupName, dirPath))
	if err := cmdChown.Run(); err != nil {
		fmt.Println(fmt.Sprintf("sudo chown -R %s:%s %s", userName, groupName, dirPath))
		fmt.Println(err)
		return err
	}
	return nil
}

func modifyDirOnMode(dirPath string, fileMode uint32) error {
	// chmod -R 755 dirPath
	cmdChmod := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))
	if err := cmdChmod.Run(); err != nil {
		fmt.Println(fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))
		fmt.Println(err)
		return err
	}
	return nil
}

func modifyDir(dirPath, userName, groupName string, fileMode uint32) error {
	if err := modifyDirOnOwner(dirPath, userName, groupName); err != nil {
		return err
	}
	if err := modifyDirOnMode(dirPath, fileMode); err != nil {
		return err
	}
	return nil
}

//
func createDir(dirPath string) error {
	// create dir for files, FileMode-0755 to make sure it can be accessed
	if err := os.MkdirAll(dirPath, FILE_MODE); err != nil {
		return err
	}
	return nil
}

func createDirWithDetail(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := createDir(dirPath); err != nil {
		return err
	}
	if err := modifyDir(
		dirPath, userName, groupName, fileMode); err != nil {
		return err
	}
	return nil
}

/*
 * == Extract Procedure ==
 * 1. open srcFile
 * 2. traverse srcFile
 * 3. if a reg file: create directory and empty dstFile
 *    else: just create directory
 * 4. write to dstFile
 * 5. close handle instance
 */

// unTarWithGzip extract a .tar.gz file by name write in GOLANG
func unTarWithGzipGo(srcFile string, dstPath string) error {
	// Open fileReader for .tar.gz file
	fr, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer fr.Close()
	// Open gzipReader with fileReader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// Open tarReader with gzipReader
	tr := tar.NewReader(gr)

	// traverse the files in tar archive
	for {
		// try to get next readable file
		header, err := tr.Next()
		switch err {
		case nil: // get a file or dir
		case io.EOF: // encount EOF, exit normally
			return nil
		default:
			return err
		}

		fmt.Printf("%s\t\t%d", header.Name, header.Size)
		switch header.Typeflag {
		case tar.TypeDir:
			createDir(
				dstPath + "/" + string(
					[]rune(header.Name)[:strings.LastIndex(header.Name, "/")]))
			fmt.Println("")
		case tar.TypeReg:
			fw, err := os.Create(dstPath + "/" + header.Name)
			if err != nil {
				return err
			}
			defer fw.Close()
			byteNum, err := io.Copy(fw, tr)
			if err != nil {
				return err
			}
			fmt.Printf("/%dByte\n", byteNum)
			fw.Close()
		default:
			return errors.New("Unknown File Type")
		}
	}
}

// unTarWithGzip extract a .tar.gz file by name write in SHELL
func unTarWithGzipShell(srcFile string, dstPath string) error {
	return nil
}
