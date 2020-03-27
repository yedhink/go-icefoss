package main
import "fmt"

type Tree struct{
	root *Node
}

type Node struct {
	left *Node
	nodeValue int
	right *Node
}

func (t *Tree) insert(chanValue chan int,outChan chan int,treeNum int,lastElement int){
	for {
		value := <-chanValue
		// fmt.Println("value in insert of Tree : ",value)
		if t.root == nil {
			t.root = &Node{nodeValue : value}
		} else {
			t.root.insert(value)
		}
		if value == lastElement {
			outChan <- treeNum
		}
	}
}

func (n *Node) insert(value int){
	if value <= n.nodeValue{
		if n.left == nil {
			n.left = &Node{nodeValue : value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{nodeValue : value}
		} else {
			n.right.insert(value)
		}
	}
}

func inOrderTraversal(n *Node){
	if n==nil{
		return
	}else{
		inOrderTraversal(n.left)
		fmt.Printf("%v,",n.nodeValue)
		inOrderTraversal(n.right)
	}
}

func main() {
	elements := [][]int{
		{1, 1, 2, 3},
		{10, 5, 1, 7, 40, 50},
	}

	trees := []Tree {
		Tree{},
		Tree{},
	}

	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	done := make(chan int)

	for i,row := range elements{
		go trees[i].insert(chans[i],done,i,elements[i][len(elements[i])-1])
		for _,item := range row{
			// fmt.Println("item value in main : ",item)
			chans[i]<-item
		}
		fmt.Printf("Inorder of Tree %v\n------\n",i)
		inOrderTraversal(trees[<-done].root)
		fmt.Println("\n")
	}
	close(done)
	close(chans[0])
	close(chans[1])
}
