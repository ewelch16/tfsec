package s3

import (
	"github.com/aquasecurity/tfsec/pkg/defsec/definition"
)

type S3 struct {
	Buckets []Bucket
}

type Bucket struct {
	*definition.Metadata
	PublicAccessBlock PublicAccessBlock
	BucketPolicy      BucketPolicy
	Encryption        Encryption
	Versioning        Versioning
	Logging           Logging
	ACL               definition.StringValue
}

func (b *Bucket) HasPublicExposureACL() bool {
	for _, publicACL := range []string{"public-read", "public-read-write", "website", "authenticated-read"} {
		if b.ACL.Value == publicACL {
			return true
		}
	}
	return false
}

type PublicAccessBlock struct {
}

type BucketPolicy struct {
}

type Logging struct {
	Enabled definition.BoolValue
}

type Versioning struct {
	Enabled definition.BoolValue
}

type Encryption struct {
	Enabled   definition.BoolValue
	Algorithm definition.StringValue
	KMSKeyId  definition.StringValue
}
