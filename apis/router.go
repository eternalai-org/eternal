package apis

import (
	"eternal-infer-worker/config"
	watcher "eternal-infer-worker/task-watcher"
	"eternal-infer-worker/types"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	version string
}

func InitRouter(port int, watcher *watcher.TaskWatcher, version string) *Router {
	return &Router{
		version: version,
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

	apiv1.GET("/chains", rt.GetChains)

	apiv1.POST("/infer", rt.Infer)
	apiv1.GET("/tasks", rt.GetTasks)
	apiv1.POST("/submit", rt.SubmitTask)

	apiv1.GET("/health", rt.health)
	apiv1.GET("/stats", rt.Stats)

	apiv1.GET("/claim-reward", rt.ClaimReward)

	apiv1.GET("/unstake", rt.Unstake)
	apiv1.GET("/claim-unstake", rt.ClaimUnstake)

	err := r.Run("0.0.0.0:" + fmt.Sprintf("%d", rt.port))
	if err != nil {
		return err
	}
	return nil
}

func (rt *Router) Stats(c *gin.Context) {
	type Stats struct {
		AssignedModel  string `json:"assigned_model"`
		ModelStatus    string `json:"model_status"`
		WorkerAddress  string `json:"worker_address"`
		WorkerBalance  string `json:"worker_balance"`
		StakeStatus    string `json:"stake_status"`
		StakedAmount   string `json:"staked_amount"`
		SessionEarning string `json:"session_earning"`
		ProcessedTasks uint64 `json:"processed_tasks"`
		MiningReward   string `json:"mining_reward"`
		UnstakeAmount  string `json:"unstake_amount"`
		UnstakeTime    string `json:"unstake_time"`
		TotalModels    int    `json:"total_models"`
		TotalMiners    int    `json:"total_miners"`
		Version        string `json:"version"`
		HasNewVersion  bool   `json:"has_new_version"`
		NewVersion     string `json:"new_version"`
	}

	unstakeAmount, unstakeTime := rt.watcher.GetUnstakeInfo()

	globalInfo, err := rt.watcher.GetHubGlobalInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	stakedStatus, err := rt.watcher.GetStakeStatus()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	hasNewVersion, newVersion := rt.watcher.HasNewVersion()

	stats := Stats{
		AssignedModel:  rt.watcher.GetAssignedModel(),
		ModelStatus:    rt.watcher.GetModelStatus(),
		WorkerAddress:  rt.watcher.GetWorkerAddress(),
		WorkerBalance:  rt.watcher.GetWorkerBalance(),
		StakeStatus:    stakedStatus,
		StakedAmount:   rt.watcher.GetStakedAmount(),
		SessionEarning: rt.watcher.GetSessionEarning(),
		ProcessedTasks: rt.watcher.GetProcessedTasks(),
		MiningReward:   rt.watcher.GetMiningReward(),
		UnstakeAmount:  unstakeAmount,
		UnstakeTime:    unstakeTime.String(),
		TotalModels:    int(globalInfo.TotalModels),
		TotalMiners:    int(globalInfo.TotalMiners),
		Version:        rt.version,
		HasNewVersion:  hasNewVersion,
		NewVersion:     newVersion,
	}

	c.JSON(http.StatusOK, APIResponse{
		Data:   stats,
		Status: http.StatusOK,
	})

	// var s string
	// s = fmt.Sprintf("[Eternal] [%v] [%v]\n\n", m.version, m.nodeMode)
	// // If the spinner is showing, render it and bail.

	// if m.expandLog {
	// 	s += "Logs: ('l' to collapse)------------------ \n\n"
	// 	for _, l := range m.logLines {
	// 		s += logStyle(l) + "\n"
	// 	}
	// } else {
	// 	s += "Logs: ('l' to expand)------------------ \n\n"
	// 	s += logStyle(m.logLines[len(m.logLines)-1]) + "\n"
	// }

	// s += "--------------------------------------- \n"

	// if !m.expandLog {
	// 	//status line
	// 	s += fmt.Sprintf("\n %s Status:%s%s\n", m.statusSpinner.View(), " ", textStyle(m.statusText))

	// 	if m.taskManager != nil {

	// 		tasks := m.taskManager.GetCurrentRunningTasks()
	// 		taskIDs := make([]string, 0, len(tasks))
	// 		for _, v := range tasks {
	// 			taskIDs = append(taskIDs, v.TaskID)
	// 		}

	// 		s += fmt.Sprintf("\n %s Current processing (%v):%s%s\n", ">", len(taskIDs), " ", textStyle(fmt.Sprintf("%v", taskIDs)))

	// 		modelLoadingSpinner := ""
	// 		if m.modelManager.GetStatus() != "loaded" {
	// 			modelLoadingSpinner = m.loadingSpinner.View()
	// 		}
	// 		s += fmt.Sprintf("\n %s Assigned Model:%s%s | ModelManager: %s %s\n", ">", " ", textStyle(m.taskManager.GetAssignedModel()), textStyle(m.modelManager.GetStatus()), modelLoadingSpinner)
	// 		s += fmt.Sprintf("\n %s Stake Status:%s%s (%s EAI)\n", ">", " ", textStyle(m.taskManager.StakeStatus()), textStyle(m.taskManager.GetStakedAmount()))
	// 		if m.taskManager.StakeStatus() == "staked" {
	// 			s += bindingShortCutStyle("   Press 'ctrl+u' to unstake and quit\n")
	// 		}

	// 		s += fmt.Sprintf("\n %s Worker Address:%s%s\n", ">", " ", textStyle(m.taskManager.GetWorkerAddress()))
	// 		s += fmt.Sprintf("\n %s Balance:%s%s EAI", ">", " ", textStyle(m.taskManager.GetWorkerBalance()))
	// 		s += fmt.Sprintf(" %s Session Earning:%s%s EAI (%s tasks)\n", "|", " ", textStyle(fmt.Sprintf("%v", m.taskManager.GetSessionEarning())), textStyle(fmt.Sprintf("%v", m.taskManager.GetProcessedTasks())))

	// 		miningAmount := m.taskManager.GetMiningReward()
	// 		if miningAmount != "0" {
	// 			s += fmt.Sprintf("\n %s Mining reward:%s%s EAI\n", ">", " ", textStyle(miningAmount))
	// 			s += bindingShortCutStyle("   Press 'e' to claim mining reward\n")
	// 		}

	// 		unstakeAmount, unstakeTime := m.taskManager.GetUnstakeInfo()
	// 		if unstakeAmount != "0" {
	// 			s += fmt.Sprintf("\n %s Pending unstake:%s%s EAI available to claim at %s\n", ">", " ", textStyle(unstakeAmount), textStyle(unstakeTime.Format("2006-01-02 15:04:05")))
	// 			s += "   Press 'ctrl+r' to reclaim stake\n"
	// 			s += "   Press 'ctrl+e' to restake and start mining again\n"
	// 		}

	// 		hubInfo, err := m.taskManager.GetHubGlobalInfo()
	// 		if err == nil {
	// 			// s += fmt.Sprintf("\n %s Hub Global Info:%s%s\n", ">", " ", textStyle(fmt.Sprintf("%v", hubInfo)))
	// 			s += fmt.Sprintf("\n %s Total models: %s | Total Miners: %s\n", ">", textStyle(fmt.Sprintf("%v", hubInfo.TotalModels)), textStyle(fmt.Sprintf("%v", hubInfo.TotalMiners)))
	// 		}
	// 	}
	// }

	// s += bindingShortCutStyle("\nPress 'q' or 'ctrl+c' to quit\n")

	// if hasNewVer, newVer := m.taskManager.HasNewVersion(); hasNewVer {
	// 	s += "\n "
	// 	s += fmt.Sprintf("%s %s %s", colorToStyle("warning").Render(m.newVersionSpinner.View()), colorToStyle("warning").Render("New version available:"), colorToStyle("warning-bg").Bold(true).Render(newVer))
	// 	s += fmt.Sprintf("\n\n")
	// 	s += colorToStyle("warning-bg").Render(fmt.Sprintf("stop and start the program to update to the new version"))
	// 	// s += colorToStyle("warning-bg").Render(fmt.Sprintf(" %s New version available:%s%s", m.newVersionSpinner.View(), " ", textStyle(newVer)))
	// 	s += "\n"
	// }

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

func (rt *Router) GetChains(c *gin.Context) {
	c.JSON(http.StatusOK, APIResponse{
		Data:   config.ChainConfigs,
		Status: http.StatusOK,
	})
}

func (rt *Router) Unstake(c *gin.Context) {
	err := rt.watcher.Unregister()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "unstake requested",
		Status:  http.StatusOK,
	})
}

func (rt *Router) ClaimReward(c *gin.Context) {
	err := rt.watcher.ClaimMiningReward()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "claimed reward",
		Status:  http.StatusOK,
	})
}

func (rt *Router) ClaimUnstake(c *gin.Context) {
	err := rt.watcher.ReclaimStake()
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "claimed reward",
		Status:  http.StatusOK,
	})
}
