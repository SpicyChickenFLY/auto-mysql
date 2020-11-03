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

	"github.com/romberli/log"
)

const (
	DEFAULT_FILE_MODE = 755
)

// TODO: complete this function to make module_test feasible
func compareFile() bool {
	return true
}

func moveFile(srcFile, dstFile string) error {
	if err := os.Rename(srcFile, dstFile); err != nil {
		log.Error("Move file encount error.")
		return err
	}
	return nil
}

func moveFileShell(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	cmdMv := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo mv %s %s", srcFile, dstFile))
	cmdMv.Stdout = &out
	cmdMv.Stderr = &stderr
	if err := cmdMv.Run(); err != nil {
		log.Warnf("cmdMv:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdMv: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func copyFileShell(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	cmdCp := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo cp -r %s %s", srcFile, dstFile))
	cmdCp.Stdout = &out
	cmdCp.Stderr = &stderr
	if err := cmdCp.Run(); err != nil {
		log.Warnf("cmdCp:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdCp: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func modifyOwner(dirPath, userName, groupName string) error {
	// sudo chown -R userName:groupName dirPath
	cmdChown := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo chown -R %s:%s %s", userName, groupName, dirPath))
	cmdChown.Stdout = &out
	cmdChown.Stderr = &stderr
	if err := cmdChown.Run(); err != nil {
		log.Warnf("cmdChown:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdChown: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func modifyMode(dirPath string, fileMode uint32) error {
	// sudo chmod -R 755 dirPath
	cmdChmod := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))
	cmdChmod.Stdout = &out
	cmdChmod.Stderr = &stderr
	if err := cmdChmod.Run(); err != nil {
		log.Warnf("cmdChmod:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdChmod: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func modifyDir(dirPath, userName, groupName string, fileMode uint32) error {
	if err := modifyOwner(dirPath, userName, groupName); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := modifyMode(dirPath, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

//
func createDir(dirPath string) error {
	// create dir for files, FileMode-0755 to make sure it can be accessed
	if err := os.MkdirAll(dirPath, DEFAULT_FILE_MODE); err != nil {
		log.Errorf("encount error:%s", err)
		return err
	}
	return nil
}

func createDirShell(dirPath string) error {
	// mkdir
	cmdMkdir := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo mkdir -p %s", dirPath))
	cmdMkdir.Stdout = &out
	cmdMkdir.Stderr = &stderr
	if err := cmdMkdir.Run(); err != nil {
		log.Warnf("cmdMkdir:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdMkdir: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}

func createDirWithDetail(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := createDir(dirPath); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := modifyDir(
		dirPath, userName, groupName, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

func createDirWithDetailShell(
	dirPath, userName, groupName string, fileMode uint32) error {
	if err := createDirShell(dirPath); err != nil {
		log.Error("encount error.")
		return err
	}
	if err := modifyDir(
		dirPath, userName, groupName, fileMode); err != nil {
		log.Error("encount error.")
		return err
	}
	return nil
}

// unTarWithGzip extract a .tar.gz file by name write in GOLANG
//  == Extract Procedure ==
//	1. open srcFile
//  2. traverse srcFile
//  3. if a reg file: create directory and empty dstFile
//     else: just create directory
//  4. write to dstFile
//  5. close handle instance
func unTarWithGzipGo(srcFile string, dstPath string) error {
	// Open fileReader for .tar.gz file
	fr, err := os.Open(srcFile)
	if err != nil {
		log.Error("encount error.")
		return err
	}
	defer fr.Close()
	// Open gzipReader with fileReader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Error("encount error.")
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
			log.Error("encount error.")
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
				fmt.Println("Untar encount error:")
				fmt.Println(err)
				log.Error("encount error.")
				return err
			}
			defer fw.Close()
			byteNum, err := io.Copy(fw, tr)
			if err != nil {
				log.Error("encount error.")
				return err
			}
			fmt.Printf("/%dByte\n", byteNum)
			fw.Close()
		default:
			log.Error("encount error.")
			return errors.New("Unknown File Type")
		}
	}
}

// unTarWithGzip extract a .tar.gz file by name write in SHELL
func unTarWithGzipShell(srcFile string, dstPath string) error {
	if err := createDirShell(dstPath); err != nil {
		return err
	}
	// tar -zxvf srcFile dstPath
	cmdTar := exec.Command(
		"/bin/sh", "-c",
		fmt.Sprintf("sudo tar -zxvf %s -C %s ", srcFile, dstPath))
	cmdTar.Stdout = &out
	cmdTar.Stderr = &stderr
	if err := cmdTar.Run(); err != nil {
		log.Warnf("cmdTar:%s:%s\n",
			err, stderr.String())
		fmt.Printf(
			"cmdTar: %s:%s\n",
			err, stderr.String())
		return err
	}
	return nil
}
