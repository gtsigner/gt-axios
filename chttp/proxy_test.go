package chttp

import (
	"fmt"
	"testing"
)

func TestGetProxyTypes(t *testing.T) {
	var types = GetProxyTypes()
	fmt.Println(types)
}
