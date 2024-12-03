package utils

type ErrorIndex int

const (
	PointerTooDeep ErrorIndex = iota
	InterfaceAsReturnValueOfConstructor
	FieldIsNotExported
	InvalidTag
	MissingTypeInTag
	UnsupportedTagType
)

var chinese = map[ErrorIndex]string{
	PointerTooDeep:                      "指针层级不可超过%v层",
	InterfaceAsReturnValueOfConstructor: "节点构造函数不可使用接口类型",
	FieldIsNotExported:                  "属性%v.%v对外不可见",
	InvalidTag:                          "属性%v.%v具有不合法的标签(%v)",
	MissingTypeInTag:                    "属性%v.%v在标签中缺失类型",
	UnsupportedTagType:                  "属性%v.%v在标签中具有不支持的类型%v",
}

var english = map[ErrorIndex]string{
	PointerTooDeep:                      "the level of pointer should not be greater than %v",
	InterfaceAsReturnValueOfConstructor: "should not use interface as the return value of the constructor",
	FieldIsNotExported:                  "%v.%v is not a public field",
	InvalidTag:                          "field %v.%v has invalid tag (%v)",
	MissingTypeInTag:                    "%v.%v misses type in tag",
	UnsupportedTagType:                  "%v.%v has unsupported type %v in tag",
}

var current map[ErrorIndex]string = chinese

func SetLanguageIndex(langIdx int) {
	switch langIdx {
	case 1:
		current = chinese
	case 2:
		current = english
	default:
		panic("unknown language index")
	}
}

func ErrorMessage(idx ErrorIndex) string {
	return current[idx]
}
