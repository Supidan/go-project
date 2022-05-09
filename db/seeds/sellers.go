package seeds

import "github.com/kristijorgji/goseeder"

func sellerSeeder(s goseeder.Seeder) {
	s.FromJson("sellers")
}
