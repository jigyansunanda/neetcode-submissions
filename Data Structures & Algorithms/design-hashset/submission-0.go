type MyHashSet struct {
	count [1000001]int
}

func Constructor() MyHashSet {
    hashMap := MyHashSet{}
	return hashMap
}

func (this *MyHashSet) Add(key int) {
    this.count[key] = 1
}

func (this *MyHashSet) Remove(key int) {
    this.count[key] = 0
}

func (this *MyHashSet) Contains(key int) bool {
    return this.count[key] > 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 