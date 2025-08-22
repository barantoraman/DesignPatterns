package command

type Command interface {
	execute()
}

type OnCommand struct {
	Device Device
}

func (c *OnCommand) execute() {
	c.Device.on()
}

type OffCommand struct {
	Device Device
}

func (c *OffCommand) execute() {
	c.Device.off()
}
