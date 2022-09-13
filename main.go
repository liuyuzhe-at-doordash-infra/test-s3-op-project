package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	/* generated from cmd:
	pulumi import aws:s3/bucket:Bucket my-test-pulumi-s3-result-from-import-tf-2nd-s3 my-test-s3bucket-my-test-s3bucket-20220913002919871500000001	*/
	pulumi.Run(func(ctx *pulumi.Context) error {
		my2ndBucket, err := s3.NewBucket(ctx, "my-test-pulumi-s3-result-from-import-tf-2nd-s3", &s3.BucketArgs{
			Arn:          pulumi.String("arn:aws:s3:::my-test-s3bucket-20220913002919871500000001"),
			Bucket:       pulumi.String("my-test-s3bucket-20220913002919871500000001"),
			HostedZoneId: pulumi.String("Z2O1EMRO9K5GLX"),
			RequestPayer: pulumi.String("BucketOwner"),
			Tags: pulumi.StringMap{
				"environment": pulumi.String("DEV"),
				"terraform":   pulumi.String("true"),
			},
			Versioning: &s3.BucketVersioningArgs{
				Enabled: pulumi.Bool(true),
			},
		}, pulumi.Import(pulumi.ID("my-test-s3bucket-20220913002919871500000001")),
		)
		//pulumi.Protect(true))
		if err != nil {
			return err
		}

		_, objerr := s3.NewBucketObject(ctx, "test-object2.txt", &s3.BucketObjectArgs{
			Bucket: my2ndBucket.ID(),
			Source: pulumi.NewFileAsset("test-object2.txt"),
		})
		if objerr != nil {
			return objerr
		}

		return nil
	})

}
