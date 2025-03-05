package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
		return fmt.Errorf("type de transaction Disponible : Revenue ou Dépense")
	}

	//classement du type
	if Type == "revenue" {
		*g.Revenue = append(*g.Revenue, Gestion{montant, Type, categorie, date})
	} else if Type == "depense" {
		*g.Depense = append(*g.Depense, Gestion{montant, Type, categorie, date})
	}

	//Insertion des données dans le fichier csv
	err := g.Enregistrement()
	if err != nil {
		return fmt.Errorf("erreur d'enregistrement")
	}

	return nil

}

// Fonction pour l'enregistement dans le fichier csv
func (g *GestionList) Enregistrement() error {
	// Ouverture du fichier
	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return fmt.Errorf("erreur à l'ouverture du fichier")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Parcours et écriture des transactions dans le fichier
	for _, Gestion := range *g.Revenue {
		err := writer.Write([]string{Gestion.Date, Gestion.Type, Gestion.Categorie, fmt.Sprintf("%d", Gestion.Montant)})
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture des transactions")
		}
	}
	for _, Gestion := range *g.Depense {
		err := writer.Write([]string{Gestion.Date, Gestion.Type, Gestion.Categorie, fmt.Sprintf("%d", Gestion.Montant)})
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture des transactions")
		}
	}

	fmt.Println("Transactions enregistrées avec succès !")

	return nil

}

// Affichage du contenu CSV (Transaction)
func (g *GestionList) Affich_Transaction() {
	fmt.Println("	Date	|	Type	|	categorie	|	Montant		")
	if g.Revenue != nil {
		for _, Gestion := range *g.Revenue {
			fmt.Printf(" %s	|	%s	|	%s	|	%d	", Gestion.Date, Gestion.Type, Gestion.Categorie, Gestion.Montant)
		}
	}

	if g.Depense != nil {
		for _, Gestion := range *g.Depense {
			fmt.Printf(" %s	|	%s	|	%s	|	%d	", Gestion.Date, Gestion.Type, Gestion.Categorie, Gestion.Montant)
		}
	}
}

func main() {

	var montant int
	var Type, categorie, date string

	//Liste pour faire le choix de Menu
	fmt.Println("============= Menu de Gestion de Transaction =============")
	fmt.Println("1. Ajout_Transaction")
	fmt.Println("2. Affich_Transaction")
	/*fmt.Println("3. Supprimer_Transaction")
	fmt.Println("4. Modifier_Transaction")*/
	fmt.Println("Veuillez choisir une option : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		// Saisi des Transactions
		fmt.Println("Veuillez saisir le montant : ")
		fmt.Scan(&montant)

		fmt.Println("Veuillez saisir le type de transaction : ")
		fmt.Scan(&Type)

		fmt.Println("Veuillez saisir la catégorie : ")
		fmt.Scan(&categorie)

		fmt.Println("Veuillez saisir la date : ")
		fmt.Scan(&date)

		// Ajout des Transactions
		gestion := GestionList{}
		err := gestion.Ajout_Transaction(montant, Type, categorie, date)
		if err != nil {
			fmt.Println(err)
		}

	case 2:
		// Affichage des Transactions
		gestion := GestionList{}
		gestion.Affich_Transaction()

	default:
		fmt.Println("Votre entré est invalide")

	}
}
