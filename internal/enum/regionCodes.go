package enums

// AWSRegion is a custom type for AWS regions
type AWSRegion string

// Define AWS regions using constants
const (
	US_EAST_1         AWSRegion = "us-east-1"      // North Virginia
	US_EAST_2         AWSRegion = "us-east-2"      // Ohio
	US_WEST_1         AWSRegion = "us-west-1"      // Northern California
	US_WEST_2         AWSRegion = "us-west-2"      // Oregon
	CA_CENTRAL_1      AWSRegion = "ca-central-1"   // Canada Central
	EU_WEST_1         AWSRegion = "eu-west-1"      // Ireland
	EU_WEST_2         AWSRegion = "eu-west-2"      // London
	EU_WEST_3         AWSRegion = "eu-west-3"      // Paris
	EU_CENTRAL_1      AWSRegion = "eu-central-1"   // Frankfurt
	AP_SOUTHEAST_1    AWSRegion = "ap-southeast-1" // Singapore
	AP_SOUTHEAST_2    AWSRegion = "ap-southeast-2" // Sydney
	AP_NORTHEAST_1    AWSRegion = "ap-northeast-1" // Tokyo
	AP_NORTHEAST_2    AWSRegion = "ap-northeast-2" // Seoul
	AP_SOUTH_1        AWSRegion = "ap-south-1"     // Mumbai
	SA_EAST_1         AWSRegion = "sa-east-1"      // São Paulo
	ME_SOUTH_1        AWSRegion = "me-south-1"     // Bahrain
	AFRICA_UK_SOUTH_1 AWSRegion = "af-south-1"     // Cape Town
	US_GOV_WEST_1     AWSRegion = "us-gov-west-1"  // AWS GovCloud (US-West)
	US_GOV_EAST_1     AWSRegion = "us-gov-east-1"  // AWS GovCloud (US-East)
	CN_NORTH_1        AWSRegion = "cn-north-1"     // Beijing
	CN_NORTHWEST_1    AWSRegion = "cn-northwest-1" // Ningxia
)
