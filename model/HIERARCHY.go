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



type HIERARCHY struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
	TEXT string `json:""` // TEXT
	WEIGHTLEFT int `json:""` // WEIGHTLEFT
	WEIGHTRIGHT int `json:""` // WEIGHTRIGHT
}

type HIERARCHYList struct {
   Recs []HIERARCHY `json:""`
}

func  (ob HIERARCHY) Create() error{
    sqlstr := `create table HIERARCHY  (ID, NAME, TEXT, WEIGHTLEFT, WEIGHTRIGHT)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob HIERARCHY) Delete() error{
   const sqlstr = `DELETE FROM HIERARCHY  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob HIERARCHY) Save() error{
   sqlstr := "update or insert into  HIERARCHY  (ID, NAME, TEXT, WEIGHTLEFT, WEIGHTRIGHT) "+
   " values (?, ?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME, ob.TEXT, ob.WEIGHTLEFT, ob.WEIGHTRIGHT)
   if err != nil {
     return err
   }
   return nil
}

func (ob HIERARCHY) Read(id string) error{
   const sqlstr = `select * FROM HIERARCHY  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.TEXT, &ob.WEIGHTLEFT, &ob.WEIGHTRIGHT,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob HIERARCHY) ReadFromJson(file string){
	var recs HIERARCHYList

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
	   ob.TEXT = recs.Recs[i].TEXT
	   ob.WEIGHTLEFT = recs.Recs[i].WEIGHTLEFT
	   ob.WEIGHTRIGHT = recs.Recs[i].WEIGHTRIGHT
	   ob.Save()
	}

}

func  (ob HIERARCHY)  TmplElem(id string) string{
   
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
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.TEXT,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.WEIGHTLEFT,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.WEIGHTRIGHT,
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

func  (ob HIERARCHY)  FormSpisok() string{
ret := ""



 
 
return ret
}