package minio

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/ChangSZ/golib/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/ChangSZ/mall-go/configs"
)

var (
	ENDPOINT    = configs.Get().Minio.Endpoint
	BUCKET_NAME = configs.Get().Minio.BucketName
	ACCESS_KEY  = configs.Get().Minio.AccessKey
	SECRET_KEY  = configs.Get().Minio.SecretKey
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) Upload(ctx context.Context, file multipart.File, filename string) (string, string, error) {
	// 创建Minio客户端对象
	minioClient, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY, SECRET_KEY, ""),
		Secure: false,
	})
	if err != nil {
		return "", "", fmt.Errorf("创建客户端失败: %w", err)
	}
	found, err := minioClient.BucketExists(ctx, BUCKET_NAME)
	if err != nil {
		return "", "", err
	}
	if !found {
		// 创建存储桶并设置只读权限
		if err := minioClient.MakeBucket(ctx, BUCKET_NAME, minio.MakeBucketOptions{}); err != nil {
			return "", "", fmt.Errorf("新建桶失败: %w", err)
		}
		// 设置存储桶策略
		policy, err := s.createBucketPolicy(BUCKET_NAME)
		if err != nil {
			return "", "", fmt.Errorf("创建桶策略失败: %w", err)
		}
		if err := minioClient.SetBucketPolicy(ctx, BUCKET_NAME, policy); err != nil {
			return "", "", fmt.Errorf("设置桶策略失败: %w", err)
		}
	} else {
		log.WithTrace(ctx).Info("存储桶已经存在！")
	}

	// 生成存储对象的名称
	if filename == "" {
		filename = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	objectName := fmt.Sprintf("%s/%s", time.Now().Format("20060102"), filename)
	// 上传文件到Minio
	_, err = minioClient.PutObject(ctx, BUCKET_NAME, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", "", fmt.Errorf("PUT文件失败: %w", err)
	}
	log.WithTrace(ctx).Infof("文件上传成功: %v", objectName)
	// 构建返回结果
	url := fmt.Sprintf("%s%s/minio/presigned-url?bucket=%s&objectName=%s",
		configs.ProjectDomain, configs.MallAdminPort, BUCKET_NAME, objectName)
	return url, filename, nil
}

/**
 * 创建存储桶的访问策略，设置为只读权限
 */
func (s *service) createBucketPolicy(bucketName string) (string, error) {
	policy := map[string]interface{}{
		"version": "2024-05-17",
		"statement": []map[string]interface{}{
			{
				"effect":    "Allow",
				"principal": "*",
				"action":    "s3:GetObject",
				"resource":  fmt.Sprintf("arn:aws:s3:::%s/*.**", bucketName),
			},
		},
	}
	policyJSON, err := json.Marshal(policy)
	return string(policyJSON), err
}

func (s *service) Delete(ctx context.Context, objectName string) error {
	// 创建Minio客户端对象
	minioClient, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY, SECRET_KEY, ""),
		Secure: false,
	})
	if err != nil {
		return fmt.Errorf("创建客户端失败: %w", err)
	}

	// 删除文件
	return minioClient.RemoveObject(ctx, BUCKET_NAME, objectName, minio.RemoveObjectOptions{})
}

func (s *service) PresignedURL(ctx context.Context, bucketName, objectName string) (string, error) {
	// 创建Minio客户端对象
	minioClient, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY, SECRET_KEY, ""),
		Secure: false,
	})
	if err != nil {
		return "", fmt.Errorf("创建客户端失败: %w", err)
	}

	// 生成预签名URL
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
