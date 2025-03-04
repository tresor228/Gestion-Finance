package main

import (
	"fmt"
	/*"os"*/)

// Ajout des Transactions
func (g *GestionList) Ajout_Transaction(montant int, Type string, categorie string, date string) error {

	// Initialisation des listes Dans le tableau
	if g.Revenue == nil {
		g.Revenue = &[]Gestion{}
	}

	if g.Depense == nil {
		g.Depense = &[]Gestion{}
	}

	// Gestion d'erreur
	if Type != "Revenue" && Type != "Dépense" {
		return fmt.Errorf("Type de transaction invalide ")
	}

	//classement du type
	if Type == "revenue" {
		*g.Revenue = append(*g.Revenue, Gestion{montant, Type, categorie, date})
	} else if Type == "depense" {
		*g.Depense = append(*g.Depense, Gestion{montant, Type, categorie, date})
	}

	//Ajout dans le fichier csv

}

//Fonction pour l'enregistement dans le fichier csv

// Affichage du contenu CSV (Transaction)
func (g *GestionList) Affich_Transaction() {
	fmt.Println("	categorie	|	Montant	|	Type	|	Date	")
	if g.Revenue != nil {
		for _, Gestion := range *g.Revenue {
			fmt.Printf(" %s	|	%d	|	%s	|	%s	", Gestion.Categorie, Gestion.Montant, Gestion.Type, Gestion.Date)
		}
	}

	if g.Depense != nil {
		for _, Gestion := range *g.Depense {
			fmt.Printf(" %s	|	%d	|	%s	|	%s	", Gestion.Categorie, Gestion.Montant, Gestion.Type, Gestion.Date)
		}
	}
}

func main() {

}
