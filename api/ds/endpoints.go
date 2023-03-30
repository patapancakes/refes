package ds

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"

	"github.com/klauspost/compress/zstd"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: request to %s\n", r.RequestURI)

	if r.Method != "POST" {
		log.Printf("ERROR: %s method not supported\n", r.Method)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: failed to read request body: %s\n", err)
		return
	}

	if len(body) == 0 {
		log.Println("ERROR: empty request body")
		return
	}

	if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		bodyUnescaped, err := url.PathUnescape(string(body))
		if err != nil {
			log.Printf("ERROR: failed to unescape request body: %s\n", err)
			return
		}

		if len(bodyUnescaped) < 5 || bodyUnescaped[:5] != "args=" { // should be safe
			log.Println("ERROR: malformed request body")
			return
		}

		body = []byte(bodyUnescaped[:5])
	}

	var response []byte
	switch r.RequestURI {
	case "/api/username": // register username
		response, err = handleUsername(body)
	case "/api/flags": // get server flags and user info
		response, err = handleFlags(body)
	case "/api/news": // get news
		response, err = handleNews(body)
	case "/api/signin": // make presence known to server?
		response, err = handleSignIn(body)
	case "/api/contestlist": // get contest list
		response, err = handleContestList(body)
	case "/api/rpglist", "/api/rpglisttitle", "/api/rpglistuname", "/api/rpglistsuid", "/api/rpglistpassword": //, "/api/myrpglist": // get rpg list of some kind
		response, err = handleRpgList(body, r.RequestURI[5:])
	case "/api/rpgdownload": // download rpg
		response, err = handleRpgDownload(body)
	case "/api/rpgreview": // review rpg
		response, err = handleRpgReview(body)
	case "/api/infomercial": // report rpg
		response, err = handleInfomercial(body)
	case "/api/rpgupload": // upload rpg
		response, err = handleRpgUpload(body)
	case "/api/rpgdelete": // delete rpg
		response, err = handleRpgDelete(body)
	default:
		err = fmt.Errorf("unknown endpoint: %s", r.RequestURI)
	}
	if err != nil {
		log.Printf("ERROR: handler for %s returned error: %s\n", r.RequestURI, err)
	
		w.WriteHeader(http.StatusNoContent) // write header so we don't cause bad gateway
		return
	}

	w.Write(response)
}

func handleUsername(body []byte) ([]byte, error) {
	usernameC := &UsernameC{}
	err := json.Unmarshal(body, usernameC)
	if err != nil {
		return nil, err
	}

	// TODO: do something here

	usernameS := &UsernameS{
		EndCode: "0",
	}

	response, err := json.Marshal(usernameS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleFlags(body []byte) ([]byte, error) {
	flagsC := &FlagsC{}
	err := json.Unmarshal(body, flagsC)
	if err != nil {
		return nil, err
	}

	flagsS := &FlagsS{
		ID: "1", // placeholder
		Region: flagsC.Region,
		Lang: flagsC.Lang,
		Maintenance: "0",
		SerchContest: "0",
		SerchFamer: "0",
		SerchOtherCountries: "1",
		ContestMode: "0",
		SUID: "1", // placeholder
		UName: base64.StdEncoding.EncodeToString([]byte("reFES User")), // placeholder
		Flag1: -1,
		Flag2: -1,
		Flag3: -1,
		EndCode: 0,
	}

	response, err := json.Marshal(flagsS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleNews(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, nil
}

func handleSignIn(body []byte) ([]byte, error) {
	signInC := &SignInC{}
	err := json.Unmarshal(body, signInC)
	if err != nil {
		return nil, err
	}

	// TODO: do something here

	signInS := &SignInS{
		EndCode: "0",
	}

	response, err := json.Marshal(signInS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleContestList(body []byte) ([]byte, error) {
	contestListC := &ContestListC{}
	err := json.Unmarshal(body, contestListC)
	if err != nil {
		return nil, err
	}

	contestListEntries, err := getContestListEntries(contestListC.Region)
	if err != nil {
		return nil, err
	}

	contestListS := &ContestListS{
		ContestListEntries: contestListEntries,
		EndCode: 0,
	}

	response, err := json.Marshal(contestListS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleRpgList(body []byte, kind string) ([]byte, error) {
	rpgListC := &RpgListC{}
	err := json.Unmarshal(body, rpgListC)
	if err != nil {
		return nil, err
	}

	// a lot of sanitization is done here

	var sort string
	direction := "asc"
	switch {
	case rpgListC.SortUpdt != "":
		sort = "updt"
		if rpgListC.SortUpdt == "desc" {
			direction = "desc"
		}
	case rpgListC.SortDLCount != "":
		sort = "dlcount"
		if rpgListC.SortDLCount == "desc" {
			direction = "desc"
		}
	case rpgListC.SortReviewAve != "":
		sort = "reviewave"
		if rpgListC.SortReviewAve == "desc" {
			direction = "desc"
		}
	}

	var filter string
	switch kind {
	case "rpglisttitle":
		filter = "title"
	case "rpglistuname":
		filter = "uname"
	case "rpglistsuid":
		filter = "suid"
	case "rpglistpassword":
		filter = "password"
	}

	var keyword string
	if filter != "" {
		decoded, err := base64.StdEncoding.DecodeString(rpgListC.Keyword)
		if err != nil {
			return nil, err
		}

		if !regexp.MustCompile("^[a-zA-Z0-9]+$").Match(decoded) {
			return nil, fmt.Errorf("keyword \"%s\" not allowed by regex", decoded)
		}

		keyword = string(decoded)
	}

	rpgListEntries, err := getRpgListEntries(rpgListC.Region, filter, keyword, sort, direction, rpgListC.RecNum, rpgListC.Offset)
	if err != nil {
		return nil ,err
	}

	rpgListS := &RpgListS{
		RpgListEntries: rpgListEntries,
		EndCode: 0,
	}

	response, err := json.Marshal(rpgListS)
	if err != nil {
		return nil ,err
	}

	return response, nil
}

func handleRpgDownload(body []byte) ([]byte, error) {
	rpgDownloadC := &RpgDownloadC{}
	err := json.Unmarshal(body, rpgDownloadC)
	if err != nil {
		return nil, err
	}

	sid, err := strconv.Atoi(rpgDownloadC.SID)
	if err != nil {
		return nil, err
	}

	public, err := getRpgPublic(sid, rpgDownloadC.Region)
	if err != nil {
		return nil, err
	}

	if !public {
		return nil, fmt.Errorf("attempt to download non-public game: %d/%s", sid, rpgDownloadC.Region)
	}

	gameDir := "games_us"
	if rpgDownloadC.Region == "JPN" {
		gameDir = "games_jp"
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/game%06d.zst", gameDir, sid))
	if err != nil {
		return nil, err
	}

	dec, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}

	decompressed, err := dec.DecodeAll(file, []byte{})
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}

func handleRpgReview(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, nil
}

func handleInfomercial(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, nil
}

func handleRpgUpload(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, nil
}

func handleRpgDelete(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, nil
}
