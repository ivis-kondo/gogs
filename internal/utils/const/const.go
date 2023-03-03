package utils

/*
Dataset Struct Type
*/
const (
	WITH_CODE     = "with_code"
	FOR_PARAMETER = "for_parameter"
)

func IsDatasetStructType(struct_type string) bool {
	if struct_type == WITH_CODE || struct_type == FOR_PARAMETER {
		return true
	}
	return false
}
