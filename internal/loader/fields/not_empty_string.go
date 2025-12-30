package fields

import (
	"encoding/json"
	"fmt"
)

type NotEmptyString string

func (s *NotEmptyString) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if raw == "" {
		return fmt.Errorf("string is empty")
	}
	*s = NotEmptyString(raw)
	return nil
}

func (s *NotEmptyString) String() string {
	return string(*s)
}
