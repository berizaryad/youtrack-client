package youtrack

import (
	"encoding/json"
	"fmt"
)

func (i *IssueCustomField) UnmarshalJSON(data []byte) error {
	type tmpT IssueCustomField
	var tmp tmpT
	var valueResultType any

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Value == nil {
		return nil
	}

	switch value := tmp.Value.(type) {
	case map[string]any:
		typ, ok := value[typeField]
		if !ok {
			return nil
		}
		typStr, ok := typ.(string)
		if !ok {
			return fmt.Errorf(typeField + " field casting to string: not a string")
		}

		if typStr == periodValueType {
			valueResultType = &PeriodValue{}
		}
	}

	valueJSON, err := json.Marshal(tmp.Value)
	if err != nil {
		return err
	}

	err = json.Unmarshal(valueJSON, valueResultType)
	if err != nil {
		return err
	}
	tmp.Value = valueResultType
	*i = IssueCustomField(tmp)

	return nil
}
