package slaves

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"llm_training_management_system/pkg/ltms_config"
)

var httpConfig *ltms_config.HttpConfig

func BindRouter(config *ltms_config.HttpConfig, r *gin.Engine) {
	httpConfig = config
	r.GET("/slave/list", ls)
}

func ls(c *gin.Context) {
	fmt.Println("start ls")
	allSlaves := ListAllSlaves()
	c.Set("data", struct {
		Size   int         `json:"size"`
		Slaves []SlaveNode `json:"slaves"`
	}{Size: len(allSlaves), Slaves: allSlaves})
}
