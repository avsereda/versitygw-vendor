package s3event

// TODO: MySQL driver imported here for the future event log implementation (!)
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type EventType string

const (
	EventObjectCreated              EventType = "s3:ObjectCreated"
	EventObjectCreatedPut           EventType = "s3:ObjectCreated:Put"
	EventObjectCreatedCopy          EventType = "s3:ObjectCreated:Copy"
	EventObjectCreatedPost          EventType = "s3:ObjectCreated:Post"
	EventObjectRemoved              EventType = "s3:ObjectRemoved"
	EventObjectRemovedDelete        EventType = "s3:ObjectRemoved:Delete"
	EventObjectRemovedDeleteObjects EventType = "s3:ObjectRemoved:DeleteObjects"
	EventObjectTagging              EventType = "s3:ObjectTagging"
	EventObjectTaggingPut           EventType = "s3:ObjectTagging:Put"
	EventObjectTaggingDelete        EventType = "s3:ObjectTagging:Delete"
	EventObjectAclPut               EventType = "s3:ObjectAcl:Put"
	EventCompleteMultipartUpload    EventType = "s3:ObjectCreated:CompleteMultipartUpload"
	EventObjectRestore              EventType = "s3:ObjectRestore"
	EventObjectRestorePost          EventType = "s3:ObjectRestore:Post"
	EventObjectRestoreCompleted     EventType = "s3:ObjectRestore:Completed"
)

type S3EventSender interface {
	SendEvent(ctx *fiber.Ctx, meta EventMeta) error
	Close() error
}

// FIXME: Original event sender was removed and should be replaced with vendor specific one if required
type NullEventSender struct{}

type EventMeta struct {
	EventName   EventType
	Bucket      string
	Key         string
	NewKey      string
	BucketOwner string
	ObjectSize  int64
	ObjectETag  string
	VersionId   string
}

func (NullEventSender) SendEvent(ctx *fiber.Ctx, meta EventMeta) error {
	// TODO: This method should implement event logging to the existing log table
	return nil
}

func (NullEventSender) Close() error {
	return nil
}
