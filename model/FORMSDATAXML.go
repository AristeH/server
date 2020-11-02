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



type FORMSDATAXML struct {
	VIEWNAME string `json:""` // VIEWNAME
	DATAXML string `json:""` // DATAXML
}

type FORMSDATAXMLList struct {
   Recs []FORMSDATAXML `json:""`
}

func  (ob FORMSDATAXML) Create() error{
    sqlstr := `create table FORMSDATAXML  (VIEWNAME, DATAXML)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob FORMSDATAXML) Delete() error{
   const sqlstr = `DELETE FROM FORMSDATAXML  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob FORMSDATAXML) Save() error{
   sqlstr := "update or insert into  FORMSDATAXML  (VIEWNAME, DATAXML) "+
   " values (?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.VIEWNAME, ob.DATAXML)
   if err != nil {
     return err
   }
   return nil
}

func (ob FORMSDATAXML) Read(id string) error{
   const sqlstr = `select * FROM FORMSDATAXML  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.VIEWNAME, &ob.DATAXML,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob FORMSDATAXML) ReadFromJson(file string){
	var recs FORMSDATAXMLList

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recs)

	for i := 0; i < len(recs.Recs); i++ {
	   ob.VIEWNAME = recs.Recs[i].VIEWNAME
	   ob.DATAXML = recs.Recs[i].DATAXML
	   ob.Save()
	}

}

func  (ob FORMSDATAXML)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "",
		Stroki: []arrayFieldSection{
			{
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.VIEWNAME,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "",
						Value:    ob.DATAXML,
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

func  (ob FORMSDATAXML)  FormSpisok() string{
ret := ""



 
 
return ret
}