//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HashicupsProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (h *jsiiProxy_HashicupsProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateHashicupsProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateHashicupsProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHashicupsProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateHashicupsProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewHashicupsProviderParameters(scope constructs.Construct, id *string, config *HashicupsProviderConfig) error {
	return nil
}

