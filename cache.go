package main

func (c *Cache) Insert(v *Video) bool {
	if c.Capacity-v.Size > 0 {
		c.Capacity -= v.Size
		c.Videos = append(c.Videos, v)
		return true
	}
	return false

}
