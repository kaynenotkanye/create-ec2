package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"flag"
	// Uncomment "strings" only if you wish to concatentate the EC2 name tag
	//"strings"
	"io/ioutil"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

logFile, err := os.OpenFile("create-ec2.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
if err != nil {
	panic(err)
}
mw := io.MultiWriter(os.Stdout, logFile)
log.SetOutput(mw)

instanceName := flag.String("n", "createdByGO", "Tagged Name of EC2 instance")
userdataFileName := flag.String("u", "sampleUserData.txt", "Name of UserData file.")
awsRegion := flag.String("r", "us-west-2", "AWS region to deploy to, defaults to us-west-2")
awsImageId := flag.String("i", "", "AWS source AMI.")
awsInstanceType := flag.String("t", "t2.micro", "AWS instance type, defaults to t2.micro")
awsKeyName := flag.String("k", "", "AWS keypair name needed for SSH")
awsSubnetId := flag.String("v", "", "Enter VPC Subnet-ID")
awsSecurityGroup := flag.String("s", "", "Enter Security Group ID")
flag.Parse()

// Uncomment the following line if you wish to perform any concatentation for EC2 name
//ec2FullName := []string{*instanceName, *awsImageId}

b, err := ioutil.ReadFile(*userdataFileName)
if err != nil {
    fmt.Print(err)
}
userdata := base64.StdEncoding.EncodeToString(b)

svc := ec2.New(session.New(&aws.Config{Region: awsRegion}))
    // Specify the details of the instance that you want to create.
    runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
        ImageId:      awsImageId,
        InstanceType: awsInstanceType,
        KeyName:      awsKeyName,
        UserData:     aws.String(userdata),
        SubnetId:     awsSubnetId,
        SecurityGroupIds: []*string{ awsSecurityGroup },
        MinCount:     aws.Int64(1),
        MaxCount:     aws.Int64(1),
    })

    if err != nil {
        log.Println("ERROR: Could not create instance", err)
        return
    }

    log.Println("Created instance", *runResult.Instances[0].InstanceId)

    // Add tags to the created instance
    _ , errtag := svc.CreateTags(&ec2.CreateTagsInput{
        Resources: []*string{runResult.Instances[0].InstanceId},
        Tags: []*ec2.Tag{
            {
                Key:   aws.String("Name"),
                Value: instanceName,
                // Uncomment the line below only if you want to concatentate on ec2FullName variable
                //Value: aws.String(strings.Join(ec2FullName, "-")),
            },
        },
    })
    if errtag != nil {
        log.Println("ERROR: Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
        return
    }

    log.Println("Successfully tagged instance")
}
