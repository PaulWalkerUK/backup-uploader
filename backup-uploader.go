package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Backup Uploader")
		fmt.Println("===============")
		fmt.Println()
		fmt.Printf("Usage: %s <file_path>\n", filepath.Base(os.Args[0]))
		fmt.Println()
		fmt.Println("Uploads the specified file to a bucket on Amazon S3.")
		fmt.Println("Please see: https://github.com/PaulWalkerUK/backup-uploader")
		fmt.Println("Version 1.0.0")
		return
	}

	filePath := os.Args[1]
	//filePath := "test.txt"

	// Read configuration from config.json
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Dir(os.Args[0]))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	log.Print("Read config")

	sendFile(viper.GetString("region"), viper.GetString("bucket"), viper.GetString("access_key"), viper.GetString("secret_key"), filePath)

	fmt.Println("File uploaded successfully!")
}

func sendFile(region string, bucket string, accessKey string, secretKey string, fullFilePath string) {
	PrintMemUsage()

	file, err := os.Open(fullFilePath)

	if err != nil {
		log.Fatalf("Error opening file to upload: %s", err)
	}
	defer file.Close()
	log.Print("Input file opened")

	PrintMemUsage()

	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}

	log.Print("Created S3 config")

	s3Session, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("Error creating S3 session: %s", err)
	}

	log.Print("Created S3 session")

	uploader := s3manager.NewUploader(s3Session)

	log.Print("Created S3 upload manager")
	PrintMemUsage()

	filename := filepath.Base(fullFilePath)

	input := &s3manager.UploadInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(filename),
		Body:         file,
		StorageClass: aws.String("DEEP_ARCHIVE"),
	}

	log.Print("Input prepared")
	PrintMemUsage()

	if true {
		log.Printf("Uploading %s...", filename)
		output, err := uploader.Upload(input)
		if err != nil {
			log.Fatalf("Error uploading: %s", err)
		}

		if false {
			log.Println(output)
		}
		log.Print("Upload complete")
	}

	PrintMemUsage()

}

// From https://gophercoding.com/print-current-memory-usage/
func PrintMemUsage() {
	//var m runtime.MemStats
	//runtime.ReadMemStats(&m)
	//log.Printf("    > Alloc = %v MiB    TotalAlloc = %v MiB    Sys = %v MiB    NumGC = %v", m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
}
