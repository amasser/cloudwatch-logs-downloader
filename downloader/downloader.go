package main
import "fmt"
import "os"
import "strings"
import "os/exec"
import "encoding/json"
import "encoding/csv"
import "strconv"
import "cloudwatch_log_downloader/util"


type LogStreamsData struct {
        LogStreams []struct {
                Arn                 string `json:"arn"`
                CreationTime        int    `json:"creationTime"`
                FirstEventTimestamp int    `json:"firstEventTimestamp"`
                LastEventTimestamp  int    `json:"lastEventTimestamp"`
                LastIngestionTime   int    `json:"lastIngestionTime"`
                LogStreamName       string `json:"logStreamName"`
                StoredBytes         int    `json:"storedBytes"`
                UploadSequenceToken string `json:"uploadSequenceToken"`
        } `json:"logStreams"`
}

type LogEventsData struct {
        Events []struct {
                IngestionTime int    `json:"ingestionTime"`
                Message       string `json:"message"`
                Timestamp     int    `json:"timestamp"`
        } `json:"events"`
        NextBackwardToken string `json:"nextBackwardToken"`
        NextForwardToken  string `json:"nextForwardToken"`
}

/**func printCommand(cmd *exec.Cmd){
        fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
        if err != nil{
                os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
        }
}

func printOutput(outs []byte){
        if len(outs) > 0{
                fmt.Printf("==> Output: %s\n", string(outs))
        }      
}*/

func getLogStreams(logGroupName string, logStreamPrefix string) (LogStreamsData){
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

func getLogEvents(logGroupName string, logStreamName string) (LogEventsData){
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

func getLogEventsWithToken(logGroupName string, logStreamName string, tokenId string) (LogEventsData){
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

func main() {
        //cmd := exec.Command("echo", "Called from Go!")
        logStreamData:= getLogStreams("/aws/lambda/mParticleEventListener", "2016/03/16/[17]")
        logStreamList := logStreamData.LogStreams;
        logStreamListLength := len(logStreamList)
        
        

        for i:=0; i<logStreamListLength; i++ {
                fmt.Println("Fetching Events for stream: " + logStreamList[i].LogStreamName)
                logEventsData := getLogEvents("/aws/lambda/mParticleEventListener",logStreamList[i].LogStreamName)
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
                    generateCSV(strings.Replace(logStreamList[i].LogStreamName + ".csv", "/", "_", -1), toBeWritten)
                    logEventsData = getLogEventsWithToken("/aws/lambda/mParticleEventListener",logStreamList[i].LogStreamName, logEventsData.NextForwardToken)
                    logEventsList = logEventsData.Events
                    logEventsListLength = len(logEventsList)
                }
        }
}


func generateCSV(fileName string, records [][]string){

        fmt.Println(records)
        csvfile, err := os.OpenFile("logs/" + fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE,0600)
        if err != nil{
                fmt.Println("Error: ", err)
                return;
        }
        defer csvfile.Close()

        writer := csv.NewWriter(csvfile)
        for _, record := range records {
                  err := writer.Write(record)
                  if err != nil {
                          fmt.Println("Error:", err)
                          return
                  }
        }
          writer.Flush()
}
