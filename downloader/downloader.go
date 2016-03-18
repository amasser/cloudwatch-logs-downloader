package main
import "fmt"
import "strings"
import "strconv"
import "time"
import "flag"
import "github.com/deep-patel/cloudwatch_log_downloader/awslogs"
import "github.com/deep-patel/cloudwatch_log_downloader/util"

func main() {

        datePtr:= flag.String("startDate", "2016/03/16", "Fetch logs from this date")
        functionVersion:= flag.String("version", "17", "function version")
        manual:= flag.Bool("custom", false, "Custom date")
        log:= flag.String("log", "/mnt/data/mparticle/logs/", "Log location")
        flag.Parse()

        var streamPattern string;
        if (*manual){
            streamPattern = *datePtr + "/[" + *functionVersion + "]";
        } else {
            streamPattern = time.Now().Add(-86400*time.Second).UTC().Format("2006/01/02") + "/[" + *functionVersion + "]";
        }

        logStreamData:= awslogs.GetLogStreams("/aws/lambda/mParticleEventListener", streamPattern)
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
                    for j:=0; j<logEventsListLength; j++ {
                            temp, err := strconv.ParseInt(strconv.Itoa((logEventsList[j].Timestamp/1000)), 10, 64)
                            if err != nil {
                                panic(err)
                            }
                            tm := time.Unix(temp, 0)

                            temp2, err := strconv.ParseInt(strconv.Itoa((logEventsList[j].IngestionTime/1000)), 10, 64)
                            if err != nil {
                                panic(err)
                            }
                            tm2 := time.Unix(temp2, 0)

                            toBeWritten[j][0] = tm.String();
                            toBeWritten[j][1] = tm2.String();
                            toBeWritten[j][2] = logEventsList[j].Message;
                    }
                    util.GenerateCSV(*log + strings.Replace(logStreamList[i].LogStreamName + ".csv", "/", "_", -1), toBeWritten)
                    logEventsData = awslogs.GetLogEventsWithToken("/aws/lambda/mParticleEventListener",logStreamList[i].LogStreamName, logEventsData.NextForwardToken)
                    logEventsList = logEventsData.Events
                    logEventsListLength = len(logEventsList)
                }
        }

        
}