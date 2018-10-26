package main

type AddCommand struct {
}

func (AddCommand) Name() string {
	return "add"
}

func (AddCommand) Exec(ctx *Context) error {
	return nil
}

func (AddCommand) Usage() string {
	return "add [table] add table into cache，`*` means all the tables belong the database you chose by command `use`"
}

func init() {
	Register(AddCommand{})
}
