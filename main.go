package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item
	for i, val := range database {
		if val.title == title {
			database[i] = edit
			changed = edit
		}
	}
	return changed
}

func DeleteItem(item Item) Item {
	var del Item
	for i, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {
	fmt.Println("Initial databse : ", database)
	a := Item{"a", "first test case"}
	b := Item{"b", "second test case"}
	c := Item{"c", "third test case"}
	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("\nDatabase after adding items: \t\t", database)
	DeleteItem(b)
	fmt.Println("\nDatabase after deleting items: \t\t", database)
	EditItem("c", Item{"x", "new first case"})
	fmt.Println("\nDatabase after editing items: \t\t", database)
	GetByName("x")
}
