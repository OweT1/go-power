package main

// Constants
var WEBSITES = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	// "http://fake-url.com", // This will make the checking take longer, since the website does not even exist
}

// Main
func main() {
	linearCheck(WEBSITES)
	concurrentCheck(WEBSITES)
	concurrentCheckChannel(WEBSITES)
}