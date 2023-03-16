package parsed

import (
	"errors"
	"fmt"

	coty "github.com/terryhay/dolly/tools/common_types"
)

// Result contains a result of nameMainCommand line arguments parsing
type Result struct {
	nameMainCommand   coty.NameCommand
	placeholdersByIDs map[coty.IDPlaceholder]*Placeholder
}

// GetCommandMainName gets nameMainCommand field
func (i *Result) GetCommandMainName() coty.NameCommand {
	if i == nil {
		return coty.NameCommandUndefined
	}
	return i.nameMainCommand
}

// GetPlaceholdersByIDs gets placeholdersByIDs field
func (i *Result) GetPlaceholdersByIDs() map[coty.IDPlaceholder]*Placeholder {
	if i == nil {
		return nil
	}
	return i.placeholdersByIDs
}

// FlagArgValue extracts flagName argument value
func (i *Result) FlagArgValue(name coty.NameFlag) (value ArgValue, err error) {
	var values []ArgValue
	values, err = i.FlagArgValues(name)
	if err != nil || len(values) == 0 {
		return ArgValueDefault, err
	}

	return values[0], nil
}

var (
	// ErrFlagArgValuesNilPointer - call by nil pointer
	ErrFlagArgValuesNilPointer = errors.New(`Result.FlagArgValues error: Result pointer is nil`)

	// ErrFlagArgValuesNoArgValues - placeholder doesn't contain parsed arguments
	ErrFlagArgValuesNoArgValues = errors.New(`Result.FlagArgValues error: placeholder doesn't contain parsed arguments`)

	// ErrFlagArgValuesNoPlaceholder - placeholder is not found by nameFlag
	ErrFlagArgValuesNoPlaceholder = errors.New(`Result.FlagArgValues error: placeholder is not found`)
)

// FlagArgValues - extract flagName argument value slice
func (i *Result) FlagArgValues(name coty.NameFlag) (values []ArgValue, err error) {
	if i == nil {
		return nil, ErrFlagArgValuesNilPointer
	}

	for _, placeholder := range i.placeholdersByIDs {
		if placeholder.GetNameFlag() != name {
			continue
		}

		if len(placeholder.GetArgData().GetArgValues()) == 0 {
			return nil, fmt.Errorf(`%w: NameFlag '%s'`, ErrFlagArgValuesNoArgValues, name)
		}
		return placeholder.GetArgData().GetArgValues(), nil
	}

	return nil, fmt.Errorf(`%w: NameFlag '%s'`, ErrFlagArgValuesNoPlaceholder, name)
}

var (
	// ErrPlaceholderArgValuesNilPointer - call by nil pointer
	ErrPlaceholderArgValuesNilPointer = errors.New(`Result.PlaceholderArgValues error: Result pointer is nil`)

	// ErrPlaceholderArgValuesNoArgValues - placeholder doesn't contain parsed arguments
	ErrPlaceholderArgValuesNoArgValues = errors.New(`Result.PlaceholderArgValues error: placeholder doesn't contain parsed arguments`)

	// ErrPlaceholderArgValuesNoPlaceholder - placeholder is not found by nameFlag
	ErrPlaceholderArgValuesNoPlaceholder = errors.New(`Result.PlaceholderArgValues error: placeholder is not found`)
)

func (i *Result) PlaceholderArgValues(id coty.IDPlaceholder) (values []ArgValue, err error) {
	if i == nil {
		return nil, ErrPlaceholderArgValuesNilPointer
	}

	placeholder, contain := i.placeholdersByIDs[id]
	if !contain {
		return nil, fmt.Errorf(`%w: IDPlaceholder '%d'`, ErrPlaceholderArgValuesNoPlaceholder, id)
	}

	if len(placeholder.GetArgData().GetArgValues()) == 0 {
		return nil, fmt.Errorf(`%w: IDPlaceholder '%d`, ErrPlaceholderArgValuesNoArgValues, id)
	}
	return placeholder.GetArgData().GetArgValues(), nil
}

func (i *Result) PlaceholderArgValue(id coty.IDPlaceholder) (value ArgValue, err error) {
	var values []ArgValue
	values, err = i.PlaceholderArgValues(id)
	if err != nil || len(values) == 0 {
		return ArgValueDefault, err
	}

	return values[0], nil
}

// PlaceholderByID returns Placeholder by IDPlaceholder
func (i *Result) PlaceholderByID(id coty.IDPlaceholder) *Placeholder {
	if i == nil {
		return nil
	}

	return i.placeholdersByIDs[id]
}

// ResultOpt contains source data for cast to Result
type ResultOpt struct {
	CommandMainName  coty.NameCommand
	PlaceholdersByID map[coty.IDPlaceholder]*PlaceholderOpt
}

// SetFlagName sets NameFlag into Placeholder by IDPlaceholder
func (opt *ResultOpt) SetFlagName(placeholderID coty.IDPlaceholder, flag coty.NameFlag) {
	if opt == nil {
		return
	}
	if opt.PlaceholdersByID == nil {
		opt.PlaceholdersByID = make(map[coty.IDPlaceholder]*PlaceholderOpt, 8)
	}

	argGroup, contain := opt.PlaceholdersByID[placeholderID]
	if !contain {
		argGroup = &PlaceholderOpt{
			ID: placeholderID,
		}
		opt.PlaceholdersByID[placeholderID] = argGroup
	}
	argGroup.Flag = flag
}

// SetArg sets ArgValue into Placeholder by IDPlaceholder
func (opt *ResultOpt) SetArg(argGroupID coty.IDPlaceholder, arg ArgValue) {
	if opt == nil {
		return
	}
	if opt.PlaceholdersByID == nil {
		opt.PlaceholdersByID = make(map[coty.IDPlaceholder]*PlaceholderOpt, 8)
	}

	argGroup, contain := opt.PlaceholdersByID[argGroupID]
	if !contain {
		argGroup = &PlaceholderOpt{
			ID: argGroupID,
		}
		opt.PlaceholdersByID[argGroupID] = argGroup
	}
	argData := argGroup.Argument
	if argData == nil {
		argData = &ArgumentOpt{
			ArgValues: make([]ArgValue, 0, 1),
		}
		argGroup.Argument = argData
	}
	argData.ArgValues = append(argData.ArgValues, arg)
}

// PlaceholderByID returns Placeholder by IDPlaceholder
func (opt *ResultOpt) PlaceholderByID(argGroupID coty.IDPlaceholder) *PlaceholderOpt {
	if opt == nil {
		return nil
	}
	return opt.PlaceholdersByID[argGroupID]
}

// PlaceholderDoesNotHaveArgs returns if the group with argGroupID doesn't have arguments
func (opt *ResultOpt) PlaceholderDoesNotHaveArgs(argGroupID coty.IDPlaceholder) bool {
	if opt == nil {
		return true
	}
	argGroup, contain := opt.PlaceholdersByID[argGroupID]
	if !contain {
		return true
	}
	if argGroup.Argument == nil {
		return true
	}
	return len(argGroup.Argument.ArgValues) <= 0
}

// MakeResult converts opt to Result pointer
func MakeResult(opt *ResultOpt) *Result {
	if opt == nil {
		return nil
	}

	if len(opt.PlaceholdersByID) == 0 {
		return &Result{
			nameMainCommand: opt.CommandMainName,
		}
	}

	argGroups := make(map[coty.IDPlaceholder]*Placeholder, len(opt.PlaceholdersByID))
	for id, data := range opt.PlaceholdersByID {
		argGroups[id] = NewPlaceholder(data)
	}

	return &Result{
		nameMainCommand:   opt.CommandMainName,
		placeholdersByIDs: argGroups,
	}
}
