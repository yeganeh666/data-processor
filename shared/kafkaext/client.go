package kafkaext

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Writer *kafka.Writer
}

func NewClient(brokerAddress string) (*Client, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{brokerAddress},
		CompressionCodec: kafka.Snappy.Codec(),
		BatchSize:        1000,
		BatchBytes:       1048576,
		BatchTimeout:     1000,
	})

	return &Client{
		Writer: writer,
	}, nil
}

func (kc *Client) SendToKafka(topic string, message []byte) error {
	err := kc.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Value: message,
		},
	)

	return err
}

func (kc *Client) Close() {
	err := kc.Writer.Close()
	if err != nil {
		log.WithError(err).Error("Kafka: Failed to close writer")
	}
}

func example() {
	// تنظیمات Kafka
	brokerAddress := "localhost:9092" // آدرس سرور Kafka
	topic := "object-storage-topic"   // نام تاپیک Kafka

	// تنظیمات سرویس object storage
	//objectStorageEndpoint := "object-storage.example.com" // آدرس سرویس object storage

	//partition := 9
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}

	// ساخت نویسنده Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{brokerAddress},
		CompressionCodec: kafka.Snappy.Codec(),
		Topic:            topic,
		BatchSize:        1000,    // تعداد پیام‌هایی که در هر بچ ارسال می‌شوند
		BatchBytes:       1048576, // حداکثر حجم داده (بایت) در هر بچ
		BatchTimeout:     1000,    // زمان مجاز برای ارسال هر بچ (میلی‌ثانیه)
		//ReadBackoffMin:   100,     // حداقل زمان تأخیر بین خواندن داده‌ها
		//ReadBackoffMax:   1000,    // حداکثر زمان تأخیر بین خواندن داده‌ها
		//ReadLagInterval:  0,       // زمان تأخیر بین خواندن داده‌ها
	})

	// تابعی برای ارسال پیام به Kafka
	sendToKafka := func(message string) {
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Value: []byte(message),
			},
			// اضافه کردن هر تنظیمات اضافی مورد نیاز
		)
		if err != nil {
			fmt.Printf("خطا در ارسال پیام به Kafka: %s\n", err.Error())
		} else {
			fmt.Println("پیام با موفقیت ارسال شد")
		}
	}

	// مثالی از ارسال داده به سرویس object storage و سپس ارسال آن به Kafka
	data := "داده‌هایی برای ذخیره در سرویس object storage"
	// اضافه کردن هر منطق یا پردازش مورد نیاز
	sendToKafka(data)

	// بستن نویسنده Kafka
	err := writer.Close()
	if err != nil {
		fmt.Printf("خطا در بستن نویسنده Kafka: %s\n", err.Error())
	}
}
