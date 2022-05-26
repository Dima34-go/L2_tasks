package command

type Cook struct {
	commands []Command
}

func (c *Cook) executeCommands() {
	for _,cmd:=range c.commands {
		cmd.execute()
	}
}
