package apis

import (
	watcher "eternal-infer-worker/task-watcher"
	"eternal-infer-worker/types"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type APIResponse struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type Router struct {
	port    int
	watcher *watcher.TaskWatcher
}

func InitRouter(port int, watcher *watcher.TaskWatcher) *Router {
	return &Router{
		port:    port,
		watcher: watcher,
	}
}
func (rt *Router) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (rt *Router) StartRouter() error {
	log.Println("starting router")

	store := persistence.NewInMemoryStore(5 * time.Second)

	_ = store

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	corCfg := cors.DefaultConfig()
	corCfg.AllowCredentials = true
	corCfg.AllowOrigins = []string{"*"}
	corCfg.AllowHeaders = append(corCfg.AllowHeaders, "Authorization")
	r.Use(cors.New(corCfg))

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	apiv1 := r.Group("/api")

	apiv1.POST("/infer", rt.Infer)
	apiv1.GET("/tasks", rt.GetTasks)
	apiv1.POST("/submit", rt.SubmitTask)

	apiv1.GET("/health", rt.health)

	err := r.Run("0.0.0.0:" + fmt.Sprintf("%d", rt.port))
	if err != nil {
		return err
	}
	return nil
}

func (rt *Router) SubmitTask(c *gin.Context) {
	var task types.TaskSubmitRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	// if task.ModeID == "" {
	// 	c.JSON(http.StatusBadRequest, APIResponse{
	// 		Error:  "model id is required",
	// 		Status: http.StatusBadRequest,
	// 	})
	// 	return
	// }

	if task.ModelContract == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  "model contract is required",
			Status: http.StatusBadRequest,
		})
		return
	}

	if task.Params == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  "params is required",
			Status: http.StatusBadRequest,
		})
		return
	}

	if task.Value == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  "value is required",
			Status: http.StatusBadRequest,
		})
		return
	}

	log.Println("submit task", task)
	err := rt.watcher.SubmitTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "task assigned",
		Status:  http.StatusOK,
	})
}

func (rt *Router) Infer(c *gin.Context) {
	var task types.TaskInfo
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	// if task.ModelAddress == "" {
	// 	c.JSON(http.StatusBadRequest, APIResponse{
	// 		Error:  "model name is required",
	// 		Status: http.StatusBadRequest,
	// 	})
	// 	return
	// }
	if task.Params == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  "params is required",
			Status: http.StatusBadRequest,
		})
		return
	}

	err := rt.watcher.AssignTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "task assigned",
		Status:  http.StatusOK,
	})

}

func (rt *Router) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Data:   rt.watcher.GetAllTasks(),
		Status: http.StatusOK,
	})
}
