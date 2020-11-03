// Code generated. DO NOT EDIT.

package model

import (
	"encoding/xml"
	"encoding/json"
	
    
	"io/ioutil"
	"os"
	"fmt"
	"пппппп/server/config"
)



type POST struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
}

type POSTList struct {
   Recs []POST `json:"должности"`
}

func  (ob POST) Create() error{
    sqlstr := `create table POST  (ID, NAME)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob POST) Delete() error{
   const sqlstr = `DELETE FROM POST  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob POST) Save() error{
   sqlstr := "update or insert into  POST  (ID, NAME) "+
   " values (?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME)
   if err != nil {
     return err
   }
   return nil
}

func (ob POST) Read(id string) error{
   const sqlstr = `select * FROM POST  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob POST) ReadFromJson(file string){
	var recs POSTList

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

func  (ob POST)  TmplElem(id string) string{
   
	v := ListForm{
		Name:  "ListForm",
		Title: "Должности",
		Stroki: []arrayFieldSection{
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "Ссылка",
						Value:   ob.ID, 
							   
							    
							           
    					Buttons: "",
					},
				},
			},
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "Наименование",
						Value:   ob.NAME, 
							   
							    
							           
    					Buttons: "",
					},
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

func  (ob POST)  FormSpisok() string{
ret := ""



 
 
return ret
}