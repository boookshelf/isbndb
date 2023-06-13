ISBNdb API Wrapper
==================

This is an unofficial Go wrapper for the [ISBNdb API](https://isbndb.com/apidocs/v2).

## Installation

`go get github.com/boookshelf/isbndb`

## Authentication

You will need an API key from ISBNdb which can be found [here](https://isbndb.com/isbn-database).
Once you have an API key, you will need to set it as the environment variable `ISBNDB_API_KEY`.

## Usage

````Go
client := isbndb.New(http.DefaultClient)
book, err := client.GetBook(context.TODO(), "1501154656")
if err != nil {
    log.Fatal(err)
}

fmt.Println(book)
````



