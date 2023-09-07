// Code generated by "enumer --type=AuthorizationServerType --trimprefix=AuthorizationServerType -json"; DO NOT EDIT.

package config

import (
	"encoding/json"
	"fmt"
)

const _AuthorizationServerTypeName = "SelfExternal"

var _AuthorizationServerTypeIndex = [...]uint8{0, 4, 12}

func (i AuthorizationServerType) String() string {
	if i < 0 || i >= AuthorizationServerType(len(_AuthorizationServerTypeIndex)-1) {
		return fmt.Sprintf("AuthorizationServerType(%d)", i)
	}
	return _AuthorizationServerTypeName[_AuthorizationServerTypeIndex[i]:_AuthorizationServerTypeIndex[i+1]]
}

var _AuthorizationServerTypeValues = []AuthorizationServerType{0, 1}

var _AuthorizationServerTypeNameToValueMap = map[string]AuthorizationServerType{
	_AuthorizationServerTypeName[0:4]:  0,
	_AuthorizationServerTypeName[4:12]: 1,
}

// AuthorizationServerTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AuthorizationServerTypeString(s string) (AuthorizationServerType, error) {
	if val, ok := _AuthorizationServerTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AuthorizationServerType values", s)
}

// AuthorizationServerTypeValues returns all values of the enum
func AuthorizationServerTypeValues() []AuthorizationServerType {
	return _AuthorizationServerTypeValues
}

// IsAAuthorizationServerType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AuthorizationServerType) IsAAuthorizationServerType() bool {
	for _, v := range _AuthorizationServerTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AuthorizationServerType
func (i AuthorizationServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AuthorizationServerType
func (i *AuthorizationServerType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AuthorizationServerType should be a string, got %s", data)
	}

	var err error
	*i, err = AuthorizationServerTypeString(s)
	return err
}
