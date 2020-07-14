// rune tree implementation
package main

// main structure - node of the tree
type Node struct {
	Value rune
	Children []*Node
}

// insert child with value in node
func (currentNode *Node) insert(value rune) *Node{
	children := []*Node{}
	newNode := Node{value, children} 
	currentNode.Children = append(currentNode.Children, &newNode)

	return &newNode
}

// get child with required value 
// return nil in case if there is no such child
func (currentNode *Node) getChildWithValue(value rune) *Node{
	for _, child := range currentNode.Children{
		if child.Value == value {
			return child
		}
	}
	return nil
}

// check if there is a specific set of values stored as a branch in a tree
func (currentNode *Node) isBranchExists(values []rune) bool{
	node := currentNode
	for _, currentValue := range values {
		node = node.getChildWithValue(currentValue)
		if node == nil {
			return false
		}
	}
	return true
}

// insert set of values as a branch in a tree
func (currentNode *Node) insertBranch(values []rune){
	// step 1 - find a disjoint node (node, that have required value, but 
	// does not have child with next reuqired value)
	notFoundValueIndex := 0
	for index, currentValue := range values {
		child := currentNode.getChildWithValue(currentValue)
		if child == nil {
			notFoundValueIndex = index
			break
		} else {
			currentNode = child
		}
	}
	
	// cut existing values
	notPresentedValues := values[notFoundValueIndex:]

	// adding childs while there is values left
	for _, value := range notPresentedValues{
		currentNode = currentNode.insert(value)
	}
}