// Open closed Principle
package openclosed

type Marker struct {
	price int
}

type Invoice struct {
	marker *Marker
	quantity int
}

type InvoiceDao struct {
	invoice *Invoice
}

func (i *InvoiceDao) saveToDB() {

}

//  Let's consider we have a logic to save invoice in db
// Now, There is one more requirement to save invoice in file
// So, what we did, we added new function within  class/struct
// which breaks the Open Close Principle which says class should be open for extension but should be close for modification
func (i *InvoiceDao) saveToFile() {

}

// Solution
type InvoiceDaoI interface {
	save()
}

type InvoicDaoDB struct {
	invoice *Invoice
}

func (i *InvoicDaoDB) save() {

}

type InvoiceDaoFile struct {
	invoice *Invoice
}

func (i *InvoiceDaoFile) save() {
	
}