package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type IKafka interface {
	Close()
	ProduceMessage(string, string, int32) error
	ConsumeMessages(string) error
}

type kafka struct {
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
}

type consumerGroupHandler struct {
}

func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		switch message.Partition {
		case 0:

		}
	}

	return nil
}

func NewIKafka() (IKafka, error) {
	producer, err := newKafkaProducer()
	if err != nil {
		return nil, err
	}

	consumer, err := newKafkaConsumer()
	if err != nil {
		return nil, err
	}

	return &kafka{
		producer: producer,
		consumer: consumer,
	}, nil
}

func (k *kafka) ProduceMessage(topic, value string, partion int32) error {
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(value),
		Partition: partion,
	}
	_, _, err := k.producer.SendMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func (k *kafka) ConsumeMessages(topic string) error {
	handler := consumerGroupHandler{}
	for {
		err := k.consumer.Consume(context.Background(), []string{topic}, &handler)
		if err != nil {
			return err
		}
	}
}

func (k *kafka) Close() {
	k.producer.Close()
	k.consumer.Close()
}

func newKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	
	producer, err := sarama.NewSyncProducer([]string{"kafka:29092"}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func newKafkaConsumer() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup([]string{"kafka:29092"}, "", config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
