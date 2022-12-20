package entity

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PK uint

func (pk PK) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, uint32(pk))
	return buf.Bytes(), err
}

func (pk *PK) UnmarshalBinary(b []byte) error {
	var result uint32
	var buf bytes.Buffer
	buf.Write(b)
	err := binary.Read(&buf, binary.BigEndian, &result)
	*pk = PK(result)
	return err
}

type FileID struct {
	primitive.ObjectID
}

func (id FileID) String() string {
	s := id.ObjectID.String()
	return s[10 : len(s)-2]
}

type File struct {
	File         io.Reader
	OrigFileName string
}

var NilFileID FileID = FileID{primitive.NilObjectID}

func ParseFileID(s string) (FileID, error) {
	objID, err := primitive.ObjectIDFromHex(s)
	return FileID{objID}, err
}

type Form struct {
	tableName struct{} `pg:"forms"`

	Id PK `pg:"id,pk" json:"id" form:"id"`

	UserId PK          `pg:"user_id" json:"-"`
	User   *AuthedUser `pg:"rel:has-one" json:"sent_by"`

	Message string    `pg:"message" json:"message" form:"message"`
	SentAt  time.Time `pg:"sent_at" json:"sent_at"`
}

type Review struct {
	tableName struct{} `pg:"reviews"`

	Id PK `pg:"id,pk" json:"id" form:"id"`

	UserId PK          `pg:"user_id" json:"-"`
	User   *AuthedUser `pg:"rel:has-one" json:"sent_by"`

	Rating   int       `pg:"rating" json:"rating" form:"rating"`
	Message  *string   `pg:"message" json:"message,omitempty" form:"message"`
	PostedAt time.Time `pg:"posted_at" json:"posted_at"`
}

type AuthedUser struct {
	tableName struct{} `pg:"users"`

	Id PK `pg:"id,pk" json:"id" form:"id" redis:"id"`

	UserName string `pg:"username" json:"username" redis:"username"`
	Password string `pg:"password" json:"-" redis:"password"`

	Email   string `pg:"email" json:"email" redis:"email"`
	Name    string `pg:"name" json:"name" redis:"name"`
	IsAdmin bool   `pg:"is_admin" json:"-" redis:"is_admin"`
}

type AuthToken struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type ErrResponse struct {
	StatusCode int
	Err        error
}

func (e *ErrResponse) Error() string {
	return e.Err.Error()
}

func (e *ErrResponse) Unwrap() error {
	return e.Err
}

type ServerError struct {
	Message   string `json:"message"`
	Location  string `json:"location"`
	ErrorCode int    `json:"code"`
	BaseError error  `json:"-"`
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%s; %s: %s (%d)", e.Message, e.Location, e.BaseError.Error(), e.ErrorCode)
}

func (e *ServerError) Unwrap() error {
	return e.BaseError
}
