package main

// Ajout des Transactions
func (g *GestionList) Ajout_Transaction(montant int, Type string, categorie string, date int) error {
	if Type == "revenue" {
		*g.Revenue = append(*g.Revenue, Gestion{montant, Type, categorie, date})
	} else if Type == "depense" {
		*g.Depense = append(*g.Depense, Gestion{montant, Type, categorie, date})
	}
	return nil
}

//Affichage des Transactions
