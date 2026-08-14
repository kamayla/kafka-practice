// ch04-producer: franz-go で Producer を書く（章本文の最終形）
//
// 事前にルートで `docker compose up -d` を実行しておくこと。
// 本書第4章の段階的なコードを 1 本にまとめた実行サンプル。
//
//	go run ./ch04-producer
//
// 送信後は次のコマンドで届いているか確認できる:
//
//	docker compose exec broker /opt/kafka/bin/kafka-console-consumer.sh \
//	    --bootstrap-server localhost:9092 --topic orders \
//	    --from-beginning --timeout-ms 3000 --property print.key=true
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
	)
	if err != nil {
		log.Fatalf("クライアント生成に失敗: %v", err)
	}
	defer cl.Close()

	ctx := context.Background()
	rec := &kgo.Record{Topic: "orders", Key: []byte("uustomer-48"), Value: []byte("order-1 created")}
	res := cl.ProduceSync(ctx, rec)
	r, err := res.First()
	if err != nil {
		log.Fatalf("送信に失敗: %v", err)
	}
	fmt.Printf("送信成功: topic=%s partition=%d offset=%d\n", r.Topic, r.Partition, r.Offset)
}
