package utils

/*
Dataset Struct Type
*/
const (
	WITH_CODE      = "with_code"
	FOR_PARAMETERS = "for_parameters"
)

func IsDatasetStructType(struct_type string) bool {
	if struct_type == WITH_CODE || struct_type == FOR_PARAMETERS {
		return true
	}
	return false
}
