package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Ajout des Transactions
func (g *GestionList) Ajout_Transaction(date string, Type string, categorie string, montant int) error {
	// Initialisation des listes Dans le tableau
	if g.Revenue == nil {
		g.Revenue = &[]Gestion{}
	}
	if g.Depense == nil {
		g.Depense = &[]Gestion{}
	}
	// Gestion d'erreur
	if Type != "Revenue" && Type != "Depense" {
		return fmt.Errorf("type de transaction Disponible : Revenue ou Depense")
	}
	//classement du type
	if Type == "Revenue" { // Correction ici: "Revenue" au lieu de "revenue"
		*g.Revenue = append(*g.Revenue, Gestion{date, Type, categorie, montant})
	} else if Type == "Depense" { // Pour être cohérent, utilisez "Depense" avec majuscule
		*g.Depense = append(*g.Depense, Gestion{date, Type, categorie, montant})
	}
	//Insertion des données dans le fichier csv
	err := g.Enregistrement()
	if err != nil {
		return fmt.Errorf("erreur d'enregistrement: %v", err)
	}
	return nil
}

// Fonction pour l'enregistement dans le fichier csv
func (g *GestionList) Enregistrement() error {
	// Ouvrir le fichier en mode création/troncation (pas append)
	file, err := os.Create("data.csv")
	if err != nil {
		return fmt.Errorf("erreur à l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Écrire l'en-tête
	err = writer.Write([]string{"Date", "Type", "Categorie", "Montant"})
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture de l'en-tête: %v", err)
	}

	// Parcours et écriture des transactions dans le fichier
	if g.Revenue != nil {
		for _, Gestion := range *g.Revenue {
			err := writer.Write([]string{Gestion.Date, Gestion.Type, Gestion.Categorie, fmt.Sprintf("%d", Gestion.Montant)})
			if err != nil {
				return fmt.Errorf("erreur lors de l'écriture des transactions de revenu: %v", err)
			}
		}
	}

	if g.Depense != nil {
		for _, Gestion := range *g.Depense {
			err := writer.Write([]string{Gestion.Date, Gestion.Type, Gestion.Categorie, fmt.Sprintf("%d", Gestion.Montant)})
			if err != nil {
				return fmt.Errorf("erreur lors de l'écriture des transactions de dépense: %v", err)
			}
		}
	}

	fmt.Println("Transactions enregistrées avec succès !")
	return nil
}

// Chargement des données depuis le fichier CSV
func (g *GestionList) ChargerDonnees() error {
	// Initialiser les listes
	g.Revenue = &[]Gestion{}
	g.Depense = &[]Gestion{}

	// Vérifier si le fichier existe
	_, err := os.Stat("data.csv")
	if os.IsNotExist(err) {
		// Le fichier n'existe pas encore, ce n'est pas une erreur
		return nil
	}

	// Ouvrir le fichier
	file, err := os.Open("data.csv")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	// Lire le fichier CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier CSV: %v", err)
	}

	// Ignorer l'en-tête
	if len(records) > 0 {
		records = records[1:]
	}

	// Parcourir les enregistrements et les ajouter aux listes appropriées
	for _, record := range records {
		if len(record) < 4 {
			continue // Ignorer les lignes mal formées
		}

		montant := 0
		fmt.Sscanf(record[3], "%d", &montant)

		gestion := Gestion{
			Date:      record[0],
			Type:      record[1],
			Categorie: record[2],
			Montant:   montant,
		}

		if gestion.Type == "Revenue" {
			*g.Revenue = append(*g.Revenue, gestion)
		} else if gestion.Type == "Depense" {
			*g.Depense = append(*g.Depense, gestion)
		}
	}

	return nil
}

// Affichage du contenu CSV (Transaction)
func (g *GestionList) Affich_Transaction() {
	fmt.Println("=== Liste des Transactions ===")
	fmt.Println(" Date 		| Type 		| Categorie 	| Montant 	")
	fmt.Println("--------------------------------------")

	if g.Revenue != nil {
		for _, Gestion := range *g.Revenue {
			fmt.Printf(" %s | %s | %s | %d\n", Gestion.Date, Gestion.Type, Gestion.Categorie, Gestion.Montant)
		}
	}

	if g.Depense != nil {
		for _, Gestion := range *g.Depense {
			fmt.Printf(" %s | %s | %s | %d\n", Gestion.Date, Gestion.Type, Gestion.Categorie, Gestion.Montant)
		}
	}
}

func main() {
	var montant int
	var Type, categorie, date string

	// Créer une instance de GestionList et charger les données existantes
	gestion := GestionList{}
	err := gestion.ChargerDonnees()
	if err != nil {
		fmt.Printf("Erreur lors du chargement des données: %v\n", err)
	}

	//Liste pour faire le choix de Menu
	fmt.Println("============= Menu de Gestion de Transaction =============")
	fmt.Println("1. Ajoutez une Transaction")
	fmt.Println("2. Affichez une Transaction")
	/*fmt.Println("3. Supprimer_Transaction")
	fmt.Println("4. Modifier_Transaction")*/
	fmt.Println("Veuillez choisir une option : ")
	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		// Saisi des Transactions
		fmt.Println("Veuillez saisir la date (format JJ/MM/AAAA) : ")
		fmt.Scan(&date)
		fmt.Println("Veuillez saisir le type de transaction (Revenue/Depense) : ")
		fmt.Scan(&Type)
		fmt.Println("Veuillez saisir la catégorie : ")
		fmt.Scan(&categorie)
		fmt.Println("Veuillez saisir le montant : ")
		fmt.Scan(&montant)

		// Ajout des Transactions
		err := gestion.Ajout_Transaction(date, Type, categorie, montant)
		if err != nil {
			fmt.Println("Erreur:", err)
		}
	case 2:
		// Affichage des Transactions
		gestion.Affich_Transaction()
	default:
		fmt.Println("Votre entrée est invalide")
	}
}
