package nullable

type String struct {
	Value    string
	HasValue bool
}

func FromString(s string) String {
	return String{
		Value:    s,
		HasValue: true,
	}
}

func NilString() String {
	return String{
		Value:    "",
		HasValue: false,
	}
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.HasValue {
		return json.Marshal(nil)
	}

	return json.Marshal(s.Value)
}

func (s *String) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &s.HasValue, &s.Value)
}

func (s String) String() string {
	return toString(s.HasValue, s.Value)
}
