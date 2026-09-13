package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(te *testing.T)  {
	var cachin *Cache = NewCache(1 * time.Second)
	cachin.Add("hola", []byte("ismael Manzanero"))
	_, ok := cachin.Get("hola")
	if !ok {
		te.Errorf("Expected to get value, got false instead")
	}	
}
func TestCleaningLoop(te *testing.T)  {
	var cachin *Cache = NewCache(1 * time.Second)
	cachin.Add("hola", []byte("ismael Manzanero"))

	time.Sleep(2 * time.Second)
	_, ok := cachin.Get("hola")
	if ok {
		te.Errorf("Expected to get false, got value instead")
	}	
}