package out

import (
    "time"
)

type PersonProfileDTOOut struct {
  UuidKey    string   `json:"uuid_key"`
  OtherField1    string   `json:"other_field_1"`
  OtherField2    int64   `json:"other_field_2"`
}

type ListPersonProfileDTOOut struct {
  Id    string   `json:"id"`
  PersonProfileId    string   `json:"person_profile_id"`
  TitleId    string   `json:"title_id"`
  TitleName    string   `json:"title_name"`
  FirstName    string   `json:"first_name"`
  LastName    string   `json:"last_name"`
  Sex    string   `json:"sex"`
  Nik    string   `json:"nik"`
  Address1    string   `json:"address_1"`
  Address2    string   `json:"address_2"`
  Address3    string   `json:"address_3"`
}