package ds

import (
	"encoding/json"
	"strings"
)

func (l *RpgListS) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &l.RpgListEntries)
	if err != nil {
		if !strings.Contains(err.Error(), "cannot unmarshal number") {
			return err
		}
	}

	// assume error code is always 0

	return nil
}

func (l RpgListS) MarshalJSON() ([]byte, error) {
	tmp := make(map[string]any)

	for id, entry := range l.RpgListEntries {
		tmp[id] = entry
	}

	tmp["endcode"] = l.EndCode

	return json.Marshal(tmp)
}
