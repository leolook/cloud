package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	dbList map[string]*sql.DB
	lock   sync.Locker
)

func init() {
	if lock == nil {
		lock = new(sync.Mutex)
	}
}

func NewClient(url string) (*sql.DB, error) {

	if v, ok := dbList[url]; ok {
		return v, nil
	}

	lock.Lock()
	defer lock.Unlock()

	db, err := sql.Open("mysql",
		url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if dbList == nil {
		dbList = make(map[string]*sql.DB)
	}
	dbList[url] = db
	return db, nil
}

func GetClient(url string) *sql.DB {
	if dbList == nil {
		return nil
	}
	if v, ok := dbList[url]; ok {
		return v
	}
	return nil
}

func RemoveClient(url string) {
	if dbList == nil {
		return
	}
	lock.Lock()
	defer lock.Unlock()
	if _, ok := dbList[url]; ok {
		delete(dbList, url)
	}
}
