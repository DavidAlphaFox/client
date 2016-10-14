package chat

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/goamz/goamz/aws"
	"github.com/keybase/client/go/chat/s3"
	"github.com/keybase/client/go/logger"
	"github.com/keybase/client/go/protocol/chat1"
)

type UploadS3Result struct {
	Region   string
	Endpoint string
	Bucket   string
	Path     string
	Size     int64
}

func UploadS3(log logger.Logger, r io.Reader, filename string, params chat1.S3AttachmentParams) (*UploadS3Result, error) {
	var body bytes.Buffer
	mpart := multipart.NewWriter(&body)

	// the order of these is important:
	if err := mpart.WriteField("key", params.ObjectKey); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("acl", params.Acl); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("X-Amz-Credential", params.Credential); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("X-Amz-Algorithm", params.Algorithm); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("X-Amz-Date", params.Date); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("Policy", params.Policy); err != nil {
		return nil, err
	}
	if err := mpart.WriteField("X-Amz-Signature", params.Signature); err != nil {
		return nil, err
	}
	part, err := mpart.CreateFormFile("file", filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	n, err := io.Copy(part, r)
	if err != nil {
		return nil, err
	}
	log.Debug("copied %d bytes to multipart", n)

	if err := mpart.Close(); err != nil {
		return nil, err
	}

	// XXX retry
	resp, err := http.Post(params.Endpoint, mpart.FormDataContentType(), &body)
	if err != nil {
		return nil, err
	}
	log.Debug("s3 post response: %+v", resp)
	bstr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	log.Debug("s3 post response body:  %s", bstr)

	// XXX check response
	res := UploadS3Result{
		Bucket: params.Bucket,
		Path:   params.ObjectKey,
		Size:   int64(n),
	}

	return &res, nil
}

func PutS3(log logger.Logger, r io.Reader, size int64, params chat1.S3Params, signer s3.Signer) (*UploadS3Result, error) {
	region := aws.Region{
		Name:             params.RegionName,
		S3Endpoint:       params.RegionEndpoint,
		S3BucketEndpoint: params.RegionBucketEndpoint,
	}
	conn := s3.New(signer, region)
	conn.AccessKey = params.AccessKey

	b := conn.Bucket(params.Bucket)
	err := b.PutReader(params.ObjectKey, r, size, "application/octet-stream", s3.ACL(params.Acl), s3.Options{})
	if err != nil {
		return nil, err
	}

	res := UploadS3Result{
		Region:   params.RegionName,
		Endpoint: params.RegionEndpoint,
		Bucket:   params.Bucket,
		Path:     params.ObjectKey,
		Size:     size,
	}

	return &res, nil
}

func DownloadAsset(ctx context.Context, log logger.Logger, params chat1.S3Params, asset chat1.Asset, w io.Writer, signer s3.Signer) error {
	region := aws.Region{
		Name:       asset.Region,
		S3Endpoint: asset.Endpoint,
	}
	conn := s3.New(signer, region)
	conn.AccessKey = params.AccessKey

	b := conn.Bucket(asset.Bucket)

	body, err := b.GetReader(asset.Path)
	defer func() {
		if body != nil {
			body.Close()
		}
	}()
	if err != nil {
		return err
	}

	n, err := io.Copy(w, body)
	if err != nil {
		return err
	}

	log.Debug("downloaded %d bytes", n)

	return nil
}
