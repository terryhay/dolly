package generate

import confYML "github.com/terryhay/dolly/generator/config_yaml"

// helpOutSection contains golang paste code with a call of help out method
type helpOutSection string

// genHelpOutSection
func genHelpOutSection(helpOutTool confYML.HelpOutTool) helpOutSection {
	if helpOutTool == confYML.HelpOutToolManStyle {
		return `
		var pageView pgv.PageView
		err = pageView.Init(tbd.NewTermBoxDecorator(nil), page.MakePage(appArgConfig))
		if err != nil {
			return nil, err
		}
		err = pageView.Run()
		if err != nil {
			return nil, err
		}

		return nil, nil`
	}
	return `
		helpOut.PrintHelpInfo(appArgConfig)
		return nil, nil`
}
