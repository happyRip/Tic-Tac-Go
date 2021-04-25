package cursor

type Cursor struct {
	X, Y int
}

func (c *Cursor) Up() {
	if c.Y > 0 {
		c.Y--
	} else {
		c.Y = 2
	}
}

func (c *Cursor) Down() {
	if c.Y < 2 {
		c.Y++
	} else {
		c.Y = 0
	}
}

func (c *Cursor) Left() {
	if c.X > 0 {
		c.X--
	} else {
		c.X = 2
	}
}

func (c *Cursor) Right() {
	if c.X < 2 {
		c.X++
	} else {
		c.X = 0
	}
}
