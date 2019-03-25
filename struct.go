package main

import "fmt"

// Pet lllll
type Pet interface {
	setName(name string)
	getName() string
	category() string
}

type dog struct {
	name string
}

func (dog dog) setName(name string) {
	dog.name = name
}

func (dog dog) getName() string {
	return dog.name
}

func (dog *dog) category() string {
	return "dog"
}

// type animalCategory struct {
// 	kingdom string
// 	phylum  string
// 	class   string
// }

// func (ac animalCategory) String() string {
// 	return fmt.Sprintf("%s %s %s", ac.kingdom, ac.phylum, ac.class)
// }

// type animal struct {
// 	animaName string
// 	animalCategory
// }

// func (an animal) String() string {
// 	return fmt.Sprintf("%s (category:%s)", an.animaName, an.animalCategory)
// }

func main() {
	// ac := animalCategory{
	// 	kingdom: "kingdom",
	// 	phylum:  "phylum",
	// 	class:   "class",
	// }

	// an := animal{
	// 	animaName:      "American Shorthair",
	// 	animalCategory: ac,
	// }

	// fmt.Println(an)

	d := dog{
		name: "gouzi",
	}
	var dp Pet = &d

	var pig Pet
	fmt.Println(pig)

	_, ok := interface{}(d).(Pet)
	fmt.Println(ok)
	_, ok = interface{}(dp).(Pet)
	fmt.Println(ok)

	var dog1 *dog
	fmt.Println("The first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("The second dog is nil. [wrap1]")
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
	}

}
