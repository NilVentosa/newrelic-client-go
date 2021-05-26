// Code generated by tutone: DO NOT EDIT
package installevents

func (a *Installevents) CreateInstallEvent(
	input InstallStatus,
) (*InstallEvent, error) {

	resp := CreateInstallEventQueryResponse{}
	vars := map[string]interface{}{
		"input": input,
	}

	if err := a.client.NerdGraphQuery(CreateInstallEventMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.InstallEvent, nil
}

type CreateInstallEventQueryResponse struct {
	InstallEvent InstallEvent `json:"CreateInstallEvent"`
}

const CreateInstallEventMutation = `mutation(
	$input: InstallStatus!,
) { createInstallEvent(
	input: $input,
) {
	displayName
	entityGuid
	error {
		details
		message
	}
	name
	status
	timestamp
	validationDurationMilliseconds
} }`

func (a *Installevents) CreateInstallMetadata(
	input InputInstallMetadata,
) (*InstallMetadata, error) {

	resp := CreateInstallMetadataQueryResponse{}
	vars := map[string]interface{}{
		"input": input,
	}

	if err := a.client.NerdGraphQuery(CreateInstallMetadataMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.InstallMetadata, nil
}

type CreateInstallMetadataQueryResponse struct {
	InstallMetadata InstallMetadata `json:"CreateInstallMetadata"`
}

const CreateInstallMetadataMutation = `mutation(
	$input: InputInstallMetadata!,
) { createInstallMetadata(
	input: $input,
) {
	cliVersion
	complete
	error {
		details
		message
	}
	hostName
	kernelArch
	kernelVersion
	logFilePath
	os
	platform
	platformFamily
	platformVersion
	redirectUrl
	targetedInstall
} }`
