package const_utils

import log "unknwon.dev/clog/v2"

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
	log.Trace("[IsParameterFolder()] name: %s", name)
	if name == INPUT_DATA {
		log.Trace("[IsParameterFolder()] name: %s is %s", name, INPUT_DATA)
		return false
	}
	if name == SOURCE {
		log.Trace("[IsParameterFolder()] name: %s is %s", name, SOURCE)
		return false
	}
	if name == CI {
		log.Trace("[IsParameterFolder()] name: %s is %s", name, CI)
		return false
	}
	log.Trace("[IsParameterFolder()] name: %s is ParameterFolder", name)
	return true
}
