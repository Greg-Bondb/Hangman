package main

import (
	"fmt"
	"os"
	"time"
)

//Cette fonction permet de vérifier si le mot donner est complet, il affichera en cas de victore deux phrases.
func Victory() {
	var compt int
	for i := 0; i < len(hidewords); i++ {
		if hidewords[i] == tab[i] {
			compt++
		}
	}
	if compt == len(tab) {
		fmt.Println("Pal mal, vous m'avez battu ^_^")
		time.Sleep(2 * time.Second)
		fmt.Println("On remet en jeu votre titre ?")
		time.Sleep(1 * time.Second)
		var a int
		for a == 0 {
			fmt.Println("-oui (pour relancer la partie) \n-non (pour arreter la partie)")
			var interaction string
			fmt.Scanf("%s\n", &interaction)

			switch interaction {
			case "oui":
				main()
				a++
			case "non":
				os.Exit(0)
				a++
			default:
				fmt.Println("Je n'ai pas compris votre requête")
			}
		}
	}
}
