package seeds

import "github.com/kristijorgji/goseeder"

func buyerSeeder(s goseeder.Seeder) {
	s.FromJson("buyers")
}
