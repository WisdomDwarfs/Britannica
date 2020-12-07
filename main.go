package main

import (
	Britannica "github.com/WisdomDwarfs/Britannica/BritannicaMathematica"
	 "fmt"
	 
)

var (
	fileOp Britannica.FileOperations = Britannica.NewInstance()
	Alice Britannica.CodexBritannica = Britannica.NewVault()
)

func main()  {

	file, err := fileOp.FileExist("iqbal"); if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(file)
	f, err , hash := Alice.CodexMathematica(file);if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("File:", f.Name())
	for i, _ := range hash {
		fmt.Println("Hash Checksum:", hash[i])
	}
}