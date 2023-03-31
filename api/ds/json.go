/*
	reFES - A RPG Maker FES server emulator
	Copyright (C) 2023  maru <maru@myyahoo.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
