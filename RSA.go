package main

import (
	"fmt"
	"math"
)

func GenerateKey(p, q int) [2][2]int {
	var n = q * p
	var e = 7
	var tmp = (q - 1) * (p - 1)
	var d int
	for d = 1; ; d++ {
		if e*d%tmp == 1 {
			break
		}
	}
	return [2][2]int{{n, e}, {n, d}} // 销毁 p、q。生成公钥私钥
}

func Encrypt(x, e, n int) int {
	var result = x % n
	var i = 1
	for ; i < e; i++ {
		result = (result * x) % n
	}
	return result
}

func Decrypt(x, d, n int) int {
	var result = x % n
	var i = 1
	for i = 1; i < d; i++ {
		result = (result * x) % n
	}
	return result
}

func modpow(x, p, m int) int {
	if p == 1 {
		return int(math.Mod(float64(x), float64(m)))
	}
	var mid int
	if (p & 1) == 0 {
		mid = p >> 1
		var tmp1 = modpow(x, mid, m)
		return int(math.Mod(float64(tmp1*tmp1), float64(m)))
	} else {
		return int(math.Mod(float64(modpow(x, p-1, m)*x), float64(m)))
	}
}

func main() {
	var p = 13
	var q = 17
	tmp := GenerateKey(p, q)
	pubKey, selfKey := tmp[0], tmp[1]
	var m = 11
	secret := modpow(m, pubKey[1], pubKey[0])
	fmt.Println(secret)
	words := modpow(secret, selfKey[1], selfKey[0])
	fmt.Println(words)
}
