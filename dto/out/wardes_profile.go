package out

import (
    "time"
)

type WardesProfileDTOOut struct {
  Id    int64   `json:"id"`
  UuidKey    string   `json:"uuid_key"`
  NexchiefAccountId    int64   `json:"nexchief_account_id"`
  UserId    int64   `json:"user_id"`
  Username    string   `json:"username"`
  PersonalName    string   `json:"personal_name"`
  Gender    string   `json:"gender"`
  Phone    string   `json:"phone"`
  Email    string   `json:"email"`
  ProvinceId    int64   `json:"province_id"`
  DistrictId    int64   `json:"district_id"`
  SubDistrictId    int64   `json:"sub_district_id"`
  UrbanVillageId    int64   `json:"urban_village_id"`
  Nik    string   `json:"nik"`
  Age    string   `json:"age"`
  Hamlet    string   `json:"hamlet"`
  Neighbour    string   `json:"neighbour"`
  VillageFamily    string   `json:"village_family"`
  VillagePopulation    string   `json:"village_population"`
  BuildingLocation    string   `json:"building_location"`
  BuildingOwnership    string   `json:"building_ownership"`
  BuildingType    string   `json:"building_type"`
  BuildingForm    string   `json:"building_form"`
  BuildingWidthFirstFloor    float64   `json:"building_width_first_floor"`
  BuildingLengthFirstFloor    float64   `json:"building_length_first_floor"`
  BuildingHeightFirstFloor    float64   `json:"building_height_first_floor"`
  BuildingWidthSecondFloor    float64   `json:"building_width_second_floor"`
  BuildingLengthSecondFloor    float64   `json:"building_length_second_floor"`
  BuildingHeightSecondFloor    float64   `json:"building_height_second_floor"`
  ParkingLotWidth    float64   `json:"parking_lot_width"`
  ParkingLotLength    float64   `json:"parking_lot_length"`
  VillageInternetConnection    string   `json:"village_internet_connection"`
  InvestmentCapital    string   `json:"investment_capital"`
  InvestmentCapitalSource    string   `json:"investment_capital_source"`
  OwnershipStatus    string   `json:"ownership_status"`
  OwnershipFixedAsset    string   `json:"ownership_fixed_asset"`
  PermissionCover    string   `json:"permission_cover"`
  ProfileCompletionStatus    string   `json:"profile_completion_status"`
  CompanyProfileId    int64   `json:"company_profile_id"`
  Code    string   `json:"code"`
  Schema    string   `json:"schema"`
  Status    string   `json:"status"`
  ActiveDate    time.Time   `json:"active_date"`
  ResignDate    time.Time   `json:"resign_date"`
  IsNexwise    bool   `json:"is_nexwise"`
  DeadlineDate    time.Time   `json:"deadline_date"`
  CreatedBy    int64   `json:"created_by"`
  CreatedClient    string   `json:"created_client"`
  CreatedAt    time.Time   `json:"created_at"`
  UpdatedBy    int64   `json:"updated_by"`
  UpdatedAt    time.Time   `json:"updated_at"`
  UpdatedClient    string   `json:"updated_client"`
  Deleted    bool   `json:"deleted"`
  NewProfileApprovalStatus    string   `json:"new_profile_approval_status"`
}