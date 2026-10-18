package main
import (
	"fmt"
	"github.com/Avikshit28/AgentDeck/internal/events"
)
func main () {
	s := events.StatusRunning
	fmt.Println("agent a1 is", s)
	fmt.Println("finished?", s.Terminal())

	s = events.StatusFailed
	fmt.Println("agent a1 is", s)
	fmt.Println("finished?", s.Terminal())
}