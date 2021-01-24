package main

import "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, o := range i.observerList {
		o.update(i.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	for i, observer := range observerList {
		if observer.getID() == observerToRemove.getID() {
			observerList[len(observerList)-1], observerList[i] = observerList[i], observerList[len(observerList)-1]
			return observerList[:len(observerList)-1]
		}
	}
	return observerList
}
