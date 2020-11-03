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



type OPERATOR struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
}

type OPERATORList struct {
   Recs []OPERATOR `json:""`
}

func  (ob OPERATOR) Create() error{
    sqlstr := `create table OPERATOR  (ID, NAME)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob OPERATOR) Delete() error{
   const sqlstr = `DELETE FROM OPERATOR  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob OPERATOR) Save() error{
   sqlstr := "update or insert into  OPERATOR  (ID, NAME) "+
   " values (?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME)
   if err != nil {
     return err
   }
   return nil
}

func (ob OPERATOR) Read(id string) error{
   const sqlstr = `select * FROM OPERATOR  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob OPERATOR) ReadFromJson(file string){
	var recs OPERATORList

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

func  (ob OPERATOR)  TmplElem(id string) string{
   
	v := ListForm{
		Name:  "ListForm",
		Title: "",
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

func  (ob OPERATOR)  FormSpisok() string{
ret := ""



 
 
return ret
}