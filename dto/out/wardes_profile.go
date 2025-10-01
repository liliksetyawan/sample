package out

import (
    "time"
)

type WardesProfileDTOOut struct {
  Id    string   `json:"id"`
  Uuidkey    string   `json:"uuid_key"`
  Nexchiefaccountid    string   `json:"nex_chief_account_id"`
  Userid    string   `json:"user_id"`
  Username    string   `json:"username"`
  Personalname    string   `json:"personal_name"`
  Gender    string   `json:"gender"`
}

type ListWardesProfileDTOOut struct {
  Id    string   `json:"id"`
  Uuidkey    string   `json:"uuid_key"`
  Nexchiefaccountid    string   `json:"nex_chief_account_id"`
  Userid    string   `json:"user_id"`
  Username    string   `json:"username"`
  Personalname    string   `json:"personal_name"`
  Gender    string   `json:"gender"`
}