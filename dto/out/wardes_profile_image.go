package out

import (
    "time"
)

type WardesProfileImageDTOOut struct {
  UuidKey    string   `json:"uuid_key"`
  ProfileImage    string   `json:"profile_image"`
  Description    string   `json:"description"`
  RecordStatus    string   `json:"record_status"`
  CreatedAt    time.Time   `json:"created_at"`
  UpdatedAt    time.Time   `json:"updated_at"`
}