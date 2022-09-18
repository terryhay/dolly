package generate

import (
	"fmt"
)

const argParserFileTemplate = `// This code was generated by dolly.generator. DO NOT EDIT

package dolly

import (
	"github.com/terryhay/dolly/pkg/dollyconf"
	"github.com/terryhay/dolly/pkg/dollyerr"
	"github.com/terryhay/dolly/pkg/helpdisplay/data"
	tbd "github.com/terryhay/dolly/pkg/helpdisplay/termbox_decorator"
	"github.com/terryhay/dolly/pkg/helpdisplay/views"
	"github.com/terryhay/dolly/pkg/parsed_data"
	"github.com/terryhay/dolly/pkg/parser"
)
%s
%s
%s

// Parse - processes command line arguments
func Parse(args []string) (res *parsed_data.ParsedData, err *dollyerr.Error) {
	appArgConfig := dollyconf.NewArgParserConfig(%s,%s,%s,%s,%s)

	if res, err = parser.Parse(appArgConfig, args); err != nil {
		return nil, err
	}

	if res.GetCommandID() == %s {
		var pageView views.PageView
		err = pageView.Init(tbd.NewTermBoxDecorator(nil), data.MakePage(appArgConfig))
		if err != nil {
			return nil, err
		}
		err = pageView.Run()
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	return res, nil
}
`

// GenArgParserFileBody applies data to argParserFileTemplate
func GenArgParserFileBody(
	commandIDListSection CommandIDListSection,
	commandNameIDListSection CommandListSection,
	flagStringIDListSection FlagStringIDListSection,
	appDescriptionSection AppDescriptionSection,
	flagDescriptionsSection FlagDescriptionsSection,
	commandDescriptionsSection CommandDescriptionsSection,
	helpCommandDescriptionSection HelpCommandDescriptionSection,
	namelessCommandDescriptionSection NamelessCommandDescriptionSection,
	helpCommandID string) string {

	return fmt.Sprintf(
		argParserFileTemplate,

		commandIDListSection,
		commandNameIDListSection,
		flagStringIDListSection,

		appDescriptionSection,
		flagDescriptionsSection,
		commandDescriptionsSection,
		helpCommandDescriptionSection,
		namelessCommandDescriptionSection,

		helpCommandID,
	)
}
