package main

// Struct de Gestion
type Gestion struct {
	Montant   int
	Type      string
	Categorie string
	Date      string
}

// Struct pour trié les revenus et les dépenses
type GestionList struct {
	Revenue *[]Gestion
	Depense *[]Gestion
}
