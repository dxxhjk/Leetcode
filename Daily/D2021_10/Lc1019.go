package D2021_10

/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
实现词典类 WordDictionary ：
	WordDictionary() 初始化词典对象
	void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
	bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false。
	word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
 */

type WordDictionary struct {
	letters [26]*WordDictionary
	finish bool
}

func Constructor1019() WordDictionary {
	dic := new(WordDictionary)
	dic.letters = *new([26]*WordDictionary)
	dic.finish = false
	return *dic
}

func (this *WordDictionary) AddWord1019(word string)  {
	if len(word) == 0 {
		this.finish = true
		return
	}
	ch := word[0] - 'A'
	if this.letters[ch] == nil {
		node := Constructor1019()
		this.letters[ch] = &node
	}
	this.letters[ch].AddWord1019(word[1:])
}

func (this *WordDictionary) Search1019(word string) bool {
	if len(word) == 0 {
		return this.finish
	}
	flag := false
	if word[0] == '.' {
		for i := 0; i < 26; i++ {
			if this.letters[i] != nil {
				flag = flag || this.letters[i].Search1019(word[1:])
			}
		}
	} else {
		ch := word[0] - 'A'
		if this.letters[ch] != nil {
			flag = flag || this.letters[ch].Search1019(word[1:])
		}
	}
	return flag
}
