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



type OBLAST struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
}

type OBLASTList struct {
   Recs []OBLAST `json:""`
}

func  (ob OBLAST) Create() error{
    sqlstr := `create table OBLAST  (ID, NAME)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob OBLAST) Delete() error{
   const sqlstr = `DELETE FROM OBLAST  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob OBLAST) Save() error{
   sqlstr := "update or insert into  OBLAST  (ID, NAME) "+
   " values (?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME)
   if err != nil {
     return err
   }
   return nil
}

func (ob OBLAST) Read(id string) error{
   const sqlstr = `select * FROM OBLAST  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob OBLAST) ReadFromJson(file string){
	var recs OBLASTList

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
	   ob.Save()
	}

}

func  (ob OBLAST)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "",
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

func  (ob OBLAST)  FormSpisok() string{
ret := ""



 
 
return ret
}