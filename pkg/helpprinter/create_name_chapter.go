package helpprinter

import (
	"fmt"
)

const nameChapter = "\u001B[1mNAME\u001B[0m\n\t\u001B[1m%s\u001B[0m – %s\n\n"

// CreateNameChapter creates name help chapter
func CreateNameChapter(appName, nameHelpInfo string) string {
	return fmt.Sprintf(nameChapter, appName, nameHelpInfo)
}
