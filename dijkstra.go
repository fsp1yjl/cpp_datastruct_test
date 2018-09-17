package main

import (
	"fmt"
)

type Vertex struct {
	Name string
}

type Edge struct {
	V_1    Vertex
	V_2    Vertex
	Weight int
}

type Graph struct {
	Nodes map[string]*Vertex
	Edges map[Vertex]map[Vertex]int
}

var (
	MAX_NUMBER = 99999
)

func (g *Graph) addNode(v *Vertex) {
	if g.Nodes == nil {
		g.Nodes = make(map[string]*Vertex, 0)
	}
	g.Nodes[v.Name] = v
}

// todo 后期增加v1, v2是否存在的判断， 以及edge是否已经存在的判断，这里默认不重复，且顶点存在
func (g *Graph) addEdge(e *Edge) {
	v1 := e.V_1
	v2 := e.V_2

	if g.Edges == nil {
		g.Edges = make(map[Vertex]map[Vertex]int, 0)
	}

	if g.Edges[v1] == nil {
		g.Edges[v1] = make(map[Vertex]int, 0)
	}
	if _, ok := g.Edges[v1][v2]; !ok {
		// todo 如果v1->v2 的边已经存在，需要做报错处理
		g.Edges[v1][v2] = e.Weight
	}

	if g.Edges[v2] == nil {
		g.Edges[v2] = make(map[Vertex]int, 0)
	}
	if _, ok := g.Edges[v2][v1]; !ok {
		// todo 如果v2->v1 的边已经存在，需要做报错处理
		g.Edges[v2][v1] = e.Weight
	}
}

func (g *Graph) initNode(nodes []string) {
	for _, v := range nodes {
		tempNode := &Vertex{Name: v}
		g.addNode(tempNode)
	}
}

func (g *Graph) initEdges(edges []*Edge) {
	for _, e := range edges {
		g.addEdge(e)
	}
}

func (g *Graph) printNodes() {

	for _, p_v := range g.Nodes {
		fmt.Printf("vertex: %+v \n", *p_v)
		for v2, weight := range g.Edges[*p_v] {
			fmt.Printf("%s -> %s : %d \n", p_v.Name, v2.Name, weight)
		}
		// fmt.Println(g.Edges[*p_v])
	}
}

// func (g *Graph) singleEdge(v1 string, v2 string, Weight int) {
// 	tempEdge := &Edge{V_1: *g.Nodes["a"], V_2: *g.Nodes["b"], Weight: 1}

// }

func (g *Graph) dijkstra(start *Vertex) {
	visitedMap := make(map[string]*Vertex)
	unvisitedMap := make(map[string]*Vertex)

	tempShortestMap := make(map[Vertex]int)

	//初始化未遍历节点
	for _, p_v := range g.Nodes {
		if p_v.Name == start.Name {
			continue
		}
		unvisitedMap[p_v.Name] = p_v
	}

	//初始化已遍历节点
	visitedMap[start.Name] = start

	// 初始化 贪心最短map
	for _, p_v := range g.Nodes {
		if p_v.Name == start.Name {
			tempShortestMap[*p_v] = 0
		} else {
			if weight, ok := g.Edges[*start][*p_v]; ok {
				tempShortestMap[*p_v] = weight
			} else {
				tempShortestMap[*p_v] = MAX_NUMBER
			}
		}
	}

	for len(unvisitedMap) != 0 {

		fmt.Printf("unvisited: %+v\n", unvisitedMap)
		// 寻找tempShortestMap中的当前未访问最小元素
		minVertex := g.findMin(tempShortestMap, unvisitedMap)

		visitedMap[minVertex.Name] = minVertex
		delete(unvisitedMap, minVertex.Name)

		tempWeight := tempShortestMap[*minVertex]

		// 更新最短路劲
		for _, p_v := range unvisitedMap {

			fmt.Println("1111111111111:", g.Edges[*minVertex])
			fmt.Println("2222222222222:", p_v.Name)
			if w, ok := g.Edges[*minVertex][*p_v]; ok {
				if w+tempWeight < tempShortestMap[*p_v] {
					tempShortestMap[*p_v] = w + tempWeight
				}
			}
		}
	}

	fmt.Println("ddddddd:", tempShortestMap)
}

func (g *Graph) findMin(tempShortestMap map[Vertex]int, unvisitedMap map[string]*Vertex) *Vertex {
	tempMinKey := ""
	var tempMinNum int
	for k, p_v := range unvisitedMap {
		if tempMinKey == "" {
			tempMinKey = k
			tempMinNum = tempShortestMap[*p_v]
		} else {
			if tempShortestMap[*p_v] < tempMinNum {
				tempMinKey = k
				tempMinNum = tempShortestMap[*p_v]
			}
		}
	}
	fmt.Println("key::::", tempMinKey)

	fmt.Println("min::", g.Nodes[tempMinKey])
	return g.Nodes[tempMinKey]
}

func (g *Graph) initShortestMap(tempShortestMap map[Vertex]int, start string) {
	// s_v := *g.Nodes[start]
	// if g.Edges[s_v]
}

func main() {
	nodes := []string{"a", "b", "c", "d", "e", "f"}
	g := Graph{}
	g.initNode(nodes)

	edges := make([]*Edge, 0)

	tempEdge := &Edge{V_1: *g.Nodes["a"], V_2: *g.Nodes["b"], Weight: 6}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["a"], V_2: *g.Nodes["d"], Weight: 1}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["b"], V_2: *g.Nodes["c"], Weight: 5}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["b"], V_2: *g.Nodes["d"], Weight: 2}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["b"], V_2: *g.Nodes["e"], Weight: 2}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["c"], V_2: *g.Nodes["e"], Weight: 5}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["c"], V_2: *g.Nodes["f"], Weight: 1}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["d"], V_2: *g.Nodes["e"], Weight: 1}
	edges = append(edges, tempEdge)

	tempEdge = &Edge{V_1: *g.Nodes["e"], V_2: *g.Nodes["f"], Weight: 1}
	edges = append(edges, tempEdge)

	g.initEdges(edges)

	// fmt.Println(g)
	g.printNodes()

	g.dijkstra(g.Nodes["b"])

}
