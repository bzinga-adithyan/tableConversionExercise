package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	csvFile, err := os.Open("inpTable.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(csvLines[1][0])

	err1 := godotenv.Load(".env")

	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err2 := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err2 != nil {
		log.Fatal("error connecting to the database: ", err2)
	}
	fmt.Println("Database bzinga opened and ready.")
	//defer db.Close()

	sqlStatement := `insert into inputTable (refid,categ,prodtitle,descp,mrp,basep) values ($1,$2,$3,$4,$5,$6)`
	for i := 1; i < len(csvLines); i++ {
		id, er := strconv.Atoi(csvLines[i][0])
		if er != nil {
			log.Fatal(er)
		}

		cate := csvLines[i][1]
		prod := csvLines[i][2]
		desc := csvLines[i][3]
		mr := csvLines[i][4]
		mr = strings.ReplaceAll(mr, " ", "")
		mrp, er := strconv.Atoi(mr)
		if er != nil {
			log.Fatal(er)
		}
		// mr := csvLines[i][4]
		bp := csvLines[i][5]

		_, err3 := db.Exec(sqlStatement, id, cate, prod, desc, mrp, bp)
		if err3 != nil {
			log.Fatalf("error in db.exec of sql statement")
		}
	}

}
