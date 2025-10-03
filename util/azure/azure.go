package azure

import (
    "github.com/nexsoft-git/nexcommon/util"
    "github.com/nexsoft-git/nexcommon/dto/in"
)

func NewAzureHelper(
    nexsoft util.AzureConfig,
) *AzureHelper {
    return &AzureHelper{
        nexsoftAzure: util.NewAzureHelper(nexsoft),
    }
}

type AzureHelper struct {
    nexsoftAzure util.AzureHelper
}

func (a *AzureHelper) UploadNexsoftPhotoProfile(
    files []in.MultipartFile,
) (
    []string,
    error,
) {
    path := "nexchief"
    return a.nexsoftAzure.UploadFileToAzure(files, path)
}