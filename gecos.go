package gecos

/*
#include <pwd.h>
#include <malloc.h>
*/
import "C"

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
)

var (
	NoGECOSEntry = fmt.Errorf("No GECOS Entry for this user")
)

type GECOS struct {
	Name        string
	Room        string
	OfficePhone string
	HomePhone   string
	Other       string
}

func getpwuid(user int) (*GECOS, error) {
	gecos := C.getpwuid(C.__uid_t(user))

	gecosString := C.GoString(gecos.pw_gecos)

	els := strings.SplitN(gecosString, ",", 5)

	if len(els) == 1 {
		return nil, NoGECOSEntry
	}
	if len(els) < 4 {
		return nil, fmt.Errorf("Expected 4 or 5 elements, got %d", len(els))
	}
	ret := GECOS{
		Name:        els[0],
		Room:        els[1],
		OfficePhone: els[2],
		HomePhone:   els[3],
	}
	if len(els) >= 5 {
		ret.Other = els[4]
	}
	return &ret, nil
}

func Lookup(user *user.User) (*GECOS, error) {
	id, err := strconv.Atoi(user.Uid)
	if err != nil {
		return nil, err
	}
	return getpwuid(id)
}
