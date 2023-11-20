package model

func (d MessageData) GetDescription() string {
	return d.Description
}

func (d HelpData) GetDescription() string {
	return d.Description
}

func (d LookingForData) GetDescription() string {
	return d.Description
}

func (d NotLookingForData) GetDescription() string {
	return d.Description
}

func (d Lectures) GetDescription() string {
	return d.Description
}

func (d ListData) GetDescription() string {
	return d.Description
}

func (d LuckData) GetDescription() string {
	return d.Description
}

func (d InvalidData) GetDescription() string {
	return "This data is invalidly parsed, please report this bug to the developer."
}
