package seeds

import "github.com/kristijorgji/goseeder"

func productSeeder(s goseeder.Seeder) {
	s.FromJson("products")
}
