package log

import (
	"bytes"
	"encoding/json"
	"log"
)

// awsLogger logger impl for awslambda
type awsLogger struct {
}

// NewAwsLogger constructor
func NewAwsLogger() LoggerImpl {
	return &awsLogger{}
}

// LogWrite logwrite
func (u *awsLogger) LogWrite(lv Level, msg string) {
	log.Printf("level:%s,msg:%v", lv, msg)
}

// LogWriteWithObj logwrite
func (u *awsLogger) LogWriteWithObj(lv Level, obj interface{}) {
	var buf bytes.Buffer
	b, _ := json.Marshal(obj)
	buf.Write(b)
	log.Printf("level:%s,obj:%v", lv, buf.String())
}

// LogWriteWithObj logwrite
func (u *awsLogger) LogWriteWithMsgAndObj(lv Level, msg string, obj interface{}) {
	var buf bytes.Buffer
	b, _ := json.Marshal(obj)
	buf.Write(b)
	log.Printf("level:%s,msg:%v,obj:%v", lv, msg, buf.String())
}
