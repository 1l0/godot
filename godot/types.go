// WARNING: This file has automatically been generated on Sun, 24 Dec 2017 06:43:30 JST.
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

// GdnativeExtNativescriptApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:52
type GdnativeExtNativescriptApiStruct struct {
	Type                          uint32
	Version                       GdnativeApiVersion
	Next                          []GdnativeApiStruct
	NativescriptRegisterClass     *func(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc InstanceCreateFunc, pDestroyFunc InstanceDestroyFunc)
	NativescriptRegisterToolClass *func(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc InstanceCreateFunc, pDestroyFunc InstanceDestroyFunc)
	NativescriptRegisterMethod    *func(pGdnativeHandle unsafe.Pointer, pName string, pFunctionName string, pAttr MethodAttributes, pMethod InstanceMethod)
	NativescriptRegisterProperty  *func(pGdnativeHandle unsafe.Pointer, pName string, pPath string, pAttr []PropertyAttributes, pSetFunc PropertySetFunc, pGetFunc PropertyGetFunc)
	NativescriptRegisterSignal    *func(pGdnativeHandle unsafe.Pointer, pName string, pSignal []Signal)
	NativescriptGetUserdata       *func(pInstance *Object) unsafe.Pointer
	reff0cd6324                   *C.godot_gdnative_ext_nativescript_api_struct
	allocsf0cd6324                interface{}
}

// GdnativeExtPluginscriptApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:59
type GdnativeExtPluginscriptApiStruct struct {
	Type                         uint32
	Version                      GdnativeApiVersion
	Next                         []GdnativeApiStruct
	PluginscriptRegisterLanguage *func(languageDesc []PluginscriptLanguageDesc)
	refc52e13a7                  *C.godot_gdnative_ext_pluginscript_api_struct
	allocsc52e13a7               interface{}
}

// GdnativeExtArvrApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:76
type GdnativeExtArvrApiStruct struct {
	Type                       uint32
	Version                    GdnativeApiVersion
	Next                       []GdnativeApiStruct
	ArvrRegisterInterface      *func(pInterface []ArvrInterfaceGdnative)
	ArvrGetWorldscale          Real
	ArvrGetReferenceFrame      Transform
	ArvrBlit                   *func(pEye int32, pRenderTarget []Rid, pScreenRect []Rect2)
	ArvrGetTexid               Int
	ArvrAddController          Int
	ArvrRemoveController       *func(pControllerId Int)
	ArvrSetControllerTransform *func(pControllerId Int, pTransform []Transform, pTracksOrientation Bool, pTracksPosition Bool)
	ArvrSetControllerButton    *func(pControllerId Int, pButton Int, pIsPressed Bool)
	ArvrSetControllerAxis      *func(pControllerId Int, pExis Int, pValue Real, pCanBeNegative Bool)
	ArvrGetControllerRumble    Real
	ref64de5bf5                *C.godot_gdnative_ext_arvr_api_struct
	allocs64de5bf5             interface{}
}

// GdnativeCoreApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:809
type GdnativeCoreApiStruct struct {
	Type                                      uint32
	Version                                   GdnativeApiVersion
	Next                                      []GdnativeApiStruct
	NumExtensions                             uint32
	Extensions                                [][]GdnativeApiStruct
	ColorNewRgba                              *func(rDest []Color, pR Real, pG Real, pB Real, pA Real)
	ColorNewRgb                               *func(rDest []Color, pR Real, pG Real, pB Real)
	ColorGetR                                 Real
	ColorSetR                                 *func(pSelf []Color, r Real)
	ColorGetG                                 Real
	ColorSetG                                 *func(pSelf []Color, g Real)
	ColorGetB                                 Real
	ColorSetB                                 *func(pSelf []Color, b Real)
	ColorGetA                                 Real
	ColorSetA                                 *func(pSelf []Color, a Real)
	ColorGetH                                 Real
	ColorGetS                                 Real
	ColorGetV                                 Real
	ColorAsString                             String
	ColorToRgba32                             Int
	ColorToArgb32                             Int
	ColorGray                                 Real
	ColorInverted                             Color
	ColorContrasted                           Color
	ColorLinearInterpolate                    Color
	ColorBlend                                Color
	ColorToHtml                               String
	ColorOperatorEqual                        Bool
	ColorOperatorLess                         Bool
	Vector2New                                *func(rDest []Vector2, pX Real, pY Real)
	Vector2AsString                           String
	Vector2Normalized                         Vector2
	Vector2Length                             Real
	Vector2Angle                              Real
	Vector2LengthSquared                      Real
	Vector2IsNormalized                       Bool
	Vector2DistanceTo                         Real
	Vector2DistanceSquaredTo                  Real
	Vector2AngleTo                            Real
	Vector2AngleToPoint                       Real
	Vector2LinearInterpolate                  Vector2
	Vector2CubicInterpolate                   Vector2
	Vector2Rotated                            Vector2
	Vector2Tangent                            Vector2
	Vector2Floor                              Vector2
	Vector2Snapped                            Vector2
	Vector2Aspect                             Real
	Vector2Dot                                Real
	Vector2Slide                              Vector2
	Vector2Bounce                             Vector2
	Vector2Reflect                            Vector2
	Vector2Abs                                Vector2
	Vector2Clamped                            Vector2
	Vector2OperatorAdd                        Vector2
	Vector2OperatorSubstract                  Vector2
	Vector2OperatorMultiplyVector             Vector2
	Vector2OperatorMultiplyScalar             Vector2
	Vector2OperatorDivideVector               Vector2
	Vector2OperatorDivideScalar               Vector2
	Vector2OperatorEqual                      Bool
	Vector2OperatorLess                       Bool
	Vector2OperatorNeg                        Vector2
	Vector2SetX                               *func(pSelf []Vector2, pX Real)
	Vector2SetY                               *func(pSelf []Vector2, pY Real)
	Vector2GetX                               Real
	Vector2GetY                               Real
	QuatNew                                   *func(rDest []Quat, pX Real, pY Real, pZ Real, pW Real)
	QuatNewWithAxisAngle                      *func(rDest []Quat, pAxis []Vector3, pAngle Real)
	QuatGetX                                  Real
	QuatSetX                                  *func(pSelf []Quat, val Real)
	QuatGetY                                  Real
	QuatSetY                                  *func(pSelf []Quat, val Real)
	QuatGetZ                                  Real
	QuatSetZ                                  *func(pSelf []Quat, val Real)
	QuatGetW                                  Real
	QuatSetW                                  *func(pSelf []Quat, val Real)
	QuatAsString                              String
	QuatLength                                Real
	QuatLengthSquared                         Real
	QuatNormalized                            Quat
	QuatIsNormalized                          Bool
	QuatInverse                               Quat
	QuatDot                                   Real
	QuatXform                                 Vector3
	QuatSlerp                                 Quat
	QuatSlerpni                               Quat
	QuatCubicSlerp                            Quat
	QuatOperatorMultiply                      Quat
	QuatOperatorAdd                           Quat
	QuatOperatorSubstract                     Quat
	QuatOperatorDivide                        Quat
	QuatOperatorEqual                         Bool
	QuatOperatorNeg                           Quat
	BasisNewWithRows                          *func(rDest []Basis, pXAxis []Vector3, pYAxis []Vector3, pZAxis []Vector3)
	BasisNewWithAxisAndAngle                  *func(rDest []Basis, pAxis []Vector3, pPhi Real)
	BasisNewWithEuler                         *func(rDest []Basis, pEuler []Vector3)
	BasisAsString                             String
	BasisInverse                              Basis
	BasisTransposed                           Basis
	BasisOrthonormalized                      Basis
	BasisDeterminant                          Real
	BasisRotated                              Basis
	BasisScaled                               Basis
	BasisGetScale                             Vector3
	BasisGetEuler                             Vector3
	BasisTdotx                                Real
	BasisTdoty                                Real
	BasisTdotz                                Real
	BasisXform                                Vector3
	BasisXformInv                             Vector3
	BasisGetOrthogonalIndex                   Int
	BasisNew                                  *func(rDest []Basis)
	BasisNewWithEulerQuat                     *func(rDest []Basis, pEuler []Quat)
	BasisGetElements                          *func(pSelf []Basis, pElements []Vector3)
	BasisGetAxis                              Vector3
	BasisSetAxis                              *func(pSelf []Basis, pAxis Int, pValue []Vector3)
	BasisGetRow                               Vector3
	BasisSetRow                               *func(pSelf []Basis, pRow Int, pValue []Vector3)
	BasisOperatorEqual                        Bool
	BasisOperatorAdd                          Basis
	BasisOperatorSubstract                    Basis
	BasisOperatorMultiplyVector               Basis
	BasisOperatorMultiplyScalar               Basis
	Vector3New                                *func(rDest []Vector3, pX Real, pY Real, pZ Real)
	Vector3AsString                           String
	Vector3MinAxis                            Int
	Vector3MaxAxis                            Int
	Vector3Length                             Real
	Vector3LengthSquared                      Real
	Vector3IsNormalized                       Bool
	Vector3Normalized                         Vector3
	Vector3Inverse                            Vector3
	Vector3Snapped                            Vector3
	Vector3Rotated                            Vector3
	Vector3LinearInterpolate                  Vector3
	Vector3CubicInterpolate                   Vector3
	Vector3Dot                                Real
	Vector3Cross                              Vector3
	Vector3Outer                              Basis
	Vector3ToDiagonalMatrix                   Basis
	Vector3Abs                                Vector3
	Vector3Floor                              Vector3
	Vector3Ceil                               Vector3
	Vector3DistanceTo                         Real
	Vector3DistanceSquaredTo                  Real
	Vector3AngleTo                            Real
	Vector3Slide                              Vector3
	Vector3Bounce                             Vector3
	Vector3Reflect                            Vector3
	Vector3OperatorAdd                        Vector3
	Vector3OperatorSubstract                  Vector3
	Vector3OperatorMultiplyVector             Vector3
	Vector3OperatorMultiplyScalar             Vector3
	Vector3OperatorDivideVector               Vector3
	Vector3OperatorDivideScalar               Vector3
	Vector3OperatorEqual                      Bool
	Vector3OperatorLess                       Bool
	Vector3OperatorNeg                        Vector3
	Vector3SetAxis                            *func(pSelf []Vector3, pAxis Vector3Axis, pVal Real)
	Vector3GetAxis                            Real
	PoolByteArrayNew                          *func(rDest []PoolByteArray)
	PoolByteArrayNewCopy                      *func(rDest []PoolByteArray, pSrc []PoolByteArray)
	PoolByteArrayNewWithArray                 *func(rDest []PoolByteArray, pA []Array)
	PoolByteArrayAppend                       *func(pSelf []PoolByteArray, pData byte)
	PoolByteArrayAppendArray                  *func(pSelf []PoolByteArray, pArray []PoolByteArray)
	PoolByteArrayInsert                       Error
	PoolByteArrayInvert                       *func(pSelf []PoolByteArray)
	PoolByteArrayPushBack                     *func(pSelf []PoolByteArray, pData byte)
	PoolByteArrayRemove                       *func(pSelf []PoolByteArray, pIdx Int)
	PoolByteArrayResize                       *func(pSelf []PoolByteArray, pSize Int)
	PoolByteArrayRead                         PoolByteArrayReadAccess
	PoolByteArrayWrite                        PoolByteArrayWriteAccess
	PoolByteArraySet                          *func(pSelf []PoolByteArray, pIdx Int, pData byte)
	PoolByteArrayGet                          Uint8T
	PoolByteArraySize                         Int
	PoolByteArrayDestroy                      *func(pSelf []PoolByteArray)
	PoolIntArrayNew                           *func(rDest []PoolIntArray)
	PoolIntArrayNewCopy                       *func(rDest []PoolIntArray, pSrc []PoolIntArray)
	PoolIntArrayNewWithArray                  *func(rDest []PoolIntArray, pA []Array)
	PoolIntArrayAppend                        *func(pSelf []PoolIntArray, pData Int)
	PoolIntArrayAppendArray                   *func(pSelf []PoolIntArray, pArray []PoolIntArray)
	PoolIntArrayInsert                        Error
	PoolIntArrayInvert                        *func(pSelf []PoolIntArray)
	PoolIntArrayPushBack                      *func(pSelf []PoolIntArray, pData Int)
	PoolIntArrayRemove                        *func(pSelf []PoolIntArray, pIdx Int)
	PoolIntArrayResize                        *func(pSelf []PoolIntArray, pSize Int)
	PoolIntArrayRead                          PoolIntArrayReadAccess
	PoolIntArrayWrite                         PoolIntArrayWriteAccess
	PoolIntArraySet                           *func(pSelf []PoolIntArray, pIdx Int, pData Int)
	PoolIntArrayGet                           Int
	PoolIntArraySize                          Int
	PoolIntArrayDestroy                       *func(pSelf []PoolIntArray)
	PoolRealArrayNew                          *func(rDest []PoolRealArray)
	PoolRealArrayNewCopy                      *func(rDest []PoolRealArray, pSrc []PoolRealArray)
	PoolRealArrayNewWithArray                 *func(rDest []PoolRealArray, pA []Array)
	PoolRealArrayAppend                       *func(pSelf []PoolRealArray, pData Real)
	PoolRealArrayAppendArray                  *func(pSelf []PoolRealArray, pArray []PoolRealArray)
	PoolRealArrayInsert                       Error
	PoolRealArrayInvert                       *func(pSelf []PoolRealArray)
	PoolRealArrayPushBack                     *func(pSelf []PoolRealArray, pData Real)
	PoolRealArrayRemove                       *func(pSelf []PoolRealArray, pIdx Int)
	PoolRealArrayResize                       *func(pSelf []PoolRealArray, pSize Int)
	PoolRealArrayRead                         PoolRealArrayReadAccess
	PoolRealArrayWrite                        PoolRealArrayWriteAccess
	PoolRealArraySet                          *func(pSelf []PoolRealArray, pIdx Int, pData Real)
	PoolRealArrayGet                          Real
	PoolRealArraySize                         Int
	PoolRealArrayDestroy                      *func(pSelf []PoolRealArray)
	PoolStringArrayNew                        *func(rDest []PoolStringArray)
	PoolStringArrayNewCopy                    *func(rDest []PoolStringArray, pSrc []PoolStringArray)
	PoolStringArrayNewWithArray               *func(rDest []PoolStringArray, pA []Array)
	PoolStringArrayAppend                     *func(pSelf []PoolStringArray, pData []String)
	PoolStringArrayAppendArray                *func(pSelf []PoolStringArray, pArray []PoolStringArray)
	PoolStringArrayInsert                     Error
	PoolStringArrayInvert                     *func(pSelf []PoolStringArray)
	PoolStringArrayPushBack                   *func(pSelf []PoolStringArray, pData []String)
	PoolStringArrayRemove                     *func(pSelf []PoolStringArray, pIdx Int)
	PoolStringArrayResize                     *func(pSelf []PoolStringArray, pSize Int)
	PoolStringArrayRead                       PoolStringArrayReadAccess
	PoolStringArrayWrite                      PoolStringArrayWriteAccess
	PoolStringArraySet                        *func(pSelf []PoolStringArray, pIdx Int, pData []String)
	PoolStringArrayGet                        String
	PoolStringArraySize                       Int
	PoolStringArrayDestroy                    *func(pSelf []PoolStringArray)
	PoolVector2ArrayNew                       *func(rDest []PoolVector2Array)
	PoolVector2ArrayNewCopy                   *func(rDest []PoolVector2Array, pSrc []PoolVector2Array)
	PoolVector2ArrayNewWithArray              *func(rDest []PoolVector2Array, pA []Array)
	PoolVector2ArrayAppend                    *func(pSelf []PoolVector2Array, pData []Vector2)
	PoolVector2ArrayAppendArray               *func(pSelf []PoolVector2Array, pArray []PoolVector2Array)
	PoolVector2ArrayInsert                    Error
	PoolVector2ArrayInvert                    *func(pSelf []PoolVector2Array)
	PoolVector2ArrayPushBack                  *func(pSelf []PoolVector2Array, pData []Vector2)
	PoolVector2ArrayRemove                    *func(pSelf []PoolVector2Array, pIdx Int)
	PoolVector2ArrayResize                    *func(pSelf []PoolVector2Array, pSize Int)
	PoolVector2ArrayRead                      PoolVector2ArrayReadAccess
	PoolVector2ArrayWrite                     PoolVector2ArrayWriteAccess
	PoolVector2ArraySet                       *func(pSelf []PoolVector2Array, pIdx Int, pData []Vector2)
	PoolVector2ArrayGet                       Vector2
	PoolVector2ArraySize                      Int
	PoolVector2ArrayDestroy                   *func(pSelf []PoolVector2Array)
	PoolVector3ArrayNew                       *func(rDest []PoolVector3Array)
	PoolVector3ArrayNewCopy                   *func(rDest []PoolVector3Array, pSrc []PoolVector3Array)
	PoolVector3ArrayNewWithArray              *func(rDest []PoolVector3Array, pA []Array)
	PoolVector3ArrayAppend                    *func(pSelf []PoolVector3Array, pData []Vector3)
	PoolVector3ArrayAppendArray               *func(pSelf []PoolVector3Array, pArray []PoolVector3Array)
	PoolVector3ArrayInsert                    Error
	PoolVector3ArrayInvert                    *func(pSelf []PoolVector3Array)
	PoolVector3ArrayPushBack                  *func(pSelf []PoolVector3Array, pData []Vector3)
	PoolVector3ArrayRemove                    *func(pSelf []PoolVector3Array, pIdx Int)
	PoolVector3ArrayResize                    *func(pSelf []PoolVector3Array, pSize Int)
	PoolVector3ArrayRead                      PoolVector3ArrayReadAccess
	PoolVector3ArrayWrite                     PoolVector3ArrayWriteAccess
	PoolVector3ArraySet                       *func(pSelf []PoolVector3Array, pIdx Int, pData []Vector3)
	PoolVector3ArrayGet                       Vector3
	PoolVector3ArraySize                      Int
	PoolVector3ArrayDestroy                   *func(pSelf []PoolVector3Array)
	PoolColorArrayNew                         *func(rDest []PoolColorArray)
	PoolColorArrayNewCopy                     *func(rDest []PoolColorArray, pSrc []PoolColorArray)
	PoolColorArrayNewWithArray                *func(rDest []PoolColorArray, pA []Array)
	PoolColorArrayAppend                      *func(pSelf []PoolColorArray, pData []Color)
	PoolColorArrayAppendArray                 *func(pSelf []PoolColorArray, pArray []PoolColorArray)
	PoolColorArrayInsert                      Error
	PoolColorArrayInvert                      *func(pSelf []PoolColorArray)
	PoolColorArrayPushBack                    *func(pSelf []PoolColorArray, pData []Color)
	PoolColorArrayRemove                      *func(pSelf []PoolColorArray, pIdx Int)
	PoolColorArrayResize                      *func(pSelf []PoolColorArray, pSize Int)
	PoolColorArrayRead                        PoolColorArrayReadAccess
	PoolColorArrayWrite                       PoolColorArrayWriteAccess
	PoolColorArraySet                         *func(pSelf []PoolColorArray, pIdx Int, pData []Color)
	PoolColorArrayGet                         Color
	PoolColorArraySize                        Int
	PoolColorArrayDestroy                     *func(pSelf []PoolColorArray)
	PoolByteArrayReadAccessPtr                Uint8T
	PoolByteArrayReadAccessOperatorAssign     *func(pRead []PoolByteArrayReadAccess, pOther []PoolByteArrayReadAccess)
	PoolByteArrayReadAccessDestroy            *func(pRead []PoolByteArrayReadAccess)
	PoolIntArrayReadAccessPtr                 Int
	PoolIntArrayReadAccessOperatorAssign      *func(pRead []PoolIntArrayReadAccess, pOther []PoolIntArrayReadAccess)
	PoolIntArrayReadAccessDestroy             *func(pRead []PoolIntArrayReadAccess)
	PoolRealArrayReadAccessPtr                Real
	PoolRealArrayReadAccessOperatorAssign     *func(pRead []PoolRealArrayReadAccess, pOther []PoolRealArrayReadAccess)
	PoolRealArrayReadAccessDestroy            *func(pRead []PoolRealArrayReadAccess)
	PoolStringArrayReadAccessPtr              String
	PoolStringArrayReadAccessOperatorAssign   *func(pRead []PoolStringArrayReadAccess, pOther []PoolStringArrayReadAccess)
	PoolStringArrayReadAccessDestroy          *func(pRead []PoolStringArrayReadAccess)
	PoolVector2ArrayReadAccessPtr             Vector2
	PoolVector2ArrayReadAccessOperatorAssign  *func(pRead []PoolVector2ArrayReadAccess, pOther []PoolVector2ArrayReadAccess)
	PoolVector2ArrayReadAccessDestroy         *func(pRead []PoolVector2ArrayReadAccess)
	PoolVector3ArrayReadAccessPtr             Vector3
	PoolVector3ArrayReadAccessOperatorAssign  *func(pRead []PoolVector3ArrayReadAccess, pOther []PoolVector3ArrayReadAccess)
	PoolVector3ArrayReadAccessDestroy         *func(pRead []PoolVector3ArrayReadAccess)
	PoolColorArrayReadAccessPtr               Color
	PoolColorArrayReadAccessOperatorAssign    *func(pRead []PoolColorArrayReadAccess, pOther []PoolColorArrayReadAccess)
	PoolColorArrayReadAccessDestroy           *func(pRead []PoolColorArrayReadAccess)
	PoolByteArrayWriteAccessPtr               Uint8T
	PoolByteArrayWriteAccessOperatorAssign    *func(pWrite []PoolByteArrayWriteAccess, pOther []PoolByteArrayWriteAccess)
	PoolByteArrayWriteAccessDestroy           *func(pWrite []PoolByteArrayWriteAccess)
	PoolIntArrayWriteAccessPtr                Int
	PoolIntArrayWriteAccessOperatorAssign     *func(pWrite []PoolIntArrayWriteAccess, pOther []PoolIntArrayWriteAccess)
	PoolIntArrayWriteAccessDestroy            *func(pWrite []PoolIntArrayWriteAccess)
	PoolRealArrayWriteAccessPtr               Real
	PoolRealArrayWriteAccessOperatorAssign    *func(pWrite []PoolRealArrayWriteAccess, pOther []PoolRealArrayWriteAccess)
	PoolRealArrayWriteAccessDestroy           *func(pWrite []PoolRealArrayWriteAccess)
	PoolStringArrayWriteAccessPtr             String
	PoolStringArrayWriteAccessOperatorAssign  *func(pWrite []PoolStringArrayWriteAccess, pOther []PoolStringArrayWriteAccess)
	PoolStringArrayWriteAccessDestroy         *func(pWrite []PoolStringArrayWriteAccess)
	PoolVector2ArrayWriteAccessPtr            Vector2
	PoolVector2ArrayWriteAccessOperatorAssign *func(pWrite []PoolVector2ArrayWriteAccess, pOther []PoolVector2ArrayWriteAccess)
	PoolVector2ArrayWriteAccessDestroy        *func(pWrite []PoolVector2ArrayWriteAccess)
	PoolVector3ArrayWriteAccessPtr            Vector3
	PoolVector3ArrayWriteAccessOperatorAssign *func(pWrite []PoolVector3ArrayWriteAccess, pOther []PoolVector3ArrayWriteAccess)
	PoolVector3ArrayWriteAccessDestroy        *func(pWrite []PoolVector3ArrayWriteAccess)
	PoolColorArrayWriteAccessPtr              Color
	PoolColorArrayWriteAccessOperatorAssign   *func(pWrite []PoolColorArrayWriteAccess, pOther []PoolColorArrayWriteAccess)
	PoolColorArrayWriteAccessDestroy          *func(pWrite []PoolColorArrayWriteAccess)
	ArrayNew                                  *func(rDest []Array)
	ArrayNewCopy                              *func(rDest []Array, pSrc []Array)
	ArrayNewPoolColorArray                    *func(rDest []Array, pPca []PoolColorArray)
	ArrayNewPoolVector3Array                  *func(rDest []Array, pPv3a []PoolVector3Array)
	ArrayNewPoolVector2Array                  *func(rDest []Array, pPv2a []PoolVector2Array)
	ArrayNewPoolStringArray                   *func(rDest []Array, pPsa []PoolStringArray)
	ArrayNewPoolRealArray                     *func(rDest []Array, pPra []PoolRealArray)
	ArrayNewPoolIntArray                      *func(rDest []Array, pPia []PoolIntArray)
	ArrayNewPoolByteArray                     *func(rDest []Array, pPba []PoolByteArray)
	ArraySet                                  *func(pSelf []Array, pIdx Int, pValue []Variant)
	ArrayGet                                  Variant
	ArrayOperatorIndex                        Variant
	ArrayOperatorIndexConst                   Variant
	ArrayAppend                               *func(pSelf []Array, pValue []Variant)
	ArrayClear                                *func(pSelf []Array)
	ArrayCount                                Int
	ArrayEmpty                                Bool
	ArrayErase                                *func(pSelf []Array, pValue []Variant)
	ArrayFront                                Variant
	ArrayBack                                 Variant
	ArrayFind                                 Int
	ArrayFindLast                             Int
	ArrayHas                                  Bool
	ArrayHash                                 Int
	ArrayInsert                               *func(pSelf []Array, pPos Int, pValue []Variant)
	ArrayInvert                               *func(pSelf []Array)
	ArrayPopBack                              Variant
	ArrayPopFront                             Variant
	ArrayPushBack                             *func(pSelf []Array, pValue []Variant)
	ArrayPushFront                            *func(pSelf []Array, pValue []Variant)
	ArrayRemove                               *func(pSelf []Array, pIdx Int)
	ArrayResize                               *func(pSelf []Array, pSize Int)
	ArrayRfind                                Int
	ArraySize                                 Int
	ArraySort                                 *func(pSelf []Array)
	ArraySortCustom                           *func(pSelf []Array, pObj *Object, pFunc []String)
	ArrayBsearch                              Int
	ArrayBsearchCustom                        Int
	ArrayDestroy                              *func(pSelf []Array)
	DictionaryNew                             *func(rDest []Dictionary)
	DictionaryNewCopy                         *func(rDest []Dictionary, pSrc []Dictionary)
	DictionaryDestroy                         *func(pSelf []Dictionary)
	DictionarySize                            Int
	DictionaryEmpty                           Bool
	DictionaryClear                           *func(pSelf []Dictionary)
	DictionaryHas                             Bool
	DictionaryHasAll                          Bool
	DictionaryErase                           *func(pSelf []Dictionary, pKey []Variant)
	DictionaryHash                            Int
	DictionaryKeys                            Array
	DictionaryValues                          Array
	DictionaryGet                             Variant
	DictionarySet                             *func(pSelf []Dictionary, pKey []Variant, pValue []Variant)
	DictionaryOperatorIndex                   Variant
	DictionaryOperatorIndexConst              Variant
	DictionaryNext                            Variant
	DictionaryOperatorEqual                   Bool
	DictionaryToJson                          String
	NodePathNew                               *func(rDest []NodePath, pFrom []String)
	NodePathNewCopy                           *func(rDest []NodePath, pSrc []NodePath)
	NodePathDestroy                           *func(pSelf []NodePath)
	NodePathAsString                          String
	NodePathIsAbsolute                        Bool
	NodePathGetNameCount                      Int
	NodePathGetName                           String
	NodePathGetSubnameCount                   Int
	NodePathGetSubname                        String
	NodePathGetConcatenatedSubnames           String
	NodePathIsEmpty                           Bool
	NodePathOperatorEqual                     Bool
	PlaneNewWithReals                         *func(rDest []Plane, pA Real, pB Real, pC Real, pD Real)
	PlaneNewWithVectors                       *func(rDest []Plane, pV1 []Vector3, pV2 []Vector3, pV3 []Vector3)
	PlaneNewWithNormal                        *func(rDest []Plane, pNormal []Vector3, pD Real)
	PlaneAsString                             String
	PlaneNormalized                           Plane
	PlaneCenter                               Vector3
	PlaneGetAnyPoint                          Vector3
	PlaneIsPointOver                          Bool
	PlaneDistanceTo                           Real
	PlaneHasPoint                             Bool
	PlaneProject                              Vector3
	PlaneIntersect3                           Bool
	PlaneIntersectsRay                        Bool
	PlaneIntersectsSegment                    Bool
	PlaneOperatorNeg                          Plane
	PlaneOperatorEqual                        Bool
	PlaneSetNormal                            *func(pSelf []Plane, pNormal []Vector3)
	PlaneGetNormal                            Vector3
	PlaneGetD                                 Real
	PlaneSetD                                 *func(pSelf []Plane, pD Real)
	Rect2NewWithPositionAndSize               *func(rDest []Rect2, pPos []Vector2, pSize []Vector2)
	Rect2New                                  *func(rDest []Rect2, pX Real, pY Real, pWidth Real, pHeight Real)
	Rect2AsString                             String
	Rect2GetArea                              Real
	Rect2Intersects                           Bool
	Rect2Encloses                             Bool
	Rect2HasNoArea                            Bool
	Rect2Clip                                 Rect2
	Rect2Merge                                Rect2
	Rect2HasPoint                             Bool
	Rect2Grow                                 Rect2
	Rect2Expand                               Rect2
	Rect2OperatorEqual                        Bool
	Rect2GetPosition                          Vector2
	Rect2GetSize                              Vector2
	Rect2SetPosition                          *func(pSelf []Rect2, pPos []Vector2)
	Rect2SetSize                              *func(pSelf []Rect2, pSize []Vector2)
	AabbNew                                   *func(rDest []Aabb, pPos []Vector3, pSize []Vector3)
	AabbGetPosition                           Vector3
	AabbSetPosition                           *func(pSelf []Aabb, pV []Vector3)
	AabbGetSize                               Vector3
	AabbSetSize                               *func(pSelf []Aabb, pV []Vector3)
	AabbAsString                              String
	AabbGetArea                               Real
	AabbHasNoArea                             Bool
	AabbHasNoSurface                          Bool
	AabbIntersects                            Bool
	AabbEncloses                              Bool
	AabbMerge                                 Aabb
	AabbIntersection                          Aabb
	AabbIntersectsPlane                       Bool
	AabbIntersectsSegment                     Bool
	AabbHasPoint                              Bool
	AabbGetSupport                            Vector3
	AabbGetLongestAxis                        Vector3
	AabbGetLongestAxisIndex                   Int
	AabbGetLongestAxisSize                    Real
	AabbGetShortestAxis                       Vector3
	AabbGetShortestAxisIndex                  Int
	AabbGetShortestAxisSize                   Real
	AabbExpand                                Aabb
	AabbGrow                                  Aabb
	AabbGetEndpoint                           Vector3
	AabbOperatorEqual                         Bool
	RidNew                                    *func(rDest []Rid)
	RidGetId                                  Int
	RidNewWithResource                        *func(rDest []Rid, pFrom *Object)
	RidOperatorEqual                          Bool
	RidOperatorLess                           Bool
	TransformNewWithAxisOrigin                *func(rDest []Transform, pXAxis []Vector3, pYAxis []Vector3, pZAxis []Vector3, pOrigin []Vector3)
	TransformNew                              *func(rDest []Transform, pBasis []Basis, pOrigin []Vector3)
	TransformGetBasis                         Basis
	TransformSetBasis                         *func(pSelf []Transform, pV []Basis)
	TransformGetOrigin                        Vector3
	TransformSetOrigin                        *func(pSelf []Transform, pV []Vector3)
	TransformAsString                         String
	TransformInverse                          Transform
	TransformAffineInverse                    Transform
	TransformOrthonormalized                  Transform
	TransformRotated                          Transform
	TransformScaled                           Transform
	TransformTranslated                       Transform
	TransformLookingAt                        Transform
	TransformXformPlane                       Plane
	TransformXformInvPlane                    Plane
	TransformNewIdentity                      *func(rDest []Transform)
	TransformOperatorEqual                    Bool
	TransformOperatorMultiply                 Transform
	TransformXformVector3                     Vector3
	TransformXformInvVector3                  Vector3
	TransformXformAabb                        Aabb
	TransformXformInvAabb                     Aabb
	Transform2dNew                            *func(rDest []Transform2d, pRot Real, pPos []Vector2)
	Transform2dNewAxisOrigin                  *func(rDest []Transform2d, pXAxis []Vector2, pYAxis []Vector2, pOrigin []Vector2)
	Transform2dAsString                       String
	Transform2dInverse                        Transform2d
	Transform2dAffineInverse                  Transform2d
	Transform2dGetRotation                    Real
	Transform2dGetOrigin                      Vector2
	Transform2dGetScale                       Vector2
	Transform2dOrthonormalized                Transform2d
	Transform2dRotated                        Transform2d
	Transform2dScaled                         Transform2d
	Transform2dTranslated                     Transform2d
	Transform2dXformVector2                   Vector2
	Transform2dXformInvVector2                Vector2
	Transform2dBasisXformVector2              Vector2
	Transform2dBasisXformInvVector2           Vector2
	Transform2dInterpolateWith                Transform2d
	Transform2dOperatorEqual                  Bool
	Transform2dOperatorMultiply               Transform2d
	Transform2dNewIdentity                    *func(rDest []Transform2d)
	Transform2dXformRect2                     Rect2
	Transform2dXformInvRect2                  Rect2
	VariantGetType                            VariantType
	VariantNewCopy                            *func(rDest []Variant, pSrc []Variant)
	VariantNewNil                             *func(rDest []Variant)
	VariantNewBool                            *func(pV []Variant, pB Bool)
	VariantNewUint                            *func(rDest []Variant, pI uint64)
	VariantNewInt                             *func(rDest []Variant, pI int64)
	VariantNewReal                            *func(rDest []Variant, pR float64)
	VariantNewString                          *func(rDest []Variant, pS []String)
	VariantNewVector2                         *func(rDest []Variant, pV2 []Vector2)
	VariantNewRect2                           *func(rDest []Variant, pRect2 []Rect2)
	VariantNewVector3                         *func(rDest []Variant, pV3 []Vector3)
	VariantNewTransform2d                     *func(rDest []Variant, pT2d []Transform2d)
	VariantNewPlane                           *func(rDest []Variant, pPlane []Plane)
	VariantNewQuat                            *func(rDest []Variant, pQuat []Quat)
	VariantNewAabb                            *func(rDest []Variant, pAabb []Aabb)
	VariantNewBasis                           *func(rDest []Variant, pBasis []Basis)
	VariantNewTransform                       *func(rDest []Variant, pTrans []Transform)
	VariantNewColor                           *func(rDest []Variant, pColor []Color)
	VariantNewNodePath                        *func(rDest []Variant, pNp []NodePath)
	VariantNewRid                             *func(rDest []Variant, pRid []Rid)
	VariantNewObject                          *func(rDest []Variant, pObj *Object)
	VariantNewDictionary                      *func(rDest []Variant, pDict []Dictionary)
	VariantNewArray                           *func(rDest []Variant, pArr []Array)
	VariantNewPoolByteArray                   *func(rDest []Variant, pPba []PoolByteArray)
	VariantNewPoolIntArray                    *func(rDest []Variant, pPia []PoolIntArray)
	VariantNewPoolRealArray                   *func(rDest []Variant, pPra []PoolRealArray)
	VariantNewPoolStringArray                 *func(rDest []Variant, pPsa []PoolStringArray)
	VariantNewPoolVector2Array                *func(rDest []Variant, pPv2a []PoolVector2Array)
	VariantNewPoolVector3Array                *func(rDest []Variant, pPv3a []PoolVector3Array)
	VariantNewPoolColorArray                  *func(rDest []Variant, pPca []PoolColorArray)
	VariantAsBool                             Bool
	VariantAsUint                             Uint64T
	VariantAsInt                              Int64T
	VariantAsReal                             *func(pSelf []Variant) float64
	VariantAsString                           String
	VariantAsVector2                          Vector2
	VariantAsRect2                            Rect2
	VariantAsVector3                          Vector3
	VariantAsTransform2d                      Transform2d
	VariantAsPlane                            Plane
	VariantAsQuat                             Quat
	VariantAsAabb                             Aabb
	VariantAsBasis                            Basis
	VariantAsTransform                        Transform
	VariantAsColor                            Color
	VariantAsNodePath                         NodePath
	VariantAsRid                              Rid
	VariantAsObject                           Object
	VariantAsDictionary                       Dictionary
	VariantAsArray                            Array
	VariantAsPoolByteArray                    PoolByteArray
	VariantAsPoolIntArray                     PoolIntArray
	VariantAsPoolRealArray                    PoolRealArray
	VariantAsPoolStringArray                  PoolStringArray
	VariantAsPoolVector2Array                 PoolVector2Array
	VariantAsPoolVector3Array                 PoolVector3Array
	VariantAsPoolColorArray                   PoolColorArray
	VariantCall                               Variant
	VariantHasMethod                          Bool
	VariantOperatorEqual                      Bool
	VariantOperatorLess                       Bool
	VariantHashCompare                        Bool
	VariantBooleanize                         Bool
	VariantDestroy                            *func(pSelf []Variant)
	StringNew                                 *func(rDest []String)
	StringNewCopy                             *func(rDest []String, pSrc []String)
	StringNewData                             *func(rDest []String, pContents string, pSize int32)
	StringNewUnicodeData                      *func(rDest []String, pContents []int32, pSize int32)
	StringGetData                             *func(pSelf []String, pDest []byte, pSize []int32)
	StringOperatorIndex                       WcharT
	StringOperatorIndexConst                  WcharT
	StringUnicodeStr                          WcharT
	StringOperatorEqual                       Bool
	StringOperatorLess                        Bool
	StringOperatorPlus                        String
	StringLength                              Int
	StringBeginsWith                          Bool
	StringBeginsWithCharArray                 Bool
	StringBigrams                             Array
	StringChr                                 String
	StringEndsWith                            Bool
	StringFind                                Int
	StringFindFrom                            Int
	StringFindmk                              Int
	StringFindmkFrom                          Int
	StringFindmkFromInPlace                   Int
	StringFindn                               Int
	StringFindnFrom                           Int
	StringFindLast                            Int
	StringFormat                              String
	StringFormatWithCustomPlaceholder         String
	StringHexEncodeBuffer                     String
	StringHexToInt                            Int
	StringHexToIntWithoutPrefix               Int
	StringInsert                              String
	StringIsNumeric                           Bool
	StringIsSubsequenceOf                     Bool
	StringIsSubsequenceOfi                    Bool
	StringLpad                                String
	StringLpadWithCustomCharacter             String
	StringMatch                               Bool
	StringMatchn                              Bool
	StringMd5                                 String
	StringNum                                 String
	StringNumInt64                            String
	StringNumInt64Capitalized                 String
	StringNumReal                             String
	StringNumScientific                       String
	StringNumWithDecimals                     String
	StringPadDecimals                         String
	StringPadZeros                            String
	StringReplaceFirst                        String
	StringReplace                             String
	StringReplacen                            String
	StringRfind                               Int
	StringRfindn                              Int
	StringRfindFrom                           Int
	StringRfindnFrom                          Int
	StringRpad                                String
	StringRpadWithCustomCharacter             String
	StringSimilarity                          Real
	StringSprintf                             String
	StringSubstr                              String
	StringToDouble                            *func(pSelf []String) float64
	StringToFloat                             Real
	StringToInt                               Int
	StringCamelcaseToUnderscore               String
	StringCamelcaseToUnderscoreLowercased     String
	StringCapitalize                          String
	StringCharToDouble                        *func(pWhat string) float64
	StringCharToInt                           Int
	StringWcharToInt                          Int64T
	StringCharToIntWithLen                    Int
	StringCharToInt64WithLen                  Int64T
	StringHexToInt64                          Int64T
	StringHexToInt64WithPrefix                Int64T
	StringToInt64                             Int64T
	StringUnicodeCharToDouble                 *func(pStr []int32, rEnd [][]int32) float64
	StringGetSliceCount                       Int
	StringGetSlice                            String
	StringGetSlicec                           String
	StringSplit                               Array
	StringSplitAllowEmpty                     Array
	StringSplitFloats                         Array
	StringSplitFloatsAllowsEmpty              Array
	StringSplitFloatsMk                       Array
	StringSplitFloatsMkAllowsEmpty            Array
	StringSplitInts                           Array
	StringSplitIntsAllowsEmpty                Array
	StringSplitIntsMk                         Array
	StringSplitIntsMkAllowsEmpty              Array
	StringSplitSpaces                         Array
	StringCharLowercase                       WcharT
	StringCharUppercase                       WcharT
	StringToLower                             String
	StringToUpper                             String
	StringGetBasename                         String
	StringGetExtension                        String
	StringLeft                                String
	StringOrdAt                               WcharT
	StringPlusFile                            String
	StringRight                               String
	StringStripEdges                          String
	StringStripEscapes                        String
	StringErase                               *func(pSelf []String, pPos Int, pChars Int)
	StringAscii                               *func(pSelf []String, result []byte)
	StringAsciiExtended                       *func(pSelf []String, result []byte)
	StringUtf8                                *func(pSelf []String, result []byte)
	StringParseUtf8                           Bool
	StringParseUtf8WithLen                    Bool
	StringCharsToUtf8                         String
	StringCharsToUtf8WithLen                  String
	StringHash                                Uint32T
	StringHash64                              Uint64T
	StringHashChars                           Uint32T
	StringHashCharsWithLen                    Uint32T
	StringHashUtf8Chars                       Uint32T
	StringHashUtf8CharsWithLen                Uint32T
	StringMd5Buffer                           PoolByteArray
	StringMd5Text                             String
	StringSha256Buffer                        PoolByteArray
	StringSha256Text                          String
	StringEmpty                               Bool
	StringGetBaseDir                          String
	StringGetFile                             String
	StringHumanizeSize                        String
	StringIsAbsPath                           Bool
	StringIsRelPath                           Bool
	StringIsResourceFile                      Bool
	StringPathTo                              String
	StringPathToFile                          String
	StringSimplifyPath                        String
	StringCEscape                             String
	StringCEscapeMultiline                    String
	StringCUnescape                           String
	StringHttpEscape                          String
	StringHttpUnescape                        String
	StringJsonEscape                          String
	StringWordWrap                            String
	StringXmlEscape                           String
	StringXmlEscapeWithQuotes                 String
	StringXmlUnescape                         String
	StringPercentDecode                       String
	StringPercentEncode                       String
	StringIsValidFloat                        Bool
	StringIsValidHexNumber                    Bool
	StringIsValidHtmlColor                    Bool
	StringIsValidIdentifier                   Bool
	StringIsValidInteger                      Bool
	StringIsValidIpAddress                    Bool
	StringDestroy                             *func(pSelf []String)
	StringNameNew                             *func(rDest []StringName, pName []String)
	StringNameNewData                         *func(rDest []StringName, pName string)
	StringNameGetName                         String
	StringNameGetHash                         Uint32T
	StringNameGetDataUniquePointer            *func(pSelf []StringName) unsafe.Pointer
	StringNameOperatorEqual                   Bool
	StringNameOperatorLess                    Bool
	StringNameDestroy                         *func(pSelf []StringName)
	ObjectDestroy                             *func(pO *Object)
	GlobalGetSingleton                        Object
	MethodBindGetMethod                       MethodBind
	MethodBindPtrcall                         *func(pMethodBind []MethodBind, pInstance *Object, pArgs []unsafe.Pointer, pRet unsafe.Pointer)
	MethodBindCall                            Variant
	GetClassConstructor                       ClassConstructor
	RegisterNativeCallType                    *func(callType string, pCallback NativeCallCb)
	Alloc                                     *func(pBytes int32) unsafe.Pointer
	Realloc                                   *func(pPtr unsafe.Pointer, pBytes int32) unsafe.Pointer
	Free                                      *func(pPtr unsafe.Pointer)
	PrintError                                *func(pDescription string, pFunction string, pFile string, pLine int32)
	PrintWarning                              *func(pDescription string, pFunction string, pFile string, pLine int32)
	Print                                     *func(pMessage []String)
	ref57717e51                               *C.godot_gdnative_core_api_struct
	allocs57717e51                            interface{}
}

// Bool type as declared in gdnative/gdnative.h:124
type Bool bool

// Int type as declared in gdnative/gdnative.h:131
type Int int32

// Real type as declared in gdnative/gdnative.h:135
type Real float32

// Object type as declared in gdnative/gdnative.h:138
type Object [0]byte

// MethodBind as declared in gdnative/gdnative.h:225
type MethodBind struct {
	DontTouchThat  [1]byte
	ref3a05c0bc    *C.godot_method_bind
	allocs3a05c0bc interface{}
}

// GdnativeApiVersion as declared in gdnative/gdnative.h:235
type GdnativeApiVersion struct {
	Major          uint32
	Minor          uint32
	ref5eed2c27    *C.godot_gdnative_api_version
	allocs5eed2c27 interface{}
}

// GdnativeApiStruct as declared in gdnative/gdnative.h:237
type GdnativeApiStruct struct {
	Type           uint32
	Version        GdnativeApiVersion
	Next           []GdnativeApiStruct
	ref45f52b65    *C.godot_gdnative_api_struct
	allocs45f52b65 interface{}
}

// GdnativeInitOptions as declared in gdnative/gdnative.h:257
type GdnativeInitOptions struct {
	InEditor              Bool
	CoreApiHash           uint64
	EditorApiHash         uint64
	NoApiHash             uint64
	ReportVersionMismatch *func(pLibrary *Object, pWhat string, pWant GdnativeApiVersion, pHave GdnativeApiVersion)
	ReportLoadingError    *func(pLibrary *Object, pWhat string)
	GdNativeLibrary       *Object
	ApiStruct             []GdnativeCoreApiStruct
	ActiveLibraryPath     []String
	reff9d34929           *C.godot_gdnative_init_options
	allocsf9d34929        interface{}
}

// GdnativeTerminateOptions as declared in gdnative/gdnative.h:261
type GdnativeTerminateOptions struct {
	InEditor       Bool
	ref80c63fdc    *C.godot_gdnative_terminate_options
	allocs80c63fdc interface{}
}

// ClassConstructor type as declared in gdnative/gdnative.h:264
type ClassConstructor func() *Object

// GdnativeInitFn type as declared in gdnative/gdnative.h:271
type GdnativeInitFn func(arg0 []GdnativeInitOptions)

// GdnativeTerminateFn type as declared in gdnative/gdnative.h:272
type GdnativeTerminateFn func(arg0 []GdnativeTerminateOptions)

// GdnativeProcedureFn type as declared in gdnative/gdnative.h:273
type GdnativeProcedureFn func(arg0 []Array) Variant

// String as declared in gdnative/string.h:46
type String struct {
	DontTouchThat  [8]byte
	ref6d1ede57    *C.godot_string
	allocs6d1ede57 interface{}
}

// Array as declared in gdnative/array.h:46
type Array struct {
	DontTouchThat  [8]byte
	refb81158a5    *C.godot_array
	allocsb81158a5 interface{}
}

// PoolArrayReadAccess as declared in gdnative/pool_arrays.h:45
type PoolArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolByteArrayReadAccess as declared in gdnative/pool_arrays.h:47
type PoolByteArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolIntArrayReadAccess as declared in gdnative/pool_arrays.h:48
type PoolIntArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolRealArrayReadAccess as declared in gdnative/pool_arrays.h:49
type PoolRealArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolStringArrayReadAccess as declared in gdnative/pool_arrays.h:50
type PoolStringArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolVector2ArrayReadAccess as declared in gdnative/pool_arrays.h:51
type PoolVector2ArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolVector3ArrayReadAccess as declared in gdnative/pool_arrays.h:52
type PoolVector3ArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolColorArrayReadAccess as declared in gdnative/pool_arrays.h:53
type PoolColorArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// PoolArrayWriteAccess as declared in gdnative/pool_arrays.h:61
type PoolArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolByteArrayWriteAccess as declared in gdnative/pool_arrays.h:63
type PoolByteArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolIntArrayWriteAccess as declared in gdnative/pool_arrays.h:64
type PoolIntArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolRealArrayWriteAccess as declared in gdnative/pool_arrays.h:65
type PoolRealArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolStringArrayWriteAccess as declared in gdnative/pool_arrays.h:66
type PoolStringArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolVector2ArrayWriteAccess as declared in gdnative/pool_arrays.h:67
type PoolVector2ArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolVector3ArrayWriteAccess as declared in gdnative/pool_arrays.h:68
type PoolVector3ArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolColorArrayWriteAccess as declared in gdnative/pool_arrays.h:69
type PoolColorArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// PoolByteArray as declared in gdnative/pool_arrays.h:79
type PoolByteArray struct {
	DontTouchThat [8]byte
	refbf60ed2    *C.godot_pool_byte_array
	allocsbf60ed2 interface{}
}

// PoolIntArray as declared in gdnative/pool_arrays.h:90
type PoolIntArray struct {
	DontTouchThat  [8]byte
	ref5d4f26e6    *C.godot_pool_int_array
	allocs5d4f26e6 interface{}
}

// PoolRealArray as declared in gdnative/pool_arrays.h:101
type PoolRealArray struct {
	DontTouchThat  [8]byte
	refc76f44c3    *C.godot_pool_real_array
	allocsc76f44c3 interface{}
}

// PoolStringArray as declared in gdnative/pool_arrays.h:112
type PoolStringArray struct {
	DontTouchThat [8]byte
	reff6fe5d9    *C.godot_pool_string_array
	allocsf6fe5d9 interface{}
}

// PoolVector2Array as declared in gdnative/pool_arrays.h:123
type PoolVector2Array struct {
	DontTouchThat  [8]byte
	ref7f6b2885    *C.godot_pool_vector2_array
	allocs7f6b2885 interface{}
}

// PoolVector3Array as declared in gdnative/pool_arrays.h:134
type PoolVector3Array struct {
	DontTouchThat  [8]byte
	refd91c2331    *C.godot_pool_vector3_array
	allocsd91c2331 interface{}
}

// PoolColorArray as declared in gdnative/pool_arrays.h:145
type PoolColorArray struct {
	DontTouchThat  [8]byte
	refd5cae78e    *C.godot_pool_color_array
	allocsd5cae78e interface{}
}

// Color as declared in gdnative/color.h:45
type Color struct {
	DontTouchThat  [16]byte
	ref7f4bfefb    *C.godot_color
	allocs7f4bfefb interface{}
}

// Vector2 as declared in gdnative/vector2.h:45
type Vector2 struct {
	DontTouchThat  [8]byte
	refbc81274e    *C.godot_vector2
	allocsbc81274e interface{}
}

// Vector3 as declared in gdnative/vector3.h:45
type Vector3 struct {
	DontTouchThat  [12]byte
	refcb8617d8    *C.godot_vector3
	allocscb8617d8 interface{}
}

// Basis as declared in gdnative/basis.h:45
type Basis struct {
	DontTouchThat  [36]byte
	ref94d3d325    *C.godot_basis
	allocs94d3d325 interface{}
}

// Quat as declared in gdnative/quat.h:45
type Quat struct {
	DontTouchThat  [16]byte
	reffaf33e0b    *C.godot_quat
	allocsfaf33e0b interface{}
}

// Variant as declared in gdnative/variant.h:45
type Variant struct {
	DontTouchThat  [24]byte
	refabb5c0da    *C.godot_variant
	allocsabb5c0da interface{}
}

// VariantCallError as declared in gdnative/variant.h:100
type VariantCallError struct {
	Error          VariantCallErrorError
	Argument       int32
	Expected       VariantType
	ref3ce71027    *C.godot_variant_call_error
	allocs3ce71027 interface{}
}

// Aabb as declared in gdnative/aabb.h:45
type Aabb struct {
	DontTouchThat  [24]byte
	ref6e3c84aa    *C.godot_aabb
	allocs6e3c84aa interface{}
}

// Plane as declared in gdnative/plane.h:45
type Plane struct {
	DontTouchThat  [16]byte
	refd8ae9b92    *C.godot_plane
	allocsd8ae9b92 interface{}
}

// Dictionary as declared in gdnative/dictionary.h:45
type Dictionary struct {
	DontTouchThat  [8]byte
	refb267a9b9    *C.godot_dictionary
	allocsb267a9b9 interface{}
}

// NodePath as declared in gdnative/node_path.h:45
type NodePath struct {
	DontTouchThat  [8]byte
	ref6c34dff3    *C.godot_node_path
	allocs6c34dff3 interface{}
}

// Rect2 as declared in gdnative/rect2.h:43
type Rect2 struct {
	DontTouchThat  [16]byte
	ref99c06d9a    *C.godot_rect2
	allocs99c06d9a interface{}
}

// Rid as declared in gdnative/rid.h:45
type Rid struct {
	DontTouchThat  [8]byte
	ref67320fc7    *C.godot_rid
	allocs67320fc7 interface{}
}

// Transform as declared in gdnative/transform.h:45
type Transform struct {
	DontTouchThat  [48]byte
	refd77658c7    *C.godot_transform
	allocsd77658c7 interface{}
}

// Transform2d as declared in gdnative/transform2d.h:45
type Transform2d struct {
	DontTouchThat [24]byte
	ref77dacf6    *C.godot_transform2d
	allocs77dacf6 interface{}
}

// StringName as declared in gdnative/string_name.h:46
type StringName struct {
	DontTouchThat  [8]byte
	ref895548fc    *C.godot_string_name
	allocs895548fc interface{}
}

// ArvrInterfaceGdnative as declared in arvr/godot_arvr.h:55
type ArvrInterfaceGdnative struct {
	Constructor                 *func(arg0 *Object) unsafe.Pointer
	Destructor                  *func(arg0 unsafe.Pointer)
	GetName                     String
	GetCapabilities             Int
	GetAnchorDetectionIsEnabled Bool
	SetAnchorDetectionIsEnabled *func(arg0 unsafe.Pointer, arg1 Bool)
	IsStereo                    Bool
	IsInitialized               Bool
	Initialize                  Bool
	Uninitialize                *func(arg0 unsafe.Pointer)
	GetRenderTargetsize         Vector2
	GetTransformForEye          Transform
	FillProjectionForEye        *func(arg0 unsafe.Pointer, arg1 []Real, arg2 Int, arg3 Real, arg4 Real, arg5 Real)
	CommitForEye                *func(arg0 unsafe.Pointer, arg1 Int, arg2 []Rid, arg3 []Rect2)
	Process                     *func(arg0 unsafe.Pointer)
	reff96a0b88                 *C.godot_arvr_interface_gdnative
	allocsf96a0b88              interface{}
}

// PropertyAttributes as declared in nativescript/godot_nativescript.h:117
type PropertyAttributes struct {
	RsetType       MethodRpcMode
	Type           Int
	Hint           PropertyHint
	HintString     String
	Usage          PropertyUsageFlags
	DefaultValue   Variant
	ref431c473b    *C.godot_property_attributes
	allocs431c473b interface{}
}

// InstanceCreateFunc as declared in nativescript/godot_nativescript.h:124
type InstanceCreateFunc struct {
	CreateFunc     *func(arg0 *Object, arg1 unsafe.Pointer) unsafe.Pointer
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	ref70ecb5db    *C.godot_instance_create_func
	allocs70ecb5db interface{}
}

// InstanceDestroyFunc as declared in nativescript/godot_nativescript.h:131
type InstanceDestroyFunc struct {
	DestroyFunc    *func(arg0 *Object, arg1 unsafe.Pointer, arg2 unsafe.Pointer)
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	refd0d05668    *C.godot_instance_destroy_func
	allocsd0d05668 interface{}
}

// MethodAttributes as declared in nativescript/godot_nativescript.h:139
type MethodAttributes struct {
	RpcType        MethodRpcMode
	ref66a6c5c9    *C.godot_method_attributes
	allocs66a6c5c9 interface{}
}

// InstanceMethod as declared in nativescript/godot_nativescript.h:146
type InstanceMethod struct {
	Method         Variant
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	ref10e1583e    *C.godot_instance_method
	allocs10e1583e interface{}
}

// PropertySetFunc as declared in nativescript/godot_nativescript.h:155
type PropertySetFunc struct {
	SetFunc       *func(arg0 *Object, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 []Variant)
	MethodData    unsafe.Pointer
	FreeFunc      *func(arg0 unsafe.Pointer)
	refc9844af    *C.godot_property_set_func
	allocsc9844af interface{}
}

// PropertyGetFunc as declared in nativescript/godot_nativescript.h:162
type PropertyGetFunc struct {
	GetFunc        Variant
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	reff4697b7e    *C.godot_property_get_func
	allocsf4697b7e interface{}
}

// SignalArgument as declared in nativescript/godot_nativescript.h:173
type SignalArgument struct {
	Name           String
	Type           Int
	Hint           PropertyHint
	HintString     String
	Usage          PropertyUsageFlags
	DefaultValue   Variant
	refc21e72ac    *C.godot_signal_argument
	allocsc21e72ac interface{}
}

// Signal as declared in nativescript/godot_nativescript.h:181
type Signal struct {
	Name           String
	NumArgs        int32
	Args           []SignalArgument
	NumDefaultArgs int32
	DefaultArgs    []Variant
	ref87acf90b    *C.godot_signal
	allocs87acf90b interface{}
}

// PluginscriptInstanceData type as declared in pluginscript/godot_pluginscript.h:40
type PluginscriptInstanceData [0]byte

// PluginscriptScriptData type as declared in pluginscript/godot_pluginscript.h:41
type PluginscriptScriptData [0]byte

// PluginscriptLanguageData type as declared in pluginscript/godot_pluginscript.h:42
type PluginscriptLanguageData [0]byte

// PluginscriptInstanceDesc as declared in pluginscript/godot_pluginscript.h:69
type PluginscriptInstanceDesc struct {
	Init                PluginscriptInstanceData
	Finish              *func(pData *PluginscriptInstanceData)
	SetProp             Bool
	GetProp             Bool
	CallMethod          Variant
	Notification        *func(pData *PluginscriptInstanceData, pNotification int32)
	GetRpcMode          MethodRpcMode
	GetRsetMode         MethodRpcMode
	RefcountIncremented *func(pData *PluginscriptInstanceData)
	RefcountDecremented *func(pData *PluginscriptInstanceData) bool
	refc0c19139         *C.godot_pluginscript_instance_desc
	allocsc0c19139      interface{}
}

// PluginscriptScriptManifest as declared in pluginscript/godot_pluginscript.h:104
type PluginscriptScriptManifest struct {
	Data           *PluginscriptScriptData
	Name           StringName
	IsTool         Bool
	Base           StringName
	MemberLines    Dictionary
	Methods        Array
	Signals        Array
	Properties     Array
	reffbf02dfd    *C.godot_pluginscript_script_manifest
	allocsfbf02dfd interface{}
}

// PluginscriptScriptDesc as declared in pluginscript/godot_pluginscript.h:110
type PluginscriptScriptDesc struct {
	Init           PluginscriptScriptManifest
	Finish         *func(pData *PluginscriptScriptData)
	InstanceDesc   PluginscriptInstanceDesc
	ref1aab3210    *C.godot_pluginscript_script_desc
	allocs1aab3210 interface{}
}

// PluginscriptProfilingData as declared in pluginscript/godot_pluginscript.h:119
type PluginscriptProfilingData struct {
	Signature      StringName
	CallCount      Int
	TotalTime      Int
	SelfTime       Int
	ref9c004e5a    *C.godot_pluginscript_profiling_data
	allocs9c004e5a interface{}
}

// PluginscriptLanguageDesc as declared in pluginscript/godot_pluginscript.h:163
type PluginscriptLanguageDesc struct {
	Name                           string
	Type                           string
	Extension                      string
	RecognizedExtensions           []string
	Init                           PluginscriptLanguageData
	Finish                         *func(pData *PluginscriptLanguageData)
	ReservedWords                  []string
	CommentDelimiters              []string
	StringDelimiters               []string
	HasNamedClasses                Bool
	SupportsBuiltinMode            Bool
	GetTemplateSourceCode          String
	Validate                       Bool
	FindFunction                   *func(pData *PluginscriptLanguageData, pFunction []String, pCode []String) int32
	MakeFunction                   String
	CompleteCode                   Error
	AutoIndentCode                 *func(pData *PluginscriptLanguageData, pCode []String, pFromLine int32, pToLine int32)
	AddGlobalConstant              *func(pData *PluginscriptLanguageData, pVariable []String, pValue []Variant)
	DebugGetError                  String
	DebugGetStackLevelCount        *func(pData *PluginscriptLanguageData) int32
	DebugGetStackLevelLine         *func(pData *PluginscriptLanguageData, pLevel int32) int32
	DebugGetStackLevelFunction     String
	DebugGetStackLevelSource       String
	DebugGetStackLevelLocals       *func(pData *PluginscriptLanguageData, pLevel int32, pLocals []PoolStringArray, pValues []Array, pMaxSubitems int32, pMaxDepth int32)
	DebugGetStackLevelMembers      *func(pData *PluginscriptLanguageData, pLevel int32, pMembers []PoolStringArray, pValues []Array, pMaxSubitems int32, pMaxDepth int32)
	DebugGetGlobals                *func(pData *PluginscriptLanguageData, pLocals []PoolStringArray, pValues []Array, pMaxSubitems int32, pMaxDepth int32)
	DebugParseStackLevelExpression String
	GetPublicFunctions             *func(pData *PluginscriptLanguageData, rFunctions []Array)
	GetPublicConstants             *func(pData *PluginscriptLanguageData, rConstants []Dictionary)
	ProfilingStart                 *func(pData *PluginscriptLanguageData)
	ProfilingStop                  *func(pData *PluginscriptLanguageData)
	ProfilingGetAccumulatedData    *func(pData *PluginscriptLanguageData, rInfo []PluginscriptProfilingData, pInfoMax int32) int32
	ProfilingGetFrameData          *func(pData *PluginscriptLanguageData, rInfo []PluginscriptProfilingData, pInfoMax int32) int32
	ProfilingFrame                 *func(pData *PluginscriptLanguageData)
	ScriptDesc                     PluginscriptScriptDesc
	refdac22bbe                    *C.godot_pluginscript_language_desc
	allocsdac22bbe                 interface{}
}
