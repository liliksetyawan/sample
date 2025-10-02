package out

import (
    "time"
)

type WardesProfileDTOOut struct {
  Name    string   `json:"name"`
  Address    string   `json:"address"`
  Phone    string   `json:"phone"`
  Email    string   `json:"email"`
  RecordStatus    string   `json:"record_status"`
  CreatedAt    time.Time   `json:"created_at"`
  UpdatedAt    time.Time   `json:"updated_at"`
}