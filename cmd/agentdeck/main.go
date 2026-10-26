package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Avikshit28/AgentDeck/internal/events"
)

// fakeAgent sends 5 text events, pausing between each.
func fakeAgent(id events.AgentID, count int,  delay time.Duration, bus *events.Bus, wg*sync.WaitGroup){
	defer wg.Done()
	for i:= 1; i<=count; i++{
		time.Sleep(delay)
		bus.Publish(events.TextDelta{Meta: events.NewMeta(id), Text:  fmt.Sprintf("line %d", i)})
	}
}
func main(){
	bus := events.NewBus(0)
	var wg sync.WaitGroup
	wg.Add(4)
	go fakeAgent("a1", 3, 100*time.Millisecond, bus, &wg)
	go fakeAgent("a2", 5, 150*time.Millisecond, bus, &wg)
	go fakeAgent("a3", 2, 250*time.Millisecond, bus, &wg)
	go fakeAgent("a4", 4, 200*time.Millisecond, bus, &wg)
	go func(){
		wg.Wait()
		bus.Close()
	}()
	for {
		select{
		case e := <-bus.Events():
			if td, ok := e.(events.TextDelta); ok{
				fmt.Printf("[%s] %s\n", td.AgentID(), td.Text)
			}
		case <-bus.Done():
			fmt.Println("all agents finished, bus closed")
			return
		}
	}
}