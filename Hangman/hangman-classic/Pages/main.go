package main

var words string       //Cette variable va stocker le mot choisi par le programme aléatoirement.
var hidewords []string //Cette variable permet de stocker les lettres de la variable "words" et les remplacer par des underscores(_)
var tab []string       //Cette variable print le mot avec les lettres trouvées par l'utilisateur.
var interaction string //Cette variable permet de stocker la lettre ou le mot rentrer par l'utilisateur.
var UseWords []string  //Cette variable permet d'afficher dans le programme en cours les lettres ou les mots utiliser par l'utilisateur.
var life int = 10      //Cette variable stock le nombre de l'utilisateur qui sera soustrait par 1 ou par 2 si la lettre ou le mot rentrer ne correspond pas.
var lvl string

func main() {
	words = ""
	hidewords = []string{}
	tab = []string{}
	interaction = ""
	UseWords = []string{}
	life = 10
	lvl = ""
	Mot()
	Hide(words)
	Random(tab)
	Letter()
}
