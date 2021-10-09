package Problems

import (
	"strings"
)

/*
The canonical path should have the following format:

The path starts with a single slash '/'.
Any two directories are separated by a single slash '/'.
The path does not end with a trailing '/'.
The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
Return the simplified canonical path.



Example 1:

Input: path = "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: path = "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: path = "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
Example 4:

Input: path = "/a/./b/../../c/"
Output: "/c"
*/

type Stack struct {
	arr []string
}

func (s *Stack) pop() string {
	ret := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ret
}

func (s *Stack) push(n string) {
	s.arr = append(s.arr, n)
}

func (s *Stack) Len() int {
	return len(s.arr)
}
func simplifyPath(path string) string {

	stack := &Stack{arr: []string{}}
	arr := strings.Split(path, "/")

	for _, v := range arr {
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if stack.Len() != 0 {
				stack.pop()
			}
			break
		default:
			stack.push(v)
		}
	}

	return strings.Join(stack.arr, "/")
}
