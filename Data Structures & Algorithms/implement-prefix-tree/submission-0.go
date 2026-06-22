type PrefixTree struct {
	isWordEnd bool
	branch [26]*PrefixTree
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	for _, c := range word {
		index := c-'a'
		if this.branch[index] == nil {
			this.branch[index] = &PrefixTree{}
		}
		this = this.branch[index]
	}
	this.isWordEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	for _, c := range word {
		index := c-'a'
		if this.branch[index] == nil {
			return false
		}
		this = this.branch[index]
	}
	return this.isWordEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for _, c := range prefix {
		index := c-'a'
		if this.branch[index] == nil {
			return false
		}
		this = this.branch[index]
	}
	return true
}
