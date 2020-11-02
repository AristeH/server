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



type USERS struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
	CODE string `json:"Код"` // CODE
	PASSWORD string `json:""` // PASSWORD
}

type USERSList struct {
   Recs []USERS `json:"Пользователи"`
}

func  (ob USERS) Create() error{
    sqlstr := `create table USERS  (ID, NAME, CODE, PASSWORD)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob USERS) Delete() error{
   const sqlstr = `DELETE FROM USERS  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob USERS) Save() error{
   sqlstr := "update or insert into  USERS  (ID, NAME, CODE, PASSWORD) "+
   " values (?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME, ob.CODE, ob.PASSWORD)
   if err != nil {
     return err
   }
   return nil
}

func (ob USERS) Read(id string) error{
   const sqlstr = `select * FROM USERS  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.CODE, &ob.PASSWORD,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob USERS) ReadFromJson(file string){
	var recs USERSList

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
	   ob.PASSWORD = recs.Recs[i].PASSWORD
	   ob.Save()
	}

}

func  (ob USERS)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "Пользователи",
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
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.PASSWORD,
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

func  (ob USERS)  FormSpisok() string{
ret := ""



 
 
return ret
}