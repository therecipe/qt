package model

type Artist struct {
	ID     int
	Name   string
	Albums map[int]*Album
}

type Album struct {
	ID    int
	Title string
	Year  int
}

type dbArrayStruct struct {
	Artist *Artist
	Album  *Album
}

var db = map[int]*Artist{
	0: &Artist{
		ID:   0,
		Name: "<all>",
	},

	1: &Artist{
		ID:   1,
		Name: "Ane Brun",
		Albums: map[int]*Album{
			1: &Album{ID: 1, Title: "Spending Time With Morgan", Year: 2003},
			2: &Album{ID: 2, Title: "A Temporary Dive", Year: 2005},
		},
	},

	2: &Artist{
		ID:   2,
		Name: "Thomas Dybdahl",
		Albums: map[int]*Album{
			3: &Album{ID: 3, Title: "...The Great October Sound", Year: 2002},
			4: &Album{ID: 4, Title: "Stray Dogs", Year: 2003},
			5: &Album{ID: 5, Title: "One day you`ll dance for me, New York City", Year: 2004},
		},
	},

	3: &Artist{
		ID:   3,
		Name: "Kaizers Orchestra",
		Albums: map[int]*Album{
			6: &Album{ID: 6, Title: "Ompa Til Du D\xc3\xb8r", Year: 2001},
			7: &Album{ID: 7, Title: "Evig Pint", Year: 2002},
			8: &Album{ID: 8, Title: "Maestro", Year: 2005},
		},
	},
}
