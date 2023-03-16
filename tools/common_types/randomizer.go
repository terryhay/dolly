package common_types

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/dolly/tools/size"
)

type typeRandStorage struct {
	errors []error

	appName                NameApp
	infoChapterName        InfoChapterNAME
	infoChapterDescription []InfoChapterDESCRIPTION

	placeholderIDs []IDPlaceholder

	commandNameTooShort NameCommand
	commandNameTooLong  NameCommand

	commandNamesShort []NameCommand
	commandNamesLong  []NameCommand
	commandNames      []NameCommand

	placeholderNames   []NamePlaceholder
	flagOneLetterNames []NameFlag
	flagShortNames     []NameFlag
	flagLongNames      []NameFlag
}

func createRandErrorSlice(size int) []error {
	res := make([]error, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, errors.New(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, "")))
	}

	return res
}

func createRandSlice[T NameCommand | IDPlaceholder | NamePlaceholder | NameFlag](size uint8, creator func() T) []T {
	uniq := make(map[T]struct{}, size)
	for len(uniq) < int(size) {
		uniq[creator()] = struct{}{}
	}

	res := make([]T, 0, size)
	for k := range uniq {
		res = append(res, k)
	}

	SortSlice(res)

	return res
}

func randName[T NameCommand](lenMin, delta size.Width) T {
	name := strings.TrimSpace(strings.ToLower(gofakeit.Color()))
	for len(name) < lenMin.Int() {
		name += strings.TrimSpace(strings.ToLower(gofakeit.Color()))
	}

	return T(limitString(name, size.MakeWidth(lenMin.Int()+1+abs(rand.Intn(delta.Int())))))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func limitString(s string, limit size.Width) string {
	if len(s) > limit.Int() {
		return s[:limit.Int()]
	}

	return s
}

func createFlagNameOneLetter() string {
	return fmt.Sprintf("-%c", 'A'+rand.Intn(26))
}

func createFlagNameShort() string {
	name := strings.TrimSpace("-" + strings.ToLower(gofakeit.Color()))
	if len(name) > LenNameFlagWithOneDashMax {
		name = name[:LenNameFlagWithOneDashMax]
	}

	return name
}

func createFlagNameLong() string {
	nameFirstPart := strings.ToLower(gofakeit.Color())
	if len(nameFirstPart) > LenNameFlagWithOneDashMax {
		nameFirstPart = nameFirstPart[:LenNameFlagWithOneDashMax]
	}

	nameSecondPart := strings.ToLower(gofakeit.Color())
	if len(nameSecondPart) > LenNameFlagWithOneDashMax {
		nameSecondPart = nameSecondPart[:LenNameFlagWithOneDashMax]
	}

	return fmt.Sprintf("--%s-%s", nameFirstPart, nameSecondPart)
}

var randStorage = func() typeRandStorage {
	args := []string{
		strings.TrimSpace(strings.ToLower(gofakeit.Color())),
		strings.TrimSpace(strings.ToLower(gofakeit.Color())),
	}
	SortSlice(args)

	const (
		countNameCommandShort = 2
		countNameCommandLong  = 3
	)

	commandNamesShort := createRandSlice[NameCommand](countNameCommandShort,
		func() NameCommand {
			return randName[NameCommand](LenNameCommandMin, size.WidthDescriptionColumnShift)
		},
	)
	commandNamesLong := createRandSlice[NameCommand](countNameCommandLong,
		func() NameCommand {
			return randName[NameCommand](size.WidthDescriptionColumnShift, LenNameCommandMax-size.WidthDescriptionColumnShift)
		},
	)

	return typeRandStorage{
		errors: createRandErrorSlice(5),

		appName:         NameApp(strings.ToLower(gofakeit.Color())),
		infoChapterName: InfoChapterNAME(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, "")),
		infoChapterDescription: []InfoChapterDESCRIPTION{
			InfoChapterDESCRIPTION(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, "")),
			InfoChapterDESCRIPTION(gofakeit.Paragraph(1, 1, rand.Intn(10)+1, "")),
		},

		placeholderIDs: createRandSlice[IDPlaceholder](5,
			func() IDPlaceholder {
				return 1 + IDPlaceholder(gofakeit.Uint8())
			},
		),

		commandNameTooShort: randName[NameCommand](1, LenNameCommandMin-2),
		commandNameTooLong:  randName[NameCommand](LenNameCommandMax, LenNameCommandMin),

		commandNamesShort: commandNamesShort,
		commandNamesLong:  commandNamesLong,

		commandNames: func() []NameCommand {
			res := make([]NameCommand, 0, countNameCommandShort+countNameCommandLong)
			res = append(res, commandNamesShort...)
			res = append(res, commandNamesLong...)

			SortSlice(res)

			return res
		}(),

		placeholderNames: createRandSlice[NamePlaceholder](2,
			func() NamePlaceholder {
				return NamePlaceholder(strings.TrimSpace(strings.ToLower(gofakeit.Color())))
			},
		),
		flagOneLetterNames: createRandSlice[NameFlag](3,
			func() NameFlag { return NameFlag(createFlagNameOneLetter()) },
		),
		flagShortNames: createRandSlice[NameFlag](3,
			func() NameFlag { return NameFlag(createFlagNameShort()) },
		),
		flagLongNames: createRandSlice[NameFlag](3,
			func() NameFlag { return NameFlag(createFlagNameLong()) },
		),
	}
}()

// RandError returns random error
func RandError() error {
	return randStorage.errors[0]
}

// RandErrorSecond random error
func RandErrorSecond() error {
	return randStorage.errors[1]
}

// RandErrorThird random error
func RandErrorThird() error {
	return randStorage.errors[2]
}

// RandErrorFourth random error
func RandErrorFourth() error {
	return randStorage.errors[3]
}

// RandErrorFifth random error
func RandErrorFifth() error {
	return randStorage.errors[4]
}

// RandNameApp returns random NameApp
func RandNameApp() NameApp {
	return randStorage.appName
}

// RandInfoChapterName returns random InfoChapterNAME
func RandInfoChapterName() InfoChapterNAME {
	return randStorage.infoChapterName
}

// RandInfoChapterDescription returns random InfoChapterDESCRIPTION
func RandInfoChapterDescription() InfoChapterDESCRIPTION {
	return randStorage.infoChapterDescription[0]
}

// RandInfoChapterDescriptionSecond returns random InfoChapterDESCRIPTION
func RandInfoChapterDescriptionSecond() InfoChapterDESCRIPTION {
	return randStorage.infoChapterDescription[1]
}

// RandIDPlaceholder returns random IDPlaceholder
func RandIDPlaceholder() IDPlaceholder {
	return randStorage.placeholderIDs[0]
}

// RandIDPlaceholderSecond returns random IDPlaceholder
func RandIDPlaceholderSecond() IDPlaceholder {
	return randStorage.placeholderIDs[1]
}

// RandIDPlaceholderThird returns random IDPlaceholder
func RandIDPlaceholderThird() IDPlaceholder {
	return randStorage.placeholderIDs[2]
}

// RandNameCommandTooShort returns random invalid too short command name
func RandNameCommandTooShort() NameCommand {
	return randStorage.commandNameTooShort
}

// RandNameCommandTooLong returns random invalid too long command name
func RandNameCommandTooLong() NameCommand {
	return randStorage.commandNameTooLong
}

// RandNameCommandShort returns random short command name
func RandNameCommandShort() NameCommand {
	return randStorage.commandNamesShort[0]
}

// RandNameCommandShortSecond returns random short command name
func RandNameCommandShortSecond() NameCommand {
	return randStorage.commandNamesShort[1]
}

// RandNameCommandLong returns random long command name
func RandNameCommandLong() NameCommand {
	return randStorage.commandNamesLong[0]
}

// RandNameCommand returns random NameCommand
func RandNameCommand() NameCommand {
	return randStorage.commandNames[0]
}

// RandNameCommandSecond returns random NameCommand
func RandNameCommandSecond() NameCommand {
	return randStorage.commandNames[1]
}

// RandNameCommandThird returns random NameCommand
func RandNameCommandThird() NameCommand {
	return randStorage.commandNames[2]
}

// RandNameCommandFourth returns random NameCommand
func RandNameCommandFourth() NameCommand {
	return randStorage.commandNames[3]
}

// RandNameCommandFifth returns random NameCommand
func RandNameCommandFifth() NameCommand {
	return randStorage.commandNames[4]
}

// RandNamePlaceholder returns random NamePlaceholder
func RandNamePlaceholder() NamePlaceholder {
	return randStorage.placeholderNames[0]
}

// RandNamePlaceholderSecond returns random NamePlaceholder
func RandNamePlaceholderSecond() NamePlaceholder {
	return randStorage.placeholderNames[1]
}

// RandNameFlagOneLetter returns random NameFlag like '-f'
func RandNameFlagOneLetter() NameFlag {
	return randStorage.flagOneLetterNames[0]
}

// RandNameFlagShort returns random NameFlag like '-fg'
func RandNameFlagShort() NameFlag {
	return randStorage.flagShortNames[0]
}

// RandNameFlagShortSecond returns random NameFlag like '-fg'
func RandNameFlagShortSecond() NameFlag {
	return randStorage.flagShortNames[1]
}

// RandNameFlagShortThird returns random NameFlag like '-fg'
func RandNameFlagShortThird() NameFlag {
	return randStorage.flagShortNames[2]
}

// RandNameFlagLong returns random NameFlag like '--name-flag'
func RandNameFlagLong() NameFlag {
	return randStorage.flagLongNames[0]
}
