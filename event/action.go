package event

//Action is a type for making action when something happens in the controller
type Action func(interface{})

//ActionData contains data for executing action
type ActionData struct {
	action   func(interface{})
	receiver interface{}
}

//Process executes action
func (actionData *ActionData) Process() {
	if actionData.action != nil && actionData.receiver != nil {
		(actionData.action)(actionData.receiver)
	}
}
