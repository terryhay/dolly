package file_writer

import (
	"fmt"
	osDecorator2 "github.com/terryhay/dolly/internal/os_decorator"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"unicode/utf8"
)

const (
	argParserDirName  = "parser"
	argParserFileName = "parser.go"
)

// Write - write file string by dir path
func Write(osd osDecorator2.OSDecorator, dirPath, fileString string) (err *dollyerr.Error) {
	err = checkDirPath(osd, dirPath)
	if err != nil {
		return err
	}

	argParserDirPath := expandPath(dirPath, argParserDirName)
	if err = checkDirPath(osd, argParserDirPath); err != nil {
		err = createArgParserDir(osd, argParserDirPath)
		if err != nil {
			return err
		}
	}

	argParserFilePath := expandPath(argParserDirPath, argParserFileName)
	err = write(osd, argParserFilePath, fileString)

	return err
}

func checkDirPath(osd osDecorator2.OSDecorator, dirPath string) *dollyerr.Error {
	if _, err := osd.Stat(dirPath); osd.IsNotExist(err) {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorInvalidGeneratePath,
			fmt.Errorf("check path exist error: %v\n", err),
		)
	}
	return nil
}

func expandPath(path, name string) string {
	generatePathRunes := []rune(path)
	backRune := generatePathRunes[len(generatePathRunes)-1]

	slashRune, _ := utf8.DecodeRuneInString("/")
	backSlashRune, _ := utf8.DecodeRuneInString("\\")

	if backRune != slashRune || backRune != backSlashRune {
		path += "/"
	}
	return path + name
}

func createArgParserDir(osd osDecorator2.OSDecorator, generatePath string) *dollyerr.Error {
	if err := osd.MkdirAll(generatePath, 0777); err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateDirError,
			fmt.Errorf("create dir error: %v\n", err))
	}
	return nil
}

func write(osd osDecorator2.OSDecorator, filePath, fileBody string) (res *dollyerr.Error) {
	file, err := osd.Create(filePath)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateFileError,
			fmt.Errorf("create file error: %v\n", err))
	}

	defer func(file osDecorator2.FileDecorator) {
		err = file.Close()
		if err != nil {
			res = dollyerr.NewError(
				dollyerr.CodeGeneratorFileCloseError,
				fmt.Errorf("can't close the file: %v\n", err))
		}
	}(file)

	err = file.WriteString(fileBody)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorWriteFileError,
			fmt.Errorf("file write error: %v\n", err))
	}

	return res
}
