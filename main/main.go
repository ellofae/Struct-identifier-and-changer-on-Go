package main

import (
	"fmt"
	"identifier"
)

type MyType struct {
	IntField    int
	StrField    string
	IntPtrField *int
	StrPtrField *string
}

func main() {
	aVar := 5
	bVar := 46
	temp := MyType{IntField: 10, StrField: "some", IntPtrField: &aVar, StrPtrField: nil}

	var tempPtr *int = &bVar

	identifier.StructIdentifier(temp)
	identifier.StructIdentifier(struct {
		TestIntField    int
		TestStringField string
	}{2, "other text"})
	identifier.StructIdentifier("hello")
	identifier.StructIdentifier(&tempPtr)
	identifier.StructIdentifier(13.4)

	fmt.Println("\nStruct before the change:")
	identifier.StructIdentifier(temp)
	identifier.StructFieldChanger(&temp, "IntField", 35)
	identifier.StructFieldChanger(&temp, "StrField", "another text")

	fmt.Println("\nStruct after the change:")
	identifier.StructIdentifier(temp)

	fmt.Printf("\nStruct's ptr value before: %v\n", *temp.IntPtrField)
	identifier.StructFieldChanger(&temp, "IntPtrField", tempPtr)
	identifier.StructIdentifier(temp)
	fmt.Printf("Struct's int ptr new value: %v\n", *temp.IntPtrField)
}
