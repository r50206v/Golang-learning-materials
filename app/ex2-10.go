package main

import "fmt"

type vehicle struct {
	doors int 
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return fmt.Sprintf("a %d doors and %s color truck with fourwheel: %t", 
		t.vehicle.doors, t.vehicle.color, t.fourWheel)
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintf("a %d doors and %s color sedan is luxurious: %t", 
		s.vehicle.doors, s.vehicle.color, s.luxury)
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) string {
	return t.transportationDevice()
}

func main() {
	truck1 := truck{
		vehicle{2, "blue"}, 
		false,
	}
	sedan1 := sedan{
		vehicle{4, "grey"}, 
		true,
	}
	
	t_desc := report(truck1)
	s_desc := report(sedan1)
	fmt.Println(t_desc)
	fmt.Println(s_desc)
}