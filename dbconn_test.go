package main

import (
	"testing"
	"encoding/gob"
	"bytes"
	_"fmt"
	"os"
	"log"
	"database/sql"
)

func TestMain(m* testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	os.Remove("./test.db")
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE photos (
	id			INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
	time			INTEGER NOT NULL, 
	base_uri		STRING NOT NULL, 
	filename		STRING NOT NULL, 
	description		TEXT NOT NULL, 
	roll_id			INTEGER NOT NULL, 
	default_version_id	INTEGER NOT NULL, 
	rating			INTEGER NULL 
);
INSERT INTO photos VALUES(1,1249579156,'file:///tmp/database/','sample.jpg','Testing!',1,2,5);
INSERT INTO photos VALUES(2,1276191607,'file:///tmp/database/','sample_canon_bibble5.jpg','',1,1,0);
INSERT INTO photos VALUES(3,1249834364,'file:///tmp/database/','sample_canon_zoombrowser.jpg','%test comment%',1,1,0);
INSERT INTO photos VALUES(4,1276191607,'file:///tmp/database/','sample_gimp_exiftool.jpg','',1,1,5);
INSERT INTO photos VALUES(5,1242995279,'file:///tmp/database/','sample_nikon1.jpg','',1,1,1);
INSERT INTO photos VALUES(6,1276191607,'file:///tmp/database/','sample_nikon1_bibble5.jpg','',1,1,0);
INSERT INTO photos VALUES(7,1167646774,'file:///tmp/database/','sample_nikon2.jpg','',1,1,0);
INSERT INTO photos VALUES(8,1276191607,'file:///tmp/database/','sample_nikon2_bibble5.jpg','',1,1,0);
INSERT INTO photos VALUES(9,1256140553,'file:///tmp/database/','sample_nikon3.jpg','                                    ',1,1,0);
INSERT INTO photos VALUES(10,1238587697,'file:///tmp/database/','sample_nikon4.jpg','                                    ',1,1,0);
INSERT INTO photos VALUES(11,1276191607,'file:///tmp/database/','sample_no_metadata.jpg','',1,1,0);
INSERT INTO photos VALUES(12,1265446642,'file:///tmp/database/','sample_null_orientation.jpg','',1,1,0);
INSERT INTO photos VALUES(13,1161575860,'file:///tmp/database/','sample_olympus1.jpg','',1,1,0);
INSERT INTO photos VALUES(14,1236006332,'file:///tmp/database/','sample_olympus2.jpg','',1,1,0);
INSERT INTO photos VALUES(15,1246010310,'file:///tmp/database/','sample_panasonic.jpg','',1,1,0);
INSERT INTO photos VALUES(16,1258799979,'file:///tmp/database/','sample_sony1.jpg','',1,1,0);
INSERT INTO photos VALUES(17,1257533767,'file:///tmp/database/','sample_sony2.jpg','',1,1,0);
INSERT INTO photos VALUES(18,1026565108,'file:///tmp/database/','sample_xap.jpg','',1,1,4);
INSERT INTO photos VALUES(19,1093249257,'file:///tmp/database/','sample_xmpcrash.jpg','',1,1,0);
INSERT INTO photos VALUES(20,1276191607,'file:///tmp/database/test/','sample_tangled1.jpg','test comment',1,1,0);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func teardown() {
	os.Remove("./test.db")
}

func Test_DbCreation(t* testing.T) {
	db := NewDb("./test.db")

	if db == nil {
		t.Error("Db object is nil")
	}
}

func Test_SelectFromTable(t* testing.T) {
	db := NewDb("./test.db")
	if db == nil {
		t.Error("Db object is nil")
	}

	tabledata := db.SelectAllFrom("photos")

	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(tabledata)
	if err != nil {
		panic(err)
	}
	var decoded []map[string]string
	d := gob.NewDecoder(b)
	err = d.Decode(&decoded)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%#v\n", decoded)
}
