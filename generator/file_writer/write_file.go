package file_writer

import (
	"fmt"
	fld "github.com/terryhay/dolly/generator/file_decorator"
	osd "github.com/terryhay/dolly/generator/os_decorator"
	"github.com/terryhay/dolly/utils/dollyerr"
	"unicode/utf8"
)

// WriteFile - write file string by dir path
func WriteFile(osDecor osd.OSDecorator, dirPath, fileString string) *dollyerr.Error {
	const (
		nameDir  = "dolly"
		nameFile = "dolly.go"
	)

	if !osDecor.IsExist(dirPath) {
		return dollyerr.NewError(dollyerr.CodeGeneratorInvalidPath, fmt.Errorf("WriteFile: path does not exist: %s", dirPath))
	}

	argParserDirPath := expandPath(dirPath, nameDir)
	if !osDecor.IsExist(argParserDirPath) {
		err := createArgParserDir(osDecor, argParserDirPath)
		if err != nil {
			return dollyerr.Append(err, fmt.Errorf(""))
		}
	}

	argParserFilePath := expandPath(argParserDirPath, nameFile)
	err := write(osDecor, argParserFilePath, fileString)

	return err
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

func createArgParserDir(osDecor osd.OSDecorator, generatePath string) *dollyerr.Error {
	if err := osDecor.MkdirAll(generatePath, 0777); err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateDirError,
			fmt.Errorf("create dir error: %s\n", err.Error()))
	}
	return nil
}

func write(osDecor osd.OSDecorator, filePath, fileBody string) (res *dollyerr.Error) {
	file, err := osDecor.Create(filePath)
	if err != nil {
		return dollyerr.NewError(
			dollyerr.CodeGeneratorCreateFileError,
			fmt.Errorf("create file error: %s\n", err.Error()))
	}

	defer func(file fld.FileDecorator) {
		err = file.Close()
		if err != nil {
			res = dollyerr.Append(err, fmt.Errorf("write: can't close the file"))
		}
	}(file)

	err = file.WriteString(fileBody)
	if err != nil {
		return dollyerr.Append(err, fmt.Errorf("write: can't write a string"))
	}

	return dollyerr.NewErrorIfItIs(
		dollyerr.CodeGeneratorWriteFileError,
		"file write error",
		err.Error())
}
