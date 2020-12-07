package BritannicaMathematica

import (
	"os"
	"io/ioutil"
	"crypto/sha256"
	"encoding/hex"
)



type FileOperations interface{
	
	// File operation 
	Mkdir(filename string) error
	FileExist(filename string) (os.FileInfo, error)
	
}

type Operations struct{}

func NewInstance() FileOperations  {
	return &Operations{}
}

func (o *Operations) Mkdir(filename string) error  {
	return os.Mkdir(filename+".txt", 0777)	
}

func (o *Operations) FileExist(filename string) (os.FileInfo, error) {
	 status, err := os.Stat(filename+".txt")
	 return status, err
}

type CodexBritannica interface{

	// Codex hold special functions 
	CodexMathematica(filename os.FileInfo)(*os.File, error, []string)

}

type Archieve struct{}

func NewVault() CodexBritannica {
	return &Archieve{}
}



func (a *Archieve) CodexMathematica(filename os.FileInfo)(*os.File, error, []string) {
	
	var file *os.File
	buf := make([]byte, 5)
	
	fileStatus ,err := os.Stat(filename.Name()); if os.IsNotExist(err) {
		return file, err , nil
	}

	file , err = os.Open(fileStatus.Name()); if err != nil {
		return file, err, nil
	}

	defer file.Close()

	content, err := ioutil.ReadFile(file.Name()); if err != nil {
		return file , err, nil
	}

	dynamicHash := make([]string, len(content)/4)

	for  valueIn, _ := range content {
		
		for i := 0; i < len(buf) ; i++ {
			
			buf[i] = content[valueIn]		
		}

		hash := AlienWorld(buf)
		dynamicHash = append(dynamicHash, hash)
	}

	return file, err, dynamicHash
}

func AlienWorld(b []byte)string{
	h := sha256.New()
	h1 := h.Sum(b)
	encode := hex.EncodeToString(h1[:6])
	return encode
}