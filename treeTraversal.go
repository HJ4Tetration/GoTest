package main

import "fmt"

type node struct {
	left  *node
	right *node
	value interface{}
}

type sharedValue interface {
	valueUpdate()
	getValue() interface{}
}

type sharedValueInt int

func (sV *sharedValueInt) getValue() interface{} {
	return *sV
}
func (sV *sharedValueInt) valueUpdate() {
	*sV++
}

func treeValueGenerator(sv *sharedValue, Rc chan chan<- interface{}) {
	for {
		c := <-Rc
		c <- (*sv).getValue()
		(*sv).valueUpdate()
	}
}

func createNodeGo(level int, Rc chan chan<- interface{}, rC chan<- *node) {
	tmpC := make(chan interface{})
	Rc <- tmpC
	myNode := &node{nil, nil, <-tmpC}
	if level > 0 {
		rCr := make(chan *node)
		go createNodeGo(level-1, Rc, rCr)
		myNode.left = createNode(level-1, Rc)
		myNode.right = <-rCr
	}
	rC <- myNode
}

func createNode(level int, Rc chan chan<- interface{}) *node {
	tmpC := make(chan interface{})
	Rc <- tmpC
	myNode := &node{nil, nil, <-tmpC}
	if level > 0 {
		rCr := make(chan *node)
		go createNodeGo(level-1, Rc, rCr)
		myNode.left = createNode(level-1, Rc)
		myNode.right = <-rCr
	}
	return myNode
}

type output struct {
	v interface{}
	l interface{}
	r interface{}
}

func traverseTree(cNode *node, c chan<- output) {
	cOutput := output{v: cNode.value}
	cl, cr := make(chan output), make(chan output)
	if cNode.left != nil {
		cOutput.l = cNode.left.value
		go traverseTree(cNode.left, cl)
	} else {
		close(cl)
	}
	if cNode.right != nil {
		cOutput.r = cNode.right.value
		go traverseTree(cNode.right, cr)
	} else {
		close(cr)
	}
	c <- cOutput
	for {
		v1, ok1 := <-cl
		v2, ok2 := <-cr
		if ok1 {
			c <- v1
		}
		if ok2 {
			c <- v2
		}
		if !ok1 && !ok2 {
			break
		}
	}
	close(c)
}

func main() {
	var svi sharedValueInt = 1
	var sv sharedValue = &svi //won't compile: sharedValue=svi
	Rc := make(chan chan<- interface{})
	go treeValueGenerator(&sv, Rc)
	root := createNode(4, Rc)
	outputC := make(chan output)
	go traverseTree(root, outputC)
	for o := range outputC {
		fmt.Printf("%v\n", o)
	}
}
