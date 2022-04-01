package optional

import "encoding/json"

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.present {
		return json.Marshal(*o.value)
	}
	return json.Marshal(nil)
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &o.value)
	if err != nil {
		return err
	}

	o.present = o.value != nil
	return err
}
