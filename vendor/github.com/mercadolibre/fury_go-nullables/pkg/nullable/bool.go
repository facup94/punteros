package nullable

type Bool struct {
	Value    bool
	HasValue bool
}

func FromBool(b bool) Bool {
	return Bool{
		Value:    b,
		HasValue: true,
	}
}

func NilBool() Bool {
	return Bool{
		Value:    false,
		HasValue: false,
	}
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.HasValue {
		return json.Marshal(nil)
	}

	return json.Marshal(b.Value)
}

func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &b.HasValue, &b.Value)
}

func (b Bool) String() string {
	return toString(b.HasValue, b.Value)
}
