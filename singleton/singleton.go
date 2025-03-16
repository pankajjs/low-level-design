package singleton

import (
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: "Hi",
		}
	})
	
	return instance
}

func (s *Singleton) SetData(data string) {
	s.data = data
}

func (s *Singleton) GetData(data string) string {
	return s.data
}

// func main(){
// 	s1:= GetInstance()
// 	s2:= GetInstance()

// 	s1.SetData("Hello Pankaj")

// 	fmt.Println(s1, s2, s1 == s2, s1.data, s2.data, s1.data == s2.data)
// }