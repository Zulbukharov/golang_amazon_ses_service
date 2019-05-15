package db

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func Init() error {
	var err error
	// db, err = buntdb.Open("data.db")
	db, err = buntdb.Open(":memory:") // Open a file that does not persist to disk./
	if err != nil {
		log.Fatal(err)
	}
	return err
	// defer db.Close()
}

//	Returns converted to <int> passed as interface{}
//
//	return: string
func genRandomFourDigitNumber() (code string) {
	rand.Seed(time.Now().UnixNano())
	code = strconv.Itoa(1000 + rand.Intn(9999-1000))
	return
}

func SetEmailGenerateCode(contact_info string) (verification_code string, err error) {
	verification_code = genRandomFourDigitNumber()
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(contact_info, verification_code, &buntdb.SetOptions{Expires: true, TTL: time.Hour})
		return err
	})
	fmt.Println(time.Second)
	return
}

func VerifyCode(contact_info, verification_code string) (res bool, err error) {
	var value string
	err = db.View(func(tx *buntdb.Tx) error {
		value, err = tx.Get(contact_info)
		if err != nil {
			return err
		}
		return nil
	})
	if value == verification_code {
		res = true
	}
	return
}
