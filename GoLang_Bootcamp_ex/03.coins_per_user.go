/*
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus,
Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed
based on the vowels contained in each name where:
a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins
and a user can’t get more than 10 coins. Print a map with each user’s name
and the amount of coins distributed. After distributing all the coins, you should
have 2 coins left.

The output should be similar to:
Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2
*/

package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {

	letterValueMap := map[string]int{
		"A": 1,
		"a": 1,
		"E": 1,
		"e": 1,
		"I": 2,
		"i": 2,
		"O": 3,
		"o": 3,
		"U": 4,
		"u": 4}

	for _, n := range users {
		v := processName(n, letterValueMap)
		distribution[n] = v
		coins = coins - v
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}

func processName(n string, letterValueMap map[string]int) int {
	v := 0
	for i := 0; i < len(n); i++ {
		var ch = string(n[i])
		if val, ok := letterValueMap[ch]; ok {
			v += val
		}
	}

	if v > 10 {
		v = 10
	}

	return v
}
