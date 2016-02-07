package main

type Bible interface {
	Source() string
	NameShort() string
	Name() string
	GetPassage(reference string) (*BiblePassage, error)
}

type BiblePassage struct {
	Html         string
	TrackingCode string
	Copyright    string
}
