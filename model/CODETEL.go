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



type CODETEL struct {
	ID string `json:"Ссылка"` // ID
	NAME string `json:"Наименование"` // NAME
	CODE string `json:"Код"` // CODE
	ID_OPERATOR string `json:"оператор"` // ID_OPERATOR
	ID_REGION string `json:"регион"` // ID_REGION
	S string `json:"с"` // S
	PO string `json:"по"` // PO
}

type CODETELList struct {
   Recs []CODETEL `json:"кодытелефонов"`
}

func  (ob CODETEL) Create() error{
    sqlstr := `create table CODETEL  (ID, NAME, CODE, ID_OPERATOR, ID_REGION, S, PO)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob CODETEL) Delete() error{
   const sqlstr = `DELETE FROM CODETEL  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob CODETEL) Save() error{
   sqlstr := "update or insert into  CODETEL  (ID, NAME, CODE, ID_OPERATOR, ID_REGION, S, PO) "+
   " values (?, ?, ?, ?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.NAME, ob.CODE, ob.ID_OPERATOR, ob.ID_REGION, ob.S, ob.PO)
   if err != nil {
     return err
   }
   return nil
}

func (ob CODETEL) Read(id string) error{
   const sqlstr = `select * FROM CODETEL  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.NAME, &ob.CODE, &ob.ID_OPERATOR, &ob.ID_REGION, &ob.S, &ob.PO,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob CODETEL) ReadFromJson(file string){
	var recs CODETELList

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
	   ob.ID_OPERATOR = recs.Recs[i].ID_OPERATOR
	   ob.ID_REGION = recs.Recs[i].ID_REGION
	   ob.S = recs.Recs[i].S
	   ob.PO = recs.Recs[i].PO
	   ob.Save()
	}

}

func  (ob CODETEL)  TmplElem(id string) string{
   
	v := ListForm{
		Name:  "ListForm",
		Title: "Коды телефонов",
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
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "Код",
						Value:   ob.CODE, 
							   
							    
							           
    					Buttons: "",
					},
				},
			},
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "Оператор",
						Value:   ob.ID_OPERATOR, 
							   
							    
							           
    					Buttons: "",
					},
				},
			},
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "Регион",
						Value:   ob.ID_REGION, 
							   
							    
							           
    					Buttons: "",
					},
				},
			},
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "С",
						Value:   ob.S, 
							   
							    
							           
    					Buttons: "",
					},
				},
			},
	        { 
	  		    Fields: []FieldSection{
					{
						Name:     "По",
						Value:   ob.PO, 
							   
							    
							           
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

func  (ob CODETEL)  FormSpisok() string{
ret := ""



 
 
return ret
}