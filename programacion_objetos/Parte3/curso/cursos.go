package cursos

import(
	"fmt"
)

type curso struct {
	name    string
	precio  float64
	isFree  bool
	userID  []uint
	classes map[uint]string
}

//Setter y getters para classes
func (c *curso) Setclasses(classes map[uint]string){c.classes = classes}
func (c *curso) Classes()map[uint]string{return c.classes}

//Setter y getters para name
func (c *curso) Setname(name string){c.name = name}
func (c *curso) Name()string{return c.name}

//Setter y getters para isfree
func (c *curso) Setisfree(isfree bool){c.isFree = isfree}
func (c *curso) Isfree()bool{return c.isFree}

//Setter y getters para userID
func (c *curso) SetUserId(userID []uint){c.userID = userID}
func (c *curso) Userid()[]uint{return c.userID}

//setter y getter precio
func (c *curso) Setprecio(precio float64){c.precio = precio}
func (c *curso) Precio() float64{ return c.precio}



func New(name string, precio float64, isfree bool)*curso{
	if precio == 0{ precio = 30}
	return &curso{
		name: name,
		precio: precio,
		isFree: isfree,
	}
}

func (c *curso) PrintClasses(){
	text := "Las clases son: "
	for _, class := range c.classes{
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

