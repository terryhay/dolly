package file_writer

import (
	"fmt"
	fld "github.com/terryhay/dolly/internal/file_decorator"
	osd "github.com/terryhay/dolly/internal/os_decorator"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"unicode/utf8"
)

// Write - write file string by dir path
func Write(osDecor osd.OSDecorator, dirPath, fileString string) (err *dollyerr.Error) {
	const (
		nameDir  = "dolly"
		nameFile = "dolly.go"
	)

	err = checkDirPath(osDecor, dirPath)
	if err != nil {
		return err
	}

	argParserDirPath := expandPath(dirPath, nameDir)
	if err = checkDirPath(osDecor, argParserDirPath); err != nil {
		err = createArgParserDir(osDecor, argParserDirPath)
		if err != nil {
			return err
		}
	}

	argParserFilePath := expandPath(argParserDirPath, nameFile)
	err = write(osDecor, argParserFilePath, fileString)

	return err
}

func checkDirPath(osd osd.OSDecorator, dirPath string) *dollyerr.Error {
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

func createArgParserDir(osDecor osd.OSDecorator, generatePath string) *dollyerr.Error {
	if err := osDecor.MkdirAll(generatePath, 0777); err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateDirError,
			fmt.Errorf("create dir error: %v\n", err))
	}
	return nil
}

func write(osDecor osd.OSDecorator, filePath, fileBody string) (res *dollyerr.Error) {
	file, err := osDecor.Create(filePath)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateFileError,
			fmt.Errorf("create file error: %v\n", err))
	}

	defer func(file fld.FileDecorator) {
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
