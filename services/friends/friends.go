package friends

import (
	"fmt"
	"go-gin-chat/models"
	"go-gin-chat/services/session"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FriendController struct{}

// GetFriendList 获取好友列表
func GetFriendList(c *gin.Context) {
	// 获取当前用户ID
	userInfo := session.GetSessionUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请先登录",
		})
		return
	}

	user_id := int(userInfo["uid"].(uint))

	// 查询好友列表
	var friends []models.Friends
	if err := models.ChatDB.Where("user_id = ?", user_id).Find(&friends).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取好友列表失败",
		})
		return
	}

	// 获取好友的用户信息
	var friendList []map[string]interface{}
	for _, friend := range friends {
		var user models.User
		if err := models.ChatDB.Where("id = ?", friend.ToUserId).First(&user).Error; err != nil {
			continue
		}
		friendInfo := map[string]interface{}{
			"user_id":   user.ID,
			"username":  user.Username,
			"avatar_id": user.AvatarId,
		}
		friendList = append(friendList, friendInfo)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": friendList,
	})
}

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	// 获取当前用户ID
	userInfo := session.GetSessionUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请先登录",
		})
		return
	}

	// 获取要添加的好友ID
	toUserIdStr := c.PostForm("to_user_id")
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}

	// 检查是否已经是好友
	var existingFriend models.Friends
	result := models.ChatDB.Where("user_id = ? AND to_user_id = ?", userInfo["uid"], toUserId).First(&existingFriend)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经是好友关系",
		})
		return
	}

	user_id := int(userInfo["uid"].(uint))

	fmt.Println(user_id, reflect.TypeOf(user_id))
	// 创建好友关系
	friend := models.Friends{
		UserId:   user_id,
		ToUserId: toUserId,
	}

	if err := models.ChatDB.Create(&friend).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "添加好友失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "添加好友成功",
	})
}

// GetRecommendFriends 获取推荐好友列表
func GetRecommendFriends(c *gin.Context) {
	// 获取当前用户ID
	userInfo := session.GetSessionUserInfo(c)
	if userInfo == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "请先登录",
		})
		return
	}

	user_id := int(userInfo["uid"].(uint))

	// 获取当前用户的好友列表
	var friends []models.Friends
	if err := models.ChatDB.Where("user_id = ?", user_id).Find(&friends).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取好友列表失败",
		})
		return
	}

	// 获取好友的好友列表
	var recommendUsers = make(map[uint]models.User)
	for _, friend := range friends {
		var friendOfFriends []models.Friends
		if err := models.ChatDB.Where("user_id = ?", friend.ToUserId).Find(&friendOfFriends).Error; err != nil {
			continue
		}

		// 获取好友的好友的用户信息
		for _, fof := range friendOfFriends {
			// 排除自己
			if fof.ToUserId == user_id {
				continue
			}

			// 排除已经是好友的用户
			var isExist bool
			for _, f := range friends {
				if f.ToUserId == fof.ToUserId {
					isExist = true
					break
				}
			}
			if isExist {
				continue
			}

			// 获取用户信息
			var user models.User
			if err := models.ChatDB.Where("id = ?", fof.ToUserId).First(&user).Error; err != nil {
				continue
			}
			recommendUsers[user.ID] = user
		}
	}

	// 转换为列表
	var recommendList []map[string]interface{}
	for _, user := range recommendUsers {
		recommendInfo := map[string]interface{}{
			"user_id":   user.ID,
			"username":  user.Username,
			"avatar_id": user.AvatarId,
		}
		recommendList = append(recommendList, recommendInfo)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取成功",
		"data": recommendList,
	})
}
