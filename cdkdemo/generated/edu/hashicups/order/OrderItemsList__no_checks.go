//go:build no_runtime_type_checking

package order

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OrderItemsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OrderItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OrderItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OrderItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OrderItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OrderItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OrderItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOrderItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

