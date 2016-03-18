package util

import "encoding/csv"
import "fmt"
import "os"

func GenerateCSV(fileName string, records [][]string){

        fmt.Println(records)
        csvfile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE,0600)
        if err != nil{
                fmt.Println("Error: ", err)
                return;
        }
        defer csvfile.Close()

        writer := csv.NewWriter(csvfile)
        for _, record := range records {
                  err := writer.Write(record)
                  if err != nil {
                          fmt.Println("Error:", err)
                          return
                  }
        }
        writer.Flush()
}
