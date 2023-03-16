package file_writer

import (
	"errors"
	"os"
	"unicode/utf8"

	"github.com/terryhay/dolly/generator/proxyes/file_proxy"
	"github.com/terryhay/dolly/generator/proxyes/os_proxy"
)

// WriteFile - write file string by dir path
func WriteFile(osDecor os_proxy.Proxy, dirPath, fileString string) error {
	const (
		nameDir  = "dolly"
		nameFile = "dolly.go"
	)

	if err := osDecor.IsExist(dirPath); err != nil {
		return err
	}

	argParserDirPath := expandPath(dirPath, nameDir)
	if errExist := osDecor.IsExist(argParserDirPath); errExist != nil {
		if errCreate := createArgParserDir(osDecor, argParserDirPath); errCreate != nil {
			return errors.Join(errExist, errCreate)
		}
	}

	argParserFilePath := expandPath(argParserDirPath, nameFile)
	return write(osDecor, argParserFilePath, fileString)
}

func expandPath(path, name string) string {
	generatePathRunes := []rune(path)
	backRune := generatePathRunes[len(generatePathRunes)-1]

	slashRune, _ := utf8.DecodeRuneInString("/")
	backSlashRune, _ := utf8.DecodeRuneInString("\\")

	if backRune != slashRune && backRune != backSlashRune {
		path += "/"
	}
	return path + name
}

func createArgParserDir(osDecor os_proxy.Proxy, generatePath string) error {
	const permissionsToWrite os.FileMode = 0777
	return osDecor.MkdirAll(generatePath, permissionsToWrite)
}

func write(decOS os_proxy.Proxy, filePath, fileBody string) (err error) {
	var decFile file_proxy.Proxy
	decFile, err = decOS.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		if errClose := decFile.Close(); errClose != nil {
			if err != nil {
				err = errors.Join(err, errClose)
				return
			}

			err = errClose
			return
		}
	}()

	err = decFile.WriteString(fileBody)
	if err != nil {
		return err
	}

	return err
}
