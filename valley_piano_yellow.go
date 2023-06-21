package main

import "fmt"

func main() {
	//Define the mission of the 'Beach Cleanup Crew'
	mission := "The mission of the Beach Cleanup Crew is to protect and preserve our fragile coastal environment through concerted beach cleanups and active beach conservation."
	fmt.Println(mission)
	
	//Create a list of items that need to be brought to the beach
	neededItems := []string{"Sunscreen", "Reusable Water Bottles", "Gloves", "Trash Bags", "Bucket", "Refillable Sanitizing Wipes"}
	
	//Print out the list of items that need to be brought to the beach
	fmt.Println("Items needed for the beach cleanup:")
	for i, item := range neededItems {
		fmt.Println(i+1, item)
	}
	
	//Define a function for cleaning up the beach
	cleanupBeach := func() {
		fmt.Println("Beach cleanup has begun!")
		fmt.Println("Pick up all the trash, fill the trash bags, and put them in the nearby bin")
		fmt.Println("Clean up any debris and place it in the appropriate bins")
		fmt.Println("Rake up any seaweed and dispose of it in designated areas")
		fmt.Println("Clean up the beach and restore it to its natural state")
	}
	
	//Invoke the cleanupBeach function
	cleanupBeach()
	
	//Create a list of safety tips
	safetyTips := []string{"Wear sunscreen to protect from sunburns", "Stay hydrated by drinking plenty of water", "Wear closed-toe shoes to protect your feet", "Avoid contact with any hazardous materials", "Be mindful of the wildlife in the area", "Always work in teams of two or more people"}
	
	//Print out the list of safety tips
	fmt.Println("Safety tips for beach cleanup:")
	for i, tip := range safetyTips {
		fmt.Println(i+1, tip)
	}
	
	//Define a function for monitoring the beach
	monitorBeach := func() {
		fmt.Println("Beach monitoring has begun!")
		fmt.Println("Look for any signs of pollution or illegal activity")
		fmt.Println("Keep an eye out for any wildlife or habitats that could be affected")
		fmt.Println("Conduct regular beach patrols to ensure the safety of visitors and the health of the environment")
		fmt.Println("Report any suspicious activity or environmental hazards to the appropriate authorities")
	}
	
	//Invoke the monitorBeach function
	monitorBeach()
	
	//Create a list of beach conservation tips
	conservationTips := []string{"Be mindful of human interactions with wildlife and their habitats", "Use ecologically friendly products", "Reduce, reuse, recycle", "Avoid littering and dumping", "Reduce your use of single-use plastic products", "Be wary of ocean currents and their effects on trash and debris"}
	
	//Print out the list of beach conservation tips
	fmt.Println("Tips for beach conservation:")
	for i, tip := range conservationTips {
		fmt.Println(i+1, tip)
	}
	
	//Define a function for conserving the beach
	conserveBeach := func() {
		fmt.Println("Beach conservation has begun!")
		fmt.Println("Help protect local wildlife by not disturbing habitats or feeding marine life")
		fmt.Println("Be aware of the potential effects of your actions, such as leaving trash behind")
		fmt.Println("Educate yourself and others on proper beach etiquette")
		fmt.Println("Spread the word about beach conservation and the importance of protecting our fragile coastal environment")
	}
	
	//Invoke the conserveBeach function
	conserveBeach()
	
	//Print out a summary of the 'Beach Cleanup Crew' mission
	fmt.Println("The Beach Cleanup Crew is a collective effort to protect and preserve our fragile coastal environment. We accomplish this through beach cleanups, beach monitoring, and beach conservation.")
}