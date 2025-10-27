package handler

import (
	"order-mock/service"

	"order-mock/model"

	"github.com/gin-gonic/gin"

	handler "order-mock/handler/model"
)

// 获取配置
func GetMockConfigHandler(c *gin.Context) {
	config, err := service.LoadOrderConfig()
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取配置成功",
		"data":    config,
	})
}

// 添加门店配置
func AddDealerConfigHandler(c *gin.Context) {

	var postBody model.OrderRequest

	if err := c.ShouldBindJSON(&postBody); err != nil {
		c.JSON(404, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	err := service.AddDealerConfig(postBody)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "添加门店成功",
	})
}

// 删除门店配置
func DeleteDealerConfigHandler(c *gin.Context) {
	var postBody handler.DeleteDealerConfigRequestParams
	if err := c.ShouldBindJSON(&postBody); err != nil {
		c.JSON(404, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	err := service.DeleteOrderConfigItem(postBody.DealerId)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除门店配置成功",
	})
}
