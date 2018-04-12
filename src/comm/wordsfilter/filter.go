/*
* @Author: huang
* @Date:   2018-04-12 15:03:44
* @Last Modified by:   huang
* @Last Modified time: 2018-04-12 15:05:35
 */
package wordsfilter

import (
	"io/ioutil"
	"strings"
)

// ============================================================================

var (
	root = create_node()
)

// ============================================================================

type trie_node_t struct {
	children map[rune]*trie_node_t
	end      bool
}

// ============================================================================

func Load(fn string) error {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(buf), "\n") {
		line = strings.Trim(line, " \r")
		if line != "" {
			add(line)
		}
	}

	return nil
}

func IsSensitive(v string) bool {
	chars := []rune(strings.ToLower(v))
	L := len(chars)

	for i := 0; i < L; i++ {
		node := root

		for j := i; j < L; j++ {
			char := chars[j]

			child := node.children[char]
			if child == nil {
				break
			} else {
				if child.end {
					return true
				}
			}

			node = child
		}
	}

	return false
}

func Filter(v string) string {
	chars_ret := []rune(v)
	chars := []rune(strings.ToLower(v))
	L := len(chars)

	for i := 0; i < L; i++ {
		node := root
		pos := -1

		for j := i; j < L; j++ {
			char := chars[j]

			child := node.children[char]
			if child == nil {
				break
			} else {
				if child.end {
					pos = j // remember and keep searching for longest match
				}
			}

			node = child
		}

		if pos >= 0 {
			for k := i; k <= pos; k++ {
				chars_ret[k] = '*'
			}
			i = pos
		}
	}

	return string(chars_ret)
}

// ============================================================================

func create_node() *trie_node_t {
	return &trie_node_t{
		children: make(map[rune]*trie_node_t),
		end:      false,
	}
}

func add(v string) {
	chars := []rune(strings.ToLower(v))
	node := root

	for _, char := range chars {
		child := node.children[char]
		if child == nil {
			child = create_node()
			node.children[char] = child
		}

		node = child
	}

	node.end = true
}
