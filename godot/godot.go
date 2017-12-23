// WARNING: This file has automatically been generated on Sun, 24 Dec 2017 06:36:13 JST.
// By https://git.io/c-for-go. DO NOT EDIT.

package godot

/*
#cgo CFLAGS: -Igodot_headers -std=c11
#cgo darwin LDFLAGS: -framework Cocoa -Wl,-undefined,dynamic_lookup
#include "gdnative_api_struct.gen.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// GodotObjectDestroy function as declared in gdnative/gdnative.h:207
func GodotObjectDestroy(pO *GodotObject) {
	cpO, _ := (C.godot_object)(unsafe.Pointer(pO)), cgoAllocsUnknown
	C.godot_object_destroy(cpO)
}

// GodotGlobalGetSingleton function as declared in gdnative/gdnative.h:215
func GodotGlobalGetSingleton(pName []byte) *GodotObject {
	cpName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pName)).Data)), cgoAllocsUnknown
	__ret := C.godot_global_get_singleton(cpName)
	__v := *(**GodotObject)(unsafe.Pointer(&__ret))
	return __v
}

// GodotGetStackBottom function as declared in gdnative/gdnative.h:219
func GodotGetStackBottom() unsafe.Pointer {
	__ret := C.godot_get_stack_bottom()
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// GodotMethodBindGetMethod function as declared in gdnative/gdnative.h:227
func GodotMethodBindGetMethod(pClassname string, pMethodname string) *GodotMethodBind {
	cpClassname, _ := unpackPCharString(pClassname)
	cpMethodname, _ := unpackPCharString(pMethodname)
	__ret := C.godot_method_bind_get_method(cpClassname, cpMethodname)
	__v := NewGodotMethodBindRef(unsafe.Pointer(__ret))
	return __v
}

// GodotMethodBindPtrcall function as declared in gdnative/gdnative.h:228
func GodotMethodBindPtrcall(pMethodBind []GodotMethodBind, pInstance *GodotObject, pArgs []unsafe.Pointer, pRet unsafe.Pointer) {
	cpMethodBind, _ := unpackArgSGodotMethodBind(pMethodBind)
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	cpArgs, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pArgs)).Data)), cgoAllocsUnknown
	cpRet, _ := pRet, cgoAllocsUnknown
	C.godot_method_bind_ptrcall(cpMethodBind, cpInstance, cpArgs, cpRet)
	packSGodotMethodBind(pMethodBind, cpMethodBind)
}

// GodotMethodBindCall function as declared in gdnative/gdnative.h:229
func GodotMethodBindCall(pMethodBind []GodotMethodBind, pInstance *GodotObject, pArgs [][]GodotVariant, pArgCount int32, pCallError []GodotVariantCallError) GodotVariant {
	cpMethodBind, _ := unpackArgSGodotMethodBind(pMethodBind)
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	cpArgs, _ := unpackArgSSGodotVariant(pArgs)
	cpArgCount, _ := (C.int)(pArgCount), cgoAllocsUnknown
	cpCallError, _ := unpackArgSGodotVariantCallError(pCallError)
	__ret := C.godot_method_bind_call(cpMethodBind, cpInstance, cpArgs, cpArgCount, cpCallError)
	packSGodotVariantCallError(pCallError, cpCallError)
	packSSGodotVariant(pArgs, cpArgs)
	packSGodotMethodBind(pMethodBind, cpMethodBind)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotGetClassConstructor function as declared in gdnative/gdnative.h:266
func GodotGetClassConstructor(pClassname string) GodotClassConstructor {
	cpClassname, _ := unpackPCharString(pClassname)
	__ret := C.godot_get_class_constructor(cpClassname)
	__v := *NewGodotClassConstructorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotGetGlobalConstants function as declared in gdnative/gdnative.h:268
func GodotGetGlobalConstants() GodotDictionary {
	__ret := C.godot_get_global_constants()
	__v := *NewGodotDictionaryRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotRegisterNativeCallType function as declared in gdnative/gdnative.h:278
func GodotRegisterNativeCallType(pCallType string, pCallback NativeCallCb) {
	cpCallType, _ := unpackPCharString(pCallType)
	cpCallback, _ := pCallback.PassValue()
	C.godot_register_native_call_type(cpCallType, cpCallback)
}

// GodotAlloc function as declared in gdnative/gdnative.h:281
func GodotAlloc(pBytes int32) unsafe.Pointer {
	cpBytes, _ := (C.int)(pBytes), cgoAllocsUnknown
	__ret := C.godot_alloc(cpBytes)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// GodotRealloc function as declared in gdnative/gdnative.h:282
func GodotRealloc(pPtr unsafe.Pointer, pBytes int32) unsafe.Pointer {
	cpPtr, _ := pPtr, cgoAllocsUnknown
	cpBytes, _ := (C.int)(pBytes), cgoAllocsUnknown
	__ret := C.godot_realloc(cpPtr, cpBytes)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// GodotFree function as declared in gdnative/gdnative.h:283
func GodotFree(pPtr unsafe.Pointer) {
	cpPtr, _ := pPtr, cgoAllocsUnknown
	C.godot_free(cpPtr)
}

// GodotPrintError function as declared in gdnative/gdnative.h:286
func GodotPrintError(pDescription string, pFunction string, pFile string, pLine int32) {
	cpDescription, _ := unpackPCharString(pDescription)
	cpFunction, _ := unpackPCharString(pFunction)
	cpFile, _ := unpackPCharString(pFile)
	cpLine, _ := (C.int)(pLine), cgoAllocsUnknown
	C.godot_print_error(cpDescription, cpFunction, cpFile, cpLine)
}

// GodotPrintWarning function as declared in gdnative/gdnative.h:287
func GodotPrintWarning(pDescription string, pFunction string, pFile string, pLine int32) {
	cpDescription, _ := unpackPCharString(pDescription)
	cpFunction, _ := unpackPCharString(pFunction)
	cpFile, _ := unpackPCharString(pFile)
	cpLine, _ := (C.int)(pLine), cgoAllocsUnknown
	C.godot_print_warning(cpDescription, cpFunction, cpFile, cpLine)
}

// GodotPrint function as declared in gdnative/gdnative.h:288
func GodotPrint(pMessage []GodotString) {
	cpMessage, _ := unpackArgSGodotString(pMessage)
	C.godot_print(cpMessage)
	packSGodotString(pMessage, cpMessage)
}

// GodotStringNew function as declared in gdnative/string.h:62
func GodotStringNew(rDest []GodotString) {
	crDest, _ := unpackArgSGodotString(rDest)
	C.godot_string_new(crDest)
	packSGodotString(rDest, crDest)
}

// GodotStringNewCopy function as declared in gdnative/string.h:63
func GodotStringNewCopy(rDest []GodotString, pSrc []GodotString) {
	crDest, _ := unpackArgSGodotString(rDest)
	cpSrc, _ := unpackArgSGodotString(pSrc)
	C.godot_string_new_copy(crDest, cpSrc)
	packSGodotString(pSrc, cpSrc)
	packSGodotString(rDest, crDest)
}

// GodotStringNewData function as declared in gdnative/string.h:64
func GodotStringNewData(rDest []GodotString, pContents string, pSize int32) {
	crDest, _ := unpackArgSGodotString(rDest)
	cpContents, _ := unpackPCharString(pContents)
	cpSize, _ := (C.int)(pSize), cgoAllocsUnknown
	C.godot_string_new_data(crDest, cpContents, cpSize)
	packSGodotString(rDest, crDest)
}

// GodotStringNewUnicodeData function as declared in gdnative/string.h:65
func GodotStringNewUnicodeData(rDest []GodotString, pContents []int32, pSize int32) {
	crDest, _ := unpackArgSGodotString(rDest)
	cpContents, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pContents)).Data)), cgoAllocsUnknown
	cpSize, _ := (C.int)(pSize), cgoAllocsUnknown
	C.godot_string_new_unicode_data(crDest, cpContents, cpSize)
	packSGodotString(rDest, crDest)
}

// GodotStringGetData function as declared in gdnative/string.h:67
func GodotStringGetData(pSelf []GodotString, pDest []byte, pSize []int32) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpDest, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pDest)).Data)), cgoAllocsUnknown
	cpSize, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pSize)).Data)), cgoAllocsUnknown
	C.godot_string_get_data(cpSelf, cpDest, cpSize)
	packSGodotString(pSelf, cpSelf)
}

// GodotStringOperatorIndex function as declared in gdnative/string.h:69
func GodotStringOperatorIndex(pSelf []GodotString, pIdx GodotInt) *int32 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_operator_index(cpSelf, cpIdx)
	packSGodotString(pSelf, cpSelf)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringOperatorIndexConst function as declared in gdnative/string.h:70
func GodotStringOperatorIndexConst(pSelf []GodotString, pIdx GodotInt) int32 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_operator_index_const(cpSelf, cpIdx)
	packSGodotString(pSelf, cpSelf)
	__v := (int32)(__ret)
	return __v
}

// GodotStringUnicodeStr function as declared in gdnative/string.h:71
func GodotStringUnicodeStr(pSelf []GodotString) *int32 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_unicode_str(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringOperatorEqual function as declared in gdnative/string.h:73
func GodotStringOperatorEqual(pSelf []GodotString, pB []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpB, _ := unpackArgSGodotString(pB)
	__ret := C.godot_string_operator_equal(cpSelf, cpB)
	packSGodotString(pB, cpB)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringOperatorLess function as declared in gdnative/string.h:74
func GodotStringOperatorLess(pSelf []GodotString, pB []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpB, _ := unpackArgSGodotString(pB)
	__ret := C.godot_string_operator_less(cpSelf, cpB)
	packSGodotString(pB, cpB)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringOperatorPlus function as declared in gdnative/string.h:75
func GodotStringOperatorPlus(pSelf []GodotString, pB []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpB, _ := unpackArgSGodotString(pB)
	__ret := C.godot_string_operator_plus(cpSelf, cpB)
	packSGodotString(pB, cpB)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringLength function as declared in gdnative/string.h:79
func GodotStringLength(pSelf []GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_length(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringBeginsWith function as declared in gdnative/string.h:83
func GodotStringBeginsWith(pSelf []GodotString, pString []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpString, _ := unpackArgSGodotString(pString)
	__ret := C.godot_string_begins_with(cpSelf, cpString)
	packSGodotString(pString, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringBeginsWithCharArray function as declared in gdnative/string.h:84
func GodotStringBeginsWithCharArray(pSelf []GodotString, pCharArray string) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpCharArray, _ := unpackPCharString(pCharArray)
	__ret := C.godot_string_begins_with_char_array(cpSelf, cpCharArray)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringBigrams function as declared in gdnative/string.h:85
func GodotStringBigrams(pSelf []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_bigrams(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringChr function as declared in gdnative/string.h:86
func GodotStringChr(pCharacter int32) GodotString {
	cpCharacter, _ := (C.wchar_t)(pCharacter), cgoAllocsUnknown
	__ret := C.godot_string_chr(cpCharacter)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringEndsWith function as declared in gdnative/string.h:87
func GodotStringEndsWith(pSelf []GodotString, pString []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpString, _ := unpackArgSGodotString(pString)
	__ret := C.godot_string_ends_with(cpSelf, cpString)
	packSGodotString(pString, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringFind function as declared in gdnative/string.h:88
func GodotStringFind(pSelf []GodotString, pWhat GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_find(cpSelf, cpWhat)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindFrom function as declared in gdnative/string.h:89
func GodotStringFindFrom(pSelf []GodotString, pWhat GodotString, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_find_from(cpSelf, cpWhat, cpFrom)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindmk function as declared in gdnative/string.h:90
func GodotStringFindmk(pSelf []GodotString, pKeys []GodotArray) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKeys, _ := unpackArgSGodotArray(pKeys)
	__ret := C.godot_string_findmk(cpSelf, cpKeys)
	packSGodotArray(pKeys, cpKeys)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindmkFrom function as declared in gdnative/string.h:91
func GodotStringFindmkFrom(pSelf []GodotString, pKeys []GodotArray, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKeys, _ := unpackArgSGodotArray(pKeys)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_findmk_from(cpSelf, cpKeys, cpFrom)
	packSGodotArray(pKeys, cpKeys)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindmkFromInPlace function as declared in gdnative/string.h:92
func GodotStringFindmkFromInPlace(pSelf []GodotString, pKeys []GodotArray, pFrom GodotInt, rKey []GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKeys, _ := unpackArgSGodotArray(pKeys)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	crKey, _ := (*C.godot_int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&rKey)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_findmk_from_in_place(cpSelf, cpKeys, cpFrom, crKey)
	packSGodotArray(pKeys, cpKeys)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindn function as declared in gdnative/string.h:93
func GodotStringFindn(pSelf []GodotString, pWhat GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_findn(cpSelf, cpWhat)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindnFrom function as declared in gdnative/string.h:94
func GodotStringFindnFrom(pSelf []GodotString, pWhat GodotString, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_findn_from(cpSelf, cpWhat, cpFrom)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFindLast function as declared in gdnative/string.h:95
func GodotStringFindLast(pSelf []GodotString, pWhat GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_find_last(cpSelf, cpWhat)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringFormat function as declared in gdnative/string.h:96
func GodotStringFormat(pSelf []GodotString, pValues []GodotVariant) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpValues, _ := unpackArgSGodotVariant(pValues)
	__ret := C.godot_string_format(cpSelf, cpValues)
	packSGodotVariant(pValues, cpValues)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringFormatWithCustomPlaceholder function as declared in gdnative/string.h:97
func GodotStringFormatWithCustomPlaceholder(pSelf []GodotString, pValues []GodotVariant, pPlaceholder string) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpValues, _ := unpackArgSGodotVariant(pValues)
	cpPlaceholder, _ := unpackPCharString(pPlaceholder)
	__ret := C.godot_string_format_with_custom_placeholder(cpSelf, cpValues, cpPlaceholder)
	packSGodotVariant(pValues, cpValues)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHexEncodeBuffer function as declared in gdnative/string.h:98
func GodotStringHexEncodeBuffer(pBuffer string, pLen GodotInt) GodotString {
	cpBuffer, _ := unpackPUint8TString(pBuffer)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hex_encode_buffer(cpBuffer, cpLen)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHexToInt function as declared in gdnative/string.h:99
func GodotStringHexToInt(pSelf []GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hex_to_int(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringHexToIntWithoutPrefix function as declared in gdnative/string.h:100
func GodotStringHexToIntWithoutPrefix(pSelf []GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hex_to_int_without_prefix(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringInsert function as declared in gdnative/string.h:101
func GodotStringInsert(pSelf []GodotString, pAtPos GodotInt, pString GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpAtPos, _ := (C.godot_int)(pAtPos), cgoAllocsUnknown
	cpString, _ := pString.PassValue()
	__ret := C.godot_string_insert(cpSelf, cpAtPos, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringIsNumeric function as declared in gdnative/string.h:102
func GodotStringIsNumeric(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_numeric(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsSubsequenceOf function as declared in gdnative/string.h:103
func GodotStringIsSubsequenceOf(pSelf []GodotString, pString []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpString, _ := unpackArgSGodotString(pString)
	__ret := C.godot_string_is_subsequence_of(cpSelf, cpString)
	packSGodotString(pString, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsSubsequenceOfi function as declared in gdnative/string.h:104
func GodotStringIsSubsequenceOfi(pSelf []GodotString, pString []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpString, _ := unpackArgSGodotString(pString)
	__ret := C.godot_string_is_subsequence_ofi(cpSelf, cpString)
	packSGodotString(pString, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringLpad function as declared in gdnative/string.h:105
func GodotStringLpad(pSelf []GodotString, pMinLength GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	__ret := C.godot_string_lpad(cpSelf, cpMinLength)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringLpadWithCustomCharacter function as declared in gdnative/string.h:106
func GodotStringLpadWithCustomCharacter(pSelf []GodotString, pMinLength GodotInt, pCharacter []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	cpCharacter, _ := unpackArgSGodotString(pCharacter)
	__ret := C.godot_string_lpad_with_custom_character(cpSelf, cpMinLength, cpCharacter)
	packSGodotString(pCharacter, cpCharacter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringMatch function as declared in gdnative/string.h:107
func GodotStringMatch(pSelf []GodotString, pWildcard []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWildcard, _ := unpackArgSGodotString(pWildcard)
	__ret := C.godot_string_match(cpSelf, cpWildcard)
	packSGodotString(pWildcard, cpWildcard)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringMatchn function as declared in gdnative/string.h:108
func GodotStringMatchn(pSelf []GodotString, pWildcard []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWildcard, _ := unpackArgSGodotString(pWildcard)
	__ret := C.godot_string_matchn(cpSelf, cpWildcard)
	packSGodotString(pWildcard, cpWildcard)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringMd5 function as declared in gdnative/string.h:109
func GodotStringMd5(pMd5 string) GodotString {
	cpMd5, _ := unpackPUint8TString(pMd5)
	__ret := C.godot_string_md5(cpMd5)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNum function as declared in gdnative/string.h:110
func GodotStringNum(pNum float64) GodotString {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num(cpNum)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNumInt64 function as declared in gdnative/string.h:111
func GodotStringNumInt64(pNum int64, pBase GodotInt) GodotString {
	cpNum, _ := (C.int64_t)(pNum), cgoAllocsUnknown
	cpBase, _ := (C.godot_int)(pBase), cgoAllocsUnknown
	__ret := C.godot_string_num_int64(cpNum, cpBase)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNumInt64Capitalized function as declared in gdnative/string.h:112
func GodotStringNumInt64Capitalized(pNum int64, pBase GodotInt, pCapitalizeHex GodotBool) GodotString {
	cpNum, _ := (C.int64_t)(pNum), cgoAllocsUnknown
	cpBase, _ := (C.godot_int)(pBase), cgoAllocsUnknown
	cpCapitalizeHex, _ := (C.godot_bool)(pCapitalizeHex), cgoAllocsUnknown
	__ret := C.godot_string_num_int64_capitalized(cpNum, cpBase, cpCapitalizeHex)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNumReal function as declared in gdnative/string.h:113
func GodotStringNumReal(pNum float64) GodotString {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num_real(cpNum)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNumScientific function as declared in gdnative/string.h:114
func GodotStringNumScientific(pNum float64) GodotString {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num_scientific(cpNum)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNumWithDecimals function as declared in gdnative/string.h:115
func GodotStringNumWithDecimals(pNum float64, pDecimals GodotInt) GodotString {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	cpDecimals, _ := (C.godot_int)(pDecimals), cgoAllocsUnknown
	__ret := C.godot_string_num_with_decimals(cpNum, cpDecimals)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringPadDecimals function as declared in gdnative/string.h:116
func GodotStringPadDecimals(pSelf []GodotString, pDigits GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpDigits, _ := (C.godot_int)(pDigits), cgoAllocsUnknown
	__ret := C.godot_string_pad_decimals(cpSelf, cpDigits)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringPadZeros function as declared in gdnative/string.h:117
func GodotStringPadZeros(pSelf []GodotString, pDigits GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpDigits, _ := (C.godot_int)(pDigits), cgoAllocsUnknown
	__ret := C.godot_string_pad_zeros(cpSelf, cpDigits)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringReplaceFirst function as declared in gdnative/string.h:118
func GodotStringReplaceFirst(pSelf []GodotString, pKey GodotString, pWith GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replace_first(cpSelf, cpKey, cpWith)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringReplace function as declared in gdnative/string.h:119
func GodotStringReplace(pSelf []GodotString, pKey GodotString, pWith GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replace(cpSelf, cpKey, cpWith)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringReplacen function as declared in gdnative/string.h:120
func GodotStringReplacen(pSelf []GodotString, pKey GodotString, pWith GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replacen(cpSelf, cpKey, cpWith)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringRfind function as declared in gdnative/string.h:121
func GodotStringRfind(pSelf []GodotString, pWhat GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_rfind(cpSelf, cpWhat)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringRfindn function as declared in gdnative/string.h:122
func GodotStringRfindn(pSelf []GodotString, pWhat GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_rfindn(cpSelf, cpWhat)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringRfindFrom function as declared in gdnative/string.h:123
func GodotStringRfindFrom(pSelf []GodotString, pWhat GodotString, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_rfind_from(cpSelf, cpWhat, cpFrom)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringRfindnFrom function as declared in gdnative/string.h:124
func GodotStringRfindnFrom(pSelf []GodotString, pWhat GodotString, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_rfindn_from(cpSelf, cpWhat, cpFrom)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringRpad function as declared in gdnative/string.h:125
func GodotStringRpad(pSelf []GodotString, pMinLength GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	__ret := C.godot_string_rpad(cpSelf, cpMinLength)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringRpadWithCustomCharacter function as declared in gdnative/string.h:126
func GodotStringRpadWithCustomCharacter(pSelf []GodotString, pMinLength GodotInt, pCharacter []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	cpCharacter, _ := unpackArgSGodotString(pCharacter)
	__ret := C.godot_string_rpad_with_custom_character(cpSelf, cpMinLength, cpCharacter)
	packSGodotString(pCharacter, cpCharacter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSimilarity function as declared in gdnative/string.h:127
func GodotStringSimilarity(pSelf []GodotString, pString []GodotString) GodotReal {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpString, _ := unpackArgSGodotString(pString)
	__ret := C.godot_string_similarity(cpSelf, cpString)
	packSGodotString(pString, cpString)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotStringSprintf function as declared in gdnative/string.h:128
func GodotStringSprintf(pSelf []GodotString, pValues []GodotArray, pError []GodotBool) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpValues, _ := unpackArgSGodotArray(pValues)
	cpError, _ := (*C.godot_bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pError)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_sprintf(cpSelf, cpValues, cpError)
	packSGodotArray(pValues, cpValues)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSubstr function as declared in gdnative/string.h:129
func GodotStringSubstr(pSelf []GodotString, pFrom GodotInt, pChars GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	cpChars, _ := (C.godot_int)(pChars), cgoAllocsUnknown
	__ret := C.godot_string_substr(cpSelf, cpFrom, cpChars)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringToDouble function as declared in gdnative/string.h:130
func GodotStringToDouble(pSelf []GodotString) float64 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_double(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (float64)(__ret)
	return __v
}

// GodotStringToFloat function as declared in gdnative/string.h:131
func GodotStringToFloat(pSelf []GodotString) GodotReal {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_float(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotStringToInt function as declared in gdnative/string.h:132
func GodotStringToInt(pSelf []GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_int(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringCamelcaseToUnderscore function as declared in gdnative/string.h:134
func GodotStringCamelcaseToUnderscore(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_camelcase_to_underscore(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCamelcaseToUnderscoreLowercased function as declared in gdnative/string.h:135
func GodotStringCamelcaseToUnderscoreLowercased(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_camelcase_to_underscore_lowercased(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCapitalize function as declared in gdnative/string.h:136
func GodotStringCapitalize(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_capitalize(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCharToDouble function as declared in gdnative/string.h:137
func GodotStringCharToDouble(pWhat string) float64 {
	cpWhat, _ := unpackPCharString(pWhat)
	__ret := C.godot_string_char_to_double(cpWhat)
	__v := (float64)(__ret)
	return __v
}

// GodotStringCharToInt function as declared in gdnative/string.h:138
func GodotStringCharToInt(pWhat string) GodotInt {
	cpWhat, _ := unpackPCharString(pWhat)
	__ret := C.godot_string_char_to_int(cpWhat)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringWcharToInt function as declared in gdnative/string.h:139
func GodotStringWcharToInt(pStr []int32) int64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_wchar_to_int(cpStr)
	__v := (int64)(__ret)
	return __v
}

// GodotStringCharToIntWithLen function as declared in gdnative/string.h:140
func GodotStringCharToIntWithLen(pWhat string, pLen GodotInt) GodotInt {
	cpWhat, _ := unpackPCharString(pWhat)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_char_to_int_with_len(cpWhat, cpLen)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringCharToInt64WithLen function as declared in gdnative/string.h:141
func GodotStringCharToInt64WithLen(pStr []int32, pLen int32) int64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	cpLen, _ := (C.int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_char_to_int64_with_len(cpStr, cpLen)
	__v := (int64)(__ret)
	return __v
}

// GodotStringHexToInt64 function as declared in gdnative/string.h:142
func GodotStringHexToInt64(pSelf []GodotString) int64 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hex_to_int64(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// GodotStringHexToInt64WithPrefix function as declared in gdnative/string.h:143
func GodotStringHexToInt64WithPrefix(pSelf []GodotString) int64 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hex_to_int64_with_prefix(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// GodotStringToInt64 function as declared in gdnative/string.h:144
func GodotStringToInt64(pSelf []GodotString) int64 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_int64(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// GodotStringUnicodeCharToDouble function as declared in gdnative/string.h:145
func GodotStringUnicodeCharToDouble(pStr []int32, rEnd [][]int32) float64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	crEnd, _ := unpackArgSSInt32(rEnd)
	__ret := C.godot_string_unicode_char_to_double(cpStr, crEnd)
	packSSInt32(rEnd, crEnd)
	__v := (float64)(__ret)
	return __v
}

// GodotStringGetSliceCount function as declared in gdnative/string.h:147
func GodotStringGetSliceCount(pSelf []GodotString, pSplitter GodotString) GodotInt {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := pSplitter.PassValue()
	__ret := C.godot_string_get_slice_count(cpSelf, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotStringGetSlice function as declared in gdnative/string.h:148
func GodotStringGetSlice(pSelf []GodotString, pSplitter GodotString, pSlice GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := pSplitter.PassValue()
	cpSlice, _ := (C.godot_int)(pSlice), cgoAllocsUnknown
	__ret := C.godot_string_get_slice(cpSelf, cpSplitter, cpSlice)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringGetSlicec function as declared in gdnative/string.h:149
func GodotStringGetSlicec(pSelf []GodotString, pSplitter int32, pSlice GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := (C.wchar_t)(pSplitter), cgoAllocsUnknown
	cpSlice, _ := (C.godot_int)(pSlice), cgoAllocsUnknown
	__ret := C.godot_string_get_slicec(cpSelf, cpSplitter, cpSlice)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplit function as declared in gdnative/string.h:151
func GodotStringSplit(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitAllowEmpty function as declared in gdnative/string.h:152
func GodotStringSplitAllowEmpty(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split_allow_empty(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitFloats function as declared in gdnative/string.h:153
func GodotStringSplitFloats(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split_floats(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitFloatsAllowsEmpty function as declared in gdnative/string.h:154
func GodotStringSplitFloatsAllowsEmpty(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split_floats_allows_empty(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitFloatsMk function as declared in gdnative/string.h:155
func GodotStringSplitFloatsMk(pSelf []GodotString, pSplitters []GodotArray) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitters, _ := unpackArgSGodotArray(pSplitters)
	__ret := C.godot_string_split_floats_mk(cpSelf, cpSplitters)
	packSGodotArray(pSplitters, cpSplitters)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitFloatsMkAllowsEmpty function as declared in gdnative/string.h:156
func GodotStringSplitFloatsMkAllowsEmpty(pSelf []GodotString, pSplitters []GodotArray) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitters, _ := unpackArgSGodotArray(pSplitters)
	__ret := C.godot_string_split_floats_mk_allows_empty(cpSelf, cpSplitters)
	packSGodotArray(pSplitters, cpSplitters)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitInts function as declared in gdnative/string.h:157
func GodotStringSplitInts(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split_ints(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitIntsAllowsEmpty function as declared in gdnative/string.h:158
func GodotStringSplitIntsAllowsEmpty(pSelf []GodotString, pSplitter []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitter, _ := unpackArgSGodotString(pSplitter)
	__ret := C.godot_string_split_ints_allows_empty(cpSelf, cpSplitter)
	packSGodotString(pSplitter, cpSplitter)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitIntsMk function as declared in gdnative/string.h:159
func GodotStringSplitIntsMk(pSelf []GodotString, pSplitters []GodotArray) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitters, _ := unpackArgSGodotArray(pSplitters)
	__ret := C.godot_string_split_ints_mk(cpSelf, cpSplitters)
	packSGodotArray(pSplitters, cpSplitters)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitIntsMkAllowsEmpty function as declared in gdnative/string.h:160
func GodotStringSplitIntsMkAllowsEmpty(pSelf []GodotString, pSplitters []GodotArray) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpSplitters, _ := unpackArgSGodotArray(pSplitters)
	__ret := C.godot_string_split_ints_mk_allows_empty(cpSelf, cpSplitters)
	packSGodotArray(pSplitters, cpSplitters)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSplitSpaces function as declared in gdnative/string.h:161
func GodotStringSplitSpaces(pSelf []GodotString) GodotArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_split_spaces(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCharLowercase function as declared in gdnative/string.h:163
func GodotStringCharLowercase(pChar int32) int32 {
	cpChar, _ := (C.wchar_t)(pChar), cgoAllocsUnknown
	__ret := C.godot_string_char_lowercase(cpChar)
	__v := (int32)(__ret)
	return __v
}

// GodotStringCharUppercase function as declared in gdnative/string.h:164
func GodotStringCharUppercase(pChar int32) int32 {
	cpChar, _ := (C.wchar_t)(pChar), cgoAllocsUnknown
	__ret := C.godot_string_char_uppercase(cpChar)
	__v := (int32)(__ret)
	return __v
}

// GodotStringToLower function as declared in gdnative/string.h:165
func GodotStringToLower(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_lower(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringToUpper function as declared in gdnative/string.h:166
func GodotStringToUpper(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_to_upper(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringGetBasename function as declared in gdnative/string.h:168
func GodotStringGetBasename(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_get_basename(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringGetExtension function as declared in gdnative/string.h:169
func GodotStringGetExtension(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_get_extension(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringLeft function as declared in gdnative/string.h:170
func GodotStringLeft(pSelf []GodotString, pPos GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	__ret := C.godot_string_left(cpSelf, cpPos)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringOrdAt function as declared in gdnative/string.h:171
func GodotStringOrdAt(pSelf []GodotString, pIdx GodotInt) int32 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_ord_at(cpSelf, cpIdx)
	packSGodotString(pSelf, cpSelf)
	__v := (int32)(__ret)
	return __v
}

// GodotStringPlusFile function as declared in gdnative/string.h:172
func GodotStringPlusFile(pSelf []GodotString, pFile []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpFile, _ := unpackArgSGodotString(pFile)
	__ret := C.godot_string_plus_file(cpSelf, cpFile)
	packSGodotString(pFile, cpFile)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringRight function as declared in gdnative/string.h:173
func GodotStringRight(pSelf []GodotString, pPos GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	__ret := C.godot_string_right(cpSelf, cpPos)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringStripEdges function as declared in gdnative/string.h:174
func GodotStringStripEdges(pSelf []GodotString, pLeft GodotBool, pRight GodotBool) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpLeft, _ := (C.godot_bool)(pLeft), cgoAllocsUnknown
	cpRight, _ := (C.godot_bool)(pRight), cgoAllocsUnknown
	__ret := C.godot_string_strip_edges(cpSelf, cpLeft, cpRight)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringStripEscapes function as declared in gdnative/string.h:175
func GodotStringStripEscapes(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_strip_escapes(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringErase function as declared in gdnative/string.h:177
func GodotStringErase(pSelf []GodotString, pPos GodotInt, pChars GodotInt) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	cpChars, _ := (C.godot_int)(pChars), cgoAllocsUnknown
	C.godot_string_erase(cpSelf, cpPos, cpChars)
	packSGodotString(pSelf, cpSelf)
}

// GodotStringAscii function as declared in gdnative/string.h:179
func GodotStringAscii(pSelf []GodotString, result []byte) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_ascii(cpSelf, cresult)
	packSGodotString(pSelf, cpSelf)
}

// GodotStringAsciiExtended function as declared in gdnative/string.h:180
func GodotStringAsciiExtended(pSelf []GodotString, result []byte) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_ascii_extended(cpSelf, cresult)
	packSGodotString(pSelf, cpSelf)
}

// GodotStringUtf8 function as declared in gdnative/string.h:181
func GodotStringUtf8(pSelf []GodotString, result []byte) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_utf8(cpSelf, cresult)
	packSGodotString(pSelf, cpSelf)
}

// GodotStringParseUtf8 function as declared in gdnative/string.h:182
func GodotStringParseUtf8(pSelf []GodotString, pUtf8 string) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpUtf8, _ := unpackPCharString(pUtf8)
	__ret := C.godot_string_parse_utf8(cpSelf, cpUtf8)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringParseUtf8WithLen function as declared in gdnative/string.h:183
func GodotStringParseUtf8WithLen(pSelf []GodotString, pUtf8 string, pLen GodotInt) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpUtf8, _ := unpackPCharString(pUtf8)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_parse_utf8_with_len(cpSelf, cpUtf8, cpLen)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringCharsToUtf8 function as declared in gdnative/string.h:184
func GodotStringCharsToUtf8(pUtf8 string) GodotString {
	cpUtf8, _ := unpackPCharString(pUtf8)
	__ret := C.godot_string_chars_to_utf8(cpUtf8)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCharsToUtf8WithLen function as declared in gdnative/string.h:185
func GodotStringCharsToUtf8WithLen(pUtf8 string, pLen GodotInt) GodotString {
	cpUtf8, _ := unpackPCharString(pUtf8)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_chars_to_utf8_with_len(cpUtf8, cpLen)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHash function as declared in gdnative/string.h:187
func GodotStringHash(pSelf []GodotString) uint32 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hash(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringHash64 function as declared in gdnative/string.h:188
func GodotStringHash64(pSelf []GodotString) uint64 {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_hash64(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (uint64)(__ret)
	return __v
}

// GodotStringHashChars function as declared in gdnative/string.h:189
func GodotStringHashChars(pCstr string) uint32 {
	cpCstr, _ := unpackPCharString(pCstr)
	__ret := C.godot_string_hash_chars(cpCstr)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringHashCharsWithLen function as declared in gdnative/string.h:190
func GodotStringHashCharsWithLen(pCstr string, pLen GodotInt) uint32 {
	cpCstr, _ := unpackPCharString(pCstr)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hash_chars_with_len(cpCstr, cpLen)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringHashUtf8Chars function as declared in gdnative/string.h:191
func GodotStringHashUtf8Chars(pStr []int32) uint32 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_hash_utf8_chars(cpStr)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringHashUtf8CharsWithLen function as declared in gdnative/string.h:192
func GodotStringHashUtf8CharsWithLen(pStr []int32, pLen GodotInt) uint32 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hash_utf8_chars_with_len(cpStr, cpLen)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringMd5Buffer function as declared in gdnative/string.h:193
func GodotStringMd5Buffer(pSelf []GodotString) GodotPoolByteArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_md5_buffer(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringMd5Text function as declared in gdnative/string.h:194
func GodotStringMd5Text(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_md5_text(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSha256Buffer function as declared in gdnative/string.h:195
func GodotStringSha256Buffer(pSelf []GodotString) GodotPoolByteArray {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_sha256_buffer(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSha256Text function as declared in gdnative/string.h:196
func GodotStringSha256Text(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_sha256_text(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringEmpty function as declared in gdnative/string.h:198
func GodotStringEmpty(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_empty(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringGetBaseDir function as declared in gdnative/string.h:201
func GodotStringGetBaseDir(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_get_base_dir(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringGetFile function as declared in gdnative/string.h:202
func GodotStringGetFile(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_get_file(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHumanizeSize function as declared in gdnative/string.h:203
func GodotStringHumanizeSize(pSize uint) GodotString {
	cpSize, _ := (C.size_t)(pSize), cgoAllocsUnknown
	__ret := C.godot_string_humanize_size(cpSize)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringIsAbsPath function as declared in gdnative/string.h:204
func GodotStringIsAbsPath(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_abs_path(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsRelPath function as declared in gdnative/string.h:205
func GodotStringIsRelPath(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_rel_path(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsResourceFile function as declared in gdnative/string.h:206
func GodotStringIsResourceFile(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_resource_file(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringPathTo function as declared in gdnative/string.h:207
func GodotStringPathTo(pSelf []GodotString, pPath []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpPath, _ := unpackArgSGodotString(pPath)
	__ret := C.godot_string_path_to(cpSelf, cpPath)
	packSGodotString(pPath, cpPath)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringPathToFile function as declared in gdnative/string.h:208
func GodotStringPathToFile(pSelf []GodotString, pPath []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpPath, _ := unpackArgSGodotString(pPath)
	__ret := C.godot_string_path_to_file(cpSelf, cpPath)
	packSGodotString(pPath, cpPath)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringSimplifyPath function as declared in gdnative/string.h:209
func GodotStringSimplifyPath(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_simplify_path(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCEscape function as declared in gdnative/string.h:211
func GodotStringCEscape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_c_escape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCEscapeMultiline function as declared in gdnative/string.h:212
func GodotStringCEscapeMultiline(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_c_escape_multiline(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringCUnescape function as declared in gdnative/string.h:213
func GodotStringCUnescape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_c_unescape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHttpEscape function as declared in gdnative/string.h:214
func GodotStringHttpEscape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_http_escape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringHttpUnescape function as declared in gdnative/string.h:215
func GodotStringHttpUnescape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_http_unescape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringJsonEscape function as declared in gdnative/string.h:216
func GodotStringJsonEscape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_json_escape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringWordWrap function as declared in gdnative/string.h:217
func GodotStringWordWrap(pSelf []GodotString, pCharsPerLine GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpCharsPerLine, _ := (C.godot_int)(pCharsPerLine), cgoAllocsUnknown
	__ret := C.godot_string_word_wrap(cpSelf, cpCharsPerLine)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringXmlEscape function as declared in gdnative/string.h:218
func GodotStringXmlEscape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_xml_escape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringXmlEscapeWithQuotes function as declared in gdnative/string.h:219
func GodotStringXmlEscapeWithQuotes(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_xml_escape_with_quotes(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringXmlUnescape function as declared in gdnative/string.h:220
func GodotStringXmlUnescape(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_xml_unescape(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringPercentDecode function as declared in gdnative/string.h:222
func GodotStringPercentDecode(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_percent_decode(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringPercentEncode function as declared in gdnative/string.h:223
func GodotStringPercentEncode(pSelf []GodotString) GodotString {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_percent_encode(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringIsValidFloat function as declared in gdnative/string.h:225
func GodotStringIsValidFloat(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_valid_float(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsValidHexNumber function as declared in gdnative/string.h:226
func GodotStringIsValidHexNumber(pSelf []GodotString, pWithPrefix GodotBool) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	cpWithPrefix, _ := (C.godot_bool)(pWithPrefix), cgoAllocsUnknown
	__ret := C.godot_string_is_valid_hex_number(cpSelf, cpWithPrefix)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsValidHtmlColor function as declared in gdnative/string.h:227
func GodotStringIsValidHtmlColor(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_valid_html_color(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsValidIdentifier function as declared in gdnative/string.h:228
func GodotStringIsValidIdentifier(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_valid_identifier(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsValidInteger function as declared in gdnative/string.h:229
func GodotStringIsValidInteger(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_valid_integer(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringIsValidIpAddress function as declared in gdnative/string.h:230
func GodotStringIsValidIpAddress(pSelf []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	__ret := C.godot_string_is_valid_ip_address(cpSelf)
	packSGodotString(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringDestroy function as declared in gdnative/string.h:232
func GodotStringDestroy(pSelf []GodotString) {
	cpSelf, _ := unpackArgSGodotString(pSelf)
	C.godot_string_destroy(cpSelf)
	packSGodotString(pSelf, cpSelf)
}

// GodotArrayNew function as declared in gdnative/array.h:63
func GodotArrayNew(rDest []GodotArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	C.godot_array_new(crDest)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewCopy function as declared in gdnative/array.h:64
func GodotArrayNewCopy(rDest []GodotArray, pSrc []GodotArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpSrc, _ := unpackArgSGodotArray(pSrc)
	C.godot_array_new_copy(crDest, cpSrc)
	packSGodotArray(pSrc, cpSrc)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolColorArray function as declared in gdnative/array.h:65
func GodotArrayNewPoolColorArray(rDest []GodotArray, pPca []GodotPoolColorArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPca, _ := unpackArgSGodotPoolColorArray(pPca)
	C.godot_array_new_pool_color_array(crDest, cpPca)
	packSGodotPoolColorArray(pPca, cpPca)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolVector3Array function as declared in gdnative/array.h:66
func GodotArrayNewPoolVector3Array(rDest []GodotArray, pPv3a []GodotPoolVector3Array) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPv3a, _ := unpackArgSGodotPoolVector3Array(pPv3a)
	C.godot_array_new_pool_vector3_array(crDest, cpPv3a)
	packSGodotPoolVector3Array(pPv3a, cpPv3a)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolVector2Array function as declared in gdnative/array.h:67
func GodotArrayNewPoolVector2Array(rDest []GodotArray, pPv2a []GodotPoolVector2Array) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPv2a, _ := unpackArgSGodotPoolVector2Array(pPv2a)
	C.godot_array_new_pool_vector2_array(crDest, cpPv2a)
	packSGodotPoolVector2Array(pPv2a, cpPv2a)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolStringArray function as declared in gdnative/array.h:68
func GodotArrayNewPoolStringArray(rDest []GodotArray, pPsa []GodotPoolStringArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPsa, _ := unpackArgSGodotPoolStringArray(pPsa)
	C.godot_array_new_pool_string_array(crDest, cpPsa)
	packSGodotPoolStringArray(pPsa, cpPsa)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolRealArray function as declared in gdnative/array.h:69
func GodotArrayNewPoolRealArray(rDest []GodotArray, pPra []GodotPoolRealArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPra, _ := unpackArgSGodotPoolRealArray(pPra)
	C.godot_array_new_pool_real_array(crDest, cpPra)
	packSGodotPoolRealArray(pPra, cpPra)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolIntArray function as declared in gdnative/array.h:70
func GodotArrayNewPoolIntArray(rDest []GodotArray, pPia []GodotPoolIntArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPia, _ := unpackArgSGodotPoolIntArray(pPia)
	C.godot_array_new_pool_int_array(crDest, cpPia)
	packSGodotPoolIntArray(pPia, cpPia)
	packSGodotArray(rDest, crDest)
}

// GodotArrayNewPoolByteArray function as declared in gdnative/array.h:71
func GodotArrayNewPoolByteArray(rDest []GodotArray, pPba []GodotPoolByteArray) {
	crDest, _ := unpackArgSGodotArray(rDest)
	cpPba, _ := unpackArgSGodotPoolByteArray(pPba)
	C.godot_array_new_pool_byte_array(crDest, cpPba)
	packSGodotPoolByteArray(pPba, cpPba)
	packSGodotArray(rDest, crDest)
}

// GodotArraySet function as declared in gdnative/array.h:73
func GodotArraySet(pSelf []GodotArray, pIdx GodotInt, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_set(cpSelf, cpIdx, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayGet function as declared in gdnative/array.h:75
func GodotArrayGet(pSelf []GodotArray, pIdx GodotInt) GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_get(cpSelf, cpIdx)
	packSGodotArray(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArrayOperatorIndex function as declared in gdnative/array.h:77
func GodotArrayOperatorIndex(pSelf []GodotArray, pIdx GodotInt) *GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_operator_index(cpSelf, cpIdx)
	packSGodotArray(pSelf, cpSelf)
	__v := NewGodotVariantRef(unsafe.Pointer(__ret))
	return __v
}

// GodotArrayOperatorIndexConst function as declared in gdnative/array.h:79
func GodotArrayOperatorIndexConst(pSelf []GodotArray, pIdx GodotInt) *GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_operator_index_const(cpSelf, cpIdx)
	packSGodotArray(pSelf, cpSelf)
	__v := NewGodotVariantRef(unsafe.Pointer(__ret))
	return __v
}

// GodotArrayAppend function as declared in gdnative/array.h:81
func GodotArrayAppend(pSelf []GodotArray, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_append(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayClear function as declared in gdnative/array.h:83
func GodotArrayClear(pSelf []GodotArray) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	C.godot_array_clear(cpSelf)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayCount function as declared in gdnative/array.h:85
func GodotArrayCount(pSelf []GodotArray, pValue []GodotVariant) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	__ret := C.godot_array_count(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayEmpty function as declared in gdnative/array.h:87
func GodotArrayEmpty(pSelf []GodotArray) GodotBool {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_empty(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotArrayErase function as declared in gdnative/array.h:89
func GodotArrayErase(pSelf []GodotArray, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_erase(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayFront function as declared in gdnative/array.h:91
func GodotArrayFront(pSelf []GodotArray) GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_front(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArrayBack function as declared in gdnative/array.h:93
func GodotArrayBack(pSelf []GodotArray) GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_back(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArrayFind function as declared in gdnative/array.h:95
func GodotArrayFind(pSelf []GodotArray, pWhat []GodotVariant, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpWhat, _ := unpackArgSGodotVariant(pWhat)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_array_find(cpSelf, cpWhat, cpFrom)
	packSGodotVariant(pWhat, cpWhat)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayFindLast function as declared in gdnative/array.h:97
func GodotArrayFindLast(pSelf []GodotArray, pWhat []GodotVariant) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpWhat, _ := unpackArgSGodotVariant(pWhat)
	__ret := C.godot_array_find_last(cpSelf, cpWhat)
	packSGodotVariant(pWhat, cpWhat)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayHas function as declared in gdnative/array.h:99
func GodotArrayHas(pSelf []GodotArray, pValue []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	__ret := C.godot_array_has(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotArrayHash function as declared in gdnative/array.h:101
func GodotArrayHash(pSelf []GodotArray) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_hash(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayInsert function as declared in gdnative/array.h:103
func GodotArrayInsert(pSelf []GodotArray, pPos GodotInt, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_insert(cpSelf, cpPos, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayInvert function as declared in gdnative/array.h:105
func GodotArrayInvert(pSelf []GodotArray) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	C.godot_array_invert(cpSelf)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayPopBack function as declared in gdnative/array.h:107
func GodotArrayPopBack(pSelf []GodotArray) GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_pop_back(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArrayPopFront function as declared in gdnative/array.h:109
func GodotArrayPopFront(pSelf []GodotArray) GodotVariant {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_pop_front(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArrayPushBack function as declared in gdnative/array.h:111
func GodotArrayPushBack(pSelf []GodotArray, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_push_back(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayPushFront function as declared in gdnative/array.h:113
func GodotArrayPushFront(pSelf []GodotArray, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_array_push_front(cpSelf, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayRemove function as declared in gdnative/array.h:115
func GodotArrayRemove(pSelf []GodotArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_array_remove(cpSelf, cpIdx)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayResize function as declared in gdnative/array.h:117
func GodotArrayResize(pSelf []GodotArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_array_resize(cpSelf, cpSize)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayRfind function as declared in gdnative/array.h:119
func GodotArrayRfind(pSelf []GodotArray, pWhat []GodotVariant, pFrom GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpWhat, _ := unpackArgSGodotVariant(pWhat)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_array_rfind(cpSelf, cpWhat, cpFrom)
	packSGodotVariant(pWhat, cpWhat)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArraySize function as declared in gdnative/array.h:121
func GodotArraySize(pSelf []GodotArray) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	__ret := C.godot_array_size(cpSelf)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArraySort function as declared in gdnative/array.h:123
func GodotArraySort(pSelf []GodotArray) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	C.godot_array_sort(cpSelf)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArraySortCustom function as declared in gdnative/array.h:125
func GodotArraySortCustom(pSelf []GodotArray, pObj *GodotObject, pFunc []GodotString) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	cpFunc, _ := unpackArgSGodotString(pFunc)
	C.godot_array_sort_custom(cpSelf, cpObj, cpFunc)
	packSGodotString(pFunc, cpFunc)
	packSGodotArray(pSelf, cpSelf)
}

// GodotArrayBsearch function as declared in gdnative/array.h:127
func GodotArrayBsearch(pSelf []GodotArray, pValue []GodotVariant, pBefore GodotBool) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	cpBefore, _ := (C.godot_bool)(pBefore), cgoAllocsUnknown
	__ret := C.godot_array_bsearch(cpSelf, cpValue, cpBefore)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayBsearchCustom function as declared in gdnative/array.h:129
func GodotArrayBsearchCustom(pSelf []GodotArray, pValue []GodotVariant, pObj *GodotObject, pFunc []GodotString, pBefore GodotBool) GodotInt {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	cpFunc, _ := unpackArgSGodotString(pFunc)
	cpBefore, _ := (C.godot_bool)(pBefore), cgoAllocsUnknown
	__ret := C.godot_array_bsearch_custom(cpSelf, cpValue, cpObj, cpFunc, cpBefore)
	packSGodotString(pFunc, cpFunc)
	packSGodotVariant(pValue, cpValue)
	packSGodotArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArrayDestroy function as declared in gdnative/array.h:131
func GodotArrayDestroy(pSelf []GodotArray) {
	cpSelf, _ := unpackArgSGodotArray(pSelf)
	C.godot_array_destroy(cpSelf)
	packSGodotArray(pSelf, cpSelf)
}

// GodotPoolByteArrayNew function as declared in gdnative/pool_arrays.h:166
func GodotPoolByteArrayNew(rDest []GodotPoolByteArray) {
	crDest, _ := unpackArgSGodotPoolByteArray(rDest)
	C.godot_pool_byte_array_new(crDest)
	packSGodotPoolByteArray(rDest, crDest)
}

// GodotPoolByteArrayNewCopy function as declared in gdnative/pool_arrays.h:167
func GodotPoolByteArrayNewCopy(rDest []GodotPoolByteArray, pSrc []GodotPoolByteArray) {
	crDest, _ := unpackArgSGodotPoolByteArray(rDest)
	cpSrc, _ := unpackArgSGodotPoolByteArray(pSrc)
	C.godot_pool_byte_array_new_copy(crDest, cpSrc)
	packSGodotPoolByteArray(pSrc, cpSrc)
	packSGodotPoolByteArray(rDest, crDest)
}

// GodotPoolByteArrayNewWithArray function as declared in gdnative/pool_arrays.h:168
func GodotPoolByteArrayNewWithArray(rDest []GodotPoolByteArray, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolByteArray(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_byte_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolByteArray(rDest, crDest)
}

// GodotPoolByteArrayAppend function as declared in gdnative/pool_arrays.h:170
func GodotPoolByteArrayAppend(pSelf []GodotPoolByteArray, pData byte) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_append(cpSelf, cpData)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayAppendArray function as declared in gdnative/pool_arrays.h:172
func GodotPoolByteArrayAppendArray(pSelf []GodotPoolByteArray, pArray []GodotPoolByteArray) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpArray, _ := unpackArgSGodotPoolByteArray(pArray)
	C.godot_pool_byte_array_append_array(cpSelf, cpArray)
	packSGodotPoolByteArray(pArray, cpArray)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayInsert function as declared in gdnative/pool_arrays.h:174
func GodotPoolByteArrayInsert(pSelf []GodotPoolByteArray, pIdx GodotInt, pData byte) GodotError {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_byte_array_insert(cpSelf, cpIdx, cpData)
	packSGodotPoolByteArray(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolByteArrayInvert function as declared in gdnative/pool_arrays.h:176
func GodotPoolByteArrayInvert(pSelf []GodotPoolByteArray) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	C.godot_pool_byte_array_invert(cpSelf)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayPushBack function as declared in gdnative/pool_arrays.h:178
func GodotPoolByteArrayPushBack(pSelf []GodotPoolByteArray, pData byte) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_push_back(cpSelf, cpData)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayRemove function as declared in gdnative/pool_arrays.h:180
func GodotPoolByteArrayRemove(pSelf []GodotPoolByteArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_byte_array_remove(cpSelf, cpIdx)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayResize function as declared in gdnative/pool_arrays.h:182
func GodotPoolByteArrayResize(pSelf []GodotPoolByteArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_byte_array_resize(cpSelf, cpSize)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayRead function as declared in gdnative/pool_arrays.h:184
func GodotPoolByteArrayRead(pSelf []GodotPoolByteArray) *GodotPoolByteArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_read(cpSelf)
	packSGodotPoolByteArray(pSelf, cpSelf)
	__v := NewGodotPoolByteArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolByteArrayWrite function as declared in gdnative/pool_arrays.h:186
func GodotPoolByteArrayWrite(pSelf []GodotPoolByteArray) *GodotPoolByteArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_write(cpSelf)
	packSGodotPoolByteArray(pSelf, cpSelf)
	__v := NewGodotPoolByteArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolByteArraySet function as declared in gdnative/pool_arrays.h:188
func GodotPoolByteArraySet(pSelf []GodotPoolByteArray, pIdx GodotInt, pData byte) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_set(cpSelf, cpIdx, cpData)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolByteArrayGet function as declared in gdnative/pool_arrays.h:189
func GodotPoolByteArrayGet(pSelf []GodotPoolByteArray, pIdx GodotInt) byte {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_byte_array_get(cpSelf, cpIdx)
	packSGodotPoolByteArray(pSelf, cpSelf)
	__v := (byte)(__ret)
	return __v
}

// GodotPoolByteArraySize function as declared in gdnative/pool_arrays.h:191
func GodotPoolByteArraySize(pSelf []GodotPoolByteArray) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_size(cpSelf)
	packSGodotPoolByteArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolByteArrayDestroy function as declared in gdnative/pool_arrays.h:193
func GodotPoolByteArrayDestroy(pSelf []GodotPoolByteArray) {
	cpSelf, _ := unpackArgSGodotPoolByteArray(pSelf)
	C.godot_pool_byte_array_destroy(cpSelf)
	packSGodotPoolByteArray(pSelf, cpSelf)
}

// GodotPoolIntArrayNew function as declared in gdnative/pool_arrays.h:197
func GodotPoolIntArrayNew(rDest []GodotPoolIntArray) {
	crDest, _ := unpackArgSGodotPoolIntArray(rDest)
	C.godot_pool_int_array_new(crDest)
	packSGodotPoolIntArray(rDest, crDest)
}

// GodotPoolIntArrayNewCopy function as declared in gdnative/pool_arrays.h:198
func GodotPoolIntArrayNewCopy(rDest []GodotPoolIntArray, pSrc []GodotPoolIntArray) {
	crDest, _ := unpackArgSGodotPoolIntArray(rDest)
	cpSrc, _ := unpackArgSGodotPoolIntArray(pSrc)
	C.godot_pool_int_array_new_copy(crDest, cpSrc)
	packSGodotPoolIntArray(pSrc, cpSrc)
	packSGodotPoolIntArray(rDest, crDest)
}

// GodotPoolIntArrayNewWithArray function as declared in gdnative/pool_arrays.h:199
func GodotPoolIntArrayNewWithArray(rDest []GodotPoolIntArray, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolIntArray(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_int_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolIntArray(rDest, crDest)
}

// GodotPoolIntArrayAppend function as declared in gdnative/pool_arrays.h:201
func GodotPoolIntArrayAppend(pSelf []GodotPoolIntArray, pData GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_append(cpSelf, cpData)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayAppendArray function as declared in gdnative/pool_arrays.h:203
func GodotPoolIntArrayAppendArray(pSelf []GodotPoolIntArray, pArray []GodotPoolIntArray) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpArray, _ := unpackArgSGodotPoolIntArray(pArray)
	C.godot_pool_int_array_append_array(cpSelf, cpArray)
	packSGodotPoolIntArray(pArray, cpArray)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayInsert function as declared in gdnative/pool_arrays.h:205
func GodotPoolIntArrayInsert(pSelf []GodotPoolIntArray, pIdx GodotInt, pData GodotInt) GodotError {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_int_array_insert(cpSelf, cpIdx, cpData)
	packSGodotPoolIntArray(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolIntArrayInvert function as declared in gdnative/pool_arrays.h:207
func GodotPoolIntArrayInvert(pSelf []GodotPoolIntArray) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	C.godot_pool_int_array_invert(cpSelf)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayPushBack function as declared in gdnative/pool_arrays.h:209
func GodotPoolIntArrayPushBack(pSelf []GodotPoolIntArray, pData GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_push_back(cpSelf, cpData)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayRemove function as declared in gdnative/pool_arrays.h:211
func GodotPoolIntArrayRemove(pSelf []GodotPoolIntArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_int_array_remove(cpSelf, cpIdx)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayResize function as declared in gdnative/pool_arrays.h:213
func GodotPoolIntArrayResize(pSelf []GodotPoolIntArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_int_array_resize(cpSelf, cpSize)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayRead function as declared in gdnative/pool_arrays.h:215
func GodotPoolIntArrayRead(pSelf []GodotPoolIntArray) *GodotPoolIntArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_read(cpSelf)
	packSGodotPoolIntArray(pSelf, cpSelf)
	__v := NewGodotPoolIntArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolIntArrayWrite function as declared in gdnative/pool_arrays.h:217
func GodotPoolIntArrayWrite(pSelf []GodotPoolIntArray) *GodotPoolIntArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_write(cpSelf)
	packSGodotPoolIntArray(pSelf, cpSelf)
	__v := NewGodotPoolIntArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolIntArraySet function as declared in gdnative/pool_arrays.h:219
func GodotPoolIntArraySet(pSelf []GodotPoolIntArray, pIdx GodotInt, pData GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_set(cpSelf, cpIdx, cpData)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolIntArrayGet function as declared in gdnative/pool_arrays.h:220
func GodotPoolIntArrayGet(pSelf []GodotPoolIntArray, pIdx GodotInt) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_int_array_get(cpSelf, cpIdx)
	packSGodotPoolIntArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolIntArraySize function as declared in gdnative/pool_arrays.h:222
func GodotPoolIntArraySize(pSelf []GodotPoolIntArray) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_size(cpSelf)
	packSGodotPoolIntArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolIntArrayDestroy function as declared in gdnative/pool_arrays.h:224
func GodotPoolIntArrayDestroy(pSelf []GodotPoolIntArray) {
	cpSelf, _ := unpackArgSGodotPoolIntArray(pSelf)
	C.godot_pool_int_array_destroy(cpSelf)
	packSGodotPoolIntArray(pSelf, cpSelf)
}

// GodotPoolRealArrayNew function as declared in gdnative/pool_arrays.h:228
func GodotPoolRealArrayNew(rDest []GodotPoolRealArray) {
	crDest, _ := unpackArgSGodotPoolRealArray(rDest)
	C.godot_pool_real_array_new(crDest)
	packSGodotPoolRealArray(rDest, crDest)
}

// GodotPoolRealArrayNewCopy function as declared in gdnative/pool_arrays.h:229
func GodotPoolRealArrayNewCopy(rDest []GodotPoolRealArray, pSrc []GodotPoolRealArray) {
	crDest, _ := unpackArgSGodotPoolRealArray(rDest)
	cpSrc, _ := unpackArgSGodotPoolRealArray(pSrc)
	C.godot_pool_real_array_new_copy(crDest, cpSrc)
	packSGodotPoolRealArray(pSrc, cpSrc)
	packSGodotPoolRealArray(rDest, crDest)
}

// GodotPoolRealArrayNewWithArray function as declared in gdnative/pool_arrays.h:230
func GodotPoolRealArrayNewWithArray(rDest []GodotPoolRealArray, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolRealArray(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_real_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolRealArray(rDest, crDest)
}

// GodotPoolRealArrayAppend function as declared in gdnative/pool_arrays.h:232
func GodotPoolRealArrayAppend(pSelf []GodotPoolRealArray, pData GodotReal) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_append(cpSelf, cpData)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayAppendArray function as declared in gdnative/pool_arrays.h:234
func GodotPoolRealArrayAppendArray(pSelf []GodotPoolRealArray, pArray []GodotPoolRealArray) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpArray, _ := unpackArgSGodotPoolRealArray(pArray)
	C.godot_pool_real_array_append_array(cpSelf, cpArray)
	packSGodotPoolRealArray(pArray, cpArray)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayInsert function as declared in gdnative/pool_arrays.h:236
func GodotPoolRealArrayInsert(pSelf []GodotPoolRealArray, pIdx GodotInt, pData GodotReal) GodotError {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_real_array_insert(cpSelf, cpIdx, cpData)
	packSGodotPoolRealArray(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolRealArrayInvert function as declared in gdnative/pool_arrays.h:238
func GodotPoolRealArrayInvert(pSelf []GodotPoolRealArray) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	C.godot_pool_real_array_invert(cpSelf)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayPushBack function as declared in gdnative/pool_arrays.h:240
func GodotPoolRealArrayPushBack(pSelf []GodotPoolRealArray, pData GodotReal) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_push_back(cpSelf, cpData)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayRemove function as declared in gdnative/pool_arrays.h:242
func GodotPoolRealArrayRemove(pSelf []GodotPoolRealArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_real_array_remove(cpSelf, cpIdx)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayResize function as declared in gdnative/pool_arrays.h:244
func GodotPoolRealArrayResize(pSelf []GodotPoolRealArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_real_array_resize(cpSelf, cpSize)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayRead function as declared in gdnative/pool_arrays.h:246
func GodotPoolRealArrayRead(pSelf []GodotPoolRealArray) *GodotPoolRealArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_read(cpSelf)
	packSGodotPoolRealArray(pSelf, cpSelf)
	__v := NewGodotPoolRealArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolRealArrayWrite function as declared in gdnative/pool_arrays.h:248
func GodotPoolRealArrayWrite(pSelf []GodotPoolRealArray) *GodotPoolRealArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_write(cpSelf)
	packSGodotPoolRealArray(pSelf, cpSelf)
	__v := NewGodotPoolRealArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolRealArraySet function as declared in gdnative/pool_arrays.h:250
func GodotPoolRealArraySet(pSelf []GodotPoolRealArray, pIdx GodotInt, pData GodotReal) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_set(cpSelf, cpIdx, cpData)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolRealArrayGet function as declared in gdnative/pool_arrays.h:251
func GodotPoolRealArrayGet(pSelf []GodotPoolRealArray, pIdx GodotInt) GodotReal {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_real_array_get(cpSelf, cpIdx)
	packSGodotPoolRealArray(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotPoolRealArraySize function as declared in gdnative/pool_arrays.h:253
func GodotPoolRealArraySize(pSelf []GodotPoolRealArray) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_size(cpSelf)
	packSGodotPoolRealArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolRealArrayDestroy function as declared in gdnative/pool_arrays.h:255
func GodotPoolRealArrayDestroy(pSelf []GodotPoolRealArray) {
	cpSelf, _ := unpackArgSGodotPoolRealArray(pSelf)
	C.godot_pool_real_array_destroy(cpSelf)
	packSGodotPoolRealArray(pSelf, cpSelf)
}

// GodotPoolStringArrayNew function as declared in gdnative/pool_arrays.h:259
func GodotPoolStringArrayNew(rDest []GodotPoolStringArray) {
	crDest, _ := unpackArgSGodotPoolStringArray(rDest)
	C.godot_pool_string_array_new(crDest)
	packSGodotPoolStringArray(rDest, crDest)
}

// GodotPoolStringArrayNewCopy function as declared in gdnative/pool_arrays.h:260
func GodotPoolStringArrayNewCopy(rDest []GodotPoolStringArray, pSrc []GodotPoolStringArray) {
	crDest, _ := unpackArgSGodotPoolStringArray(rDest)
	cpSrc, _ := unpackArgSGodotPoolStringArray(pSrc)
	C.godot_pool_string_array_new_copy(crDest, cpSrc)
	packSGodotPoolStringArray(pSrc, cpSrc)
	packSGodotPoolStringArray(rDest, crDest)
}

// GodotPoolStringArrayNewWithArray function as declared in gdnative/pool_arrays.h:261
func GodotPoolStringArrayNewWithArray(rDest []GodotPoolStringArray, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolStringArray(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_string_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolStringArray(rDest, crDest)
}

// GodotPoolStringArrayAppend function as declared in gdnative/pool_arrays.h:263
func GodotPoolStringArrayAppend(pSelf []GodotPoolStringArray, pData []GodotString) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpData, _ := unpackArgSGodotString(pData)
	C.godot_pool_string_array_append(cpSelf, cpData)
	packSGodotString(pData, cpData)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayAppendArray function as declared in gdnative/pool_arrays.h:265
func GodotPoolStringArrayAppendArray(pSelf []GodotPoolStringArray, pArray []GodotPoolStringArray) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpArray, _ := unpackArgSGodotPoolStringArray(pArray)
	C.godot_pool_string_array_append_array(cpSelf, cpArray)
	packSGodotPoolStringArray(pArray, cpArray)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayInsert function as declared in gdnative/pool_arrays.h:267
func GodotPoolStringArrayInsert(pSelf []GodotPoolStringArray, pIdx GodotInt, pData []GodotString) GodotError {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotString(pData)
	__ret := C.godot_pool_string_array_insert(cpSelf, cpIdx, cpData)
	packSGodotString(pData, cpData)
	packSGodotPoolStringArray(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolStringArrayInvert function as declared in gdnative/pool_arrays.h:269
func GodotPoolStringArrayInvert(pSelf []GodotPoolStringArray) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	C.godot_pool_string_array_invert(cpSelf)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayPushBack function as declared in gdnative/pool_arrays.h:271
func GodotPoolStringArrayPushBack(pSelf []GodotPoolStringArray, pData []GodotString) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpData, _ := unpackArgSGodotString(pData)
	C.godot_pool_string_array_push_back(cpSelf, cpData)
	packSGodotString(pData, cpData)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayRemove function as declared in gdnative/pool_arrays.h:273
func GodotPoolStringArrayRemove(pSelf []GodotPoolStringArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_string_array_remove(cpSelf, cpIdx)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayResize function as declared in gdnative/pool_arrays.h:275
func GodotPoolStringArrayResize(pSelf []GodotPoolStringArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_string_array_resize(cpSelf, cpSize)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayRead function as declared in gdnative/pool_arrays.h:277
func GodotPoolStringArrayRead(pSelf []GodotPoolStringArray) *GodotPoolStringArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_read(cpSelf)
	packSGodotPoolStringArray(pSelf, cpSelf)
	__v := NewGodotPoolStringArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolStringArrayWrite function as declared in gdnative/pool_arrays.h:279
func GodotPoolStringArrayWrite(pSelf []GodotPoolStringArray) *GodotPoolStringArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_write(cpSelf)
	packSGodotPoolStringArray(pSelf, cpSelf)
	__v := NewGodotPoolStringArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolStringArraySet function as declared in gdnative/pool_arrays.h:281
func GodotPoolStringArraySet(pSelf []GodotPoolStringArray, pIdx GodotInt, pData []GodotString) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotString(pData)
	C.godot_pool_string_array_set(cpSelf, cpIdx, cpData)
	packSGodotString(pData, cpData)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolStringArrayGet function as declared in gdnative/pool_arrays.h:282
func GodotPoolStringArrayGet(pSelf []GodotPoolStringArray, pIdx GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_string_array_get(cpSelf, cpIdx)
	packSGodotPoolStringArray(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolStringArraySize function as declared in gdnative/pool_arrays.h:284
func GodotPoolStringArraySize(pSelf []GodotPoolStringArray) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_size(cpSelf)
	packSGodotPoolStringArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolStringArrayDestroy function as declared in gdnative/pool_arrays.h:286
func GodotPoolStringArrayDestroy(pSelf []GodotPoolStringArray) {
	cpSelf, _ := unpackArgSGodotPoolStringArray(pSelf)
	C.godot_pool_string_array_destroy(cpSelf)
	packSGodotPoolStringArray(pSelf, cpSelf)
}

// GodotPoolVector2ArrayNew function as declared in gdnative/pool_arrays.h:290
func GodotPoolVector2ArrayNew(rDest []GodotPoolVector2Array) {
	crDest, _ := unpackArgSGodotPoolVector2Array(rDest)
	C.godot_pool_vector2_array_new(crDest)
	packSGodotPoolVector2Array(rDest, crDest)
}

// GodotPoolVector2ArrayNewCopy function as declared in gdnative/pool_arrays.h:291
func GodotPoolVector2ArrayNewCopy(rDest []GodotPoolVector2Array, pSrc []GodotPoolVector2Array) {
	crDest, _ := unpackArgSGodotPoolVector2Array(rDest)
	cpSrc, _ := unpackArgSGodotPoolVector2Array(pSrc)
	C.godot_pool_vector2_array_new_copy(crDest, cpSrc)
	packSGodotPoolVector2Array(pSrc, cpSrc)
	packSGodotPoolVector2Array(rDest, crDest)
}

// GodotPoolVector2ArrayNewWithArray function as declared in gdnative/pool_arrays.h:292
func GodotPoolVector2ArrayNewWithArray(rDest []GodotPoolVector2Array, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolVector2Array(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_vector2_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolVector2Array(rDest, crDest)
}

// GodotPoolVector2ArrayAppend function as declared in gdnative/pool_arrays.h:294
func GodotPoolVector2ArrayAppend(pSelf []GodotPoolVector2Array, pData []GodotVector2) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpData, _ := unpackArgSGodotVector2(pData)
	C.godot_pool_vector2_array_append(cpSelf, cpData)
	packSGodotVector2(pData, cpData)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayAppendArray function as declared in gdnative/pool_arrays.h:296
func GodotPoolVector2ArrayAppendArray(pSelf []GodotPoolVector2Array, pArray []GodotPoolVector2Array) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpArray, _ := unpackArgSGodotPoolVector2Array(pArray)
	C.godot_pool_vector2_array_append_array(cpSelf, cpArray)
	packSGodotPoolVector2Array(pArray, cpArray)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayInsert function as declared in gdnative/pool_arrays.h:298
func GodotPoolVector2ArrayInsert(pSelf []GodotPoolVector2Array, pIdx GodotInt, pData []GodotVector2) GodotError {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotVector2(pData)
	__ret := C.godot_pool_vector2_array_insert(cpSelf, cpIdx, cpData)
	packSGodotVector2(pData, cpData)
	packSGodotPoolVector2Array(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolVector2ArrayInvert function as declared in gdnative/pool_arrays.h:300
func GodotPoolVector2ArrayInvert(pSelf []GodotPoolVector2Array) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	C.godot_pool_vector2_array_invert(cpSelf)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayPushBack function as declared in gdnative/pool_arrays.h:302
func GodotPoolVector2ArrayPushBack(pSelf []GodotPoolVector2Array, pData []GodotVector2) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpData, _ := unpackArgSGodotVector2(pData)
	C.godot_pool_vector2_array_push_back(cpSelf, cpData)
	packSGodotVector2(pData, cpData)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayRemove function as declared in gdnative/pool_arrays.h:304
func GodotPoolVector2ArrayRemove(pSelf []GodotPoolVector2Array, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_vector2_array_remove(cpSelf, cpIdx)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayResize function as declared in gdnative/pool_arrays.h:306
func GodotPoolVector2ArrayResize(pSelf []GodotPoolVector2Array, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_vector2_array_resize(cpSelf, cpSize)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayRead function as declared in gdnative/pool_arrays.h:308
func GodotPoolVector2ArrayRead(pSelf []GodotPoolVector2Array) *GodotPoolVector2ArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_read(cpSelf)
	packSGodotPoolVector2Array(pSelf, cpSelf)
	__v := NewGodotPoolVector2ArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector2ArrayWrite function as declared in gdnative/pool_arrays.h:310
func GodotPoolVector2ArrayWrite(pSelf []GodotPoolVector2Array) *GodotPoolVector2ArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_write(cpSelf)
	packSGodotPoolVector2Array(pSelf, cpSelf)
	__v := NewGodotPoolVector2ArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector2ArraySet function as declared in gdnative/pool_arrays.h:312
func GodotPoolVector2ArraySet(pSelf []GodotPoolVector2Array, pIdx GodotInt, pData []GodotVector2) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotVector2(pData)
	C.godot_pool_vector2_array_set(cpSelf, cpIdx, cpData)
	packSGodotVector2(pData, cpData)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector2ArrayGet function as declared in gdnative/pool_arrays.h:313
func GodotPoolVector2ArrayGet(pSelf []GodotPoolVector2Array, pIdx GodotInt) GodotVector2 {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_vector2_array_get(cpSelf, cpIdx)
	packSGodotPoolVector2Array(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolVector2ArraySize function as declared in gdnative/pool_arrays.h:315
func GodotPoolVector2ArraySize(pSelf []GodotPoolVector2Array) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_size(cpSelf)
	packSGodotPoolVector2Array(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolVector2ArrayDestroy function as declared in gdnative/pool_arrays.h:317
func GodotPoolVector2ArrayDestroy(pSelf []GodotPoolVector2Array) {
	cpSelf, _ := unpackArgSGodotPoolVector2Array(pSelf)
	C.godot_pool_vector2_array_destroy(cpSelf)
	packSGodotPoolVector2Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayNew function as declared in gdnative/pool_arrays.h:321
func GodotPoolVector3ArrayNew(rDest []GodotPoolVector3Array) {
	crDest, _ := unpackArgSGodotPoolVector3Array(rDest)
	C.godot_pool_vector3_array_new(crDest)
	packSGodotPoolVector3Array(rDest, crDest)
}

// GodotPoolVector3ArrayNewCopy function as declared in gdnative/pool_arrays.h:322
func GodotPoolVector3ArrayNewCopy(rDest []GodotPoolVector3Array, pSrc []GodotPoolVector3Array) {
	crDest, _ := unpackArgSGodotPoolVector3Array(rDest)
	cpSrc, _ := unpackArgSGodotPoolVector3Array(pSrc)
	C.godot_pool_vector3_array_new_copy(crDest, cpSrc)
	packSGodotPoolVector3Array(pSrc, cpSrc)
	packSGodotPoolVector3Array(rDest, crDest)
}

// GodotPoolVector3ArrayNewWithArray function as declared in gdnative/pool_arrays.h:323
func GodotPoolVector3ArrayNewWithArray(rDest []GodotPoolVector3Array, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolVector3Array(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_vector3_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolVector3Array(rDest, crDest)
}

// GodotPoolVector3ArrayAppend function as declared in gdnative/pool_arrays.h:325
func GodotPoolVector3ArrayAppend(pSelf []GodotPoolVector3Array, pData []GodotVector3) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpData, _ := unpackArgSGodotVector3(pData)
	C.godot_pool_vector3_array_append(cpSelf, cpData)
	packSGodotVector3(pData, cpData)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayAppendArray function as declared in gdnative/pool_arrays.h:327
func GodotPoolVector3ArrayAppendArray(pSelf []GodotPoolVector3Array, pArray []GodotPoolVector3Array) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpArray, _ := unpackArgSGodotPoolVector3Array(pArray)
	C.godot_pool_vector3_array_append_array(cpSelf, cpArray)
	packSGodotPoolVector3Array(pArray, cpArray)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayInsert function as declared in gdnative/pool_arrays.h:329
func GodotPoolVector3ArrayInsert(pSelf []GodotPoolVector3Array, pIdx GodotInt, pData []GodotVector3) GodotError {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotVector3(pData)
	__ret := C.godot_pool_vector3_array_insert(cpSelf, cpIdx, cpData)
	packSGodotVector3(pData, cpData)
	packSGodotPoolVector3Array(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolVector3ArrayInvert function as declared in gdnative/pool_arrays.h:331
func GodotPoolVector3ArrayInvert(pSelf []GodotPoolVector3Array) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	C.godot_pool_vector3_array_invert(cpSelf)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayPushBack function as declared in gdnative/pool_arrays.h:333
func GodotPoolVector3ArrayPushBack(pSelf []GodotPoolVector3Array, pData []GodotVector3) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpData, _ := unpackArgSGodotVector3(pData)
	C.godot_pool_vector3_array_push_back(cpSelf, cpData)
	packSGodotVector3(pData, cpData)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayRemove function as declared in gdnative/pool_arrays.h:335
func GodotPoolVector3ArrayRemove(pSelf []GodotPoolVector3Array, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_vector3_array_remove(cpSelf, cpIdx)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayResize function as declared in gdnative/pool_arrays.h:337
func GodotPoolVector3ArrayResize(pSelf []GodotPoolVector3Array, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_vector3_array_resize(cpSelf, cpSize)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayRead function as declared in gdnative/pool_arrays.h:339
func GodotPoolVector3ArrayRead(pSelf []GodotPoolVector3Array) *GodotPoolVector3ArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_read(cpSelf)
	packSGodotPoolVector3Array(pSelf, cpSelf)
	__v := NewGodotPoolVector3ArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector3ArrayWrite function as declared in gdnative/pool_arrays.h:341
func GodotPoolVector3ArrayWrite(pSelf []GodotPoolVector3Array) *GodotPoolVector3ArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_write(cpSelf)
	packSGodotPoolVector3Array(pSelf, cpSelf)
	__v := NewGodotPoolVector3ArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector3ArraySet function as declared in gdnative/pool_arrays.h:343
func GodotPoolVector3ArraySet(pSelf []GodotPoolVector3Array, pIdx GodotInt, pData []GodotVector3) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotVector3(pData)
	C.godot_pool_vector3_array_set(cpSelf, cpIdx, cpData)
	packSGodotVector3(pData, cpData)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolVector3ArrayGet function as declared in gdnative/pool_arrays.h:344
func GodotPoolVector3ArrayGet(pSelf []GodotPoolVector3Array, pIdx GodotInt) GodotVector3 {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_vector3_array_get(cpSelf, cpIdx)
	packSGodotPoolVector3Array(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolVector3ArraySize function as declared in gdnative/pool_arrays.h:346
func GodotPoolVector3ArraySize(pSelf []GodotPoolVector3Array) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_size(cpSelf)
	packSGodotPoolVector3Array(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolVector3ArrayDestroy function as declared in gdnative/pool_arrays.h:348
func GodotPoolVector3ArrayDestroy(pSelf []GodotPoolVector3Array) {
	cpSelf, _ := unpackArgSGodotPoolVector3Array(pSelf)
	C.godot_pool_vector3_array_destroy(cpSelf)
	packSGodotPoolVector3Array(pSelf, cpSelf)
}

// GodotPoolColorArrayNew function as declared in gdnative/pool_arrays.h:352
func GodotPoolColorArrayNew(rDest []GodotPoolColorArray) {
	crDest, _ := unpackArgSGodotPoolColorArray(rDest)
	C.godot_pool_color_array_new(crDest)
	packSGodotPoolColorArray(rDest, crDest)
}

// GodotPoolColorArrayNewCopy function as declared in gdnative/pool_arrays.h:353
func GodotPoolColorArrayNewCopy(rDest []GodotPoolColorArray, pSrc []GodotPoolColorArray) {
	crDest, _ := unpackArgSGodotPoolColorArray(rDest)
	cpSrc, _ := unpackArgSGodotPoolColorArray(pSrc)
	C.godot_pool_color_array_new_copy(crDest, cpSrc)
	packSGodotPoolColorArray(pSrc, cpSrc)
	packSGodotPoolColorArray(rDest, crDest)
}

// GodotPoolColorArrayNewWithArray function as declared in gdnative/pool_arrays.h:354
func GodotPoolColorArrayNewWithArray(rDest []GodotPoolColorArray, pA []GodotArray) {
	crDest, _ := unpackArgSGodotPoolColorArray(rDest)
	cpA, _ := unpackArgSGodotArray(pA)
	C.godot_pool_color_array_new_with_array(crDest, cpA)
	packSGodotArray(pA, cpA)
	packSGodotPoolColorArray(rDest, crDest)
}

// GodotPoolColorArrayAppend function as declared in gdnative/pool_arrays.h:356
func GodotPoolColorArrayAppend(pSelf []GodotPoolColorArray, pData []GodotColor) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpData, _ := unpackArgSGodotColor(pData)
	C.godot_pool_color_array_append(cpSelf, cpData)
	packSGodotColor(pData, cpData)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayAppendArray function as declared in gdnative/pool_arrays.h:358
func GodotPoolColorArrayAppendArray(pSelf []GodotPoolColorArray, pArray []GodotPoolColorArray) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpArray, _ := unpackArgSGodotPoolColorArray(pArray)
	C.godot_pool_color_array_append_array(cpSelf, cpArray)
	packSGodotPoolColorArray(pArray, cpArray)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayInsert function as declared in gdnative/pool_arrays.h:360
func GodotPoolColorArrayInsert(pSelf []GodotPoolColorArray, pIdx GodotInt, pData []GodotColor) GodotError {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotColor(pData)
	__ret := C.godot_pool_color_array_insert(cpSelf, cpIdx, cpData)
	packSGodotColor(pData, cpData)
	packSGodotPoolColorArray(pSelf, cpSelf)
	__v := (GodotError)(__ret)
	return __v
}

// GodotPoolColorArrayInvert function as declared in gdnative/pool_arrays.h:362
func GodotPoolColorArrayInvert(pSelf []GodotPoolColorArray) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	C.godot_pool_color_array_invert(cpSelf)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayPushBack function as declared in gdnative/pool_arrays.h:364
func GodotPoolColorArrayPushBack(pSelf []GodotPoolColorArray, pData []GodotColor) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpData, _ := unpackArgSGodotColor(pData)
	C.godot_pool_color_array_push_back(cpSelf, cpData)
	packSGodotColor(pData, cpData)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayRemove function as declared in gdnative/pool_arrays.h:366
func GodotPoolColorArrayRemove(pSelf []GodotPoolColorArray, pIdx GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_color_array_remove(cpSelf, cpIdx)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayResize function as declared in gdnative/pool_arrays.h:368
func GodotPoolColorArrayResize(pSelf []GodotPoolColorArray, pSize GodotInt) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_color_array_resize(cpSelf, cpSize)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayRead function as declared in gdnative/pool_arrays.h:370
func GodotPoolColorArrayRead(pSelf []GodotPoolColorArray) *GodotPoolColorArrayReadAccess {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_read(cpSelf)
	packSGodotPoolColorArray(pSelf, cpSelf)
	__v := NewGodotPoolColorArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolColorArrayWrite function as declared in gdnative/pool_arrays.h:372
func GodotPoolColorArrayWrite(pSelf []GodotPoolColorArray) *GodotPoolColorArrayWriteAccess {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_write(cpSelf)
	packSGodotPoolColorArray(pSelf, cpSelf)
	__v := NewGodotPoolColorArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolColorArraySet function as declared in gdnative/pool_arrays.h:374
func GodotPoolColorArraySet(pSelf []GodotPoolColorArray, pIdx GodotInt, pData []GodotColor) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSGodotColor(pData)
	C.godot_pool_color_array_set(cpSelf, cpIdx, cpData)
	packSGodotColor(pData, cpData)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolColorArrayGet function as declared in gdnative/pool_arrays.h:375
func GodotPoolColorArrayGet(pSelf []GodotPoolColorArray, pIdx GodotInt) GodotColor {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_color_array_get(cpSelf, cpIdx)
	packSGodotPoolColorArray(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolColorArraySize function as declared in gdnative/pool_arrays.h:377
func GodotPoolColorArraySize(pSelf []GodotPoolColorArray) GodotInt {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_size(cpSelf)
	packSGodotPoolColorArray(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotPoolColorArrayDestroy function as declared in gdnative/pool_arrays.h:379
func GodotPoolColorArrayDestroy(pSelf []GodotPoolColorArray) {
	cpSelf, _ := unpackArgSGodotPoolColorArray(pSelf)
	C.godot_pool_color_array_destroy(cpSelf)
	packSGodotPoolColorArray(pSelf, cpSelf)
}

// GodotPoolByteArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:385
func GodotPoolByteArrayReadAccessPtr(pRead []GodotPoolByteArrayReadAccess) string {
	cpRead, _ := unpackArgSGodotPoolByteArrayReadAccess(pRead)
	__ret := C.godot_pool_byte_array_read_access_ptr(cpRead)
	packSGodotPoolByteArrayReadAccess(pRead, cpRead)
	__v := packPUint8TString(__ret)
	return __v
}

// GodotPoolByteArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:386
func GodotPoolByteArrayReadAccessOperatorAssign(pRead []GodotPoolByteArrayReadAccess, pOther []GodotPoolByteArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolByteArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolByteArrayReadAccess(pOther)
	C.godot_pool_byte_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolByteArrayReadAccess(pOther, cpOther)
	packSGodotPoolByteArrayReadAccess(pRead, cpRead)
}

// GodotPoolByteArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:387
func GodotPoolByteArrayReadAccessDestroy(pRead []GodotPoolByteArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolByteArrayReadAccess(pRead)
	C.godot_pool_byte_array_read_access_destroy(cpRead)
	packSGodotPoolByteArrayReadAccess(pRead, cpRead)
}

// GodotPoolIntArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:389
func GodotPoolIntArrayReadAccessPtr(pRead []GodotPoolIntArrayReadAccess) *GodotInt {
	cpRead, _ := unpackArgSGodotPoolIntArrayReadAccess(pRead)
	__ret := C.godot_pool_int_array_read_access_ptr(cpRead)
	packSGodotPoolIntArrayReadAccess(pRead, cpRead)
	__v := *(**GodotInt)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolIntArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:390
func GodotPoolIntArrayReadAccessOperatorAssign(pRead []GodotPoolIntArrayReadAccess, pOther []GodotPoolIntArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolIntArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolIntArrayReadAccess(pOther)
	C.godot_pool_int_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolIntArrayReadAccess(pOther, cpOther)
	packSGodotPoolIntArrayReadAccess(pRead, cpRead)
}

// GodotPoolIntArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:391
func GodotPoolIntArrayReadAccessDestroy(pRead []GodotPoolIntArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolIntArrayReadAccess(pRead)
	C.godot_pool_int_array_read_access_destroy(cpRead)
	packSGodotPoolIntArrayReadAccess(pRead, cpRead)
}

// GodotPoolRealArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:393
func GodotPoolRealArrayReadAccessPtr(pRead []GodotPoolRealArrayReadAccess) *GodotReal {
	cpRead, _ := unpackArgSGodotPoolRealArrayReadAccess(pRead)
	__ret := C.godot_pool_real_array_read_access_ptr(cpRead)
	packSGodotPoolRealArrayReadAccess(pRead, cpRead)
	__v := *(**GodotReal)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolRealArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:394
func GodotPoolRealArrayReadAccessOperatorAssign(pRead []GodotPoolRealArrayReadAccess, pOther []GodotPoolRealArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolRealArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolRealArrayReadAccess(pOther)
	C.godot_pool_real_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolRealArrayReadAccess(pOther, cpOther)
	packSGodotPoolRealArrayReadAccess(pRead, cpRead)
}

// GodotPoolRealArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:395
func GodotPoolRealArrayReadAccessDestroy(pRead []GodotPoolRealArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolRealArrayReadAccess(pRead)
	C.godot_pool_real_array_read_access_destroy(cpRead)
	packSGodotPoolRealArrayReadAccess(pRead, cpRead)
}

// GodotPoolStringArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:397
func GodotPoolStringArrayReadAccessPtr(pRead []GodotPoolStringArrayReadAccess) *GodotString {
	cpRead, _ := unpackArgSGodotPoolStringArrayReadAccess(pRead)
	__ret := C.godot_pool_string_array_read_access_ptr(cpRead)
	packSGodotPoolStringArrayReadAccess(pRead, cpRead)
	__v := NewGodotStringRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolStringArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:398
func GodotPoolStringArrayReadAccessOperatorAssign(pRead []GodotPoolStringArrayReadAccess, pOther []GodotPoolStringArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolStringArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolStringArrayReadAccess(pOther)
	C.godot_pool_string_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolStringArrayReadAccess(pOther, cpOther)
	packSGodotPoolStringArrayReadAccess(pRead, cpRead)
}

// GodotPoolStringArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:399
func GodotPoolStringArrayReadAccessDestroy(pRead []GodotPoolStringArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolStringArrayReadAccess(pRead)
	C.godot_pool_string_array_read_access_destroy(cpRead)
	packSGodotPoolStringArrayReadAccess(pRead, cpRead)
}

// GodotPoolVector2ArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:401
func GodotPoolVector2ArrayReadAccessPtr(pRead []GodotPoolVector2ArrayReadAccess) *GodotVector2 {
	cpRead, _ := unpackArgSGodotPoolVector2ArrayReadAccess(pRead)
	__ret := C.godot_pool_vector2_array_read_access_ptr(cpRead)
	packSGodotPoolVector2ArrayReadAccess(pRead, cpRead)
	__v := NewGodotVector2Ref(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector2ArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:402
func GodotPoolVector2ArrayReadAccessOperatorAssign(pRead []GodotPoolVector2ArrayReadAccess, pOther []GodotPoolVector2ArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolVector2ArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolVector2ArrayReadAccess(pOther)
	C.godot_pool_vector2_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolVector2ArrayReadAccess(pOther, cpOther)
	packSGodotPoolVector2ArrayReadAccess(pRead, cpRead)
}

// GodotPoolVector2ArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:403
func GodotPoolVector2ArrayReadAccessDestroy(pRead []GodotPoolVector2ArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolVector2ArrayReadAccess(pRead)
	C.godot_pool_vector2_array_read_access_destroy(cpRead)
	packSGodotPoolVector2ArrayReadAccess(pRead, cpRead)
}

// GodotPoolVector3ArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:405
func GodotPoolVector3ArrayReadAccessPtr(pRead []GodotPoolVector3ArrayReadAccess) *GodotVector3 {
	cpRead, _ := unpackArgSGodotPoolVector3ArrayReadAccess(pRead)
	__ret := C.godot_pool_vector3_array_read_access_ptr(cpRead)
	packSGodotPoolVector3ArrayReadAccess(pRead, cpRead)
	__v := NewGodotVector3Ref(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector3ArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:406
func GodotPoolVector3ArrayReadAccessOperatorAssign(pRead []GodotPoolVector3ArrayReadAccess, pOther []GodotPoolVector3ArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolVector3ArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolVector3ArrayReadAccess(pOther)
	C.godot_pool_vector3_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolVector3ArrayReadAccess(pOther, cpOther)
	packSGodotPoolVector3ArrayReadAccess(pRead, cpRead)
}

// GodotPoolVector3ArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:407
func GodotPoolVector3ArrayReadAccessDestroy(pRead []GodotPoolVector3ArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolVector3ArrayReadAccess(pRead)
	C.godot_pool_vector3_array_read_access_destroy(cpRead)
	packSGodotPoolVector3ArrayReadAccess(pRead, cpRead)
}

// GodotPoolColorArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:409
func GodotPoolColorArrayReadAccessPtr(pRead []GodotPoolColorArrayReadAccess) *GodotColor {
	cpRead, _ := unpackArgSGodotPoolColorArrayReadAccess(pRead)
	__ret := C.godot_pool_color_array_read_access_ptr(cpRead)
	packSGodotPoolColorArrayReadAccess(pRead, cpRead)
	__v := NewGodotColorRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolColorArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:410
func GodotPoolColorArrayReadAccessOperatorAssign(pRead []GodotPoolColorArrayReadAccess, pOther []GodotPoolColorArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolColorArrayReadAccess(pRead)
	cpOther, _ := unpackArgSGodotPoolColorArrayReadAccess(pOther)
	C.godot_pool_color_array_read_access_operator_assign(cpRead, cpOther)
	packSGodotPoolColorArrayReadAccess(pOther, cpOther)
	packSGodotPoolColorArrayReadAccess(pRead, cpRead)
}

// GodotPoolColorArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:411
func GodotPoolColorArrayReadAccessDestroy(pRead []GodotPoolColorArrayReadAccess) {
	cpRead, _ := unpackArgSGodotPoolColorArrayReadAccess(pRead)
	C.godot_pool_color_array_read_access_destroy(cpRead)
	packSGodotPoolColorArrayReadAccess(pRead, cpRead)
}

// GodotPoolByteArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:417
func GodotPoolByteArrayWriteAccessPtr(pWrite []GodotPoolByteArrayWriteAccess) *byte {
	cpWrite, _ := unpackArgSGodotPoolByteArrayWriteAccess(pWrite)
	__ret := C.godot_pool_byte_array_write_access_ptr(cpWrite)
	packSGodotPoolByteArrayWriteAccess(pWrite, cpWrite)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolByteArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:418
func GodotPoolByteArrayWriteAccessOperatorAssign(pWrite []GodotPoolByteArrayWriteAccess, pOther []GodotPoolByteArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolByteArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolByteArrayWriteAccess(pOther)
	C.godot_pool_byte_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolByteArrayWriteAccess(pOther, cpOther)
	packSGodotPoolByteArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolByteArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:419
func GodotPoolByteArrayWriteAccessDestroy(pWrite []GodotPoolByteArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolByteArrayWriteAccess(pWrite)
	C.godot_pool_byte_array_write_access_destroy(cpWrite)
	packSGodotPoolByteArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolIntArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:421
func GodotPoolIntArrayWriteAccessPtr(pWrite []GodotPoolIntArrayWriteAccess) *GodotInt {
	cpWrite, _ := unpackArgSGodotPoolIntArrayWriteAccess(pWrite)
	__ret := C.godot_pool_int_array_write_access_ptr(cpWrite)
	packSGodotPoolIntArrayWriteAccess(pWrite, cpWrite)
	__v := *(**GodotInt)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolIntArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:422
func GodotPoolIntArrayWriteAccessOperatorAssign(pWrite []GodotPoolIntArrayWriteAccess, pOther []GodotPoolIntArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolIntArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolIntArrayWriteAccess(pOther)
	C.godot_pool_int_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolIntArrayWriteAccess(pOther, cpOther)
	packSGodotPoolIntArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolIntArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:423
func GodotPoolIntArrayWriteAccessDestroy(pWrite []GodotPoolIntArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolIntArrayWriteAccess(pWrite)
	C.godot_pool_int_array_write_access_destroy(cpWrite)
	packSGodotPoolIntArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolRealArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:425
func GodotPoolRealArrayWriteAccessPtr(pWrite []GodotPoolRealArrayWriteAccess) *GodotReal {
	cpWrite, _ := unpackArgSGodotPoolRealArrayWriteAccess(pWrite)
	__ret := C.godot_pool_real_array_write_access_ptr(cpWrite)
	packSGodotPoolRealArrayWriteAccess(pWrite, cpWrite)
	__v := *(**GodotReal)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPoolRealArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:426
func GodotPoolRealArrayWriteAccessOperatorAssign(pWrite []GodotPoolRealArrayWriteAccess, pOther []GodotPoolRealArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolRealArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolRealArrayWriteAccess(pOther)
	C.godot_pool_real_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolRealArrayWriteAccess(pOther, cpOther)
	packSGodotPoolRealArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolRealArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:427
func GodotPoolRealArrayWriteAccessDestroy(pWrite []GodotPoolRealArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolRealArrayWriteAccess(pWrite)
	C.godot_pool_real_array_write_access_destroy(cpWrite)
	packSGodotPoolRealArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolStringArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:429
func GodotPoolStringArrayWriteAccessPtr(pWrite []GodotPoolStringArrayWriteAccess) *GodotString {
	cpWrite, _ := unpackArgSGodotPoolStringArrayWriteAccess(pWrite)
	__ret := C.godot_pool_string_array_write_access_ptr(cpWrite)
	packSGodotPoolStringArrayWriteAccess(pWrite, cpWrite)
	__v := NewGodotStringRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolStringArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:430
func GodotPoolStringArrayWriteAccessOperatorAssign(pWrite []GodotPoolStringArrayWriteAccess, pOther []GodotPoolStringArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolStringArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolStringArrayWriteAccess(pOther)
	C.godot_pool_string_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolStringArrayWriteAccess(pOther, cpOther)
	packSGodotPoolStringArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolStringArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:431
func GodotPoolStringArrayWriteAccessDestroy(pWrite []GodotPoolStringArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolStringArrayWriteAccess(pWrite)
	C.godot_pool_string_array_write_access_destroy(cpWrite)
	packSGodotPoolStringArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolVector2ArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:433
func GodotPoolVector2ArrayWriteAccessPtr(pWrite []GodotPoolVector2ArrayWriteAccess) *GodotVector2 {
	cpWrite, _ := unpackArgSGodotPoolVector2ArrayWriteAccess(pWrite)
	__ret := C.godot_pool_vector2_array_write_access_ptr(cpWrite)
	packSGodotPoolVector2ArrayWriteAccess(pWrite, cpWrite)
	__v := NewGodotVector2Ref(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector2ArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:434
func GodotPoolVector2ArrayWriteAccessOperatorAssign(pWrite []GodotPoolVector2ArrayWriteAccess, pOther []GodotPoolVector2ArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolVector2ArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolVector2ArrayWriteAccess(pOther)
	C.godot_pool_vector2_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolVector2ArrayWriteAccess(pOther, cpOther)
	packSGodotPoolVector2ArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolVector2ArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:435
func GodotPoolVector2ArrayWriteAccessDestroy(pWrite []GodotPoolVector2ArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolVector2ArrayWriteAccess(pWrite)
	C.godot_pool_vector2_array_write_access_destroy(cpWrite)
	packSGodotPoolVector2ArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolVector3ArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:437
func GodotPoolVector3ArrayWriteAccessPtr(pWrite []GodotPoolVector3ArrayWriteAccess) *GodotVector3 {
	cpWrite, _ := unpackArgSGodotPoolVector3ArrayWriteAccess(pWrite)
	__ret := C.godot_pool_vector3_array_write_access_ptr(cpWrite)
	packSGodotPoolVector3ArrayWriteAccess(pWrite, cpWrite)
	__v := NewGodotVector3Ref(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolVector3ArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:438
func GodotPoolVector3ArrayWriteAccessOperatorAssign(pWrite []GodotPoolVector3ArrayWriteAccess, pOther []GodotPoolVector3ArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolVector3ArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolVector3ArrayWriteAccess(pOther)
	C.godot_pool_vector3_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolVector3ArrayWriteAccess(pOther, cpOther)
	packSGodotPoolVector3ArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolVector3ArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:439
func GodotPoolVector3ArrayWriteAccessDestroy(pWrite []GodotPoolVector3ArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolVector3ArrayWriteAccess(pWrite)
	C.godot_pool_vector3_array_write_access_destroy(cpWrite)
	packSGodotPoolVector3ArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolColorArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:441
func GodotPoolColorArrayWriteAccessPtr(pWrite []GodotPoolColorArrayWriteAccess) *GodotColor {
	cpWrite, _ := unpackArgSGodotPoolColorArrayWriteAccess(pWrite)
	__ret := C.godot_pool_color_array_write_access_ptr(cpWrite)
	packSGodotPoolColorArrayWriteAccess(pWrite, cpWrite)
	__v := NewGodotColorRef(unsafe.Pointer(__ret))
	return __v
}

// GodotPoolColorArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:442
func GodotPoolColorArrayWriteAccessOperatorAssign(pWrite []GodotPoolColorArrayWriteAccess, pOther []GodotPoolColorArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolColorArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSGodotPoolColorArrayWriteAccess(pOther)
	C.godot_pool_color_array_write_access_operator_assign(cpWrite, cpOther)
	packSGodotPoolColorArrayWriteAccess(pOther, cpOther)
	packSGodotPoolColorArrayWriteAccess(pWrite, cpWrite)
}

// GodotPoolColorArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:443
func GodotPoolColorArrayWriteAccessDestroy(pWrite []GodotPoolColorArrayWriteAccess) {
	cpWrite, _ := unpackArgSGodotPoolColorArrayWriteAccess(pWrite)
	C.godot_pool_color_array_write_access_destroy(cpWrite)
	packSGodotPoolColorArrayWriteAccess(pWrite, cpWrite)
}

// GodotColorNewRgba function as declared in gdnative/color.h:60
func GodotColorNewRgba(rDest []GodotColor, pR GodotReal, pG GodotReal, pB GodotReal, pA GodotReal) {
	crDest, _ := unpackArgSGodotColor(rDest)
	cpR, _ := (C.godot_real)(pR), cgoAllocsUnknown
	cpG, _ := (C.godot_real)(pG), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	cpA, _ := (C.godot_real)(pA), cgoAllocsUnknown
	C.godot_color_new_rgba(crDest, cpR, cpG, cpB, cpA)
	packSGodotColor(rDest, crDest)
}

// GodotColorNewRgb function as declared in gdnative/color.h:61
func GodotColorNewRgb(rDest []GodotColor, pR GodotReal, pG GodotReal, pB GodotReal) {
	crDest, _ := unpackArgSGodotColor(rDest)
	cpR, _ := (C.godot_real)(pR), cgoAllocsUnknown
	cpG, _ := (C.godot_real)(pG), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	C.godot_color_new_rgb(crDest, cpR, cpG, cpB)
	packSGodotColor(rDest, crDest)
}

// GodotColorGetR function as declared in gdnative/color.h:63
func GodotColorGetR(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_r(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorSetR function as declared in gdnative/color.h:64
func GodotColorSetR(pSelf []GodotColor, r GodotReal) {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cr, _ := (C.godot_real)(r), cgoAllocsUnknown
	C.godot_color_set_r(cpSelf, cr)
	packSGodotColor(pSelf, cpSelf)
}

// GodotColorGetG function as declared in gdnative/color.h:66
func GodotColorGetG(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_g(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorSetG function as declared in gdnative/color.h:67
func GodotColorSetG(pSelf []GodotColor, g GodotReal) {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cg, _ := (C.godot_real)(g), cgoAllocsUnknown
	C.godot_color_set_g(cpSelf, cg)
	packSGodotColor(pSelf, cpSelf)
}

// GodotColorGetB function as declared in gdnative/color.h:69
func GodotColorGetB(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_b(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorSetB function as declared in gdnative/color.h:70
func GodotColorSetB(pSelf []GodotColor, b GodotReal) {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cb, _ := (C.godot_real)(b), cgoAllocsUnknown
	C.godot_color_set_b(cpSelf, cb)
	packSGodotColor(pSelf, cpSelf)
}

// GodotColorGetA function as declared in gdnative/color.h:72
func GodotColorGetA(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_a(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorSetA function as declared in gdnative/color.h:73
func GodotColorSetA(pSelf []GodotColor, a GodotReal) {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	ca, _ := (C.godot_real)(a), cgoAllocsUnknown
	C.godot_color_set_a(cpSelf, ca)
	packSGodotColor(pSelf, cpSelf)
}

// GodotColorGetH function as declared in gdnative/color.h:75
func GodotColorGetH(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_h(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorGetS function as declared in gdnative/color.h:76
func GodotColorGetS(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_s(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorGetV function as declared in gdnative/color.h:77
func GodotColorGetV(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_get_v(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorAsString function as declared in gdnative/color.h:79
func GodotColorAsString(pSelf []GodotColor) GodotString {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_as_string(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorToRgba32 function as declared in gdnative/color.h:81
func GodotColorToRgba32(pSelf []GodotColor) GodotInt {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_to_rgba32(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotColorToArgb32 function as declared in gdnative/color.h:83
func GodotColorToArgb32(pSelf []GodotColor) GodotInt {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_to_argb32(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotColorGray function as declared in gdnative/color.h:85
func GodotColorGray(pSelf []GodotColor) GodotReal {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_gray(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotColorInverted function as declared in gdnative/color.h:87
func GodotColorInverted(pSelf []GodotColor) GodotColor {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_inverted(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorContrasted function as declared in gdnative/color.h:89
func GodotColorContrasted(pSelf []GodotColor) GodotColor {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	__ret := C.godot_color_contrasted(cpSelf)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorLinearInterpolate function as declared in gdnative/color.h:91
func GodotColorLinearInterpolate(pSelf []GodotColor, pB []GodotColor, pT GodotReal) GodotColor {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cpB, _ := unpackArgSGodotColor(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_color_linear_interpolate(cpSelf, cpB, cpT)
	packSGodotColor(pB, cpB)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorBlend function as declared in gdnative/color.h:93
func GodotColorBlend(pSelf []GodotColor, pOver []GodotColor) GodotColor {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cpOver, _ := unpackArgSGodotColor(pOver)
	__ret := C.godot_color_blend(cpSelf, cpOver)
	packSGodotColor(pOver, cpOver)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorToHtml function as declared in gdnative/color.h:95
func GodotColorToHtml(pSelf []GodotColor, pWithAlpha GodotBool) GodotString {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cpWithAlpha, _ := (C.godot_bool)(pWithAlpha), cgoAllocsUnknown
	__ret := C.godot_color_to_html(cpSelf, cpWithAlpha)
	packSGodotColor(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotColorOperatorEqual function as declared in gdnative/color.h:97
func GodotColorOperatorEqual(pSelf []GodotColor, pB []GodotColor) GodotBool {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cpB, _ := unpackArgSGodotColor(pB)
	__ret := C.godot_color_operator_equal(cpSelf, cpB)
	packSGodotColor(pB, cpB)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotColorOperatorLess function as declared in gdnative/color.h:99
func GodotColorOperatorLess(pSelf []GodotColor, pB []GodotColor) GodotBool {
	cpSelf, _ := unpackArgSGodotColor(pSelf)
	cpB, _ := unpackArgSGodotColor(pB)
	__ret := C.godot_color_operator_less(cpSelf, cpB)
	packSGodotColor(pB, cpB)
	packSGodotColor(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector2New function as declared in gdnative/vector2.h:59
func GodotVector2New(rDest []GodotVector2, pX GodotReal, pY GodotReal) {
	crDest, _ := unpackArgSGodotVector2(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	C.godot_vector2_new(crDest, cpX, cpY)
	packSGodotVector2(rDest, crDest)
}

// GodotVector2AsString function as declared in gdnative/vector2.h:61
func GodotVector2AsString(pSelf []GodotVector2) GodotString {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_as_string(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Normalized function as declared in gdnative/vector2.h:63
func GodotVector2Normalized(pSelf []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_normalized(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Length function as declared in gdnative/vector2.h:65
func GodotVector2Length(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_length(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2Angle function as declared in gdnative/vector2.h:67
func GodotVector2Angle(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_angle(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2LengthSquared function as declared in gdnative/vector2.h:69
func GodotVector2LengthSquared(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_length_squared(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2IsNormalized function as declared in gdnative/vector2.h:71
func GodotVector2IsNormalized(pSelf []GodotVector2) GodotBool {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_is_normalized(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector2DistanceTo function as declared in gdnative/vector2.h:73
func GodotVector2DistanceTo(pSelf []GodotVector2, pTo []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpTo, _ := unpackArgSGodotVector2(pTo)
	__ret := C.godot_vector2_distance_to(cpSelf, cpTo)
	packSGodotVector2(pTo, cpTo)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2DistanceSquaredTo function as declared in gdnative/vector2.h:75
func GodotVector2DistanceSquaredTo(pSelf []GodotVector2, pTo []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpTo, _ := unpackArgSGodotVector2(pTo)
	__ret := C.godot_vector2_distance_squared_to(cpSelf, cpTo)
	packSGodotVector2(pTo, cpTo)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2AngleTo function as declared in gdnative/vector2.h:77
func GodotVector2AngleTo(pSelf []GodotVector2, pTo []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpTo, _ := unpackArgSGodotVector2(pTo)
	__ret := C.godot_vector2_angle_to(cpSelf, cpTo)
	packSGodotVector2(pTo, cpTo)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2AngleToPoint function as declared in gdnative/vector2.h:79
func GodotVector2AngleToPoint(pSelf []GodotVector2, pTo []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpTo, _ := unpackArgSGodotVector2(pTo)
	__ret := C.godot_vector2_angle_to_point(cpSelf, cpTo)
	packSGodotVector2(pTo, cpTo)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2LinearInterpolate function as declared in gdnative/vector2.h:81
func GodotVector2LinearInterpolate(pSelf []GodotVector2, pB []GodotVector2, pT GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector2_linear_interpolate(cpSelf, cpB, cpT)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2CubicInterpolate function as declared in gdnative/vector2.h:83
func GodotVector2CubicInterpolate(pSelf []GodotVector2, pB []GodotVector2, pPreA []GodotVector2, pPostB []GodotVector2, pT GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	cpPreA, _ := unpackArgSGodotVector2(pPreA)
	cpPostB, _ := unpackArgSGodotVector2(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector2_cubic_interpolate(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSGodotVector2(pPostB, cpPostB)
	packSGodotVector2(pPreA, cpPreA)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Rotated function as declared in gdnative/vector2.h:85
func GodotVector2Rotated(pSelf []GodotVector2, pPhi GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_vector2_rotated(cpSelf, cpPhi)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Tangent function as declared in gdnative/vector2.h:87
func GodotVector2Tangent(pSelf []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_tangent(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Floor function as declared in gdnative/vector2.h:89
func GodotVector2Floor(pSelf []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_floor(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Snapped function as declared in gdnative/vector2.h:91
func GodotVector2Snapped(pSelf []GodotVector2, pBy []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpBy, _ := unpackArgSGodotVector2(pBy)
	__ret := C.godot_vector2_snapped(cpSelf, cpBy)
	packSGodotVector2(pBy, cpBy)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Aspect function as declared in gdnative/vector2.h:93
func GodotVector2Aspect(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_aspect(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2Dot function as declared in gdnative/vector2.h:95
func GodotVector2Dot(pSelf []GodotVector2, pWith []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpWith, _ := unpackArgSGodotVector2(pWith)
	__ret := C.godot_vector2_dot(cpSelf, cpWith)
	packSGodotVector2(pWith, cpWith)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2Slide function as declared in gdnative/vector2.h:97
func GodotVector2Slide(pSelf []GodotVector2, pN []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpN, _ := unpackArgSGodotVector2(pN)
	__ret := C.godot_vector2_slide(cpSelf, cpN)
	packSGodotVector2(pN, cpN)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Bounce function as declared in gdnative/vector2.h:99
func GodotVector2Bounce(pSelf []GodotVector2, pN []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpN, _ := unpackArgSGodotVector2(pN)
	__ret := C.godot_vector2_bounce(cpSelf, cpN)
	packSGodotVector2(pN, cpN)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Reflect function as declared in gdnative/vector2.h:101
func GodotVector2Reflect(pSelf []GodotVector2, pN []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpN, _ := unpackArgSGodotVector2(pN)
	__ret := C.godot_vector2_reflect(cpSelf, cpN)
	packSGodotVector2(pN, cpN)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Abs function as declared in gdnative/vector2.h:103
func GodotVector2Abs(pSelf []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_abs(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2Clamped function as declared in gdnative/vector2.h:105
func GodotVector2Clamped(pSelf []GodotVector2, pLength GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpLength, _ := (C.godot_real)(pLength), cgoAllocsUnknown
	__ret := C.godot_vector2_clamped(cpSelf, cpLength)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorAdd function as declared in gdnative/vector2.h:107
func GodotVector2OperatorAdd(pSelf []GodotVector2, pB []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_add(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorSubstract function as declared in gdnative/vector2.h:109
func GodotVector2OperatorSubstract(pSelf []GodotVector2, pB []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_substract(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorMultiplyVector function as declared in gdnative/vector2.h:111
func GodotVector2OperatorMultiplyVector(pSelf []GodotVector2, pB []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_multiply_vector(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorMultiplyScalar function as declared in gdnative/vector2.h:113
func GodotVector2OperatorMultiplyScalar(pSelf []GodotVector2, pB GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector2_operator_multiply_scalar(cpSelf, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorDivideVector function as declared in gdnative/vector2.h:115
func GodotVector2OperatorDivideVector(pSelf []GodotVector2, pB []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_divide_vector(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorDivideScalar function as declared in gdnative/vector2.h:117
func GodotVector2OperatorDivideScalar(pSelf []GodotVector2, pB GodotReal) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector2_operator_divide_scalar(cpSelf, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2OperatorEqual function as declared in gdnative/vector2.h:119
func GodotVector2OperatorEqual(pSelf []GodotVector2, pB []GodotVector2) GodotBool {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_equal(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector2OperatorLess function as declared in gdnative/vector2.h:121
func GodotVector2OperatorLess(pSelf []GodotVector2, pB []GodotVector2) GodotBool {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpB, _ := unpackArgSGodotVector2(pB)
	__ret := C.godot_vector2_operator_less(cpSelf, cpB)
	packSGodotVector2(pB, cpB)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector2OperatorNeg function as declared in gdnative/vector2.h:123
func GodotVector2OperatorNeg(pSelf []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_operator_neg(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector2SetX function as declared in gdnative/vector2.h:125
func GodotVector2SetX(pSelf []GodotVector2, pX GodotReal) {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	C.godot_vector2_set_x(cpSelf, cpX)
	packSGodotVector2(pSelf, cpSelf)
}

// GodotVector2SetY function as declared in gdnative/vector2.h:127
func GodotVector2SetY(pSelf []GodotVector2, pY GodotReal) {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	C.godot_vector2_set_y(cpSelf, cpY)
	packSGodotVector2(pSelf, cpSelf)
}

// GodotVector2GetX function as declared in gdnative/vector2.h:129
func GodotVector2GetX(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_get_x(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector2GetY function as declared in gdnative/vector2.h:131
func GodotVector2GetY(pSelf []GodotVector2) GodotReal {
	cpSelf, _ := unpackArgSGodotVector2(pSelf)
	__ret := C.godot_vector2_get_y(cpSelf)
	packSGodotVector2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3New function as declared in gdnative/vector3.h:66
func GodotVector3New(rDest []GodotVector3, pX GodotReal, pY GodotReal, pZ GodotReal) {
	crDest, _ := unpackArgSGodotVector3(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpZ, _ := (C.godot_real)(pZ), cgoAllocsUnknown
	C.godot_vector3_new(crDest, cpX, cpY, cpZ)
	packSGodotVector3(rDest, crDest)
}

// GodotVector3AsString function as declared in gdnative/vector3.h:68
func GodotVector3AsString(pSelf []GodotVector3) GodotString {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_as_string(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3MinAxis function as declared in gdnative/vector3.h:70
func GodotVector3MinAxis(pSelf []GodotVector3) GodotInt {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_min_axis(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotVector3MaxAxis function as declared in gdnative/vector3.h:72
func GodotVector3MaxAxis(pSelf []GodotVector3) GodotInt {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_max_axis(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotVector3Length function as declared in gdnative/vector3.h:74
func GodotVector3Length(pSelf []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_length(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3LengthSquared function as declared in gdnative/vector3.h:76
func GodotVector3LengthSquared(pSelf []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_length_squared(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3IsNormalized function as declared in gdnative/vector3.h:78
func GodotVector3IsNormalized(pSelf []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_is_normalized(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector3Normalized function as declared in gdnative/vector3.h:80
func GodotVector3Normalized(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_normalized(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Inverse function as declared in gdnative/vector3.h:82
func GodotVector3Inverse(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_inverse(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Snapped function as declared in gdnative/vector3.h:84
func GodotVector3Snapped(pSelf []GodotVector3, pBy []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpBy, _ := unpackArgSGodotVector3(pBy)
	__ret := C.godot_vector3_snapped(cpSelf, cpBy)
	packSGodotVector3(pBy, cpBy)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Rotated function as declared in gdnative/vector3.h:86
func GodotVector3Rotated(pSelf []GodotVector3, pAxis []GodotVector3, pPhi GodotReal) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpAxis, _ := unpackArgSGodotVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_vector3_rotated(cpSelf, cpAxis, cpPhi)
	packSGodotVector3(pAxis, cpAxis)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3LinearInterpolate function as declared in gdnative/vector3.h:88
func GodotVector3LinearInterpolate(pSelf []GodotVector3, pB []GodotVector3, pT GodotReal) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector3_linear_interpolate(cpSelf, cpB, cpT)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3CubicInterpolate function as declared in gdnative/vector3.h:90
func GodotVector3CubicInterpolate(pSelf []GodotVector3, pB []GodotVector3, pPreA []GodotVector3, pPostB []GodotVector3, pT GodotReal) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	cpPreA, _ := unpackArgSGodotVector3(pPreA)
	cpPostB, _ := unpackArgSGodotVector3(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector3_cubic_interpolate(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSGodotVector3(pPostB, cpPostB)
	packSGodotVector3(pPreA, cpPreA)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Dot function as declared in gdnative/vector3.h:92
func GodotVector3Dot(pSelf []GodotVector3, pB []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_dot(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3Cross function as declared in gdnative/vector3.h:94
func GodotVector3Cross(pSelf []GodotVector3, pB []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_cross(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Outer function as declared in gdnative/vector3.h:96
func GodotVector3Outer(pSelf []GodotVector3, pB []GodotVector3) GodotBasis {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_outer(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3ToDiagonalMatrix function as declared in gdnative/vector3.h:98
func GodotVector3ToDiagonalMatrix(pSelf []GodotVector3) GodotBasis {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_to_diagonal_matrix(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Abs function as declared in gdnative/vector3.h:100
func GodotVector3Abs(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_abs(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Floor function as declared in gdnative/vector3.h:102
func GodotVector3Floor(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_floor(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Ceil function as declared in gdnative/vector3.h:104
func GodotVector3Ceil(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_ceil(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3DistanceTo function as declared in gdnative/vector3.h:106
func GodotVector3DistanceTo(pSelf []GodotVector3, pB []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_distance_to(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3DistanceSquaredTo function as declared in gdnative/vector3.h:108
func GodotVector3DistanceSquaredTo(pSelf []GodotVector3, pB []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_distance_squared_to(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3AngleTo function as declared in gdnative/vector3.h:110
func GodotVector3AngleTo(pSelf []GodotVector3, pTo []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpTo, _ := unpackArgSGodotVector3(pTo)
	__ret := C.godot_vector3_angle_to(cpSelf, cpTo)
	packSGodotVector3(pTo, cpTo)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotVector3Slide function as declared in gdnative/vector3.h:112
func GodotVector3Slide(pSelf []GodotVector3, pN []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpN, _ := unpackArgSGodotVector3(pN)
	__ret := C.godot_vector3_slide(cpSelf, cpN)
	packSGodotVector3(pN, cpN)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Bounce function as declared in gdnative/vector3.h:114
func GodotVector3Bounce(pSelf []GodotVector3, pN []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpN, _ := unpackArgSGodotVector3(pN)
	__ret := C.godot_vector3_bounce(cpSelf, cpN)
	packSGodotVector3(pN, cpN)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3Reflect function as declared in gdnative/vector3.h:116
func GodotVector3Reflect(pSelf []GodotVector3, pN []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpN, _ := unpackArgSGodotVector3(pN)
	__ret := C.godot_vector3_reflect(cpSelf, cpN)
	packSGodotVector3(pN, cpN)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorAdd function as declared in gdnative/vector3.h:118
func GodotVector3OperatorAdd(pSelf []GodotVector3, pB []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_add(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorSubstract function as declared in gdnative/vector3.h:120
func GodotVector3OperatorSubstract(pSelf []GodotVector3, pB []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_substract(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorMultiplyVector function as declared in gdnative/vector3.h:122
func GodotVector3OperatorMultiplyVector(pSelf []GodotVector3, pB []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_multiply_vector(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorMultiplyScalar function as declared in gdnative/vector3.h:124
func GodotVector3OperatorMultiplyScalar(pSelf []GodotVector3, pB GodotReal) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector3_operator_multiply_scalar(cpSelf, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorDivideVector function as declared in gdnative/vector3.h:126
func GodotVector3OperatorDivideVector(pSelf []GodotVector3, pB []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_divide_vector(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorDivideScalar function as declared in gdnative/vector3.h:128
func GodotVector3OperatorDivideScalar(pSelf []GodotVector3, pB GodotReal) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector3_operator_divide_scalar(cpSelf, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3OperatorEqual function as declared in gdnative/vector3.h:130
func GodotVector3OperatorEqual(pSelf []GodotVector3, pB []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_equal(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector3OperatorLess function as declared in gdnative/vector3.h:132
func GodotVector3OperatorLess(pSelf []GodotVector3, pB []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpB, _ := unpackArgSGodotVector3(pB)
	__ret := C.godot_vector3_operator_less(cpSelf, cpB)
	packSGodotVector3(pB, cpB)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVector3OperatorNeg function as declared in gdnative/vector3.h:134
func GodotVector3OperatorNeg(pSelf []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	__ret := C.godot_vector3_operator_neg(cpSelf)
	packSGodotVector3(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVector3SetAxis function as declared in gdnative/vector3.h:136
func GodotVector3SetAxis(pSelf []GodotVector3, pAxis GodotVector3Axis, pVal GodotReal) {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpAxis, _ := (C.godot_vector3_axis)(pAxis), cgoAllocsUnknown
	cpVal, _ := (C.godot_real)(pVal), cgoAllocsUnknown
	C.godot_vector3_set_axis(cpSelf, cpAxis, cpVal)
	packSGodotVector3(pSelf, cpSelf)
}

// GodotVector3GetAxis function as declared in gdnative/vector3.h:138
func GodotVector3GetAxis(pSelf []GodotVector3, pAxis GodotVector3Axis) GodotReal {
	cpSelf, _ := unpackArgSGodotVector3(pSelf)
	cpAxis, _ := (C.godot_vector3_axis)(pAxis), cgoAllocsUnknown
	__ret := C.godot_vector3_get_axis(cpSelf, cpAxis)
	packSGodotVector3(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotBasisNewWithRows function as declared in gdnative/basis.h:61
func GodotBasisNewWithRows(rDest []GodotBasis, pXAxis []GodotVector3, pYAxis []GodotVector3, pZAxis []GodotVector3) {
	crDest, _ := unpackArgSGodotBasis(rDest)
	cpXAxis, _ := unpackArgSGodotVector3(pXAxis)
	cpYAxis, _ := unpackArgSGodotVector3(pYAxis)
	cpZAxis, _ := unpackArgSGodotVector3(pZAxis)
	C.godot_basis_new_with_rows(crDest, cpXAxis, cpYAxis, cpZAxis)
	packSGodotVector3(pZAxis, cpZAxis)
	packSGodotVector3(pYAxis, cpYAxis)
	packSGodotVector3(pXAxis, cpXAxis)
	packSGodotBasis(rDest, crDest)
}

// GodotBasisNewWithAxisAndAngle function as declared in gdnative/basis.h:62
func GodotBasisNewWithAxisAndAngle(rDest []GodotBasis, pAxis []GodotVector3, pPhi GodotReal) {
	crDest, _ := unpackArgSGodotBasis(rDest)
	cpAxis, _ := unpackArgSGodotVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	C.godot_basis_new_with_axis_and_angle(crDest, cpAxis, cpPhi)
	packSGodotVector3(pAxis, cpAxis)
	packSGodotBasis(rDest, crDest)
}

// GodotBasisNewWithEuler function as declared in gdnative/basis.h:63
func GodotBasisNewWithEuler(rDest []GodotBasis, pEuler []GodotVector3) {
	crDest, _ := unpackArgSGodotBasis(rDest)
	cpEuler, _ := unpackArgSGodotVector3(pEuler)
	C.godot_basis_new_with_euler(crDest, cpEuler)
	packSGodotVector3(pEuler, cpEuler)
	packSGodotBasis(rDest, crDest)
}

// GodotBasisAsString function as declared in gdnative/basis.h:65
func GodotBasisAsString(pSelf []GodotBasis) GodotString {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_as_string(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisInverse function as declared in gdnative/basis.h:67
func GodotBasisInverse(pSelf []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_inverse(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisTransposed function as declared in gdnative/basis.h:69
func GodotBasisTransposed(pSelf []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_transposed(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisOrthonormalized function as declared in gdnative/basis.h:71
func GodotBasisOrthonormalized(pSelf []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_orthonormalized(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisDeterminant function as declared in gdnative/basis.h:73
func GodotBasisDeterminant(pSelf []GodotBasis) GodotReal {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_determinant(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotBasisRotated function as declared in gdnative/basis.h:75
func GodotBasisRotated(pSelf []GodotBasis, pAxis []GodotVector3, pPhi GodotReal) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpAxis, _ := unpackArgSGodotVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_basis_rotated(cpSelf, cpAxis, cpPhi)
	packSGodotVector3(pAxis, cpAxis)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisScaled function as declared in gdnative/basis.h:77
func GodotBasisScaled(pSelf []GodotBasis, pScale []GodotVector3) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpScale, _ := unpackArgSGodotVector3(pScale)
	__ret := C.godot_basis_scaled(cpSelf, cpScale)
	packSGodotVector3(pScale, cpScale)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisGetScale function as declared in gdnative/basis.h:79
func GodotBasisGetScale(pSelf []GodotBasis) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_get_scale(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisGetEuler function as declared in gdnative/basis.h:81
func GodotBasisGetEuler(pSelf []GodotBasis) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_get_euler(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisTdotx function as declared in gdnative/basis.h:83
func GodotBasisTdotx(pSelf []GodotBasis, pWith []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpWith, _ := unpackArgSGodotVector3(pWith)
	__ret := C.godot_basis_tdotx(cpSelf, cpWith)
	packSGodotVector3(pWith, cpWith)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotBasisTdoty function as declared in gdnative/basis.h:85
func GodotBasisTdoty(pSelf []GodotBasis, pWith []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpWith, _ := unpackArgSGodotVector3(pWith)
	__ret := C.godot_basis_tdoty(cpSelf, cpWith)
	packSGodotVector3(pWith, cpWith)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotBasisTdotz function as declared in gdnative/basis.h:87
func GodotBasisTdotz(pSelf []GodotBasis, pWith []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpWith, _ := unpackArgSGodotVector3(pWith)
	__ret := C.godot_basis_tdotz(cpSelf, cpWith)
	packSGodotVector3(pWith, cpWith)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotBasisXform function as declared in gdnative/basis.h:89
func GodotBasisXform(pSelf []GodotBasis, pV []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	__ret := C.godot_basis_xform(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisXformInv function as declared in gdnative/basis.h:91
func GodotBasisXformInv(pSelf []GodotBasis, pV []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	__ret := C.godot_basis_xform_inv(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisGetOrthogonalIndex function as declared in gdnative/basis.h:93
func GodotBasisGetOrthogonalIndex(pSelf []GodotBasis) GodotInt {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	__ret := C.godot_basis_get_orthogonal_index(cpSelf)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotBasisNew function as declared in gdnative/basis.h:95
func GodotBasisNew(rDest []GodotBasis) {
	crDest, _ := unpackArgSGodotBasis(rDest)
	C.godot_basis_new(crDest)
	packSGodotBasis(rDest, crDest)
}

// GodotBasisNewWithEulerQuat function as declared in gdnative/basis.h:97
func GodotBasisNewWithEulerQuat(rDest []GodotBasis, pEuler []GodotQuat) {
	crDest, _ := unpackArgSGodotBasis(rDest)
	cpEuler, _ := unpackArgSGodotQuat(pEuler)
	C.godot_basis_new_with_euler_quat(crDest, cpEuler)
	packSGodotQuat(pEuler, cpEuler)
	packSGodotBasis(rDest, crDest)
}

// GodotBasisGetElements function as declared in gdnative/basis.h:100
func GodotBasisGetElements(pSelf []GodotBasis, pElements []GodotVector3) {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpElements, _ := unpackArgSGodotVector3(pElements)
	C.godot_basis_get_elements(cpSelf, cpElements)
	packSGodotVector3(pElements, cpElements)
	packSGodotBasis(pSelf, cpSelf)
}

// GodotBasisGetAxis function as declared in gdnative/basis.h:102
func GodotBasisGetAxis(pSelf []GodotBasis, pAxis GodotInt) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	__ret := C.godot_basis_get_axis(cpSelf, cpAxis)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisSetAxis function as declared in gdnative/basis.h:104
func GodotBasisSetAxis(pSelf []GodotBasis, pAxis GodotInt, pValue []GodotVector3) {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	cpValue, _ := unpackArgSGodotVector3(pValue)
	C.godot_basis_set_axis(cpSelf, cpAxis, cpValue)
	packSGodotVector3(pValue, cpValue)
	packSGodotBasis(pSelf, cpSelf)
}

// GodotBasisGetRow function as declared in gdnative/basis.h:106
func GodotBasisGetRow(pSelf []GodotBasis, pRow GodotInt) GodotVector3 {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpRow, _ := (C.godot_int)(pRow), cgoAllocsUnknown
	__ret := C.godot_basis_get_row(cpSelf, cpRow)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisSetRow function as declared in gdnative/basis.h:108
func GodotBasisSetRow(pSelf []GodotBasis, pRow GodotInt, pValue []GodotVector3) {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpRow, _ := (C.godot_int)(pRow), cgoAllocsUnknown
	cpValue, _ := unpackArgSGodotVector3(pValue)
	C.godot_basis_set_row(cpSelf, cpRow, cpValue)
	packSGodotVector3(pValue, cpValue)
	packSGodotBasis(pSelf, cpSelf)
}

// GodotBasisOperatorEqual function as declared in gdnative/basis.h:110
func GodotBasisOperatorEqual(pSelf []GodotBasis, pB []GodotBasis) GodotBool {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpB, _ := unpackArgSGodotBasis(pB)
	__ret := C.godot_basis_operator_equal(cpSelf, cpB)
	packSGodotBasis(pB, cpB)
	packSGodotBasis(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotBasisOperatorAdd function as declared in gdnative/basis.h:112
func GodotBasisOperatorAdd(pSelf []GodotBasis, pB []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpB, _ := unpackArgSGodotBasis(pB)
	__ret := C.godot_basis_operator_add(cpSelf, cpB)
	packSGodotBasis(pB, cpB)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisOperatorSubstract function as declared in gdnative/basis.h:114
func GodotBasisOperatorSubstract(pSelf []GodotBasis, pB []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpB, _ := unpackArgSGodotBasis(pB)
	__ret := C.godot_basis_operator_substract(cpSelf, cpB)
	packSGodotBasis(pB, cpB)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisOperatorMultiplyVector function as declared in gdnative/basis.h:116
func GodotBasisOperatorMultiplyVector(pSelf []GodotBasis, pB []GodotBasis) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpB, _ := unpackArgSGodotBasis(pB)
	__ret := C.godot_basis_operator_multiply_vector(cpSelf, cpB)
	packSGodotBasis(pB, cpB)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotBasisOperatorMultiplyScalar function as declared in gdnative/basis.h:118
func GodotBasisOperatorMultiplyScalar(pSelf []GodotBasis, pB GodotReal) GodotBasis {
	cpSelf, _ := unpackArgSGodotBasis(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_basis_operator_multiply_scalar(cpSelf, cpB)
	packSGodotBasis(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatNew function as declared in gdnative/quat.h:60
func GodotQuatNew(rDest []GodotQuat, pX GodotReal, pY GodotReal, pZ GodotReal, pW GodotReal) {
	crDest, _ := unpackArgSGodotQuat(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpZ, _ := (C.godot_real)(pZ), cgoAllocsUnknown
	cpW, _ := (C.godot_real)(pW), cgoAllocsUnknown
	C.godot_quat_new(crDest, cpX, cpY, cpZ, cpW)
	packSGodotQuat(rDest, crDest)
}

// GodotQuatNewWithAxisAngle function as declared in gdnative/quat.h:61
func GodotQuatNewWithAxisAngle(rDest []GodotQuat, pAxis []GodotVector3, pAngle GodotReal) {
	crDest, _ := unpackArgSGodotQuat(rDest)
	cpAxis, _ := unpackArgSGodotVector3(pAxis)
	cpAngle, _ := (C.godot_real)(pAngle), cgoAllocsUnknown
	C.godot_quat_new_with_axis_angle(crDest, cpAxis, cpAngle)
	packSGodotVector3(pAxis, cpAxis)
	packSGodotQuat(rDest, crDest)
}

// GodotQuatGetX function as declared in gdnative/quat.h:63
func GodotQuatGetX(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_get_x(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatSetX function as declared in gdnative/quat.h:64
func GodotQuatSetX(pSelf []GodotQuat, val GodotReal) {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_x(cpSelf, cval)
	packSGodotQuat(pSelf, cpSelf)
}

// GodotQuatGetY function as declared in gdnative/quat.h:66
func GodotQuatGetY(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_get_y(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatSetY function as declared in gdnative/quat.h:67
func GodotQuatSetY(pSelf []GodotQuat, val GodotReal) {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_y(cpSelf, cval)
	packSGodotQuat(pSelf, cpSelf)
}

// GodotQuatGetZ function as declared in gdnative/quat.h:69
func GodotQuatGetZ(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_get_z(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatSetZ function as declared in gdnative/quat.h:70
func GodotQuatSetZ(pSelf []GodotQuat, val GodotReal) {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_z(cpSelf, cval)
	packSGodotQuat(pSelf, cpSelf)
}

// GodotQuatGetW function as declared in gdnative/quat.h:72
func GodotQuatGetW(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_get_w(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatSetW function as declared in gdnative/quat.h:73
func GodotQuatSetW(pSelf []GodotQuat, val GodotReal) {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_w(cpSelf, cval)
	packSGodotQuat(pSelf, cpSelf)
}

// GodotQuatAsString function as declared in gdnative/quat.h:75
func GodotQuatAsString(pSelf []GodotQuat) GodotString {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_as_string(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatLength function as declared in gdnative/quat.h:77
func GodotQuatLength(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_length(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatLengthSquared function as declared in gdnative/quat.h:79
func GodotQuatLengthSquared(pSelf []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_length_squared(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatNormalized function as declared in gdnative/quat.h:81
func GodotQuatNormalized(pSelf []GodotQuat) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_normalized(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatIsNormalized function as declared in gdnative/quat.h:83
func GodotQuatIsNormalized(pSelf []GodotQuat) GodotBool {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_is_normalized(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotQuatInverse function as declared in gdnative/quat.h:85
func GodotQuatInverse(pSelf []GodotQuat) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_inverse(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatDot function as declared in gdnative/quat.h:87
func GodotQuatDot(pSelf []GodotQuat, pB []GodotQuat) GodotReal {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	__ret := C.godot_quat_dot(cpSelf, cpB)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotQuatXform function as declared in gdnative/quat.h:89
func GodotQuatXform(pSelf []GodotQuat, pV []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	__ret := C.godot_quat_xform(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatSlerp function as declared in gdnative/quat.h:91
func GodotQuatSlerp(pSelf []GodotQuat, pB []GodotQuat, pT GodotReal) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_slerp(cpSelf, cpB, cpT)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatSlerpni function as declared in gdnative/quat.h:93
func GodotQuatSlerpni(pSelf []GodotQuat, pB []GodotQuat, pT GodotReal) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_slerpni(cpSelf, cpB, cpT)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatCubicSlerp function as declared in gdnative/quat.h:95
func GodotQuatCubicSlerp(pSelf []GodotQuat, pB []GodotQuat, pPreA []GodotQuat, pPostB []GodotQuat, pT GodotReal) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	cpPreA, _ := unpackArgSGodotQuat(pPreA)
	cpPostB, _ := unpackArgSGodotQuat(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_cubic_slerp(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSGodotQuat(pPostB, cpPostB)
	packSGodotQuat(pPreA, cpPreA)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatOperatorMultiply function as declared in gdnative/quat.h:97
func GodotQuatOperatorMultiply(pSelf []GodotQuat, pB GodotReal) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_quat_operator_multiply(cpSelf, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatOperatorAdd function as declared in gdnative/quat.h:99
func GodotQuatOperatorAdd(pSelf []GodotQuat, pB []GodotQuat) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	__ret := C.godot_quat_operator_add(cpSelf, cpB)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatOperatorSubstract function as declared in gdnative/quat.h:101
func GodotQuatOperatorSubstract(pSelf []GodotQuat, pB []GodotQuat) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	__ret := C.godot_quat_operator_substract(cpSelf, cpB)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatOperatorDivide function as declared in gdnative/quat.h:103
func GodotQuatOperatorDivide(pSelf []GodotQuat, pB GodotReal) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_quat_operator_divide(cpSelf, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotQuatOperatorEqual function as declared in gdnative/quat.h:105
func GodotQuatOperatorEqual(pSelf []GodotQuat, pB []GodotQuat) GodotBool {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	cpB, _ := unpackArgSGodotQuat(pB)
	__ret := C.godot_quat_operator_equal(cpSelf, cpB)
	packSGodotQuat(pB, cpB)
	packSGodotQuat(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotQuatOperatorNeg function as declared in gdnative/quat.h:107
func GodotQuatOperatorNeg(pSelf []GodotQuat) GodotQuat {
	cpSelf, _ := unpackArgSGodotQuat(pSelf)
	__ret := C.godot_quat_operator_neg(cpSelf)
	packSGodotQuat(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantGetType function as declared in gdnative/variant.h:131
func GodotVariantGetType(pV []GodotVariant) GodotVariantType {
	cpV, _ := unpackArgSGodotVariant(pV)
	__ret := C.godot_variant_get_type(cpV)
	packSGodotVariant(pV, cpV)
	__v := (GodotVariantType)(__ret)
	return __v
}

// GodotVariantNewCopy function as declared in gdnative/variant.h:133
func GodotVariantNewCopy(rDest []GodotVariant, pSrc []GodotVariant) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpSrc, _ := unpackArgSGodotVariant(pSrc)
	C.godot_variant_new_copy(crDest, cpSrc)
	packSGodotVariant(pSrc, cpSrc)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewNil function as declared in gdnative/variant.h:135
func GodotVariantNewNil(rDest []GodotVariant) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	C.godot_variant_new_nil(crDest)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewBool function as declared in gdnative/variant.h:137
func GodotVariantNewBool(pV []GodotVariant, pB GodotBool) {
	cpV, _ := unpackArgSGodotVariant(pV)
	cpB, _ := (C.godot_bool)(pB), cgoAllocsUnknown
	C.godot_variant_new_bool(cpV, cpB)
	packSGodotVariant(pV, cpV)
}

// GodotVariantNewUint function as declared in gdnative/variant.h:138
func GodotVariantNewUint(rDest []GodotVariant, pI uint64) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpI, _ := (C.uint64_t)(pI), cgoAllocsUnknown
	C.godot_variant_new_uint(crDest, cpI)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewInt function as declared in gdnative/variant.h:139
func GodotVariantNewInt(rDest []GodotVariant, pI int64) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpI, _ := (C.int64_t)(pI), cgoAllocsUnknown
	C.godot_variant_new_int(crDest, cpI)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewReal function as declared in gdnative/variant.h:140
func GodotVariantNewReal(rDest []GodotVariant, pR float64) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpR, _ := (C.double)(pR), cgoAllocsUnknown
	C.godot_variant_new_real(crDest, cpR)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewString function as declared in gdnative/variant.h:141
func GodotVariantNewString(rDest []GodotVariant, pS []GodotString) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpS, _ := unpackArgSGodotString(pS)
	C.godot_variant_new_string(crDest, cpS)
	packSGodotString(pS, cpS)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewVector2 function as declared in gdnative/variant.h:142
func GodotVariantNewVector2(rDest []GodotVariant, pV2 []GodotVector2) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpV2, _ := unpackArgSGodotVector2(pV2)
	C.godot_variant_new_vector2(crDest, cpV2)
	packSGodotVector2(pV2, cpV2)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewRect2 function as declared in gdnative/variant.h:143
func GodotVariantNewRect2(rDest []GodotVariant, pRect2 []GodotRect2) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpRect2, _ := unpackArgSGodotRect2(pRect2)
	C.godot_variant_new_rect2(crDest, cpRect2)
	packSGodotRect2(pRect2, cpRect2)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewVector3 function as declared in gdnative/variant.h:144
func GodotVariantNewVector3(rDest []GodotVariant, pV3 []GodotVector3) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpV3, _ := unpackArgSGodotVector3(pV3)
	C.godot_variant_new_vector3(crDest, cpV3)
	packSGodotVector3(pV3, cpV3)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewTransform2d function as declared in gdnative/variant.h:145
func GodotVariantNewTransform2d(rDest []GodotVariant, pT2d []GodotTransform2d) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpT2d, _ := unpackArgSGodotTransform2d(pT2d)
	C.godot_variant_new_transform2d(crDest, cpT2d)
	packSGodotTransform2d(pT2d, cpT2d)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPlane function as declared in gdnative/variant.h:146
func GodotVariantNewPlane(rDest []GodotVariant, pPlane []GodotPlane) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPlane, _ := unpackArgSGodotPlane(pPlane)
	C.godot_variant_new_plane(crDest, cpPlane)
	packSGodotPlane(pPlane, cpPlane)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewQuat function as declared in gdnative/variant.h:147
func GodotVariantNewQuat(rDest []GodotVariant, pQuat []GodotQuat) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpQuat, _ := unpackArgSGodotQuat(pQuat)
	C.godot_variant_new_quat(crDest, cpQuat)
	packSGodotQuat(pQuat, cpQuat)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewAabb function as declared in gdnative/variant.h:148
func GodotVariantNewAabb(rDest []GodotVariant, pAabb []GodotAabb) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpAabb, _ := unpackArgSGodotAabb(pAabb)
	C.godot_variant_new_aabb(crDest, cpAabb)
	packSGodotAabb(pAabb, cpAabb)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewBasis function as declared in gdnative/variant.h:149
func GodotVariantNewBasis(rDest []GodotVariant, pBasis []GodotBasis) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpBasis, _ := unpackArgSGodotBasis(pBasis)
	C.godot_variant_new_basis(crDest, cpBasis)
	packSGodotBasis(pBasis, cpBasis)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewTransform function as declared in gdnative/variant.h:150
func GodotVariantNewTransform(rDest []GodotVariant, pTrans []GodotTransform) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpTrans, _ := unpackArgSGodotTransform(pTrans)
	C.godot_variant_new_transform(crDest, cpTrans)
	packSGodotTransform(pTrans, cpTrans)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewColor function as declared in gdnative/variant.h:151
func GodotVariantNewColor(rDest []GodotVariant, pColor []GodotColor) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpColor, _ := unpackArgSGodotColor(pColor)
	C.godot_variant_new_color(crDest, cpColor)
	packSGodotColor(pColor, cpColor)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewNodePath function as declared in gdnative/variant.h:152
func GodotVariantNewNodePath(rDest []GodotVariant, pNp []GodotNodePath) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpNp, _ := unpackArgSGodotNodePath(pNp)
	C.godot_variant_new_node_path(crDest, cpNp)
	packSGodotNodePath(pNp, cpNp)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewRid function as declared in gdnative/variant.h:153
func GodotVariantNewRid(rDest []GodotVariant, pRid []GodotRid) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpRid, _ := unpackArgSGodotRid(pRid)
	C.godot_variant_new_rid(crDest, cpRid)
	packSGodotRid(pRid, cpRid)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewObject function as declared in gdnative/variant.h:154
func GodotVariantNewObject(rDest []GodotVariant, pObj *GodotObject) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	C.godot_variant_new_object(crDest, cpObj)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewDictionary function as declared in gdnative/variant.h:155
func GodotVariantNewDictionary(rDest []GodotVariant, pDict []GodotDictionary) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpDict, _ := unpackArgSGodotDictionary(pDict)
	C.godot_variant_new_dictionary(crDest, cpDict)
	packSGodotDictionary(pDict, cpDict)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewArray function as declared in gdnative/variant.h:156
func GodotVariantNewArray(rDest []GodotVariant, pArr []GodotArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpArr, _ := unpackArgSGodotArray(pArr)
	C.godot_variant_new_array(crDest, cpArr)
	packSGodotArray(pArr, cpArr)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolByteArray function as declared in gdnative/variant.h:157
func GodotVariantNewPoolByteArray(rDest []GodotVariant, pPba []GodotPoolByteArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPba, _ := unpackArgSGodotPoolByteArray(pPba)
	C.godot_variant_new_pool_byte_array(crDest, cpPba)
	packSGodotPoolByteArray(pPba, cpPba)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolIntArray function as declared in gdnative/variant.h:158
func GodotVariantNewPoolIntArray(rDest []GodotVariant, pPia []GodotPoolIntArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPia, _ := unpackArgSGodotPoolIntArray(pPia)
	C.godot_variant_new_pool_int_array(crDest, cpPia)
	packSGodotPoolIntArray(pPia, cpPia)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolRealArray function as declared in gdnative/variant.h:159
func GodotVariantNewPoolRealArray(rDest []GodotVariant, pPra []GodotPoolRealArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPra, _ := unpackArgSGodotPoolRealArray(pPra)
	C.godot_variant_new_pool_real_array(crDest, cpPra)
	packSGodotPoolRealArray(pPra, cpPra)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolStringArray function as declared in gdnative/variant.h:160
func GodotVariantNewPoolStringArray(rDest []GodotVariant, pPsa []GodotPoolStringArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPsa, _ := unpackArgSGodotPoolStringArray(pPsa)
	C.godot_variant_new_pool_string_array(crDest, cpPsa)
	packSGodotPoolStringArray(pPsa, cpPsa)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolVector2Array function as declared in gdnative/variant.h:161
func GodotVariantNewPoolVector2Array(rDest []GodotVariant, pPv2a []GodotPoolVector2Array) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPv2a, _ := unpackArgSGodotPoolVector2Array(pPv2a)
	C.godot_variant_new_pool_vector2_array(crDest, cpPv2a)
	packSGodotPoolVector2Array(pPv2a, cpPv2a)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolVector3Array function as declared in gdnative/variant.h:162
func GodotVariantNewPoolVector3Array(rDest []GodotVariant, pPv3a []GodotPoolVector3Array) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPv3a, _ := unpackArgSGodotPoolVector3Array(pPv3a)
	C.godot_variant_new_pool_vector3_array(crDest, cpPv3a)
	packSGodotPoolVector3Array(pPv3a, cpPv3a)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantNewPoolColorArray function as declared in gdnative/variant.h:163
func GodotVariantNewPoolColorArray(rDest []GodotVariant, pPca []GodotPoolColorArray) {
	crDest, _ := unpackArgSGodotVariant(rDest)
	cpPca, _ := unpackArgSGodotPoolColorArray(pPca)
	C.godot_variant_new_pool_color_array(crDest, cpPca)
	packSGodotPoolColorArray(pPca, cpPca)
	packSGodotVariant(rDest, crDest)
}

// GodotVariantAsBool function as declared in gdnative/variant.h:165
func GodotVariantAsBool(pSelf []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_bool(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantAsUint function as declared in gdnative/variant.h:166
func GodotVariantAsUint(pSelf []GodotVariant) uint64 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_uint(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := (uint64)(__ret)
	return __v
}

// GodotVariantAsInt function as declared in gdnative/variant.h:167
func GodotVariantAsInt(pSelf []GodotVariant) int64 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_int(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// GodotVariantAsReal function as declared in gdnative/variant.h:168
func GodotVariantAsReal(pSelf []GodotVariant) float64 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_real(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := (float64)(__ret)
	return __v
}

// GodotVariantAsString function as declared in gdnative/variant.h:169
func GodotVariantAsString(pSelf []GodotVariant) GodotString {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_string(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsVector2 function as declared in gdnative/variant.h:170
func GodotVariantAsVector2(pSelf []GodotVariant) GodotVector2 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_vector2(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsRect2 function as declared in gdnative/variant.h:171
func GodotVariantAsRect2(pSelf []GodotVariant) GodotRect2 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_rect2(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsVector3 function as declared in gdnative/variant.h:172
func GodotVariantAsVector3(pSelf []GodotVariant) GodotVector3 {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_vector3(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsTransform2d function as declared in gdnative/variant.h:173
func GodotVariantAsTransform2d(pSelf []GodotVariant) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_transform2d(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPlane function as declared in gdnative/variant.h:174
func GodotVariantAsPlane(pSelf []GodotVariant) GodotPlane {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_plane(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsQuat function as declared in gdnative/variant.h:175
func GodotVariantAsQuat(pSelf []GodotVariant) GodotQuat {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_quat(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsAabb function as declared in gdnative/variant.h:176
func GodotVariantAsAabb(pSelf []GodotVariant) GodotAabb {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_aabb(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsBasis function as declared in gdnative/variant.h:177
func GodotVariantAsBasis(pSelf []GodotVariant) GodotBasis {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_basis(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsTransform function as declared in gdnative/variant.h:178
func GodotVariantAsTransform(pSelf []GodotVariant) GodotTransform {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_transform(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsColor function as declared in gdnative/variant.h:179
func GodotVariantAsColor(pSelf []GodotVariant) GodotColor {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_color(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotColorRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsNodePath function as declared in gdnative/variant.h:180
func GodotVariantAsNodePath(pSelf []GodotVariant) GodotNodePath {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_node_path(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotNodePathRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsRid function as declared in gdnative/variant.h:181
func GodotVariantAsRid(pSelf []GodotVariant) GodotRid {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_rid(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotRidRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsObject function as declared in gdnative/variant.h:182
func GodotVariantAsObject(pSelf []GodotVariant) *GodotObject {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_object(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *(**GodotObject)(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsDictionary function as declared in gdnative/variant.h:183
func GodotVariantAsDictionary(pSelf []GodotVariant) GodotDictionary {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_dictionary(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotDictionaryRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsArray function as declared in gdnative/variant.h:184
func GodotVariantAsArray(pSelf []GodotVariant) GodotArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolByteArray function as declared in gdnative/variant.h:185
func GodotVariantAsPoolByteArray(pSelf []GodotVariant) GodotPoolByteArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_byte_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolIntArray function as declared in gdnative/variant.h:186
func GodotVariantAsPoolIntArray(pSelf []GodotVariant) GodotPoolIntArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_int_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolIntArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolRealArray function as declared in gdnative/variant.h:187
func GodotVariantAsPoolRealArray(pSelf []GodotVariant) GodotPoolRealArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_real_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolRealArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolStringArray function as declared in gdnative/variant.h:188
func GodotVariantAsPoolStringArray(pSelf []GodotVariant) GodotPoolStringArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_string_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolStringArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolVector2Array function as declared in gdnative/variant.h:189
func GodotVariantAsPoolVector2Array(pSelf []GodotVariant) GodotPoolVector2Array {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_vector2_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolVector2ArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolVector3Array function as declared in gdnative/variant.h:190
func GodotVariantAsPoolVector3Array(pSelf []GodotVariant) GodotPoolVector3Array {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_vector3_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolVector3ArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantAsPoolColorArray function as declared in gdnative/variant.h:191
func GodotVariantAsPoolColorArray(pSelf []GodotVariant) GodotPoolColorArray {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_as_pool_color_array(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotPoolColorArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantCall function as declared in gdnative/variant.h:193
func GodotVariantCall(pSelf []GodotVariant, pMethod []GodotString, pArgs [][]GodotVariant, pArgcount GodotInt, rError []GodotVariantCallError) GodotVariant {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	cpMethod, _ := unpackArgSGodotString(pMethod)
	cpArgs, _ := unpackArgSSGodotVariant(pArgs)
	cpArgcount, _ := (C.godot_int)(pArgcount), cgoAllocsUnknown
	crError, _ := unpackArgSGodotVariantCallError(rError)
	__ret := C.godot_variant_call(cpSelf, cpMethod, cpArgs, cpArgcount, crError)
	packSGodotVariantCallError(rError, crError)
	packSSGodotVariant(pArgs, cpArgs)
	packSGodotString(pMethod, cpMethod)
	packSGodotVariant(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotVariantHasMethod function as declared in gdnative/variant.h:195
func GodotVariantHasMethod(pSelf []GodotVariant, pMethod []GodotString) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	cpMethod, _ := unpackArgSGodotString(pMethod)
	__ret := C.godot_variant_has_method(cpSelf, cpMethod)
	packSGodotString(pMethod, cpMethod)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantOperatorEqual function as declared in gdnative/variant.h:197
func GodotVariantOperatorEqual(pSelf []GodotVariant, pOther []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	cpOther, _ := unpackArgSGodotVariant(pOther)
	__ret := C.godot_variant_operator_equal(cpSelf, cpOther)
	packSGodotVariant(pOther, cpOther)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantOperatorLess function as declared in gdnative/variant.h:198
func GodotVariantOperatorLess(pSelf []GodotVariant, pOther []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	cpOther, _ := unpackArgSGodotVariant(pOther)
	__ret := C.godot_variant_operator_less(cpSelf, cpOther)
	packSGodotVariant(pOther, cpOther)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantHashCompare function as declared in gdnative/variant.h:200
func GodotVariantHashCompare(pSelf []GodotVariant, pOther []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	cpOther, _ := unpackArgSGodotVariant(pOther)
	__ret := C.godot_variant_hash_compare(cpSelf, cpOther)
	packSGodotVariant(pOther, cpOther)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantBooleanize function as declared in gdnative/variant.h:202
func GodotVariantBooleanize(pSelf []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	__ret := C.godot_variant_booleanize(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotVariantDestroy function as declared in gdnative/variant.h:204
func GodotVariantDestroy(pSelf []GodotVariant) {
	cpSelf, _ := unpackArgSGodotVariant(pSelf)
	C.godot_variant_destroy(cpSelf)
	packSGodotVariant(pSelf, cpSelf)
}

// GodotAabbNew function as declared in gdnative/aabb.h:61
func GodotAabbNew(rDest []GodotAabb, pPos []GodotVector3, pSize []GodotVector3) {
	crDest, _ := unpackArgSGodotAabb(rDest)
	cpPos, _ := unpackArgSGodotVector3(pPos)
	cpSize, _ := unpackArgSGodotVector3(pSize)
	C.godot_aabb_new(crDest, cpPos, cpSize)
	packSGodotVector3(pSize, cpSize)
	packSGodotVector3(pPos, cpPos)
	packSGodotAabb(rDest, crDest)
}

// GodotAabbGetPosition function as declared in gdnative/aabb.h:63
func GodotAabbGetPosition(pSelf []GodotAabb) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_position(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbSetPosition function as declared in gdnative/aabb.h:64
func GodotAabbSetPosition(pSelf []GodotAabb, pV []GodotVector3) {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	C.godot_aabb_set_position(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotAabb(pSelf, cpSelf)
}

// GodotAabbGetSize function as declared in gdnative/aabb.h:66
func GodotAabbGetSize(pSelf []GodotAabb) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_size(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbSetSize function as declared in gdnative/aabb.h:67
func GodotAabbSetSize(pSelf []GodotAabb, pV []GodotVector3) {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	C.godot_aabb_set_size(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotAabb(pSelf, cpSelf)
}

// GodotAabbAsString function as declared in gdnative/aabb.h:69
func GodotAabbAsString(pSelf []GodotAabb) GodotString {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_as_string(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGetArea function as declared in gdnative/aabb.h:71
func GodotAabbGetArea(pSelf []GodotAabb) GodotReal {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_area(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotAabbHasNoArea function as declared in gdnative/aabb.h:73
func GodotAabbHasNoArea(pSelf []GodotAabb) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_has_no_area(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbHasNoSurface function as declared in gdnative/aabb.h:75
func GodotAabbHasNoSurface(pSelf []GodotAabb) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_has_no_surface(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbIntersects function as declared in gdnative/aabb.h:77
func GodotAabbIntersects(pSelf []GodotAabb, pWith []GodotAabb) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpWith, _ := unpackArgSGodotAabb(pWith)
	__ret := C.godot_aabb_intersects(cpSelf, cpWith)
	packSGodotAabb(pWith, cpWith)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbEncloses function as declared in gdnative/aabb.h:79
func GodotAabbEncloses(pSelf []GodotAabb, pWith []GodotAabb) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpWith, _ := unpackArgSGodotAabb(pWith)
	__ret := C.godot_aabb_encloses(cpSelf, cpWith)
	packSGodotAabb(pWith, cpWith)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbMerge function as declared in gdnative/aabb.h:81
func GodotAabbMerge(pSelf []GodotAabb, pWith []GodotAabb) GodotAabb {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpWith, _ := unpackArgSGodotAabb(pWith)
	__ret := C.godot_aabb_merge(cpSelf, cpWith)
	packSGodotAabb(pWith, cpWith)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbIntersection function as declared in gdnative/aabb.h:83
func GodotAabbIntersection(pSelf []GodotAabb, pWith []GodotAabb) GodotAabb {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpWith, _ := unpackArgSGodotAabb(pWith)
	__ret := C.godot_aabb_intersection(cpSelf, cpWith)
	packSGodotAabb(pWith, cpWith)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbIntersectsPlane function as declared in gdnative/aabb.h:85
func GodotAabbIntersectsPlane(pSelf []GodotAabb, pPlane []GodotPlane) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpPlane, _ := unpackArgSGodotPlane(pPlane)
	__ret := C.godot_aabb_intersects_plane(cpSelf, cpPlane)
	packSGodotPlane(pPlane, cpPlane)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbIntersectsSegment function as declared in gdnative/aabb.h:87
func GodotAabbIntersectsSegment(pSelf []GodotAabb, pFrom []GodotVector3, pTo []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpFrom, _ := unpackArgSGodotVector3(pFrom)
	cpTo, _ := unpackArgSGodotVector3(pTo)
	__ret := C.godot_aabb_intersects_segment(cpSelf, cpFrom, cpTo)
	packSGodotVector3(pTo, cpTo)
	packSGodotVector3(pFrom, cpFrom)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbHasPoint function as declared in gdnative/aabb.h:89
func GodotAabbHasPoint(pSelf []GodotAabb, pPoint []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpPoint, _ := unpackArgSGodotVector3(pPoint)
	__ret := C.godot_aabb_has_point(cpSelf, cpPoint)
	packSGodotVector3(pPoint, cpPoint)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotAabbGetSupport function as declared in gdnative/aabb.h:91
func GodotAabbGetSupport(pSelf []GodotAabb, pDir []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpDir, _ := unpackArgSGodotVector3(pDir)
	__ret := C.godot_aabb_get_support(cpSelf, cpDir)
	packSGodotVector3(pDir, cpDir)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGetLongestAxis function as declared in gdnative/aabb.h:93
func GodotAabbGetLongestAxis(pSelf []GodotAabb) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGetLongestAxisIndex function as declared in gdnative/aabb.h:95
func GodotAabbGetLongestAxisIndex(pSelf []GodotAabb) GodotInt {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis_index(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotAabbGetLongestAxisSize function as declared in gdnative/aabb.h:97
func GodotAabbGetLongestAxisSize(pSelf []GodotAabb) GodotReal {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis_size(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotAabbGetShortestAxis function as declared in gdnative/aabb.h:99
func GodotAabbGetShortestAxis(pSelf []GodotAabb) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGetShortestAxisIndex function as declared in gdnative/aabb.h:101
func GodotAabbGetShortestAxisIndex(pSelf []GodotAabb) GodotInt {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis_index(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotAabbGetShortestAxisSize function as declared in gdnative/aabb.h:103
func GodotAabbGetShortestAxisSize(pSelf []GodotAabb) GodotReal {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis_size(cpSelf)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotAabbExpand function as declared in gdnative/aabb.h:105
func GodotAabbExpand(pSelf []GodotAabb, pToPoint []GodotVector3) GodotAabb {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpToPoint, _ := unpackArgSGodotVector3(pToPoint)
	__ret := C.godot_aabb_expand(cpSelf, cpToPoint)
	packSGodotVector3(pToPoint, cpToPoint)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGrow function as declared in gdnative/aabb.h:107
func GodotAabbGrow(pSelf []GodotAabb, pBy GodotReal) GodotAabb {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpBy, _ := (C.godot_real)(pBy), cgoAllocsUnknown
	__ret := C.godot_aabb_grow(cpSelf, cpBy)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbGetEndpoint function as declared in gdnative/aabb.h:109
func GodotAabbGetEndpoint(pSelf []GodotAabb, pIdx GodotInt) GodotVector3 {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_aabb_get_endpoint(cpSelf, cpIdx)
	packSGodotAabb(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotAabbOperatorEqual function as declared in gdnative/aabb.h:111
func GodotAabbOperatorEqual(pSelf []GodotAabb, pB []GodotAabb) GodotBool {
	cpSelf, _ := unpackArgSGodotAabb(pSelf)
	cpB, _ := unpackArgSGodotAabb(pB)
	__ret := C.godot_aabb_operator_equal(cpSelf, cpB)
	packSGodotAabb(pB, cpB)
	packSGodotAabb(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneNewWithReals function as declared in gdnative/plane.h:60
func GodotPlaneNewWithReals(rDest []GodotPlane, pA GodotReal, pB GodotReal, pC GodotReal, pD GodotReal) {
	crDest, _ := unpackArgSGodotPlane(rDest)
	cpA, _ := (C.godot_real)(pA), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	cpC, _ := (C.godot_real)(pC), cgoAllocsUnknown
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_new_with_reals(crDest, cpA, cpB, cpC, cpD)
	packSGodotPlane(rDest, crDest)
}

// GodotPlaneNewWithVectors function as declared in gdnative/plane.h:61
func GodotPlaneNewWithVectors(rDest []GodotPlane, pV1 []GodotVector3, pV2 []GodotVector3, pV3 []GodotVector3) {
	crDest, _ := unpackArgSGodotPlane(rDest)
	cpV1, _ := unpackArgSGodotVector3(pV1)
	cpV2, _ := unpackArgSGodotVector3(pV2)
	cpV3, _ := unpackArgSGodotVector3(pV3)
	C.godot_plane_new_with_vectors(crDest, cpV1, cpV2, cpV3)
	packSGodotVector3(pV3, cpV3)
	packSGodotVector3(pV2, cpV2)
	packSGodotVector3(pV1, cpV1)
	packSGodotPlane(rDest, crDest)
}

// GodotPlaneNewWithNormal function as declared in gdnative/plane.h:62
func GodotPlaneNewWithNormal(rDest []GodotPlane, pNormal []GodotVector3, pD GodotReal) {
	crDest, _ := unpackArgSGodotPlane(rDest)
	cpNormal, _ := unpackArgSGodotVector3(pNormal)
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_new_with_normal(crDest, cpNormal, cpD)
	packSGodotVector3(pNormal, cpNormal)
	packSGodotPlane(rDest, crDest)
}

// GodotPlaneAsString function as declared in gdnative/plane.h:64
func GodotPlaneAsString(pSelf []GodotPlane) GodotString {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_as_string(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneNormalized function as declared in gdnative/plane.h:66
func GodotPlaneNormalized(pSelf []GodotPlane) GodotPlane {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_normalized(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneCenter function as declared in gdnative/plane.h:68
func GodotPlaneCenter(pSelf []GodotPlane) GodotVector3 {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_center(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneGetAnyPoint function as declared in gdnative/plane.h:70
func GodotPlaneGetAnyPoint(pSelf []GodotPlane) GodotVector3 {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_get_any_point(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneIsPointOver function as declared in gdnative/plane.h:72
func GodotPlaneIsPointOver(pSelf []GodotPlane, pPoint []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpPoint, _ := unpackArgSGodotVector3(pPoint)
	__ret := C.godot_plane_is_point_over(cpSelf, cpPoint)
	packSGodotVector3(pPoint, cpPoint)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneDistanceTo function as declared in gdnative/plane.h:74
func GodotPlaneDistanceTo(pSelf []GodotPlane, pPoint []GodotVector3) GodotReal {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpPoint, _ := unpackArgSGodotVector3(pPoint)
	__ret := C.godot_plane_distance_to(cpSelf, cpPoint)
	packSGodotVector3(pPoint, cpPoint)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotPlaneHasPoint function as declared in gdnative/plane.h:76
func GodotPlaneHasPoint(pSelf []GodotPlane, pPoint []GodotVector3, pEpsilon GodotReal) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpPoint, _ := unpackArgSGodotVector3(pPoint)
	cpEpsilon, _ := (C.godot_real)(pEpsilon), cgoAllocsUnknown
	__ret := C.godot_plane_has_point(cpSelf, cpPoint, cpEpsilon)
	packSGodotVector3(pPoint, cpPoint)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneProject function as declared in gdnative/plane.h:78
func GodotPlaneProject(pSelf []GodotPlane, pPoint []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpPoint, _ := unpackArgSGodotVector3(pPoint)
	__ret := C.godot_plane_project(cpSelf, cpPoint)
	packSGodotVector3(pPoint, cpPoint)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneIntersect3 function as declared in gdnative/plane.h:80
func GodotPlaneIntersect3(pSelf []GodotPlane, rDest []GodotVector3, pB []GodotPlane, pC []GodotPlane) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	crDest, _ := unpackArgSGodotVector3(rDest)
	cpB, _ := unpackArgSGodotPlane(pB)
	cpC, _ := unpackArgSGodotPlane(pC)
	__ret := C.godot_plane_intersect_3(cpSelf, crDest, cpB, cpC)
	packSGodotPlane(pC, cpC)
	packSGodotPlane(pB, cpB)
	packSGodotVector3(rDest, crDest)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneIntersectsRay function as declared in gdnative/plane.h:82
func GodotPlaneIntersectsRay(pSelf []GodotPlane, rDest []GodotVector3, pFrom []GodotVector3, pDir []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	crDest, _ := unpackArgSGodotVector3(rDest)
	cpFrom, _ := unpackArgSGodotVector3(pFrom)
	cpDir, _ := unpackArgSGodotVector3(pDir)
	__ret := C.godot_plane_intersects_ray(cpSelf, crDest, cpFrom, cpDir)
	packSGodotVector3(pDir, cpDir)
	packSGodotVector3(pFrom, cpFrom)
	packSGodotVector3(rDest, crDest)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneIntersectsSegment function as declared in gdnative/plane.h:84
func GodotPlaneIntersectsSegment(pSelf []GodotPlane, rDest []GodotVector3, pBegin []GodotVector3, pEnd []GodotVector3) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	crDest, _ := unpackArgSGodotVector3(rDest)
	cpBegin, _ := unpackArgSGodotVector3(pBegin)
	cpEnd, _ := unpackArgSGodotVector3(pEnd)
	__ret := C.godot_plane_intersects_segment(cpSelf, crDest, cpBegin, cpEnd)
	packSGodotVector3(pEnd, cpEnd)
	packSGodotVector3(pBegin, cpBegin)
	packSGodotVector3(rDest, crDest)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneOperatorNeg function as declared in gdnative/plane.h:86
func GodotPlaneOperatorNeg(pSelf []GodotPlane) GodotPlane {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_operator_neg(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneOperatorEqual function as declared in gdnative/plane.h:88
func GodotPlaneOperatorEqual(pSelf []GodotPlane, pB []GodotPlane) GodotBool {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpB, _ := unpackArgSGodotPlane(pB)
	__ret := C.godot_plane_operator_equal(cpSelf, cpB)
	packSGodotPlane(pB, cpB)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotPlaneSetNormal function as declared in gdnative/plane.h:90
func GodotPlaneSetNormal(pSelf []GodotPlane, pNormal []GodotVector3) {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpNormal, _ := unpackArgSGodotVector3(pNormal)
	C.godot_plane_set_normal(cpSelf, cpNormal)
	packSGodotVector3(pNormal, cpNormal)
	packSGodotPlane(pSelf, cpSelf)
}

// GodotPlaneGetNormal function as declared in gdnative/plane.h:92
func GodotPlaneGetNormal(pSelf []GodotPlane) GodotVector3 {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_get_normal(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotPlaneGetD function as declared in gdnative/plane.h:94
func GodotPlaneGetD(pSelf []GodotPlane) GodotReal {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	__ret := C.godot_plane_get_d(cpSelf)
	packSGodotPlane(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotPlaneSetD function as declared in gdnative/plane.h:96
func GodotPlaneSetD(pSelf []GodotPlane, pD GodotReal) {
	cpSelf, _ := unpackArgSGodotPlane(pSelf)
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_set_d(cpSelf, cpD)
	packSGodotPlane(pSelf, cpSelf)
}

// GodotDictionaryNew function as declared in gdnative/dictionary.h:61
func GodotDictionaryNew(rDest []GodotDictionary) {
	crDest, _ := unpackArgSGodotDictionary(rDest)
	C.godot_dictionary_new(crDest)
	packSGodotDictionary(rDest, crDest)
}

// GodotDictionaryNewCopy function as declared in gdnative/dictionary.h:62
func GodotDictionaryNewCopy(rDest []GodotDictionary, pSrc []GodotDictionary) {
	crDest, _ := unpackArgSGodotDictionary(rDest)
	cpSrc, _ := unpackArgSGodotDictionary(pSrc)
	C.godot_dictionary_new_copy(crDest, cpSrc)
	packSGodotDictionary(pSrc, cpSrc)
	packSGodotDictionary(rDest, crDest)
}

// GodotDictionaryDestroy function as declared in gdnative/dictionary.h:63
func GodotDictionaryDestroy(pSelf []GodotDictionary) {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	C.godot_dictionary_destroy(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
}

// GodotDictionarySize function as declared in gdnative/dictionary.h:65
func GodotDictionarySize(pSelf []GodotDictionary) GodotInt {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_size(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotDictionaryEmpty function as declared in gdnative/dictionary.h:67
func GodotDictionaryEmpty(pSelf []GodotDictionary) GodotBool {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_empty(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotDictionaryClear function as declared in gdnative/dictionary.h:69
func GodotDictionaryClear(pSelf []GodotDictionary) {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	C.godot_dictionary_clear(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
}

// GodotDictionaryHas function as declared in gdnative/dictionary.h:71
func GodotDictionaryHas(pSelf []GodotDictionary, pKey []GodotVariant) GodotBool {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	__ret := C.godot_dictionary_has(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotDictionaryHasAll function as declared in gdnative/dictionary.h:73
func GodotDictionaryHasAll(pSelf []GodotDictionary, pKeys []GodotArray) GodotBool {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKeys, _ := unpackArgSGodotArray(pKeys)
	__ret := C.godot_dictionary_has_all(cpSelf, cpKeys)
	packSGodotArray(pKeys, cpKeys)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotDictionaryErase function as declared in gdnative/dictionary.h:75
func GodotDictionaryErase(pSelf []GodotDictionary, pKey []GodotVariant) {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	C.godot_dictionary_erase(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
}

// GodotDictionaryHash function as declared in gdnative/dictionary.h:77
func GodotDictionaryHash(pSelf []GodotDictionary) GodotInt {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_hash(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotDictionaryKeys function as declared in gdnative/dictionary.h:79
func GodotDictionaryKeys(pSelf []GodotDictionary) GodotArray {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_keys(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotDictionaryValues function as declared in gdnative/dictionary.h:81
func GodotDictionaryValues(pSelf []GodotDictionary) GodotArray {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_values(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := *NewGodotArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotDictionaryGet function as declared in gdnative/dictionary.h:83
func GodotDictionaryGet(pSelf []GodotDictionary, pKey []GodotVariant) GodotVariant {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	__ret := C.godot_dictionary_get(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
	__v := *NewGodotVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotDictionarySet function as declared in gdnative/dictionary.h:84
func GodotDictionarySet(pSelf []GodotDictionary, pKey []GodotVariant, pValue []GodotVariant) {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	cpValue, _ := unpackArgSGodotVariant(pValue)
	C.godot_dictionary_set(cpSelf, cpKey, cpValue)
	packSGodotVariant(pValue, cpValue)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
}

// GodotDictionaryOperatorIndex function as declared in gdnative/dictionary.h:86
func GodotDictionaryOperatorIndex(pSelf []GodotDictionary, pKey []GodotVariant) *GodotVariant {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	__ret := C.godot_dictionary_operator_index(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
	__v := NewGodotVariantRef(unsafe.Pointer(__ret))
	return __v
}

// GodotDictionaryOperatorIndexConst function as declared in gdnative/dictionary.h:88
func GodotDictionaryOperatorIndexConst(pSelf []GodotDictionary, pKey []GodotVariant) *GodotVariant {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	__ret := C.godot_dictionary_operator_index_const(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
	__v := NewGodotVariantRef(unsafe.Pointer(__ret))
	return __v
}

// GodotDictionaryNext function as declared in gdnative/dictionary.h:90
func GodotDictionaryNext(pSelf []GodotDictionary, pKey []GodotVariant) *GodotVariant {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpKey, _ := unpackArgSGodotVariant(pKey)
	__ret := C.godot_dictionary_next(cpSelf, cpKey)
	packSGodotVariant(pKey, cpKey)
	packSGodotDictionary(pSelf, cpSelf)
	__v := NewGodotVariantRef(unsafe.Pointer(__ret))
	return __v
}

// GodotDictionaryOperatorEqual function as declared in gdnative/dictionary.h:92
func GodotDictionaryOperatorEqual(pSelf []GodotDictionary, pB []GodotDictionary) GodotBool {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	cpB, _ := unpackArgSGodotDictionary(pB)
	__ret := C.godot_dictionary_operator_equal(cpSelf, cpB)
	packSGodotDictionary(pB, cpB)
	packSGodotDictionary(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotDictionaryToJson function as declared in gdnative/dictionary.h:94
func GodotDictionaryToJson(pSelf []GodotDictionary) GodotString {
	cpSelf, _ := unpackArgSGodotDictionary(pSelf)
	__ret := C.godot_dictionary_to_json(cpSelf)
	packSGodotDictionary(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotNodePathNew function as declared in gdnative/node_path.h:60
func GodotNodePathNew(rDest []GodotNodePath, pFrom []GodotString) {
	crDest, _ := unpackArgSGodotNodePath(rDest)
	cpFrom, _ := unpackArgSGodotString(pFrom)
	C.godot_node_path_new(crDest, cpFrom)
	packSGodotString(pFrom, cpFrom)
	packSGodotNodePath(rDest, crDest)
}

// GodotNodePathNewCopy function as declared in gdnative/node_path.h:61
func GodotNodePathNewCopy(rDest []GodotNodePath, pSrc []GodotNodePath) {
	crDest, _ := unpackArgSGodotNodePath(rDest)
	cpSrc, _ := unpackArgSGodotNodePath(pSrc)
	C.godot_node_path_new_copy(crDest, cpSrc)
	packSGodotNodePath(pSrc, cpSrc)
	packSGodotNodePath(rDest, crDest)
}

// GodotNodePathDestroy function as declared in gdnative/node_path.h:62
func GodotNodePathDestroy(pSelf []GodotNodePath) {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	C.godot_node_path_destroy(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
}

// GodotNodePathAsString function as declared in gdnative/node_path.h:64
func GodotNodePathAsString(pSelf []GodotNodePath) GodotString {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_as_string(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotNodePathIsAbsolute function as declared in gdnative/node_path.h:66
func GodotNodePathIsAbsolute(pSelf []GodotNodePath) GodotBool {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_is_absolute(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotNodePathGetNameCount function as declared in gdnative/node_path.h:68
func GodotNodePathGetNameCount(pSelf []GodotNodePath) GodotInt {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_get_name_count(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotNodePathGetName function as declared in gdnative/node_path.h:70
func GodotNodePathGetName(pSelf []GodotNodePath, pIdx GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_node_path_get_name(cpSelf, cpIdx)
	packSGodotNodePath(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotNodePathGetSubnameCount function as declared in gdnative/node_path.h:72
func GodotNodePathGetSubnameCount(pSelf []GodotNodePath) GodotInt {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_get_subname_count(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotNodePathGetSubname function as declared in gdnative/node_path.h:74
func GodotNodePathGetSubname(pSelf []GodotNodePath, pIdx GodotInt) GodotString {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_node_path_get_subname(cpSelf, cpIdx)
	packSGodotNodePath(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotNodePathGetConcatenatedSubnames function as declared in gdnative/node_path.h:76
func GodotNodePathGetConcatenatedSubnames(pSelf []GodotNodePath) GodotString {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_get_concatenated_subnames(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotNodePathIsEmpty function as declared in gdnative/node_path.h:78
func GodotNodePathIsEmpty(pSelf []GodotNodePath) GodotBool {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	__ret := C.godot_node_path_is_empty(cpSelf)
	packSGodotNodePath(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotNodePathOperatorEqual function as declared in gdnative/node_path.h:80
func GodotNodePathOperatorEqual(pSelf []GodotNodePath, pB []GodotNodePath) GodotBool {
	cpSelf, _ := unpackArgSGodotNodePath(pSelf)
	cpB, _ := unpackArgSGodotNodePath(pB)
	__ret := C.godot_node_path_operator_equal(cpSelf, cpB)
	packSGodotNodePath(pB, cpB)
	packSGodotNodePath(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2NewWithPositionAndSize function as declared in gdnative/rect2.h:58
func GodotRect2NewWithPositionAndSize(rDest []GodotRect2, pPos []GodotVector2, pSize []GodotVector2) {
	crDest, _ := unpackArgSGodotRect2(rDest)
	cpPos, _ := unpackArgSGodotVector2(pPos)
	cpSize, _ := unpackArgSGodotVector2(pSize)
	C.godot_rect2_new_with_position_and_size(crDest, cpPos, cpSize)
	packSGodotVector2(pSize, cpSize)
	packSGodotVector2(pPos, cpPos)
	packSGodotRect2(rDest, crDest)
}

// GodotRect2New function as declared in gdnative/rect2.h:59
func GodotRect2New(rDest []GodotRect2, pX GodotReal, pY GodotReal, pWidth GodotReal, pHeight GodotReal) {
	crDest, _ := unpackArgSGodotRect2(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpWidth, _ := (C.godot_real)(pWidth), cgoAllocsUnknown
	cpHeight, _ := (C.godot_real)(pHeight), cgoAllocsUnknown
	C.godot_rect2_new(crDest, cpX, cpY, cpWidth, cpHeight)
	packSGodotRect2(rDest, crDest)
}

// GodotRect2AsString function as declared in gdnative/rect2.h:61
func GodotRect2AsString(pSelf []GodotRect2) GodotString {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	__ret := C.godot_rect2_as_string(cpSelf)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2GetArea function as declared in gdnative/rect2.h:63
func GodotRect2GetArea(pSelf []GodotRect2) GodotReal {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	__ret := C.godot_rect2_get_area(cpSelf)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotRect2Intersects function as declared in gdnative/rect2.h:65
func GodotRect2Intersects(pSelf []GodotRect2, pB []GodotRect2) GodotBool {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpB, _ := unpackArgSGodotRect2(pB)
	__ret := C.godot_rect2_intersects(cpSelf, cpB)
	packSGodotRect2(pB, cpB)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2Encloses function as declared in gdnative/rect2.h:67
func GodotRect2Encloses(pSelf []GodotRect2, pB []GodotRect2) GodotBool {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpB, _ := unpackArgSGodotRect2(pB)
	__ret := C.godot_rect2_encloses(cpSelf, cpB)
	packSGodotRect2(pB, cpB)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2HasNoArea function as declared in gdnative/rect2.h:69
func GodotRect2HasNoArea(pSelf []GodotRect2) GodotBool {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	__ret := C.godot_rect2_has_no_area(cpSelf)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2Clip function as declared in gdnative/rect2.h:71
func GodotRect2Clip(pSelf []GodotRect2, pB []GodotRect2) GodotRect2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpB, _ := unpackArgSGodotRect2(pB)
	__ret := C.godot_rect2_clip(cpSelf, cpB)
	packSGodotRect2(pB, cpB)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2Merge function as declared in gdnative/rect2.h:73
func GodotRect2Merge(pSelf []GodotRect2, pB []GodotRect2) GodotRect2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpB, _ := unpackArgSGodotRect2(pB)
	__ret := C.godot_rect2_merge(cpSelf, cpB)
	packSGodotRect2(pB, cpB)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2HasPoint function as declared in gdnative/rect2.h:75
func GodotRect2HasPoint(pSelf []GodotRect2, pPoint []GodotVector2) GodotBool {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpPoint, _ := unpackArgSGodotVector2(pPoint)
	__ret := C.godot_rect2_has_point(cpSelf, cpPoint)
	packSGodotVector2(pPoint, cpPoint)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2Grow function as declared in gdnative/rect2.h:77
func GodotRect2Grow(pSelf []GodotRect2, pBy GodotReal) GodotRect2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpBy, _ := (C.godot_real)(pBy), cgoAllocsUnknown
	__ret := C.godot_rect2_grow(cpSelf, cpBy)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2Expand function as declared in gdnative/rect2.h:79
func GodotRect2Expand(pSelf []GodotRect2, pTo []GodotVector2) GodotRect2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpTo, _ := unpackArgSGodotVector2(pTo)
	__ret := C.godot_rect2_expand(cpSelf, cpTo)
	packSGodotVector2(pTo, cpTo)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2OperatorEqual function as declared in gdnative/rect2.h:81
func GodotRect2OperatorEqual(pSelf []GodotRect2, pB []GodotRect2) GodotBool {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpB, _ := unpackArgSGodotRect2(pB)
	__ret := C.godot_rect2_operator_equal(cpSelf, cpB)
	packSGodotRect2(pB, cpB)
	packSGodotRect2(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRect2GetPosition function as declared in gdnative/rect2.h:83
func GodotRect2GetPosition(pSelf []GodotRect2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	__ret := C.godot_rect2_get_position(cpSelf)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2GetSize function as declared in gdnative/rect2.h:85
func GodotRect2GetSize(pSelf []GodotRect2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	__ret := C.godot_rect2_get_size(cpSelf)
	packSGodotRect2(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotRect2SetPosition function as declared in gdnative/rect2.h:87
func GodotRect2SetPosition(pSelf []GodotRect2, pPos []GodotVector2) {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpPos, _ := unpackArgSGodotVector2(pPos)
	C.godot_rect2_set_position(cpSelf, cpPos)
	packSGodotVector2(pPos, cpPos)
	packSGodotRect2(pSelf, cpSelf)
}

// GodotRect2SetSize function as declared in gdnative/rect2.h:89
func GodotRect2SetSize(pSelf []GodotRect2, pSize []GodotVector2) {
	cpSelf, _ := unpackArgSGodotRect2(pSelf)
	cpSize, _ := unpackArgSGodotVector2(pSize)
	C.godot_rect2_set_size(cpSelf, cpSize)
	packSGodotVector2(pSize, cpSize)
	packSGodotRect2(pSelf, cpSelf)
}

// GodotRidNew function as declared in gdnative/rid.h:59
func GodotRidNew(rDest []GodotRid) {
	crDest, _ := unpackArgSGodotRid(rDest)
	C.godot_rid_new(crDest)
	packSGodotRid(rDest, crDest)
}

// GodotRidGetId function as declared in gdnative/rid.h:61
func GodotRidGetId(pSelf []GodotRid) GodotInt {
	cpSelf, _ := unpackArgSGodotRid(pSelf)
	__ret := C.godot_rid_get_id(cpSelf)
	packSGodotRid(pSelf, cpSelf)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotRidNewWithResource function as declared in gdnative/rid.h:63
func GodotRidNewWithResource(rDest []GodotRid, pFrom *GodotObject) {
	crDest, _ := unpackArgSGodotRid(rDest)
	cpFrom, _ := (C.godot_object)(unsafe.Pointer(pFrom)), cgoAllocsUnknown
	C.godot_rid_new_with_resource(crDest, cpFrom)
	packSGodotRid(rDest, crDest)
}

// GodotRidOperatorEqual function as declared in gdnative/rid.h:65
func GodotRidOperatorEqual(pSelf []GodotRid, pB []GodotRid) GodotBool {
	cpSelf, _ := unpackArgSGodotRid(pSelf)
	cpB, _ := unpackArgSGodotRid(pB)
	__ret := C.godot_rid_operator_equal(cpSelf, cpB)
	packSGodotRid(pB, cpB)
	packSGodotRid(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotRidOperatorLess function as declared in gdnative/rid.h:67
func GodotRidOperatorLess(pSelf []GodotRid, pB []GodotRid) GodotBool {
	cpSelf, _ := unpackArgSGodotRid(pSelf)
	cpB, _ := unpackArgSGodotRid(pB)
	__ret := C.godot_rid_operator_less(cpSelf, cpB)
	packSGodotRid(pB, cpB)
	packSGodotRid(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotTransformNewWithAxisOrigin function as declared in gdnative/transform.h:62
func GodotTransformNewWithAxisOrigin(rDest []GodotTransform, pXAxis []GodotVector3, pYAxis []GodotVector3, pZAxis []GodotVector3, pOrigin []GodotVector3) {
	crDest, _ := unpackArgSGodotTransform(rDest)
	cpXAxis, _ := unpackArgSGodotVector3(pXAxis)
	cpYAxis, _ := unpackArgSGodotVector3(pYAxis)
	cpZAxis, _ := unpackArgSGodotVector3(pZAxis)
	cpOrigin, _ := unpackArgSGodotVector3(pOrigin)
	C.godot_transform_new_with_axis_origin(crDest, cpXAxis, cpYAxis, cpZAxis, cpOrigin)
	packSGodotVector3(pOrigin, cpOrigin)
	packSGodotVector3(pZAxis, cpZAxis)
	packSGodotVector3(pYAxis, cpYAxis)
	packSGodotVector3(pXAxis, cpXAxis)
	packSGodotTransform(rDest, crDest)
}

// GodotTransformNew function as declared in gdnative/transform.h:63
func GodotTransformNew(rDest []GodotTransform, pBasis []GodotBasis, pOrigin []GodotVector3) {
	crDest, _ := unpackArgSGodotTransform(rDest)
	cpBasis, _ := unpackArgSGodotBasis(pBasis)
	cpOrigin, _ := unpackArgSGodotVector3(pOrigin)
	C.godot_transform_new(crDest, cpBasis, cpOrigin)
	packSGodotVector3(pOrigin, cpOrigin)
	packSGodotBasis(pBasis, cpBasis)
	packSGodotTransform(rDest, crDest)
}

// GodotTransformGetBasis function as declared in gdnative/transform.h:65
func GodotTransformGetBasis(pSelf []GodotTransform) GodotBasis {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_get_basis(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformSetBasis function as declared in gdnative/transform.h:66
func GodotTransformSetBasis(pSelf []GodotTransform, pV []GodotBasis) {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotBasis(pV)
	C.godot_transform_set_basis(cpSelf, cpV)
	packSGodotBasis(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
}

// GodotTransformGetOrigin function as declared in gdnative/transform.h:68
func GodotTransformGetOrigin(pSelf []GodotTransform) GodotVector3 {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_get_origin(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformSetOrigin function as declared in gdnative/transform.h:69
func GodotTransformSetOrigin(pSelf []GodotTransform, pV []GodotVector3) {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	C.godot_transform_set_origin(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
}

// GodotTransformAsString function as declared in gdnative/transform.h:71
func GodotTransformAsString(pSelf []GodotTransform) GodotString {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_as_string(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformInverse function as declared in gdnative/transform.h:73
func GodotTransformInverse(pSelf []GodotTransform) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_inverse(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformAffineInverse function as declared in gdnative/transform.h:75
func GodotTransformAffineInverse(pSelf []GodotTransform) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_affine_inverse(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformOrthonormalized function as declared in gdnative/transform.h:77
func GodotTransformOrthonormalized(pSelf []GodotTransform) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	__ret := C.godot_transform_orthonormalized(cpSelf)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformRotated function as declared in gdnative/transform.h:79
func GodotTransformRotated(pSelf []GodotTransform, pAxis []GodotVector3, pPhi GodotReal) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpAxis, _ := unpackArgSGodotVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_transform_rotated(cpSelf, cpAxis, cpPhi)
	packSGodotVector3(pAxis, cpAxis)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformScaled function as declared in gdnative/transform.h:81
func GodotTransformScaled(pSelf []GodotTransform, pScale []GodotVector3) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpScale, _ := unpackArgSGodotVector3(pScale)
	__ret := C.godot_transform_scaled(cpSelf, cpScale)
	packSGodotVector3(pScale, cpScale)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformTranslated function as declared in gdnative/transform.h:83
func GodotTransformTranslated(pSelf []GodotTransform, pOfs []GodotVector3) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpOfs, _ := unpackArgSGodotVector3(pOfs)
	__ret := C.godot_transform_translated(cpSelf, cpOfs)
	packSGodotVector3(pOfs, cpOfs)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformLookingAt function as declared in gdnative/transform.h:85
func GodotTransformLookingAt(pSelf []GodotTransform, pTarget []GodotVector3, pUp []GodotVector3) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpTarget, _ := unpackArgSGodotVector3(pTarget)
	cpUp, _ := unpackArgSGodotVector3(pUp)
	__ret := C.godot_transform_looking_at(cpSelf, cpTarget, cpUp)
	packSGodotVector3(pUp, cpUp)
	packSGodotVector3(pTarget, cpTarget)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformPlane function as declared in gdnative/transform.h:87
func GodotTransformXformPlane(pSelf []GodotTransform, pV []GodotPlane) GodotPlane {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotPlane(pV)
	__ret := C.godot_transform_xform_plane(cpSelf, cpV)
	packSGodotPlane(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformInvPlane function as declared in gdnative/transform.h:89
func GodotTransformXformInvPlane(pSelf []GodotTransform, pV []GodotPlane) GodotPlane {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotPlane(pV)
	__ret := C.godot_transform_xform_inv_plane(cpSelf, cpV)
	packSGodotPlane(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformNewIdentity function as declared in gdnative/transform.h:91
func GodotTransformNewIdentity(rDest []GodotTransform) {
	crDest, _ := unpackArgSGodotTransform(rDest)
	C.godot_transform_new_identity(crDest)
	packSGodotTransform(rDest, crDest)
}

// GodotTransformOperatorEqual function as declared in gdnative/transform.h:93
func GodotTransformOperatorEqual(pSelf []GodotTransform, pB []GodotTransform) GodotBool {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpB, _ := unpackArgSGodotTransform(pB)
	__ret := C.godot_transform_operator_equal(cpSelf, cpB)
	packSGodotTransform(pB, cpB)
	packSGodotTransform(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotTransformOperatorMultiply function as declared in gdnative/transform.h:95
func GodotTransformOperatorMultiply(pSelf []GodotTransform, pB []GodotTransform) GodotTransform {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpB, _ := unpackArgSGodotTransform(pB)
	__ret := C.godot_transform_operator_multiply(cpSelf, cpB)
	packSGodotTransform(pB, cpB)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformVector3 function as declared in gdnative/transform.h:97
func GodotTransformXformVector3(pSelf []GodotTransform, pV []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	__ret := C.godot_transform_xform_vector3(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformInvVector3 function as declared in gdnative/transform.h:99
func GodotTransformXformInvVector3(pSelf []GodotTransform, pV []GodotVector3) GodotVector3 {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotVector3(pV)
	__ret := C.godot_transform_xform_inv_vector3(cpSelf, cpV)
	packSGodotVector3(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformAabb function as declared in gdnative/transform.h:101
func GodotTransformXformAabb(pSelf []GodotTransform, pV []GodotAabb) GodotAabb {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotAabb(pV)
	__ret := C.godot_transform_xform_aabb(cpSelf, cpV)
	packSGodotAabb(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransformXformInvAabb function as declared in gdnative/transform.h:103
func GodotTransformXformInvAabb(pSelf []GodotTransform, pV []GodotAabb) GodotAabb {
	cpSelf, _ := unpackArgSGodotTransform(pSelf)
	cpV, _ := unpackArgSGodotAabb(pV)
	__ret := C.godot_transform_xform_inv_aabb(cpSelf, cpV)
	packSGodotAabb(pV, cpV)
	packSGodotTransform(pSelf, cpSelf)
	__v := *NewGodotAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dNew function as declared in gdnative/transform2d.h:61
func GodotTransform2dNew(rDest []GodotTransform2d, pRot GodotReal, pPos []GodotVector2) {
	crDest, _ := unpackArgSGodotTransform2d(rDest)
	cpRot, _ := (C.godot_real)(pRot), cgoAllocsUnknown
	cpPos, _ := unpackArgSGodotVector2(pPos)
	C.godot_transform2d_new(crDest, cpRot, cpPos)
	packSGodotVector2(pPos, cpPos)
	packSGodotTransform2d(rDest, crDest)
}

// GodotTransform2dNewAxisOrigin function as declared in gdnative/transform2d.h:62
func GodotTransform2dNewAxisOrigin(rDest []GodotTransform2d, pXAxis []GodotVector2, pYAxis []GodotVector2, pOrigin []GodotVector2) {
	crDest, _ := unpackArgSGodotTransform2d(rDest)
	cpXAxis, _ := unpackArgSGodotVector2(pXAxis)
	cpYAxis, _ := unpackArgSGodotVector2(pYAxis)
	cpOrigin, _ := unpackArgSGodotVector2(pOrigin)
	C.godot_transform2d_new_axis_origin(crDest, cpXAxis, cpYAxis, cpOrigin)
	packSGodotVector2(pOrigin, cpOrigin)
	packSGodotVector2(pYAxis, cpYAxis)
	packSGodotVector2(pXAxis, cpXAxis)
	packSGodotTransform2d(rDest, crDest)
}

// GodotTransform2dAsString function as declared in gdnative/transform2d.h:64
func GodotTransform2dAsString(pSelf []GodotTransform2d) GodotString {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_as_string(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dInverse function as declared in gdnative/transform2d.h:66
func GodotTransform2dInverse(pSelf []GodotTransform2d) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_inverse(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dAffineInverse function as declared in gdnative/transform2d.h:68
func GodotTransform2dAffineInverse(pSelf []GodotTransform2d) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_affine_inverse(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dGetRotation function as declared in gdnative/transform2d.h:70
func GodotTransform2dGetRotation(pSelf []GodotTransform2d) GodotReal {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_get_rotation(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotTransform2dGetOrigin function as declared in gdnative/transform2d.h:72
func GodotTransform2dGetOrigin(pSelf []GodotTransform2d) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_get_origin(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dGetScale function as declared in gdnative/transform2d.h:74
func GodotTransform2dGetScale(pSelf []GodotTransform2d) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_get_scale(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dOrthonormalized function as declared in gdnative/transform2d.h:76
func GodotTransform2dOrthonormalized(pSelf []GodotTransform2d) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	__ret := C.godot_transform2d_orthonormalized(cpSelf)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dRotated function as declared in gdnative/transform2d.h:78
func GodotTransform2dRotated(pSelf []GodotTransform2d, pPhi GodotReal) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_transform2d_rotated(cpSelf, cpPhi)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dScaled function as declared in gdnative/transform2d.h:80
func GodotTransform2dScaled(pSelf []GodotTransform2d, pScale []GodotVector2) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpScale, _ := unpackArgSGodotVector2(pScale)
	__ret := C.godot_transform2d_scaled(cpSelf, cpScale)
	packSGodotVector2(pScale, cpScale)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dTranslated function as declared in gdnative/transform2d.h:82
func GodotTransform2dTranslated(pSelf []GodotTransform2d, pOffset []GodotVector2) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpOffset, _ := unpackArgSGodotVector2(pOffset)
	__ret := C.godot_transform2d_translated(cpSelf, cpOffset)
	packSGodotVector2(pOffset, cpOffset)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dXformVector2 function as declared in gdnative/transform2d.h:84
func GodotTransform2dXformVector2(pSelf []GodotTransform2d, pV []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotVector2(pV)
	__ret := C.godot_transform2d_xform_vector2(cpSelf, cpV)
	packSGodotVector2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dXformInvVector2 function as declared in gdnative/transform2d.h:86
func GodotTransform2dXformInvVector2(pSelf []GodotTransform2d, pV []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotVector2(pV)
	__ret := C.godot_transform2d_xform_inv_vector2(cpSelf, cpV)
	packSGodotVector2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dBasisXformVector2 function as declared in gdnative/transform2d.h:88
func GodotTransform2dBasisXformVector2(pSelf []GodotTransform2d, pV []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotVector2(pV)
	__ret := C.godot_transform2d_basis_xform_vector2(cpSelf, cpV)
	packSGodotVector2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dBasisXformInvVector2 function as declared in gdnative/transform2d.h:90
func GodotTransform2dBasisXformInvVector2(pSelf []GodotTransform2d, pV []GodotVector2) GodotVector2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotVector2(pV)
	__ret := C.godot_transform2d_basis_xform_inv_vector2(cpSelf, cpV)
	packSGodotVector2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dInterpolateWith function as declared in gdnative/transform2d.h:92
func GodotTransform2dInterpolateWith(pSelf []GodotTransform2d, pM []GodotTransform2d, pC GodotReal) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpM, _ := unpackArgSGodotTransform2d(pM)
	cpC, _ := (C.godot_real)(pC), cgoAllocsUnknown
	__ret := C.godot_transform2d_interpolate_with(cpSelf, cpM, cpC)
	packSGodotTransform2d(pM, cpM)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dOperatorEqual function as declared in gdnative/transform2d.h:94
func GodotTransform2dOperatorEqual(pSelf []GodotTransform2d, pB []GodotTransform2d) GodotBool {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpB, _ := unpackArgSGodotTransform2d(pB)
	__ret := C.godot_transform2d_operator_equal(cpSelf, cpB)
	packSGodotTransform2d(pB, cpB)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotTransform2dOperatorMultiply function as declared in gdnative/transform2d.h:96
func GodotTransform2dOperatorMultiply(pSelf []GodotTransform2d, pB []GodotTransform2d) GodotTransform2d {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpB, _ := unpackArgSGodotTransform2d(pB)
	__ret := C.godot_transform2d_operator_multiply(cpSelf, cpB)
	packSGodotTransform2d(pB, cpB)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dNewIdentity function as declared in gdnative/transform2d.h:98
func GodotTransform2dNewIdentity(rDest []GodotTransform2d) {
	crDest, _ := unpackArgSGodotTransform2d(rDest)
	C.godot_transform2d_new_identity(crDest)
	packSGodotTransform2d(rDest, crDest)
}

// GodotTransform2dXformRect2 function as declared in gdnative/transform2d.h:100
func GodotTransform2dXformRect2(pSelf []GodotTransform2d, pV []GodotRect2) GodotRect2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotRect2(pV)
	__ret := C.godot_transform2d_xform_rect2(cpSelf, cpV)
	packSGodotRect2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotTransform2dXformInvRect2 function as declared in gdnative/transform2d.h:102
func GodotTransform2dXformInvRect2(pSelf []GodotTransform2d, pV []GodotRect2) GodotRect2 {
	cpSelf, _ := unpackArgSGodotTransform2d(pSelf)
	cpV, _ := unpackArgSGodotRect2(pV)
	__ret := C.godot_transform2d_xform_inv_rect2(cpSelf, cpV)
	packSGodotRect2(pV, cpV)
	packSGodotTransform2d(pSelf, cpSelf)
	__v := *NewGodotRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNameNew function as declared in gdnative/string_name.h:60
func GodotStringNameNew(rDest []GodotStringName, pName []GodotString) {
	crDest, _ := unpackArgSGodotStringName(rDest)
	cpName, _ := unpackArgSGodotString(pName)
	C.godot_string_name_new(crDest, cpName)
	packSGodotString(pName, cpName)
	packSGodotStringName(rDest, crDest)
}

// GodotStringNameNewData function as declared in gdnative/string_name.h:61
func GodotStringNameNewData(rDest []GodotStringName, pName string) {
	crDest, _ := unpackArgSGodotStringName(rDest)
	cpName, _ := unpackPCharString(pName)
	C.godot_string_name_new_data(crDest, cpName)
	packSGodotStringName(rDest, crDest)
}

// GodotStringNameGetName function as declared in gdnative/string_name.h:63
func GodotStringNameGetName(pSelf []GodotStringName) GodotString {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	__ret := C.godot_string_name_get_name(cpSelf)
	packSGodotStringName(pSelf, cpSelf)
	__v := *NewGodotStringRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNameGetHash function as declared in gdnative/string_name.h:65
func GodotStringNameGetHash(pSelf []GodotStringName) uint32 {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	__ret := C.godot_string_name_get_hash(cpSelf)
	packSGodotStringName(pSelf, cpSelf)
	__v := (uint32)(__ret)
	return __v
}

// GodotStringNameGetDataUniquePointer function as declared in gdnative/string_name.h:66
func GodotStringNameGetDataUniquePointer(pSelf []GodotStringName) unsafe.Pointer {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	__ret := C.godot_string_name_get_data_unique_pointer(cpSelf)
	packSGodotStringName(pSelf, cpSelf)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// GodotStringNameOperatorEqual function as declared in gdnative/string_name.h:68
func GodotStringNameOperatorEqual(pSelf []GodotStringName, pOther []GodotStringName) GodotBool {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	cpOther, _ := unpackArgSGodotStringName(pOther)
	__ret := C.godot_string_name_operator_equal(cpSelf, cpOther)
	packSGodotStringName(pOther, cpOther)
	packSGodotStringName(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringNameOperatorLess function as declared in gdnative/string_name.h:69
func GodotStringNameOperatorLess(pSelf []GodotStringName, pOther []GodotStringName) GodotBool {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	cpOther, _ := unpackArgSGodotStringName(pOther)
	__ret := C.godot_string_name_operator_less(cpSelf, cpOther)
	packSGodotStringName(pOther, cpOther)
	packSGodotStringName(pSelf, cpSelf)
	__v := (GodotBool)(__ret)
	return __v
}

// GodotStringNameDestroy function as declared in gdnative/string_name.h:71
func GodotStringNameDestroy(pSelf []GodotStringName) {
	cpSelf, _ := unpackArgSGodotStringName(pSelf)
	C.godot_string_name_destroy(cpSelf)
	packSGodotStringName(pSelf, cpSelf)
}

// GodotArvrRegisterInterface function as declared in arvr/godot_arvr.h:57
func GodotArvrRegisterInterface(pInterface []GodotArvrInterfaceGdnative) {
	cpInterface, _ := unpackArgSGodotArvrInterfaceGdnative(pInterface)
	C.godot_arvr_register_interface(cpInterface)
	packSGodotArvrInterfaceGdnative(pInterface, cpInterface)
}

// GodotArvrGetWorldscale function as declared in arvr/godot_arvr.h:60
func GodotArvrGetWorldscale() GodotReal {
	__ret := C.godot_arvr_get_worldscale()
	__v := (GodotReal)(__ret)
	return __v
}

// GodotArvrGetReferenceFrame function as declared in arvr/godot_arvr.h:61
func GodotArvrGetReferenceFrame() GodotTransform {
	__ret := C.godot_arvr_get_reference_frame()
	__v := *NewGodotTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// GodotArvrBlit function as declared in arvr/godot_arvr.h:64
func GodotArvrBlit(pEye GodotInt, pRenderTarget []GodotRid, pRect []GodotRect2) {
	cpEye, _ := (C.godot_int)(pEye), cgoAllocsUnknown
	cpRenderTarget, _ := unpackArgSGodotRid(pRenderTarget)
	cpRect, _ := unpackArgSGodotRect2(pRect)
	C.godot_arvr_blit(cpEye, cpRenderTarget, cpRect)
	packSGodotRect2(pRect, cpRect)
	packSGodotRid(pRenderTarget, cpRenderTarget)
}

// GodotArvrGetTexid function as declared in arvr/godot_arvr.h:65
func GodotArvrGetTexid(pRenderTarget []GodotRid) GodotInt {
	cpRenderTarget, _ := unpackArgSGodotRid(pRenderTarget)
	__ret := C.godot_arvr_get_texid(cpRenderTarget)
	packSGodotRid(pRenderTarget, cpRenderTarget)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArvrAddController function as declared in arvr/godot_arvr.h:68
func GodotArvrAddController(pDeviceName []byte, pHand GodotInt, pTracksOrientation GodotBool, pTracksPosition GodotBool) GodotInt {
	cpDeviceName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pDeviceName)).Data)), cgoAllocsUnknown
	cpHand, _ := (C.godot_int)(pHand), cgoAllocsUnknown
	cpTracksOrientation, _ := (C.godot_bool)(pTracksOrientation), cgoAllocsUnknown
	cpTracksPosition, _ := (C.godot_bool)(pTracksPosition), cgoAllocsUnknown
	__ret := C.godot_arvr_add_controller(cpDeviceName, cpHand, cpTracksOrientation, cpTracksPosition)
	__v := (GodotInt)(__ret)
	return __v
}

// GodotArvrRemoveController function as declared in arvr/godot_arvr.h:69
func GodotArvrRemoveController(pControllerId GodotInt) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	C.godot_arvr_remove_controller(cpControllerId)
}

// GodotArvrSetControllerTransform function as declared in arvr/godot_arvr.h:70
func GodotArvrSetControllerTransform(pControllerId GodotInt, pTransform []GodotTransform, pTracksOrientation GodotBool, pTracksPosition GodotBool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpTransform, _ := unpackArgSGodotTransform(pTransform)
	cpTracksOrientation, _ := (C.godot_bool)(pTracksOrientation), cgoAllocsUnknown
	cpTracksPosition, _ := (C.godot_bool)(pTracksPosition), cgoAllocsUnknown
	C.godot_arvr_set_controller_transform(cpControllerId, cpTransform, cpTracksOrientation, cpTracksPosition)
	packSGodotTransform(pTransform, cpTransform)
}

// GodotArvrSetControllerButton function as declared in arvr/godot_arvr.h:71
func GodotArvrSetControllerButton(pControllerId GodotInt, pButton GodotInt, pIsPressed GodotBool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpButton, _ := (C.godot_int)(pButton), cgoAllocsUnknown
	cpIsPressed, _ := (C.godot_bool)(pIsPressed), cgoAllocsUnknown
	C.godot_arvr_set_controller_button(cpControllerId, cpButton, cpIsPressed)
}

// GodotArvrSetControllerAxis function as declared in arvr/godot_arvr.h:72
func GodotArvrSetControllerAxis(pControllerId GodotInt, pAxis GodotInt, pValue GodotReal, pCanBeNegative GodotBool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	cpValue, _ := (C.godot_real)(pValue), cgoAllocsUnknown
	cpCanBeNegative, _ := (C.godot_bool)(pCanBeNegative), cgoAllocsUnknown
	C.godot_arvr_set_controller_axis(cpControllerId, cpAxis, cpValue, cpCanBeNegative)
}

// GodotArvrGetControllerRumble function as declared in arvr/godot_arvr.h:73
func GodotArvrGetControllerRumble(pControllerId GodotInt) GodotReal {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	__ret := C.godot_arvr_get_controller_rumble(cpControllerId)
	__v := (GodotReal)(__ret)
	return __v
}

// GodotNativescriptRegisterClass function as declared in nativescript/godot_nativescript.h:133
func GodotNativescriptRegisterClass(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc GodotInstanceCreateFunc, pDestroyFunc GodotInstanceDestroyFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpBase, _ := unpackPCharString(pBase)
	cpCreateFunc, _ := pCreateFunc.PassValue()
	cpDestroyFunc, _ := pDestroyFunc.PassValue()
	C.godot_nativescript_register_class(cpGdnativeHandle, cpName, cpBase, cpCreateFunc, cpDestroyFunc)
}

// GodotNativescriptRegisterToolClass function as declared in nativescript/godot_nativescript.h:135
func GodotNativescriptRegisterToolClass(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc GodotInstanceCreateFunc, pDestroyFunc GodotInstanceDestroyFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpBase, _ := unpackPCharString(pBase)
	cpCreateFunc, _ := pCreateFunc.PassValue()
	cpDestroyFunc, _ := pDestroyFunc.PassValue()
	C.godot_nativescript_register_tool_class(cpGdnativeHandle, cpName, cpBase, cpCreateFunc, cpDestroyFunc)
}

// GodotNativescriptRegisterMethod function as declared in nativescript/godot_nativescript.h:148
func GodotNativescriptRegisterMethod(pGdnativeHandle unsafe.Pointer, pName string, pFunctionName string, pAttr GodotMethodAttributes, pMethod GodotInstanceMethod) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpFunctionName, _ := unpackPCharString(pFunctionName)
	cpAttr, _ := pAttr.PassValue()
	cpMethod, _ := pMethod.PassValue()
	C.godot_nativescript_register_method(cpGdnativeHandle, cpName, cpFunctionName, cpAttr, cpMethod)
}

// GodotNativescriptRegisterProperty function as declared in nativescript/godot_nativescript.h:164
func GodotNativescriptRegisterProperty(pGdnativeHandle unsafe.Pointer, pName string, pPath string, pAttr []GodotPropertyAttributes, pSetFunc GodotPropertySetFunc, pGetFunc GodotPropertyGetFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpPath, _ := unpackPCharString(pPath)
	cpAttr, _ := unpackArgSGodotPropertyAttributes(pAttr)
	cpSetFunc, _ := pSetFunc.PassValue()
	cpGetFunc, _ := pGetFunc.PassValue()
	C.godot_nativescript_register_property(cpGdnativeHandle, cpName, cpPath, cpAttr, cpSetFunc, cpGetFunc)
	packSGodotPropertyAttributes(pAttr, cpAttr)
}

// GodotNativescriptRegisterSignal function as declared in nativescript/godot_nativescript.h:183
func GodotNativescriptRegisterSignal(pGdnativeHandle unsafe.Pointer, pName string, pSignal []GodotSignal) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpSignal, _ := unpackArgSGodotSignal(pSignal)
	C.godot_nativescript_register_signal(cpGdnativeHandle, cpName, cpSignal)
	packSGodotSignal(pSignal, cpSignal)
}

// GodotNativescriptGetUserdata function as declared in nativescript/godot_nativescript.h:185
func GodotNativescriptGetUserdata(pInstance *GodotObject) unsafe.Pointer {
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	__ret := C.godot_nativescript_get_userdata(cpInstance)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// GodotPluginscriptRegisterLanguage function as declared in pluginscript/godot_pluginscript.h:165
func GodotPluginscriptRegisterLanguage(languageDesc []GodotPluginscriptLanguageDesc) {
	clanguageDesc, _ := unpackArgSGodotPluginscriptLanguageDesc(languageDesc)
	C.godot_pluginscript_register_language(clanguageDesc)
	packSGodotPluginscriptLanguageDesc(languageDesc, clanguageDesc)
}
