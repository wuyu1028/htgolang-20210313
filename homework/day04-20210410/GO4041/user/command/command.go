package command

import "strconv"

var id int = 1

var Commands map[string]func() = map[string]func(){}

var Prompts [][2]string = [][2]string{}

func Register(desc string, callback func()) {
	Commands[strconv.Itoa(id)] = callback
	Prompts = append(Prompts, [2]string{strconv.Itoa(id), desc})
	id++
}
