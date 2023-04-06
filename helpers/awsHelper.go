package helper

import (
	"log"

	"github.com/TirthShingala/the-user/constants"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3_BUCKET_NAME = constants.S3BucketName
const S3_REGION = constants.S3Region
const PRE_SIGNED_URL_EXPIRATION = constants.PreSignedUrlExpiration

var sess, _ = session.NewSession(&aws.Config{
	Region: aws.String(S3_REGION)},
)

var svc = s3.New(sess)

func SignedURL(key *string, ContentLength *int64) (preSignURL string) {
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:        aws.String(S3_BUCKET_NAME),
		Key:           key,
		ContentLength: ContentLength,
	})

	preSignURL, err := req.Presign(PRE_SIGNED_URL_EXPIRATION)

	if err != nil {
		log.Println("Failed to sign request", err)
		return ""
	}

	return preSignURL
}
