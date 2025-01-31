// Dependency Inversion Principle
package dependencyinversion

import "fmt"

type DaoInterface interface {
	save()
}

// implement DaoInterface
type Dao struct {

}

func (d *Dao) save() {
	fmt.Print("save")
}

type Service struct {
	//  depend on concrete implementation
	dao *Dao
}

// Here it breaks the dependency inversion principle
// which says the high level module should not depend upon low level module rather it should 
// depend upon interface

// Solution
type ServiceS struct {
	dao DaoInterface
}

func main() {
	dao := &Dao{}
	s := &ServiceS{
		dao: dao,
	}
	s.dao.save()
}