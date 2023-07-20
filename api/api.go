/*
	reFES - A RPG Maker FES server emulator
	Copyright (C) 2023  maru <maru@myyahoo.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package api

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"unicode/utf16"
	"unicode/utf8"
)

func Init(proto *string, address *string) error {
	http.HandleFunc("/", handleRequest)

	log.Printf("INFO: server starting on %s\n", *address)

	listener, err := net.Listen(*proto, *address)
	if err != nil {
		return err
	}

	if *proto == "unix" {
		os.Chmod(*address, 0777)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		return err
	}

	return nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
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

		body = []byte(bodyUnescaped[5:])
	}

	var response []byte
	switch r.RequestURI {
	case "/api/username": // register username
		response, err = handleUsername(body)
	case "/api/flags": // get server flags and user info
		response, err = handleFlags(body)
	case "/api/signin": // make presence known to server?
		response, err = handleSignIn(body)
	case "/api/news": // get news
		response, err = handleNews(body)
	case "/api/contestlist": // get contest list
		response, err = handleContestList(body)
	case "/api/rpglist", "/api/rpglisttitle", "/api/rpglistuname", "/api/rpglistsuid", "/api/rpglistpassword": // get rpg list of some kind
		response, err = handleRpgList(body, r.RequestURI[12:])
	case "/api/myrpglist": // get your uploaded rpgs
		response, err = handleMyRpgList(body)
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

		w.WriteHeader(http.StatusBadRequest) // write header so we don't cause bad gateway
		return
	}

	if utf8.Valid(response) {
		respUtf16 := utf16.Encode([]rune(string(response)))

		response = make([]byte, len(respUtf16)*2)
		for i, v := range respUtf16 {
			binary.LittleEndian.PutUint16(response[i*2:i*2+2], v)
		}
	}

	w.Write(response)
}
