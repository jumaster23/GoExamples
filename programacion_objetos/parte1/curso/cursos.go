package cursos

import(
	"fmt"
)

type curso struct {
	Name    string
	Precio  float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

func New(name string, precio float64, isfree bool)*curso{
	if precio == 0{ precio = 30}
	return &curso{
		Name: name,
		Precio: precio,
		IsFree: isfree,
	}
}

func (c *curso) PrintClasses(){
	text := "Las clases son: "
	for _, class := range c.Classes{
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c *curso) Changeprices(price float64){
	c.Precio = price

}