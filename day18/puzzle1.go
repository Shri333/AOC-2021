package day18

import (
	"math"
	"strconv"
	"strings"
)

type pair struct {
	val    int
	left   *pair
	right  *pair
	parent *pair
}

type node struct {
	p     *pair
	depth int
}

func (p *pair) String() string {
	sb := &strings.Builder{}
	if p.left == nil {
		sb.WriteString(strconv.Itoa(p.val))
	} else {
		sb.WriteRune('[')
		sb.WriteString(p.left.String())
		sb.WriteRune(',')
		sb.WriteString(p.right.String())
		sb.WriteRune(']')
	}

	return sb.String()
}

func (p *pair) reduce() {
	reduced := true
	for reduced {
		reduced = false
		nodes := []*node{}
		p.inorder(&nodes, 0)
		for i, current := range nodes {
			if current.depth > 4 {
				explode(nodes, i)
				reduced = true
				break
			}
		}
		if !reduced {
			for _, current := range nodes {
				if current.p.val >= 10 {
					split(current.p)
					reduced = true
					break
				}
			}
		}
	}
}

func (p *pair) inorder(nodes *[]*node, depth int) {
	if p != nil {
		p.left.inorder(nodes, depth+1)
		if p.val != -1 {
			*nodes = append(*nodes, &node{p, depth})
		}
		p.right.inorder(nodes, depth+1)
	}
}

func (p *pair) magnitude() int {
	if p.left == nil {
		return p.val
	}

	return 3*p.left.magnitude() + 2*p.right.magnitude()
}

func puzzle1(numbers []string) int {
	temp := 0
	ptr := &temp
	root := parse(numbers[0], ptr)

	for i := 1; i < len(numbers); i++ {
		*ptr = 0
		left, right := root, parse(numbers[i], ptr)
		root = &pair{-1, left, right, nil}
		left.parent, right.parent = root, root
		root.reduce()
	}

	return root.magnitude()
}

func parse(number string, i *int) *pair {
	p := &pair{-1, nil, nil, nil}

	if number[*i] == '[' {
		*i++
		left := parse(number, i)
		left.parent, p.left = p, left
	}

	if number[*i] == ',' {
		*i++
		right := parse(number, i)
		right.parent, p.right = p, right
	}

	if number[*i] == ']' {
		*i++
		return p
	}

	start := *i
	for number[*i] != '[' && number[*i] != ',' && number[*i] != ']' {
		*i++
	}

	p.val, _ = strconv.Atoi(number[start:*i])
	return p
}

func explode(nodes []*node, i int) {
	if i > 0 {
		prev, left := nodes[i-1], nodes[i]
		prev.p.val += left.p.val
	}

	if i < len(nodes)-2 {
		right, next := nodes[i+1], nodes[i+2]
		next.p.val += right.p.val
	}

	parent := nodes[i].p.parent
	parent.val, parent.left, parent.right = 0, nil, nil
}

func split(p *pair) {
	half := float64(p.val) / 2
	leftVal := int(math.Floor(half))
	rightVal := int(math.Ceil(half))
	p.left = &pair{leftVal, nil, nil, p}
	p.right = &pair{rightVal, nil, nil, p}
	p.val = -1
}
