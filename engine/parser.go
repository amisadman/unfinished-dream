package engine

import (
	"errors"
	"strings"
)

type Statement struct {
	Type string
	Raw  string
}

func Parse(sql string) (*Statement, error) {
	upperSQL := strings.ToUpper(sql)

	if strings.HasPrefix(upperSQL, "SELECT") {
		return &Statement{Type: "SELECT", Raw: sql}, nil
	}
	if strings.HasPrefix(upperSQL, "INSERT") {
		return &Statement{Type: "INSERT", Raw: sql}, nil
	}
	if strings.HasPrefix(upperSQL, "CREATE") {
		return &Statement{Type: "CREATE", Raw: sql}, nil
	}

	return nil, errors.New("unrecognized SQL syntax. the dream doesn't go that deep yet.")
}
