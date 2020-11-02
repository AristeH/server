// Code generated. DO NOT EDIT.

package model

import (
	"encoding/json"
  
 "strings"

    "time"
	"my/server/config"
	"io/ioutil"
	"os"
	"fmt"
)



type PERSONEL struct {
	ID int `json:"Ссылка"` // ID
	ID_PERSON string `json:"физическоелицо"` // ID_PERSON
	ID_POST string `json:"Должности"` // ID_POST
	ID_DEPARTMENT string `json:"Подразделения"` // ID_DEPARTMENT
	PERSONNUMBER string `json:"табельныйномер"` // PERSONNUMBER
	EMPLOYMENTTYPE bool `json:"видзанятости"` // EMPLOYMENTTYPE
	EVENTSTART string `json:"видсобытияначало"` // EVENTSTART
	DATESTART time.Time `json:"датасобытияначало"` // DATESTART
	EVENTEND string `json:"видсобытияокончание"` // EVENTEND
	DATEEND time.Time `json:"датасобытияокончание"` // DATEEND
}

type PERSONELList struct {
   Recs []PERSONEL `json:"перемещения"`
}

func  (ob PERSONEL) Create() error{
    sqlstr := `create table PERSONEL  (ID, ID_PERSON, ID_POST, ID_DEPARTMENT, PERSONNUMBER, EMPLOYMENTTYPE, EVENTSTART, DATESTART, EVENTEND, DATEEND)
                         CONSTRAINT CODETEL_PK PRIMARY KEY (ID));`
    _, err := config.DB.Exec(sqlstr)
   	if err != nil {
   	   return err
	}
	return nil
}

func (ob PERSONEL) Delete() error{
   const sqlstr = `DELETE FROM PERSONEL  WHERE ID =  ?`
		_, err := config.DB.Exec(sqlstr, ob.ID)
		if err != nil {
			return err
		}
		return nil
}

func (ob PERSONEL) Save() error{
   sqlstr := "update or insert into  PERSONEL  (ID, ID_PERSON, ID_POST, ID_DEPARTMENT, PERSONNUMBER, EMPLOYMENTTYPE, EVENTSTART, DATESTART, EVENTEND, DATEEND) "+
   " values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)" +
   " matching (ID)"
   _, err := config.DB.Exec(sqlstr,  ob.ID, ob.ID_PERSON, ob.ID_POST, ob.ID_DEPARTMENT, ob.PERSONNUMBER, ob.EMPLOYMENTTYPE, ob.EVENTSTART, ob.DATESTART, ob.EVENTEND, ob.DATEEND)
   if err != nil {
     return err
   }
   return nil
}

func (ob PERSONEL) Read(id string) error{
   const sqlstr = `select * FROM PERSONEL  WHERE ID =  ?`
   row := config.DB.QueryRow(sqlstr, id)
   err := row.Scan( &ob.ID, &ob.ID_PERSON, &ob.ID_POST, &ob.ID_DEPARTMENT, &ob.PERSONNUMBER, &ob.EMPLOYMENTTYPE, &ob.EVENTSTART, &ob.DATESTART, &ob.EVENTEND, &ob.DATEEND,)
   if err != nil {
	 return err
   }
   return nil
}

func (ob PERSONEL) ReadFromJson(file string){
	var recs PERSONELList

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recs)

	for i := 0; i < len(recs.Recs); i++ {
	   ob.ID = recs.Recs[i].ID
	   ob.ID_PERSON = recs.Recs[i].ID_PERSON
	   ob.ID_POST = recs.Recs[i].ID_POST
	   ob.ID_DEPARTMENT = recs.Recs[i].ID_DEPARTMENT
	   ob.PERSONNUMBER = recs.Recs[i].PERSONNUMBER
	   ob.EMPLOYMENTTYPE = recs.Recs[i].EMPLOYMENTTYPE
	   ob.EVENTSTART = recs.Recs[i].EVENTSTART
	   ob.DATESTART = recs.Recs[i].DATESTART
	   ob.EVENTEND = recs.Recs[i].EVENTEND
	   ob.DATEEND = recs.Recs[i].DATEEND
	   ob.Save()
	}

}

func  (ob PERSONEL)  TmplElem(id string) string{
   
	v := listform{
		Name:  "listform",
		Title: "Кадровые данные",
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
						Name:     "Физическое лицо",
						Value:    ob.ID_PERSON,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Должности",
						Value:    ob.ID_POST,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Подразделение",
						Value:    ob.ID_DEPARTMENT,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Табельный номер",
						Value:    ob.PERSONNUMBER,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид занятости",
						Value:    ob.EMPLOYMENTTYPE,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид события начало",
						Value:    ob.EVENTSTART,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Дата события начало",
						Value:    ob.DATESTART,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Вид события окончание",
						Value:    ob.EVENTEND,
						Buttons: "",
					},
				},
			},
	  		Fields: []FieldSection{
					{
						Name:     "Дата события окончание",
						Value:    ob.DATEEND,
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

func  (ob PERSONEL)  FormSpisok() string{
ret := ""



 
 
return ret
}