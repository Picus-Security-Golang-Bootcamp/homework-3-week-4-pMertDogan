// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    book, err := UnmarshalBook(bytes)
//    bytes, err = book.Marshal()

package book

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pMertDogan/picusWeek4/domain/author"
	"gorm.io/gorm"
)

type Books []Book

type Book struct {
	gorm.Model
	ID            string
	AuthorID      string
	BookName      string
	NumberOfPages int64
	StockCount    int64
	Price         int64
	Isbn          string
	StockCode     string
	//Authors
	//type Authors []Author
	Author author.Author `gorm:"foreignKey:AuthorID;references:AuthorID"`
}

// fromJson
func UnmarshalBook(data []byte) (Books, error) {
	var r Books
	err := json.Unmarshal(data, &r)
	return r, err
}

//toJson
func (r *Books) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

//override fmt.print() presensation
func (b Book) String() string {
	return fmt.Sprintf("ID : %s, Name : %s Price %d, StockCount %d,AuthorID %s,  CreatedAt %s, DeletedAt %s  *** Author INFO : Name  %s \n ", b.ID, b.BookName, b.Price, b.StockCount, b.AuthorID,b.CreatedAt.Format("2006-01-02 15:04:05"), b.DeletedAt.Time.Format("2006-01-02 15:04:05"),b.Author.Name)
}

//read from file and return
func FromFile(locationOfFile string) (Books, error) {

	//read our data from dumy json file
	dat, err := os.ReadFile(locationOfFile)
	if err != nil {
		// check json file :)
		return Books{}, err
	}

	//convert json to to authors
	authors, err := UnmarshalBook(dat)

	if err != nil {
		fmt.Print(string(dat))

		// printUsageAndExit("Unable convert to struct")
	}
	return authors, err

}
