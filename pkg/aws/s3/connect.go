package s3

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func (s *s3Client) connect(ctx context.Context) (err error) {

	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           s.config.Endpoint,
					SigningRegion: s.config.Region,
				}, nil
			}
			// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
			//nolint:exhaustruct // aws struct, not our struct
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		},
	)

	awsCfnOptFns := []func(*config.LoadOptions) error{
		config.WithRegion(s.config.Region),
		config.WithEndpointResolverWithOptions(customResolver),
	}
	// Static config (only for local development)
	if s.config.StaticAccessKey != nil && s.config.StaticSecretKey != nil {
		slog.WarnContext(ctx, "don't use static credentials in production")
		awsCfnOptFns = append(
			awsCfnOptFns,
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				*s.config.StaticAccessKey,
				*s.config.StaticSecretKey,
				"local")),
		)
	}
	cfg, err := config.LoadDefaultConfig(
		ctx,
		awsCfnOptFns...,
	)

	if err != nil {
		slog.ErrorContext(ctx, "s3 connect error #1",
			slog.String("S3_REGION", s.config.Region),
			slog.String("S3_ENDPOINT", s.config.Endpoint),
			slog.String("S3_BUCKET", s.config.Bucket),
		)
		return err
	}

	// Assume a role if the role ARN is set
	if s.config.Role != nil {
		// Create the credentials from AssumeRoleProvider to assume the role
		// referenced by the "myRoleARN" ARN.
		stsSvc := sts.NewFromConfig(cfg)
		creds := stscreds.NewAssumeRoleProvider(stsSvc, *s.config.Role)
		cfg.Credentials = aws.NewCredentialsCache(creds)
	}

	s.client = s3.NewFromConfig(cfg)
	s.presignClient = s3.NewPresignClient(s.client)

	return nil
}
