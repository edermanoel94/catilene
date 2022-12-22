package tranfer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Transfer struct {
	ParentFolder string
}

var parentFolder string

func Run() {

}

func batchTransfer(extensions map[string][]string) {
	for _, filepaths := range extensions {
		for _, f := range filepaths {
			transfer(f)
		}
	}
}

func extensionCleaned(filename string) string {
	return strings.ToLower(strings.TrimLeft(filepath.Ext(filename), "."))
}

func createDir(dirName string) error {
	return os.MkdirAll(fmt.Sprintf("%s%s%s", parentFolder, string(os.PathSeparator), dirName), 0755)
}

func copyFile(dirTarget string, src *os.File) error {

	targetFolder := fmt.Sprintf("%s%s%s",
		parentFolder,
		string(os.PathSeparator),
		dirTarget,
	)

	fileInfo, err := src.Stat()

	if err != nil {
		return err
	}

	dst, err := os.Create(targetFolder + string(os.PathSeparator) + fileInfo.Name())

	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return err
	}

	return nil
}

func transfer(filepathSrc string) error {

	fileSrc, err := os.Open(filepathSrc)

	if err != nil {
		return err
	}

	defer fileSrc.Close()

	var dirTarget string

	switch extensionCleaned(filepathSrc) {

	case "jpg", "jpeg", "png", "aae":
		dirTarget = "photos"
	case "mp3":
		dirTarget = "music"
	case "mp4", "mov":
		dirTarget = "video"
	case "pdf", "docx", "doc", "rtf":
		dirTarget = "document"
	case "xls", "xlsm", "xlsx":
		dirTarget = "excel"
	case "ppt", "pptx":
		dirTarget = "presentation"
	case "zip", "rar":
		dirTarget = "compression"
	default:
		dirTarget = "others"
	}

	if err := createDir(dirTarget); err != nil {
		return err
	}

	return copyFile(dirTarget, fileSrc)
}
