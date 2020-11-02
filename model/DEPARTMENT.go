// Code generated. DO NOT EDIT.

package model

import (
	"encoding/json"
  
  "strconv"

  
 "strings"

  
	"my/server/config"
	"io/ioutil"
	"os"
	"fmt"
)



type DEPARTMENT struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
	LEVEL int `json:"Уровень"` // LEVEL
	ID_PARENT string `json:"Родитель"` // ID_PARENT
}

type DEPARTMENTList struct {
   Recs []DEPARTMENT `json:"подразделения"`
}

func  (ob DEPARTMENT) Create() error{
    sqlstr := `create table DEPARTMENT  (ID, NAME, LEVEL, ID_PARENT)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob DEPARTMENT) Delete() error{
   const sqlstr = `DELETE FROM DEPARTMENT  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob DEPARTMENT) Save() error{
   sqlstr := "update or insert into  DEPARTMENT  (ID, NAME, LEVEL, ID_PARENT) "+
   " values (?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME, ob.LEVEL, ob.ID_PARENT)
   if err != nil {
     return err
   }
   return nil
}

func (ob DEPARTMENT) Read(id string) error{
   const sqlstr = `select * FROM DEPARTMENT  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.LEVEL, &ob.ID_PARENT,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob DEPARTMENT) ReadFromJson(file string){
	var recs DEPARTMENTList

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
	   ob.LEVEL = recs.Recs[i].LEVEL
	   ob.ID_PARENT = recs.Recs[i].ID_PARENT
	   ob.Save()
	}

}

func  (ob DEPARTMENT)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "Подразделения",
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
						Name:     "Уровень",
						Value:    ob.LEVEL,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Папка",
						Value:    ob.ID_PARENT,
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

func  (ob DEPARTMENT)  FormSpisok() string{
ret := ""



 
 
return ret
}