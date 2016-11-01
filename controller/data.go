package controller

type KeyData struct {
	Key   string
	Value int
}

func UpdateCompoundedValues(keyController Controller, compoundedData int, keys []KeyData) int {
	for _, key := range keys {
		if compoundedData >= key.Value {
			compoundedData -= key.Value
			keyController.SetKey(key.Key, true)
		} else {
			keyController.SetKey(key.Key, false)
		}
	}
	return compoundedData
}
