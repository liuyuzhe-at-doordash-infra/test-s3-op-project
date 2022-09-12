package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	//pulumi.Run(func(ctx *pulumi.Context) error {
	//	// Create an AWS resource (S3 Bucket)
	//	bucket, err := s3.NewBucket(ctx, "my-tmp-test-pulumi-bucket1", nil)
	//	if err != nil {
	//		return err
	//	}
	//
	//	// Export the name of the bucket
	//	ctx.Export("bucketName", bucket.ID())
	//	return nil
	//})

	/* generated from cmd:
	pulumi import aws:s3/bucket:Bucket my-test-pulumi-s3-result-from-import-tf-s3 my-test-s3bucket-20220912203139557000000001
	*/
	pulumi.Run(func(ctx *pulumi.Context) error {
		mybucket, err := s3.NewBucket(ctx, "my-test-pulumi-s3-result-from-import-tf-s3", &s3.BucketArgs{
			Arn:          pulumi.String("arn:aws:s3:::my-test-s3bucket-20220912203139557000000001"),
			Bucket:       pulumi.String("my-test-s3bucket-20220912203139557000000001"),
			HostedZoneId: pulumi.String("Z2O1EMRO9K5GLX"),
			RequestPayer: pulumi.String("BucketOwner"),
			Tags: pulumi.StringMap{
				"environment": pulumi.String("DEV"),
				"terraform":   pulumi.String("true"),
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		}, pulumi.Protect(true))

		_, err = s3.NewBucketObject(ctx, "test-object2.txt", &s3.BucketObjectArgs{
			Bucket: mybucket.ID(),
			Source: pulumi.NewFileAsset("test-object2.txt"),
		})

		if err != nil {
			return err
		}
		return nil
	})

}
