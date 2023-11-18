package core

var Version = "Nasus Server v0.0.0" // VERSION OF SOFTWARE

// This is probably the most hardcoded part of this project (I don't like hardcode)

// Resolves user args
func ResolveArgs(args []string) []string {
	return args
}

type param struct {
	name string
	req  bool
	opts []string
}

type arg struct {
	name   string
	desc   string
	call   []string
	run    func(any) any
	params []param
	subarg []arg
}

var Args = []arg{
	// arg{
	// 	name: "help",
	// 	desc: "This help menu",
	// 	call: []string{"h"},
	// 	run:  helpFunc,
	// },
}