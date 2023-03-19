package conio

import (
	"bufio"
	"context"
	"errors"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	mHost   string
	mUser   string
	mSecret string
	mBucket string
)

// Connect and get object stream
func Stream(objPath string) (*bufio.Scanner, error) {
	const fprobe = "con.Stream"
	var (
		client *minio.Client
		object *minio.Object
		err    error
	)

	if client, err = minio.New(mHost, &minio.Options{
		Creds:  credentials.NewStaticV4(mUser, mSecret, ""),
		Secure: false,
	}); err != nil {
		Probeln(fprobe, err.Error())
		return nil, err
	}
	if object, err = client.GetObject(
		context.Background(),
		mBucket,
		objPath,
		minio.GetObjectOptions{}); err != nil {
		Probeln(fprobe, err.Error())
		return nil, err
	}
	return bufio.NewScanner(object), nil
}

// Setup and test connection
func Setup(host, user, secret, bucket string) error {
	const fprobe = "con.Setup"

	if client, err := minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(user, secret, ""),
		Secure: false,
	}); err != nil {
		Probeln(fprobe, err.Error())
		return err
	} else if ok, err := client.BucketExists(context.Background(), bucket); err != nil {
		Probeln(fprobe, err.Error())
		return err
	} else if !ok {
		err = errors.New("bucket not found error")
		Probeln(fprobe, err.Error())
		return err
	}

	mHost = host
	mUser = user
	mSecret = secret
	mBucket = bucket
	Probeln(fprobe, fmt.Sprintf("Successfully connected to (%s)", mHost))
	return nil
}
