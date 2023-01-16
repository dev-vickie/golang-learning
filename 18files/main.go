package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Let's manipulate them files in golang")
	content := "This is my first golang book" //Have some text content

	file, err := os.Create("./myfirstbook.txt") //create file using the os.Create() method - pass the filename.txt
	if err != nil {
		panic(err)
	}

	//this returns the length of the content text in the file
	length, err := io.WriteString(file, content) //add the created content into the file using io.WriteString() method
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is:", length) //print the length
	defer file.Close()                    //set to close the file once done
	readFile("./myfirstbook.txt")
}
func readFile(filename string) { //this method reads the file in the passed filename
	databyte, err := os.ReadFile(filename) //read the file using the os.ReadFile() method - pass the ./filename.txs
	if err != nil {
		panic(err)
	}
	fmt.Println("The data in the file is : \n", string(databyte)) //since data is read in bytes,use the string() method to convert it to text
}
