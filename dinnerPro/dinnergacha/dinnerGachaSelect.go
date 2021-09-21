package dinnergacha

import (
	"math/rand"
	"time"
)

func DinnerMenuGacha() (string, string) {

	// メニューの設定の為の乱数作成
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10) + 1

	// 和食の場合
	if x >= 1 && x <= 4 {

		menuW := menu[0]
		wList := washokuList()

		return menuW, wList

		// 洋食の場合
	} else if x >= 5 && x <= 7 {

		menuY := menu[1]
		yList := youshokuList()

		return menuY, yList

		// 中華の場合
	} else {

		menuC := menu[2]
		cList := chuukaList()

		return menuC, cList

	}

}

// 和食の場合設定
func washokuList() string {

	// washokuMenuの乱数作成
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(3) + 1

	washoku := washokuMenu[x]

	return washoku

}

// 洋食の場合設定
func youshokuList() string {

	// youshokuMenuの乱数作成
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(3) + 1

	youshoku := youshokuMenu[x]

	return youshoku

}

// 中華の場合設定
func chuukaList() string {

	// chuukaMenuの乱数作成
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(3) + 1

	chuuka := chuukaMenu[x]

	return chuuka

}
