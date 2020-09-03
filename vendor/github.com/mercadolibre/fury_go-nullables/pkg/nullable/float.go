package nullable

type Float struct {
	Value    float64
	HasValue bool
}

func FromFloat(f float64) Float {
	return Float{
		Value:    f,
		HasValue: true,
	}
}

func NilFloat() Float {
	return Float{
		Value:    0,
		HasValue: false,
	}
}

func (f Float) MarshalJSON() ([]byte, error) {
	return marshall(f.HasValue, f.Value)
}

func (f *Float) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &f.HasValue, &f.Value)
}

func (f Float) String() string {
	return toString(f.HasValue, f.Value)
}
