package common

import (
	"github.com/falcolee/transgo/common/utils/gologger"
)

const banner = `

_________ _______  _______  _        _______  _______  _______ 
\__   __/(  ____ )(  ___  )( (    /|(  ____ \(  ____ \(  ___  )
   ) (   | (    )|| (   ) ||  \  ( || (    \/| (    \/| (   ) |
   | |   | (____)|| (___) ||   \ | || (_____ | |      | |   | |
   | |   |     __)|  ___  || (\ \) |(_____  )| | ____ | |   | |
   | |   | (\ (   | (   ) || | \   |      ) || | \_  )| |   | |
   | |   | ) \ \__| )   ( || )  \  |/\____) || (___) || (___) |
   )_(   |/   \__/|/     \||/    )_)\_______)(_______)(_______)

`

func Banner() {
	gologger.Printf("%sBuilt At: %s\nGo Version: %s\nAuthor: %s\nTag: %s\nVersion: %s\n\n", banner, BuiltAt, GoVersion, GitAuthor, GitTag, version)
	gologger.Printf("\t\thttps://github.com/falcolee/transgo\n\n")
	gologger.Labelf("请勿用于非法用途，开发人员不承担任何责任，也不对任何滥用或损坏负责.\n")
}
