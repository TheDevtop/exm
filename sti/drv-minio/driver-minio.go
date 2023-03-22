package drvminio

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/TheDevtop/go-probes"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const DriverName = "minio"

var (
	minioHost   = os.Getenv("S3HOST")
	minioUser   = os.Getenv("S3USER")
	minioSecret = os.Getenv("S3SECRET")
	minioBucket = os.Getenv("S3BUCKET")
)

func Stream(path string) (*bufio.Scanner, error) {
	var (
		pb     = probes.NewLogProbe("drvminio.Stream", os.Stderr)
		client *minio.Client
		object *minio.Object
		err    error
	)

	if client, err = minio.New(minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioSecret, ""),
		Secure: false,
	}); err != nil {
		pb.Probe(err.Error())
		return nil, err
	}
	if object, err = client.GetObject(
		context.Background(),
		minioBucket,
		path,
		minio.GetObjectOptions{}); err != nil {
		pb.Probe(err.Error())
		return nil, err
	}
	return bufio.NewScanner(object), nil
}

func Setup() error {
	pb := probes.NewLogProbe("drvminio.Setup", os.Stderr)

	if client, err := minio.New(minioHost, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioSecret, ""),
		Secure: false,
	}); err != nil {
		pb.Probe(err.Error())
		return err
	} else if ok, err := client.BucketExists(context.Background(), minioBucket); err != nil {
		pb.Probe(err.Error())
		return err
	} else if !ok {
		err = errors.New("bucket not found error")
		pb.Probe(err.Error())
		return err
	}

	pb.Probe(fmt.Sprintf("Successfully connected (%s)", minioHost))
	return nil
}
