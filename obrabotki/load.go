package obarabotki

import (
	"io"
	"strings"
	"bufio"
	"os"
	"encoding/csv"
	"log"
)


// LoadCsvFile загрузка csv файла utf8
func LoadCsvFile(filename string) {
	var f, t []string
	var table string

	csvFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()
	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ';'
	k := 0

	for {
		record, err := r.Read()
		if k == 0 {
			f = make([]string, len(record), len(record)*2)
			copy(f, record[0:])
		}

		if k == 1 {
			t = make([]string, len(record), len(record)*2)
			copy(t, record[0:])
			log.Println(t)
			rf := csv.NewReader(strings.NewReader(t[0]))
			rf.Comma = '.'
			recordf, _ := rf.Read()
			// удалим Маркер последовательности байтов bom
			f[0] = f[0][3:len(f[0])]

			switch f[0] {
			case "ОбъектТЧ":
				table = recordf[2]
			case "Объект":
				table = recordf[1]
			}
			log.Println(table)
		}

		k++
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		switch table {
		case "Должности":
			//fmt.Println(table)

		}

		log.Println(record)

	}

}