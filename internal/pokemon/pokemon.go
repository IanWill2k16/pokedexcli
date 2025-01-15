package pokemon

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          Stats
	Types          []string
}

type Stats struct {
	HP             int
	Attack         int
	Defense        int
	SpecialAttack  int
	SpecialDefense int
	Speed          int
}
