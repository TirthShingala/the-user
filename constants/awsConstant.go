package constants

import "time"

const (
	S3BucketName = "videosdk-live-storage-dev"
	S3Region     = "ap-south-1"

	PreSignedUrlExpiration = 5 * time.Minute

	CDN_BASE_URL = "https://cdn.example.com"
)
