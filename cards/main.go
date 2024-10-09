package main

// import "fmt"

func main(){
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // only for new variables, will not work for reasigning
	// card := newCard()

	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()	


	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()

}


// func newCard() string {
// 	return "Five of Diamonds"
// }