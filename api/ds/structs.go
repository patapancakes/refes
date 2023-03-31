/*
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

package ds

type GenericC struct {
	Region string `json:"region"`
	Lang   string `json:"lang"`
	Token  string `json:"token"`
}

type GenericS struct {
	EndCode int `json:"EndCode"`
}

// /api/username
type UsernameC struct {
	Region string `json:"region"`
	Lang   string `json:"lang"`
	Token  string `json:"token"`
	UName  string `json:"uname"`
}
type UsernameS GenericS

// /api/flags
type FlagsC GenericC
type FlagsS struct {
	ID                  string `json:"id"`
	Region              string `json:"region"`
	Lang                string `json:"lang"`
	Maintenance         string `json:"maintenance"`
	SerchContest        string `json:"serchcontest"`        // intentionally misspelled
	SerchFamer          string `json:"serchfamer"`          // intentionally misspelled
	SerchOtherCountries string `json:"serchothercountries"` // intentionally misspelled
	ContestMode         string `json:"contestmode"`
	SUID                string `json:"suid"`
	UName               string `json:"uname"`
	Flag1               int    `json:"flag1"`
	Flag2               int    `json:"flag2"`
	Flag3               int    `json:"flag3"`
	EndCode             int    `json:"endcode"`
}

// /api/signin
type SignInC GenericC
type SignInS GenericS

// /api/news (response is data)
type NewsC GenericC

// /api/contestlist
type ContestListC GenericC
type ContestListS struct {
	ContestListEntries map[string]ContestListEntry
	EndCode            int `json:"endcode"`
}
type ContestListEntry struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ApplyStart  string `json:"apply_start"`
	ApplyEnd    string `json:"apply_end"`
	ReviewStart string `json:"review_start"`
	ReviewEnd   string `json:"review_end"`
	ExcStart    string `json:"exc_start"`
	ExcEnd      string `json:"exc_end"`
	NowDate     string `json:"now_date"`
}

// /api/rpglist
type RpgListC struct {
	StartUpdt     int    `json:"startupdt"`
	Contest       int    `json:"contest"`
	SortUpdt      string `json:"sortupdt"`      // used for sorting
	SortDLCount   string `json:"sortdlcount"`   // used for sorting
	SortReviewAve string `json:"sortreviewave"` // used for sorting
	Keyword       string `json:"keyword"`       // used for searching
	Offset        int    `json:"offset"`
	RecNum        int    `json:"recnum"`
	Award         int    `json:"award"`
	Famer         int    `json:"famer"`
	Region        string `json:"region"`
	Lang          string `json:"lang"`
	Token         string `json:"token"`
}
type RpgListS struct {
	RpgListEntries map[string]RpgListEntry
	EndCode        int `json:"endcode"`
}
type RpgListEntry struct {
	SID            string `json:"sid"`
	SUID           string `json:"suid"`
	Title          string `json:"title"`
	UName          string `json:"uname"`
	Password       string `json:"password"`
	Updt           string `json:"updt"`
	DataBlockSize  string `json:"datablocksize"`
	Version        string `json:"version"`
	PackageVersion string `json:"packageversion"`
	ReviewAve      string `json:"reviewave"`
	Lang           string `json:"lang"`
	Edit           string `json:"edit"`
	Attribute      string `json:"attribute"`
	Award          string `json:"award"`
	Famer          string `json:"famer"`
	Comment        string `json:"comment"`
	Contest        string `json:"contest"`
	Owner          string `json:"owner"`
	DLCount        string `json:"dlcount"`

	Genre1  string `json:"genre1,omitempty"`  // Fantasy
	Genre2  string `json:"genre2,omitempty"`  // SF
	Genre3  string `json:"genre3,omitempty"`  // School Life
	Genre4  string `json:"genre4,omitempty"`  // Modern
	Genre5  string `json:"genre5,omitempty"`  // Japanese
	Genre6  string `json:"genre6,omitempty"`  // Adventure
	Genre7  string `json:"genre7,omitempty"`  // Puzzle
	Genre8  string `json:"genre8,omitempty"`  // Novel
	Genre9  string `json:"genre9,omitempty"`  // Hunt
	Genre10 string `json:"genre10,omitempty"` // Original
	Genre11 string `json:"genre11,omitempty"` // Romance
	Genre12 string `json:"genre12,omitempty"` // Training
	Genre13 string `json:"genre13,omitempty"` // Riddlle
	Genre14 string `json:"genre14,omitempty"` // Horror
	Genre15 string `json:"genre15,omitempty"` // Mystery
	Genre16 string `json:"genre16,omitempty"` // Classic
	Genre17 string `json:"genre17,omitempty"` // Comical
	Genre18 string `json:"genre18,omitempty"` // Serious
	Genre19 string `json:"genre19,omitempty"` // Heartful
	Genre20 string `json:"genre20,omitempty"` // Dark
	Genre21 string `json:"genre21,omitempty"` // Children's
	Genre22 string `json:"genre22,omitempty"` // Adult
	Genre23 string `json:"genre23,omitempty"` // For Men
	Genre24 string `json:"genre24,omitempty"` // For Women
	Genre25 string `json:"genre25,omitempty"` // Short
	Genre26 string `json:"genre26,omitempty"` // Long
	Genre27 string `json:"genre27,omitempty"` // Easy
	Genre28 string `json:"genre28,omitempty"` // Difficult
	Genre29 string `json:"genre29,omitempty"` // Collabo.
	Genre30 string `json:"genre30,omitempty"` // No Battle
	Genre31 string `json:"genre31,omitempty"` // Mini Game
	Genre32 string `json:"genre32,omitempty"` // Open Tech
	Genre33 string `json:"genre33,omitempty"` // Movie NG
	Genre34 string `json:"genre34,omitempty"` // Updated
}

// generic
/*type RpgListFilterC struct {
	Keyword string `json:"keyword"`
	Offset  int    `json:"offset"`
	RecNum  int    `json:"recnum"`
	Region  string `json:"region"`
	Lang    string `json:"lang"`
	Token   string `json:"token"`
}*/

// /api/rpglisttitle
type RpgListTitleC RpgListC
type RpgListTitleS RpgListS

// /api/rpglistuname
type RpgListUNameC RpgListC
type RpgListUNameS RpgListS

// /api/rpglistsuid
type RpgListSUIDC RpgListC
type RpgListSUIDS RpgListS

// /api/rpglistpassword
type RpgListPasswordC RpgListC
type RpgListPasswordS RpgListS

// /api/myrpglist
type MyRpgListC GenericC
type MyRpgListS RpgListS

// /api/rpgdownload (response is data)
type RpgDownloadC struct {
	Ver      string `json:"ver"`
	SID      string `json:"sid"`
	Region   string `json:"region"`
	Language string `json:"lang"`
	Token    string `json:"token"`
}

// /api/rpgreview
type RpgReviewC struct {
	Review   int    `json:"review"`
	SID      int    `json:"sid"`
	Region   string `json:"region"`
	Language string `json:"lang"`
	Token    string `json:"token"`
}
type RpgReviewS GenericS

// /api/infomercial (reporting)
type InfomercialC struct {
	SID    int    `json:"sid"`
	Info1  int    `json:"info1"`
	Info2  int    `json:"info2"`
	Info3  int    `json:"info3"`
	Info4  int    `json:"info4"`
	Info5  int    `json:"info5"`
	Info6  int    `json:"info6"`
	Region string `json:"region"`
	Lang   string `json:"lang"`
	Token  string `json:"token"`
	Text   string `json:"text"`
}
type InfomercialS GenericS

// /api/rgpupload (request is data)
type RpgUploadC struct {
	Title          string `json:"title"`
	Version        string `json:"version"`
	Lang           string `json:"lang"`
	PackageVersion int    `json:"packageversion"`
	Edit           int    `json:"edit"`
	Attribute      int    `json:"attribute"`
	Comment        string `json:"comment"`
	Owner          int    `json:"owner"`
	Crc32          int    `json:"crc32"`
	DataBlockSize  int    `json:"datablocksize"`
	Region         string `json:"region"`
	Token          string `json:"token"`

	Genre1  string `json:"genre1"`  // Fantasy
	Genre2  string `json:"genre2"`  // SF
	Genre3  string `json:"genre3"`  // School Life
	Genre4  string `json:"genre4"`  // Modern
	Genre5  string `json:"genre5"`  // Japanese
	Genre6  string `json:"genre6"`  // Adventure
	Genre7  string `json:"genre7"`  // Puzzle
	Genre8  string `json:"genre8"`  // Novel
	Genre9  string `json:"genre9"`  // Hunt
	Genre10 string `json:"genre10"` // Original
	Genre11 string `json:"genre11"` // Romance
	Genre12 string `json:"genre12"` // Training
	Genre13 string `json:"genre13"` // Riddlle
	Genre14 string `json:"genre14"` // Horror
	Genre15 string `json:"genre15"` // Mystery
	Genre16 string `json:"genre16"` // Classic
	Genre17 string `json:"genre17"` // Comical
	Genre18 string `json:"genre18"` // Serious
	Genre19 string `json:"genre19"` // Heartful
	Genre20 string `json:"genre20"` // Dark
	Genre21 string `json:"genre21"` // Children's
	Genre22 string `json:"genre22"` // Adult
	Genre23 string `json:"genre23"` // For Men
	Genre24 string `json:"genre24"` // For Women
	Genre25 string `json:"genre25"` // Short
	Genre26 string `json:"genre26"` // Long
	Genre27 string `json:"genre27"` // Easy
	Genre28 string `json:"genre28"` // Difficult
	Genre29 string `json:"genre29"` // Collabo.
	Genre30 string `json:"genre30"` // No Battle
	Genre31 string `json:"genre31"` // Mini Game
	Genre32 string `json:"genre32"` // Open Tech
	Genre33 string `json:"genre33"` // Movie NG
	Genre34 string `json:"genre34"` // Updated
}
type RpgUploadS GenericS

// /api/rpgdelete
type RpgDeleteC struct {
	Region   string `json:"region"`
	Language string `json:"lang"`
	Token    string `json:"token"`
	SID      int    `json:"sid"`
}
type RpgDeleteS GenericS
