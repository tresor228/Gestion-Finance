package main

// Struct de Gestion
type Gestion struct {
	Montant   int
	Type      string
	Categorie string
	Date      int
}

// Struct pour trié les revenus et les dépenses
type GestionList struct {
	Revenue *[]Gestion
	Depense *[]Gestion
}
