// Attini resources
package attini_cdk_lib


type AwsVpcConfiguration struct {
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	SecurityGroups *map[string]interface{} `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	Subnets *map[string]interface{} `field:"optional" json:"subnets" yaml:"subnets"`
}

