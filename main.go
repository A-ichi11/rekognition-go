package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

const imagePath = "image/"

var imageName1 = "gun.png"
var imageName2 = "human.jpg"
var imageName3 = "sakura.jpeg"
var imageName4 = "水着.png"

func main() {
	// imageNameを変更して実行することで、画像を変更します。
	// ここ雑です。すいません。
	file := imagePath + imageName4

	// 画像ファイルを取得
	imageFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// 最後に画像ファイルを閉じます
	defer imageFile.Close()

	// 画像ファイルのデータを読み込み
	bytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// セッション作成
	sess := session.Must(session.NewSession())

	// Rekognitionクライアントを作成
	svc := rekognition.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	// パラメータを設定
	params := &rekognition.DetectModerationLabelsInput{
		Image: &rekognition.Image{
			Bytes: bytes,
		},
	}

	// DetectModerationLabelsを実行。resultに結果のLabelが入ります。
	result, err := svc.DetectModerationLabels(params)
	// 結果を出力
	fmt.Println(result)
}
