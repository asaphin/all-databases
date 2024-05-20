package app

type EntityKey []interface{}

func NewEntityKey(keyComponents ...interface{}) EntityKey {
	return keyComponents
}

func (k *EntityKey) String(index int) string {
	if index >= 0 && index < len(*k) {
		return (*k)[index].(string)
	}

	return ""
}

func (k *EntityKey) Int(index int) int {
	if index >= 0 && index < len(*k) {
		return (*k)[index].(int)
	}

	return 0
}

type Entity struct {
	Resource string    `json:"resource"`
	Key      EntityKey `json:"key"`
}

func NewEntity(resource string, keyComponents ...interface{}) *Entity {
	return &Entity{
		Resource: resource,
		Key:      keyComponents,
	}
}

type Action struct {
	name     string
	function func() error
}
