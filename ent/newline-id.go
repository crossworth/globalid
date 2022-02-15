package ent

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var typeLookupRef = map[string]string{
	"13": "user",
}

var typeLookupTable = map[string]string{
	"user": "13",
}

func newLineMarshal(g GlobalID) string {
	id := fmt.Sprintf("%s\n%s", typeLookupTable[g.Type], g.ID)
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func init() {
	marshalGlobalID = newLineMarshal

	unmarshalGlobalID = func(v interface{}) (string, string, error) {
		id := v.(string)
		b, err := base64.URLEncoding.DecodeString(id)
		if err != nil {
			return "", "", err
		}
		tid := strings.Split(string(b), "\n")
		if len(tid) != 2 {
			return "", "", fmt.Errorf("invalid global identifier format %q", id)
		}
		return typeLookupRef[tid[0]], tid[1], nil
	}
}
