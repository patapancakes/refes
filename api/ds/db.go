package ds

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user   = ""
	pass   = ""
	host   = ""
	dbname = ""
)

var db = newDbConn()

func newDbConn() *sql.DB {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, dbname))
	if err != nil {
		panic(err)
	}

	return conn
}

func getContestListEntries(region string) (map[string]ContestListEntry, error) {
	table := "contests_us"
	if region == "JPN" {
		table = "contests_jp"
	}

	results, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		return nil, err
	}

	defer results.Close()

	contestListEntries := make(map[string]ContestListEntry)
	for cnt := 0; results.Next(); cnt++ {
		var id int
		var name string
		var applyStart, applyEnd, reviewStart, reviewEnd, excStart, excEnd time.Time
		err := results.Scan(&id, &name, &applyStart, &applyEnd, &reviewStart, &reviewEnd, &excStart, &excEnd)
		if err != nil {
			return nil, err
		}

		contestListEntries[strconv.Itoa(cnt)] = ContestListEntry{
			ID:          strconv.Itoa(id),
			Name:        base64.StdEncoding.EncodeToString([]byte(name)),
			ApplyStart:  applyStart.Format("2006-01-02 15:04:05"),
			ApplyEnd:    applyEnd.Format("2006-01-02 15:04:05"),
			ReviewStart: applyStart.Format("2006-01-02 15:04:05"),
			ReviewEnd:   applyEnd.Format("2006-01-02 15:04:05"),
			ExcStart:    excStart.Format("2006-01-02 15:04:05"),
			ExcEnd:      excEnd.Format("2006-01-02 15:04:05"),
		}
	}

	return contestListEntries, nil
}

// filter | title - uname - suid - password
// sort | updt - dlcount - reviewave
// direction | asc - desc
func getRpgListEntries(region, filter, keyword, sort, direction string, count, offset int) (map[string]RpgListEntry, error) {
	table := "games_us"
	if region == "JPN" {
		table = "games_jp"
	}

	query := "SELECT * FROM " + table

	if filter != "" {
		query += " WHERE " + filter + " ="

		if filter == "password" {
			query += " " + keyword // do not use wildcard for password filter
		} else {
			query += " LIKE \"%" + keyword + "%\""
		}
	}

	if sort != "" {
		query += " ORDER BY " + sort + " " + direction
	}

	if count > 0 {
		query += " LIMIT " + strconv.Itoa(count)

		if offset > 0 { // nested because OFFSET without LIMIT is pointless
			query += " OFFSET " + strconv.Itoa(offset)
		}
	}

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer results.Close()

	rpgListEntries := make(map[string]RpgListEntry)
	for cnt := 0; results.Next(); cnt++ {
		var sid, suid, datablocksize, version, packageversion, edit, attribute, award, famer, contest, owner, dlcount int
		var title, uname, password, lang, comment, genre string
		var updt time.Time
		var reviewave float64
		err := results.Scan(&sid, &suid, &title, &uname, &password, &updt, &datablocksize, &version, &packageversion, &reviewave, &lang, &edit, &attribute, &award, &famer, &comment, &contest, &owner, &genre, &dlcount)
		if err != nil {
			return nil, err
		}

		rpgListEntry := RpgListEntry{
			SID:            strconv.Itoa(sid),
			SUID:           strconv.Itoa(suid),
			Title:          base64.StdEncoding.EncodeToString([]byte(title)),
			UName:          base64.StdEncoding.EncodeToString([]byte(uname)),
			Password:       password,
			Updt:           updt.Format("2006-01-02 15:04:05"),
			DataBlockSize:  strconv.Itoa(datablocksize),
			Version:        strconv.Itoa(version),
			PackageVersion: strconv.Itoa(packageversion),
			ReviewAve:      strconv.FormatFloat(reviewave, 'f', 5, 64),
			Lang:           lang,
			Edit:           strconv.Itoa(edit),
			Attribute:      strconv.Itoa(attribute),
			Award:          strconv.Itoa(award),
			Famer:          strconv.Itoa(famer),
			Comment:        base64.StdEncoding.EncodeToString([]byte(comment)),
			Contest:        strconv.Itoa(contest),
			Owner:          strconv.Itoa(owner),
			DLCount:        strconv.Itoa(dlcount),
		}

		// the genre system is so bad there's probably no better way to do this
		for _, str := range strings.Split(genre, ",") {
			switch str {
			case "1":
				rpgListEntry.Genre1 = "1"
			case "2":
				rpgListEntry.Genre2 = "1"
			case "3":
				rpgListEntry.Genre3 = "1"
			case "4":
				rpgListEntry.Genre4 = "1"
			case "5":
				rpgListEntry.Genre5 = "1"
			case "6":
				rpgListEntry.Genre6 = "1"
			case "7":
				rpgListEntry.Genre7 = "1"
			case "8":
				rpgListEntry.Genre8 = "1"
			case "9":
				rpgListEntry.Genre9 = "1"
			case "10":
				rpgListEntry.Genre10 = "1"
			case "11":
				rpgListEntry.Genre11 = "1"
			case "12":
				rpgListEntry.Genre12 = "1"
			case "13":
				rpgListEntry.Genre13 = "1"
			case "14":
				rpgListEntry.Genre14 = "1"
			case "15":
				rpgListEntry.Genre15 = "1"
			case "16":
				rpgListEntry.Genre16 = "1"
			case "17":
				rpgListEntry.Genre17 = "1"
			case "18":
				rpgListEntry.Genre18 = "1"
			case "19":
				rpgListEntry.Genre19 = "1"
			case "20":
				rpgListEntry.Genre20 = "1"
			case "21":
				rpgListEntry.Genre21 = "1"
			case "22":
				rpgListEntry.Genre22 = "1"
			case "23":
				rpgListEntry.Genre23 = "1"
			case "24":
				rpgListEntry.Genre24 = "1"
			case "25":
				rpgListEntry.Genre25 = "1"
			case "26":
				rpgListEntry.Genre26 = "1"
			case "27":
				rpgListEntry.Genre27 = "1"
			case "28":
				rpgListEntry.Genre28 = "1"
			case "29":
				rpgListEntry.Genre29 = "1"
			case "30":
				rpgListEntry.Genre30 = "1"
			case "31":
				rpgListEntry.Genre31 = "1"
			case "32":
				rpgListEntry.Genre32 = "1"
			case "33":
				rpgListEntry.Genre33 = "1"
			case "34":
				rpgListEntry.Genre34 = "1"
			}
		}

		rpgListEntries[strconv.Itoa(cnt)] = rpgListEntry
	}

	return rpgListEntries, nil
}

func getRpgPublic(sid int, region string) (bool, error) {
	table := "games_us"
	if region == "JPN" {
		table = "games_jp"
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM "+table+" WHERE suid = ? AND region = ?", sid, region).Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
