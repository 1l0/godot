// WARNING: This file has automatically been generated on Sun, 24 Dec 2017 07:11:06 JST.
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

// ObjectDestroy function as declared in gdnative/gdnative.h:207
func ObjectDestroy(pO *Object) {
	cpO, _ := (C.godot_object)(unsafe.Pointer(pO)), cgoAllocsUnknown
	C.godot_object_destroy(cpO)
}

// GlobalGetSingleton function as declared in gdnative/gdnative.h:215
func GlobalGetSingleton(pName []byte) *Object {
	cpName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pName)).Data)), cgoAllocsUnknown
	__ret := C.godot_global_get_singleton(cpName)
	__v := *(**Object)(unsafe.Pointer(&__ret))
	return __v
}

// GetStackBottom function as declared in gdnative/gdnative.h:219
func GetStackBottom() unsafe.Pointer {
	__ret := C.godot_get_stack_bottom()
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// MethodBindGetMethod function as declared in gdnative/gdnative.h:227
func MethodBindGetMethod(pClassname string, pMethodname string) *MethodBind {
	cpClassname, _ := unpackPCharString(pClassname)
	cpMethodname, _ := unpackPCharString(pMethodname)
	__ret := C.godot_method_bind_get_method(cpClassname, cpMethodname)
	__v := NewMethodBindRef(unsafe.Pointer(__ret))
	return __v
}

// MethodBindPtrcall function as declared in gdnative/gdnative.h:228
func MethodBindPtrcall(pMethodBind []MethodBind, pInstance *Object, pArgs []unsafe.Pointer, pRet unsafe.Pointer) {
	cpMethodBind, _ := unpackArgSMethodBind(pMethodBind)
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	cpArgs, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pArgs)).Data)), cgoAllocsUnknown
	cpRet, _ := pRet, cgoAllocsUnknown
	C.godot_method_bind_ptrcall(cpMethodBind, cpInstance, cpArgs, cpRet)
	packSMethodBind(pMethodBind, cpMethodBind)
}

// MethodBindCall function as declared in gdnative/gdnative.h:229
func MethodBindCall(pMethodBind []MethodBind, pInstance *Object, pArgs [][]Variant, pArgCount int32, pCallError []VariantCallError) Variant {
	cpMethodBind, _ := unpackArgSMethodBind(pMethodBind)
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	cpArgs, _ := unpackArgSSVariant(pArgs)
	cpArgCount, _ := (C.int)(pArgCount), cgoAllocsUnknown
	cpCallError, _ := unpackArgSVariantCallError(pCallError)
	__ret := C.godot_method_bind_call(cpMethodBind, cpInstance, cpArgs, cpArgCount, cpCallError)
	packSVariantCallError(pCallError, cpCallError)
	packSSVariant(pArgs, cpArgs)
	packSMethodBind(pMethodBind, cpMethodBind)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// GetClassConstructor function as declared in gdnative/gdnative.h:266
func GetClassConstructor(pClassname string) ClassConstructor {
	cpClassname, _ := unpackPCharString(pClassname)
	__ret := C.godot_get_class_constructor(cpClassname)
	__v := *NewClassConstructorRef(unsafe.Pointer(&__ret))
	return __v
}

// GetGlobalConstants function as declared in gdnative/gdnative.h:268
func GetGlobalConstants() Dictionary {
	__ret := C.godot_get_global_constants()
	__v := *NewDictionaryRef(unsafe.Pointer(&__ret))
	return __v
}

// RegisterNativeCallType function as declared in gdnative/gdnative.h:278
func RegisterNativeCallType(pCallType string, pCallback NativeCallCb) {
	cpCallType, _ := unpackPCharString(pCallType)
	cpCallback, _ := pCallback.PassValue()
	C.godot_register_native_call_type(cpCallType, cpCallback)
}

// Alloc function as declared in gdnative/gdnative.h:281
func Alloc(pBytes int32) unsafe.Pointer {
	cpBytes, _ := (C.int)(pBytes), cgoAllocsUnknown
	__ret := C.godot_alloc(cpBytes)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// Realloc function as declared in gdnative/gdnative.h:282
func Realloc(pPtr unsafe.Pointer, pBytes int32) unsafe.Pointer {
	cpPtr, _ := pPtr, cgoAllocsUnknown
	cpBytes, _ := (C.int)(pBytes), cgoAllocsUnknown
	__ret := C.godot_realloc(cpPtr, cpBytes)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// Free function as declared in gdnative/gdnative.h:283
func Free(pPtr unsafe.Pointer) {
	cpPtr, _ := pPtr, cgoAllocsUnknown
	C.godot_free(cpPtr)
}

// PrintError function as declared in gdnative/gdnative.h:286
func PrintError(pDescription string, pFunction string, pFile string, pLine int32) {
	cpDescription, _ := unpackPCharString(pDescription)
	cpFunction, _ := unpackPCharString(pFunction)
	cpFile, _ := unpackPCharString(pFile)
	cpLine, _ := (C.int)(pLine), cgoAllocsUnknown
	C.godot_print_error(cpDescription, cpFunction, cpFile, cpLine)
}

// PrintWarning function as declared in gdnative/gdnative.h:287
func PrintWarning(pDescription string, pFunction string, pFile string, pLine int32) {
	cpDescription, _ := unpackPCharString(pDescription)
	cpFunction, _ := unpackPCharString(pFunction)
	cpFile, _ := unpackPCharString(pFile)
	cpLine, _ := (C.int)(pLine), cgoAllocsUnknown
	C.godot_print_warning(cpDescription, cpFunction, cpFile, cpLine)
}

// Print function as declared in gdnative/gdnative.h:288
func Print(pMessage []String) {
	cpMessage, _ := unpackArgSString(pMessage)
	C.godot_print(cpMessage)
	packSString(pMessage, cpMessage)
}

// StringNew function as declared in gdnative/string.h:62
func StringNew(rDest []String) {
	crDest, _ := unpackArgSString(rDest)
	C.godot_string_new(crDest)
	packSString(rDest, crDest)
}

// StringNewCopy function as declared in gdnative/string.h:63
func StringNewCopy(rDest []String, pSrc []String) {
	crDest, _ := unpackArgSString(rDest)
	cpSrc, _ := unpackArgSString(pSrc)
	C.godot_string_new_copy(crDest, cpSrc)
	packSString(pSrc, cpSrc)
	packSString(rDest, crDest)
}

// StringNewData function as declared in gdnative/string.h:64
func StringNewData(rDest []String, pContents string, pSize int32) {
	crDest, _ := unpackArgSString(rDest)
	cpContents, _ := unpackPCharString(pContents)
	cpSize, _ := (C.int)(pSize), cgoAllocsUnknown
	C.godot_string_new_data(crDest, cpContents, cpSize)
	packSString(rDest, crDest)
}

// StringNewUnicodeData function as declared in gdnative/string.h:65
func StringNewUnicodeData(rDest []String, pContents []int32, pSize int32) {
	crDest, _ := unpackArgSString(rDest)
	cpContents, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pContents)).Data)), cgoAllocsUnknown
	cpSize, _ := (C.int)(pSize), cgoAllocsUnknown
	C.godot_string_new_unicode_data(crDest, cpContents, cpSize)
	packSString(rDest, crDest)
}

// StringGetData function as declared in gdnative/string.h:67
func StringGetData(pSelf []String, pDest []byte, pSize []int32) {
	cpSelf, _ := unpackArgSString(pSelf)
	cpDest, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pDest)).Data)), cgoAllocsUnknown
	cpSize, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pSize)).Data)), cgoAllocsUnknown
	C.godot_string_get_data(cpSelf, cpDest, cpSize)
	packSString(pSelf, cpSelf)
}

// StringOperatorIndex function as declared in gdnative/string.h:69
func StringOperatorIndex(pSelf []String, pIdx Int) *int32 {
	cpSelf, _ := unpackArgSString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_operator_index(cpSelf, cpIdx)
	packSString(pSelf, cpSelf)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// StringOperatorIndexConst function as declared in gdnative/string.h:70
func StringOperatorIndexConst(pSelf []String, pIdx Int) int32 {
	cpSelf, _ := unpackArgSString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_operator_index_const(cpSelf, cpIdx)
	packSString(pSelf, cpSelf)
	__v := (int32)(__ret)
	return __v
}

// StringUnicodeStr function as declared in gdnative/string.h:71
func StringUnicodeStr(pSelf []String) *int32 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_unicode_str(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *(**int32)(unsafe.Pointer(&__ret))
	return __v
}

// StringOperatorEqual function as declared in gdnative/string.h:73
func StringOperatorEqual(pSelf []String, pB []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpB, _ := unpackArgSString(pB)
	__ret := C.godot_string_operator_equal(cpSelf, cpB)
	packSString(pB, cpB)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringOperatorLess function as declared in gdnative/string.h:74
func StringOperatorLess(pSelf []String, pB []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpB, _ := unpackArgSString(pB)
	__ret := C.godot_string_operator_less(cpSelf, cpB)
	packSString(pB, cpB)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringOperatorPlus function as declared in gdnative/string.h:75
func StringOperatorPlus(pSelf []String, pB []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpB, _ := unpackArgSString(pB)
	__ret := C.godot_string_operator_plus(cpSelf, cpB)
	packSString(pB, cpB)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringLength function as declared in gdnative/string.h:79
func StringLength(pSelf []String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_length(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringBeginsWith function as declared in gdnative/string.h:83
func StringBeginsWith(pSelf []String, pString []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpString, _ := unpackArgSString(pString)
	__ret := C.godot_string_begins_with(cpSelf, cpString)
	packSString(pString, cpString)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringBeginsWithCharArray function as declared in gdnative/string.h:84
func StringBeginsWithCharArray(pSelf []String, pCharArray string) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpCharArray, _ := unpackPCharString(pCharArray)
	__ret := C.godot_string_begins_with_char_array(cpSelf, cpCharArray)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringBigrams function as declared in gdnative/string.h:85
func StringBigrams(pSelf []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_bigrams(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringChr function as declared in gdnative/string.h:86
func StringChr(pCharacter int32) String {
	cpCharacter, _ := (C.wchar_t)(pCharacter), cgoAllocsUnknown
	__ret := C.godot_string_chr(cpCharacter)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringEndsWith function as declared in gdnative/string.h:87
func StringEndsWith(pSelf []String, pString []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpString, _ := unpackArgSString(pString)
	__ret := C.godot_string_ends_with(cpSelf, cpString)
	packSString(pString, cpString)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringFind function as declared in gdnative/string.h:88
func StringFind(pSelf []String, pWhat String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_find(cpSelf, cpWhat)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindFrom function as declared in gdnative/string.h:89
func StringFindFrom(pSelf []String, pWhat String, pFrom Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_find_from(cpSelf, cpWhat, cpFrom)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindmk function as declared in gdnative/string.h:90
func StringFindmk(pSelf []String, pKeys []Array) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKeys, _ := unpackArgSArray(pKeys)
	__ret := C.godot_string_findmk(cpSelf, cpKeys)
	packSArray(pKeys, cpKeys)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindmkFrom function as declared in gdnative/string.h:91
func StringFindmkFrom(pSelf []String, pKeys []Array, pFrom Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKeys, _ := unpackArgSArray(pKeys)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_findmk_from(cpSelf, cpKeys, cpFrom)
	packSArray(pKeys, cpKeys)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindmkFromInPlace function as declared in gdnative/string.h:92
func StringFindmkFromInPlace(pSelf []String, pKeys []Array, pFrom Int, rKey []Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKeys, _ := unpackArgSArray(pKeys)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	crKey, _ := (*C.godot_int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&rKey)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_findmk_from_in_place(cpSelf, cpKeys, cpFrom, crKey)
	packSArray(pKeys, cpKeys)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindn function as declared in gdnative/string.h:93
func StringFindn(pSelf []String, pWhat String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_findn(cpSelf, cpWhat)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindnFrom function as declared in gdnative/string.h:94
func StringFindnFrom(pSelf []String, pWhat String, pFrom Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_findn_from(cpSelf, cpWhat, cpFrom)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFindLast function as declared in gdnative/string.h:95
func StringFindLast(pSelf []String, pWhat String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_find_last(cpSelf, cpWhat)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringFormat function as declared in gdnative/string.h:96
func StringFormat(pSelf []String, pValues []Variant) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpValues, _ := unpackArgSVariant(pValues)
	__ret := C.godot_string_format(cpSelf, cpValues)
	packSVariant(pValues, cpValues)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringFormatWithCustomPlaceholder function as declared in gdnative/string.h:97
func StringFormatWithCustomPlaceholder(pSelf []String, pValues []Variant, pPlaceholder string) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpValues, _ := unpackArgSVariant(pValues)
	cpPlaceholder, _ := unpackPCharString(pPlaceholder)
	__ret := C.godot_string_format_with_custom_placeholder(cpSelf, cpValues, cpPlaceholder)
	packSVariant(pValues, cpValues)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHexEncodeBuffer function as declared in gdnative/string.h:98
func StringHexEncodeBuffer(pBuffer string, pLen Int) String {
	cpBuffer, _ := unpackPUint8TString(pBuffer)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hex_encode_buffer(cpBuffer, cpLen)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHexToInt function as declared in gdnative/string.h:99
func StringHexToInt(pSelf []String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hex_to_int(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringHexToIntWithoutPrefix function as declared in gdnative/string.h:100
func StringHexToIntWithoutPrefix(pSelf []String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hex_to_int_without_prefix(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringInsert function as declared in gdnative/string.h:101
func StringInsert(pSelf []String, pAtPos Int, pString String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpAtPos, _ := (C.godot_int)(pAtPos), cgoAllocsUnknown
	cpString, _ := pString.PassValue()
	__ret := C.godot_string_insert(cpSelf, cpAtPos, cpString)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringIsNumeric function as declared in gdnative/string.h:102
func StringIsNumeric(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_numeric(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsSubsequenceOf function as declared in gdnative/string.h:103
func StringIsSubsequenceOf(pSelf []String, pString []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpString, _ := unpackArgSString(pString)
	__ret := C.godot_string_is_subsequence_of(cpSelf, cpString)
	packSString(pString, cpString)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsSubsequenceOfi function as declared in gdnative/string.h:104
func StringIsSubsequenceOfi(pSelf []String, pString []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpString, _ := unpackArgSString(pString)
	__ret := C.godot_string_is_subsequence_ofi(cpSelf, cpString)
	packSString(pString, cpString)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringLpad function as declared in gdnative/string.h:105
func StringLpad(pSelf []String, pMinLength Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	__ret := C.godot_string_lpad(cpSelf, cpMinLength)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringLpadWithCustomCharacter function as declared in gdnative/string.h:106
func StringLpadWithCustomCharacter(pSelf []String, pMinLength Int, pCharacter []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	cpCharacter, _ := unpackArgSString(pCharacter)
	__ret := C.godot_string_lpad_with_custom_character(cpSelf, cpMinLength, cpCharacter)
	packSString(pCharacter, cpCharacter)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringMatch function as declared in gdnative/string.h:107
func StringMatch(pSelf []String, pWildcard []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWildcard, _ := unpackArgSString(pWildcard)
	__ret := C.godot_string_match(cpSelf, cpWildcard)
	packSString(pWildcard, cpWildcard)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringMatchn function as declared in gdnative/string.h:108
func StringMatchn(pSelf []String, pWildcard []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWildcard, _ := unpackArgSString(pWildcard)
	__ret := C.godot_string_matchn(cpSelf, cpWildcard)
	packSString(pWildcard, cpWildcard)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringMd5 function as declared in gdnative/string.h:109
func StringMd5(pMd5 string) String {
	cpMd5, _ := unpackPUint8TString(pMd5)
	__ret := C.godot_string_md5(cpMd5)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNum function as declared in gdnative/string.h:110
func StringNum(pNum float64) String {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num(cpNum)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNumInt64 function as declared in gdnative/string.h:111
func StringNumInt64(pNum int64, pBase Int) String {
	cpNum, _ := (C.int64_t)(pNum), cgoAllocsUnknown
	cpBase, _ := (C.godot_int)(pBase), cgoAllocsUnknown
	__ret := C.godot_string_num_int64(cpNum, cpBase)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNumInt64Capitalized function as declared in gdnative/string.h:112
func StringNumInt64Capitalized(pNum int64, pBase Int, pCapitalizeHex Bool) String {
	cpNum, _ := (C.int64_t)(pNum), cgoAllocsUnknown
	cpBase, _ := (C.godot_int)(pBase), cgoAllocsUnknown
	cpCapitalizeHex, _ := (C.godot_bool)(pCapitalizeHex), cgoAllocsUnknown
	__ret := C.godot_string_num_int64_capitalized(cpNum, cpBase, cpCapitalizeHex)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNumReal function as declared in gdnative/string.h:113
func StringNumReal(pNum float64) String {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num_real(cpNum)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNumScientific function as declared in gdnative/string.h:114
func StringNumScientific(pNum float64) String {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	__ret := C.godot_string_num_scientific(cpNum)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNumWithDecimals function as declared in gdnative/string.h:115
func StringNumWithDecimals(pNum float64, pDecimals Int) String {
	cpNum, _ := (C.double)(pNum), cgoAllocsUnknown
	cpDecimals, _ := (C.godot_int)(pDecimals), cgoAllocsUnknown
	__ret := C.godot_string_num_with_decimals(cpNum, cpDecimals)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringPadDecimals function as declared in gdnative/string.h:116
func StringPadDecimals(pSelf []String, pDigits Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpDigits, _ := (C.godot_int)(pDigits), cgoAllocsUnknown
	__ret := C.godot_string_pad_decimals(cpSelf, cpDigits)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringPadZeros function as declared in gdnative/string.h:117
func StringPadZeros(pSelf []String, pDigits Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpDigits, _ := (C.godot_int)(pDigits), cgoAllocsUnknown
	__ret := C.godot_string_pad_zeros(cpSelf, cpDigits)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringReplaceFirst function as declared in gdnative/string.h:118
func StringReplaceFirst(pSelf []String, pKey String, pWith String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replace_first(cpSelf, cpKey, cpWith)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringReplace function as declared in gdnative/string.h:119
func StringReplace(pSelf []String, pKey String, pWith String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replace(cpSelf, cpKey, cpWith)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringReplacen function as declared in gdnative/string.h:120
func StringReplacen(pSelf []String, pKey String, pWith String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpKey, _ := pKey.PassValue()
	cpWith, _ := pWith.PassValue()
	__ret := C.godot_string_replacen(cpSelf, cpKey, cpWith)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringRfind function as declared in gdnative/string.h:121
func StringRfind(pSelf []String, pWhat String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_rfind(cpSelf, cpWhat)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringRfindn function as declared in gdnative/string.h:122
func StringRfindn(pSelf []String, pWhat String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	__ret := C.godot_string_rfindn(cpSelf, cpWhat)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringRfindFrom function as declared in gdnative/string.h:123
func StringRfindFrom(pSelf []String, pWhat String, pFrom Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_rfind_from(cpSelf, cpWhat, cpFrom)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringRfindnFrom function as declared in gdnative/string.h:124
func StringRfindnFrom(pSelf []String, pWhat String, pFrom Int) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWhat, _ := pWhat.PassValue()
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_string_rfindn_from(cpSelf, cpWhat, cpFrom)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringRpad function as declared in gdnative/string.h:125
func StringRpad(pSelf []String, pMinLength Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	__ret := C.godot_string_rpad(cpSelf, cpMinLength)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringRpadWithCustomCharacter function as declared in gdnative/string.h:126
func StringRpadWithCustomCharacter(pSelf []String, pMinLength Int, pCharacter []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpMinLength, _ := (C.godot_int)(pMinLength), cgoAllocsUnknown
	cpCharacter, _ := unpackArgSString(pCharacter)
	__ret := C.godot_string_rpad_with_custom_character(cpSelf, cpMinLength, cpCharacter)
	packSString(pCharacter, cpCharacter)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSimilarity function as declared in gdnative/string.h:127
func StringSimilarity(pSelf []String, pString []String) Real {
	cpSelf, _ := unpackArgSString(pSelf)
	cpString, _ := unpackArgSString(pString)
	__ret := C.godot_string_similarity(cpSelf, cpString)
	packSString(pString, cpString)
	packSString(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// StringSprintf function as declared in gdnative/string.h:128
func StringSprintf(pSelf []String, pValues []Array, pError []Bool) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpValues, _ := unpackArgSArray(pValues)
	cpError, _ := (*C.godot_bool)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pError)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_sprintf(cpSelf, cpValues, cpError)
	packSArray(pValues, cpValues)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSubstr function as declared in gdnative/string.h:129
func StringSubstr(pSelf []String, pFrom Int, pChars Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	cpChars, _ := (C.godot_int)(pChars), cgoAllocsUnknown
	__ret := C.godot_string_substr(cpSelf, cpFrom, cpChars)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringToDouble function as declared in gdnative/string.h:130
func StringToDouble(pSelf []String) float64 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_double(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (float64)(__ret)
	return __v
}

// StringToFloat function as declared in gdnative/string.h:131
func StringToFloat(pSelf []String) Real {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_float(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// StringToInt function as declared in gdnative/string.h:132
func StringToInt(pSelf []String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_int(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringCamelcaseToUnderscore function as declared in gdnative/string.h:134
func StringCamelcaseToUnderscore(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_camelcase_to_underscore(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCamelcaseToUnderscoreLowercased function as declared in gdnative/string.h:135
func StringCamelcaseToUnderscoreLowercased(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_camelcase_to_underscore_lowercased(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCapitalize function as declared in gdnative/string.h:136
func StringCapitalize(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_capitalize(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCharToDouble function as declared in gdnative/string.h:137
func StringCharToDouble(pWhat string) float64 {
	cpWhat, _ := unpackPCharString(pWhat)
	__ret := C.godot_string_char_to_double(cpWhat)
	__v := (float64)(__ret)
	return __v
}

// StringCharToInt function as declared in gdnative/string.h:138
func StringCharToInt(pWhat string) Int {
	cpWhat, _ := unpackPCharString(pWhat)
	__ret := C.godot_string_char_to_int(cpWhat)
	__v := (Int)(__ret)
	return __v
}

// StringWcharToInt function as declared in gdnative/string.h:139
func StringWcharToInt(pStr []int32) int64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_wchar_to_int(cpStr)
	__v := (int64)(__ret)
	return __v
}

// StringCharToIntWithLen function as declared in gdnative/string.h:140
func StringCharToIntWithLen(pWhat string, pLen Int) Int {
	cpWhat, _ := unpackPCharString(pWhat)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_char_to_int_with_len(cpWhat, cpLen)
	__v := (Int)(__ret)
	return __v
}

// StringCharToInt64WithLen function as declared in gdnative/string.h:141
func StringCharToInt64WithLen(pStr []int32, pLen int32) int64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	cpLen, _ := (C.int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_char_to_int64_with_len(cpStr, cpLen)
	__v := (int64)(__ret)
	return __v
}

// StringHexToInt64 function as declared in gdnative/string.h:142
func StringHexToInt64(pSelf []String) int64 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hex_to_int64(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// StringHexToInt64WithPrefix function as declared in gdnative/string.h:143
func StringHexToInt64WithPrefix(pSelf []String) int64 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hex_to_int64_with_prefix(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// StringToInt64 function as declared in gdnative/string.h:144
func StringToInt64(pSelf []String) int64 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_int64(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// StringUnicodeCharToDouble function as declared in gdnative/string.h:145
func StringUnicodeCharToDouble(pStr []int32, rEnd [][]int32) float64 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	crEnd, _ := unpackArgSSInt32(rEnd)
	__ret := C.godot_string_unicode_char_to_double(cpStr, crEnd)
	packSSInt32(rEnd, crEnd)
	__v := (float64)(__ret)
	return __v
}

// StringGetSliceCount function as declared in gdnative/string.h:147
func StringGetSliceCount(pSelf []String, pSplitter String) Int {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := pSplitter.PassValue()
	__ret := C.godot_string_get_slice_count(cpSelf, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// StringGetSlice function as declared in gdnative/string.h:148
func StringGetSlice(pSelf []String, pSplitter String, pSlice Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := pSplitter.PassValue()
	cpSlice, _ := (C.godot_int)(pSlice), cgoAllocsUnknown
	__ret := C.godot_string_get_slice(cpSelf, cpSplitter, cpSlice)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringGetSlicec function as declared in gdnative/string.h:149
func StringGetSlicec(pSelf []String, pSplitter int32, pSlice Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := (C.wchar_t)(pSplitter), cgoAllocsUnknown
	cpSlice, _ := (C.godot_int)(pSlice), cgoAllocsUnknown
	__ret := C.godot_string_get_slicec(cpSelf, cpSplitter, cpSlice)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplit function as declared in gdnative/string.h:151
func StringSplit(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitAllowEmpty function as declared in gdnative/string.h:152
func StringSplitAllowEmpty(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split_allow_empty(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitFloats function as declared in gdnative/string.h:153
func StringSplitFloats(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split_floats(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitFloatsAllowsEmpty function as declared in gdnative/string.h:154
func StringSplitFloatsAllowsEmpty(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split_floats_allows_empty(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitFloatsMk function as declared in gdnative/string.h:155
func StringSplitFloatsMk(pSelf []String, pSplitters []Array) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitters, _ := unpackArgSArray(pSplitters)
	__ret := C.godot_string_split_floats_mk(cpSelf, cpSplitters)
	packSArray(pSplitters, cpSplitters)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitFloatsMkAllowsEmpty function as declared in gdnative/string.h:156
func StringSplitFloatsMkAllowsEmpty(pSelf []String, pSplitters []Array) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitters, _ := unpackArgSArray(pSplitters)
	__ret := C.godot_string_split_floats_mk_allows_empty(cpSelf, cpSplitters)
	packSArray(pSplitters, cpSplitters)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitInts function as declared in gdnative/string.h:157
func StringSplitInts(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split_ints(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitIntsAllowsEmpty function as declared in gdnative/string.h:158
func StringSplitIntsAllowsEmpty(pSelf []String, pSplitter []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitter, _ := unpackArgSString(pSplitter)
	__ret := C.godot_string_split_ints_allows_empty(cpSelf, cpSplitter)
	packSString(pSplitter, cpSplitter)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitIntsMk function as declared in gdnative/string.h:159
func StringSplitIntsMk(pSelf []String, pSplitters []Array) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitters, _ := unpackArgSArray(pSplitters)
	__ret := C.godot_string_split_ints_mk(cpSelf, cpSplitters)
	packSArray(pSplitters, cpSplitters)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitIntsMkAllowsEmpty function as declared in gdnative/string.h:160
func StringSplitIntsMkAllowsEmpty(pSelf []String, pSplitters []Array) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	cpSplitters, _ := unpackArgSArray(pSplitters)
	__ret := C.godot_string_split_ints_mk_allows_empty(cpSelf, cpSplitters)
	packSArray(pSplitters, cpSplitters)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSplitSpaces function as declared in gdnative/string.h:161
func StringSplitSpaces(pSelf []String) Array {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_split_spaces(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCharLowercase function as declared in gdnative/string.h:163
func StringCharLowercase(pChar int32) int32 {
	cpChar, _ := (C.wchar_t)(pChar), cgoAllocsUnknown
	__ret := C.godot_string_char_lowercase(cpChar)
	__v := (int32)(__ret)
	return __v
}

// StringCharUppercase function as declared in gdnative/string.h:164
func StringCharUppercase(pChar int32) int32 {
	cpChar, _ := (C.wchar_t)(pChar), cgoAllocsUnknown
	__ret := C.godot_string_char_uppercase(cpChar)
	__v := (int32)(__ret)
	return __v
}

// StringToLower function as declared in gdnative/string.h:165
func StringToLower(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_lower(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringToUpper function as declared in gdnative/string.h:166
func StringToUpper(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_to_upper(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringGetBasename function as declared in gdnative/string.h:168
func StringGetBasename(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_get_basename(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringGetExtension function as declared in gdnative/string.h:169
func StringGetExtension(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_get_extension(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringLeft function as declared in gdnative/string.h:170
func StringLeft(pSelf []String, pPos Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	__ret := C.godot_string_left(cpSelf, cpPos)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringOrdAt function as declared in gdnative/string.h:171
func StringOrdAt(pSelf []String, pIdx Int) int32 {
	cpSelf, _ := unpackArgSString(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_string_ord_at(cpSelf, cpIdx)
	packSString(pSelf, cpSelf)
	__v := (int32)(__ret)
	return __v
}

// StringPlusFile function as declared in gdnative/string.h:172
func StringPlusFile(pSelf []String, pFile []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpFile, _ := unpackArgSString(pFile)
	__ret := C.godot_string_plus_file(cpSelf, cpFile)
	packSString(pFile, cpFile)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringRight function as declared in gdnative/string.h:173
func StringRight(pSelf []String, pPos Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	__ret := C.godot_string_right(cpSelf, cpPos)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringStripEdges function as declared in gdnative/string.h:174
func StringStripEdges(pSelf []String, pLeft Bool, pRight Bool) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpLeft, _ := (C.godot_bool)(pLeft), cgoAllocsUnknown
	cpRight, _ := (C.godot_bool)(pRight), cgoAllocsUnknown
	__ret := C.godot_string_strip_edges(cpSelf, cpLeft, cpRight)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringStripEscapes function as declared in gdnative/string.h:175
func StringStripEscapes(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_strip_escapes(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringErase function as declared in gdnative/string.h:177
func StringErase(pSelf []String, pPos Int, pChars Int) {
	cpSelf, _ := unpackArgSString(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	cpChars, _ := (C.godot_int)(pChars), cgoAllocsUnknown
	C.godot_string_erase(cpSelf, cpPos, cpChars)
	packSString(pSelf, cpSelf)
}

// StringAscii function as declared in gdnative/string.h:179
func StringAscii(pSelf []String, result []byte) {
	cpSelf, _ := unpackArgSString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_ascii(cpSelf, cresult)
	packSString(pSelf, cpSelf)
}

// StringAsciiExtended function as declared in gdnative/string.h:180
func StringAsciiExtended(pSelf []String, result []byte) {
	cpSelf, _ := unpackArgSString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_ascii_extended(cpSelf, cresult)
	packSString(pSelf, cpSelf)
}

// StringUtf8 function as declared in gdnative/string.h:181
func StringUtf8(pSelf []String, result []byte) {
	cpSelf, _ := unpackArgSString(pSelf)
	cresult, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&result)).Data)), cgoAllocsUnknown
	C.godot_string_utf8(cpSelf, cresult)
	packSString(pSelf, cpSelf)
}

// StringParseUtf8 function as declared in gdnative/string.h:182
func StringParseUtf8(pSelf []String, pUtf8 string) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpUtf8, _ := unpackPCharString(pUtf8)
	__ret := C.godot_string_parse_utf8(cpSelf, cpUtf8)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringParseUtf8WithLen function as declared in gdnative/string.h:183
func StringParseUtf8WithLen(pSelf []String, pUtf8 string, pLen Int) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpUtf8, _ := unpackPCharString(pUtf8)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_parse_utf8_with_len(cpSelf, cpUtf8, cpLen)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringCharsToUtf8 function as declared in gdnative/string.h:184
func StringCharsToUtf8(pUtf8 string) String {
	cpUtf8, _ := unpackPCharString(pUtf8)
	__ret := C.godot_string_chars_to_utf8(cpUtf8)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCharsToUtf8WithLen function as declared in gdnative/string.h:185
func StringCharsToUtf8WithLen(pUtf8 string, pLen Int) String {
	cpUtf8, _ := unpackPCharString(pUtf8)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_chars_to_utf8_with_len(cpUtf8, cpLen)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHash function as declared in gdnative/string.h:187
func StringHash(pSelf []String) uint32 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hash(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (uint32)(__ret)
	return __v
}

// StringHash64 function as declared in gdnative/string.h:188
func StringHash64(pSelf []String) uint64 {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_hash64(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (uint64)(__ret)
	return __v
}

// StringHashChars function as declared in gdnative/string.h:189
func StringHashChars(pCstr string) uint32 {
	cpCstr, _ := unpackPCharString(pCstr)
	__ret := C.godot_string_hash_chars(cpCstr)
	__v := (uint32)(__ret)
	return __v
}

// StringHashCharsWithLen function as declared in gdnative/string.h:190
func StringHashCharsWithLen(pCstr string, pLen Int) uint32 {
	cpCstr, _ := unpackPCharString(pCstr)
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hash_chars_with_len(cpCstr, cpLen)
	__v := (uint32)(__ret)
	return __v
}

// StringHashUtf8Chars function as declared in gdnative/string.h:191
func StringHashUtf8Chars(pStr []int32) uint32 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	__ret := C.godot_string_hash_utf8_chars(cpStr)
	__v := (uint32)(__ret)
	return __v
}

// StringHashUtf8CharsWithLen function as declared in gdnative/string.h:192
func StringHashUtf8CharsWithLen(pStr []int32, pLen Int) uint32 {
	cpStr, _ := (*C.wchar_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pStr)).Data)), cgoAllocsUnknown
	cpLen, _ := (C.godot_int)(pLen), cgoAllocsUnknown
	__ret := C.godot_string_hash_utf8_chars_with_len(cpStr, cpLen)
	__v := (uint32)(__ret)
	return __v
}

// StringMd5Buffer function as declared in gdnative/string.h:193
func StringMd5Buffer(pSelf []String) PoolByteArray {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_md5_buffer(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringMd5Text function as declared in gdnative/string.h:194
func StringMd5Text(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_md5_text(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSha256Buffer function as declared in gdnative/string.h:195
func StringSha256Buffer(pSelf []String) PoolByteArray {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_sha256_buffer(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSha256Text function as declared in gdnative/string.h:196
func StringSha256Text(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_sha256_text(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringEmpty function as declared in gdnative/string.h:198
func StringEmpty(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_empty(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringGetBaseDir function as declared in gdnative/string.h:201
func StringGetBaseDir(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_get_base_dir(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringGetFile function as declared in gdnative/string.h:202
func StringGetFile(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_get_file(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHumanizeSize function as declared in gdnative/string.h:203
func StringHumanizeSize(pSize uint) String {
	cpSize, _ := (C.size_t)(pSize), cgoAllocsUnknown
	__ret := C.godot_string_humanize_size(cpSize)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringIsAbsPath function as declared in gdnative/string.h:204
func StringIsAbsPath(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_abs_path(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsRelPath function as declared in gdnative/string.h:205
func StringIsRelPath(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_rel_path(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsResourceFile function as declared in gdnative/string.h:206
func StringIsResourceFile(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_resource_file(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringPathTo function as declared in gdnative/string.h:207
func StringPathTo(pSelf []String, pPath []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpPath, _ := unpackArgSString(pPath)
	__ret := C.godot_string_path_to(cpSelf, cpPath)
	packSString(pPath, cpPath)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringPathToFile function as declared in gdnative/string.h:208
func StringPathToFile(pSelf []String, pPath []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpPath, _ := unpackArgSString(pPath)
	__ret := C.godot_string_path_to_file(cpSelf, cpPath)
	packSString(pPath, cpPath)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringSimplifyPath function as declared in gdnative/string.h:209
func StringSimplifyPath(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_simplify_path(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCEscape function as declared in gdnative/string.h:211
func StringCEscape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_c_escape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCEscapeMultiline function as declared in gdnative/string.h:212
func StringCEscapeMultiline(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_c_escape_multiline(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringCUnescape function as declared in gdnative/string.h:213
func StringCUnescape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_c_unescape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHttpEscape function as declared in gdnative/string.h:214
func StringHttpEscape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_http_escape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringHttpUnescape function as declared in gdnative/string.h:215
func StringHttpUnescape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_http_unescape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringJsonEscape function as declared in gdnative/string.h:216
func StringJsonEscape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_json_escape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringWordWrap function as declared in gdnative/string.h:217
func StringWordWrap(pSelf []String, pCharsPerLine Int) String {
	cpSelf, _ := unpackArgSString(pSelf)
	cpCharsPerLine, _ := (C.godot_int)(pCharsPerLine), cgoAllocsUnknown
	__ret := C.godot_string_word_wrap(cpSelf, cpCharsPerLine)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringXmlEscape function as declared in gdnative/string.h:218
func StringXmlEscape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_xml_escape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringXmlEscapeWithQuotes function as declared in gdnative/string.h:219
func StringXmlEscapeWithQuotes(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_xml_escape_with_quotes(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringXmlUnescape function as declared in gdnative/string.h:220
func StringXmlUnescape(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_xml_unescape(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringPercentDecode function as declared in gdnative/string.h:222
func StringPercentDecode(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_percent_decode(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringPercentEncode function as declared in gdnative/string.h:223
func StringPercentEncode(pSelf []String) String {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_percent_encode(cpSelf)
	packSString(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringIsValidFloat function as declared in gdnative/string.h:225
func StringIsValidFloat(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_valid_float(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsValidHexNumber function as declared in gdnative/string.h:226
func StringIsValidHexNumber(pSelf []String, pWithPrefix Bool) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	cpWithPrefix, _ := (C.godot_bool)(pWithPrefix), cgoAllocsUnknown
	__ret := C.godot_string_is_valid_hex_number(cpSelf, cpWithPrefix)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsValidHtmlColor function as declared in gdnative/string.h:227
func StringIsValidHtmlColor(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_valid_html_color(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsValidIdentifier function as declared in gdnative/string.h:228
func StringIsValidIdentifier(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_valid_identifier(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsValidInteger function as declared in gdnative/string.h:229
func StringIsValidInteger(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_valid_integer(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringIsValidIpAddress function as declared in gdnative/string.h:230
func StringIsValidIpAddress(pSelf []String) Bool {
	cpSelf, _ := unpackArgSString(pSelf)
	__ret := C.godot_string_is_valid_ip_address(cpSelf)
	packSString(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringDestroy function as declared in gdnative/string.h:232
func StringDestroy(pSelf []String) {
	cpSelf, _ := unpackArgSString(pSelf)
	C.godot_string_destroy(cpSelf)
	packSString(pSelf, cpSelf)
}

// ArrayNew function as declared in gdnative/array.h:63
func ArrayNew(rDest []Array) {
	crDest, _ := unpackArgSArray(rDest)
	C.godot_array_new(crDest)
	packSArray(rDest, crDest)
}

// ArrayNewCopy function as declared in gdnative/array.h:64
func ArrayNewCopy(rDest []Array, pSrc []Array) {
	crDest, _ := unpackArgSArray(rDest)
	cpSrc, _ := unpackArgSArray(pSrc)
	C.godot_array_new_copy(crDest, cpSrc)
	packSArray(pSrc, cpSrc)
	packSArray(rDest, crDest)
}

// ArrayNewPoolColorArray function as declared in gdnative/array.h:65
func ArrayNewPoolColorArray(rDest []Array, pPca []PoolColorArray) {
	crDest, _ := unpackArgSArray(rDest)
	cpPca, _ := unpackArgSPoolColorArray(pPca)
	C.godot_array_new_pool_color_array(crDest, cpPca)
	packSPoolColorArray(pPca, cpPca)
	packSArray(rDest, crDest)
}

// ArrayNewPoolVector3Array function as declared in gdnative/array.h:66
func ArrayNewPoolVector3Array(rDest []Array, pPv3a []PoolVector3Array) {
	crDest, _ := unpackArgSArray(rDest)
	cpPv3a, _ := unpackArgSPoolVector3Array(pPv3a)
	C.godot_array_new_pool_vector3_array(crDest, cpPv3a)
	packSPoolVector3Array(pPv3a, cpPv3a)
	packSArray(rDest, crDest)
}

// ArrayNewPoolVector2Array function as declared in gdnative/array.h:67
func ArrayNewPoolVector2Array(rDest []Array, pPv2a []PoolVector2Array) {
	crDest, _ := unpackArgSArray(rDest)
	cpPv2a, _ := unpackArgSPoolVector2Array(pPv2a)
	C.godot_array_new_pool_vector2_array(crDest, cpPv2a)
	packSPoolVector2Array(pPv2a, cpPv2a)
	packSArray(rDest, crDest)
}

// ArrayNewPoolStringArray function as declared in gdnative/array.h:68
func ArrayNewPoolStringArray(rDest []Array, pPsa []PoolStringArray) {
	crDest, _ := unpackArgSArray(rDest)
	cpPsa, _ := unpackArgSPoolStringArray(pPsa)
	C.godot_array_new_pool_string_array(crDest, cpPsa)
	packSPoolStringArray(pPsa, cpPsa)
	packSArray(rDest, crDest)
}

// ArrayNewPoolRealArray function as declared in gdnative/array.h:69
func ArrayNewPoolRealArray(rDest []Array, pPra []PoolRealArray) {
	crDest, _ := unpackArgSArray(rDest)
	cpPra, _ := unpackArgSPoolRealArray(pPra)
	C.godot_array_new_pool_real_array(crDest, cpPra)
	packSPoolRealArray(pPra, cpPra)
	packSArray(rDest, crDest)
}

// ArrayNewPoolIntArray function as declared in gdnative/array.h:70
func ArrayNewPoolIntArray(rDest []Array, pPia []PoolIntArray) {
	crDest, _ := unpackArgSArray(rDest)
	cpPia, _ := unpackArgSPoolIntArray(pPia)
	C.godot_array_new_pool_int_array(crDest, cpPia)
	packSPoolIntArray(pPia, cpPia)
	packSArray(rDest, crDest)
}

// ArrayNewPoolByteArray function as declared in gdnative/array.h:71
func ArrayNewPoolByteArray(rDest []Array, pPba []PoolByteArray) {
	crDest, _ := unpackArgSArray(rDest)
	cpPba, _ := unpackArgSPoolByteArray(pPba)
	C.godot_array_new_pool_byte_array(crDest, cpPba)
	packSPoolByteArray(pPba, cpPba)
	packSArray(rDest, crDest)
}

// ArraySet function as declared in gdnative/array.h:73
func ArraySet(pSelf []Array, pIdx Int, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_set(cpSelf, cpIdx, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayGet function as declared in gdnative/array.h:75
func ArrayGet(pSelf []Array, pIdx Int) Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_get(cpSelf, cpIdx)
	packSArray(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// ArrayOperatorIndex function as declared in gdnative/array.h:77
func ArrayOperatorIndex(pSelf []Array, pIdx Int) *Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_operator_index(cpSelf, cpIdx)
	packSArray(pSelf, cpSelf)
	__v := NewVariantRef(unsafe.Pointer(__ret))
	return __v
}

// ArrayOperatorIndexConst function as declared in gdnative/array.h:79
func ArrayOperatorIndexConst(pSelf []Array, pIdx Int) *Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_array_operator_index_const(cpSelf, cpIdx)
	packSArray(pSelf, cpSelf)
	__v := NewVariantRef(unsafe.Pointer(__ret))
	return __v
}

// ArrayAppend function as declared in gdnative/array.h:81
func ArrayAppend(pSelf []Array, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_append(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayClear function as declared in gdnative/array.h:83
func ArrayClear(pSelf []Array) {
	cpSelf, _ := unpackArgSArray(pSelf)
	C.godot_array_clear(cpSelf)
	packSArray(pSelf, cpSelf)
}

// ArrayCount function as declared in gdnative/array.h:85
func ArrayCount(pSelf []Array, pValue []Variant) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	__ret := C.godot_array_count(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayEmpty function as declared in gdnative/array.h:87
func ArrayEmpty(pSelf []Array) Bool {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_empty(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// ArrayErase function as declared in gdnative/array.h:89
func ArrayErase(pSelf []Array, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_erase(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayFront function as declared in gdnative/array.h:91
func ArrayFront(pSelf []Array) Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_front(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// ArrayBack function as declared in gdnative/array.h:93
func ArrayBack(pSelf []Array) Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_back(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// ArrayFind function as declared in gdnative/array.h:95
func ArrayFind(pSelf []Array, pWhat []Variant, pFrom Int) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpWhat, _ := unpackArgSVariant(pWhat)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_array_find(cpSelf, cpWhat, cpFrom)
	packSVariant(pWhat, cpWhat)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayFindLast function as declared in gdnative/array.h:97
func ArrayFindLast(pSelf []Array, pWhat []Variant) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpWhat, _ := unpackArgSVariant(pWhat)
	__ret := C.godot_array_find_last(cpSelf, cpWhat)
	packSVariant(pWhat, cpWhat)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayHas function as declared in gdnative/array.h:99
func ArrayHas(pSelf []Array, pValue []Variant) Bool {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	__ret := C.godot_array_has(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// ArrayHash function as declared in gdnative/array.h:101
func ArrayHash(pSelf []Array) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_hash(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayInsert function as declared in gdnative/array.h:103
func ArrayInsert(pSelf []Array, pPos Int, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpPos, _ := (C.godot_int)(pPos), cgoAllocsUnknown
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_insert(cpSelf, cpPos, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayInvert function as declared in gdnative/array.h:105
func ArrayInvert(pSelf []Array) {
	cpSelf, _ := unpackArgSArray(pSelf)
	C.godot_array_invert(cpSelf)
	packSArray(pSelf, cpSelf)
}

// ArrayPopBack function as declared in gdnative/array.h:107
func ArrayPopBack(pSelf []Array) Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_pop_back(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// ArrayPopFront function as declared in gdnative/array.h:109
func ArrayPopFront(pSelf []Array) Variant {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_pop_front(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// ArrayPushBack function as declared in gdnative/array.h:111
func ArrayPushBack(pSelf []Array, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_push_back(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayPushFront function as declared in gdnative/array.h:113
func ArrayPushFront(pSelf []Array, pValue []Variant) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_array_push_front(cpSelf, cpValue)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
}

// ArrayRemove function as declared in gdnative/array.h:115
func ArrayRemove(pSelf []Array, pIdx Int) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_array_remove(cpSelf, cpIdx)
	packSArray(pSelf, cpSelf)
}

// ArrayResize function as declared in gdnative/array.h:117
func ArrayResize(pSelf []Array, pSize Int) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_array_resize(cpSelf, cpSize)
	packSArray(pSelf, cpSelf)
}

// ArrayRfind function as declared in gdnative/array.h:119
func ArrayRfind(pSelf []Array, pWhat []Variant, pFrom Int) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpWhat, _ := unpackArgSVariant(pWhat)
	cpFrom, _ := (C.godot_int)(pFrom), cgoAllocsUnknown
	__ret := C.godot_array_rfind(cpSelf, cpWhat, cpFrom)
	packSVariant(pWhat, cpWhat)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArraySize function as declared in gdnative/array.h:121
func ArraySize(pSelf []Array) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	__ret := C.godot_array_size(cpSelf)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArraySort function as declared in gdnative/array.h:123
func ArraySort(pSelf []Array) {
	cpSelf, _ := unpackArgSArray(pSelf)
	C.godot_array_sort(cpSelf)
	packSArray(pSelf, cpSelf)
}

// ArraySortCustom function as declared in gdnative/array.h:125
func ArraySortCustom(pSelf []Array, pObj *Object, pFunc []String) {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	cpFunc, _ := unpackArgSString(pFunc)
	C.godot_array_sort_custom(cpSelf, cpObj, cpFunc)
	packSString(pFunc, cpFunc)
	packSArray(pSelf, cpSelf)
}

// ArrayBsearch function as declared in gdnative/array.h:127
func ArrayBsearch(pSelf []Array, pValue []Variant, pBefore Bool) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	cpBefore, _ := (C.godot_bool)(pBefore), cgoAllocsUnknown
	__ret := C.godot_array_bsearch(cpSelf, cpValue, cpBefore)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayBsearchCustom function as declared in gdnative/array.h:129
func ArrayBsearchCustom(pSelf []Array, pValue []Variant, pObj *Object, pFunc []String, pBefore Bool) Int {
	cpSelf, _ := unpackArgSArray(pSelf)
	cpValue, _ := unpackArgSVariant(pValue)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	cpFunc, _ := unpackArgSString(pFunc)
	cpBefore, _ := (C.godot_bool)(pBefore), cgoAllocsUnknown
	__ret := C.godot_array_bsearch_custom(cpSelf, cpValue, cpObj, cpFunc, cpBefore)
	packSString(pFunc, cpFunc)
	packSVariant(pValue, cpValue)
	packSArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ArrayDestroy function as declared in gdnative/array.h:131
func ArrayDestroy(pSelf []Array) {
	cpSelf, _ := unpackArgSArray(pSelf)
	C.godot_array_destroy(cpSelf)
	packSArray(pSelf, cpSelf)
}

// PoolByteArrayNew function as declared in gdnative/pool_arrays.h:166
func PoolByteArrayNew(rDest []PoolByteArray) {
	crDest, _ := unpackArgSPoolByteArray(rDest)
	C.godot_pool_byte_array_new(crDest)
	packSPoolByteArray(rDest, crDest)
}

// PoolByteArrayNewCopy function as declared in gdnative/pool_arrays.h:167
func PoolByteArrayNewCopy(rDest []PoolByteArray, pSrc []PoolByteArray) {
	crDest, _ := unpackArgSPoolByteArray(rDest)
	cpSrc, _ := unpackArgSPoolByteArray(pSrc)
	C.godot_pool_byte_array_new_copy(crDest, cpSrc)
	packSPoolByteArray(pSrc, cpSrc)
	packSPoolByteArray(rDest, crDest)
}

// PoolByteArrayNewWithArray function as declared in gdnative/pool_arrays.h:168
func PoolByteArrayNewWithArray(rDest []PoolByteArray, pA []Array) {
	crDest, _ := unpackArgSPoolByteArray(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_byte_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolByteArray(rDest, crDest)
}

// PoolByteArrayAppend function as declared in gdnative/pool_arrays.h:170
func PoolByteArrayAppend(pSelf []PoolByteArray, pData byte) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_append(cpSelf, cpData)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayAppendArray function as declared in gdnative/pool_arrays.h:172
func PoolByteArrayAppendArray(pSelf []PoolByteArray, pArray []PoolByteArray) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpArray, _ := unpackArgSPoolByteArray(pArray)
	C.godot_pool_byte_array_append_array(cpSelf, cpArray)
	packSPoolByteArray(pArray, cpArray)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayInsert function as declared in gdnative/pool_arrays.h:174
func PoolByteArrayInsert(pSelf []PoolByteArray, pIdx Int, pData byte) Error {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_byte_array_insert(cpSelf, cpIdx, cpData)
	packSPoolByteArray(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolByteArrayInvert function as declared in gdnative/pool_arrays.h:176
func PoolByteArrayInvert(pSelf []PoolByteArray) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	C.godot_pool_byte_array_invert(cpSelf)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayPushBack function as declared in gdnative/pool_arrays.h:178
func PoolByteArrayPushBack(pSelf []PoolByteArray, pData byte) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_push_back(cpSelf, cpData)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayRemove function as declared in gdnative/pool_arrays.h:180
func PoolByteArrayRemove(pSelf []PoolByteArray, pIdx Int) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_byte_array_remove(cpSelf, cpIdx)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayResize function as declared in gdnative/pool_arrays.h:182
func PoolByteArrayResize(pSelf []PoolByteArray, pSize Int) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_byte_array_resize(cpSelf, cpSize)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayRead function as declared in gdnative/pool_arrays.h:184
func PoolByteArrayRead(pSelf []PoolByteArray) *PoolByteArrayReadAccess {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_read(cpSelf)
	packSPoolByteArray(pSelf, cpSelf)
	__v := NewPoolByteArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolByteArrayWrite function as declared in gdnative/pool_arrays.h:186
func PoolByteArrayWrite(pSelf []PoolByteArray) *PoolByteArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_write(cpSelf)
	packSPoolByteArray(pSelf, cpSelf)
	__v := NewPoolByteArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolByteArraySet function as declared in gdnative/pool_arrays.h:188
func PoolByteArraySet(pSelf []PoolByteArray, pIdx Int, pData byte) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.uint8_t)(pData), cgoAllocsUnknown
	C.godot_pool_byte_array_set(cpSelf, cpIdx, cpData)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolByteArrayGet function as declared in gdnative/pool_arrays.h:189
func PoolByteArrayGet(pSelf []PoolByteArray, pIdx Int) byte {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_byte_array_get(cpSelf, cpIdx)
	packSPoolByteArray(pSelf, cpSelf)
	__v := (byte)(__ret)
	return __v
}

// PoolByteArraySize function as declared in gdnative/pool_arrays.h:191
func PoolByteArraySize(pSelf []PoolByteArray) Int {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	__ret := C.godot_pool_byte_array_size(cpSelf)
	packSPoolByteArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolByteArrayDestroy function as declared in gdnative/pool_arrays.h:193
func PoolByteArrayDestroy(pSelf []PoolByteArray) {
	cpSelf, _ := unpackArgSPoolByteArray(pSelf)
	C.godot_pool_byte_array_destroy(cpSelf)
	packSPoolByteArray(pSelf, cpSelf)
}

// PoolIntArrayNew function as declared in gdnative/pool_arrays.h:197
func PoolIntArrayNew(rDest []PoolIntArray) {
	crDest, _ := unpackArgSPoolIntArray(rDest)
	C.godot_pool_int_array_new(crDest)
	packSPoolIntArray(rDest, crDest)
}

// PoolIntArrayNewCopy function as declared in gdnative/pool_arrays.h:198
func PoolIntArrayNewCopy(rDest []PoolIntArray, pSrc []PoolIntArray) {
	crDest, _ := unpackArgSPoolIntArray(rDest)
	cpSrc, _ := unpackArgSPoolIntArray(pSrc)
	C.godot_pool_int_array_new_copy(crDest, cpSrc)
	packSPoolIntArray(pSrc, cpSrc)
	packSPoolIntArray(rDest, crDest)
}

// PoolIntArrayNewWithArray function as declared in gdnative/pool_arrays.h:199
func PoolIntArrayNewWithArray(rDest []PoolIntArray, pA []Array) {
	crDest, _ := unpackArgSPoolIntArray(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_int_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolIntArray(rDest, crDest)
}

// PoolIntArrayAppend function as declared in gdnative/pool_arrays.h:201
func PoolIntArrayAppend(pSelf []PoolIntArray, pData Int) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_append(cpSelf, cpData)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayAppendArray function as declared in gdnative/pool_arrays.h:203
func PoolIntArrayAppendArray(pSelf []PoolIntArray, pArray []PoolIntArray) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpArray, _ := unpackArgSPoolIntArray(pArray)
	C.godot_pool_int_array_append_array(cpSelf, cpArray)
	packSPoolIntArray(pArray, cpArray)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayInsert function as declared in gdnative/pool_arrays.h:205
func PoolIntArrayInsert(pSelf []PoolIntArray, pIdx Int, pData Int) Error {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_int_array_insert(cpSelf, cpIdx, cpData)
	packSPoolIntArray(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolIntArrayInvert function as declared in gdnative/pool_arrays.h:207
func PoolIntArrayInvert(pSelf []PoolIntArray) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	C.godot_pool_int_array_invert(cpSelf)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayPushBack function as declared in gdnative/pool_arrays.h:209
func PoolIntArrayPushBack(pSelf []PoolIntArray, pData Int) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_push_back(cpSelf, cpData)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayRemove function as declared in gdnative/pool_arrays.h:211
func PoolIntArrayRemove(pSelf []PoolIntArray, pIdx Int) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_int_array_remove(cpSelf, cpIdx)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayResize function as declared in gdnative/pool_arrays.h:213
func PoolIntArrayResize(pSelf []PoolIntArray, pSize Int) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_int_array_resize(cpSelf, cpSize)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayRead function as declared in gdnative/pool_arrays.h:215
func PoolIntArrayRead(pSelf []PoolIntArray) *PoolIntArrayReadAccess {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_read(cpSelf)
	packSPoolIntArray(pSelf, cpSelf)
	__v := NewPoolIntArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolIntArrayWrite function as declared in gdnative/pool_arrays.h:217
func PoolIntArrayWrite(pSelf []PoolIntArray) *PoolIntArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_write(cpSelf)
	packSPoolIntArray(pSelf, cpSelf)
	__v := NewPoolIntArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolIntArraySet function as declared in gdnative/pool_arrays.h:219
func PoolIntArraySet(pSelf []PoolIntArray, pIdx Int, pData Int) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_int)(pData), cgoAllocsUnknown
	C.godot_pool_int_array_set(cpSelf, cpIdx, cpData)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolIntArrayGet function as declared in gdnative/pool_arrays.h:220
func PoolIntArrayGet(pSelf []PoolIntArray, pIdx Int) Int {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_int_array_get(cpSelf, cpIdx)
	packSPoolIntArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolIntArraySize function as declared in gdnative/pool_arrays.h:222
func PoolIntArraySize(pSelf []PoolIntArray) Int {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	__ret := C.godot_pool_int_array_size(cpSelf)
	packSPoolIntArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolIntArrayDestroy function as declared in gdnative/pool_arrays.h:224
func PoolIntArrayDestroy(pSelf []PoolIntArray) {
	cpSelf, _ := unpackArgSPoolIntArray(pSelf)
	C.godot_pool_int_array_destroy(cpSelf)
	packSPoolIntArray(pSelf, cpSelf)
}

// PoolRealArrayNew function as declared in gdnative/pool_arrays.h:228
func PoolRealArrayNew(rDest []PoolRealArray) {
	crDest, _ := unpackArgSPoolRealArray(rDest)
	C.godot_pool_real_array_new(crDest)
	packSPoolRealArray(rDest, crDest)
}

// PoolRealArrayNewCopy function as declared in gdnative/pool_arrays.h:229
func PoolRealArrayNewCopy(rDest []PoolRealArray, pSrc []PoolRealArray) {
	crDest, _ := unpackArgSPoolRealArray(rDest)
	cpSrc, _ := unpackArgSPoolRealArray(pSrc)
	C.godot_pool_real_array_new_copy(crDest, cpSrc)
	packSPoolRealArray(pSrc, cpSrc)
	packSPoolRealArray(rDest, crDest)
}

// PoolRealArrayNewWithArray function as declared in gdnative/pool_arrays.h:230
func PoolRealArrayNewWithArray(rDest []PoolRealArray, pA []Array) {
	crDest, _ := unpackArgSPoolRealArray(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_real_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolRealArray(rDest, crDest)
}

// PoolRealArrayAppend function as declared in gdnative/pool_arrays.h:232
func PoolRealArrayAppend(pSelf []PoolRealArray, pData Real) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_append(cpSelf, cpData)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayAppendArray function as declared in gdnative/pool_arrays.h:234
func PoolRealArrayAppendArray(pSelf []PoolRealArray, pArray []PoolRealArray) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpArray, _ := unpackArgSPoolRealArray(pArray)
	C.godot_pool_real_array_append_array(cpSelf, cpArray)
	packSPoolRealArray(pArray, cpArray)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayInsert function as declared in gdnative/pool_arrays.h:236
func PoolRealArrayInsert(pSelf []PoolRealArray, pIdx Int, pData Real) Error {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	__ret := C.godot_pool_real_array_insert(cpSelf, cpIdx, cpData)
	packSPoolRealArray(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolRealArrayInvert function as declared in gdnative/pool_arrays.h:238
func PoolRealArrayInvert(pSelf []PoolRealArray) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	C.godot_pool_real_array_invert(cpSelf)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayPushBack function as declared in gdnative/pool_arrays.h:240
func PoolRealArrayPushBack(pSelf []PoolRealArray, pData Real) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_push_back(cpSelf, cpData)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayRemove function as declared in gdnative/pool_arrays.h:242
func PoolRealArrayRemove(pSelf []PoolRealArray, pIdx Int) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_real_array_remove(cpSelf, cpIdx)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayResize function as declared in gdnative/pool_arrays.h:244
func PoolRealArrayResize(pSelf []PoolRealArray, pSize Int) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_real_array_resize(cpSelf, cpSize)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayRead function as declared in gdnative/pool_arrays.h:246
func PoolRealArrayRead(pSelf []PoolRealArray) *PoolRealArrayReadAccess {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_read(cpSelf)
	packSPoolRealArray(pSelf, cpSelf)
	__v := NewPoolRealArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolRealArrayWrite function as declared in gdnative/pool_arrays.h:248
func PoolRealArrayWrite(pSelf []PoolRealArray) *PoolRealArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_write(cpSelf)
	packSPoolRealArray(pSelf, cpSelf)
	__v := NewPoolRealArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolRealArraySet function as declared in gdnative/pool_arrays.h:250
func PoolRealArraySet(pSelf []PoolRealArray, pIdx Int, pData Real) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := (C.godot_real)(pData), cgoAllocsUnknown
	C.godot_pool_real_array_set(cpSelf, cpIdx, cpData)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolRealArrayGet function as declared in gdnative/pool_arrays.h:251
func PoolRealArrayGet(pSelf []PoolRealArray, pIdx Int) Real {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_real_array_get(cpSelf, cpIdx)
	packSPoolRealArray(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// PoolRealArraySize function as declared in gdnative/pool_arrays.h:253
func PoolRealArraySize(pSelf []PoolRealArray) Int {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	__ret := C.godot_pool_real_array_size(cpSelf)
	packSPoolRealArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolRealArrayDestroy function as declared in gdnative/pool_arrays.h:255
func PoolRealArrayDestroy(pSelf []PoolRealArray) {
	cpSelf, _ := unpackArgSPoolRealArray(pSelf)
	C.godot_pool_real_array_destroy(cpSelf)
	packSPoolRealArray(pSelf, cpSelf)
}

// PoolStringArrayNew function as declared in gdnative/pool_arrays.h:259
func PoolStringArrayNew(rDest []PoolStringArray) {
	crDest, _ := unpackArgSPoolStringArray(rDest)
	C.godot_pool_string_array_new(crDest)
	packSPoolStringArray(rDest, crDest)
}

// PoolStringArrayNewCopy function as declared in gdnative/pool_arrays.h:260
func PoolStringArrayNewCopy(rDest []PoolStringArray, pSrc []PoolStringArray) {
	crDest, _ := unpackArgSPoolStringArray(rDest)
	cpSrc, _ := unpackArgSPoolStringArray(pSrc)
	C.godot_pool_string_array_new_copy(crDest, cpSrc)
	packSPoolStringArray(pSrc, cpSrc)
	packSPoolStringArray(rDest, crDest)
}

// PoolStringArrayNewWithArray function as declared in gdnative/pool_arrays.h:261
func PoolStringArrayNewWithArray(rDest []PoolStringArray, pA []Array) {
	crDest, _ := unpackArgSPoolStringArray(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_string_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolStringArray(rDest, crDest)
}

// PoolStringArrayAppend function as declared in gdnative/pool_arrays.h:263
func PoolStringArrayAppend(pSelf []PoolStringArray, pData []String) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpData, _ := unpackArgSString(pData)
	C.godot_pool_string_array_append(cpSelf, cpData)
	packSString(pData, cpData)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayAppendArray function as declared in gdnative/pool_arrays.h:265
func PoolStringArrayAppendArray(pSelf []PoolStringArray, pArray []PoolStringArray) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpArray, _ := unpackArgSPoolStringArray(pArray)
	C.godot_pool_string_array_append_array(cpSelf, cpArray)
	packSPoolStringArray(pArray, cpArray)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayInsert function as declared in gdnative/pool_arrays.h:267
func PoolStringArrayInsert(pSelf []PoolStringArray, pIdx Int, pData []String) Error {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSString(pData)
	__ret := C.godot_pool_string_array_insert(cpSelf, cpIdx, cpData)
	packSString(pData, cpData)
	packSPoolStringArray(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolStringArrayInvert function as declared in gdnative/pool_arrays.h:269
func PoolStringArrayInvert(pSelf []PoolStringArray) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	C.godot_pool_string_array_invert(cpSelf)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayPushBack function as declared in gdnative/pool_arrays.h:271
func PoolStringArrayPushBack(pSelf []PoolStringArray, pData []String) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpData, _ := unpackArgSString(pData)
	C.godot_pool_string_array_push_back(cpSelf, cpData)
	packSString(pData, cpData)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayRemove function as declared in gdnative/pool_arrays.h:273
func PoolStringArrayRemove(pSelf []PoolStringArray, pIdx Int) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_string_array_remove(cpSelf, cpIdx)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayResize function as declared in gdnative/pool_arrays.h:275
func PoolStringArrayResize(pSelf []PoolStringArray, pSize Int) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_string_array_resize(cpSelf, cpSize)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayRead function as declared in gdnative/pool_arrays.h:277
func PoolStringArrayRead(pSelf []PoolStringArray) *PoolStringArrayReadAccess {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_read(cpSelf)
	packSPoolStringArray(pSelf, cpSelf)
	__v := NewPoolStringArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolStringArrayWrite function as declared in gdnative/pool_arrays.h:279
func PoolStringArrayWrite(pSelf []PoolStringArray) *PoolStringArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_write(cpSelf)
	packSPoolStringArray(pSelf, cpSelf)
	__v := NewPoolStringArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolStringArraySet function as declared in gdnative/pool_arrays.h:281
func PoolStringArraySet(pSelf []PoolStringArray, pIdx Int, pData []String) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSString(pData)
	C.godot_pool_string_array_set(cpSelf, cpIdx, cpData)
	packSString(pData, cpData)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolStringArrayGet function as declared in gdnative/pool_arrays.h:282
func PoolStringArrayGet(pSelf []PoolStringArray, pIdx Int) String {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_string_array_get(cpSelf, cpIdx)
	packSPoolStringArray(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// PoolStringArraySize function as declared in gdnative/pool_arrays.h:284
func PoolStringArraySize(pSelf []PoolStringArray) Int {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	__ret := C.godot_pool_string_array_size(cpSelf)
	packSPoolStringArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolStringArrayDestroy function as declared in gdnative/pool_arrays.h:286
func PoolStringArrayDestroy(pSelf []PoolStringArray) {
	cpSelf, _ := unpackArgSPoolStringArray(pSelf)
	C.godot_pool_string_array_destroy(cpSelf)
	packSPoolStringArray(pSelf, cpSelf)
}

// PoolVector2ArrayNew function as declared in gdnative/pool_arrays.h:290
func PoolVector2ArrayNew(rDest []PoolVector2Array) {
	crDest, _ := unpackArgSPoolVector2Array(rDest)
	C.godot_pool_vector2_array_new(crDest)
	packSPoolVector2Array(rDest, crDest)
}

// PoolVector2ArrayNewCopy function as declared in gdnative/pool_arrays.h:291
func PoolVector2ArrayNewCopy(rDest []PoolVector2Array, pSrc []PoolVector2Array) {
	crDest, _ := unpackArgSPoolVector2Array(rDest)
	cpSrc, _ := unpackArgSPoolVector2Array(pSrc)
	C.godot_pool_vector2_array_new_copy(crDest, cpSrc)
	packSPoolVector2Array(pSrc, cpSrc)
	packSPoolVector2Array(rDest, crDest)
}

// PoolVector2ArrayNewWithArray function as declared in gdnative/pool_arrays.h:292
func PoolVector2ArrayNewWithArray(rDest []PoolVector2Array, pA []Array) {
	crDest, _ := unpackArgSPoolVector2Array(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_vector2_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolVector2Array(rDest, crDest)
}

// PoolVector2ArrayAppend function as declared in gdnative/pool_arrays.h:294
func PoolVector2ArrayAppend(pSelf []PoolVector2Array, pData []Vector2) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpData, _ := unpackArgSVector2(pData)
	C.godot_pool_vector2_array_append(cpSelf, cpData)
	packSVector2(pData, cpData)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayAppendArray function as declared in gdnative/pool_arrays.h:296
func PoolVector2ArrayAppendArray(pSelf []PoolVector2Array, pArray []PoolVector2Array) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpArray, _ := unpackArgSPoolVector2Array(pArray)
	C.godot_pool_vector2_array_append_array(cpSelf, cpArray)
	packSPoolVector2Array(pArray, cpArray)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayInsert function as declared in gdnative/pool_arrays.h:298
func PoolVector2ArrayInsert(pSelf []PoolVector2Array, pIdx Int, pData []Vector2) Error {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSVector2(pData)
	__ret := C.godot_pool_vector2_array_insert(cpSelf, cpIdx, cpData)
	packSVector2(pData, cpData)
	packSPoolVector2Array(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolVector2ArrayInvert function as declared in gdnative/pool_arrays.h:300
func PoolVector2ArrayInvert(pSelf []PoolVector2Array) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	C.godot_pool_vector2_array_invert(cpSelf)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayPushBack function as declared in gdnative/pool_arrays.h:302
func PoolVector2ArrayPushBack(pSelf []PoolVector2Array, pData []Vector2) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpData, _ := unpackArgSVector2(pData)
	C.godot_pool_vector2_array_push_back(cpSelf, cpData)
	packSVector2(pData, cpData)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayRemove function as declared in gdnative/pool_arrays.h:304
func PoolVector2ArrayRemove(pSelf []PoolVector2Array, pIdx Int) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_vector2_array_remove(cpSelf, cpIdx)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayResize function as declared in gdnative/pool_arrays.h:306
func PoolVector2ArrayResize(pSelf []PoolVector2Array, pSize Int) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_vector2_array_resize(cpSelf, cpSize)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayRead function as declared in gdnative/pool_arrays.h:308
func PoolVector2ArrayRead(pSelf []PoolVector2Array) *PoolVector2ArrayReadAccess {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_read(cpSelf)
	packSPoolVector2Array(pSelf, cpSelf)
	__v := NewPoolVector2ArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolVector2ArrayWrite function as declared in gdnative/pool_arrays.h:310
func PoolVector2ArrayWrite(pSelf []PoolVector2Array) *PoolVector2ArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_write(cpSelf)
	packSPoolVector2Array(pSelf, cpSelf)
	__v := NewPoolVector2ArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolVector2ArraySet function as declared in gdnative/pool_arrays.h:312
func PoolVector2ArraySet(pSelf []PoolVector2Array, pIdx Int, pData []Vector2) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSVector2(pData)
	C.godot_pool_vector2_array_set(cpSelf, cpIdx, cpData)
	packSVector2(pData, cpData)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector2ArrayGet function as declared in gdnative/pool_arrays.h:313
func PoolVector2ArrayGet(pSelf []PoolVector2Array, pIdx Int) Vector2 {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_vector2_array_get(cpSelf, cpIdx)
	packSPoolVector2Array(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// PoolVector2ArraySize function as declared in gdnative/pool_arrays.h:315
func PoolVector2ArraySize(pSelf []PoolVector2Array) Int {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	__ret := C.godot_pool_vector2_array_size(cpSelf)
	packSPoolVector2Array(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolVector2ArrayDestroy function as declared in gdnative/pool_arrays.h:317
func PoolVector2ArrayDestroy(pSelf []PoolVector2Array) {
	cpSelf, _ := unpackArgSPoolVector2Array(pSelf)
	C.godot_pool_vector2_array_destroy(cpSelf)
	packSPoolVector2Array(pSelf, cpSelf)
}

// PoolVector3ArrayNew function as declared in gdnative/pool_arrays.h:321
func PoolVector3ArrayNew(rDest []PoolVector3Array) {
	crDest, _ := unpackArgSPoolVector3Array(rDest)
	C.godot_pool_vector3_array_new(crDest)
	packSPoolVector3Array(rDest, crDest)
}

// PoolVector3ArrayNewCopy function as declared in gdnative/pool_arrays.h:322
func PoolVector3ArrayNewCopy(rDest []PoolVector3Array, pSrc []PoolVector3Array) {
	crDest, _ := unpackArgSPoolVector3Array(rDest)
	cpSrc, _ := unpackArgSPoolVector3Array(pSrc)
	C.godot_pool_vector3_array_new_copy(crDest, cpSrc)
	packSPoolVector3Array(pSrc, cpSrc)
	packSPoolVector3Array(rDest, crDest)
}

// PoolVector3ArrayNewWithArray function as declared in gdnative/pool_arrays.h:323
func PoolVector3ArrayNewWithArray(rDest []PoolVector3Array, pA []Array) {
	crDest, _ := unpackArgSPoolVector3Array(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_vector3_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolVector3Array(rDest, crDest)
}

// PoolVector3ArrayAppend function as declared in gdnative/pool_arrays.h:325
func PoolVector3ArrayAppend(pSelf []PoolVector3Array, pData []Vector3) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpData, _ := unpackArgSVector3(pData)
	C.godot_pool_vector3_array_append(cpSelf, cpData)
	packSVector3(pData, cpData)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayAppendArray function as declared in gdnative/pool_arrays.h:327
func PoolVector3ArrayAppendArray(pSelf []PoolVector3Array, pArray []PoolVector3Array) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpArray, _ := unpackArgSPoolVector3Array(pArray)
	C.godot_pool_vector3_array_append_array(cpSelf, cpArray)
	packSPoolVector3Array(pArray, cpArray)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayInsert function as declared in gdnative/pool_arrays.h:329
func PoolVector3ArrayInsert(pSelf []PoolVector3Array, pIdx Int, pData []Vector3) Error {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSVector3(pData)
	__ret := C.godot_pool_vector3_array_insert(cpSelf, cpIdx, cpData)
	packSVector3(pData, cpData)
	packSPoolVector3Array(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolVector3ArrayInvert function as declared in gdnative/pool_arrays.h:331
func PoolVector3ArrayInvert(pSelf []PoolVector3Array) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	C.godot_pool_vector3_array_invert(cpSelf)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayPushBack function as declared in gdnative/pool_arrays.h:333
func PoolVector3ArrayPushBack(pSelf []PoolVector3Array, pData []Vector3) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpData, _ := unpackArgSVector3(pData)
	C.godot_pool_vector3_array_push_back(cpSelf, cpData)
	packSVector3(pData, cpData)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayRemove function as declared in gdnative/pool_arrays.h:335
func PoolVector3ArrayRemove(pSelf []PoolVector3Array, pIdx Int) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_vector3_array_remove(cpSelf, cpIdx)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayResize function as declared in gdnative/pool_arrays.h:337
func PoolVector3ArrayResize(pSelf []PoolVector3Array, pSize Int) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_vector3_array_resize(cpSelf, cpSize)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayRead function as declared in gdnative/pool_arrays.h:339
func PoolVector3ArrayRead(pSelf []PoolVector3Array) *PoolVector3ArrayReadAccess {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_read(cpSelf)
	packSPoolVector3Array(pSelf, cpSelf)
	__v := NewPoolVector3ArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolVector3ArrayWrite function as declared in gdnative/pool_arrays.h:341
func PoolVector3ArrayWrite(pSelf []PoolVector3Array) *PoolVector3ArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_write(cpSelf)
	packSPoolVector3Array(pSelf, cpSelf)
	__v := NewPoolVector3ArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolVector3ArraySet function as declared in gdnative/pool_arrays.h:343
func PoolVector3ArraySet(pSelf []PoolVector3Array, pIdx Int, pData []Vector3) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSVector3(pData)
	C.godot_pool_vector3_array_set(cpSelf, cpIdx, cpData)
	packSVector3(pData, cpData)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolVector3ArrayGet function as declared in gdnative/pool_arrays.h:344
func PoolVector3ArrayGet(pSelf []PoolVector3Array, pIdx Int) Vector3 {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_vector3_array_get(cpSelf, cpIdx)
	packSPoolVector3Array(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// PoolVector3ArraySize function as declared in gdnative/pool_arrays.h:346
func PoolVector3ArraySize(pSelf []PoolVector3Array) Int {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	__ret := C.godot_pool_vector3_array_size(cpSelf)
	packSPoolVector3Array(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolVector3ArrayDestroy function as declared in gdnative/pool_arrays.h:348
func PoolVector3ArrayDestroy(pSelf []PoolVector3Array) {
	cpSelf, _ := unpackArgSPoolVector3Array(pSelf)
	C.godot_pool_vector3_array_destroy(cpSelf)
	packSPoolVector3Array(pSelf, cpSelf)
}

// PoolColorArrayNew function as declared in gdnative/pool_arrays.h:352
func PoolColorArrayNew(rDest []PoolColorArray) {
	crDest, _ := unpackArgSPoolColorArray(rDest)
	C.godot_pool_color_array_new(crDest)
	packSPoolColorArray(rDest, crDest)
}

// PoolColorArrayNewCopy function as declared in gdnative/pool_arrays.h:353
func PoolColorArrayNewCopy(rDest []PoolColorArray, pSrc []PoolColorArray) {
	crDest, _ := unpackArgSPoolColorArray(rDest)
	cpSrc, _ := unpackArgSPoolColorArray(pSrc)
	C.godot_pool_color_array_new_copy(crDest, cpSrc)
	packSPoolColorArray(pSrc, cpSrc)
	packSPoolColorArray(rDest, crDest)
}

// PoolColorArrayNewWithArray function as declared in gdnative/pool_arrays.h:354
func PoolColorArrayNewWithArray(rDest []PoolColorArray, pA []Array) {
	crDest, _ := unpackArgSPoolColorArray(rDest)
	cpA, _ := unpackArgSArray(pA)
	C.godot_pool_color_array_new_with_array(crDest, cpA)
	packSArray(pA, cpA)
	packSPoolColorArray(rDest, crDest)
}

// PoolColorArrayAppend function as declared in gdnative/pool_arrays.h:356
func PoolColorArrayAppend(pSelf []PoolColorArray, pData []Color) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpData, _ := unpackArgSColor(pData)
	C.godot_pool_color_array_append(cpSelf, cpData)
	packSColor(pData, cpData)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayAppendArray function as declared in gdnative/pool_arrays.h:358
func PoolColorArrayAppendArray(pSelf []PoolColorArray, pArray []PoolColorArray) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpArray, _ := unpackArgSPoolColorArray(pArray)
	C.godot_pool_color_array_append_array(cpSelf, cpArray)
	packSPoolColorArray(pArray, cpArray)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayInsert function as declared in gdnative/pool_arrays.h:360
func PoolColorArrayInsert(pSelf []PoolColorArray, pIdx Int, pData []Color) Error {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSColor(pData)
	__ret := C.godot_pool_color_array_insert(cpSelf, cpIdx, cpData)
	packSColor(pData, cpData)
	packSPoolColorArray(pSelf, cpSelf)
	__v := (Error)(__ret)
	return __v
}

// PoolColorArrayInvert function as declared in gdnative/pool_arrays.h:362
func PoolColorArrayInvert(pSelf []PoolColorArray) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	C.godot_pool_color_array_invert(cpSelf)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayPushBack function as declared in gdnative/pool_arrays.h:364
func PoolColorArrayPushBack(pSelf []PoolColorArray, pData []Color) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpData, _ := unpackArgSColor(pData)
	C.godot_pool_color_array_push_back(cpSelf, cpData)
	packSColor(pData, cpData)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayRemove function as declared in gdnative/pool_arrays.h:366
func PoolColorArrayRemove(pSelf []PoolColorArray, pIdx Int) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	C.godot_pool_color_array_remove(cpSelf, cpIdx)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayResize function as declared in gdnative/pool_arrays.h:368
func PoolColorArrayResize(pSelf []PoolColorArray, pSize Int) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpSize, _ := (C.godot_int)(pSize), cgoAllocsUnknown
	C.godot_pool_color_array_resize(cpSelf, cpSize)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayRead function as declared in gdnative/pool_arrays.h:370
func PoolColorArrayRead(pSelf []PoolColorArray) *PoolColorArrayReadAccess {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_read(cpSelf)
	packSPoolColorArray(pSelf, cpSelf)
	__v := NewPoolColorArrayReadAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolColorArrayWrite function as declared in gdnative/pool_arrays.h:372
func PoolColorArrayWrite(pSelf []PoolColorArray) *PoolColorArrayWriteAccess {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_write(cpSelf)
	packSPoolColorArray(pSelf, cpSelf)
	__v := NewPoolColorArrayWriteAccessRef(unsafe.Pointer(__ret))
	return __v
}

// PoolColorArraySet function as declared in gdnative/pool_arrays.h:374
func PoolColorArraySet(pSelf []PoolColorArray, pIdx Int, pData []Color) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	cpData, _ := unpackArgSColor(pData)
	C.godot_pool_color_array_set(cpSelf, cpIdx, cpData)
	packSColor(pData, cpData)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolColorArrayGet function as declared in gdnative/pool_arrays.h:375
func PoolColorArrayGet(pSelf []PoolColorArray, pIdx Int) Color {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_pool_color_array_get(cpSelf, cpIdx)
	packSPoolColorArray(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// PoolColorArraySize function as declared in gdnative/pool_arrays.h:377
func PoolColorArraySize(pSelf []PoolColorArray) Int {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	__ret := C.godot_pool_color_array_size(cpSelf)
	packSPoolColorArray(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// PoolColorArrayDestroy function as declared in gdnative/pool_arrays.h:379
func PoolColorArrayDestroy(pSelf []PoolColorArray) {
	cpSelf, _ := unpackArgSPoolColorArray(pSelf)
	C.godot_pool_color_array_destroy(cpSelf)
	packSPoolColorArray(pSelf, cpSelf)
}

// PoolByteArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:385
func PoolByteArrayReadAccessPtr(pRead []PoolByteArrayReadAccess) string {
	cpRead, _ := unpackArgSPoolByteArrayReadAccess(pRead)
	__ret := C.godot_pool_byte_array_read_access_ptr(cpRead)
	packSPoolByteArrayReadAccess(pRead, cpRead)
	__v := packPUint8TString(__ret)
	return __v
}

// PoolByteArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:386
func PoolByteArrayReadAccessOperatorAssign(pRead []PoolByteArrayReadAccess, pOther []PoolByteArrayReadAccess) {
	cpRead, _ := unpackArgSPoolByteArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolByteArrayReadAccess(pOther)
	C.godot_pool_byte_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolByteArrayReadAccess(pOther, cpOther)
	packSPoolByteArrayReadAccess(pRead, cpRead)
}

// PoolByteArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:387
func PoolByteArrayReadAccessDestroy(pRead []PoolByteArrayReadAccess) {
	cpRead, _ := unpackArgSPoolByteArrayReadAccess(pRead)
	C.godot_pool_byte_array_read_access_destroy(cpRead)
	packSPoolByteArrayReadAccess(pRead, cpRead)
}

// PoolIntArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:389
func PoolIntArrayReadAccessPtr(pRead []PoolIntArrayReadAccess) *Int {
	cpRead, _ := unpackArgSPoolIntArrayReadAccess(pRead)
	__ret := C.godot_pool_int_array_read_access_ptr(cpRead)
	packSPoolIntArrayReadAccess(pRead, cpRead)
	__v := *(**Int)(unsafe.Pointer(&__ret))
	return __v
}

// PoolIntArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:390
func PoolIntArrayReadAccessOperatorAssign(pRead []PoolIntArrayReadAccess, pOther []PoolIntArrayReadAccess) {
	cpRead, _ := unpackArgSPoolIntArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolIntArrayReadAccess(pOther)
	C.godot_pool_int_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolIntArrayReadAccess(pOther, cpOther)
	packSPoolIntArrayReadAccess(pRead, cpRead)
}

// PoolIntArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:391
func PoolIntArrayReadAccessDestroy(pRead []PoolIntArrayReadAccess) {
	cpRead, _ := unpackArgSPoolIntArrayReadAccess(pRead)
	C.godot_pool_int_array_read_access_destroy(cpRead)
	packSPoolIntArrayReadAccess(pRead, cpRead)
}

// PoolRealArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:393
func PoolRealArrayReadAccessPtr(pRead []PoolRealArrayReadAccess) *Real {
	cpRead, _ := unpackArgSPoolRealArrayReadAccess(pRead)
	__ret := C.godot_pool_real_array_read_access_ptr(cpRead)
	packSPoolRealArrayReadAccess(pRead, cpRead)
	__v := *(**Real)(unsafe.Pointer(&__ret))
	return __v
}

// PoolRealArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:394
func PoolRealArrayReadAccessOperatorAssign(pRead []PoolRealArrayReadAccess, pOther []PoolRealArrayReadAccess) {
	cpRead, _ := unpackArgSPoolRealArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolRealArrayReadAccess(pOther)
	C.godot_pool_real_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolRealArrayReadAccess(pOther, cpOther)
	packSPoolRealArrayReadAccess(pRead, cpRead)
}

// PoolRealArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:395
func PoolRealArrayReadAccessDestroy(pRead []PoolRealArrayReadAccess) {
	cpRead, _ := unpackArgSPoolRealArrayReadAccess(pRead)
	C.godot_pool_real_array_read_access_destroy(cpRead)
	packSPoolRealArrayReadAccess(pRead, cpRead)
}

// PoolStringArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:397
func PoolStringArrayReadAccessPtr(pRead []PoolStringArrayReadAccess) *String {
	cpRead, _ := unpackArgSPoolStringArrayReadAccess(pRead)
	__ret := C.godot_pool_string_array_read_access_ptr(cpRead)
	packSPoolStringArrayReadAccess(pRead, cpRead)
	__v := NewStringRef(unsafe.Pointer(__ret))
	return __v
}

// PoolStringArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:398
func PoolStringArrayReadAccessOperatorAssign(pRead []PoolStringArrayReadAccess, pOther []PoolStringArrayReadAccess) {
	cpRead, _ := unpackArgSPoolStringArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolStringArrayReadAccess(pOther)
	C.godot_pool_string_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolStringArrayReadAccess(pOther, cpOther)
	packSPoolStringArrayReadAccess(pRead, cpRead)
}

// PoolStringArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:399
func PoolStringArrayReadAccessDestroy(pRead []PoolStringArrayReadAccess) {
	cpRead, _ := unpackArgSPoolStringArrayReadAccess(pRead)
	C.godot_pool_string_array_read_access_destroy(cpRead)
	packSPoolStringArrayReadAccess(pRead, cpRead)
}

// PoolVector2ArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:401
func PoolVector2ArrayReadAccessPtr(pRead []PoolVector2ArrayReadAccess) *Vector2 {
	cpRead, _ := unpackArgSPoolVector2ArrayReadAccess(pRead)
	__ret := C.godot_pool_vector2_array_read_access_ptr(cpRead)
	packSPoolVector2ArrayReadAccess(pRead, cpRead)
	__v := NewVector2Ref(unsafe.Pointer(__ret))
	return __v
}

// PoolVector2ArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:402
func PoolVector2ArrayReadAccessOperatorAssign(pRead []PoolVector2ArrayReadAccess, pOther []PoolVector2ArrayReadAccess) {
	cpRead, _ := unpackArgSPoolVector2ArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolVector2ArrayReadAccess(pOther)
	C.godot_pool_vector2_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolVector2ArrayReadAccess(pOther, cpOther)
	packSPoolVector2ArrayReadAccess(pRead, cpRead)
}

// PoolVector2ArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:403
func PoolVector2ArrayReadAccessDestroy(pRead []PoolVector2ArrayReadAccess) {
	cpRead, _ := unpackArgSPoolVector2ArrayReadAccess(pRead)
	C.godot_pool_vector2_array_read_access_destroy(cpRead)
	packSPoolVector2ArrayReadAccess(pRead, cpRead)
}

// PoolVector3ArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:405
func PoolVector3ArrayReadAccessPtr(pRead []PoolVector3ArrayReadAccess) *Vector3 {
	cpRead, _ := unpackArgSPoolVector3ArrayReadAccess(pRead)
	__ret := C.godot_pool_vector3_array_read_access_ptr(cpRead)
	packSPoolVector3ArrayReadAccess(pRead, cpRead)
	__v := NewVector3Ref(unsafe.Pointer(__ret))
	return __v
}

// PoolVector3ArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:406
func PoolVector3ArrayReadAccessOperatorAssign(pRead []PoolVector3ArrayReadAccess, pOther []PoolVector3ArrayReadAccess) {
	cpRead, _ := unpackArgSPoolVector3ArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolVector3ArrayReadAccess(pOther)
	C.godot_pool_vector3_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolVector3ArrayReadAccess(pOther, cpOther)
	packSPoolVector3ArrayReadAccess(pRead, cpRead)
}

// PoolVector3ArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:407
func PoolVector3ArrayReadAccessDestroy(pRead []PoolVector3ArrayReadAccess) {
	cpRead, _ := unpackArgSPoolVector3ArrayReadAccess(pRead)
	C.godot_pool_vector3_array_read_access_destroy(cpRead)
	packSPoolVector3ArrayReadAccess(pRead, cpRead)
}

// PoolColorArrayReadAccessPtr function as declared in gdnative/pool_arrays.h:409
func PoolColorArrayReadAccessPtr(pRead []PoolColorArrayReadAccess) *Color {
	cpRead, _ := unpackArgSPoolColorArrayReadAccess(pRead)
	__ret := C.godot_pool_color_array_read_access_ptr(cpRead)
	packSPoolColorArrayReadAccess(pRead, cpRead)
	__v := NewColorRef(unsafe.Pointer(__ret))
	return __v
}

// PoolColorArrayReadAccessOperatorAssign function as declared in gdnative/pool_arrays.h:410
func PoolColorArrayReadAccessOperatorAssign(pRead []PoolColorArrayReadAccess, pOther []PoolColorArrayReadAccess) {
	cpRead, _ := unpackArgSPoolColorArrayReadAccess(pRead)
	cpOther, _ := unpackArgSPoolColorArrayReadAccess(pOther)
	C.godot_pool_color_array_read_access_operator_assign(cpRead, cpOther)
	packSPoolColorArrayReadAccess(pOther, cpOther)
	packSPoolColorArrayReadAccess(pRead, cpRead)
}

// PoolColorArrayReadAccessDestroy function as declared in gdnative/pool_arrays.h:411
func PoolColorArrayReadAccessDestroy(pRead []PoolColorArrayReadAccess) {
	cpRead, _ := unpackArgSPoolColorArrayReadAccess(pRead)
	C.godot_pool_color_array_read_access_destroy(cpRead)
	packSPoolColorArrayReadAccess(pRead, cpRead)
}

// PoolByteArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:417
func PoolByteArrayWriteAccessPtr(pWrite []PoolByteArrayWriteAccess) *byte {
	cpWrite, _ := unpackArgSPoolByteArrayWriteAccess(pWrite)
	__ret := C.godot_pool_byte_array_write_access_ptr(cpWrite)
	packSPoolByteArrayWriteAccess(pWrite, cpWrite)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// PoolByteArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:418
func PoolByteArrayWriteAccessOperatorAssign(pWrite []PoolByteArrayWriteAccess, pOther []PoolByteArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolByteArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolByteArrayWriteAccess(pOther)
	C.godot_pool_byte_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolByteArrayWriteAccess(pOther, cpOther)
	packSPoolByteArrayWriteAccess(pWrite, cpWrite)
}

// PoolByteArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:419
func PoolByteArrayWriteAccessDestroy(pWrite []PoolByteArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolByteArrayWriteAccess(pWrite)
	C.godot_pool_byte_array_write_access_destroy(cpWrite)
	packSPoolByteArrayWriteAccess(pWrite, cpWrite)
}

// PoolIntArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:421
func PoolIntArrayWriteAccessPtr(pWrite []PoolIntArrayWriteAccess) *Int {
	cpWrite, _ := unpackArgSPoolIntArrayWriteAccess(pWrite)
	__ret := C.godot_pool_int_array_write_access_ptr(cpWrite)
	packSPoolIntArrayWriteAccess(pWrite, cpWrite)
	__v := *(**Int)(unsafe.Pointer(&__ret))
	return __v
}

// PoolIntArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:422
func PoolIntArrayWriteAccessOperatorAssign(pWrite []PoolIntArrayWriteAccess, pOther []PoolIntArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolIntArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolIntArrayWriteAccess(pOther)
	C.godot_pool_int_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolIntArrayWriteAccess(pOther, cpOther)
	packSPoolIntArrayWriteAccess(pWrite, cpWrite)
}

// PoolIntArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:423
func PoolIntArrayWriteAccessDestroy(pWrite []PoolIntArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolIntArrayWriteAccess(pWrite)
	C.godot_pool_int_array_write_access_destroy(cpWrite)
	packSPoolIntArrayWriteAccess(pWrite, cpWrite)
}

// PoolRealArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:425
func PoolRealArrayWriteAccessPtr(pWrite []PoolRealArrayWriteAccess) *Real {
	cpWrite, _ := unpackArgSPoolRealArrayWriteAccess(pWrite)
	__ret := C.godot_pool_real_array_write_access_ptr(cpWrite)
	packSPoolRealArrayWriteAccess(pWrite, cpWrite)
	__v := *(**Real)(unsafe.Pointer(&__ret))
	return __v
}

// PoolRealArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:426
func PoolRealArrayWriteAccessOperatorAssign(pWrite []PoolRealArrayWriteAccess, pOther []PoolRealArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolRealArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolRealArrayWriteAccess(pOther)
	C.godot_pool_real_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolRealArrayWriteAccess(pOther, cpOther)
	packSPoolRealArrayWriteAccess(pWrite, cpWrite)
}

// PoolRealArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:427
func PoolRealArrayWriteAccessDestroy(pWrite []PoolRealArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolRealArrayWriteAccess(pWrite)
	C.godot_pool_real_array_write_access_destroy(cpWrite)
	packSPoolRealArrayWriteAccess(pWrite, cpWrite)
}

// PoolStringArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:429
func PoolStringArrayWriteAccessPtr(pWrite []PoolStringArrayWriteAccess) *String {
	cpWrite, _ := unpackArgSPoolStringArrayWriteAccess(pWrite)
	__ret := C.godot_pool_string_array_write_access_ptr(cpWrite)
	packSPoolStringArrayWriteAccess(pWrite, cpWrite)
	__v := NewStringRef(unsafe.Pointer(__ret))
	return __v
}

// PoolStringArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:430
func PoolStringArrayWriteAccessOperatorAssign(pWrite []PoolStringArrayWriteAccess, pOther []PoolStringArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolStringArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolStringArrayWriteAccess(pOther)
	C.godot_pool_string_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolStringArrayWriteAccess(pOther, cpOther)
	packSPoolStringArrayWriteAccess(pWrite, cpWrite)
}

// PoolStringArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:431
func PoolStringArrayWriteAccessDestroy(pWrite []PoolStringArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolStringArrayWriteAccess(pWrite)
	C.godot_pool_string_array_write_access_destroy(cpWrite)
	packSPoolStringArrayWriteAccess(pWrite, cpWrite)
}

// PoolVector2ArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:433
func PoolVector2ArrayWriteAccessPtr(pWrite []PoolVector2ArrayWriteAccess) *Vector2 {
	cpWrite, _ := unpackArgSPoolVector2ArrayWriteAccess(pWrite)
	__ret := C.godot_pool_vector2_array_write_access_ptr(cpWrite)
	packSPoolVector2ArrayWriteAccess(pWrite, cpWrite)
	__v := NewVector2Ref(unsafe.Pointer(__ret))
	return __v
}

// PoolVector2ArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:434
func PoolVector2ArrayWriteAccessOperatorAssign(pWrite []PoolVector2ArrayWriteAccess, pOther []PoolVector2ArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolVector2ArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolVector2ArrayWriteAccess(pOther)
	C.godot_pool_vector2_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolVector2ArrayWriteAccess(pOther, cpOther)
	packSPoolVector2ArrayWriteAccess(pWrite, cpWrite)
}

// PoolVector2ArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:435
func PoolVector2ArrayWriteAccessDestroy(pWrite []PoolVector2ArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolVector2ArrayWriteAccess(pWrite)
	C.godot_pool_vector2_array_write_access_destroy(cpWrite)
	packSPoolVector2ArrayWriteAccess(pWrite, cpWrite)
}

// PoolVector3ArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:437
func PoolVector3ArrayWriteAccessPtr(pWrite []PoolVector3ArrayWriteAccess) *Vector3 {
	cpWrite, _ := unpackArgSPoolVector3ArrayWriteAccess(pWrite)
	__ret := C.godot_pool_vector3_array_write_access_ptr(cpWrite)
	packSPoolVector3ArrayWriteAccess(pWrite, cpWrite)
	__v := NewVector3Ref(unsafe.Pointer(__ret))
	return __v
}

// PoolVector3ArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:438
func PoolVector3ArrayWriteAccessOperatorAssign(pWrite []PoolVector3ArrayWriteAccess, pOther []PoolVector3ArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolVector3ArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolVector3ArrayWriteAccess(pOther)
	C.godot_pool_vector3_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolVector3ArrayWriteAccess(pOther, cpOther)
	packSPoolVector3ArrayWriteAccess(pWrite, cpWrite)
}

// PoolVector3ArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:439
func PoolVector3ArrayWriteAccessDestroy(pWrite []PoolVector3ArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolVector3ArrayWriteAccess(pWrite)
	C.godot_pool_vector3_array_write_access_destroy(cpWrite)
	packSPoolVector3ArrayWriteAccess(pWrite, cpWrite)
}

// PoolColorArrayWriteAccessPtr function as declared in gdnative/pool_arrays.h:441
func PoolColorArrayWriteAccessPtr(pWrite []PoolColorArrayWriteAccess) *Color {
	cpWrite, _ := unpackArgSPoolColorArrayWriteAccess(pWrite)
	__ret := C.godot_pool_color_array_write_access_ptr(cpWrite)
	packSPoolColorArrayWriteAccess(pWrite, cpWrite)
	__v := NewColorRef(unsafe.Pointer(__ret))
	return __v
}

// PoolColorArrayWriteAccessOperatorAssign function as declared in gdnative/pool_arrays.h:442
func PoolColorArrayWriteAccessOperatorAssign(pWrite []PoolColorArrayWriteAccess, pOther []PoolColorArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolColorArrayWriteAccess(pWrite)
	cpOther, _ := unpackArgSPoolColorArrayWriteAccess(pOther)
	C.godot_pool_color_array_write_access_operator_assign(cpWrite, cpOther)
	packSPoolColorArrayWriteAccess(pOther, cpOther)
	packSPoolColorArrayWriteAccess(pWrite, cpWrite)
}

// PoolColorArrayWriteAccessDestroy function as declared in gdnative/pool_arrays.h:443
func PoolColorArrayWriteAccessDestroy(pWrite []PoolColorArrayWriteAccess) {
	cpWrite, _ := unpackArgSPoolColorArrayWriteAccess(pWrite)
	C.godot_pool_color_array_write_access_destroy(cpWrite)
	packSPoolColorArrayWriteAccess(pWrite, cpWrite)
}

// ColorNewRgba function as declared in gdnative/color.h:60
func ColorNewRgba(rDest []Color, pR Real, pG Real, pB Real, pA Real) {
	crDest, _ := unpackArgSColor(rDest)
	cpR, _ := (C.godot_real)(pR), cgoAllocsUnknown
	cpG, _ := (C.godot_real)(pG), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	cpA, _ := (C.godot_real)(pA), cgoAllocsUnknown
	C.godot_color_new_rgba(crDest, cpR, cpG, cpB, cpA)
	packSColor(rDest, crDest)
}

// ColorNewRgb function as declared in gdnative/color.h:61
func ColorNewRgb(rDest []Color, pR Real, pG Real, pB Real) {
	crDest, _ := unpackArgSColor(rDest)
	cpR, _ := (C.godot_real)(pR), cgoAllocsUnknown
	cpG, _ := (C.godot_real)(pG), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	C.godot_color_new_rgb(crDest, cpR, cpG, cpB)
	packSColor(rDest, crDest)
}

// ColorGetR function as declared in gdnative/color.h:63
func ColorGetR(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_r(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorSetR function as declared in gdnative/color.h:64
func ColorSetR(pSelf []Color, r Real) {
	cpSelf, _ := unpackArgSColor(pSelf)
	cr, _ := (C.godot_real)(r), cgoAllocsUnknown
	C.godot_color_set_r(cpSelf, cr)
	packSColor(pSelf, cpSelf)
}

// ColorGetG function as declared in gdnative/color.h:66
func ColorGetG(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_g(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorSetG function as declared in gdnative/color.h:67
func ColorSetG(pSelf []Color, g Real) {
	cpSelf, _ := unpackArgSColor(pSelf)
	cg, _ := (C.godot_real)(g), cgoAllocsUnknown
	C.godot_color_set_g(cpSelf, cg)
	packSColor(pSelf, cpSelf)
}

// ColorGetB function as declared in gdnative/color.h:69
func ColorGetB(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_b(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorSetB function as declared in gdnative/color.h:70
func ColorSetB(pSelf []Color, b Real) {
	cpSelf, _ := unpackArgSColor(pSelf)
	cb, _ := (C.godot_real)(b), cgoAllocsUnknown
	C.godot_color_set_b(cpSelf, cb)
	packSColor(pSelf, cpSelf)
}

// ColorGetA function as declared in gdnative/color.h:72
func ColorGetA(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_a(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorSetA function as declared in gdnative/color.h:73
func ColorSetA(pSelf []Color, a Real) {
	cpSelf, _ := unpackArgSColor(pSelf)
	ca, _ := (C.godot_real)(a), cgoAllocsUnknown
	C.godot_color_set_a(cpSelf, ca)
	packSColor(pSelf, cpSelf)
}

// ColorGetH function as declared in gdnative/color.h:75
func ColorGetH(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_h(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorGetS function as declared in gdnative/color.h:76
func ColorGetS(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_s(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorGetV function as declared in gdnative/color.h:77
func ColorGetV(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_get_v(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorAsString function as declared in gdnative/color.h:79
func ColorAsString(pSelf []Color) String {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_as_string(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorToRgba32 function as declared in gdnative/color.h:81
func ColorToRgba32(pSelf []Color) Int {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_to_rgba32(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ColorToArgb32 function as declared in gdnative/color.h:83
func ColorToArgb32(pSelf []Color) Int {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_to_argb32(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// ColorGray function as declared in gdnative/color.h:85
func ColorGray(pSelf []Color) Real {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_gray(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// ColorInverted function as declared in gdnative/color.h:87
func ColorInverted(pSelf []Color) Color {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_inverted(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorContrasted function as declared in gdnative/color.h:89
func ColorContrasted(pSelf []Color) Color {
	cpSelf, _ := unpackArgSColor(pSelf)
	__ret := C.godot_color_contrasted(cpSelf)
	packSColor(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorLinearInterpolate function as declared in gdnative/color.h:91
func ColorLinearInterpolate(pSelf []Color, pB []Color, pT Real) Color {
	cpSelf, _ := unpackArgSColor(pSelf)
	cpB, _ := unpackArgSColor(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_color_linear_interpolate(cpSelf, cpB, cpT)
	packSColor(pB, cpB)
	packSColor(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorBlend function as declared in gdnative/color.h:93
func ColorBlend(pSelf []Color, pOver []Color) Color {
	cpSelf, _ := unpackArgSColor(pSelf)
	cpOver, _ := unpackArgSColor(pOver)
	__ret := C.godot_color_blend(cpSelf, cpOver)
	packSColor(pOver, cpOver)
	packSColor(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorToHtml function as declared in gdnative/color.h:95
func ColorToHtml(pSelf []Color, pWithAlpha Bool) String {
	cpSelf, _ := unpackArgSColor(pSelf)
	cpWithAlpha, _ := (C.godot_bool)(pWithAlpha), cgoAllocsUnknown
	__ret := C.godot_color_to_html(cpSelf, cpWithAlpha)
	packSColor(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// ColorOperatorEqual function as declared in gdnative/color.h:97
func ColorOperatorEqual(pSelf []Color, pB []Color) Bool {
	cpSelf, _ := unpackArgSColor(pSelf)
	cpB, _ := unpackArgSColor(pB)
	__ret := C.godot_color_operator_equal(cpSelf, cpB)
	packSColor(pB, cpB)
	packSColor(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// ColorOperatorLess function as declared in gdnative/color.h:99
func ColorOperatorLess(pSelf []Color, pB []Color) Bool {
	cpSelf, _ := unpackArgSColor(pSelf)
	cpB, _ := unpackArgSColor(pB)
	__ret := C.godot_color_operator_less(cpSelf, cpB)
	packSColor(pB, cpB)
	packSColor(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector2New function as declared in gdnative/vector2.h:59
func Vector2New(rDest []Vector2, pX Real, pY Real) {
	crDest, _ := unpackArgSVector2(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	C.godot_vector2_new(crDest, cpX, cpY)
	packSVector2(rDest, crDest)
}

// Vector2AsString function as declared in gdnative/vector2.h:61
func Vector2AsString(pSelf []Vector2) String {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_as_string(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Normalized function as declared in gdnative/vector2.h:63
func Vector2Normalized(pSelf []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_normalized(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Length function as declared in gdnative/vector2.h:65
func Vector2Length(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_length(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2Angle function as declared in gdnative/vector2.h:67
func Vector2Angle(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_angle(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2LengthSquared function as declared in gdnative/vector2.h:69
func Vector2LengthSquared(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_length_squared(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2IsNormalized function as declared in gdnative/vector2.h:71
func Vector2IsNormalized(pSelf []Vector2) Bool {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_is_normalized(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector2DistanceTo function as declared in gdnative/vector2.h:73
func Vector2DistanceTo(pSelf []Vector2, pTo []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpTo, _ := unpackArgSVector2(pTo)
	__ret := C.godot_vector2_distance_to(cpSelf, cpTo)
	packSVector2(pTo, cpTo)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2DistanceSquaredTo function as declared in gdnative/vector2.h:75
func Vector2DistanceSquaredTo(pSelf []Vector2, pTo []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpTo, _ := unpackArgSVector2(pTo)
	__ret := C.godot_vector2_distance_squared_to(cpSelf, cpTo)
	packSVector2(pTo, cpTo)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2AngleTo function as declared in gdnative/vector2.h:77
func Vector2AngleTo(pSelf []Vector2, pTo []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpTo, _ := unpackArgSVector2(pTo)
	__ret := C.godot_vector2_angle_to(cpSelf, cpTo)
	packSVector2(pTo, cpTo)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2AngleToPoint function as declared in gdnative/vector2.h:79
func Vector2AngleToPoint(pSelf []Vector2, pTo []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpTo, _ := unpackArgSVector2(pTo)
	__ret := C.godot_vector2_angle_to_point(cpSelf, cpTo)
	packSVector2(pTo, cpTo)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2LinearInterpolate function as declared in gdnative/vector2.h:81
func Vector2LinearInterpolate(pSelf []Vector2, pB []Vector2, pT Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector2_linear_interpolate(cpSelf, cpB, cpT)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2CubicInterpolate function as declared in gdnative/vector2.h:83
func Vector2CubicInterpolate(pSelf []Vector2, pB []Vector2, pPreA []Vector2, pPostB []Vector2, pT Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	cpPreA, _ := unpackArgSVector2(pPreA)
	cpPostB, _ := unpackArgSVector2(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector2_cubic_interpolate(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSVector2(pPostB, cpPostB)
	packSVector2(pPreA, cpPreA)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Rotated function as declared in gdnative/vector2.h:85
func Vector2Rotated(pSelf []Vector2, pPhi Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_vector2_rotated(cpSelf, cpPhi)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Tangent function as declared in gdnative/vector2.h:87
func Vector2Tangent(pSelf []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_tangent(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Floor function as declared in gdnative/vector2.h:89
func Vector2Floor(pSelf []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_floor(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Snapped function as declared in gdnative/vector2.h:91
func Vector2Snapped(pSelf []Vector2, pBy []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpBy, _ := unpackArgSVector2(pBy)
	__ret := C.godot_vector2_snapped(cpSelf, cpBy)
	packSVector2(pBy, cpBy)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Aspect function as declared in gdnative/vector2.h:93
func Vector2Aspect(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_aspect(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2Dot function as declared in gdnative/vector2.h:95
func Vector2Dot(pSelf []Vector2, pWith []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpWith, _ := unpackArgSVector2(pWith)
	__ret := C.godot_vector2_dot(cpSelf, cpWith)
	packSVector2(pWith, cpWith)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2Slide function as declared in gdnative/vector2.h:97
func Vector2Slide(pSelf []Vector2, pN []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpN, _ := unpackArgSVector2(pN)
	__ret := C.godot_vector2_slide(cpSelf, cpN)
	packSVector2(pN, cpN)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Bounce function as declared in gdnative/vector2.h:99
func Vector2Bounce(pSelf []Vector2, pN []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpN, _ := unpackArgSVector2(pN)
	__ret := C.godot_vector2_bounce(cpSelf, cpN)
	packSVector2(pN, cpN)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Reflect function as declared in gdnative/vector2.h:101
func Vector2Reflect(pSelf []Vector2, pN []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpN, _ := unpackArgSVector2(pN)
	__ret := C.godot_vector2_reflect(cpSelf, cpN)
	packSVector2(pN, cpN)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Abs function as declared in gdnative/vector2.h:103
func Vector2Abs(pSelf []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_abs(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2Clamped function as declared in gdnative/vector2.h:105
func Vector2Clamped(pSelf []Vector2, pLength Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpLength, _ := (C.godot_real)(pLength), cgoAllocsUnknown
	__ret := C.godot_vector2_clamped(cpSelf, cpLength)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorAdd function as declared in gdnative/vector2.h:107
func Vector2OperatorAdd(pSelf []Vector2, pB []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_add(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorSubstract function as declared in gdnative/vector2.h:109
func Vector2OperatorSubstract(pSelf []Vector2, pB []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_substract(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorMultiplyVector function as declared in gdnative/vector2.h:111
func Vector2OperatorMultiplyVector(pSelf []Vector2, pB []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_multiply_vector(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorMultiplyScalar function as declared in gdnative/vector2.h:113
func Vector2OperatorMultiplyScalar(pSelf []Vector2, pB Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector2_operator_multiply_scalar(cpSelf, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorDivideVector function as declared in gdnative/vector2.h:115
func Vector2OperatorDivideVector(pSelf []Vector2, pB []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_divide_vector(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorDivideScalar function as declared in gdnative/vector2.h:117
func Vector2OperatorDivideScalar(pSelf []Vector2, pB Real) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector2_operator_divide_scalar(cpSelf, cpB)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2OperatorEqual function as declared in gdnative/vector2.h:119
func Vector2OperatorEqual(pSelf []Vector2, pB []Vector2) Bool {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_equal(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector2OperatorLess function as declared in gdnative/vector2.h:121
func Vector2OperatorLess(pSelf []Vector2, pB []Vector2) Bool {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpB, _ := unpackArgSVector2(pB)
	__ret := C.godot_vector2_operator_less(cpSelf, cpB)
	packSVector2(pB, cpB)
	packSVector2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector2OperatorNeg function as declared in gdnative/vector2.h:123
func Vector2OperatorNeg(pSelf []Vector2) Vector2 {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_operator_neg(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector2SetX function as declared in gdnative/vector2.h:125
func Vector2SetX(pSelf []Vector2, pX Real) {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	C.godot_vector2_set_x(cpSelf, cpX)
	packSVector2(pSelf, cpSelf)
}

// Vector2SetY function as declared in gdnative/vector2.h:127
func Vector2SetY(pSelf []Vector2, pY Real) {
	cpSelf, _ := unpackArgSVector2(pSelf)
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	C.godot_vector2_set_y(cpSelf, cpY)
	packSVector2(pSelf, cpSelf)
}

// Vector2GetX function as declared in gdnative/vector2.h:129
func Vector2GetX(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_get_x(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector2GetY function as declared in gdnative/vector2.h:131
func Vector2GetY(pSelf []Vector2) Real {
	cpSelf, _ := unpackArgSVector2(pSelf)
	__ret := C.godot_vector2_get_y(cpSelf)
	packSVector2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3New function as declared in gdnative/vector3.h:66
func Vector3New(rDest []Vector3, pX Real, pY Real, pZ Real) {
	crDest, _ := unpackArgSVector3(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpZ, _ := (C.godot_real)(pZ), cgoAllocsUnknown
	C.godot_vector3_new(crDest, cpX, cpY, cpZ)
	packSVector3(rDest, crDest)
}

// Vector3AsString function as declared in gdnative/vector3.h:68
func Vector3AsString(pSelf []Vector3) String {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_as_string(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// Vector3MinAxis function as declared in gdnative/vector3.h:70
func Vector3MinAxis(pSelf []Vector3) Int {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_min_axis(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// Vector3MaxAxis function as declared in gdnative/vector3.h:72
func Vector3MaxAxis(pSelf []Vector3) Int {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_max_axis(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// Vector3Length function as declared in gdnative/vector3.h:74
func Vector3Length(pSelf []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_length(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3LengthSquared function as declared in gdnative/vector3.h:76
func Vector3LengthSquared(pSelf []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_length_squared(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3IsNormalized function as declared in gdnative/vector3.h:78
func Vector3IsNormalized(pSelf []Vector3) Bool {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_is_normalized(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector3Normalized function as declared in gdnative/vector3.h:80
func Vector3Normalized(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_normalized(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Inverse function as declared in gdnative/vector3.h:82
func Vector3Inverse(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_inverse(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Snapped function as declared in gdnative/vector3.h:84
func Vector3Snapped(pSelf []Vector3, pBy []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpBy, _ := unpackArgSVector3(pBy)
	__ret := C.godot_vector3_snapped(cpSelf, cpBy)
	packSVector3(pBy, cpBy)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Rotated function as declared in gdnative/vector3.h:86
func Vector3Rotated(pSelf []Vector3, pAxis []Vector3, pPhi Real) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpAxis, _ := unpackArgSVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_vector3_rotated(cpSelf, cpAxis, cpPhi)
	packSVector3(pAxis, cpAxis)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3LinearInterpolate function as declared in gdnative/vector3.h:88
func Vector3LinearInterpolate(pSelf []Vector3, pB []Vector3, pT Real) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector3_linear_interpolate(cpSelf, cpB, cpT)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3CubicInterpolate function as declared in gdnative/vector3.h:90
func Vector3CubicInterpolate(pSelf []Vector3, pB []Vector3, pPreA []Vector3, pPostB []Vector3, pT Real) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	cpPreA, _ := unpackArgSVector3(pPreA)
	cpPostB, _ := unpackArgSVector3(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_vector3_cubic_interpolate(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSVector3(pPostB, cpPostB)
	packSVector3(pPreA, cpPreA)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Dot function as declared in gdnative/vector3.h:92
func Vector3Dot(pSelf []Vector3, pB []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_dot(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3Cross function as declared in gdnative/vector3.h:94
func Vector3Cross(pSelf []Vector3, pB []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_cross(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Outer function as declared in gdnative/vector3.h:96
func Vector3Outer(pSelf []Vector3, pB []Vector3) Basis {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_outer(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// Vector3ToDiagonalMatrix function as declared in gdnative/vector3.h:98
func Vector3ToDiagonalMatrix(pSelf []Vector3) Basis {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_to_diagonal_matrix(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Abs function as declared in gdnative/vector3.h:100
func Vector3Abs(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_abs(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Floor function as declared in gdnative/vector3.h:102
func Vector3Floor(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_floor(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Ceil function as declared in gdnative/vector3.h:104
func Vector3Ceil(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_ceil(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3DistanceTo function as declared in gdnative/vector3.h:106
func Vector3DistanceTo(pSelf []Vector3, pB []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_distance_to(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3DistanceSquaredTo function as declared in gdnative/vector3.h:108
func Vector3DistanceSquaredTo(pSelf []Vector3, pB []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_distance_squared_to(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3AngleTo function as declared in gdnative/vector3.h:110
func Vector3AngleTo(pSelf []Vector3, pTo []Vector3) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpTo, _ := unpackArgSVector3(pTo)
	__ret := C.godot_vector3_angle_to(cpSelf, cpTo)
	packSVector3(pTo, cpTo)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Vector3Slide function as declared in gdnative/vector3.h:112
func Vector3Slide(pSelf []Vector3, pN []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpN, _ := unpackArgSVector3(pN)
	__ret := C.godot_vector3_slide(cpSelf, cpN)
	packSVector3(pN, cpN)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Bounce function as declared in gdnative/vector3.h:114
func Vector3Bounce(pSelf []Vector3, pN []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpN, _ := unpackArgSVector3(pN)
	__ret := C.godot_vector3_bounce(cpSelf, cpN)
	packSVector3(pN, cpN)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3Reflect function as declared in gdnative/vector3.h:116
func Vector3Reflect(pSelf []Vector3, pN []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpN, _ := unpackArgSVector3(pN)
	__ret := C.godot_vector3_reflect(cpSelf, cpN)
	packSVector3(pN, cpN)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorAdd function as declared in gdnative/vector3.h:118
func Vector3OperatorAdd(pSelf []Vector3, pB []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_add(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorSubstract function as declared in gdnative/vector3.h:120
func Vector3OperatorSubstract(pSelf []Vector3, pB []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_substract(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorMultiplyVector function as declared in gdnative/vector3.h:122
func Vector3OperatorMultiplyVector(pSelf []Vector3, pB []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_multiply_vector(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorMultiplyScalar function as declared in gdnative/vector3.h:124
func Vector3OperatorMultiplyScalar(pSelf []Vector3, pB Real) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector3_operator_multiply_scalar(cpSelf, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorDivideVector function as declared in gdnative/vector3.h:126
func Vector3OperatorDivideVector(pSelf []Vector3, pB []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_divide_vector(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorDivideScalar function as declared in gdnative/vector3.h:128
func Vector3OperatorDivideScalar(pSelf []Vector3, pB Real) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_vector3_operator_divide_scalar(cpSelf, cpB)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3OperatorEqual function as declared in gdnative/vector3.h:130
func Vector3OperatorEqual(pSelf []Vector3, pB []Vector3) Bool {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_equal(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector3OperatorLess function as declared in gdnative/vector3.h:132
func Vector3OperatorLess(pSelf []Vector3, pB []Vector3) Bool {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpB, _ := unpackArgSVector3(pB)
	__ret := C.godot_vector3_operator_less(cpSelf, cpB)
	packSVector3(pB, cpB)
	packSVector3(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Vector3OperatorNeg function as declared in gdnative/vector3.h:134
func Vector3OperatorNeg(pSelf []Vector3) Vector3 {
	cpSelf, _ := unpackArgSVector3(pSelf)
	__ret := C.godot_vector3_operator_neg(cpSelf)
	packSVector3(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// Vector3SetAxis function as declared in gdnative/vector3.h:136
func Vector3SetAxis(pSelf []Vector3, pAxis Vector3Axis, pVal Real) {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpAxis, _ := (C.godot_vector3_axis)(pAxis), cgoAllocsUnknown
	cpVal, _ := (C.godot_real)(pVal), cgoAllocsUnknown
	C.godot_vector3_set_axis(cpSelf, cpAxis, cpVal)
	packSVector3(pSelf, cpSelf)
}

// Vector3GetAxis function as declared in gdnative/vector3.h:138
func Vector3GetAxis(pSelf []Vector3, pAxis Vector3Axis) Real {
	cpSelf, _ := unpackArgSVector3(pSelf)
	cpAxis, _ := (C.godot_vector3_axis)(pAxis), cgoAllocsUnknown
	__ret := C.godot_vector3_get_axis(cpSelf, cpAxis)
	packSVector3(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// BasisNewWithRows function as declared in gdnative/basis.h:61
func BasisNewWithRows(rDest []Basis, pXAxis []Vector3, pYAxis []Vector3, pZAxis []Vector3) {
	crDest, _ := unpackArgSBasis(rDest)
	cpXAxis, _ := unpackArgSVector3(pXAxis)
	cpYAxis, _ := unpackArgSVector3(pYAxis)
	cpZAxis, _ := unpackArgSVector3(pZAxis)
	C.godot_basis_new_with_rows(crDest, cpXAxis, cpYAxis, cpZAxis)
	packSVector3(pZAxis, cpZAxis)
	packSVector3(pYAxis, cpYAxis)
	packSVector3(pXAxis, cpXAxis)
	packSBasis(rDest, crDest)
}

// BasisNewWithAxisAndAngle function as declared in gdnative/basis.h:62
func BasisNewWithAxisAndAngle(rDest []Basis, pAxis []Vector3, pPhi Real) {
	crDest, _ := unpackArgSBasis(rDest)
	cpAxis, _ := unpackArgSVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	C.godot_basis_new_with_axis_and_angle(crDest, cpAxis, cpPhi)
	packSVector3(pAxis, cpAxis)
	packSBasis(rDest, crDest)
}

// BasisNewWithEuler function as declared in gdnative/basis.h:63
func BasisNewWithEuler(rDest []Basis, pEuler []Vector3) {
	crDest, _ := unpackArgSBasis(rDest)
	cpEuler, _ := unpackArgSVector3(pEuler)
	C.godot_basis_new_with_euler(crDest, cpEuler)
	packSVector3(pEuler, cpEuler)
	packSBasis(rDest, crDest)
}

// BasisAsString function as declared in gdnative/basis.h:65
func BasisAsString(pSelf []Basis) String {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_as_string(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisInverse function as declared in gdnative/basis.h:67
func BasisInverse(pSelf []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_inverse(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisTransposed function as declared in gdnative/basis.h:69
func BasisTransposed(pSelf []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_transposed(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisOrthonormalized function as declared in gdnative/basis.h:71
func BasisOrthonormalized(pSelf []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_orthonormalized(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisDeterminant function as declared in gdnative/basis.h:73
func BasisDeterminant(pSelf []Basis) Real {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_determinant(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// BasisRotated function as declared in gdnative/basis.h:75
func BasisRotated(pSelf []Basis, pAxis []Vector3, pPhi Real) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpAxis, _ := unpackArgSVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_basis_rotated(cpSelf, cpAxis, cpPhi)
	packSVector3(pAxis, cpAxis)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisScaled function as declared in gdnative/basis.h:77
func BasisScaled(pSelf []Basis, pScale []Vector3) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpScale, _ := unpackArgSVector3(pScale)
	__ret := C.godot_basis_scaled(cpSelf, cpScale)
	packSVector3(pScale, cpScale)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisGetScale function as declared in gdnative/basis.h:79
func BasisGetScale(pSelf []Basis) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_get_scale(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisGetEuler function as declared in gdnative/basis.h:81
func BasisGetEuler(pSelf []Basis) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_get_euler(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisTdotx function as declared in gdnative/basis.h:83
func BasisTdotx(pSelf []Basis, pWith []Vector3) Real {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpWith, _ := unpackArgSVector3(pWith)
	__ret := C.godot_basis_tdotx(cpSelf, cpWith)
	packSVector3(pWith, cpWith)
	packSBasis(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// BasisTdoty function as declared in gdnative/basis.h:85
func BasisTdoty(pSelf []Basis, pWith []Vector3) Real {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpWith, _ := unpackArgSVector3(pWith)
	__ret := C.godot_basis_tdoty(cpSelf, cpWith)
	packSVector3(pWith, cpWith)
	packSBasis(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// BasisTdotz function as declared in gdnative/basis.h:87
func BasisTdotz(pSelf []Basis, pWith []Vector3) Real {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpWith, _ := unpackArgSVector3(pWith)
	__ret := C.godot_basis_tdotz(cpSelf, cpWith)
	packSVector3(pWith, cpWith)
	packSBasis(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// BasisXform function as declared in gdnative/basis.h:89
func BasisXform(pSelf []Basis, pV []Vector3) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	__ret := C.godot_basis_xform(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisXformInv function as declared in gdnative/basis.h:91
func BasisXformInv(pSelf []Basis, pV []Vector3) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	__ret := C.godot_basis_xform_inv(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisGetOrthogonalIndex function as declared in gdnative/basis.h:93
func BasisGetOrthogonalIndex(pSelf []Basis) Int {
	cpSelf, _ := unpackArgSBasis(pSelf)
	__ret := C.godot_basis_get_orthogonal_index(cpSelf)
	packSBasis(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// BasisNew function as declared in gdnative/basis.h:95
func BasisNew(rDest []Basis) {
	crDest, _ := unpackArgSBasis(rDest)
	C.godot_basis_new(crDest)
	packSBasis(rDest, crDest)
}

// BasisNewWithEulerQuat function as declared in gdnative/basis.h:97
func BasisNewWithEulerQuat(rDest []Basis, pEuler []Quat) {
	crDest, _ := unpackArgSBasis(rDest)
	cpEuler, _ := unpackArgSQuat(pEuler)
	C.godot_basis_new_with_euler_quat(crDest, cpEuler)
	packSQuat(pEuler, cpEuler)
	packSBasis(rDest, crDest)
}

// BasisGetElements function as declared in gdnative/basis.h:100
func BasisGetElements(pSelf []Basis, pElements []Vector3) {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpElements, _ := unpackArgSVector3(pElements)
	C.godot_basis_get_elements(cpSelf, cpElements)
	packSVector3(pElements, cpElements)
	packSBasis(pSelf, cpSelf)
}

// BasisGetAxis function as declared in gdnative/basis.h:102
func BasisGetAxis(pSelf []Basis, pAxis Int) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	__ret := C.godot_basis_get_axis(cpSelf, cpAxis)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisSetAxis function as declared in gdnative/basis.h:104
func BasisSetAxis(pSelf []Basis, pAxis Int, pValue []Vector3) {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	cpValue, _ := unpackArgSVector3(pValue)
	C.godot_basis_set_axis(cpSelf, cpAxis, cpValue)
	packSVector3(pValue, cpValue)
	packSBasis(pSelf, cpSelf)
}

// BasisGetRow function as declared in gdnative/basis.h:106
func BasisGetRow(pSelf []Basis, pRow Int) Vector3 {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpRow, _ := (C.godot_int)(pRow), cgoAllocsUnknown
	__ret := C.godot_basis_get_row(cpSelf, cpRow)
	packSBasis(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// BasisSetRow function as declared in gdnative/basis.h:108
func BasisSetRow(pSelf []Basis, pRow Int, pValue []Vector3) {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpRow, _ := (C.godot_int)(pRow), cgoAllocsUnknown
	cpValue, _ := unpackArgSVector3(pValue)
	C.godot_basis_set_row(cpSelf, cpRow, cpValue)
	packSVector3(pValue, cpValue)
	packSBasis(pSelf, cpSelf)
}

// BasisOperatorEqual function as declared in gdnative/basis.h:110
func BasisOperatorEqual(pSelf []Basis, pB []Basis) Bool {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpB, _ := unpackArgSBasis(pB)
	__ret := C.godot_basis_operator_equal(cpSelf, cpB)
	packSBasis(pB, cpB)
	packSBasis(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// BasisOperatorAdd function as declared in gdnative/basis.h:112
func BasisOperatorAdd(pSelf []Basis, pB []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpB, _ := unpackArgSBasis(pB)
	__ret := C.godot_basis_operator_add(cpSelf, cpB)
	packSBasis(pB, cpB)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisOperatorSubstract function as declared in gdnative/basis.h:114
func BasisOperatorSubstract(pSelf []Basis, pB []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpB, _ := unpackArgSBasis(pB)
	__ret := C.godot_basis_operator_substract(cpSelf, cpB)
	packSBasis(pB, cpB)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisOperatorMultiplyVector function as declared in gdnative/basis.h:116
func BasisOperatorMultiplyVector(pSelf []Basis, pB []Basis) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpB, _ := unpackArgSBasis(pB)
	__ret := C.godot_basis_operator_multiply_vector(cpSelf, cpB)
	packSBasis(pB, cpB)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// BasisOperatorMultiplyScalar function as declared in gdnative/basis.h:118
func BasisOperatorMultiplyScalar(pSelf []Basis, pB Real) Basis {
	cpSelf, _ := unpackArgSBasis(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_basis_operator_multiply_scalar(cpSelf, cpB)
	packSBasis(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatNew function as declared in gdnative/quat.h:60
func QuatNew(rDest []Quat, pX Real, pY Real, pZ Real, pW Real) {
	crDest, _ := unpackArgSQuat(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpZ, _ := (C.godot_real)(pZ), cgoAllocsUnknown
	cpW, _ := (C.godot_real)(pW), cgoAllocsUnknown
	C.godot_quat_new(crDest, cpX, cpY, cpZ, cpW)
	packSQuat(rDest, crDest)
}

// QuatNewWithAxisAngle function as declared in gdnative/quat.h:61
func QuatNewWithAxisAngle(rDest []Quat, pAxis []Vector3, pAngle Real) {
	crDest, _ := unpackArgSQuat(rDest)
	cpAxis, _ := unpackArgSVector3(pAxis)
	cpAngle, _ := (C.godot_real)(pAngle), cgoAllocsUnknown
	C.godot_quat_new_with_axis_angle(crDest, cpAxis, cpAngle)
	packSVector3(pAxis, cpAxis)
	packSQuat(rDest, crDest)
}

// QuatGetX function as declared in gdnative/quat.h:63
func QuatGetX(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_get_x(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatSetX function as declared in gdnative/quat.h:64
func QuatSetX(pSelf []Quat, val Real) {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_x(cpSelf, cval)
	packSQuat(pSelf, cpSelf)
}

// QuatGetY function as declared in gdnative/quat.h:66
func QuatGetY(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_get_y(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatSetY function as declared in gdnative/quat.h:67
func QuatSetY(pSelf []Quat, val Real) {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_y(cpSelf, cval)
	packSQuat(pSelf, cpSelf)
}

// QuatGetZ function as declared in gdnative/quat.h:69
func QuatGetZ(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_get_z(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatSetZ function as declared in gdnative/quat.h:70
func QuatSetZ(pSelf []Quat, val Real) {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_z(cpSelf, cval)
	packSQuat(pSelf, cpSelf)
}

// QuatGetW function as declared in gdnative/quat.h:72
func QuatGetW(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_get_w(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatSetW function as declared in gdnative/quat.h:73
func QuatSetW(pSelf []Quat, val Real) {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cval, _ := (C.godot_real)(val), cgoAllocsUnknown
	C.godot_quat_set_w(cpSelf, cval)
	packSQuat(pSelf, cpSelf)
}

// QuatAsString function as declared in gdnative/quat.h:75
func QuatAsString(pSelf []Quat) String {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_as_string(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatLength function as declared in gdnative/quat.h:77
func QuatLength(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_length(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatLengthSquared function as declared in gdnative/quat.h:79
func QuatLengthSquared(pSelf []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_length_squared(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatNormalized function as declared in gdnative/quat.h:81
func QuatNormalized(pSelf []Quat) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_normalized(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatIsNormalized function as declared in gdnative/quat.h:83
func QuatIsNormalized(pSelf []Quat) Bool {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_is_normalized(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// QuatInverse function as declared in gdnative/quat.h:85
func QuatInverse(pSelf []Quat) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_inverse(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatDot function as declared in gdnative/quat.h:87
func QuatDot(pSelf []Quat, pB []Quat) Real {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	__ret := C.godot_quat_dot(cpSelf, cpB)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// QuatXform function as declared in gdnative/quat.h:89
func QuatXform(pSelf []Quat, pV []Vector3) Vector3 {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	__ret := C.godot_quat_xform(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSQuat(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// QuatSlerp function as declared in gdnative/quat.h:91
func QuatSlerp(pSelf []Quat, pB []Quat, pT Real) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_slerp(cpSelf, cpB, cpT)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatSlerpni function as declared in gdnative/quat.h:93
func QuatSlerpni(pSelf []Quat, pB []Quat, pT Real) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_slerpni(cpSelf, cpB, cpT)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatCubicSlerp function as declared in gdnative/quat.h:95
func QuatCubicSlerp(pSelf []Quat, pB []Quat, pPreA []Quat, pPostB []Quat, pT Real) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	cpPreA, _ := unpackArgSQuat(pPreA)
	cpPostB, _ := unpackArgSQuat(pPostB)
	cpT, _ := (C.godot_real)(pT), cgoAllocsUnknown
	__ret := C.godot_quat_cubic_slerp(cpSelf, cpB, cpPreA, cpPostB, cpT)
	packSQuat(pPostB, cpPostB)
	packSQuat(pPreA, cpPreA)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatOperatorMultiply function as declared in gdnative/quat.h:97
func QuatOperatorMultiply(pSelf []Quat, pB Real) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_quat_operator_multiply(cpSelf, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatOperatorAdd function as declared in gdnative/quat.h:99
func QuatOperatorAdd(pSelf []Quat, pB []Quat) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	__ret := C.godot_quat_operator_add(cpSelf, cpB)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatOperatorSubstract function as declared in gdnative/quat.h:101
func QuatOperatorSubstract(pSelf []Quat, pB []Quat) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	__ret := C.godot_quat_operator_substract(cpSelf, cpB)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatOperatorDivide function as declared in gdnative/quat.h:103
func QuatOperatorDivide(pSelf []Quat, pB Real) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	__ret := C.godot_quat_operator_divide(cpSelf, cpB)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// QuatOperatorEqual function as declared in gdnative/quat.h:105
func QuatOperatorEqual(pSelf []Quat, pB []Quat) Bool {
	cpSelf, _ := unpackArgSQuat(pSelf)
	cpB, _ := unpackArgSQuat(pB)
	__ret := C.godot_quat_operator_equal(cpSelf, cpB)
	packSQuat(pB, cpB)
	packSQuat(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// QuatOperatorNeg function as declared in gdnative/quat.h:107
func QuatOperatorNeg(pSelf []Quat) Quat {
	cpSelf, _ := unpackArgSQuat(pSelf)
	__ret := C.godot_quat_operator_neg(cpSelf)
	packSQuat(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantGetType function as declared in gdnative/variant.h:131
func VariantGetType(pV []Variant) VariantType {
	cpV, _ := unpackArgSVariant(pV)
	__ret := C.godot_variant_get_type(cpV)
	packSVariant(pV, cpV)
	__v := (VariantType)(__ret)
	return __v
}

// VariantNewCopy function as declared in gdnative/variant.h:133
func VariantNewCopy(rDest []Variant, pSrc []Variant) {
	crDest, _ := unpackArgSVariant(rDest)
	cpSrc, _ := unpackArgSVariant(pSrc)
	C.godot_variant_new_copy(crDest, cpSrc)
	packSVariant(pSrc, cpSrc)
	packSVariant(rDest, crDest)
}

// VariantNewNil function as declared in gdnative/variant.h:135
func VariantNewNil(rDest []Variant) {
	crDest, _ := unpackArgSVariant(rDest)
	C.godot_variant_new_nil(crDest)
	packSVariant(rDest, crDest)
}

// VariantNewBool function as declared in gdnative/variant.h:137
func VariantNewBool(pV []Variant, pB Bool) {
	cpV, _ := unpackArgSVariant(pV)
	cpB, _ := (C.godot_bool)(pB), cgoAllocsUnknown
	C.godot_variant_new_bool(cpV, cpB)
	packSVariant(pV, cpV)
}

// VariantNewUint function as declared in gdnative/variant.h:138
func VariantNewUint(rDest []Variant, pI uint64) {
	crDest, _ := unpackArgSVariant(rDest)
	cpI, _ := (C.uint64_t)(pI), cgoAllocsUnknown
	C.godot_variant_new_uint(crDest, cpI)
	packSVariant(rDest, crDest)
}

// VariantNewInt function as declared in gdnative/variant.h:139
func VariantNewInt(rDest []Variant, pI int64) {
	crDest, _ := unpackArgSVariant(rDest)
	cpI, _ := (C.int64_t)(pI), cgoAllocsUnknown
	C.godot_variant_new_int(crDest, cpI)
	packSVariant(rDest, crDest)
}

// VariantNewReal function as declared in gdnative/variant.h:140
func VariantNewReal(rDest []Variant, pR float64) {
	crDest, _ := unpackArgSVariant(rDest)
	cpR, _ := (C.double)(pR), cgoAllocsUnknown
	C.godot_variant_new_real(crDest, cpR)
	packSVariant(rDest, crDest)
}

// VariantNewString function as declared in gdnative/variant.h:141
func VariantNewString(rDest []Variant, pS []String) {
	crDest, _ := unpackArgSVariant(rDest)
	cpS, _ := unpackArgSString(pS)
	C.godot_variant_new_string(crDest, cpS)
	packSString(pS, cpS)
	packSVariant(rDest, crDest)
}

// VariantNewVector2 function as declared in gdnative/variant.h:142
func VariantNewVector2(rDest []Variant, pV2 []Vector2) {
	crDest, _ := unpackArgSVariant(rDest)
	cpV2, _ := unpackArgSVector2(pV2)
	C.godot_variant_new_vector2(crDest, cpV2)
	packSVector2(pV2, cpV2)
	packSVariant(rDest, crDest)
}

// VariantNewRect2 function as declared in gdnative/variant.h:143
func VariantNewRect2(rDest []Variant, pRect2 []Rect2) {
	crDest, _ := unpackArgSVariant(rDest)
	cpRect2, _ := unpackArgSRect2(pRect2)
	C.godot_variant_new_rect2(crDest, cpRect2)
	packSRect2(pRect2, cpRect2)
	packSVariant(rDest, crDest)
}

// VariantNewVector3 function as declared in gdnative/variant.h:144
func VariantNewVector3(rDest []Variant, pV3 []Vector3) {
	crDest, _ := unpackArgSVariant(rDest)
	cpV3, _ := unpackArgSVector3(pV3)
	C.godot_variant_new_vector3(crDest, cpV3)
	packSVector3(pV3, cpV3)
	packSVariant(rDest, crDest)
}

// VariantNewTransform2d function as declared in gdnative/variant.h:145
func VariantNewTransform2d(rDest []Variant, pT2d []Transform2d) {
	crDest, _ := unpackArgSVariant(rDest)
	cpT2d, _ := unpackArgSTransform2d(pT2d)
	C.godot_variant_new_transform2d(crDest, cpT2d)
	packSTransform2d(pT2d, cpT2d)
	packSVariant(rDest, crDest)
}

// VariantNewPlane function as declared in gdnative/variant.h:146
func VariantNewPlane(rDest []Variant, pPlane []Plane) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPlane, _ := unpackArgSPlane(pPlane)
	C.godot_variant_new_plane(crDest, cpPlane)
	packSPlane(pPlane, cpPlane)
	packSVariant(rDest, crDest)
}

// VariantNewQuat function as declared in gdnative/variant.h:147
func VariantNewQuat(rDest []Variant, pQuat []Quat) {
	crDest, _ := unpackArgSVariant(rDest)
	cpQuat, _ := unpackArgSQuat(pQuat)
	C.godot_variant_new_quat(crDest, cpQuat)
	packSQuat(pQuat, cpQuat)
	packSVariant(rDest, crDest)
}

// VariantNewAabb function as declared in gdnative/variant.h:148
func VariantNewAabb(rDest []Variant, pAabb []Aabb) {
	crDest, _ := unpackArgSVariant(rDest)
	cpAabb, _ := unpackArgSAabb(pAabb)
	C.godot_variant_new_aabb(crDest, cpAabb)
	packSAabb(pAabb, cpAabb)
	packSVariant(rDest, crDest)
}

// VariantNewBasis function as declared in gdnative/variant.h:149
func VariantNewBasis(rDest []Variant, pBasis []Basis) {
	crDest, _ := unpackArgSVariant(rDest)
	cpBasis, _ := unpackArgSBasis(pBasis)
	C.godot_variant_new_basis(crDest, cpBasis)
	packSBasis(pBasis, cpBasis)
	packSVariant(rDest, crDest)
}

// VariantNewTransform function as declared in gdnative/variant.h:150
func VariantNewTransform(rDest []Variant, pTrans []Transform) {
	crDest, _ := unpackArgSVariant(rDest)
	cpTrans, _ := unpackArgSTransform(pTrans)
	C.godot_variant_new_transform(crDest, cpTrans)
	packSTransform(pTrans, cpTrans)
	packSVariant(rDest, crDest)
}

// VariantNewColor function as declared in gdnative/variant.h:151
func VariantNewColor(rDest []Variant, pColor []Color) {
	crDest, _ := unpackArgSVariant(rDest)
	cpColor, _ := unpackArgSColor(pColor)
	C.godot_variant_new_color(crDest, cpColor)
	packSColor(pColor, cpColor)
	packSVariant(rDest, crDest)
}

// VariantNewNodePath function as declared in gdnative/variant.h:152
func VariantNewNodePath(rDest []Variant, pNp []NodePath) {
	crDest, _ := unpackArgSVariant(rDest)
	cpNp, _ := unpackArgSNodePath(pNp)
	C.godot_variant_new_node_path(crDest, cpNp)
	packSNodePath(pNp, cpNp)
	packSVariant(rDest, crDest)
}

// VariantNewRid function as declared in gdnative/variant.h:153
func VariantNewRid(rDest []Variant, pRid []Rid) {
	crDest, _ := unpackArgSVariant(rDest)
	cpRid, _ := unpackArgSRid(pRid)
	C.godot_variant_new_rid(crDest, cpRid)
	packSRid(pRid, cpRid)
	packSVariant(rDest, crDest)
}

// VariantNewObject function as declared in gdnative/variant.h:154
func VariantNewObject(rDest []Variant, pObj *Object) {
	crDest, _ := unpackArgSVariant(rDest)
	cpObj, _ := (C.godot_object)(unsafe.Pointer(pObj)), cgoAllocsUnknown
	C.godot_variant_new_object(crDest, cpObj)
	packSVariant(rDest, crDest)
}

// VariantNewDictionary function as declared in gdnative/variant.h:155
func VariantNewDictionary(rDest []Variant, pDict []Dictionary) {
	crDest, _ := unpackArgSVariant(rDest)
	cpDict, _ := unpackArgSDictionary(pDict)
	C.godot_variant_new_dictionary(crDest, cpDict)
	packSDictionary(pDict, cpDict)
	packSVariant(rDest, crDest)
}

// VariantNewArray function as declared in gdnative/variant.h:156
func VariantNewArray(rDest []Variant, pArr []Array) {
	crDest, _ := unpackArgSVariant(rDest)
	cpArr, _ := unpackArgSArray(pArr)
	C.godot_variant_new_array(crDest, cpArr)
	packSArray(pArr, cpArr)
	packSVariant(rDest, crDest)
}

// VariantNewPoolByteArray function as declared in gdnative/variant.h:157
func VariantNewPoolByteArray(rDest []Variant, pPba []PoolByteArray) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPba, _ := unpackArgSPoolByteArray(pPba)
	C.godot_variant_new_pool_byte_array(crDest, cpPba)
	packSPoolByteArray(pPba, cpPba)
	packSVariant(rDest, crDest)
}

// VariantNewPoolIntArray function as declared in gdnative/variant.h:158
func VariantNewPoolIntArray(rDest []Variant, pPia []PoolIntArray) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPia, _ := unpackArgSPoolIntArray(pPia)
	C.godot_variant_new_pool_int_array(crDest, cpPia)
	packSPoolIntArray(pPia, cpPia)
	packSVariant(rDest, crDest)
}

// VariantNewPoolRealArray function as declared in gdnative/variant.h:159
func VariantNewPoolRealArray(rDest []Variant, pPra []PoolRealArray) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPra, _ := unpackArgSPoolRealArray(pPra)
	C.godot_variant_new_pool_real_array(crDest, cpPra)
	packSPoolRealArray(pPra, cpPra)
	packSVariant(rDest, crDest)
}

// VariantNewPoolStringArray function as declared in gdnative/variant.h:160
func VariantNewPoolStringArray(rDest []Variant, pPsa []PoolStringArray) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPsa, _ := unpackArgSPoolStringArray(pPsa)
	C.godot_variant_new_pool_string_array(crDest, cpPsa)
	packSPoolStringArray(pPsa, cpPsa)
	packSVariant(rDest, crDest)
}

// VariantNewPoolVector2Array function as declared in gdnative/variant.h:161
func VariantNewPoolVector2Array(rDest []Variant, pPv2a []PoolVector2Array) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPv2a, _ := unpackArgSPoolVector2Array(pPv2a)
	C.godot_variant_new_pool_vector2_array(crDest, cpPv2a)
	packSPoolVector2Array(pPv2a, cpPv2a)
	packSVariant(rDest, crDest)
}

// VariantNewPoolVector3Array function as declared in gdnative/variant.h:162
func VariantNewPoolVector3Array(rDest []Variant, pPv3a []PoolVector3Array) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPv3a, _ := unpackArgSPoolVector3Array(pPv3a)
	C.godot_variant_new_pool_vector3_array(crDest, cpPv3a)
	packSPoolVector3Array(pPv3a, cpPv3a)
	packSVariant(rDest, crDest)
}

// VariantNewPoolColorArray function as declared in gdnative/variant.h:163
func VariantNewPoolColorArray(rDest []Variant, pPca []PoolColorArray) {
	crDest, _ := unpackArgSVariant(rDest)
	cpPca, _ := unpackArgSPoolColorArray(pPca)
	C.godot_variant_new_pool_color_array(crDest, cpPca)
	packSPoolColorArray(pPca, cpPca)
	packSVariant(rDest, crDest)
}

// VariantAsBool function as declared in gdnative/variant.h:165
func VariantAsBool(pSelf []Variant) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_bool(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantAsUint function as declared in gdnative/variant.h:166
func VariantAsUint(pSelf []Variant) uint64 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_uint(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := (uint64)(__ret)
	return __v
}

// VariantAsInt function as declared in gdnative/variant.h:167
func VariantAsInt(pSelf []Variant) int64 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_int(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := (int64)(__ret)
	return __v
}

// VariantAsReal function as declared in gdnative/variant.h:168
func VariantAsReal(pSelf []Variant) float64 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_real(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := (float64)(__ret)
	return __v
}

// VariantAsString function as declared in gdnative/variant.h:169
func VariantAsString(pSelf []Variant) String {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_string(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsVector2 function as declared in gdnative/variant.h:170
func VariantAsVector2(pSelf []Variant) Vector2 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_vector2(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsRect2 function as declared in gdnative/variant.h:171
func VariantAsRect2(pSelf []Variant) Rect2 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_rect2(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsVector3 function as declared in gdnative/variant.h:172
func VariantAsVector3(pSelf []Variant) Vector3 {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_vector3(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsTransform2d function as declared in gdnative/variant.h:173
func VariantAsTransform2d(pSelf []Variant) Transform2d {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_transform2d(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPlane function as declared in gdnative/variant.h:174
func VariantAsPlane(pSelf []Variant) Plane {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_plane(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsQuat function as declared in gdnative/variant.h:175
func VariantAsQuat(pSelf []Variant) Quat {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_quat(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewQuatRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsAabb function as declared in gdnative/variant.h:176
func VariantAsAabb(pSelf []Variant) Aabb {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_aabb(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsBasis function as declared in gdnative/variant.h:177
func VariantAsBasis(pSelf []Variant) Basis {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_basis(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsTransform function as declared in gdnative/variant.h:178
func VariantAsTransform(pSelf []Variant) Transform {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_transform(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsColor function as declared in gdnative/variant.h:179
func VariantAsColor(pSelf []Variant) Color {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_color(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewColorRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsNodePath function as declared in gdnative/variant.h:180
func VariantAsNodePath(pSelf []Variant) NodePath {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_node_path(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewNodePathRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsRid function as declared in gdnative/variant.h:181
func VariantAsRid(pSelf []Variant) Rid {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_rid(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewRidRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsObject function as declared in gdnative/variant.h:182
func VariantAsObject(pSelf []Variant) *Object {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_object(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *(**Object)(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsDictionary function as declared in gdnative/variant.h:183
func VariantAsDictionary(pSelf []Variant) Dictionary {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_dictionary(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewDictionaryRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsArray function as declared in gdnative/variant.h:184
func VariantAsArray(pSelf []Variant) Array {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolByteArray function as declared in gdnative/variant.h:185
func VariantAsPoolByteArray(pSelf []Variant) PoolByteArray {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_byte_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolByteArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolIntArray function as declared in gdnative/variant.h:186
func VariantAsPoolIntArray(pSelf []Variant) PoolIntArray {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_int_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolIntArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolRealArray function as declared in gdnative/variant.h:187
func VariantAsPoolRealArray(pSelf []Variant) PoolRealArray {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_real_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolRealArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolStringArray function as declared in gdnative/variant.h:188
func VariantAsPoolStringArray(pSelf []Variant) PoolStringArray {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_string_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolStringArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolVector2Array function as declared in gdnative/variant.h:189
func VariantAsPoolVector2Array(pSelf []Variant) PoolVector2Array {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_vector2_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolVector2ArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolVector3Array function as declared in gdnative/variant.h:190
func VariantAsPoolVector3Array(pSelf []Variant) PoolVector3Array {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_vector3_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolVector3ArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantAsPoolColorArray function as declared in gdnative/variant.h:191
func VariantAsPoolColorArray(pSelf []Variant) PoolColorArray {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_as_pool_color_array(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := *NewPoolColorArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantCall function as declared in gdnative/variant.h:193
func VariantCall(pSelf []Variant, pMethod []String, pArgs [][]Variant, pArgcount Int, rError []VariantCallError) Variant {
	cpSelf, _ := unpackArgSVariant(pSelf)
	cpMethod, _ := unpackArgSString(pMethod)
	cpArgs, _ := unpackArgSSVariant(pArgs)
	cpArgcount, _ := (C.godot_int)(pArgcount), cgoAllocsUnknown
	crError, _ := unpackArgSVariantCallError(rError)
	__ret := C.godot_variant_call(cpSelf, cpMethod, cpArgs, cpArgcount, crError)
	packSVariantCallError(rError, crError)
	packSSVariant(pArgs, cpArgs)
	packSString(pMethod, cpMethod)
	packSVariant(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// VariantHasMethod function as declared in gdnative/variant.h:195
func VariantHasMethod(pSelf []Variant, pMethod []String) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	cpMethod, _ := unpackArgSString(pMethod)
	__ret := C.godot_variant_has_method(cpSelf, cpMethod)
	packSString(pMethod, cpMethod)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantOperatorEqual function as declared in gdnative/variant.h:197
func VariantOperatorEqual(pSelf []Variant, pOther []Variant) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	cpOther, _ := unpackArgSVariant(pOther)
	__ret := C.godot_variant_operator_equal(cpSelf, cpOther)
	packSVariant(pOther, cpOther)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantOperatorLess function as declared in gdnative/variant.h:198
func VariantOperatorLess(pSelf []Variant, pOther []Variant) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	cpOther, _ := unpackArgSVariant(pOther)
	__ret := C.godot_variant_operator_less(cpSelf, cpOther)
	packSVariant(pOther, cpOther)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantHashCompare function as declared in gdnative/variant.h:200
func VariantHashCompare(pSelf []Variant, pOther []Variant) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	cpOther, _ := unpackArgSVariant(pOther)
	__ret := C.godot_variant_hash_compare(cpSelf, cpOther)
	packSVariant(pOther, cpOther)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantBooleanize function as declared in gdnative/variant.h:202
func VariantBooleanize(pSelf []Variant) Bool {
	cpSelf, _ := unpackArgSVariant(pSelf)
	__ret := C.godot_variant_booleanize(cpSelf)
	packSVariant(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// VariantDestroy function as declared in gdnative/variant.h:204
func VariantDestroy(pSelf []Variant) {
	cpSelf, _ := unpackArgSVariant(pSelf)
	C.godot_variant_destroy(cpSelf)
	packSVariant(pSelf, cpSelf)
}

// AabbNew function as declared in gdnative/aabb.h:61
func AabbNew(rDest []Aabb, pPos []Vector3, pSize []Vector3) {
	crDest, _ := unpackArgSAabb(rDest)
	cpPos, _ := unpackArgSVector3(pPos)
	cpSize, _ := unpackArgSVector3(pSize)
	C.godot_aabb_new(crDest, cpPos, cpSize)
	packSVector3(pSize, cpSize)
	packSVector3(pPos, cpPos)
	packSAabb(rDest, crDest)
}

// AabbGetPosition function as declared in gdnative/aabb.h:63
func AabbGetPosition(pSelf []Aabb) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_position(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbSetPosition function as declared in gdnative/aabb.h:64
func AabbSetPosition(pSelf []Aabb, pV []Vector3) {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	C.godot_aabb_set_position(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSAabb(pSelf, cpSelf)
}

// AabbGetSize function as declared in gdnative/aabb.h:66
func AabbGetSize(pSelf []Aabb) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_size(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbSetSize function as declared in gdnative/aabb.h:67
func AabbSetSize(pSelf []Aabb, pV []Vector3) {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	C.godot_aabb_set_size(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSAabb(pSelf, cpSelf)
}

// AabbAsString function as declared in gdnative/aabb.h:69
func AabbAsString(pSelf []Aabb) String {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_as_string(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// AabbGetArea function as declared in gdnative/aabb.h:71
func AabbGetArea(pSelf []Aabb) Real {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_area(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// AabbHasNoArea function as declared in gdnative/aabb.h:73
func AabbHasNoArea(pSelf []Aabb) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_has_no_area(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbHasNoSurface function as declared in gdnative/aabb.h:75
func AabbHasNoSurface(pSelf []Aabb) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_has_no_surface(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbIntersects function as declared in gdnative/aabb.h:77
func AabbIntersects(pSelf []Aabb, pWith []Aabb) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpWith, _ := unpackArgSAabb(pWith)
	__ret := C.godot_aabb_intersects(cpSelf, cpWith)
	packSAabb(pWith, cpWith)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbEncloses function as declared in gdnative/aabb.h:79
func AabbEncloses(pSelf []Aabb, pWith []Aabb) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpWith, _ := unpackArgSAabb(pWith)
	__ret := C.godot_aabb_encloses(cpSelf, cpWith)
	packSAabb(pWith, cpWith)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbMerge function as declared in gdnative/aabb.h:81
func AabbMerge(pSelf []Aabb, pWith []Aabb) Aabb {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpWith, _ := unpackArgSAabb(pWith)
	__ret := C.godot_aabb_merge(cpSelf, cpWith)
	packSAabb(pWith, cpWith)
	packSAabb(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// AabbIntersection function as declared in gdnative/aabb.h:83
func AabbIntersection(pSelf []Aabb, pWith []Aabb) Aabb {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpWith, _ := unpackArgSAabb(pWith)
	__ret := C.godot_aabb_intersection(cpSelf, cpWith)
	packSAabb(pWith, cpWith)
	packSAabb(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// AabbIntersectsPlane function as declared in gdnative/aabb.h:85
func AabbIntersectsPlane(pSelf []Aabb, pPlane []Plane) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpPlane, _ := unpackArgSPlane(pPlane)
	__ret := C.godot_aabb_intersects_plane(cpSelf, cpPlane)
	packSPlane(pPlane, cpPlane)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbIntersectsSegment function as declared in gdnative/aabb.h:87
func AabbIntersectsSegment(pSelf []Aabb, pFrom []Vector3, pTo []Vector3) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpFrom, _ := unpackArgSVector3(pFrom)
	cpTo, _ := unpackArgSVector3(pTo)
	__ret := C.godot_aabb_intersects_segment(cpSelf, cpFrom, cpTo)
	packSVector3(pTo, cpTo)
	packSVector3(pFrom, cpFrom)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbHasPoint function as declared in gdnative/aabb.h:89
func AabbHasPoint(pSelf []Aabb, pPoint []Vector3) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpPoint, _ := unpackArgSVector3(pPoint)
	__ret := C.godot_aabb_has_point(cpSelf, cpPoint)
	packSVector3(pPoint, cpPoint)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// AabbGetSupport function as declared in gdnative/aabb.h:91
func AabbGetSupport(pSelf []Aabb, pDir []Vector3) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpDir, _ := unpackArgSVector3(pDir)
	__ret := C.godot_aabb_get_support(cpSelf, cpDir)
	packSVector3(pDir, cpDir)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbGetLongestAxis function as declared in gdnative/aabb.h:93
func AabbGetLongestAxis(pSelf []Aabb) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbGetLongestAxisIndex function as declared in gdnative/aabb.h:95
func AabbGetLongestAxisIndex(pSelf []Aabb) Int {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis_index(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// AabbGetLongestAxisSize function as declared in gdnative/aabb.h:97
func AabbGetLongestAxisSize(pSelf []Aabb) Real {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_longest_axis_size(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// AabbGetShortestAxis function as declared in gdnative/aabb.h:99
func AabbGetShortestAxis(pSelf []Aabb) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbGetShortestAxisIndex function as declared in gdnative/aabb.h:101
func AabbGetShortestAxisIndex(pSelf []Aabb) Int {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis_index(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// AabbGetShortestAxisSize function as declared in gdnative/aabb.h:103
func AabbGetShortestAxisSize(pSelf []Aabb) Real {
	cpSelf, _ := unpackArgSAabb(pSelf)
	__ret := C.godot_aabb_get_shortest_axis_size(cpSelf)
	packSAabb(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// AabbExpand function as declared in gdnative/aabb.h:105
func AabbExpand(pSelf []Aabb, pToPoint []Vector3) Aabb {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpToPoint, _ := unpackArgSVector3(pToPoint)
	__ret := C.godot_aabb_expand(cpSelf, cpToPoint)
	packSVector3(pToPoint, cpToPoint)
	packSAabb(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// AabbGrow function as declared in gdnative/aabb.h:107
func AabbGrow(pSelf []Aabb, pBy Real) Aabb {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpBy, _ := (C.godot_real)(pBy), cgoAllocsUnknown
	__ret := C.godot_aabb_grow(cpSelf, cpBy)
	packSAabb(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// AabbGetEndpoint function as declared in gdnative/aabb.h:109
func AabbGetEndpoint(pSelf []Aabb, pIdx Int) Vector3 {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_aabb_get_endpoint(cpSelf, cpIdx)
	packSAabb(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// AabbOperatorEqual function as declared in gdnative/aabb.h:111
func AabbOperatorEqual(pSelf []Aabb, pB []Aabb) Bool {
	cpSelf, _ := unpackArgSAabb(pSelf)
	cpB, _ := unpackArgSAabb(pB)
	__ret := C.godot_aabb_operator_equal(cpSelf, cpB)
	packSAabb(pB, cpB)
	packSAabb(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneNewWithReals function as declared in gdnative/plane.h:60
func PlaneNewWithReals(rDest []Plane, pA Real, pB Real, pC Real, pD Real) {
	crDest, _ := unpackArgSPlane(rDest)
	cpA, _ := (C.godot_real)(pA), cgoAllocsUnknown
	cpB, _ := (C.godot_real)(pB), cgoAllocsUnknown
	cpC, _ := (C.godot_real)(pC), cgoAllocsUnknown
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_new_with_reals(crDest, cpA, cpB, cpC, cpD)
	packSPlane(rDest, crDest)
}

// PlaneNewWithVectors function as declared in gdnative/plane.h:61
func PlaneNewWithVectors(rDest []Plane, pV1 []Vector3, pV2 []Vector3, pV3 []Vector3) {
	crDest, _ := unpackArgSPlane(rDest)
	cpV1, _ := unpackArgSVector3(pV1)
	cpV2, _ := unpackArgSVector3(pV2)
	cpV3, _ := unpackArgSVector3(pV3)
	C.godot_plane_new_with_vectors(crDest, cpV1, cpV2, cpV3)
	packSVector3(pV3, cpV3)
	packSVector3(pV2, cpV2)
	packSVector3(pV1, cpV1)
	packSPlane(rDest, crDest)
}

// PlaneNewWithNormal function as declared in gdnative/plane.h:62
func PlaneNewWithNormal(rDest []Plane, pNormal []Vector3, pD Real) {
	crDest, _ := unpackArgSPlane(rDest)
	cpNormal, _ := unpackArgSVector3(pNormal)
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_new_with_normal(crDest, cpNormal, cpD)
	packSVector3(pNormal, cpNormal)
	packSPlane(rDest, crDest)
}

// PlaneAsString function as declared in gdnative/plane.h:64
func PlaneAsString(pSelf []Plane) String {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_as_string(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// PlaneNormalized function as declared in gdnative/plane.h:66
func PlaneNormalized(pSelf []Plane) Plane {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_normalized(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// PlaneCenter function as declared in gdnative/plane.h:68
func PlaneCenter(pSelf []Plane) Vector3 {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_center(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// PlaneGetAnyPoint function as declared in gdnative/plane.h:70
func PlaneGetAnyPoint(pSelf []Plane) Vector3 {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_get_any_point(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// PlaneIsPointOver function as declared in gdnative/plane.h:72
func PlaneIsPointOver(pSelf []Plane, pPoint []Vector3) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpPoint, _ := unpackArgSVector3(pPoint)
	__ret := C.godot_plane_is_point_over(cpSelf, cpPoint)
	packSVector3(pPoint, cpPoint)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneDistanceTo function as declared in gdnative/plane.h:74
func PlaneDistanceTo(pSelf []Plane, pPoint []Vector3) Real {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpPoint, _ := unpackArgSVector3(pPoint)
	__ret := C.godot_plane_distance_to(cpSelf, cpPoint)
	packSVector3(pPoint, cpPoint)
	packSPlane(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// PlaneHasPoint function as declared in gdnative/plane.h:76
func PlaneHasPoint(pSelf []Plane, pPoint []Vector3, pEpsilon Real) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpPoint, _ := unpackArgSVector3(pPoint)
	cpEpsilon, _ := (C.godot_real)(pEpsilon), cgoAllocsUnknown
	__ret := C.godot_plane_has_point(cpSelf, cpPoint, cpEpsilon)
	packSVector3(pPoint, cpPoint)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneProject function as declared in gdnative/plane.h:78
func PlaneProject(pSelf []Plane, pPoint []Vector3) Vector3 {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpPoint, _ := unpackArgSVector3(pPoint)
	__ret := C.godot_plane_project(cpSelf, cpPoint)
	packSVector3(pPoint, cpPoint)
	packSPlane(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// PlaneIntersect3 function as declared in gdnative/plane.h:80
func PlaneIntersect3(pSelf []Plane, rDest []Vector3, pB []Plane, pC []Plane) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	crDest, _ := unpackArgSVector3(rDest)
	cpB, _ := unpackArgSPlane(pB)
	cpC, _ := unpackArgSPlane(pC)
	__ret := C.godot_plane_intersect_3(cpSelf, crDest, cpB, cpC)
	packSPlane(pC, cpC)
	packSPlane(pB, cpB)
	packSVector3(rDest, crDest)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneIntersectsRay function as declared in gdnative/plane.h:82
func PlaneIntersectsRay(pSelf []Plane, rDest []Vector3, pFrom []Vector3, pDir []Vector3) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	crDest, _ := unpackArgSVector3(rDest)
	cpFrom, _ := unpackArgSVector3(pFrom)
	cpDir, _ := unpackArgSVector3(pDir)
	__ret := C.godot_plane_intersects_ray(cpSelf, crDest, cpFrom, cpDir)
	packSVector3(pDir, cpDir)
	packSVector3(pFrom, cpFrom)
	packSVector3(rDest, crDest)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneIntersectsSegment function as declared in gdnative/plane.h:84
func PlaneIntersectsSegment(pSelf []Plane, rDest []Vector3, pBegin []Vector3, pEnd []Vector3) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	crDest, _ := unpackArgSVector3(rDest)
	cpBegin, _ := unpackArgSVector3(pBegin)
	cpEnd, _ := unpackArgSVector3(pEnd)
	__ret := C.godot_plane_intersects_segment(cpSelf, crDest, cpBegin, cpEnd)
	packSVector3(pEnd, cpEnd)
	packSVector3(pBegin, cpBegin)
	packSVector3(rDest, crDest)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneOperatorNeg function as declared in gdnative/plane.h:86
func PlaneOperatorNeg(pSelf []Plane) Plane {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_operator_neg(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// PlaneOperatorEqual function as declared in gdnative/plane.h:88
func PlaneOperatorEqual(pSelf []Plane, pB []Plane) Bool {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpB, _ := unpackArgSPlane(pB)
	__ret := C.godot_plane_operator_equal(cpSelf, cpB)
	packSPlane(pB, cpB)
	packSPlane(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// PlaneSetNormal function as declared in gdnative/plane.h:90
func PlaneSetNormal(pSelf []Plane, pNormal []Vector3) {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpNormal, _ := unpackArgSVector3(pNormal)
	C.godot_plane_set_normal(cpSelf, cpNormal)
	packSVector3(pNormal, cpNormal)
	packSPlane(pSelf, cpSelf)
}

// PlaneGetNormal function as declared in gdnative/plane.h:92
func PlaneGetNormal(pSelf []Plane) Vector3 {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_get_normal(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// PlaneGetD function as declared in gdnative/plane.h:94
func PlaneGetD(pSelf []Plane) Real {
	cpSelf, _ := unpackArgSPlane(pSelf)
	__ret := C.godot_plane_get_d(cpSelf)
	packSPlane(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// PlaneSetD function as declared in gdnative/plane.h:96
func PlaneSetD(pSelf []Plane, pD Real) {
	cpSelf, _ := unpackArgSPlane(pSelf)
	cpD, _ := (C.godot_real)(pD), cgoAllocsUnknown
	C.godot_plane_set_d(cpSelf, cpD)
	packSPlane(pSelf, cpSelf)
}

// DictionaryNew function as declared in gdnative/dictionary.h:61
func DictionaryNew(rDest []Dictionary) {
	crDest, _ := unpackArgSDictionary(rDest)
	C.godot_dictionary_new(crDest)
	packSDictionary(rDest, crDest)
}

// DictionaryNewCopy function as declared in gdnative/dictionary.h:62
func DictionaryNewCopy(rDest []Dictionary, pSrc []Dictionary) {
	crDest, _ := unpackArgSDictionary(rDest)
	cpSrc, _ := unpackArgSDictionary(pSrc)
	C.godot_dictionary_new_copy(crDest, cpSrc)
	packSDictionary(pSrc, cpSrc)
	packSDictionary(rDest, crDest)
}

// DictionaryDestroy function as declared in gdnative/dictionary.h:63
func DictionaryDestroy(pSelf []Dictionary) {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	C.godot_dictionary_destroy(cpSelf)
	packSDictionary(pSelf, cpSelf)
}

// DictionarySize function as declared in gdnative/dictionary.h:65
func DictionarySize(pSelf []Dictionary) Int {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_size(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// DictionaryEmpty function as declared in gdnative/dictionary.h:67
func DictionaryEmpty(pSelf []Dictionary) Bool {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_empty(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// DictionaryClear function as declared in gdnative/dictionary.h:69
func DictionaryClear(pSelf []Dictionary) {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	C.godot_dictionary_clear(cpSelf)
	packSDictionary(pSelf, cpSelf)
}

// DictionaryHas function as declared in gdnative/dictionary.h:71
func DictionaryHas(pSelf []Dictionary, pKey []Variant) Bool {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	__ret := C.godot_dictionary_has(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// DictionaryHasAll function as declared in gdnative/dictionary.h:73
func DictionaryHasAll(pSelf []Dictionary, pKeys []Array) Bool {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKeys, _ := unpackArgSArray(pKeys)
	__ret := C.godot_dictionary_has_all(cpSelf, cpKeys)
	packSArray(pKeys, cpKeys)
	packSDictionary(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// DictionaryErase function as declared in gdnative/dictionary.h:75
func DictionaryErase(pSelf []Dictionary, pKey []Variant) {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	C.godot_dictionary_erase(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
}

// DictionaryHash function as declared in gdnative/dictionary.h:77
func DictionaryHash(pSelf []Dictionary) Int {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_hash(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// DictionaryKeys function as declared in gdnative/dictionary.h:79
func DictionaryKeys(pSelf []Dictionary) Array {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_keys(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// DictionaryValues function as declared in gdnative/dictionary.h:81
func DictionaryValues(pSelf []Dictionary) Array {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_values(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := *NewArrayRef(unsafe.Pointer(&__ret))
	return __v
}

// DictionaryGet function as declared in gdnative/dictionary.h:83
func DictionaryGet(pSelf []Dictionary, pKey []Variant) Variant {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	__ret := C.godot_dictionary_get(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
	__v := *NewVariantRef(unsafe.Pointer(&__ret))
	return __v
}

// DictionarySet function as declared in gdnative/dictionary.h:84
func DictionarySet(pSelf []Dictionary, pKey []Variant, pValue []Variant) {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	cpValue, _ := unpackArgSVariant(pValue)
	C.godot_dictionary_set(cpSelf, cpKey, cpValue)
	packSVariant(pValue, cpValue)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
}

// DictionaryOperatorIndex function as declared in gdnative/dictionary.h:86
func DictionaryOperatorIndex(pSelf []Dictionary, pKey []Variant) *Variant {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	__ret := C.godot_dictionary_operator_index(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
	__v := NewVariantRef(unsafe.Pointer(__ret))
	return __v
}

// DictionaryOperatorIndexConst function as declared in gdnative/dictionary.h:88
func DictionaryOperatorIndexConst(pSelf []Dictionary, pKey []Variant) *Variant {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	__ret := C.godot_dictionary_operator_index_const(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
	__v := NewVariantRef(unsafe.Pointer(__ret))
	return __v
}

// DictionaryNext function as declared in gdnative/dictionary.h:90
func DictionaryNext(pSelf []Dictionary, pKey []Variant) *Variant {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpKey, _ := unpackArgSVariant(pKey)
	__ret := C.godot_dictionary_next(cpSelf, cpKey)
	packSVariant(pKey, cpKey)
	packSDictionary(pSelf, cpSelf)
	__v := NewVariantRef(unsafe.Pointer(__ret))
	return __v
}

// DictionaryOperatorEqual function as declared in gdnative/dictionary.h:92
func DictionaryOperatorEqual(pSelf []Dictionary, pB []Dictionary) Bool {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	cpB, _ := unpackArgSDictionary(pB)
	__ret := C.godot_dictionary_operator_equal(cpSelf, cpB)
	packSDictionary(pB, cpB)
	packSDictionary(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// DictionaryToJson function as declared in gdnative/dictionary.h:94
func DictionaryToJson(pSelf []Dictionary) String {
	cpSelf, _ := unpackArgSDictionary(pSelf)
	__ret := C.godot_dictionary_to_json(cpSelf)
	packSDictionary(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// NodePathNew function as declared in gdnative/node_path.h:60
func NodePathNew(rDest []NodePath, pFrom []String) {
	crDest, _ := unpackArgSNodePath(rDest)
	cpFrom, _ := unpackArgSString(pFrom)
	C.godot_node_path_new(crDest, cpFrom)
	packSString(pFrom, cpFrom)
	packSNodePath(rDest, crDest)
}

// NodePathNewCopy function as declared in gdnative/node_path.h:61
func NodePathNewCopy(rDest []NodePath, pSrc []NodePath) {
	crDest, _ := unpackArgSNodePath(rDest)
	cpSrc, _ := unpackArgSNodePath(pSrc)
	C.godot_node_path_new_copy(crDest, cpSrc)
	packSNodePath(pSrc, cpSrc)
	packSNodePath(rDest, crDest)
}

// NodePathDestroy function as declared in gdnative/node_path.h:62
func NodePathDestroy(pSelf []NodePath) {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	C.godot_node_path_destroy(cpSelf)
	packSNodePath(pSelf, cpSelf)
}

// NodePathAsString function as declared in gdnative/node_path.h:64
func NodePathAsString(pSelf []NodePath) String {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_as_string(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// NodePathIsAbsolute function as declared in gdnative/node_path.h:66
func NodePathIsAbsolute(pSelf []NodePath) Bool {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_is_absolute(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// NodePathGetNameCount function as declared in gdnative/node_path.h:68
func NodePathGetNameCount(pSelf []NodePath) Int {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_get_name_count(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// NodePathGetName function as declared in gdnative/node_path.h:70
func NodePathGetName(pSelf []NodePath, pIdx Int) String {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_node_path_get_name(cpSelf, cpIdx)
	packSNodePath(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// NodePathGetSubnameCount function as declared in gdnative/node_path.h:72
func NodePathGetSubnameCount(pSelf []NodePath) Int {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_get_subname_count(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// NodePathGetSubname function as declared in gdnative/node_path.h:74
func NodePathGetSubname(pSelf []NodePath, pIdx Int) String {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	cpIdx, _ := (C.godot_int)(pIdx), cgoAllocsUnknown
	__ret := C.godot_node_path_get_subname(cpSelf, cpIdx)
	packSNodePath(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// NodePathGetConcatenatedSubnames function as declared in gdnative/node_path.h:76
func NodePathGetConcatenatedSubnames(pSelf []NodePath) String {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_get_concatenated_subnames(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// NodePathIsEmpty function as declared in gdnative/node_path.h:78
func NodePathIsEmpty(pSelf []NodePath) Bool {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	__ret := C.godot_node_path_is_empty(cpSelf)
	packSNodePath(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// NodePathOperatorEqual function as declared in gdnative/node_path.h:80
func NodePathOperatorEqual(pSelf []NodePath, pB []NodePath) Bool {
	cpSelf, _ := unpackArgSNodePath(pSelf)
	cpB, _ := unpackArgSNodePath(pB)
	__ret := C.godot_node_path_operator_equal(cpSelf, cpB)
	packSNodePath(pB, cpB)
	packSNodePath(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2NewWithPositionAndSize function as declared in gdnative/rect2.h:58
func Rect2NewWithPositionAndSize(rDest []Rect2, pPos []Vector2, pSize []Vector2) {
	crDest, _ := unpackArgSRect2(rDest)
	cpPos, _ := unpackArgSVector2(pPos)
	cpSize, _ := unpackArgSVector2(pSize)
	C.godot_rect2_new_with_position_and_size(crDest, cpPos, cpSize)
	packSVector2(pSize, cpSize)
	packSVector2(pPos, cpPos)
	packSRect2(rDest, crDest)
}

// Rect2New function as declared in gdnative/rect2.h:59
func Rect2New(rDest []Rect2, pX Real, pY Real, pWidth Real, pHeight Real) {
	crDest, _ := unpackArgSRect2(rDest)
	cpX, _ := (C.godot_real)(pX), cgoAllocsUnknown
	cpY, _ := (C.godot_real)(pY), cgoAllocsUnknown
	cpWidth, _ := (C.godot_real)(pWidth), cgoAllocsUnknown
	cpHeight, _ := (C.godot_real)(pHeight), cgoAllocsUnknown
	C.godot_rect2_new(crDest, cpX, cpY, cpWidth, cpHeight)
	packSRect2(rDest, crDest)
}

// Rect2AsString function as declared in gdnative/rect2.h:61
func Rect2AsString(pSelf []Rect2) String {
	cpSelf, _ := unpackArgSRect2(pSelf)
	__ret := C.godot_rect2_as_string(cpSelf)
	packSRect2(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// Rect2GetArea function as declared in gdnative/rect2.h:63
func Rect2GetArea(pSelf []Rect2) Real {
	cpSelf, _ := unpackArgSRect2(pSelf)
	__ret := C.godot_rect2_get_area(cpSelf)
	packSRect2(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Rect2Intersects function as declared in gdnative/rect2.h:65
func Rect2Intersects(pSelf []Rect2, pB []Rect2) Bool {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpB, _ := unpackArgSRect2(pB)
	__ret := C.godot_rect2_intersects(cpSelf, cpB)
	packSRect2(pB, cpB)
	packSRect2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2Encloses function as declared in gdnative/rect2.h:67
func Rect2Encloses(pSelf []Rect2, pB []Rect2) Bool {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpB, _ := unpackArgSRect2(pB)
	__ret := C.godot_rect2_encloses(cpSelf, cpB)
	packSRect2(pB, cpB)
	packSRect2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2HasNoArea function as declared in gdnative/rect2.h:69
func Rect2HasNoArea(pSelf []Rect2) Bool {
	cpSelf, _ := unpackArgSRect2(pSelf)
	__ret := C.godot_rect2_has_no_area(cpSelf)
	packSRect2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2Clip function as declared in gdnative/rect2.h:71
func Rect2Clip(pSelf []Rect2, pB []Rect2) Rect2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpB, _ := unpackArgSRect2(pB)
	__ret := C.godot_rect2_clip(cpSelf, cpB)
	packSRect2(pB, cpB)
	packSRect2(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2Merge function as declared in gdnative/rect2.h:73
func Rect2Merge(pSelf []Rect2, pB []Rect2) Rect2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpB, _ := unpackArgSRect2(pB)
	__ret := C.godot_rect2_merge(cpSelf, cpB)
	packSRect2(pB, cpB)
	packSRect2(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2HasPoint function as declared in gdnative/rect2.h:75
func Rect2HasPoint(pSelf []Rect2, pPoint []Vector2) Bool {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpPoint, _ := unpackArgSVector2(pPoint)
	__ret := C.godot_rect2_has_point(cpSelf, cpPoint)
	packSVector2(pPoint, cpPoint)
	packSRect2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2Grow function as declared in gdnative/rect2.h:77
func Rect2Grow(pSelf []Rect2, pBy Real) Rect2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpBy, _ := (C.godot_real)(pBy), cgoAllocsUnknown
	__ret := C.godot_rect2_grow(cpSelf, cpBy)
	packSRect2(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2Expand function as declared in gdnative/rect2.h:79
func Rect2Expand(pSelf []Rect2, pTo []Vector2) Rect2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpTo, _ := unpackArgSVector2(pTo)
	__ret := C.godot_rect2_expand(cpSelf, cpTo)
	packSVector2(pTo, cpTo)
	packSRect2(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2OperatorEqual function as declared in gdnative/rect2.h:81
func Rect2OperatorEqual(pSelf []Rect2, pB []Rect2) Bool {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpB, _ := unpackArgSRect2(pB)
	__ret := C.godot_rect2_operator_equal(cpSelf, cpB)
	packSRect2(pB, cpB)
	packSRect2(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Rect2GetPosition function as declared in gdnative/rect2.h:83
func Rect2GetPosition(pSelf []Rect2) Vector2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	__ret := C.godot_rect2_get_position(cpSelf)
	packSRect2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2GetSize function as declared in gdnative/rect2.h:85
func Rect2GetSize(pSelf []Rect2) Vector2 {
	cpSelf, _ := unpackArgSRect2(pSelf)
	__ret := C.godot_rect2_get_size(cpSelf)
	packSRect2(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Rect2SetPosition function as declared in gdnative/rect2.h:87
func Rect2SetPosition(pSelf []Rect2, pPos []Vector2) {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpPos, _ := unpackArgSVector2(pPos)
	C.godot_rect2_set_position(cpSelf, cpPos)
	packSVector2(pPos, cpPos)
	packSRect2(pSelf, cpSelf)
}

// Rect2SetSize function as declared in gdnative/rect2.h:89
func Rect2SetSize(pSelf []Rect2, pSize []Vector2) {
	cpSelf, _ := unpackArgSRect2(pSelf)
	cpSize, _ := unpackArgSVector2(pSize)
	C.godot_rect2_set_size(cpSelf, cpSize)
	packSVector2(pSize, cpSize)
	packSRect2(pSelf, cpSelf)
}

// RidNew function as declared in gdnative/rid.h:59
func RidNew(rDest []Rid) {
	crDest, _ := unpackArgSRid(rDest)
	C.godot_rid_new(crDest)
	packSRid(rDest, crDest)
}

// RidGetId function as declared in gdnative/rid.h:61
func RidGetId(pSelf []Rid) Int {
	cpSelf, _ := unpackArgSRid(pSelf)
	__ret := C.godot_rid_get_id(cpSelf)
	packSRid(pSelf, cpSelf)
	__v := (Int)(__ret)
	return __v
}

// RidNewWithResource function as declared in gdnative/rid.h:63
func RidNewWithResource(rDest []Rid, pFrom *Object) {
	crDest, _ := unpackArgSRid(rDest)
	cpFrom, _ := (C.godot_object)(unsafe.Pointer(pFrom)), cgoAllocsUnknown
	C.godot_rid_new_with_resource(crDest, cpFrom)
	packSRid(rDest, crDest)
}

// RidOperatorEqual function as declared in gdnative/rid.h:65
func RidOperatorEqual(pSelf []Rid, pB []Rid) Bool {
	cpSelf, _ := unpackArgSRid(pSelf)
	cpB, _ := unpackArgSRid(pB)
	__ret := C.godot_rid_operator_equal(cpSelf, cpB)
	packSRid(pB, cpB)
	packSRid(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// RidOperatorLess function as declared in gdnative/rid.h:67
func RidOperatorLess(pSelf []Rid, pB []Rid) Bool {
	cpSelf, _ := unpackArgSRid(pSelf)
	cpB, _ := unpackArgSRid(pB)
	__ret := C.godot_rid_operator_less(cpSelf, cpB)
	packSRid(pB, cpB)
	packSRid(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// TransformNewWithAxisOrigin function as declared in gdnative/transform.h:62
func TransformNewWithAxisOrigin(rDest []Transform, pXAxis []Vector3, pYAxis []Vector3, pZAxis []Vector3, pOrigin []Vector3) {
	crDest, _ := unpackArgSTransform(rDest)
	cpXAxis, _ := unpackArgSVector3(pXAxis)
	cpYAxis, _ := unpackArgSVector3(pYAxis)
	cpZAxis, _ := unpackArgSVector3(pZAxis)
	cpOrigin, _ := unpackArgSVector3(pOrigin)
	C.godot_transform_new_with_axis_origin(crDest, cpXAxis, cpYAxis, cpZAxis, cpOrigin)
	packSVector3(pOrigin, cpOrigin)
	packSVector3(pZAxis, cpZAxis)
	packSVector3(pYAxis, cpYAxis)
	packSVector3(pXAxis, cpXAxis)
	packSTransform(rDest, crDest)
}

// TransformNew function as declared in gdnative/transform.h:63
func TransformNew(rDest []Transform, pBasis []Basis, pOrigin []Vector3) {
	crDest, _ := unpackArgSTransform(rDest)
	cpBasis, _ := unpackArgSBasis(pBasis)
	cpOrigin, _ := unpackArgSVector3(pOrigin)
	C.godot_transform_new(crDest, cpBasis, cpOrigin)
	packSVector3(pOrigin, cpOrigin)
	packSBasis(pBasis, cpBasis)
	packSTransform(rDest, crDest)
}

// TransformGetBasis function as declared in gdnative/transform.h:65
func TransformGetBasis(pSelf []Transform) Basis {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_get_basis(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewBasisRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformSetBasis function as declared in gdnative/transform.h:66
func TransformSetBasis(pSelf []Transform, pV []Basis) {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSBasis(pV)
	C.godot_transform_set_basis(cpSelf, cpV)
	packSBasis(pV, cpV)
	packSTransform(pSelf, cpSelf)
}

// TransformGetOrigin function as declared in gdnative/transform.h:68
func TransformGetOrigin(pSelf []Transform) Vector3 {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_get_origin(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// TransformSetOrigin function as declared in gdnative/transform.h:69
func TransformSetOrigin(pSelf []Transform, pV []Vector3) {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	C.godot_transform_set_origin(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSTransform(pSelf, cpSelf)
}

// TransformAsString function as declared in gdnative/transform.h:71
func TransformAsString(pSelf []Transform) String {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_as_string(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformInverse function as declared in gdnative/transform.h:73
func TransformInverse(pSelf []Transform) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_inverse(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformAffineInverse function as declared in gdnative/transform.h:75
func TransformAffineInverse(pSelf []Transform) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_affine_inverse(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformOrthonormalized function as declared in gdnative/transform.h:77
func TransformOrthonormalized(pSelf []Transform) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	__ret := C.godot_transform_orthonormalized(cpSelf)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformRotated function as declared in gdnative/transform.h:79
func TransformRotated(pSelf []Transform, pAxis []Vector3, pPhi Real) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpAxis, _ := unpackArgSVector3(pAxis)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_transform_rotated(cpSelf, cpAxis, cpPhi)
	packSVector3(pAxis, cpAxis)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformScaled function as declared in gdnative/transform.h:81
func TransformScaled(pSelf []Transform, pScale []Vector3) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpScale, _ := unpackArgSVector3(pScale)
	__ret := C.godot_transform_scaled(cpSelf, cpScale)
	packSVector3(pScale, cpScale)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformTranslated function as declared in gdnative/transform.h:83
func TransformTranslated(pSelf []Transform, pOfs []Vector3) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpOfs, _ := unpackArgSVector3(pOfs)
	__ret := C.godot_transform_translated(cpSelf, cpOfs)
	packSVector3(pOfs, cpOfs)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformLookingAt function as declared in gdnative/transform.h:85
func TransformLookingAt(pSelf []Transform, pTarget []Vector3, pUp []Vector3) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpTarget, _ := unpackArgSVector3(pTarget)
	cpUp, _ := unpackArgSVector3(pUp)
	__ret := C.godot_transform_looking_at(cpSelf, cpTarget, cpUp)
	packSVector3(pUp, cpUp)
	packSVector3(pTarget, cpTarget)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformPlane function as declared in gdnative/transform.h:87
func TransformXformPlane(pSelf []Transform, pV []Plane) Plane {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSPlane(pV)
	__ret := C.godot_transform_xform_plane(cpSelf, cpV)
	packSPlane(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformInvPlane function as declared in gdnative/transform.h:89
func TransformXformInvPlane(pSelf []Transform, pV []Plane) Plane {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSPlane(pV)
	__ret := C.godot_transform_xform_inv_plane(cpSelf, cpV)
	packSPlane(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewPlaneRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformNewIdentity function as declared in gdnative/transform.h:91
func TransformNewIdentity(rDest []Transform) {
	crDest, _ := unpackArgSTransform(rDest)
	C.godot_transform_new_identity(crDest)
	packSTransform(rDest, crDest)
}

// TransformOperatorEqual function as declared in gdnative/transform.h:93
func TransformOperatorEqual(pSelf []Transform, pB []Transform) Bool {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpB, _ := unpackArgSTransform(pB)
	__ret := C.godot_transform_operator_equal(cpSelf, cpB)
	packSTransform(pB, cpB)
	packSTransform(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// TransformOperatorMultiply function as declared in gdnative/transform.h:95
func TransformOperatorMultiply(pSelf []Transform, pB []Transform) Transform {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpB, _ := unpackArgSTransform(pB)
	__ret := C.godot_transform_operator_multiply(cpSelf, cpB)
	packSTransform(pB, cpB)
	packSTransform(pSelf, cpSelf)
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformVector3 function as declared in gdnative/transform.h:97
func TransformXformVector3(pSelf []Transform, pV []Vector3) Vector3 {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	__ret := C.godot_transform_xform_vector3(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformInvVector3 function as declared in gdnative/transform.h:99
func TransformXformInvVector3(pSelf []Transform, pV []Vector3) Vector3 {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSVector3(pV)
	__ret := C.godot_transform_xform_inv_vector3(cpSelf, cpV)
	packSVector3(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewVector3Ref(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformAabb function as declared in gdnative/transform.h:101
func TransformXformAabb(pSelf []Transform, pV []Aabb) Aabb {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSAabb(pV)
	__ret := C.godot_transform_xform_aabb(cpSelf, cpV)
	packSAabb(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// TransformXformInvAabb function as declared in gdnative/transform.h:103
func TransformXformInvAabb(pSelf []Transform, pV []Aabb) Aabb {
	cpSelf, _ := unpackArgSTransform(pSelf)
	cpV, _ := unpackArgSAabb(pV)
	__ret := C.godot_transform_xform_inv_aabb(cpSelf, cpV)
	packSAabb(pV, cpV)
	packSTransform(pSelf, cpSelf)
	__v := *NewAabbRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dNew function as declared in gdnative/transform2d.h:61
func Transform2dNew(rDest []Transform2d, pRot Real, pPos []Vector2) {
	crDest, _ := unpackArgSTransform2d(rDest)
	cpRot, _ := (C.godot_real)(pRot), cgoAllocsUnknown
	cpPos, _ := unpackArgSVector2(pPos)
	C.godot_transform2d_new(crDest, cpRot, cpPos)
	packSVector2(pPos, cpPos)
	packSTransform2d(rDest, crDest)
}

// Transform2dNewAxisOrigin function as declared in gdnative/transform2d.h:62
func Transform2dNewAxisOrigin(rDest []Transform2d, pXAxis []Vector2, pYAxis []Vector2, pOrigin []Vector2) {
	crDest, _ := unpackArgSTransform2d(rDest)
	cpXAxis, _ := unpackArgSVector2(pXAxis)
	cpYAxis, _ := unpackArgSVector2(pYAxis)
	cpOrigin, _ := unpackArgSVector2(pOrigin)
	C.godot_transform2d_new_axis_origin(crDest, cpXAxis, cpYAxis, cpOrigin)
	packSVector2(pOrigin, cpOrigin)
	packSVector2(pYAxis, cpYAxis)
	packSVector2(pXAxis, cpXAxis)
	packSTransform2d(rDest, crDest)
}

// Transform2dAsString function as declared in gdnative/transform2d.h:64
func Transform2dAsString(pSelf []Transform2d) String {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_as_string(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dInverse function as declared in gdnative/transform2d.h:66
func Transform2dInverse(pSelf []Transform2d) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_inverse(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dAffineInverse function as declared in gdnative/transform2d.h:68
func Transform2dAffineInverse(pSelf []Transform2d) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_affine_inverse(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dGetRotation function as declared in gdnative/transform2d.h:70
func Transform2dGetRotation(pSelf []Transform2d) Real {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_get_rotation(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := (Real)(__ret)
	return __v
}

// Transform2dGetOrigin function as declared in gdnative/transform2d.h:72
func Transform2dGetOrigin(pSelf []Transform2d) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_get_origin(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dGetScale function as declared in gdnative/transform2d.h:74
func Transform2dGetScale(pSelf []Transform2d) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_get_scale(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dOrthonormalized function as declared in gdnative/transform2d.h:76
func Transform2dOrthonormalized(pSelf []Transform2d) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	__ret := C.godot_transform2d_orthonormalized(cpSelf)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dRotated function as declared in gdnative/transform2d.h:78
func Transform2dRotated(pSelf []Transform2d, pPhi Real) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpPhi, _ := (C.godot_real)(pPhi), cgoAllocsUnknown
	__ret := C.godot_transform2d_rotated(cpSelf, cpPhi)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dScaled function as declared in gdnative/transform2d.h:80
func Transform2dScaled(pSelf []Transform2d, pScale []Vector2) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpScale, _ := unpackArgSVector2(pScale)
	__ret := C.godot_transform2d_scaled(cpSelf, cpScale)
	packSVector2(pScale, cpScale)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dTranslated function as declared in gdnative/transform2d.h:82
func Transform2dTranslated(pSelf []Transform2d, pOffset []Vector2) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpOffset, _ := unpackArgSVector2(pOffset)
	__ret := C.godot_transform2d_translated(cpSelf, cpOffset)
	packSVector2(pOffset, cpOffset)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dXformVector2 function as declared in gdnative/transform2d.h:84
func Transform2dXformVector2(pSelf []Transform2d, pV []Vector2) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSVector2(pV)
	__ret := C.godot_transform2d_xform_vector2(cpSelf, cpV)
	packSVector2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dXformInvVector2 function as declared in gdnative/transform2d.h:86
func Transform2dXformInvVector2(pSelf []Transform2d, pV []Vector2) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSVector2(pV)
	__ret := C.godot_transform2d_xform_inv_vector2(cpSelf, cpV)
	packSVector2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dBasisXformVector2 function as declared in gdnative/transform2d.h:88
func Transform2dBasisXformVector2(pSelf []Transform2d, pV []Vector2) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSVector2(pV)
	__ret := C.godot_transform2d_basis_xform_vector2(cpSelf, cpV)
	packSVector2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dBasisXformInvVector2 function as declared in gdnative/transform2d.h:90
func Transform2dBasisXformInvVector2(pSelf []Transform2d, pV []Vector2) Vector2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSVector2(pV)
	__ret := C.godot_transform2d_basis_xform_inv_vector2(cpSelf, cpV)
	packSVector2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewVector2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dInterpolateWith function as declared in gdnative/transform2d.h:92
func Transform2dInterpolateWith(pSelf []Transform2d, pM []Transform2d, pC Real) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpM, _ := unpackArgSTransform2d(pM)
	cpC, _ := (C.godot_real)(pC), cgoAllocsUnknown
	__ret := C.godot_transform2d_interpolate_with(cpSelf, cpM, cpC)
	packSTransform2d(pM, cpM)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dOperatorEqual function as declared in gdnative/transform2d.h:94
func Transform2dOperatorEqual(pSelf []Transform2d, pB []Transform2d) Bool {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpB, _ := unpackArgSTransform2d(pB)
	__ret := C.godot_transform2d_operator_equal(cpSelf, cpB)
	packSTransform2d(pB, cpB)
	packSTransform2d(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// Transform2dOperatorMultiply function as declared in gdnative/transform2d.h:96
func Transform2dOperatorMultiply(pSelf []Transform2d, pB []Transform2d) Transform2d {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpB, _ := unpackArgSTransform2d(pB)
	__ret := C.godot_transform2d_operator_multiply(cpSelf, cpB)
	packSTransform2d(pB, cpB)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewTransform2dRef(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dNewIdentity function as declared in gdnative/transform2d.h:98
func Transform2dNewIdentity(rDest []Transform2d) {
	crDest, _ := unpackArgSTransform2d(rDest)
	C.godot_transform2d_new_identity(crDest)
	packSTransform2d(rDest, crDest)
}

// Transform2dXformRect2 function as declared in gdnative/transform2d.h:100
func Transform2dXformRect2(pSelf []Transform2d, pV []Rect2) Rect2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSRect2(pV)
	__ret := C.godot_transform2d_xform_rect2(cpSelf, cpV)
	packSRect2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// Transform2dXformInvRect2 function as declared in gdnative/transform2d.h:102
func Transform2dXformInvRect2(pSelf []Transform2d, pV []Rect2) Rect2 {
	cpSelf, _ := unpackArgSTransform2d(pSelf)
	cpV, _ := unpackArgSRect2(pV)
	__ret := C.godot_transform2d_xform_inv_rect2(cpSelf, cpV)
	packSRect2(pV, cpV)
	packSTransform2d(pSelf, cpSelf)
	__v := *NewRect2Ref(unsafe.Pointer(&__ret))
	return __v
}

// StringNameNew function as declared in gdnative/string_name.h:60
func StringNameNew(rDest []StringName, pName []String) {
	crDest, _ := unpackArgSStringName(rDest)
	cpName, _ := unpackArgSString(pName)
	C.godot_string_name_new(crDest, cpName)
	packSString(pName, cpName)
	packSStringName(rDest, crDest)
}

// StringNameNewData function as declared in gdnative/string_name.h:61
func StringNameNewData(rDest []StringName, pName string) {
	crDest, _ := unpackArgSStringName(rDest)
	cpName, _ := unpackPCharString(pName)
	C.godot_string_name_new_data(crDest, cpName)
	packSStringName(rDest, crDest)
}

// StringNameGetName function as declared in gdnative/string_name.h:63
func StringNameGetName(pSelf []StringName) String {
	cpSelf, _ := unpackArgSStringName(pSelf)
	__ret := C.godot_string_name_get_name(cpSelf)
	packSStringName(pSelf, cpSelf)
	__v := *NewStringRef(unsafe.Pointer(&__ret))
	return __v
}

// StringNameGetHash function as declared in gdnative/string_name.h:65
func StringNameGetHash(pSelf []StringName) uint32 {
	cpSelf, _ := unpackArgSStringName(pSelf)
	__ret := C.godot_string_name_get_hash(cpSelf)
	packSStringName(pSelf, cpSelf)
	__v := (uint32)(__ret)
	return __v
}

// StringNameGetDataUniquePointer function as declared in gdnative/string_name.h:66
func StringNameGetDataUniquePointer(pSelf []StringName) unsafe.Pointer {
	cpSelf, _ := unpackArgSStringName(pSelf)
	__ret := C.godot_string_name_get_data_unique_pointer(cpSelf)
	packSStringName(pSelf, cpSelf)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// StringNameOperatorEqual function as declared in gdnative/string_name.h:68
func StringNameOperatorEqual(pSelf []StringName, pOther []StringName) Bool {
	cpSelf, _ := unpackArgSStringName(pSelf)
	cpOther, _ := unpackArgSStringName(pOther)
	__ret := C.godot_string_name_operator_equal(cpSelf, cpOther)
	packSStringName(pOther, cpOther)
	packSStringName(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringNameOperatorLess function as declared in gdnative/string_name.h:69
func StringNameOperatorLess(pSelf []StringName, pOther []StringName) Bool {
	cpSelf, _ := unpackArgSStringName(pSelf)
	cpOther, _ := unpackArgSStringName(pOther)
	__ret := C.godot_string_name_operator_less(cpSelf, cpOther)
	packSStringName(pOther, cpOther)
	packSStringName(pSelf, cpSelf)
	__v := (Bool)(__ret)
	return __v
}

// StringNameDestroy function as declared in gdnative/string_name.h:71
func StringNameDestroy(pSelf []StringName) {
	cpSelf, _ := unpackArgSStringName(pSelf)
	C.godot_string_name_destroy(cpSelf)
	packSStringName(pSelf, cpSelf)
}

// ArvrRegisterInterface function as declared in arvr/godot_arvr.h:57
func ArvrRegisterInterface(pInterface []ArvrInterfaceGdnative) {
	cpInterface, _ := unpackArgSArvrInterfaceGdnative(pInterface)
	C.godot_arvr_register_interface(cpInterface)
	packSArvrInterfaceGdnative(pInterface, cpInterface)
}

// ArvrGetWorldscale function as declared in arvr/godot_arvr.h:60
func ArvrGetWorldscale() Real {
	__ret := C.godot_arvr_get_worldscale()
	__v := (Real)(__ret)
	return __v
}

// ArvrGetReferenceFrame function as declared in arvr/godot_arvr.h:61
func ArvrGetReferenceFrame() Transform {
	__ret := C.godot_arvr_get_reference_frame()
	__v := *NewTransformRef(unsafe.Pointer(&__ret))
	return __v
}

// ArvrBlit function as declared in arvr/godot_arvr.h:64
func ArvrBlit(pEye Int, pRenderTarget []Rid, pRect []Rect2) {
	cpEye, _ := (C.godot_int)(pEye), cgoAllocsUnknown
	cpRenderTarget, _ := unpackArgSRid(pRenderTarget)
	cpRect, _ := unpackArgSRect2(pRect)
	C.godot_arvr_blit(cpEye, cpRenderTarget, cpRect)
	packSRect2(pRect, cpRect)
	packSRid(pRenderTarget, cpRenderTarget)
}

// ArvrGetTexid function as declared in arvr/godot_arvr.h:65
func ArvrGetTexid(pRenderTarget []Rid) Int {
	cpRenderTarget, _ := unpackArgSRid(pRenderTarget)
	__ret := C.godot_arvr_get_texid(cpRenderTarget)
	packSRid(pRenderTarget, cpRenderTarget)
	__v := (Int)(__ret)
	return __v
}

// ArvrAddController function as declared in arvr/godot_arvr.h:68
func ArvrAddController(pDeviceName []byte, pHand Int, pTracksOrientation Bool, pTracksPosition Bool) Int {
	cpDeviceName, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&pDeviceName)).Data)), cgoAllocsUnknown
	cpHand, _ := (C.godot_int)(pHand), cgoAllocsUnknown
	cpTracksOrientation, _ := (C.godot_bool)(pTracksOrientation), cgoAllocsUnknown
	cpTracksPosition, _ := (C.godot_bool)(pTracksPosition), cgoAllocsUnknown
	__ret := C.godot_arvr_add_controller(cpDeviceName, cpHand, cpTracksOrientation, cpTracksPosition)
	__v := (Int)(__ret)
	return __v
}

// ArvrRemoveController function as declared in arvr/godot_arvr.h:69
func ArvrRemoveController(pControllerId Int) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	C.godot_arvr_remove_controller(cpControllerId)
}

// ArvrSetControllerTransform function as declared in arvr/godot_arvr.h:70
func ArvrSetControllerTransform(pControllerId Int, pTransform []Transform, pTracksOrientation Bool, pTracksPosition Bool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpTransform, _ := unpackArgSTransform(pTransform)
	cpTracksOrientation, _ := (C.godot_bool)(pTracksOrientation), cgoAllocsUnknown
	cpTracksPosition, _ := (C.godot_bool)(pTracksPosition), cgoAllocsUnknown
	C.godot_arvr_set_controller_transform(cpControllerId, cpTransform, cpTracksOrientation, cpTracksPosition)
	packSTransform(pTransform, cpTransform)
}

// ArvrSetControllerButton function as declared in arvr/godot_arvr.h:71
func ArvrSetControllerButton(pControllerId Int, pButton Int, pIsPressed Bool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpButton, _ := (C.godot_int)(pButton), cgoAllocsUnknown
	cpIsPressed, _ := (C.godot_bool)(pIsPressed), cgoAllocsUnknown
	C.godot_arvr_set_controller_button(cpControllerId, cpButton, cpIsPressed)
}

// ArvrSetControllerAxis function as declared in arvr/godot_arvr.h:72
func ArvrSetControllerAxis(pControllerId Int, pAxis Int, pValue Real, pCanBeNegative Bool) {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	cpAxis, _ := (C.godot_int)(pAxis), cgoAllocsUnknown
	cpValue, _ := (C.godot_real)(pValue), cgoAllocsUnknown
	cpCanBeNegative, _ := (C.godot_bool)(pCanBeNegative), cgoAllocsUnknown
	C.godot_arvr_set_controller_axis(cpControllerId, cpAxis, cpValue, cpCanBeNegative)
}

// ArvrGetControllerRumble function as declared in arvr/godot_arvr.h:73
func ArvrGetControllerRumble(pControllerId Int) Real {
	cpControllerId, _ := (C.godot_int)(pControllerId), cgoAllocsUnknown
	__ret := C.godot_arvr_get_controller_rumble(cpControllerId)
	__v := (Real)(__ret)
	return __v
}

// NativescriptRegisterClass function as declared in nativescript/godot_nativescript.h:133
func NativescriptRegisterClass(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc InstanceCreateFunc, pDestroyFunc InstanceDestroyFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpBase, _ := unpackPCharString(pBase)
	cpCreateFunc, _ := pCreateFunc.PassValue()
	cpDestroyFunc, _ := pDestroyFunc.PassValue()
	C.godot_nativescript_register_class(cpGdnativeHandle, cpName, cpBase, cpCreateFunc, cpDestroyFunc)
}

// NativescriptRegisterToolClass function as declared in nativescript/godot_nativescript.h:135
func NativescriptRegisterToolClass(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc InstanceCreateFunc, pDestroyFunc InstanceDestroyFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpBase, _ := unpackPCharString(pBase)
	cpCreateFunc, _ := pCreateFunc.PassValue()
	cpDestroyFunc, _ := pDestroyFunc.PassValue()
	C.godot_nativescript_register_tool_class(cpGdnativeHandle, cpName, cpBase, cpCreateFunc, cpDestroyFunc)
}

// NativescriptRegisterMethod function as declared in nativescript/godot_nativescript.h:148
func NativescriptRegisterMethod(pGdnativeHandle unsafe.Pointer, pName string, pFunctionName string, pAttr MethodAttributes, pMethod InstanceMethod) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpFunctionName, _ := unpackPCharString(pFunctionName)
	cpAttr, _ := pAttr.PassValue()
	cpMethod, _ := pMethod.PassValue()
	C.godot_nativescript_register_method(cpGdnativeHandle, cpName, cpFunctionName, cpAttr, cpMethod)
}

// NativescriptRegisterProperty function as declared in nativescript/godot_nativescript.h:164
func NativescriptRegisterProperty(pGdnativeHandle unsafe.Pointer, pName string, pPath string, pAttr []PropertyAttributes, pSetFunc PropertySetFunc, pGetFunc PropertyGetFunc) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpPath, _ := unpackPCharString(pPath)
	cpAttr, _ := unpackArgSPropertyAttributes(pAttr)
	cpSetFunc, _ := pSetFunc.PassValue()
	cpGetFunc, _ := pGetFunc.PassValue()
	C.godot_nativescript_register_property(cpGdnativeHandle, cpName, cpPath, cpAttr, cpSetFunc, cpGetFunc)
	packSPropertyAttributes(pAttr, cpAttr)
}

// NativescriptRegisterSignal function as declared in nativescript/godot_nativescript.h:183
func NativescriptRegisterSignal(pGdnativeHandle unsafe.Pointer, pName string, pSignal []Signal) {
	cpGdnativeHandle, _ := pGdnativeHandle, cgoAllocsUnknown
	cpName, _ := unpackPCharString(pName)
	cpSignal, _ := unpackArgSSignal(pSignal)
	C.godot_nativescript_register_signal(cpGdnativeHandle, cpName, cpSignal)
	packSSignal(pSignal, cpSignal)
}

// NativescriptGetUserdata function as declared in nativescript/godot_nativescript.h:185
func NativescriptGetUserdata(pInstance *Object) unsafe.Pointer {
	cpInstance, _ := (C.godot_object)(unsafe.Pointer(pInstance)), cgoAllocsUnknown
	__ret := C.godot_nativescript_get_userdata(cpInstance)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// PluginscriptRegisterLanguage function as declared in pluginscript/godot_pluginscript.h:165
func PluginscriptRegisterLanguage(languageDesc []PluginscriptLanguageDesc) {
	clanguageDesc, _ := unpackArgSPluginscriptLanguageDesc(languageDesc)
	C.godot_pluginscript_register_language(clanguageDesc)
	packSPluginscriptLanguageDesc(languageDesc, clanguageDesc)
}
