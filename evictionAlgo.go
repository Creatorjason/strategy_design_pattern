package main


type evictionAlgo interface{
	evict(*cache)
}