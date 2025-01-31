// Interface Segregation Principle
package interfacesegregation

type RestaurantEmpolyee interface{
	cook()
	takeOrder()
}

type Chef struct {}
func (c *Chef) cook() {}
func (c *Chef) takeOrder() { 
	// Here it is breaking interface segregation principle which says
	// client should not implement unnecessary functions
	// not the job of chef
}

// Solution
type RestaurantChef interface {
	cook()
}
type RestaurantWaiter interface {
	takeOrder()
}

type ChefS struct {}
func (c *ChefS) cook () {}

type Waiter struct {}
func (w *Waiter) takeOrder() {}
