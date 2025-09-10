package main

import "fmt"

// rw-rw-rw {file} owner group
// don't consider leverage
type Cfg struct {
	Auth  uint16 // 333
	Owner string
	Group string
}

type Data string

type Resource struct {
	Cfg
	Data
}

type GroupName string
type UserName string

var UserResource = map[UserName]GroupName{
	"user1": "admin",
}

var Example = map[string]Resource{
	"f1": {
		Cfg: Cfg{
			Auth:  333,
			Owner: "zelin",
			Group: "root",
		},
		Data: "data1",
	},
	"f2": {
		Cfg: Cfg{
			Auth:  330,
			Owner: "zelin",
			Group: "admin",
		},
		Data: "data2",
	},
}

func main() {
	fmt.Println(Read("root", "f1"))
	fmt.Println(Write("user", "f2"))
}

func Read(user, file string) bool {
	return true
}

func Write(user, file string) bool {
	return true
}

func CheckAccess(user, file, oper string) (bool, error) {
	if user == "root" {
		return true, nil
	}

	fileResource, ok := Example[file]
	if !ok {
		return false, fmt.Errorf("no such file")
	}

	if user == fileResource.Owner {
		if CheckOwner(fileResource.Auth, oper) {
			return true, nil
		}
	} else if g, ok := UserResource[UserName(user)]; ok && string(g) == fileResource.Group {
		if CheckGroup(fileResource.Auth, oper) {
			return true, nil
		}
	} else {
		if CheckOther(fileResource.Auth, oper) {
			return true, nil
		}
	}

	return false, nil
}

func CheckOwner(auth uint16, oper string) bool {
	if oper == "read" {
		if (auth & 32) != 0 {
			return true
		}
	} else if oper == "write" {
		if (auth & 16) != 0 {
			return true
		}
	}
	return false
}

func CheckGroup(auth uint16, oper string) bool {
	if oper == "read" {
		if (auth & 8) != 0 {
			return true
		}
	} else if oper == "write" {
		if (auth & 4) != 0 {
			return true
		}
	}
	return false
}

func CheckOther(auth uint16, oper string) bool {
	if oper == "read" {
		if (auth & 2) != 0 {
			return true
		}
	} else if oper == "write" {
		if (auth & 1) != 0 {
			return true
		}
	}
	return false
}
