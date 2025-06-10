package aws

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"

	"musicapi/lambdaSrc/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	BucketName   = ""
	BucketRegion = ""
)

func UpdateTrackList(trackName string, userName string) error {

	trackListBytes, err := LoadS3Object(userName)
	if err != nil {
		err = CreateNewTrackList(trackName, userName)
		if err != nil {
			return err
		}
		return nil
	}

	var trackList types.TrackList
	err = json.Unmarshal(trackListBytes, &trackList)
	if err != nil {
		return err
	}

	trackList.TrackIDs = append(trackList.TrackIDs, trackName)

	trackListBytes, err = json.Marshal(trackList)
	if err != nil {
		return err
	}

	err = SaveS3Object(userName, trackListBytes)
	if err != nil {
		return err
	}

	return nil
}

func CreateNewTrackList(trackName string, userName string) error {

	newTrackList := types.TrackList{
		TrackIDs: []string{
			trackName,
		},
	}

	newTrackListBytes, err := json.Marshal(newTrackList)
	if err != nil {
		return err
	}

	err = SaveS3Object(userName, newTrackListBytes)
	if err != nil {
		return err
	}

	return nil
}

func GetTrackList(userName string) (types.TrackList, error) {

	trackListBytes, err := LoadS3Object(userName)
	if err != nil {
		return types.TrackList{}, err
	}

	var trackList types.TrackList
	err = json.Unmarshal(trackListBytes, &trackList)
	if err != nil {
		return types.TrackList{}, err
	}

	return trackList, nil
}

func GetRandomTrack(userName string) (string, error) {

	trackListBytes, err := LoadS3Object(userName)
	if err != nil {
		return "", err
	}

	var trackList types.TrackList
	err = json.Unmarshal(trackListBytes, &trackList)
	if err != nil {
		return "", err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(trackList.TrackIDs) - 1)

	return trackList.TrackIDs[randomIndex], nil
}

func LoadS3Object(objectName string) ([]byte, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(BucketRegion),
	})
	if err != nil {
		return []byte{}, err
	}

	s3Client := s3.New(sess)

	s3Response, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(objectName),
	})
	if err != nil {
		return []byte{}, err
	}

	objectBytes, err := io.ReadAll(s3Response.Body)
	if err != nil {
		return []byte{}, err
	}

	return objectBytes, nil
}

func SaveS3Object(objectName string, object []byte) error {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(BucketRegion),
	})
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.New(sess)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(BucketName),
		Key:         aws.String(objectName),
		Body:        bytes.NewReader(object),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return err
	}

	return nil
}
