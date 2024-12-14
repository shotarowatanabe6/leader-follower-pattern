package handler

import (
	repository "leader-follower-pattern/domain/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	InMemoryRepo repository.IMemoryRepository
	RedisRepo    repository.IRedisRepository
}

func NewHandler() Handler {
	return Handler{
		InMemoryRepo: repository.NewInMemoryRepository(),
		RedisRepo:    repository.NewRedisRepository(),
	}
}

func (h *Handler) Run() {
	r := gin.Default()
	r.POST("/set", h.Set)
	r.GET("/get", h.Get)
	r.Run()
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (h *Handler) Set(c *gin.Context) {
	var reqs []SetRequest
	err := c.BindJSON(&reqs)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		err := h.RedisRepo.Set(c, req.Key, req.Value)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func (h *Handler) Get(c *gin.Context) {
	key := c.Query("key")
	value, err := h.RedisRepo.Get(c, key)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"value": value})
}
