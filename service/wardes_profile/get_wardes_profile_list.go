package wardes_profile

import (
    "github.com/nexsoft-git/nexcommon/model"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dto/in"
    "github.com/nexsoft-git/nexcommon/services"
    "nexsoft.co.id/example/dto/out"
    "nexsoft.co.id/example/repository"
    out2 "github.com/nexsoft-git/nexcommon/dto/out"
    "github.com/nexsoft-git/nexlogger/log"
)

func (s *wardes_profileService) GetWardesProfileList(
    ctx *context.ContextModel,
    searchParam []model.SearchParam,
    dto in.GetListRequest,
) (
    header map[string]string,
    output interface{},
    err error,
) {

    var (
        dbResult []repository.WardesProfileModel
        result   []out.ListWardesProfileDTOOut
    )

    dbResult, err = s.wardesProfileDAO.GetListWardesProfile(ctx, dto, searchParam)
    if err != nil {
        log.Error().
            Err(err).
            Caller().
            Msg("Error Found When Get List Wardes Profile data")
        return
    }

    for i := 0; i < len(dbResult); i++ {
        repo := dbResult[i]
        result = append(result, out.ListWardesProfileDTOOut{
            ID:                           repo.ID.Int64,
            UUIDKey:                      repo.UUIDKey.String,
            NexchiefAccountID:            repo.NexchiefAccountID.Int64,
            UserID:                       repo.UserID.Int64,
            Username:                     repo.Username.String,
            PersonalName:                 repo.PersonalName.String,
            Gender:                       repo.Gender.String,
            Phone:                        repo.Phone.String,
            Email:                        repo.Email.String,
            ProvinceID:                   repo.ProvinceID.Int64,
            DistrictID:                   repo.DistrictID.Int64,
            SubDistrictID:                repo.SubDistrictID.Int64,
            UrbanVillageID:               repo.UrbanVillageID.Int64,
            Nik:                          repo.Nik.String,
            Age:                          repo.Age.String,
            Hamlet:                       repo.Hamlet.String,
            Neighbour:                    repo.Neighbour.String,
            VillageFamily:                repo.VillageFamily.String,
            VillagePopulation:            repo.VillagePopulation.String,
            BuildingLocation:             repo.BuildingLocation.String,
            BuildingOwnership:            repo.BuildingOwnership.String,
            BuildingType:                 repo.BuildingType.String,
            BuildingForm:                 repo.BuildingForm.String,
            BuildingWidthFirstFloor:      repo.BuildingWidthFirstFloor.Float64,
            BuildingLengthFirstFloor:     repo.BuildingLengthFirstFloor.Float64,
            BuildingHeightFirstFloor:     repo.BuildingHeightFirstFloor.Float64,
            BuildingWidthSecondFloor:     repo.BuildingWidthSecondFloor.Float64,
            BuildingLengthSecondFloor:    repo.BuildingLengthSecondFloor.Float64,
            BuildingHeightSecondFloor:    repo.BuildingHeightSecondFloor.Float64,
            ParkingLotWidth:              repo.ParkingLotWidth.Float64,
            ParkingLotLength:             repo.ParkingLotLength.Float64,
            VillageInternetConnection:    repo.VillageInternetConnection.String,
            InvestmentCapital:            repo.InvestmentCapital.String,
            InvestmentCapitalSource:      repo.InvestmentCapitalSource.String,
            OwnershipStatus:              repo.OwnershipStatus.String,
            OwnershipFixedAsset:          repo.OwnershipFixedAsset.String,
            PermissionCover:              repo.PermissionCover.String,
            ProfileCompletionStatus:      repo.ProfileCompletionStatus.String,
            CompanyProfileID:             repo.CompanyProfileID.Int64,
            Code:                         repo.Code.String,
            Schema:                       repo.Schema.String,
            Status:                       repo.Status.String,
            ActiveDate:                   repo.ActiveDate.Time,
            ResignDate:                   repo.ResignDate.Time,
            IsNexwise:                    repo.IsNexwise.Bool,
            DeadlineDate:                 repo.DeadlineDate.Time,
            CreatedBy:                    repo.CreatedBy.Int64,
            CreatedClient:                repo.CreatedClient.String,
            CreatedAt:                    repo.CreatedAt.Time,
            UpdatedBy:                    repo.UpdatedBy.Int64,
            UpdatedAt:                    repo.UpdatedAt.Time,
            UpdatedClient:                repo.UpdatedClient.String,
            Deleted:                      repo.Deleted.Bool,
            NewProfileApprovalStatus:     repo.NewProfileApprovalStatus.String,
        })
    }

    output = out2.DefaultResponsePayloadMessage{
        Status: out2.DefaultResponsePayloadStatus{
            Code:    "OK",
            Message: s.bundles.ReadMessageBundle("wardes_profile", "SUCCESS_GET_WARDES_PROFILE_LIST_MESSAGE", ctx.AuthAccessTokenModel.Locale, nil),
        },
        Data: result,
    }

    return
}