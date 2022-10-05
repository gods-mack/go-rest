package main
import (
	"fmt"
)

type user struct {
	name string
	email string
	user_id string
	mobile int
}


type link_node struct {
	data int 
	next *link_node
}

var head *link_node = nil
var tail *link_node = nil



func insert_node (_d int) {

	if head == nil {
		head = &link_node{data:_d, next:nil}
		tail = head
	} else {
		new_node := &link_node{data:_d , next:nil}
		tail.next = new_node
		tail = tail.next
	}
}




func main() {
	
	
	fmt.Println("Hlso")
	var u  = user{name:"Manish", email:"manish@gmails.com", user_id:"manish", mobile:804500871}
	fmt.Println(u)

	u1 := user{}
	u2 := user{"usej", "use@mail.com", "godsja", 70384848332}
	u3 := user{user_id:"godsmack"}

	fmt.Println("u1", u1)
	fmt.Println("u2", u2)
	fmt.Println("u3", u3)

	node   := &u2
	fmt.Println("proint ", (*node).email)

	insert_node(12)
	insert_node(14)
	insert_node(16)
	insert_node(18)

	h := head
	for h!=nil{
		fmt.Print(h.data, " ")
		h  = h.next

	}
}
