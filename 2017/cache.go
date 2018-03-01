package main

func (c *Cache) Insert(v *Video) bool {

	for _, v1 := range c.Videos {
		if v1.index == v.index {
			return false
		}
	}
	if c.Capacity-v.Size > 0 {
		c.Capacity -= v.Size
		c.Videos = append(c.Videos, v)
		return true
	}
	return false

}
