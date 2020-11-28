package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Lista struct {
	name  string
	lista []*Producto
}

type Producto struct {
	name   string
	price  float64
	store  string
	amount int
}

func (list *Lista) addToTheList(product *Producto) {
	list.lista = append(list.lista, product)
}

func addProduct(list *Lista) {
	newProduct := &Producto{}
	defer list.addToTheList(newProduct)
	name := read(newProduct, "Nombre de Porducto:")
	price := read(newProduct, "Precio del producto:")
	store := read(newProduct, "Tienda donde se puede comprar:")
	amount := read(newProduct, "Cantidad que debes comprar: ")
	newProduct.name = name
	newProduct.price, _ = strconv.ParseFloat(price, 32)
	newProduct.store = store
	newProduct.amount, _ = strconv.Atoi(amount)

}

func read(pruduct *Producto, text string) (input string) {
	println(text)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input = scan.Text()
	return
}

func (list *Lista) precioLista() (price float64) {
	price = 0
	for i := 0; i < len(list.lista); i++ {
		price += (list.lista[i].price * float64(list.lista[i].amount))
	}
	return

}

func main() {

	product1 := &Producto{
		name:   "Cerveza",
		price:  0.98,
		store:  "Mercadona",
		amount: 1,
	}

	list := &Lista{
		name: "Lista de Mama",
		lista: []*Producto{
			product1,
		},
	}

	addProduct(list)

	fmt.Println("Suma total de la compra: ", list.precioLista(), "€")
}
