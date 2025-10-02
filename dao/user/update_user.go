package user

import (
    "fmt"
    "database/sql"
    "github.com/nexsoft-git/nexcommon/context"
    "github.com/nexsoft-git/nexcommon/dao"
    "nexsoft.co.id/example/repository"
)

// Descriptions: Update an existing user record
func (d *userPostgresqlSQLDAO) UpdateUser(
     ctx *context.ContextModel,
     tx *sql.Tx,
     dtoIn repository.UserModel,
) (
     error, 
) {
    query := fmt.Sprintf(`
        UPDATE %s SET
            auth_user_id = $1, person_profile_id = $2, client_id = $3, user_code = $4,
            username = $5, last_token = $6, signature_key = $7, ip_whitelist = $8,
            device = $9, level = $10, locale = $11, additional_info = $12, alias = $13,
            is_user_nex_acc = $14, mobile_access = $15, website_access = $16,
            wdi_portal_access = $17, domain_nd6 = $18, username_nd6 = $19,
            password_nd6 = $20, domain_nexvin = $21, username_nexvin = $22,
            password_nexvin = $23, updated_by = $24, updated_at = $25, updated_client = $26
        WHERE uuid_key = $27
    `, dao.GetDBTable(ctx, "user"))

    stmt, err := tx.Prepare(query)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(
        dtoIn.AuthUserID.Int64, dtoIn.PersonProfileID.Int64, dtoIn.ClientID.String, dtoIn.UserCode.String,
        dtoIn.Username.String, dtoIn.LastToken.String, dtoIn.SignatureKey.String, dtoIn.IPWhitelist.String,
        dtoIn.Device.String, dtoIn.Level.Int64, dtoIn.Locale.String, dtoIn.AdditionalInfo.String,
        dtoIn.Alias.String, dtoIn.IsUserNexAcc.Bool, dtoIn.MobileAccess.Bool, dtoIn.WebsiteAccess.Bool,
        dtoIn.WDIPortalAccess.Bool, dtoIn.DomainNd6.String, dtoIn.UsernameNd6.String,
        dtoIn.PasswordNd6.String, dtoIn.DomainNexvin.String, dtoIn.UsernameNexvin.String,
        dtoIn.PasswordNexvin.String, ctx.Limitation.UserID, dtoIn.UpdatedAt.Time,
        ctx.AuthAccessTokenModel.ClientID, dtoIn.UUIDKey.String,
    )
    return err
}
