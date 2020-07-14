package main

import "testing"

func TestInsert(t *testing.T){
	node := Node{' ', []*Node{}}
	node.insert('1')
	if !(node.Children[0].Value == '1'){
		t.Errorf("Children wasn't added correctly")
	}
}

func TestGetChildWithValue(t *testing.T){
	node := Node{' ', []*Node{}}
	node.insert('1')
	node.insert('2')

	child := node.getChildWithValue('2')

	if child.Value != '2'{
		t.Errorf("Child value is not correct; expected: %v, got: %v", '2', child.Value)
	}
}

func TestGetChildWithValueNegative(t *testing.T){
	node := Node{' ', []*Node{}}
	node.insert('1')
	node.insert('2')

	child := node.getChildWithValue('3')

	if child != nil{
		t.Errorf("getChildWithValue is not correct; expected: %v, got: %v", nil, child)
	}
}

func TestInsertSingleBranch(t *testing.T){
	node := Node{' ', []*Node{}}
	values := []rune{'1', '2'}
	node.insertBranch(values)

	child1 := node.Children[0]
	child2 := child1.Children[0]

	if child1.Value != '1'{
		t.Errorf("Child 1 value is not correct; expected: %v, got: %v", '1', child1.Value)
	}
	if child2.Value != '2'{
		t.Errorf("Child 2 value is not correct; expected: %v, got: %v", '2', child2.Value)
	}
}

func TestGetChildInBranch(t *testing.T){
	node := &Node{' ', []*Node{}}
	values := []rune{'1', '2', '3', '4'}
	node.insertBranch(values)

	for _, value := range values{
		node = node.getChildWithValue(value)
		if node == nil {
			t.Errorf("Can't get child with value %v", value)
		}
	}
}

func TestInsertMultipleBranches(t* testing.T){
	node := Node{' ', []*Node{}}
	values := []rune{'1', '2', '3', '4'}
	node.insertBranch(values)
	values = []rune{'1', '2', '5', '6'}
	node.insertBranch(values)

	child := node.Children[0].Children[0]

	if len(child.Children) != 2{
		t.Errorf("Element have incorrect number of children, expected %v, got %v", 2, len(child.Children))
	}
	if child.Children[0].Value != '3'{
		t.Errorf("Child have incorrect value; expected %v, got %v", '3', child.Children[0].Value)
	}
	if child.Children[1].Value != '5'{
		t.Errorf("Child have incorrect value; expected %v, got %v", '3', child.Children[1].Value)
	}
}

func TestIsBranchExists(t* testing.T){
	node := Node{' ', []*Node{}}
	values := []rune{'1', '2', '3', '4'}
	node.insertBranch(values)

	branchExists := node.isBranchExists(values)
	if !branchExists{
		t.Errorf("Required branch is not exists")
	}
}

func TestIsBranchExistsNegative(t* testing.T){
	node := Node{' ', []*Node{}}
	values := []rune{'1', '2', '3', '4'}
	node.insertBranch(values)
	negative_values := []rune{'1', '2', '4', '5'}

	branchExists := node.isBranchExists(negative_values)
	if branchExists{
		t.Errorf("Required branch should not exists")
	}
}