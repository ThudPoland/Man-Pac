package event

import "go/types"

//Action is a type for making action when something happens in the controller
type Action func(*interface{})

//ActionData contains data for executing action
type ActionData struct {
	sourceType *types.Type
	action     *Action
	dataSource *interface{}
}

//Process executes action
func (actionData *ActionData) Process() {
	if actionData.action != nil && actionData.dataSource != nil {
		(*actionData.action)(actionData.dataSource)
	}
}
