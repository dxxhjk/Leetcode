package D2021_05

/*
1734. 解码异或后的排列
给你一个整数数组 perm ，它是前 n 个正整数的排列，且 n 是个奇数。
它被加密成另一个长度为 n - 1 的整数数组 encoded，满足 encoded[i] = perm[i] XOR perm[i + 1]。
比方说，如果 perm = [1,3,2] ，那么 encoded = [2,1]。
给你 encoded 数组，请你返回原始数组 perm 。题目保证答案存在且唯一。
 */

/*
思路：
利用 n 为奇数，
perm[1] ^ perm[2] = encoded[1]， 
perm[3] ^ perm[4] = encoded[3]，
...
perm[n - 2] ^ perm[n - 1] = encoded[n - 2]
只要将 encoded 的偶数位都异或起来就可以获得 perm[1] 到 perm[n - 1]的异或，
用该值与 1 到 n 全部数字的异或做异或，即可获得 perm[0].
用 perm[0] ^ encoded[0] = perm[1],
perm[1] ^ encoded[1] = perm[2]
...
即可递推出 perm 数列。
 */

func Lc511Decode(encoded []int) []int {
	sumWithoutFirst := 0
	sum := 0
	for i := 1; i < len(encoded); i += 2 {
		sumWithoutFirst ^= encoded[i]
	}
	for i := 1; i <= len(encoded) + 1; i++ {
		sum ^= i
	}
	perm := make([]int, 0)
	perm = append(perm, sum ^ sumWithoutFirst)
	for i := 0; i < len(encoded); i++ {
		perm = append(perm, perm[i] ^ encoded[i])
	}
	return perm
}