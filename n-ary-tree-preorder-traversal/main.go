// NOT SOLVED
// NOT SOLVED
// NOT SOLVED

package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	stack := []*Node{}
	stack = append(stack, root)
	index := 0
	result := []int{}

	for {
		result = append(result, stack[index].Val)

		if len(stack[index].Children) > 0 {
			for k, v := range stack[index].Children {
				if k == 0 {
					if len(stack) > index+1 {
						stack[index+1] = v
					} else {
						stack = append(stack, v)
					}
					index++
					break
				}

				if len(stack) > index+1 {
					if stack[index+1] == v && len(stack[index].Children) > k+1 {
						stack[index+1] = stack[index].Children[k+1]
						index++
						break
					}
				}

			}
		} else {
			index--
		}

		if index == 0 {
			break
		}
	}

	return result
}

func main() {

}
