type MyHashMap struct {
	count [1000001]int
}

func Constructor() MyHashMap {
    hashMap := MyHashMap{}
	for i := range 1000001 {
		hashMap.count[i] = -1
	}
	return hashMap
}

func (this *MyHashMap) Put(key int, value int) {
    this.count[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.count[key]
}

func (this *MyHashMap) Remove(key int) {
    this.count[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */