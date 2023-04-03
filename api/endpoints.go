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
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/klauspost/compress/zstd"
)

func handleUsername(body []byte) ([]byte, error) {
	usernameC := &UsernameC{}
	err := json.Unmarshal(body, usernameC)
	if err != nil {
		return nil, err
	}

	// TODO: do something here

	usernameS := &UsernameS{
		EndCode: 0,
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
		Id:                  "1", // placeholder
		Region:              flagsC.Region,
		Lang:                flagsC.Lang,
		Maintenance:         "0",
		SerchContest:        "0",
		SerchFamer:          "0",
		SerchOtherCountries: "1",
		ContestMode:         "0",
		Suid:                "1",                                                     // placeholder
		Uname:               base64.StdEncoding.EncodeToString([]byte("reFES User")), // placeholder
		Flag1:               -1,
		Flag2:               -1,
		Flag3:               -1,
		EndCode:             0,
	}

	response, err := json.Marshal(flagsS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleSignIn(body []byte) ([]byte, error) {
	signInC := &SignInC{}
	err := json.Unmarshal(body, signInC)
	if err != nil {
		return nil, err
	}

	// TODO: do something here

	signInS := &SignInS{
		EndCode: 0,
	}

	response, err := json.Marshal(signInS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleNews(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, errors.New("not implemented")
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
		EndCode:            0,
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
	direction := "ASC"
	switch {
	case rpgListC.SortUpdt != "":
		sort = "updt"
		if rpgListC.SortUpdt == "desc" {
			direction = "DESC"
		}
	case rpgListC.SortDlCount != "":
		sort = "dlcount"
		if rpgListC.SortDlCount == "desc" {
			direction = "DESC"
		}
	case rpgListC.SortReviewAve != "":
		sort = "reviewave"
		if rpgListC.SortReviewAve == "desc" {
			direction = "DESC"
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
		decoded, err := base64.RawStdEncoding.DecodeString(rpgListC.Keyword)
		if err != nil {
			return nil, err
		}

		keyword = string(decoded)
	}

	rpgListEntries, err := getRpgListEntries(rpgListC.Region, filter, keyword, sort, direction, rpgListC.Contest, rpgListC.Famer, rpgListC.RecNum, rpgListC.Offset)
	if err != nil {
		return nil, err
	}

	rpgListS := &RpgListS{
		RpgListEntries: rpgListEntries,
		EndCode:        0,
	}

	response, err := json.Marshal(rpgListS)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleRpgDownload(body []byte) ([]byte, error) {
	rpgDownloadC := &RpgDownloadC{}
	err := json.Unmarshal(body, rpgDownloadC)
	if err != nil {
		return nil, err
	}

	public, err := getRpgPublic(rpgDownloadC.Sid, rpgDownloadC.Region)
	if err != nil {
		return nil, err
	}

	if !public {
		return nil, fmt.Errorf("attempt to download non-public game: %d/%s", rpgDownloadC.Sid, rpgDownloadC.Region)
	}

	gameDir := "games_us"
	if rpgDownloadC.Region == "JPN" {
		gameDir = "games_jp"
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/game%06d.zst", gameDir, rpgDownloadC.Sid))
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
	return nil, errors.New("not implemented")
}

func handleInfomercial(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, errors.New("not implemented")
}

func handleRpgUpload(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, errors.New("not implemented")
}

func handleRpgDelete(body []byte) ([]byte, error) {
	// TODO: do something here
	return nil, errors.New("not implemented")
}
