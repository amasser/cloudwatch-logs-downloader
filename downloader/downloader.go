package main
import "fmt"
import "strings"
import "strconv"
import "github.com/deep-patel/cloudwatch_log_downloader/awslogs"
import "github.com/deep-patel/cloudwatch_log_downloader/util"
func main() {
        //cmd := exec.Command("echo", "Called from Go!")
        logStreamData:= awslogs.GetLogStreams("/aws/lambda/mParticleEventListener", "2016/03/16/[17]")
        logStreamList := logStreamData.LogStreams;
        logStreamListLength := len(logStreamList)
        
        

        for i:=0; i<logStreamListLength; i++ {
                fmt.Println("Fetching Events for stream: " + logStreamList[i].LogStreamName)
                logEventsData := awslogs.GetLogEvents("/aws/lambda/mParticleEventListener",logStreamList[i].LogStreamName)
                logEventsList:= logEventsData.Events
                logEventsListLength:= len(logEventsList)
                for logEventsListLength!=0 {
                    toBeWritten := make([][]string, logEventsListLength)
                    for i := range toBeWritten {
                        toBeWritten[i] = make([]string, 3)
                    }
                    for j:=0; j<logEventsListLength; j++{
                            toBeWritten[j][0] = strconv.Itoa(logEventsList[j].Timestamp);
                            toBeWritten[j][1] = strconv.Itoa(logEventsList[j].IngestionTime);
                            toBeWritten[j][2] = logEventsList[j].Message;
                    }
                    util.GenerateCSV(strings.Replace(logStreamList[i].LogStreamName + ".csv", "/", "_", -1), toBeWritten)
                    logEventsData = awslogs.GetLogEventsWithToken("/aws/lambda/mParticleEventListener",logStreamList[i].LogStreamName, logEventsData.NextForwardToken)
                    logEventsList = logEventsData.Events
                    logEventsListLength = len(logEventsList)
                }
        }
}


