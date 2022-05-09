package seeds

import "github.com/kristijorgji/goseeder"

func init() {
	goseeder.Register(buyerSeeder)
	goseeder.Register(sellerSeeder)
	goseeder.Register(productSeeder)
}
