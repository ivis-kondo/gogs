package db

import (
	"encoding/csv"
	"fmt"
	"os"

	log "unknwon.dev/clog/v2"
)

// RCOS spesific code
type Affiliation struct {
	ID          int64
	Name        string
	Url         string `xorm:"UNIQUE NOT NULL" gorm:"UNIQUE"`
	Alias       string
	Description string
}

// RCOS spesific code.
// RegisterAffiliation register Research Laboratory from a csv file.
// TODO: うまくいったらinternal/route/install.go GlobalInitへ
func InitAffiliation() {

	filePath := "conf/affiliation/affiliation.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to get %s file: %v", filePath, err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#' // If the line starts with #, treat it as a comment
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to read %s file: %v", filePath, err)
	}

	orgs := make([]Affiliation, 0, len(rows))

	for _, v := range rows {
		orgs = append(orgs, Affiliation{
			Name: v[0],
			Url:  v[1],
		})
	}

	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		log.Fatal("Failed to begin a transaction : %v", err)
	}
	if _, err = sess.Insert(orgs); err != nil {
		log.Fatal("Failed to insert affiliation : %v", err)
	}
}

// RCOS spesific code.
// GetAffiliationList return map like {Affiliation.ID:Affliation.Name}.
func GetAffiliationList() (map[int64]string, error) {

	var beans []*Affiliation
	err := x.Find(&beans)
	list := make(map[int64]string)

	for _, bean := range beans {
		list[bean.ID] = bean.Name
	}

	return list, err
}

// RCOS spesific code.
// GetAffiliationByID returns an affiliation by given ID.
func GetAffiliationByID(id int64) (*Affiliation, error) {

	affi := new(Affiliation)
	has, err := x.ID(id).Get(affi)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("no affiliation:id= %v", id)
	}
	return affi, nil

}
