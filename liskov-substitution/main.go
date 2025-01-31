// Liskov Substitution Principle
package liskovsubstitution

type Vehicle interface {
	startEngine()
}

type Car struct {}
func (c *Car) startEngine() {
	// 
}

type BiCycle struct {}
func (b *BiCycle) startEngine() {
	panic("Engine not found")
}

func main() {
	car := &Car{}
	car.startEngine()

	bicycle := &BiCycle{}
	// This is not possible because bicycle does not have engine
	//  which breaks the liskov substitution principle which says a subclass shoud be able to perform action of parent without breaking the current behaviour
	// Subclass should be able to replace parent class
	bicycle.startEngine()
}


// Solution
// Create separate interface for different vehicle type
type EngineVehicle interface {
	startEngine()
}
type EngineLessVehicle interface {
	satrt()
}
type CarS struct {}
func (c *CarS) startEngine () {}

type BiCycleS struct {}
func (b *BiCycleS) start(){}
