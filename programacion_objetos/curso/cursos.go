package cursos

import(
	"fmt"
)

type Curso struct {
	Name    string
	Precio  float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

func (c *Curso) PrintClasses(){
	text := "Las clases son: "
	for _, class := range c.Classes{
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c *Curso) Changeprices(price float64){
	c.Precio = price

}