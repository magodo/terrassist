package output

import types "types"

func expandTypeNamedBoolSliceAlias(input []interface{}) types.TypeNamedBoolSliceAlias {
	if len(input) == 0 {
		return nil
	}
	output := make(types.TypeNamedBoolSliceAlias, 0)
	for _, elem := range input {
		output = append(output, elem.(types.Bool))
	}
	return output
}
