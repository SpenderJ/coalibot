package Beer

type Bar struct {
	ID   string
	Name string
	Abre []string
}

var Bars = []Bar{
	Bar{
		ID:   "6xCYWThMea",
		Name: "SPRITZ",
	},
	Bar{
		ID:   "TS738y2M42",
		Name: "BRASSERIE DU THÉÂTRE",
		Abre: []string{
			"cdt",
		},
	},
	Bar{
		ID:   "6AGtVK30Dx",
		Name: "MOTY",
	},
}
