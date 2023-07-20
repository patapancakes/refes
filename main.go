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

package main

import (
	"flag"
	"log"
	"refes/api"
)

func main() {
	proto := flag.String("proto", "tcp", "protocol to use (\"tcp\", \"unix\", etc)")
	addr := flag.String("addr", "0.0.0.0:8100", "address to listen on")
	flag.Parse()

	err := api.Init(proto, addr)
	if err != nil {
		log.Fatalln(err)
	}
}
