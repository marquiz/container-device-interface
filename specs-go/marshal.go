/*
   Copyright © The CDI Authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package specs

import (
	"encoding/json"
	"fmt"

	"go.yaml.in/yaml/v3"
)

const anyMinorString = "any"

// MarshalJSON implements json.Marshaler.
func (m DeviceCgroupMinor) MarshalJSON() ([]byte, error) {
	if m == DeviceCgroupMinorAny {
		return json.Marshal(anyMinorString)
	}
	return json.Marshal(int64(m))
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *DeviceCgroupMinor) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil && s == anyMinorString {
		*m = DeviceCgroupMinorAny
		return nil
	}

	var n int64
	if err := json.Unmarshal(data, &n); err != nil {
		return fmt.Errorf("invalid device cgroup minor, must be an integer or %q", anyMinorString)
	}
	*m = DeviceCgroupMinor(n)
	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (m DeviceCgroupMinor) MarshalYAML() (any, error) {
	if m == DeviceCgroupMinorAny {
		return anyMinorString, nil
	}
	return int64(m), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (m *DeviceCgroupMinor) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind == yaml.ScalarNode && node.Value == anyMinorString {
		*m = DeviceCgroupMinorAny
		return nil
	}

	var n int64
	if err := node.Decode(&n); err != nil {
		return fmt.Errorf("invalid device cgroup minor, must be an integer or %q", anyMinorString)
	}
	*m = DeviceCgroupMinor(n)
	return nil
}
