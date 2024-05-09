package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Provider interface {
	Download(
		ctx context.Context,
		inputURLs []string,
	) (
		mapURLFilepath map[string]string, err error)

	Zip(ctx context.Context, folderName string, fileNames map[string]string) (err error)
	UploadFolderZip(
		ctx context.Context,
		folderName string,
		localPath string,
		s3Path string,
	) (s3URL string, err error)
	/*
		# GetS3ObjContentT1

		get content of S3 object and transfer into buffer
	*/
	GetS3ObjContentT1(
		ctx context.Context,
		s3Path string,
		fileName string,
	) (contentBuf io.Reader, err error)
	/*
		# GetS3ObjContent2

		get content of S3 object and transfer into buffer
	*/
	GetS3ObjContentT2(
		ctx context.Context,
		s3Key string,
	) (contentBuf io.Reader, err error)
	// UploadContent .. upload content to s3 with s3Path and fileName.
	// Auto generate s3Path and fileName if they are nil.
	UploadContent(
		ctx context.Context,
		buffer io.Reader,
		s3Path *string,
		fileName *string,
	) (s3Region string, s3Bucket string, s3Key string, err error)
	UploadJSON(
		ctx context.Context,
		jsonableStruct any,
	) (s3Location string, err error)
	GetFileName(ctx context.Context, s3URL string) (fileName string, err error)
	GetBucketPath(ctx context.Context, s3URL string) (bucketPath string, err error)
	DeleteS3ObjT1(
		ctx context.Context,
		s3Path string,
		fileName string,
	) (err error)
	DeleteS3ObjT2(
		ctx context.Context,
		s3Key string,
	) (err error)

	/*
		PresignGetObject

		using url to put object
			curl -X GET "presigned URL"
	*/
	PresignGetObject(
		ctx context.Context,
		s3ObjectKey string,
		lifetimeSecs int64,
	) (getObjectURL string, err error)
	/*
		PresignPutObject

		using url to put object
			curl -X PUT -T "/path/to/file" "presigned URL"
	*/
	PresignPutObject(
		ctx context.Context,
		s3ObjectKey string,
		lifetimeSecs int64,
	) (putObjectURL string, err error)

	/*
		PresignPutObject

		using url to put object
			curl -X DELETE "presigned URL"
	*/
	PresignDeleteObject(
		ctx context.Context,
		s3ObjectKey string,
		lifetimeSecs int64,
	) (deleteObjectURL string, err error)
}

type S3Config struct {
	Endpoint        string
	Region          string
	Bucket          string
	Role            *string
	StaticAccessKey *string
	StaticSecretKey *string
}

type s3Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	config        *S3Config
}

func New(
	config *S3Config,
) S3Provider {
	s3provdier := &s3Client{
		config: config,
	}

	err := s3provdier.connect(context.TODO())
	if err != nil {
		panic(err)
	}

	return s3provdier
}
