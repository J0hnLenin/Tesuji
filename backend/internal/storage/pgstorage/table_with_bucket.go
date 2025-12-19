package pgstorage

import "fmt"

func tableWithBacket(bucket bucketNum) string {
	return fmt.Sprintf("%s%d.%s", bucketPrefix, bucket, tableName)
}