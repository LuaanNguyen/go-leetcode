package find_eventual_safe_states


func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    safe := make(map[int]bool)

    res := make([]int, 0)
    for i := range n {
        if DFS(i, safe, graph) {
            res = append(res, i)
        }
    }
    return res
}

func DFS(i int, s map[int]bool, g [][]int ) bool {
    if _, exists := s[i]; exists {
        return s[i]
    }
    s[i] = false
    for _, nei := range g[i] {
        if !DFS(nei, s, g) {
            return false
        }
    }
    s[i] = true
    return true
}