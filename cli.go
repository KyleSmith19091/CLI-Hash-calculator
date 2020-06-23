package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main(){
    app := cli.NewApp()
    app.Name = "Hash Calculator"
    app.Usage = "A quick and easy way to calculate a hash of a file or literal value"

    myFlags := []cli.Flag{
        &cli.StringFlag{
            Name: "hash",
        },
        &cli.StringFlag{
            Name: "filename",
        },
        &cli.StringFlag{
            Name: "value",
        },
    }

    app.Commands = []*cli.Command{
        {
            Name: "hash",
            Usage: "Calculates a hash using specifed hasing algorithm to the provided file",
            Flags: myFlags,
            Action: func(c *cli.Context) error{
                // Get the data from the specified file
                var data []byte
                var err error
                // Check if the user provides a file and a value
                if(len(c.String("filename"))>0 && len(c.String("value"))>0){
                    fmt.Println("One argument at a time")
                }else if(len(c.String("filename"))>0){ // Now check if the user has provided a filename
                    data, err = ioutil.ReadFile(c.String("filename"))
                }else if(len(c.String("value")) > 0){ // Now check if the user has provided some literal value
                    err = nil
                    data = []byte(c.String("value"))
                }
                if err != nil{
                    log.Fatal(err)
                    return err
                }

                // Apply specified hashing algorithm
                switch c.String("hash"){
                    case "md5":
                        fmt.Printf("Md5: %x\n\n", md5.Sum(data))
                    case "sha1":
                        fmt.Printf("Sha1: %x\n\n",sha1.Sum(data))
                    case "sha256":
                        fmt.Printf("Sha256: %x\n\n",sha256.Sum256(data))
                    case "sha512":
                        fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
                    default:
                        fmt.Println("Hashing algorithm not found, Please try again")
                }

                // Hash the file and output results
                return nil
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil{
        log.Fatal(err)
    }

}
