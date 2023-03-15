package db

import (
	"encoding/csv"
	"fmt"
	"os"

	log "unknwon.dev/clog/v2"
)

// RCOS spesific code
type Affiliation struct {
	ID            int64
	Name          string
	DisplayedName string
	Url           string `xorm:"UNIQUE NOT NULL" gorm:"UNIQUE"`
	Alias         string
	Description   string
}

// RCOS spesific code.
// InitAffiliation inserts or updates affiliation's table from a csv file.
func InitAffiliation() {
	filePath := "conf/affiliation/affiliation.csv"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open %s file: %v", filePath, err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to read %s file: %v", filePath, err)
		return
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
		return
	}

	for _, org := range orgs {
		bean := &Affiliation{Url: org.Url}
		has, err := sess.Get(bean)
		if err != nil {
			log.Fatal("Failed to get affiliation : %v", err)
			sess.Rollback()
			return
		} else if has {
			if _, err = sess.Where("url = ?", org.Url).Update(org); err != nil {
				log.Fatal("Failed to update affiliation : %v", err)
				sess.Rollback()
				return
			}
		} else {
			if _, err = sess.Insert(org); err != nil {
				log.Fatal("Failed to insert affiliation : %v", err)
				sess.Rollback()
				return
			}
		}
	}

	sess.Commit()
}

// RCOS spesific code.
// GetAffiliationList return map like {Affiliation.ID:Affliation.Name}.
func GetAffiliationList() (map[int64]string, error) {

	var beans []*Affiliation
	err := x.Find(&beans)
	list := make(map[int64]string)

	for _, bean := range beans {
		list[bean.ID] = bean.DisplayedName
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
