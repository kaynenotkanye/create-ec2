# Golang app to create EC2 from AMI
This is a Golang app that will launch an EC2 instance from a source AMI
while also passing along a userdata script upon EC2 launch.
### Example usage
The following command will create an EC2 based off of ami-abcd1234 in the us-west-2 region placed in  subnet-abcd1234 (notice that we only need subnet-id and not VPC ID) assigned to security groupo sg-abcd1234 while passing along the sampleUserData.txt file with the instance type of t3.medium.
```sh
create-ec2 -r us-west-2 \
-n NameOfEC2instance \
-i ami-abcd1234 \
-u sampleUserData.txt \
-k MyKeyPair \
-v subnet-abcd1234 \
-s sg-abcd1234 \
-t t3.medium
```
The following command is essentially the same, except it passes in the ami ID through a text file (the typical usage for this is through a CI/CD pipeline as the ami.txt can be from a build artifact, or even passed along from another pipeline.)
```sh
create-ec2 -r us-west-2 \
-n NameOfEC2instance \
-i $(cat ami.txt) \
-u sampleUserData.txt \
-k MyKeyPair \
-v subnet-abcd1234 \
-s sg-abcd1234 \
-t t3.medium
```
