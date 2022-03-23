package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	postgres "github.com/pMertDogan/picusWeek4/common/db"
	"github.com/pMertDogan/picusWeek4/domain/author"
	"github.com/pMertDogan/picusWeek4/domain/book"
)

var usageInformation = ` 

****
Just fill the .env file with your database informations. 
If you get an error message please check parameters inside the  .env file
****

`

//init app config
func init() {
	//Load env
	err := godotenv.Load()
	if err != nil {
		printUsageAndExit("Error loading .env file. \n ERROR : " + err.Error())
	}

}

func main() {
	//make database connection
	db, err := postgres.ConnectPostgresDB()
	if err != nil {
		log.Fatal("Postgres cannot init: \n ", err)
	}
	log.Println("Postgres connected!!")

	//create repositorys for each domain
	authorRepository := author.NewAuthorRepository(db)
	bookRepository := book.NewBookRepository(db)

	//migrate database struct changes.
	authorRepository.Migrations()
	bookRepository.Migrations()

	//save source data readed by json file to SQL
	readFilesAndSaveThemToDB(authorRepository, bookRepository)


	// SUPPORTED Methods
	b, err := bookRepository.FindByName("Lord")
	// // b, err := bookRepository.GetByID("2")
	// b, err := bookRepository.GetBooksWithAuthors()
	a, err := authorRepository.FindByName("J.R.R")

	if err != nil {
		log.Fatal("Unable read all data " + err.Error())
	}

	//To test String override
	//if its return BOOKS not Book
	// for _,v := range b{
	// 	fmt.Print(v)
	// }

	fmt.Println(a)
	fmt.Println(b)

}

func readFilesAndSaveThemToDB(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	//read files
	// authors := readAuthorsFromFile()
	authors, authorsErr := author.FromFile(os.Getenv("sourceAuthorsJsonLocation"))

	if authorsErr != nil {
		printUsageAndExit("unable read Authors from file " + authorsErr.Error())
	}


	books, bookErr := book.FromFile(os.Getenv("sourceBooksJsonLocation"))

	if bookErr != nil {
		printUsageAndExit("unable read Books from file " + bookErr.Error())
	}

	authorRepository.InsertSampleData(authors)
	bookRepository.InsertSampleData(books)

	log.Println("Sample Datas imported, source is JSON File\n \n ")
}

//print optional meesage and exit with error code (1)
func printUsageAndExit(optionalText string) {
	fmt.Println(optionalText)
	fmt.Println(usageInformation)
	os.Exit(1)

}
