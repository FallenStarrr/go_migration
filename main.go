package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}


func main() {

	sql_migration, err := os.Create("C:/Users/u14911/Desktop/Maquette/Миграции/migration2.sql")
	check(err)

	type docs []struct {
		Name        string `json:"name"`
		ParentID    string `json:"parent_id"`
		DeepestNode bool   `json:"deepest_node"`
	}
	filea,_ := ioutil.ReadFile("C:/Users/u14911/Desktop/Maquette/Миграции/docs.json")

	ads:= docs{}
	json.Unmarshal(filea,&ads)


	defer sql_migration.Close()

	query := ""
	for _, v := range ads {
		if v.ParentID == "" {
			v.ParentID = "null"
		}else {
			v.ParentID = "'"+v.ParentID+"'"
		}
		id := uuid.NewString()
		query += fmt.Sprintf(
			`

				 insert into document_type (id, doc_type, parent_id, deepest_node, expiration_time)
				 values (%q, %q, %v, %v, 99);
                 insert into roles(id, role)
				 values (%q,'COMMON-READ');
				 insert into roles(id, role)
				 values (%q,'COMMON-WRITE');

				

				 `,   id,  v.Name, v.ParentID, v.DeepestNode,  id,id)


	}
        query = strings.Replace(query, "\"", "'", -1)
	fmt.Println(query)
	sql_migration.WriteString(query)

	
}
