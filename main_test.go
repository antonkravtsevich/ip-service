package main

import "testing"

func TestIsArrayContainValue(t *testing.T){
	inputArray := []string{"1", "2", "3", "4", "5"}

	exists := isArrayContainValue(inputArray, "1")
	if !exists {
		t.Errorf("isArrayContainValue failed, expected %v, got %v", true, exists)
	}

	exists = isArrayContainValue(inputArray, "6")
	if exists {
		t.Errorf("isArrayContainValue failed, expected %v, got %v", false, exists)
	}
}

func TestStartIPsCounter(t *testing.T){
	commandsChannel := startIPsCounter()
	responseChannel := make(chan int)

	// check amount of unique IPs before anything will be added
	command := Command{GetCount, "", responseChannel}
	commandsChannel <- command
	count := <- responseChannel

	if count != 0 {
		t.Errorf("startIPsCounter failed, expected %v, got %v", 0, count)
	}

	// add new unique IP
	command = Command{AddIp, "1.2.3.4", responseChannel}
	commandsChannel <- command
	<- responseChannel

	// check if unique IPs count is increased
	command = Command{GetCount, "", responseChannel}
	commandsChannel <- command
	count = <- responseChannel
	if count != 1 {
		t.Errorf("startIPsCounter failed, expected %v, got %v", 1, count)
	}

	// add same IP again
	command = Command{AddIp, "1.2.3.4", responseChannel}
	commandsChannel <- command
	<- responseChannel

	// unique IPs count should not be increased
	command = Command{GetCount, "", responseChannel}
	commandsChannel <- command
	count = <- responseChannel

	if count != 1 {
		t.Errorf("startIPsCounter failed, expected %v, got %v", 1, count)
	}
}