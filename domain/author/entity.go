package author

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/gorm"
)


type Authors []Author

type Author struct {
	gorm.Model
	AuthorID string  `gorm:"foreignKey:AuthorID;references:AuthorID"` //`gorm:"type:varchar(100);column:ID_Author2"` SQL type and column name can be selected
	Name     string 
}

//Change table name
// func (Author) TableName() string {
// 	return "Author"
// }

//override default string method fmt.println get this value 
func (a *Author) String() string {
	return fmt.Sprintf("ID : %s, Name : %s , CreatedAt %s", a.AuthorID, a.Name,a.CreatedAt.Format("2006-01-02 15:04:05"),
)
}

func FromFile(locationOfFile string) (Authors,error){
	
	//read our data from dumy json file
	dat, err := os.ReadFile(locationOfFile)
	if err != nil {
		// check json file :)
		return Authors{},err
	}

	//convert json to to authors
	authors, err := UnmarshalAuthors(dat)

	if err != nil {
		fmt.Print(string(dat))

		// printUsageAndExit("Unable convert to struct")
	}
	return authors , err



}

//just paste authors.json to quicktype or any other tool
// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authors, err := UnmarshalAuthors(bytes)
//    bytes, err = authors.Marshal()





//fromJson
func UnmarshalAuthors(data []byte) (Authors, error) {
	var r Authors
	err := json.Unmarshal(data, &r)
	return r, err
}
//toJson
func (r *Authors) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


