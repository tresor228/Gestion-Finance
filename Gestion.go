package main

// Struct de Gestion
type Gestion struct {
	Date      string
	Type      string
	Categorie string
	Montant   int
}

// Struct pour trié les revenus et les dépenses
type GestionList struct {
	Revenue *[]Gestion
	Depense *[]Gestion
}
