package nullable

import "encoding/json"

type Int struct {
	Value    int64
	HasValue bool
}

func FromInt(i int64) Int {
	return Int{
		Value:    i,
		HasValue: true,
	}
}

func NilInt() Int {
	return Int{
		Value:    0,
		HasValue: false,
	}
}

func (i Int) MarshalJSON() ([]byte, error) {
	if !i.HasValue {
		return json.Marshal(nil)
	}

	return json.Marshal(i.Value)
}

func (i *Int) UnmarshalJSON(data []byte) (err error) {
	return unmarshal(data, &i.HasValue, &i.Value)
}

func (i Int) String() string {
	return toString(i.HasValue, i.Value)
}
