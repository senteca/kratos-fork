package selfservice

type FormField struct {
	Name     string `json:"name"`
	Type     string `json:"type,omitempty"`
	Required bool   `json:"required,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Error    string `json:"error,omitempty"`
}

type FormFields map[string]FormField

func (fs FormFields) Reset() {
	for k, f := range fs {
		f.Error = ""
		f.Value = ""
		fs[k] = f
	}
}

func (fs FormFields) SetValue(name string, value interface{}) {
	var field FormField
	if ff, ok := fs[name]; ok {
		field = ff
	}

	field.Name = name
	field.Value = value
	fs[name] = field
}

func (fs FormFields) SetError(name, err string) {
	var field FormField
	if ff, ok := fs[name]; ok {
		field = ff
	}

	field.Name = name
	field.Error = err
	fs[name] = field
}
