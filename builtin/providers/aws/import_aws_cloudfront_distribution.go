package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsCloudFrontDistributionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	conn := meta.(*AWSClient).cloudfrontconn
	id := d.Id()
	resp, err := conn.GetDistributionConfig(&cloudfront.GetDistributionConfigInput{
		Id: aws.String(id),
	})

	if err != nil {
		return nil, err
	}

	distConfig := resp.DistributionConfig
	results := make([]*schema.ResourceData, 1)
	err = flattenDistributionConfig(d, distConfig)
	if err != nil {
		return nil, err
	}
	results[0] = d
	return results, nil
}
