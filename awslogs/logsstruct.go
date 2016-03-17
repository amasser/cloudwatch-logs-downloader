package awslogs

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