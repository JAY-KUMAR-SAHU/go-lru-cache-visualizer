package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var cacheCapacity int = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

type Cache struct {
	Queue Queue
	Hash  map[string]*Node
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: make(map[string]*Node)}
}

func (c *Cache) Remove(n *Node) {
	left := n.Left
	right := n.Right
	left.Right = right
	right.Left = left
	c.Queue.Length--
	delete(c.Hash, n.Val)
}

func (c *Cache) Add(n *Node) {
	r := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = r
	r.Left = n
	c.Queue.Length++
	if c.Queue.Length > cacheCapacity {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Check(word string) {
	var node *Node
	if val, ok := c.Hash[word]; ok {
		c.Remove(val)
		node = val
	} else {
		node = &Node{Val: word}
	}
	c.Add(node)
	c.Hash[word] = node
}

func (c *Cache) Display() []string {
	result := []string{}
	node := c.Queue.Head.Right
	for i := 0; i < c.Queue.Length; i++ {
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

func addWord(c echo.Context) error {
	word := c.Param("word")
	if word == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Word is required"})
	}
	cache.Check(word)

	return c.JSON(http.StatusOK, map[string]string{"message": "Word added successfully"})
}

func getCacheState(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cache":    cache.Display(),
		"capacity": cacheCapacity,
		"length":   cache.Queue.Length,
	})
}

func setCacheSize(c echo.Context) error {
	sizeParam := c.Param("size")
	size, err := strconv.Atoi(sizeParam)
	if err != nil || size <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cache size"})
	}
	cacheCapacity = size

	if cache.Queue.Length > cacheCapacity {
		for cache.Queue.Length > cacheCapacity {
			cache.Remove(cache.Queue.Tail.Left)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cache capacity updated successfully"})
}

var cache = NewCache()

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/add/:word", addWord)
	e.GET("/cache", getCacheState)
	e.POST("/set-cache-size/:size", setCacheSize)
	e.Logger.Fatal(e.Start(":8080"))
}
