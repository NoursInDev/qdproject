package tplinker

import (
	"log"
	"qdproject/extensions/global"
)

func NeedTP(x, y *int) {
	if *x == global.Data.Tp1x && *y == global.Data.Tp1y {
		tp1to2(x, y)
		log.Println(global.Data.TPNameExtension, ":", global.Data.Tp1x, global.Data.Tp1y, "to", global.Data.Tp2x, global.Data.Tp2y, "| 1 > 2")
	} else if *x == global.Data.Tp2x && *y == global.Data.Tp2y {
		tp2to1(x, y)
		log.Println(global.Data.TPNameExtension, ":", global.Data.Tp2x, global.Data.Tp2y, "to", global.Data.Tp1x, global.Data.Tp1y, "| 2 > 1")
	}
}

func tp1to2(x, y *int) {
	*x = global.Data.Tp2x
	*y = global.Data.Tp2y
}

func tp2to1(x, y *int) {
	*x = global.Data.Tp1x
	*y = global.Data.Tp1y
}
