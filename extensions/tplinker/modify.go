package tplinker

import (
	"log"
	"qdproject/extensions/global"
)

func Modify(x, y *int, mapcontent *[][]int) {

	if (*mapcontent)[*y][*x] > 76 && (*mapcontent)[*y][*x] < 88 && global.Data.Active {
		NeedTP(x, y)
	} else {
		tpnumber := (global.Data.NbTP + 1) % 2

		if global.Data.Active == true { // 2 tp déjà définis, remise du bloc initial
			switch {
			case tpnumber == 1:
				(*mapcontent)[global.Data.Tp1y][global.Data.Tp1x] = global.Data.Ptp1
			case tpnumber == 0:
				(*mapcontent)[global.Data.Tp2y][global.Data.Tp2x] = global.Data.Ptp2
			}
		}

		switch { // coordonnées du tp
		case tpnumber == 1:
			global.Data.Ptp1 = (*mapcontent)[*y][*x]
			global.Data.Tp1x = *x
			global.Data.Tp1y = *y
			log.Println(global.Data.TPNameExtension, ":", "posé en", *x, *y, "(entity - 1)")
		case tpnumber == 0:
			global.Data.Ptp2 = (*mapcontent)[*y][*x]
			global.Data.Tp2x = *x
			global.Data.Tp2y = *y
			log.Println(global.Data.TPNameExtension, ":", "posé en", *x, *y, "(entity - 2)")
		case tpnumber != 2 && tpnumber != 1:
			log.Println(global.Data.TPNameExtension, ":", "Erreur: tpnumber hors 1, 2")
		}
		global.Data.NbTP = global.Data.NbTP + 1 // maj nb tps

		(*mapcontent)[*y][*x] = 77 // mise du bloc tp sur fullcontent

		if global.Data.NbTP >= 2 { // définition activié
			global.Data.Active = true
		}
	}
}
