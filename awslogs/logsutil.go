package awslogs

import "os/exec"
import "encoding/json"
import "github.com/deep-patel/cloudwatch_log_downloader/util"


func GetLogStreams(logGroupName string, logStreamPrefix string) (LogStreamsData){
        //aws --profile pso-lambda-user logs describe-log-streams --log-group-name /aws/lambda/mParticleEventListener --log-stream-name-prefix 2016/03/16/[17]
        
        cmd := exec.Command("aws", "--profile", "pso-lambda-user", "logs", "describe-log-streams", 
                                "--log-group-name", logGroupName, 
                                "--log-stream-name-prefix", logStreamPrefix);
        util.PrintCommand(cmd)
        output, err := cmd.CombinedOutput()
        util.PrintError(err)

        //printOutput(output)

        logStreamsData:= LogStreamsData{}
        json.Unmarshal(output, &logStreamsData)
        return logStreamsData;
}

func GetLogEvents(logGroupName string, logStreamName string) (LogEventsData){
        //aws --profile pso-lambda-user logs get-log-events --log-group-name /aws/lambda/mParticleEventListener --log-stream-name 2016/03/16/[17]45faa53aac5c44e4aec07c182189642c
        cmd := exec.Command("aws", "--profile", "pso-lambda-user", "logs", "get-log-events", 
                                "--log-group-name", logGroupName, 
                                "--log-stream-name", logStreamName);
        util.PrintCommand(cmd)
        output, err := cmd.CombinedOutput()
        util.PrintError(err)

        logEventsData:= LogEventsData{}
        json.Unmarshal(output, &logEventsData)
        return logEventsData;
}


func GetLogEventsWithToken(logGroupName string, logStreamName string, tokenId string) (LogEventsData){
        //aws --profile pso-lambda-user logs get-log-events --log-group-name /aws/lambda/mParticleEventListener --log-stream-name 2016/03/16/[17]45faa53aac5c44e4aec07c182189642c
        cmd := exec.Command("aws", "--profile", "pso-lambda-user", "logs", "get-log-events", 
                                "--log-group-name", logGroupName, 
                                "--log-stream-name", logStreamName,
                                "--next-token", tokenId);
        util.PrintCommand(cmd)
        output, err := cmd.CombinedOutput()
        util.PrintError(err)

        logEventsData:= LogEventsData{}
        json.Unmarshal(output, &logEventsData)
        return logEventsData;
}