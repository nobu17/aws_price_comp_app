package info

import "os"

// GetItemObservSQSURL get sqs obs url from enviroment
func GetItemObservSQSURL() string {
	return os.Getenv("ITEM_OBS_SQS_URL")
}
