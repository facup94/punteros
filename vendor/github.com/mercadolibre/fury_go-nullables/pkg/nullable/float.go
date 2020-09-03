package nullable

import "encoding/json"

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
	if !f.HasValue {
		return json.Marshal(nil)
	}

	return json.Marshal(f.Value)
}

func (f *Float) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &f.HasValue, &f.Value)
}

func (f Float) String() string {
	return toString(f.HasValue, f.Value)
}
