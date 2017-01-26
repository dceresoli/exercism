package robotname

import (
	"math/rand"
)

var RobotFactory = make(map[string]bool)

func genName() string {
	res := ""
	for {
		res += string('A' + rand.Intn(26))
		res += string('A' + rand.Intn(26))
		res += string('0' + rand.Intn(10))
		res += string('0' + rand.Intn(10))
		res += string('0' + rand.Intn(10))
		if RobotFactory[res] == false {
			RobotFactory[res] = true
			break
		}
	}
	return res
}

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = genName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = genName()
}
