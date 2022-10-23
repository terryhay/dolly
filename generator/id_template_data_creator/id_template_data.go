package id_template_data_creator

// IDTemplateData - page for fill up templates
type IDTemplateData struct {
	id       string
	nameID   string
	callName string
	comment  string
}

// NewIDTemplateData - IDTemplateData object constructor
func NewIDTemplateData(id, stringID, callName, comment string) *IDTemplateData {
	return &IDTemplateData{
		id:       id,
		nameID:   stringID,
		callName: callName,
		comment:  comment,
	}
}

// GetID - id field getter
func (i *IDTemplateData) GetID() string {
	if i == nil {
		return ""
	}
	return i.id
}

// GetNameID - nameID field getter
func (i *IDTemplateData) GetNameID() string {
	if i == nil {
		return ""
	}
	return i.nameID
}

// GetCallName - field callName getter
func (i *IDTemplateData) GetCallName() string {
	if i == nil {
		return ""
	}
	return i.callName
}

// GetComment - field comment getter
func (i *IDTemplateData) GetComment() string {
	if i == nil {
		return ""
	}
	return i.comment
}
