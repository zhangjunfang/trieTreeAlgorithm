package main

import (
    "bytes"
    "fmt"
)

type node struct {
    Word []byte  //保存一个单词
    Children []*node //保存该单词的后继节点

    //表明从根到这个节点所经过的所有字母组合起来是否构成了一个单词
    //譬如我添加了hello这个单词，查找hel，exist则为false，说明不存在这个单词
    Exist bool
}

type TrieTree struct {
    Root node
}

func main() {
    trie := new(TrieTree)
    trie.Add("hello")
    fmt.Println(trie.Search("hell"))
    fmt.Println(trie.Search("hello2"))

}

func (tree *TrieTree) Search(word string) bool {
    current := &tree.Root
    chars := []byte(word)
    for i := 0; i < len(chars); i++ {
        s := chars[i : i+1]
        nodes := current.Children
        index, exist := tree.BinarySearch(nodes, s)
        if !exist {
            //fmt.Println(i, s)
            return false
        }
        current = nodes[index]
    }

    return current.Exist
}
func (tree *TrieTree) Add(word string) {
    //把一个单词增加到当前的前缀树中
    current := &tree.Root
    chars := []byte(word)
    //把单词的每个字母都添加到前缀树中去，（有则忽略掉）
    for i := 0; i < len(chars); i++ {
        s := chars[i : i+1]
        //fmt.Println(s)
        nodes := &current.Children
        index, exist := tree.BinarySearch(*nodes, s)
        //fmt.Println(index, exist)
        if !exist {
            //这个字母在前缀树中没有则添加进去
            //fmt.Println("index:", index)
            *nodes = append(*nodes, nil)

            copy((*nodes)[index+1:], (*nodes)[index:])
            (*nodes)[index] = &node{Word: s, Exist: false}
            //fmt.Println("insert:", (*nodes)[0].Word)

        }

        current.Children = *nodes
        current = (*nodes)[index]

    }
    current.Exist = true //单词添加完毕，末位节点设为true表示这个从跟到这个节点的路线所构成的单词存在

}

//二分查找。譬如字母a这个节点是第一个节点，则其子节点，包含了所有以a开头的单词
//该二分查找，就是找前缀树里，以a开头的单词的第二个字母是否含有s所代表的那个字母
func (tree *TrieTree) BinarySearch(nodes []*node, s []byte) (int, bool) {
    start := 0
    end := len(nodes) - 1

    if end == -1 {
        return 0, false
    }
    compareFirst := bytes.Compare(s, nodes[0].Word)

    if compareFirst < 0 {
        //说明不存在
        fmt.Println("compare first:", s, nodes[0].Word)
        return 0, false
    } else if compareFirst == 0 {
        return 0, true
    }
    compareLast := bytes.Compare(s, nodes[end].Word)
    if compareLast > 0 {
        //说明不存在
        return end + 1, false
    } else if compareLast == 0 {
        return end, true
    }

    current := end / 2

    //为什么是end-start>1 而不是>0呢
    if end-start > 1 {
        compareCurrent := bytes.Compare(s, nodes[current].Word)
        if compareCurrent > 0 {
            start = current
            current = (end + start) / 2
        } else {
            end = current
            current = (end + start) / 2
        }
    }

    return end, true

}
