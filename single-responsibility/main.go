// Single Responsibility Principle
package singleresponsibility

type Marker struct {
	price int
}

type Invoice struct {
	marker *Marker
	quantity int
}

func NewInvoice(marker *Marker, quantity int) *Invoice {
	return &Invoice{
		marker: marker,
		quantity: quantity,
	}
}

func (i *Invoice) calculateTotal() int {
	return i.marker.price * i.quantity
}

func (i *Invoice) printInvoice(){
	
}

func (i *Invoice) saveInvoiceToDB(){

}

// Here Invoice has multiple responsibilities
// Invoice can be changed because of multiple reasons which breaks the single responsibility principle
// Solution

type InvoicePrinter struct {
	invoice *Invoice
}

func NewInvoicePrinter(invoice *Invoice) *InvoicePrinter {
	return &InvoicePrinter{
		invoice: invoice,
	}
}

func (i *InvoicePrinter) print() {

}

type InvoiceDao struct {
	invoice * Invoice
}

func NewInvoiceDao(invoice * Invoice) *InvoiceDao {
	return &InvoiceDao{
		invoice: invoice,
	}
}

func (i *InvoiceDao) save(){
	
}



