package models

var returns map[string]interface{}

func NewReturns(result interface{}, msg interface{}) map[string]interface{} {
	returns = make(map[string]interface{})

	returns["response"] = result
	returns["metaInfo"] = msg

	return returns
}
