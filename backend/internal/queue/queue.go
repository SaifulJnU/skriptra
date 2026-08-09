// Package queue carries ingestion work from the API to the workers.
package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	streamName  = "SKRIPTRA"
	subjectSet  = "skriptra.>"
	SubjectFile = "skriptra.document.uploaded"
	consumerID  = "ingest-workers"
)

// IngestMessage is the work item. It carries a storage key rather than the file
// bytes: NATS is a coordination system, not a blob store, and a 50 MB message
// would be an abuse of it.
type IngestMessage struct {
	DocumentID uuid.UUID `json:"documentId"`
	CourseID   uuid.UUID `json:"courseId"`
	Filename   string    `json:"filename"`
	StorageKey string    `json:"storageKey"`
}

type Queue struct {
	conn   *nats.Conn
	js     jetstream.JetStream
	stream jetstream.Stream
}

// Connect opens a JetStream-backed connection.
//
// JetStream rather than core NATS because ingestion jobs must survive a worker
// restart. Core NATS is fire-and-forget: a crash mid-ingest would leave a
// document stuck at "queued" forever with nothing to retry it.
func Connect(ctx context.Context, url string) (*Queue, error) {
	conn, err := nats.Connect(url,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("connect nats: %w", err)
	}

	js, err := jetstream.New(conn)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("jetstream: %w", err)
	}

	stream, err := js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  []string{subjectSet},
		Retention: jetstream.WorkQueuePolicy,
		MaxAge:    24 * time.Hour,
	})
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("create stream: %w", err)
	}

	return &Queue{conn: conn, js: js, stream: stream}, nil
}

func (q *Queue) Close() {
	if q.conn != nil {
		q.conn.Drain() //nolint:errcheck // best effort on shutdown
	}
}

func (q *Queue) PublishIngest(ctx context.Context, documentID, courseID uuid.UUID, filename, storageKey string) error {
	payload, err := json.Marshal(IngestMessage{
		DocumentID: documentID, CourseID: courseID,
		Filename: filename, StorageKey: storageKey,
	})
	if err != nil {
		return err
	}
	_, err = q.js.Publish(ctx, SubjectFile, payload)
	return err
}

// Consume runs handler for each job until ctx is cancelled.
//
// A durable pull consumer, so several workers can share the queue and an
// unacknowledged message is redelivered rather than lost.
func (q *Queue) Consume(ctx context.Context, handler func(context.Context, IngestMessage) error) error {
	consumer, err := q.stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:       consumerID,
		AckPolicy:     jetstream.AckExplicitPolicy,
		MaxDeliver:    3,
		AckWait:       10 * time.Minute, // a long paper on a local model is slow
		FilterSubject: SubjectFile,
	})
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}

	sub, err := consumer.Consume(func(msg jetstream.Msg) {
		var job IngestMessage
		if err := json.Unmarshal(msg.Data(), &job); err != nil {
			// Unparseable payloads will never succeed. Terminate rather than
			// redeliver them until MaxDeliver is exhausted.
			_ = msg.Term()
			return
		}
		if err := handler(ctx, job); err != nil {
			// Nak triggers redelivery; the pipeline has already recorded the
			// failure against the document either way.
			_ = msg.Nak()
			return
		}
		_ = msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("consume: %w", err)
	}
	defer sub.Stop()

	<-ctx.Done()
	return nil
}
