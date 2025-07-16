package types

type Seeding string

const (
	Seed      Seeding = "seed"
	Reseed    Seeding = "reseed"
	Truncated Seeding = "truncate"
)
