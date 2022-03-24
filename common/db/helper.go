package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/pMertDogan/picusWeek4/domain/author"
	"github.com/pMertDogan/picusWeek4/domain/book"
)



//Simple helper file to test predefined operations by executing them :)

func MigrateDatabase(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	authorRepository.Migrations()
	bookRepository.Migrations()
}

func DropTables(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) error {
	err := bookRepository.DropTable()
	if err != nil {
		// log.Fatal("drop err 1 " + err.Error())
		return err
	}
	err = authorRepository.DropTable()
	if err != nil {
		// log.Fatal("drop err 2 " + err.Error())
		return err
	}
	fmt.Println("Succesfuly dropped tables")
	os.Exit(0)
	return nil

}

func SoftDeleteTest(bookRepository *book.BookRepository) {
	a, err := bookRepository.GetByID("2")
	if err != nil {
		log.Fatal("err 1 " + err.Error())
	}
	bookRepository.SoftDeleteBook("2")
	b, err2 := bookRepository.GetByID("2")

	if err2 != nil {
		log.Fatal("ITs normal to get this error  cause deleted_ad is not  null " + err2.Error())
	}
	fmt.Println(a)
	fmt.Println(b)
}

func UpdateBookTest(bookRepository *book.BookRepository) (*book.Book, error, error, *book.Book) {
	a, err := bookRepository.GetByID("2")
	err2 := bookRepository.UpdateBookQuantity("2", "763")
	b, _ := bookRepository.GetByID("2")

	if err != nil {
		log.Fatal("Unable read all data " + err.Error())
	}
	if err2 != nil {
		log.Fatal("Unable read all data " + err2.Error())
	}

	fmt.Println(a)
	fmt.Println(b)
	return a, err, err2, b
}

func ReadFilesAndSaveThemToDB(authorRepository *author.AuthorRepository, bookRepository *book.BookRepository) {
	//read files
	// authors := readAuthorsFromFile()
	authors, authorsErr := author.FromFile(os.Getenv("sourceAuthorsJsonLocation"))

	if authorsErr != nil {
		log.Fatal("unable read Authors from file " + authorsErr.Error())
	}

	books, bookErr := book.FromFile(os.Getenv("sourceBooksJsonLocation"))

	if bookErr != nil {
		log.Fatal("unable read Books from file " + bookErr.Error())
	}

	authorRepository.InsertSampleData(authors)
	bookRepository.InsertSampleData(books)

	log.Println("Sample Datas imported, source is JSON File\n \n ")
}


