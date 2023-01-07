package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Cette fonction nous permet de choisir un mot aléatoirement dans le fichier words.txt.
func Mot() {
	fmt.Println("Bonjour, mon nom est GlaDoS, je suis une intelligence artificielle.")
	time.Sleep(2 * time.Second)
	fmt.Println("Aujourd'hui nous allons tester vos compétences, vous êtes prêt ? C'est partie !")
	time.Sleep(2 * time.Second)
	var interaction int
	var a int
	for a == 0 {
		fmt.Println("Choississez un niveau : \n -1 (niveau facile) \n -2 (niveau moyen) \n -3 (niveau difficile) ")
		fmt.Scanf("%d\n", &interaction)
		switch interaction {
		case 1:
			lvl = "../Annexe/words.txt"
			a++
		case 2:
			lvl = "../Annexe/words2.txt"
			a++
		case 3:
			lvl = "../Annexe/words3.txt"
			a++
		default:
			fmt.Println("Je n'ai pas compris votre requête")
		}
	}
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J") //Permet de clear le terminal.
	file, err := ioutil.ReadFile(lvl)
	if err != nil {
		log.Fatal(err)
	}
	answers := string(file)
	table := strings.Split(answers, "\n")
	rand.Seed(time.Now().UnixNano())
	words = (table[rand.Intn(len(table))])
	for i := 0; i < len(words); i++ {
		tab = append(tab, string(words[i]))
	}
}

func Random(tab []string) {
	n := len(words)/2 - 1
	for i := 0; i < n; i++ {
		alea := rand.Intn(len(words))
		hidewords[alea] = tab[alea]
	}
	PrintLetter()
}

//Cette fonction laisse le choix à l'utilisateur de choisir d'écrire une lettre directement ou d'écrire le mot en entier.
func Verifword(interaction string) {
	var Rep []string
	var compteur int
	for i := 0; i < len(interaction); i++ {
		Rep = append(Rep, string(interaction[i]))
		if len(Rep) > len(tab) {
			compteur = 0
		} else if Rep[i] == tab[i] {
			compteur++
		}
	}
	if compteur == len(tab) {
		for i := 0; i < len(tab); i++ {
			hidewords[i] = Rep[i]
		}
		fmt.Println(hidewords)
		fmt.Println("Bien jouer c'était le bon mot.")
		Victory()
	} else {
		fmt.Println("Dommage ce n'est pas le bon mot.")
		if life == 1 {
			life = 0
			fmt.Println("Il vous reste :", life, "vies")
		} else {
			life = life - 2
			fmt.Println("Il vous reste :", life, "vie")
		}
		Position()
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println(hidewords)
	Letter()
}

//Cette fonction permet de remplacer les lettres du mot choisis par des underscore (_).
func Hide(words string) {
	for i := 0; i < len(words); i++ {
		hidewords = append(hidewords, "_ ")
	}
}

//Cette fonction permet de vérifier la lettre écrite et elle affiche les lettres déjà testés.
func Word(number string) {
	var a int
	for i := 0; i < len(UseWords); i++ {
		if number == UseWords[i] {
			a++
		}
	}
	if a == 0 {
		UseWords = append(UseWords, number)
	} else {
		fmt.Println("Vous avez déjà écris cette lettre.")
		fmt.Println("Lettres déjà testées :", UseWords)
		fmt.Println("Il vous reste :", life, "vies")
		fmt.Println("-----------------------------------------------")
		Letter()
	}
	fmt.Println("Lettres déjà testées :", UseWords)
}
