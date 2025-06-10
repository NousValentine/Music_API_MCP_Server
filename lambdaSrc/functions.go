package main

import (
	"musicapi/lambdaSrc/aws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTrack(c *gin.Context) {
	trackID := c.Query("trackID")
	userName := c.Query("userName")

	if trackID == "" || userName == "" {
		c.JSON(http.StatusBadRequest, nil)
	}

	err := aws.UpdateTrackList(trackID, userName)
	if err != nil {
		c.JSON(http.StatusFailedDependency, err)
	}

	c.JSON(http.StatusOK, nil)
}

func GetAllTracks(c *gin.Context) {

	userName := c.Query("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, nil)
	}

	trackIDList, err := aws.GetTrackList(userName)
	if err != nil {
		c.JSON(http.StatusFailedDependency, err)
	}

	c.JSON(http.StatusOK, trackIDList)
}

func GetRandomTrack(c *gin.Context) {

	userName := c.Query("userName")
	if userName == "" {
		c.JSON(http.StatusBadRequest, nil)
	}

	randomTrackID, err := aws.GetRandomTrack(userName)
	if err != nil {
		c.JSON(http.StatusFailedDependency, err)
	}
	c.JSON(http.StatusOK, randomTrackID)
}
