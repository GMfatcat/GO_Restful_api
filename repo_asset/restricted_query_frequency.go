package main

import (
	"fmt"
	"time"
)

var (
	userList      []string
	userQueries   map[string]*UserInfo
	blacklist     map[string]*BlacklistInfo
	maxQueries    = 10
	blacklistTime = 10 * time.Minute
)

// 定義使用者的查詢資訊
type UserInfo struct {
	QueryCount  int       // 查詢次數
	LastQueryAt time.Time // 最後一次查詢的時間
}

// 定義黑名單資訊
type BlacklistInfo struct {
	BlacklistedUntil time.Time // 黑名單結束時間
}

/* -------- Functions --------*/
func query() string {
	return "Query Success"
}

func isBlackListed(user string) bool {
	// check if user is blacklisted, only process when user is blacklisted
	info, ok := blacklist[user]
	if !ok {
		return false
	}
	// check if the time pass the blacklisttime
	// 使用 time.Before() 檢查是否在指定時間之前
	// 查看是否在blacklistuntil時間之前
	if time.Now().Before(info.BlacklistedUntil) {
		return true
	}
	// delete this info from blacklist since the time pass
	delete(blacklist, user)
	return false
}

func updateUserInfo(user string) {

	// get user info, add into userQueries if not exist
	info, ok := userQueries[user]
	if !ok {
		info = &UserInfo{QueryCount: 1, LastQueryAt: time.Now()}
		userQueries[user] = info
		fmt.Printf("%s : %s\n", user, query())
		return
	}
	// check if last query has already exceeded 1 minutes -> querycount = 1
	if time.Since(info.LastQueryAt) > time.Minute {
		info.QueryCount = 1
		info.LastQueryAt = time.Now()
		fmt.Printf("%s : %s\n", user, query())
		return
	}
	// check if querycount >= 10 --> if yes send into blacklist
	if info.QueryCount >= maxQueries {
		blacklist[user] = &BlacklistInfo{
			BlacklistedUntil: time.Now().Add(blacklistTime),
		}
		fmt.Printf("%s : Add to blacklist\n", user)
		return
	}

	info.QueryCount++
	info.LastQueryAt = time.Now()
	fmt.Printf("%s : %s\n", user, query())
}

func Simulation(userList []string) {
	// Loop
	for index, user := range userList {
		// check if blacklist, no query for blacklist
		isblacklist := isBlackListed(user)
		if isblacklist {
			fmt.Printf("User%d %s is in Blacklist\n", index, user)
			continue
		}
		updateUserInfo(user)
	}
}

func init() {
	userList = make([]string, 15) // 15 users
	userQueries = make(map[string]*UserInfo)
	blacklist = make(map[string]*BlacklistInfo)
}

func main() {
	// Get user list (all user is Tom)
	for i := range userList {
		userList[i] = "Tom"
	}
	// Simulation of Restriction Query Frequency
	Simulation(userList)
	// End of Simulation
	fmt.Println("End of Simulation")
}
