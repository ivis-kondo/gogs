package const_utils

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

func GetForParameters() string {
	return FOR_PARAMETERS
}

func GetWithCode() string {
	return WITH_CODE
}

/*
Folder Name in Dataset
*/
const (
	INPUT_DATA  = "input_data"
	SOURCE      = "source"
	OUTPUT_DATA = "output_data"
	CI          = "ci"
)

func Get_INPUT_DATA() string {
	return INPUT_DATA
}

func Get_SOURCE() string {
	return SOURCE
}

func Get_OUTPUT_DATA() string {
	return OUTPUT_DATA
}

func Get_CI() string {
	return CI
}

func IsParameterFolder(name string) bool {
	if name == INPUT_DATA {
		return false
	}
	if name == SOURCE {
		return false
	}
	if name == CI {
		return false
	}
	return true
}
