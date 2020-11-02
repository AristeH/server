// Code generated. DO NOT EDIT.

package model

import (
	"encoding/json"
  
 "strings"

  
	"my/server/config"
	"io/ioutil"
	"os"
	"fmt"
)



type PERSON struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
	CODE string `json:"Код"` // CODE
}

type PERSONList struct {
   Recs []PERSON `json:"ФизическиеЛица"`
}

func  (ob PERSON) Create() error{
    sqlstr := `create table PERSON  (ID, NAME, CODE)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob PERSON) Delete() error{
   const sqlstr = `DELETE FROM PERSON  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob PERSON) Save() error{
   sqlstr := "update or insert into  PERSON  (ID, NAME, CODE) "+
   " values (?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME, ob.CODE)
   if err != nil {
     return err
   }
   return nil
}

func (ob PERSON) Read(id string) error{
   const sqlstr = `select * FROM PERSON  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.CODE,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob PERSON) ReadFromJson(file string){
	var recs PERSONList

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recs)

	for i := 0; i < len(recs.Recs); i++ {
	   ob.ID = recs.Recs[i].ID
	   ob.NAME = recs.Recs[i].NAME
	   ob.CODE = recs.Recs[i].CODE
	   ob.Save()
	}

}

func  (ob PERSON)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "Физические лица",
		Stroki: []arrayFieldSection{
			{
	  		Fields: []FieldSection{
					{
						Name:     "Ссылка",
						Value:    ob.ID,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Наименование",
						Value:    ob.NAME,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Код",
						Value:    ob.CODE,
						Buttons: "",
					},
				},
			},

		
		Buttons: []Button{
			{
				Name:       "Войти",
				Parameters: "login",
				Image:      "",
			},
			{
				Name:       "Отмена",
				Parameters: "cancel",
				Image:      "",
			},
		},
	}

	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return string(output)

}

func  (ob PERSON)  FormSpisok() string{
ret := ""



 
 
return ret
}