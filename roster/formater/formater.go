package formater

import (
	"fmt"

	"github.com/stephenwu2020/fabric-cases/roster/datatype"
)

func PrintPerson(index int, person *datatype.Person) {
	fmt.Printf("(%d) %s: \n", index, person.Name)
	fmt.Println("  Id\t\t:", person.Id)
	fmt.Println("  Name\t\t:", person.Name)
	fmt.Println("  Age\t\t:", person.Age)
	fmt.Println("  Gender\t:", person.Gender)
	year, month, day := person.Birth.Date()
	fmt.Printf("  Birth\t\t: %04d-%02d-%02d\n", year, month, day)
	fmt.Println("  BirthPlace\t:", person.BirthPlace)
	fmt.Println("  GroupTags\t:", person.GroupTags)
	fmt.Println("  HistroyId\t:", person.HistroyId)
}

func PrintRecord(index int, record *datatype.Record) {
	fmt.Printf("(%d) Record: \n", index)
	fmt.Println("  Id\t\t:", record.Id)
	fmt.Println("  Content\t:", record.Content)
	fmt.Println("  Comment\t:", record.Comment)
	year, month, day := record.Time.Date()
	fmt.Printf("  Time\t\t: %d-%02d-%02d\n", year, month, day)
}
