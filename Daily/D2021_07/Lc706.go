package D2021_07

import (
	"sort"
	"strconv"
)

/*
1418. 点菜展示表
给你一个数组 orders，表示客户在餐厅中完成的订单，确切地说， orders[i]=[customerNamei,tableNumberi,foodItemi] ，
其中 customerNamei 是客户的姓名，tableNumberi 是客户所在餐桌的桌号，而 foodItemi 是客户点的餐品名称。
请你返回该餐厅的 点菜展示表 。在这张表中，表中第一行为标题，其第一列为餐桌桌号 “Table” ，后面每一列都是按字母顺序排列的餐品名称。
接下来每一行中的项则表示每张餐桌订购的相应餐品数量，第一列应当填对应的桌号，后面依次填写下单的餐品数量。
注意：客户姓名不是点菜展示表的一部分。此外，表中的数据行应该按餐桌桌号升序排列。
 */

/*
思路：
一个map存菜品名字，一个map套小map存客户点单
 */

func Lc706DisplayTable(orders [][]string) [][]string {
	dishes := make(map[string]struct{})
	custom := make(map[int]map[string]int)
	for _, order := range orders {
		dishes[order[2]] = struct{}{}
		id, _ := strconv.Atoi(order[1])
		if custom[id] == nil {
			custom[id] = make(map[string]int)
		}
		custom[id][order[2]]++
	}
	dishesSlice := make([]string, 0)
	for dish := range dishes {
		dishesSlice = append(dishesSlice, dish)
		sort.Strings(dishesSlice)
	}
	ids := make([]int, 0)
	for id := range custom {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	ans := make([][]string, 0)
	tableHead := make([]string, 1)
	tableHead[0] = "Table"
	for _, dish := range dishesSlice {
		tableHead = append(tableHead, dish)
	}
	ans = append(ans, tableHead)
	for _, id := range ids {
		tableLine := make([]string, 0)
		tableLine = append(tableLine, strconv.Itoa(id))
		for _, dish := range dishesSlice {
			tableLine = append(tableLine, strconv.Itoa(custom[id][dish]))
		}
		ans = append(ans, tableLine)
	}
	return ans
}