package D2021_11

/*
677. 键值映射
实现一个 MapSum 类，支持两个方法，insert 和 sum：
MapSum() 初始化 MapSum 对象。
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。
如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
 */

type Lc1114MapSum struct {
	value int
	children [26]*Lc1114MapSum
}

func Lc1114Constructor() Lc1114MapSum {
	return *new(Lc1114MapSum)
}

func (this *Lc1114MapSum) Lc1114Insert(key string, val int)  {
	if len(key) == 0 {
		this.value = val
		return
	}
	if this.children[key[0] - 'a'] == nil {
		newNode := Lc1114Constructor()
		this.children[key[0] - 'a'] = &newNode
	}
	this.children[key[0] - 'a'].Lc1114Insert(string([]byte(key)[1:]), val)
}

func (this *Lc1114MapSum) Lc1114Sum(prefix string) int {
	if len(prefix) == 0 {
		sum := this.value
		for _, child := range this.children {
			if child != nil {
				sum += child.Lc1114Sum("")
			}
		}
		return sum
	}
	if this.children[prefix[0] - 'a'] == nil {
		return 0
	} else {
		return this.children[prefix[0] - 'a'].Lc1114Sum(string([]byte(prefix)[1:]))
	}
}
