package observer

import "fmt"

type Observer interface {
	update(string)
	getID() string
}

type Customer struct {
	ID string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, itemName)
}

func (c *Customer) getID() string {
	return c.ID
}
