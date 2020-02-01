package trie

type TreeNode struct {
	isWord bool
	next map[rune]*TreeNode
}

type TrieTree struct {
	root *TreeNode
	Size int
}

func NewTreeNode() *TreeNode {
	node := new(TreeNode)
	node.isWord = false
	node.next = make(map[rune]*TreeNode, 26)
	return node
}

func InitTrieTree() *TrieTree {
	return &TrieTree{root: NewTreeNode(), Size: 0}
}

func BuildTrieTree(words []string) *TrieTree {
	tree := InitTrieTree()
	for _, word := range words {
		tree.Add(word)
	}
	return tree
}

// 添加单词 word
func (tree *TrieTree) Add(word string) {
	current := tree.root
	chars := []rune(word)
	for _, v := range chars {
		next := current.next[v]
		if next == nil {
			// 该路径不存在
			current.next[v] = NewTreeNode()
		}
		current = current.next[v]
	}
	if !current.isWord {
		// 不存在这个单词
		// 则设置为 word
		tree.Size++
		current.isWord = true
	}
}

// 查询单词 word 是否存在
func (tree *TrieTree) Contains(word string) bool {
	current := tree.root
	chars := []rune(word)
	for _, v := range chars {
		next := current.next[v]
		if next == nil {
			// 该路径不存在
			return false
		}
		current = next
	}
	// 判断是前缀还是单词
	return current.isWord
}

// 查询前缀 prefix 是否存在
func (tree *TrieTree) ContainsPrefix(prefix string) bool {
	current := tree.root
	chars := []rune(prefix)
	for _, v := range chars {
		next := current.next[v]
		if next == nil {
			// 该路径不存在
			return false
		}
		current = next
	}
	return true
}