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

// GodotGdnativeExtNativescriptApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:52
type GodotGdnativeExtNativescriptApiStruct struct {
	Type                               uint32
	Version                            GodotGdnativeApiVersion
	Next                               []GodotGdnativeApiStruct
	GodotNativescriptRegisterClass     *func(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc GodotInstanceCreateFunc, pDestroyFunc GodotInstanceDestroyFunc)
	GodotNativescriptRegisterToolClass *func(pGdnativeHandle unsafe.Pointer, pName string, pBase string, pCreateFunc GodotInstanceCreateFunc, pDestroyFunc GodotInstanceDestroyFunc)
	GodotNativescriptRegisterMethod    *func(pGdnativeHandle unsafe.Pointer, pName string, pFunctionName string, pAttr GodotMethodAttributes, pMethod GodotInstanceMethod)
	GodotNativescriptRegisterProperty  *func(pGdnativeHandle unsafe.Pointer, pName string, pPath string, pAttr []GodotPropertyAttributes, pSetFunc GodotPropertySetFunc, pGetFunc GodotPropertyGetFunc)
	GodotNativescriptRegisterSignal    *func(pGdnativeHandle unsafe.Pointer, pName string, pSignal []GodotSignal)
	GodotNativescriptGetUserdata       *func(pInstance *GodotObject) unsafe.Pointer
	reff0cd6324                        *C.godot_gdnative_ext_nativescript_api_struct
	allocsf0cd6324                     interface{}
}

// GodotGdnativeExtPluginscriptApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:59
type GodotGdnativeExtPluginscriptApiStruct struct {
	Type                              uint32
	Version                           GodotGdnativeApiVersion
	Next                              []GodotGdnativeApiStruct
	GodotPluginscriptRegisterLanguage *func(languageDesc []GodotPluginscriptLanguageDesc)
	refc52e13a7                       *C.godot_gdnative_ext_pluginscript_api_struct
	allocsc52e13a7                    interface{}
}

// GodotGdnativeExtArvrApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:76
type GodotGdnativeExtArvrApiStruct struct {
	Type                            uint32
	Version                         GodotGdnativeApiVersion
	Next                            []GodotGdnativeApiStruct
	GodotArvrRegisterInterface      *func(pInterface []GodotArvrInterfaceGdnative)
	GodotArvrGetWorldscale          GodotReal
	GodotArvrGetReferenceFrame      GodotTransform
	GodotArvrBlit                   *func(pEye int32, pRenderTarget []GodotRid, pScreenRect []GodotRect2)
	GodotArvrGetTexid               GodotInt
	GodotArvrAddController          GodotInt
	GodotArvrRemoveController       *func(pControllerId GodotInt)
	GodotArvrSetControllerTransform *func(pControllerId GodotInt, pTransform []GodotTransform, pTracksOrientation GodotBool, pTracksPosition GodotBool)
	GodotArvrSetControllerButton    *func(pControllerId GodotInt, pButton GodotInt, pIsPressed GodotBool)
	GodotArvrSetControllerAxis      *func(pControllerId GodotInt, pExis GodotInt, pValue GodotReal, pCanBeNegative GodotBool)
	GodotArvrGetControllerRumble    GodotReal
	ref64de5bf5                     *C.godot_gdnative_ext_arvr_api_struct
	allocs64de5bf5                  interface{}
}

// GodotGdnativeCoreApiStruct as declared in godot_headers/gdnative_api_struct.gen.h:809
type GodotGdnativeCoreApiStruct struct {
	Type                                           uint32
	Version                                        GodotGdnativeApiVersion
	Next                                           []GodotGdnativeApiStruct
	NumExtensions                                  uint32
	Extensions                                     [][]GodotGdnativeApiStruct
	GodotColorNewRgba                              *func(rDest []GodotColor, pR GodotReal, pG GodotReal, pB GodotReal, pA GodotReal)
	GodotColorNewRgb                               *func(rDest []GodotColor, pR GodotReal, pG GodotReal, pB GodotReal)
	GodotColorGetR                                 GodotReal
	GodotColorSetR                                 *func(pSelf []GodotColor, r GodotReal)
	GodotColorGetG                                 GodotReal
	GodotColorSetG                                 *func(pSelf []GodotColor, g GodotReal)
	GodotColorGetB                                 GodotReal
	GodotColorSetB                                 *func(pSelf []GodotColor, b GodotReal)
	GodotColorGetA                                 GodotReal
	GodotColorSetA                                 *func(pSelf []GodotColor, a GodotReal)
	GodotColorGetH                                 GodotReal
	GodotColorGetS                                 GodotReal
	GodotColorGetV                                 GodotReal
	GodotColorAsString                             GodotString
	GodotColorToRgba32                             GodotInt
	GodotColorToArgb32                             GodotInt
	GodotColorGray                                 GodotReal
	GodotColorInverted                             GodotColor
	GodotColorContrasted                           GodotColor
	GodotColorLinearInterpolate                    GodotColor
	GodotColorBlend                                GodotColor
	GodotColorToHtml                               GodotString
	GodotColorOperatorEqual                        GodotBool
	GodotColorOperatorLess                         GodotBool
	GodotVector2New                                *func(rDest []GodotVector2, pX GodotReal, pY GodotReal)
	GodotVector2AsString                           GodotString
	GodotVector2Normalized                         GodotVector2
	GodotVector2Length                             GodotReal
	GodotVector2Angle                              GodotReal
	GodotVector2LengthSquared                      GodotReal
	GodotVector2IsNormalized                       GodotBool
	GodotVector2DistanceTo                         GodotReal
	GodotVector2DistanceSquaredTo                  GodotReal
	GodotVector2AngleTo                            GodotReal
	GodotVector2AngleToPoint                       GodotReal
	GodotVector2LinearInterpolate                  GodotVector2
	GodotVector2CubicInterpolate                   GodotVector2
	GodotVector2Rotated                            GodotVector2
	GodotVector2Tangent                            GodotVector2
	GodotVector2Floor                              GodotVector2
	GodotVector2Snapped                            GodotVector2
	GodotVector2Aspect                             GodotReal
	GodotVector2Dot                                GodotReal
	GodotVector2Slide                              GodotVector2
	GodotVector2Bounce                             GodotVector2
	GodotVector2Reflect                            GodotVector2
	GodotVector2Abs                                GodotVector2
	GodotVector2Clamped                            GodotVector2
	GodotVector2OperatorAdd                        GodotVector2
	GodotVector2OperatorSubstract                  GodotVector2
	GodotVector2OperatorMultiplyVector             GodotVector2
	GodotVector2OperatorMultiplyScalar             GodotVector2
	GodotVector2OperatorDivideVector               GodotVector2
	GodotVector2OperatorDivideScalar               GodotVector2
	GodotVector2OperatorEqual                      GodotBool
	GodotVector2OperatorLess                       GodotBool
	GodotVector2OperatorNeg                        GodotVector2
	GodotVector2SetX                               *func(pSelf []GodotVector2, pX GodotReal)
	GodotVector2SetY                               *func(pSelf []GodotVector2, pY GodotReal)
	GodotVector2GetX                               GodotReal
	GodotVector2GetY                               GodotReal
	GodotQuatNew                                   *func(rDest []GodotQuat, pX GodotReal, pY GodotReal, pZ GodotReal, pW GodotReal)
	GodotQuatNewWithAxisAngle                      *func(rDest []GodotQuat, pAxis []GodotVector3, pAngle GodotReal)
	GodotQuatGetX                                  GodotReal
	GodotQuatSetX                                  *func(pSelf []GodotQuat, val GodotReal)
	GodotQuatGetY                                  GodotReal
	GodotQuatSetY                                  *func(pSelf []GodotQuat, val GodotReal)
	GodotQuatGetZ                                  GodotReal
	GodotQuatSetZ                                  *func(pSelf []GodotQuat, val GodotReal)
	GodotQuatGetW                                  GodotReal
	GodotQuatSetW                                  *func(pSelf []GodotQuat, val GodotReal)
	GodotQuatAsString                              GodotString
	GodotQuatLength                                GodotReal
	GodotQuatLengthSquared                         GodotReal
	GodotQuatNormalized                            GodotQuat
	GodotQuatIsNormalized                          GodotBool
	GodotQuatInverse                               GodotQuat
	GodotQuatDot                                   GodotReal
	GodotQuatXform                                 GodotVector3
	GodotQuatSlerp                                 GodotQuat
	GodotQuatSlerpni                               GodotQuat
	GodotQuatCubicSlerp                            GodotQuat
	GodotQuatOperatorMultiply                      GodotQuat
	GodotQuatOperatorAdd                           GodotQuat
	GodotQuatOperatorSubstract                     GodotQuat
	GodotQuatOperatorDivide                        GodotQuat
	GodotQuatOperatorEqual                         GodotBool
	GodotQuatOperatorNeg                           GodotQuat
	GodotBasisNewWithRows                          *func(rDest []GodotBasis, pXAxis []GodotVector3, pYAxis []GodotVector3, pZAxis []GodotVector3)
	GodotBasisNewWithAxisAndAngle                  *func(rDest []GodotBasis, pAxis []GodotVector3, pPhi GodotReal)
	GodotBasisNewWithEuler                         *func(rDest []GodotBasis, pEuler []GodotVector3)
	GodotBasisAsString                             GodotString
	GodotBasisInverse                              GodotBasis
	GodotBasisTransposed                           GodotBasis
	GodotBasisOrthonormalized                      GodotBasis
	GodotBasisDeterminant                          GodotReal
	GodotBasisRotated                              GodotBasis
	GodotBasisScaled                               GodotBasis
	GodotBasisGetScale                             GodotVector3
	GodotBasisGetEuler                             GodotVector3
	GodotBasisTdotx                                GodotReal
	GodotBasisTdoty                                GodotReal
	GodotBasisTdotz                                GodotReal
	GodotBasisXform                                GodotVector3
	GodotBasisXformInv                             GodotVector3
	GodotBasisGetOrthogonalIndex                   GodotInt
	GodotBasisNew                                  *func(rDest []GodotBasis)
	GodotBasisNewWithEulerQuat                     *func(rDest []GodotBasis, pEuler []GodotQuat)
	GodotBasisGetElements                          *func(pSelf []GodotBasis, pElements []GodotVector3)
	GodotBasisGetAxis                              GodotVector3
	GodotBasisSetAxis                              *func(pSelf []GodotBasis, pAxis GodotInt, pValue []GodotVector3)
	GodotBasisGetRow                               GodotVector3
	GodotBasisSetRow                               *func(pSelf []GodotBasis, pRow GodotInt, pValue []GodotVector3)
	GodotBasisOperatorEqual                        GodotBool
	GodotBasisOperatorAdd                          GodotBasis
	GodotBasisOperatorSubstract                    GodotBasis
	GodotBasisOperatorMultiplyVector               GodotBasis
	GodotBasisOperatorMultiplyScalar               GodotBasis
	GodotVector3New                                *func(rDest []GodotVector3, pX GodotReal, pY GodotReal, pZ GodotReal)
	GodotVector3AsString                           GodotString
	GodotVector3MinAxis                            GodotInt
	GodotVector3MaxAxis                            GodotInt
	GodotVector3Length                             GodotReal
	GodotVector3LengthSquared                      GodotReal
	GodotVector3IsNormalized                       GodotBool
	GodotVector3Normalized                         GodotVector3
	GodotVector3Inverse                            GodotVector3
	GodotVector3Snapped                            GodotVector3
	GodotVector3Rotated                            GodotVector3
	GodotVector3LinearInterpolate                  GodotVector3
	GodotVector3CubicInterpolate                   GodotVector3
	GodotVector3Dot                                GodotReal
	GodotVector3Cross                              GodotVector3
	GodotVector3Outer                              GodotBasis
	GodotVector3ToDiagonalMatrix                   GodotBasis
	GodotVector3Abs                                GodotVector3
	GodotVector3Floor                              GodotVector3
	GodotVector3Ceil                               GodotVector3
	GodotVector3DistanceTo                         GodotReal
	GodotVector3DistanceSquaredTo                  GodotReal
	GodotVector3AngleTo                            GodotReal
	GodotVector3Slide                              GodotVector3
	GodotVector3Bounce                             GodotVector3
	GodotVector3Reflect                            GodotVector3
	GodotVector3OperatorAdd                        GodotVector3
	GodotVector3OperatorSubstract                  GodotVector3
	GodotVector3OperatorMultiplyVector             GodotVector3
	GodotVector3OperatorMultiplyScalar             GodotVector3
	GodotVector3OperatorDivideVector               GodotVector3
	GodotVector3OperatorDivideScalar               GodotVector3
	GodotVector3OperatorEqual                      GodotBool
	GodotVector3OperatorLess                       GodotBool
	GodotVector3OperatorNeg                        GodotVector3
	GodotVector3SetAxis                            *func(pSelf []GodotVector3, pAxis GodotVector3Axis, pVal GodotReal)
	GodotVector3GetAxis                            GodotReal
	GodotPoolByteArrayNew                          *func(rDest []GodotPoolByteArray)
	GodotPoolByteArrayNewCopy                      *func(rDest []GodotPoolByteArray, pSrc []GodotPoolByteArray)
	GodotPoolByteArrayNewWithArray                 *func(rDest []GodotPoolByteArray, pA []GodotArray)
	GodotPoolByteArrayAppend                       *func(pSelf []GodotPoolByteArray, pData byte)
	GodotPoolByteArrayAppendArray                  *func(pSelf []GodotPoolByteArray, pArray []GodotPoolByteArray)
	GodotPoolByteArrayInsert                       GodotError
	GodotPoolByteArrayInvert                       *func(pSelf []GodotPoolByteArray)
	GodotPoolByteArrayPushBack                     *func(pSelf []GodotPoolByteArray, pData byte)
	GodotPoolByteArrayRemove                       *func(pSelf []GodotPoolByteArray, pIdx GodotInt)
	GodotPoolByteArrayResize                       *func(pSelf []GodotPoolByteArray, pSize GodotInt)
	GodotPoolByteArrayRead                         GodotPoolByteArrayReadAccess
	GodotPoolByteArrayWrite                        GodotPoolByteArrayWriteAccess
	GodotPoolByteArraySet                          *func(pSelf []GodotPoolByteArray, pIdx GodotInt, pData byte)
	GodotPoolByteArrayGet                          Uint8T
	GodotPoolByteArraySize                         GodotInt
	GodotPoolByteArrayDestroy                      *func(pSelf []GodotPoolByteArray)
	GodotPoolIntArrayNew                           *func(rDest []GodotPoolIntArray)
	GodotPoolIntArrayNewCopy                       *func(rDest []GodotPoolIntArray, pSrc []GodotPoolIntArray)
	GodotPoolIntArrayNewWithArray                  *func(rDest []GodotPoolIntArray, pA []GodotArray)
	GodotPoolIntArrayAppend                        *func(pSelf []GodotPoolIntArray, pData GodotInt)
	GodotPoolIntArrayAppendArray                   *func(pSelf []GodotPoolIntArray, pArray []GodotPoolIntArray)
	GodotPoolIntArrayInsert                        GodotError
	GodotPoolIntArrayInvert                        *func(pSelf []GodotPoolIntArray)
	GodotPoolIntArrayPushBack                      *func(pSelf []GodotPoolIntArray, pData GodotInt)
	GodotPoolIntArrayRemove                        *func(pSelf []GodotPoolIntArray, pIdx GodotInt)
	GodotPoolIntArrayResize                        *func(pSelf []GodotPoolIntArray, pSize GodotInt)
	GodotPoolIntArrayRead                          GodotPoolIntArrayReadAccess
	GodotPoolIntArrayWrite                         GodotPoolIntArrayWriteAccess
	GodotPoolIntArraySet                           *func(pSelf []GodotPoolIntArray, pIdx GodotInt, pData GodotInt)
	GodotPoolIntArrayGet                           GodotInt
	GodotPoolIntArraySize                          GodotInt
	GodotPoolIntArrayDestroy                       *func(pSelf []GodotPoolIntArray)
	GodotPoolRealArrayNew                          *func(rDest []GodotPoolRealArray)
	GodotPoolRealArrayNewCopy                      *func(rDest []GodotPoolRealArray, pSrc []GodotPoolRealArray)
	GodotPoolRealArrayNewWithArray                 *func(rDest []GodotPoolRealArray, pA []GodotArray)
	GodotPoolRealArrayAppend                       *func(pSelf []GodotPoolRealArray, pData GodotReal)
	GodotPoolRealArrayAppendArray                  *func(pSelf []GodotPoolRealArray, pArray []GodotPoolRealArray)
	GodotPoolRealArrayInsert                       GodotError
	GodotPoolRealArrayInvert                       *func(pSelf []GodotPoolRealArray)
	GodotPoolRealArrayPushBack                     *func(pSelf []GodotPoolRealArray, pData GodotReal)
	GodotPoolRealArrayRemove                       *func(pSelf []GodotPoolRealArray, pIdx GodotInt)
	GodotPoolRealArrayResize                       *func(pSelf []GodotPoolRealArray, pSize GodotInt)
	GodotPoolRealArrayRead                         GodotPoolRealArrayReadAccess
	GodotPoolRealArrayWrite                        GodotPoolRealArrayWriteAccess
	GodotPoolRealArraySet                          *func(pSelf []GodotPoolRealArray, pIdx GodotInt, pData GodotReal)
	GodotPoolRealArrayGet                          GodotReal
	GodotPoolRealArraySize                         GodotInt
	GodotPoolRealArrayDestroy                      *func(pSelf []GodotPoolRealArray)
	GodotPoolStringArrayNew                        *func(rDest []GodotPoolStringArray)
	GodotPoolStringArrayNewCopy                    *func(rDest []GodotPoolStringArray, pSrc []GodotPoolStringArray)
	GodotPoolStringArrayNewWithArray               *func(rDest []GodotPoolStringArray, pA []GodotArray)
	GodotPoolStringArrayAppend                     *func(pSelf []GodotPoolStringArray, pData []GodotString)
	GodotPoolStringArrayAppendArray                *func(pSelf []GodotPoolStringArray, pArray []GodotPoolStringArray)
	GodotPoolStringArrayInsert                     GodotError
	GodotPoolStringArrayInvert                     *func(pSelf []GodotPoolStringArray)
	GodotPoolStringArrayPushBack                   *func(pSelf []GodotPoolStringArray, pData []GodotString)
	GodotPoolStringArrayRemove                     *func(pSelf []GodotPoolStringArray, pIdx GodotInt)
	GodotPoolStringArrayResize                     *func(pSelf []GodotPoolStringArray, pSize GodotInt)
	GodotPoolStringArrayRead                       GodotPoolStringArrayReadAccess
	GodotPoolStringArrayWrite                      GodotPoolStringArrayWriteAccess
	GodotPoolStringArraySet                        *func(pSelf []GodotPoolStringArray, pIdx GodotInt, pData []GodotString)
	GodotPoolStringArrayGet                        GodotString
	GodotPoolStringArraySize                       GodotInt
	GodotPoolStringArrayDestroy                    *func(pSelf []GodotPoolStringArray)
	GodotPoolVector2ArrayNew                       *func(rDest []GodotPoolVector2Array)
	GodotPoolVector2ArrayNewCopy                   *func(rDest []GodotPoolVector2Array, pSrc []GodotPoolVector2Array)
	GodotPoolVector2ArrayNewWithArray              *func(rDest []GodotPoolVector2Array, pA []GodotArray)
	GodotPoolVector2ArrayAppend                    *func(pSelf []GodotPoolVector2Array, pData []GodotVector2)
	GodotPoolVector2ArrayAppendArray               *func(pSelf []GodotPoolVector2Array, pArray []GodotPoolVector2Array)
	GodotPoolVector2ArrayInsert                    GodotError
	GodotPoolVector2ArrayInvert                    *func(pSelf []GodotPoolVector2Array)
	GodotPoolVector2ArrayPushBack                  *func(pSelf []GodotPoolVector2Array, pData []GodotVector2)
	GodotPoolVector2ArrayRemove                    *func(pSelf []GodotPoolVector2Array, pIdx GodotInt)
	GodotPoolVector2ArrayResize                    *func(pSelf []GodotPoolVector2Array, pSize GodotInt)
	GodotPoolVector2ArrayRead                      GodotPoolVector2ArrayReadAccess
	GodotPoolVector2ArrayWrite                     GodotPoolVector2ArrayWriteAccess
	GodotPoolVector2ArraySet                       *func(pSelf []GodotPoolVector2Array, pIdx GodotInt, pData []GodotVector2)
	GodotPoolVector2ArrayGet                       GodotVector2
	GodotPoolVector2ArraySize                      GodotInt
	GodotPoolVector2ArrayDestroy                   *func(pSelf []GodotPoolVector2Array)
	GodotPoolVector3ArrayNew                       *func(rDest []GodotPoolVector3Array)
	GodotPoolVector3ArrayNewCopy                   *func(rDest []GodotPoolVector3Array, pSrc []GodotPoolVector3Array)
	GodotPoolVector3ArrayNewWithArray              *func(rDest []GodotPoolVector3Array, pA []GodotArray)
	GodotPoolVector3ArrayAppend                    *func(pSelf []GodotPoolVector3Array, pData []GodotVector3)
	GodotPoolVector3ArrayAppendArray               *func(pSelf []GodotPoolVector3Array, pArray []GodotPoolVector3Array)
	GodotPoolVector3ArrayInsert                    GodotError
	GodotPoolVector3ArrayInvert                    *func(pSelf []GodotPoolVector3Array)
	GodotPoolVector3ArrayPushBack                  *func(pSelf []GodotPoolVector3Array, pData []GodotVector3)
	GodotPoolVector3ArrayRemove                    *func(pSelf []GodotPoolVector3Array, pIdx GodotInt)
	GodotPoolVector3ArrayResize                    *func(pSelf []GodotPoolVector3Array, pSize GodotInt)
	GodotPoolVector3ArrayRead                      GodotPoolVector3ArrayReadAccess
	GodotPoolVector3ArrayWrite                     GodotPoolVector3ArrayWriteAccess
	GodotPoolVector3ArraySet                       *func(pSelf []GodotPoolVector3Array, pIdx GodotInt, pData []GodotVector3)
	GodotPoolVector3ArrayGet                       GodotVector3
	GodotPoolVector3ArraySize                      GodotInt
	GodotPoolVector3ArrayDestroy                   *func(pSelf []GodotPoolVector3Array)
	GodotPoolColorArrayNew                         *func(rDest []GodotPoolColorArray)
	GodotPoolColorArrayNewCopy                     *func(rDest []GodotPoolColorArray, pSrc []GodotPoolColorArray)
	GodotPoolColorArrayNewWithArray                *func(rDest []GodotPoolColorArray, pA []GodotArray)
	GodotPoolColorArrayAppend                      *func(pSelf []GodotPoolColorArray, pData []GodotColor)
	GodotPoolColorArrayAppendArray                 *func(pSelf []GodotPoolColorArray, pArray []GodotPoolColorArray)
	GodotPoolColorArrayInsert                      GodotError
	GodotPoolColorArrayInvert                      *func(pSelf []GodotPoolColorArray)
	GodotPoolColorArrayPushBack                    *func(pSelf []GodotPoolColorArray, pData []GodotColor)
	GodotPoolColorArrayRemove                      *func(pSelf []GodotPoolColorArray, pIdx GodotInt)
	GodotPoolColorArrayResize                      *func(pSelf []GodotPoolColorArray, pSize GodotInt)
	GodotPoolColorArrayRead                        GodotPoolColorArrayReadAccess
	GodotPoolColorArrayWrite                       GodotPoolColorArrayWriteAccess
	GodotPoolColorArraySet                         *func(pSelf []GodotPoolColorArray, pIdx GodotInt, pData []GodotColor)
	GodotPoolColorArrayGet                         GodotColor
	GodotPoolColorArraySize                        GodotInt
	GodotPoolColorArrayDestroy                     *func(pSelf []GodotPoolColorArray)
	GodotPoolByteArrayReadAccessPtr                Uint8T
	GodotPoolByteArrayReadAccessOperatorAssign     *func(pRead []GodotPoolByteArrayReadAccess, pOther []GodotPoolByteArrayReadAccess)
	GodotPoolByteArrayReadAccessDestroy            *func(pRead []GodotPoolByteArrayReadAccess)
	GodotPoolIntArrayReadAccessPtr                 GodotInt
	GodotPoolIntArrayReadAccessOperatorAssign      *func(pRead []GodotPoolIntArrayReadAccess, pOther []GodotPoolIntArrayReadAccess)
	GodotPoolIntArrayReadAccessDestroy             *func(pRead []GodotPoolIntArrayReadAccess)
	GodotPoolRealArrayReadAccessPtr                GodotReal
	GodotPoolRealArrayReadAccessOperatorAssign     *func(pRead []GodotPoolRealArrayReadAccess, pOther []GodotPoolRealArrayReadAccess)
	GodotPoolRealArrayReadAccessDestroy            *func(pRead []GodotPoolRealArrayReadAccess)
	GodotPoolStringArrayReadAccessPtr              GodotString
	GodotPoolStringArrayReadAccessOperatorAssign   *func(pRead []GodotPoolStringArrayReadAccess, pOther []GodotPoolStringArrayReadAccess)
	GodotPoolStringArrayReadAccessDestroy          *func(pRead []GodotPoolStringArrayReadAccess)
	GodotPoolVector2ArrayReadAccessPtr             GodotVector2
	GodotPoolVector2ArrayReadAccessOperatorAssign  *func(pRead []GodotPoolVector2ArrayReadAccess, pOther []GodotPoolVector2ArrayReadAccess)
	GodotPoolVector2ArrayReadAccessDestroy         *func(pRead []GodotPoolVector2ArrayReadAccess)
	GodotPoolVector3ArrayReadAccessPtr             GodotVector3
	GodotPoolVector3ArrayReadAccessOperatorAssign  *func(pRead []GodotPoolVector3ArrayReadAccess, pOther []GodotPoolVector3ArrayReadAccess)
	GodotPoolVector3ArrayReadAccessDestroy         *func(pRead []GodotPoolVector3ArrayReadAccess)
	GodotPoolColorArrayReadAccessPtr               GodotColor
	GodotPoolColorArrayReadAccessOperatorAssign    *func(pRead []GodotPoolColorArrayReadAccess, pOther []GodotPoolColorArrayReadAccess)
	GodotPoolColorArrayReadAccessDestroy           *func(pRead []GodotPoolColorArrayReadAccess)
	GodotPoolByteArrayWriteAccessPtr               Uint8T
	GodotPoolByteArrayWriteAccessOperatorAssign    *func(pWrite []GodotPoolByteArrayWriteAccess, pOther []GodotPoolByteArrayWriteAccess)
	GodotPoolByteArrayWriteAccessDestroy           *func(pWrite []GodotPoolByteArrayWriteAccess)
	GodotPoolIntArrayWriteAccessPtr                GodotInt
	GodotPoolIntArrayWriteAccessOperatorAssign     *func(pWrite []GodotPoolIntArrayWriteAccess, pOther []GodotPoolIntArrayWriteAccess)
	GodotPoolIntArrayWriteAccessDestroy            *func(pWrite []GodotPoolIntArrayWriteAccess)
	GodotPoolRealArrayWriteAccessPtr               GodotReal
	GodotPoolRealArrayWriteAccessOperatorAssign    *func(pWrite []GodotPoolRealArrayWriteAccess, pOther []GodotPoolRealArrayWriteAccess)
	GodotPoolRealArrayWriteAccessDestroy           *func(pWrite []GodotPoolRealArrayWriteAccess)
	GodotPoolStringArrayWriteAccessPtr             GodotString
	GodotPoolStringArrayWriteAccessOperatorAssign  *func(pWrite []GodotPoolStringArrayWriteAccess, pOther []GodotPoolStringArrayWriteAccess)
	GodotPoolStringArrayWriteAccessDestroy         *func(pWrite []GodotPoolStringArrayWriteAccess)
	GodotPoolVector2ArrayWriteAccessPtr            GodotVector2
	GodotPoolVector2ArrayWriteAccessOperatorAssign *func(pWrite []GodotPoolVector2ArrayWriteAccess, pOther []GodotPoolVector2ArrayWriteAccess)
	GodotPoolVector2ArrayWriteAccessDestroy        *func(pWrite []GodotPoolVector2ArrayWriteAccess)
	GodotPoolVector3ArrayWriteAccessPtr            GodotVector3
	GodotPoolVector3ArrayWriteAccessOperatorAssign *func(pWrite []GodotPoolVector3ArrayWriteAccess, pOther []GodotPoolVector3ArrayWriteAccess)
	GodotPoolVector3ArrayWriteAccessDestroy        *func(pWrite []GodotPoolVector3ArrayWriteAccess)
	GodotPoolColorArrayWriteAccessPtr              GodotColor
	GodotPoolColorArrayWriteAccessOperatorAssign   *func(pWrite []GodotPoolColorArrayWriteAccess, pOther []GodotPoolColorArrayWriteAccess)
	GodotPoolColorArrayWriteAccessDestroy          *func(pWrite []GodotPoolColorArrayWriteAccess)
	GodotArrayNew                                  *func(rDest []GodotArray)
	GodotArrayNewCopy                              *func(rDest []GodotArray, pSrc []GodotArray)
	GodotArrayNewPoolColorArray                    *func(rDest []GodotArray, pPca []GodotPoolColorArray)
	GodotArrayNewPoolVector3Array                  *func(rDest []GodotArray, pPv3a []GodotPoolVector3Array)
	GodotArrayNewPoolVector2Array                  *func(rDest []GodotArray, pPv2a []GodotPoolVector2Array)
	GodotArrayNewPoolStringArray                   *func(rDest []GodotArray, pPsa []GodotPoolStringArray)
	GodotArrayNewPoolRealArray                     *func(rDest []GodotArray, pPra []GodotPoolRealArray)
	GodotArrayNewPoolIntArray                      *func(rDest []GodotArray, pPia []GodotPoolIntArray)
	GodotArrayNewPoolByteArray                     *func(rDest []GodotArray, pPba []GodotPoolByteArray)
	GodotArraySet                                  *func(pSelf []GodotArray, pIdx GodotInt, pValue []GodotVariant)
	GodotArrayGet                                  GodotVariant
	GodotArrayOperatorIndex                        GodotVariant
	GodotArrayOperatorIndexConst                   GodotVariant
	GodotArrayAppend                               *func(pSelf []GodotArray, pValue []GodotVariant)
	GodotArrayClear                                *func(pSelf []GodotArray)
	GodotArrayCount                                GodotInt
	GodotArrayEmpty                                GodotBool
	GodotArrayErase                                *func(pSelf []GodotArray, pValue []GodotVariant)
	GodotArrayFront                                GodotVariant
	GodotArrayBack                                 GodotVariant
	GodotArrayFind                                 GodotInt
	GodotArrayFindLast                             GodotInt
	GodotArrayHas                                  GodotBool
	GodotArrayHash                                 GodotInt
	GodotArrayInsert                               *func(pSelf []GodotArray, pPos GodotInt, pValue []GodotVariant)
	GodotArrayInvert                               *func(pSelf []GodotArray)
	GodotArrayPopBack                              GodotVariant
	GodotArrayPopFront                             GodotVariant
	GodotArrayPushBack                             *func(pSelf []GodotArray, pValue []GodotVariant)
	GodotArrayPushFront                            *func(pSelf []GodotArray, pValue []GodotVariant)
	GodotArrayRemove                               *func(pSelf []GodotArray, pIdx GodotInt)
	GodotArrayResize                               *func(pSelf []GodotArray, pSize GodotInt)
	GodotArrayRfind                                GodotInt
	GodotArraySize                                 GodotInt
	GodotArraySort                                 *func(pSelf []GodotArray)
	GodotArraySortCustom                           *func(pSelf []GodotArray, pObj *GodotObject, pFunc []GodotString)
	GodotArrayBsearch                              GodotInt
	GodotArrayBsearchCustom                        GodotInt
	GodotArrayDestroy                              *func(pSelf []GodotArray)
	GodotDictionaryNew                             *func(rDest []GodotDictionary)
	GodotDictionaryNewCopy                         *func(rDest []GodotDictionary, pSrc []GodotDictionary)
	GodotDictionaryDestroy                         *func(pSelf []GodotDictionary)
	GodotDictionarySize                            GodotInt
	GodotDictionaryEmpty                           GodotBool
	GodotDictionaryClear                           *func(pSelf []GodotDictionary)
	GodotDictionaryHas                             GodotBool
	GodotDictionaryHasAll                          GodotBool
	GodotDictionaryErase                           *func(pSelf []GodotDictionary, pKey []GodotVariant)
	GodotDictionaryHash                            GodotInt
	GodotDictionaryKeys                            GodotArray
	GodotDictionaryValues                          GodotArray
	GodotDictionaryGet                             GodotVariant
	GodotDictionarySet                             *func(pSelf []GodotDictionary, pKey []GodotVariant, pValue []GodotVariant)
	GodotDictionaryOperatorIndex                   GodotVariant
	GodotDictionaryOperatorIndexConst              GodotVariant
	GodotDictionaryNext                            GodotVariant
	GodotDictionaryOperatorEqual                   GodotBool
	GodotDictionaryToJson                          GodotString
	GodotNodePathNew                               *func(rDest []GodotNodePath, pFrom []GodotString)
	GodotNodePathNewCopy                           *func(rDest []GodotNodePath, pSrc []GodotNodePath)
	GodotNodePathDestroy                           *func(pSelf []GodotNodePath)
	GodotNodePathAsString                          GodotString
	GodotNodePathIsAbsolute                        GodotBool
	GodotNodePathGetNameCount                      GodotInt
	GodotNodePathGetName                           GodotString
	GodotNodePathGetSubnameCount                   GodotInt
	GodotNodePathGetSubname                        GodotString
	GodotNodePathGetConcatenatedSubnames           GodotString
	GodotNodePathIsEmpty                           GodotBool
	GodotNodePathOperatorEqual                     GodotBool
	GodotPlaneNewWithReals                         *func(rDest []GodotPlane, pA GodotReal, pB GodotReal, pC GodotReal, pD GodotReal)
	GodotPlaneNewWithVectors                       *func(rDest []GodotPlane, pV1 []GodotVector3, pV2 []GodotVector3, pV3 []GodotVector3)
	GodotPlaneNewWithNormal                        *func(rDest []GodotPlane, pNormal []GodotVector3, pD GodotReal)
	GodotPlaneAsString                             GodotString
	GodotPlaneNormalized                           GodotPlane
	GodotPlaneCenter                               GodotVector3
	GodotPlaneGetAnyPoint                          GodotVector3
	GodotPlaneIsPointOver                          GodotBool
	GodotPlaneDistanceTo                           GodotReal
	GodotPlaneHasPoint                             GodotBool
	GodotPlaneProject                              GodotVector3
	GodotPlaneIntersect3                           GodotBool
	GodotPlaneIntersectsRay                        GodotBool
	GodotPlaneIntersectsSegment                    GodotBool
	GodotPlaneOperatorNeg                          GodotPlane
	GodotPlaneOperatorEqual                        GodotBool
	GodotPlaneSetNormal                            *func(pSelf []GodotPlane, pNormal []GodotVector3)
	GodotPlaneGetNormal                            GodotVector3
	GodotPlaneGetD                                 GodotReal
	GodotPlaneSetD                                 *func(pSelf []GodotPlane, pD GodotReal)
	GodotRect2NewWithPositionAndSize               *func(rDest []GodotRect2, pPos []GodotVector2, pSize []GodotVector2)
	GodotRect2New                                  *func(rDest []GodotRect2, pX GodotReal, pY GodotReal, pWidth GodotReal, pHeight GodotReal)
	GodotRect2AsString                             GodotString
	GodotRect2GetArea                              GodotReal
	GodotRect2Intersects                           GodotBool
	GodotRect2Encloses                             GodotBool
	GodotRect2HasNoArea                            GodotBool
	GodotRect2Clip                                 GodotRect2
	GodotRect2Merge                                GodotRect2
	GodotRect2HasPoint                             GodotBool
	GodotRect2Grow                                 GodotRect2
	GodotRect2Expand                               GodotRect2
	GodotRect2OperatorEqual                        GodotBool
	GodotRect2GetPosition                          GodotVector2
	GodotRect2GetSize                              GodotVector2
	GodotRect2SetPosition                          *func(pSelf []GodotRect2, pPos []GodotVector2)
	GodotRect2SetSize                              *func(pSelf []GodotRect2, pSize []GodotVector2)
	GodotAabbNew                                   *func(rDest []GodotAabb, pPos []GodotVector3, pSize []GodotVector3)
	GodotAabbGetPosition                           GodotVector3
	GodotAabbSetPosition                           *func(pSelf []GodotAabb, pV []GodotVector3)
	GodotAabbGetSize                               GodotVector3
	GodotAabbSetSize                               *func(pSelf []GodotAabb, pV []GodotVector3)
	GodotAabbAsString                              GodotString
	GodotAabbGetArea                               GodotReal
	GodotAabbHasNoArea                             GodotBool
	GodotAabbHasNoSurface                          GodotBool
	GodotAabbIntersects                            GodotBool
	GodotAabbEncloses                              GodotBool
	GodotAabbMerge                                 GodotAabb
	GodotAabbIntersection                          GodotAabb
	GodotAabbIntersectsPlane                       GodotBool
	GodotAabbIntersectsSegment                     GodotBool
	GodotAabbHasPoint                              GodotBool
	GodotAabbGetSupport                            GodotVector3
	GodotAabbGetLongestAxis                        GodotVector3
	GodotAabbGetLongestAxisIndex                   GodotInt
	GodotAabbGetLongestAxisSize                    GodotReal
	GodotAabbGetShortestAxis                       GodotVector3
	GodotAabbGetShortestAxisIndex                  GodotInt
	GodotAabbGetShortestAxisSize                   GodotReal
	GodotAabbExpand                                GodotAabb
	GodotAabbGrow                                  GodotAabb
	GodotAabbGetEndpoint                           GodotVector3
	GodotAabbOperatorEqual                         GodotBool
	GodotRidNew                                    *func(rDest []GodotRid)
	GodotRidGetId                                  GodotInt
	GodotRidNewWithResource                        *func(rDest []GodotRid, pFrom *GodotObject)
	GodotRidOperatorEqual                          GodotBool
	GodotRidOperatorLess                           GodotBool
	GodotTransformNewWithAxisOrigin                *func(rDest []GodotTransform, pXAxis []GodotVector3, pYAxis []GodotVector3, pZAxis []GodotVector3, pOrigin []GodotVector3)
	GodotTransformNew                              *func(rDest []GodotTransform, pBasis []GodotBasis, pOrigin []GodotVector3)
	GodotTransformGetBasis                         GodotBasis
	GodotTransformSetBasis                         *func(pSelf []GodotTransform, pV []GodotBasis)
	GodotTransformGetOrigin                        GodotVector3
	GodotTransformSetOrigin                        *func(pSelf []GodotTransform, pV []GodotVector3)
	GodotTransformAsString                         GodotString
	GodotTransformInverse                          GodotTransform
	GodotTransformAffineInverse                    GodotTransform
	GodotTransformOrthonormalized                  GodotTransform
	GodotTransformRotated                          GodotTransform
	GodotTransformScaled                           GodotTransform
	GodotTransformTranslated                       GodotTransform
	GodotTransformLookingAt                        GodotTransform
	GodotTransformXformPlane                       GodotPlane
	GodotTransformXformInvPlane                    GodotPlane
	GodotTransformNewIdentity                      *func(rDest []GodotTransform)
	GodotTransformOperatorEqual                    GodotBool
	GodotTransformOperatorMultiply                 GodotTransform
	GodotTransformXformVector3                     GodotVector3
	GodotTransformXformInvVector3                  GodotVector3
	GodotTransformXformAabb                        GodotAabb
	GodotTransformXformInvAabb                     GodotAabb
	GodotTransform2dNew                            *func(rDest []GodotTransform2d, pRot GodotReal, pPos []GodotVector2)
	GodotTransform2dNewAxisOrigin                  *func(rDest []GodotTransform2d, pXAxis []GodotVector2, pYAxis []GodotVector2, pOrigin []GodotVector2)
	GodotTransform2dAsString                       GodotString
	GodotTransform2dInverse                        GodotTransform2d
	GodotTransform2dAffineInverse                  GodotTransform2d
	GodotTransform2dGetRotation                    GodotReal
	GodotTransform2dGetOrigin                      GodotVector2
	GodotTransform2dGetScale                       GodotVector2
	GodotTransform2dOrthonormalized                GodotTransform2d
	GodotTransform2dRotated                        GodotTransform2d
	GodotTransform2dScaled                         GodotTransform2d
	GodotTransform2dTranslated                     GodotTransform2d
	GodotTransform2dXformVector2                   GodotVector2
	GodotTransform2dXformInvVector2                GodotVector2
	GodotTransform2dBasisXformVector2              GodotVector2
	GodotTransform2dBasisXformInvVector2           GodotVector2
	GodotTransform2dInterpolateWith                GodotTransform2d
	GodotTransform2dOperatorEqual                  GodotBool
	GodotTransform2dOperatorMultiply               GodotTransform2d
	GodotTransform2dNewIdentity                    *func(rDest []GodotTransform2d)
	GodotTransform2dXformRect2                     GodotRect2
	GodotTransform2dXformInvRect2                  GodotRect2
	GodotVariantGetType                            GodotVariantType
	GodotVariantNewCopy                            *func(rDest []GodotVariant, pSrc []GodotVariant)
	GodotVariantNewNil                             *func(rDest []GodotVariant)
	GodotVariantNewBool                            *func(pV []GodotVariant, pB GodotBool)
	GodotVariantNewUint                            *func(rDest []GodotVariant, pI uint64)
	GodotVariantNewInt                             *func(rDest []GodotVariant, pI int64)
	GodotVariantNewReal                            *func(rDest []GodotVariant, pR float64)
	GodotVariantNewString                          *func(rDest []GodotVariant, pS []GodotString)
	GodotVariantNewVector2                         *func(rDest []GodotVariant, pV2 []GodotVector2)
	GodotVariantNewRect2                           *func(rDest []GodotVariant, pRect2 []GodotRect2)
	GodotVariantNewVector3                         *func(rDest []GodotVariant, pV3 []GodotVector3)
	GodotVariantNewTransform2d                     *func(rDest []GodotVariant, pT2d []GodotTransform2d)
	GodotVariantNewPlane                           *func(rDest []GodotVariant, pPlane []GodotPlane)
	GodotVariantNewQuat                            *func(rDest []GodotVariant, pQuat []GodotQuat)
	GodotVariantNewAabb                            *func(rDest []GodotVariant, pAabb []GodotAabb)
	GodotVariantNewBasis                           *func(rDest []GodotVariant, pBasis []GodotBasis)
	GodotVariantNewTransform                       *func(rDest []GodotVariant, pTrans []GodotTransform)
	GodotVariantNewColor                           *func(rDest []GodotVariant, pColor []GodotColor)
	GodotVariantNewNodePath                        *func(rDest []GodotVariant, pNp []GodotNodePath)
	GodotVariantNewRid                             *func(rDest []GodotVariant, pRid []GodotRid)
	GodotVariantNewObject                          *func(rDest []GodotVariant, pObj *GodotObject)
	GodotVariantNewDictionary                      *func(rDest []GodotVariant, pDict []GodotDictionary)
	GodotVariantNewArray                           *func(rDest []GodotVariant, pArr []GodotArray)
	GodotVariantNewPoolByteArray                   *func(rDest []GodotVariant, pPba []GodotPoolByteArray)
	GodotVariantNewPoolIntArray                    *func(rDest []GodotVariant, pPia []GodotPoolIntArray)
	GodotVariantNewPoolRealArray                   *func(rDest []GodotVariant, pPra []GodotPoolRealArray)
	GodotVariantNewPoolStringArray                 *func(rDest []GodotVariant, pPsa []GodotPoolStringArray)
	GodotVariantNewPoolVector2Array                *func(rDest []GodotVariant, pPv2a []GodotPoolVector2Array)
	GodotVariantNewPoolVector3Array                *func(rDest []GodotVariant, pPv3a []GodotPoolVector3Array)
	GodotVariantNewPoolColorArray                  *func(rDest []GodotVariant, pPca []GodotPoolColorArray)
	GodotVariantAsBool                             GodotBool
	GodotVariantAsUint                             Uint64T
	GodotVariantAsInt                              Int64T
	GodotVariantAsReal                             *func(pSelf []GodotVariant) float64
	GodotVariantAsString                           GodotString
	GodotVariantAsVector2                          GodotVector2
	GodotVariantAsRect2                            GodotRect2
	GodotVariantAsVector3                          GodotVector3
	GodotVariantAsTransform2d                      GodotTransform2d
	GodotVariantAsPlane                            GodotPlane
	GodotVariantAsQuat                             GodotQuat
	GodotVariantAsAabb                             GodotAabb
	GodotVariantAsBasis                            GodotBasis
	GodotVariantAsTransform                        GodotTransform
	GodotVariantAsColor                            GodotColor
	GodotVariantAsNodePath                         GodotNodePath
	GodotVariantAsRid                              GodotRid
	GodotVariantAsObject                           GodotObject
	GodotVariantAsDictionary                       GodotDictionary
	GodotVariantAsArray                            GodotArray
	GodotVariantAsPoolByteArray                    GodotPoolByteArray
	GodotVariantAsPoolIntArray                     GodotPoolIntArray
	GodotVariantAsPoolRealArray                    GodotPoolRealArray
	GodotVariantAsPoolStringArray                  GodotPoolStringArray
	GodotVariantAsPoolVector2Array                 GodotPoolVector2Array
	GodotVariantAsPoolVector3Array                 GodotPoolVector3Array
	GodotVariantAsPoolColorArray                   GodotPoolColorArray
	GodotVariantCall                               GodotVariant
	GodotVariantHasMethod                          GodotBool
	GodotVariantOperatorEqual                      GodotBool
	GodotVariantOperatorLess                       GodotBool
	GodotVariantHashCompare                        GodotBool
	GodotVariantBooleanize                         GodotBool
	GodotVariantDestroy                            *func(pSelf []GodotVariant)
	GodotStringNew                                 *func(rDest []GodotString)
	GodotStringNewCopy                             *func(rDest []GodotString, pSrc []GodotString)
	GodotStringNewData                             *func(rDest []GodotString, pContents string, pSize int32)
	GodotStringNewUnicodeData                      *func(rDest []GodotString, pContents []int32, pSize int32)
	GodotStringGetData                             *func(pSelf []GodotString, pDest []byte, pSize []int32)
	GodotStringOperatorIndex                       WcharT
	GodotStringOperatorIndexConst                  WcharT
	GodotStringUnicodeStr                          WcharT
	GodotStringOperatorEqual                       GodotBool
	GodotStringOperatorLess                        GodotBool
	GodotStringOperatorPlus                        GodotString
	GodotStringLength                              GodotInt
	GodotStringBeginsWith                          GodotBool
	GodotStringBeginsWithCharArray                 GodotBool
	GodotStringBigrams                             GodotArray
	GodotStringChr                                 GodotString
	GodotStringEndsWith                            GodotBool
	GodotStringFind                                GodotInt
	GodotStringFindFrom                            GodotInt
	GodotStringFindmk                              GodotInt
	GodotStringFindmkFrom                          GodotInt
	GodotStringFindmkFromInPlace                   GodotInt
	GodotStringFindn                               GodotInt
	GodotStringFindnFrom                           GodotInt
	GodotStringFindLast                            GodotInt
	GodotStringFormat                              GodotString
	GodotStringFormatWithCustomPlaceholder         GodotString
	GodotStringHexEncodeBuffer                     GodotString
	GodotStringHexToInt                            GodotInt
	GodotStringHexToIntWithoutPrefix               GodotInt
	GodotStringInsert                              GodotString
	GodotStringIsNumeric                           GodotBool
	GodotStringIsSubsequenceOf                     GodotBool
	GodotStringIsSubsequenceOfi                    GodotBool
	GodotStringLpad                                GodotString
	GodotStringLpadWithCustomCharacter             GodotString
	GodotStringMatch                               GodotBool
	GodotStringMatchn                              GodotBool
	GodotStringMd5                                 GodotString
	GodotStringNum                                 GodotString
	GodotStringNumInt64                            GodotString
	GodotStringNumInt64Capitalized                 GodotString
	GodotStringNumReal                             GodotString
	GodotStringNumScientific                       GodotString
	GodotStringNumWithDecimals                     GodotString
	GodotStringPadDecimals                         GodotString
	GodotStringPadZeros                            GodotString
	GodotStringReplaceFirst                        GodotString
	GodotStringReplace                             GodotString
	GodotStringReplacen                            GodotString
	GodotStringRfind                               GodotInt
	GodotStringRfindn                              GodotInt
	GodotStringRfindFrom                           GodotInt
	GodotStringRfindnFrom                          GodotInt
	GodotStringRpad                                GodotString
	GodotStringRpadWithCustomCharacter             GodotString
	GodotStringSimilarity                          GodotReal
	GodotStringSprintf                             GodotString
	GodotStringSubstr                              GodotString
	GodotStringToDouble                            *func(pSelf []GodotString) float64
	GodotStringToFloat                             GodotReal
	GodotStringToInt                               GodotInt
	GodotStringCamelcaseToUnderscore               GodotString
	GodotStringCamelcaseToUnderscoreLowercased     GodotString
	GodotStringCapitalize                          GodotString
	GodotStringCharToDouble                        *func(pWhat string) float64
	GodotStringCharToInt                           GodotInt
	GodotStringWcharToInt                          Int64T
	GodotStringCharToIntWithLen                    GodotInt
	GodotStringCharToInt64WithLen                  Int64T
	GodotStringHexToInt64                          Int64T
	GodotStringHexToInt64WithPrefix                Int64T
	GodotStringToInt64                             Int64T
	GodotStringUnicodeCharToDouble                 *func(pStr []int32, rEnd [][]int32) float64
	GodotStringGetSliceCount                       GodotInt
	GodotStringGetSlice                            GodotString
	GodotStringGetSlicec                           GodotString
	GodotStringSplit                               GodotArray
	GodotStringSplitAllowEmpty                     GodotArray
	GodotStringSplitFloats                         GodotArray
	GodotStringSplitFloatsAllowsEmpty              GodotArray
	GodotStringSplitFloatsMk                       GodotArray
	GodotStringSplitFloatsMkAllowsEmpty            GodotArray
	GodotStringSplitInts                           GodotArray
	GodotStringSplitIntsAllowsEmpty                GodotArray
	GodotStringSplitIntsMk                         GodotArray
	GodotStringSplitIntsMkAllowsEmpty              GodotArray
	GodotStringSplitSpaces                         GodotArray
	GodotStringCharLowercase                       WcharT
	GodotStringCharUppercase                       WcharT
	GodotStringToLower                             GodotString
	GodotStringToUpper                             GodotString
	GodotStringGetBasename                         GodotString
	GodotStringGetExtension                        GodotString
	GodotStringLeft                                GodotString
	GodotStringOrdAt                               WcharT
	GodotStringPlusFile                            GodotString
	GodotStringRight                               GodotString
	GodotStringStripEdges                          GodotString
	GodotStringStripEscapes                        GodotString
	GodotStringErase                               *func(pSelf []GodotString, pPos GodotInt, pChars GodotInt)
	GodotStringAscii                               *func(pSelf []GodotString, result []byte)
	GodotStringAsciiExtended                       *func(pSelf []GodotString, result []byte)
	GodotStringUtf8                                *func(pSelf []GodotString, result []byte)
	GodotStringParseUtf8                           GodotBool
	GodotStringParseUtf8WithLen                    GodotBool
	GodotStringCharsToUtf8                         GodotString
	GodotStringCharsToUtf8WithLen                  GodotString
	GodotStringHash                                Uint32T
	GodotStringHash64                              Uint64T
	GodotStringHashChars                           Uint32T
	GodotStringHashCharsWithLen                    Uint32T
	GodotStringHashUtf8Chars                       Uint32T
	GodotStringHashUtf8CharsWithLen                Uint32T
	GodotStringMd5Buffer                           GodotPoolByteArray
	GodotStringMd5Text                             GodotString
	GodotStringSha256Buffer                        GodotPoolByteArray
	GodotStringSha256Text                          GodotString
	GodotStringEmpty                               GodotBool
	GodotStringGetBaseDir                          GodotString
	GodotStringGetFile                             GodotString
	GodotStringHumanizeSize                        GodotString
	GodotStringIsAbsPath                           GodotBool
	GodotStringIsRelPath                           GodotBool
	GodotStringIsResourceFile                      GodotBool
	GodotStringPathTo                              GodotString
	GodotStringPathToFile                          GodotString
	GodotStringSimplifyPath                        GodotString
	GodotStringCEscape                             GodotString
	GodotStringCEscapeMultiline                    GodotString
	GodotStringCUnescape                           GodotString
	GodotStringHttpEscape                          GodotString
	GodotStringHttpUnescape                        GodotString
	GodotStringJsonEscape                          GodotString
	GodotStringWordWrap                            GodotString
	GodotStringXmlEscape                           GodotString
	GodotStringXmlEscapeWithQuotes                 GodotString
	GodotStringXmlUnescape                         GodotString
	GodotStringPercentDecode                       GodotString
	GodotStringPercentEncode                       GodotString
	GodotStringIsValidFloat                        GodotBool
	GodotStringIsValidHexNumber                    GodotBool
	GodotStringIsValidHtmlColor                    GodotBool
	GodotStringIsValidIdentifier                   GodotBool
	GodotStringIsValidInteger                      GodotBool
	GodotStringIsValidIpAddress                    GodotBool
	GodotStringDestroy                             *func(pSelf []GodotString)
	GodotStringNameNew                             *func(rDest []GodotStringName, pName []GodotString)
	GodotStringNameNewData                         *func(rDest []GodotStringName, pName string)
	GodotStringNameGetName                         GodotString
	GodotStringNameGetHash                         Uint32T
	GodotStringNameGetDataUniquePointer            *func(pSelf []GodotStringName) unsafe.Pointer
	GodotStringNameOperatorEqual                   GodotBool
	GodotStringNameOperatorLess                    GodotBool
	GodotStringNameDestroy                         *func(pSelf []GodotStringName)
	GodotObjectDestroy                             *func(pO *GodotObject)
	GodotGlobalGetSingleton                        GodotObject
	GodotMethodBindGetMethod                       GodotMethodBind
	GodotMethodBindPtrcall                         *func(pMethodBind []GodotMethodBind, pInstance *GodotObject, pArgs []unsafe.Pointer, pRet unsafe.Pointer)
	GodotMethodBindCall                            GodotVariant
	GodotGetClassConstructor                       GodotClassConstructor
	GodotRegisterNativeCallType                    *func(callType string, pCallback NativeCallCb)
	GodotAlloc                                     *func(pBytes int32) unsafe.Pointer
	GodotRealloc                                   *func(pPtr unsafe.Pointer, pBytes int32) unsafe.Pointer
	GodotFree                                      *func(pPtr unsafe.Pointer)
	GodotPrintError                                *func(pDescription string, pFunction string, pFile string, pLine int32)
	GodotPrintWarning                              *func(pDescription string, pFunction string, pFile string, pLine int32)
	GodotPrint                                     *func(pMessage []GodotString)
	ref57717e51                                    *C.godot_gdnative_core_api_struct
	allocs57717e51                                 interface{}
}

// GodotBool type as declared in gdnative/gdnative.h:124
type GodotBool bool

// GodotInt type as declared in gdnative/gdnative.h:131
type GodotInt int32

// GodotReal type as declared in gdnative/gdnative.h:135
type GodotReal float32

// GodotObject type as declared in gdnative/gdnative.h:138
type GodotObject [0]byte

// GodotMethodBind as declared in gdnative/gdnative.h:225
type GodotMethodBind struct {
	DontTouchThat  [1]byte
	ref3a05c0bc    *C.godot_method_bind
	allocs3a05c0bc interface{}
}

// GodotGdnativeApiVersion as declared in gdnative/gdnative.h:235
type GodotGdnativeApiVersion struct {
	Major          uint32
	Minor          uint32
	ref5eed2c27    *C.godot_gdnative_api_version
	allocs5eed2c27 interface{}
}

// GodotGdnativeApiStruct as declared in gdnative/gdnative.h:237
type GodotGdnativeApiStruct struct {
	Type           uint32
	Version        GodotGdnativeApiVersion
	Next           []GodotGdnativeApiStruct
	ref45f52b65    *C.godot_gdnative_api_struct
	allocs45f52b65 interface{}
}

// GodotGdnativeInitOptions as declared in gdnative/gdnative.h:257
type GodotGdnativeInitOptions struct {
	InEditor              GodotBool
	CoreApiHash           uint64
	EditorApiHash         uint64
	NoApiHash             uint64
	ReportVersionMismatch *func(pLibrary *GodotObject, pWhat string, pWant GodotGdnativeApiVersion, pHave GodotGdnativeApiVersion)
	ReportLoadingError    *func(pLibrary *GodotObject, pWhat string)
	GdNativeLibrary       *GodotObject
	ApiStruct             []GodotGdnativeCoreApiStruct
	ActiveLibraryPath     []GodotString
	reff9d34929           *C.godot_gdnative_init_options
	allocsf9d34929        interface{}
}

// GodotGdnativeTerminateOptions as declared in gdnative/gdnative.h:261
type GodotGdnativeTerminateOptions struct {
	InEditor       GodotBool
	ref80c63fdc    *C.godot_gdnative_terminate_options
	allocs80c63fdc interface{}
}

// GodotClassConstructor type as declared in gdnative/gdnative.h:264
type GodotClassConstructor func() *GodotObject

// GodotGdnativeInitFn type as declared in gdnative/gdnative.h:271
type GodotGdnativeInitFn func(arg0 []GodotGdnativeInitOptions)

// GodotGdnativeTerminateFn type as declared in gdnative/gdnative.h:272
type GodotGdnativeTerminateFn func(arg0 []GodotGdnativeTerminateOptions)

// GodotGdnativeProcedureFn type as declared in gdnative/gdnative.h:273
type GodotGdnativeProcedureFn func(arg0 []GodotArray) GodotVariant

// GodotString as declared in gdnative/string.h:46
type GodotString struct {
	DontTouchThat  [8]byte
	ref6d1ede57    *C.godot_string
	allocs6d1ede57 interface{}
}

// GodotArray as declared in gdnative/array.h:46
type GodotArray struct {
	DontTouchThat  [8]byte
	refb81158a5    *C.godot_array
	allocsb81158a5 interface{}
}

// GodotPoolArrayReadAccess as declared in gdnative/pool_arrays.h:45
type GodotPoolArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolByteArrayReadAccess as declared in gdnative/pool_arrays.h:47
type GodotPoolByteArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolIntArrayReadAccess as declared in gdnative/pool_arrays.h:48
type GodotPoolIntArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolRealArrayReadAccess as declared in gdnative/pool_arrays.h:49
type GodotPoolRealArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolStringArrayReadAccess as declared in gdnative/pool_arrays.h:50
type GodotPoolStringArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolVector2ArrayReadAccess as declared in gdnative/pool_arrays.h:51
type GodotPoolVector2ArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolVector3ArrayReadAccess as declared in gdnative/pool_arrays.h:52
type GodotPoolVector3ArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolColorArrayReadAccess as declared in gdnative/pool_arrays.h:53
type GodotPoolColorArrayReadAccess struct {
	DontTouchThat  [1]byte
	ref172179be    *C.godot_pool_array_read_access
	allocs172179be interface{}
}

// GodotPoolArrayWriteAccess as declared in gdnative/pool_arrays.h:61
type GodotPoolArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolByteArrayWriteAccess as declared in gdnative/pool_arrays.h:63
type GodotPoolByteArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolIntArrayWriteAccess as declared in gdnative/pool_arrays.h:64
type GodotPoolIntArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolRealArrayWriteAccess as declared in gdnative/pool_arrays.h:65
type GodotPoolRealArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolStringArrayWriteAccess as declared in gdnative/pool_arrays.h:66
type GodotPoolStringArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolVector2ArrayWriteAccess as declared in gdnative/pool_arrays.h:67
type GodotPoolVector2ArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolVector3ArrayWriteAccess as declared in gdnative/pool_arrays.h:68
type GodotPoolVector3ArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolColorArrayWriteAccess as declared in gdnative/pool_arrays.h:69
type GodotPoolColorArrayWriteAccess struct {
	DontTouchThat [1]byte
	refbe0648b    *C.godot_pool_array_write_access
	allocsbe0648b interface{}
}

// GodotPoolByteArray as declared in gdnative/pool_arrays.h:79
type GodotPoolByteArray struct {
	DontTouchThat [8]byte
	refbf60ed2    *C.godot_pool_byte_array
	allocsbf60ed2 interface{}
}

// GodotPoolIntArray as declared in gdnative/pool_arrays.h:90
type GodotPoolIntArray struct {
	DontTouchThat  [8]byte
	ref5d4f26e6    *C.godot_pool_int_array
	allocs5d4f26e6 interface{}
}

// GodotPoolRealArray as declared in gdnative/pool_arrays.h:101
type GodotPoolRealArray struct {
	DontTouchThat  [8]byte
	refc76f44c3    *C.godot_pool_real_array
	allocsc76f44c3 interface{}
}

// GodotPoolStringArray as declared in gdnative/pool_arrays.h:112
type GodotPoolStringArray struct {
	DontTouchThat [8]byte
	reff6fe5d9    *C.godot_pool_string_array
	allocsf6fe5d9 interface{}
}

// GodotPoolVector2Array as declared in gdnative/pool_arrays.h:123
type GodotPoolVector2Array struct {
	DontTouchThat  [8]byte
	ref7f6b2885    *C.godot_pool_vector2_array
	allocs7f6b2885 interface{}
}

// GodotPoolVector3Array as declared in gdnative/pool_arrays.h:134
type GodotPoolVector3Array struct {
	DontTouchThat  [8]byte
	refd91c2331    *C.godot_pool_vector3_array
	allocsd91c2331 interface{}
}

// GodotPoolColorArray as declared in gdnative/pool_arrays.h:145
type GodotPoolColorArray struct {
	DontTouchThat  [8]byte
	refd5cae78e    *C.godot_pool_color_array
	allocsd5cae78e interface{}
}

// GodotColor as declared in gdnative/color.h:45
type GodotColor struct {
	DontTouchThat  [16]byte
	ref7f4bfefb    *C.godot_color
	allocs7f4bfefb interface{}
}

// GodotVector2 as declared in gdnative/vector2.h:45
type GodotVector2 struct {
	DontTouchThat  [8]byte
	refbc81274e    *C.godot_vector2
	allocsbc81274e interface{}
}

// GodotVector3 as declared in gdnative/vector3.h:45
type GodotVector3 struct {
	DontTouchThat  [12]byte
	refcb8617d8    *C.godot_vector3
	allocscb8617d8 interface{}
}

// GodotBasis as declared in gdnative/basis.h:45
type GodotBasis struct {
	DontTouchThat  [36]byte
	ref94d3d325    *C.godot_basis
	allocs94d3d325 interface{}
}

// GodotQuat as declared in gdnative/quat.h:45
type GodotQuat struct {
	DontTouchThat  [16]byte
	reffaf33e0b    *C.godot_quat
	allocsfaf33e0b interface{}
}

// GodotVariant as declared in gdnative/variant.h:45
type GodotVariant struct {
	DontTouchThat  [24]byte
	refabb5c0da    *C.godot_variant
	allocsabb5c0da interface{}
}

// GodotVariantCallError as declared in gdnative/variant.h:100
type GodotVariantCallError struct {
	Error          GodotVariantCallErrorError
	Argument       int32
	Expected       GodotVariantType
	ref3ce71027    *C.godot_variant_call_error
	allocs3ce71027 interface{}
}

// GodotAabb as declared in gdnative/aabb.h:45
type GodotAabb struct {
	DontTouchThat  [24]byte
	ref6e3c84aa    *C.godot_aabb
	allocs6e3c84aa interface{}
}

// GodotPlane as declared in gdnative/plane.h:45
type GodotPlane struct {
	DontTouchThat  [16]byte
	refd8ae9b92    *C.godot_plane
	allocsd8ae9b92 interface{}
}

// GodotDictionary as declared in gdnative/dictionary.h:45
type GodotDictionary struct {
	DontTouchThat  [8]byte
	refb267a9b9    *C.godot_dictionary
	allocsb267a9b9 interface{}
}

// GodotNodePath as declared in gdnative/node_path.h:45
type GodotNodePath struct {
	DontTouchThat  [8]byte
	ref6c34dff3    *C.godot_node_path
	allocs6c34dff3 interface{}
}

// GodotRect2 as declared in gdnative/rect2.h:43
type GodotRect2 struct {
	DontTouchThat  [16]byte
	ref99c06d9a    *C.godot_rect2
	allocs99c06d9a interface{}
}

// GodotRid as declared in gdnative/rid.h:45
type GodotRid struct {
	DontTouchThat  [8]byte
	ref67320fc7    *C.godot_rid
	allocs67320fc7 interface{}
}

// GodotTransform as declared in gdnative/transform.h:45
type GodotTransform struct {
	DontTouchThat  [48]byte
	refd77658c7    *C.godot_transform
	allocsd77658c7 interface{}
}

// GodotTransform2d as declared in gdnative/transform2d.h:45
type GodotTransform2d struct {
	DontTouchThat [24]byte
	ref77dacf6    *C.godot_transform2d
	allocs77dacf6 interface{}
}

// GodotStringName as declared in gdnative/string_name.h:46
type GodotStringName struct {
	DontTouchThat  [8]byte
	ref895548fc    *C.godot_string_name
	allocs895548fc interface{}
}

// GodotArvrInterfaceGdnative as declared in arvr/godot_arvr.h:55
type GodotArvrInterfaceGdnative struct {
	Constructor                 *func(arg0 *GodotObject) unsafe.Pointer
	Destructor                  *func(arg0 unsafe.Pointer)
	GetName                     GodotString
	GetCapabilities             GodotInt
	GetAnchorDetectionIsEnabled GodotBool
	SetAnchorDetectionIsEnabled *func(arg0 unsafe.Pointer, arg1 GodotBool)
	IsStereo                    GodotBool
	IsInitialized               GodotBool
	Initialize                  GodotBool
	Uninitialize                *func(arg0 unsafe.Pointer)
	GetRenderTargetsize         GodotVector2
	GetTransformForEye          GodotTransform
	FillProjectionForEye        *func(arg0 unsafe.Pointer, arg1 []GodotReal, arg2 GodotInt, arg3 GodotReal, arg4 GodotReal, arg5 GodotReal)
	CommitForEye                *func(arg0 unsafe.Pointer, arg1 GodotInt, arg2 []GodotRid, arg3 []GodotRect2)
	Process                     *func(arg0 unsafe.Pointer)
	reff96a0b88                 *C.godot_arvr_interface_gdnative
	allocsf96a0b88              interface{}
}

// GodotPropertyAttributes as declared in nativescript/godot_nativescript.h:117
type GodotPropertyAttributes struct {
	RsetType       GodotMethodRpcMode
	Type           GodotInt
	Hint           GodotPropertyHint
	HintString     GodotString
	Usage          GodotPropertyUsageFlags
	DefaultValue   GodotVariant
	ref431c473b    *C.godot_property_attributes
	allocs431c473b interface{}
}

// GodotInstanceCreateFunc as declared in nativescript/godot_nativescript.h:124
type GodotInstanceCreateFunc struct {
	CreateFunc     *func(arg0 *GodotObject, arg1 unsafe.Pointer) unsafe.Pointer
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	ref70ecb5db    *C.godot_instance_create_func
	allocs70ecb5db interface{}
}

// GodotInstanceDestroyFunc as declared in nativescript/godot_nativescript.h:131
type GodotInstanceDestroyFunc struct {
	DestroyFunc    *func(arg0 *GodotObject, arg1 unsafe.Pointer, arg2 unsafe.Pointer)
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	refd0d05668    *C.godot_instance_destroy_func
	allocsd0d05668 interface{}
}

// GodotMethodAttributes as declared in nativescript/godot_nativescript.h:139
type GodotMethodAttributes struct {
	RpcType        GodotMethodRpcMode
	ref66a6c5c9    *C.godot_method_attributes
	allocs66a6c5c9 interface{}
}

// GodotInstanceMethod as declared in nativescript/godot_nativescript.h:146
type GodotInstanceMethod struct {
	Method         GodotVariant
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	ref10e1583e    *C.godot_instance_method
	allocs10e1583e interface{}
}

// GodotPropertySetFunc as declared in nativescript/godot_nativescript.h:155
type GodotPropertySetFunc struct {
	SetFunc       *func(arg0 *GodotObject, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 []GodotVariant)
	MethodData    unsafe.Pointer
	FreeFunc      *func(arg0 unsafe.Pointer)
	refc9844af    *C.godot_property_set_func
	allocsc9844af interface{}
}

// GodotPropertyGetFunc as declared in nativescript/godot_nativescript.h:162
type GodotPropertyGetFunc struct {
	GetFunc        GodotVariant
	MethodData     unsafe.Pointer
	FreeFunc       *func(arg0 unsafe.Pointer)
	reff4697b7e    *C.godot_property_get_func
	allocsf4697b7e interface{}
}

// GodotSignalArgument as declared in nativescript/godot_nativescript.h:173
type GodotSignalArgument struct {
	Name           GodotString
	Type           GodotInt
	Hint           GodotPropertyHint
	HintString     GodotString
	Usage          GodotPropertyUsageFlags
	DefaultValue   GodotVariant
	refc21e72ac    *C.godot_signal_argument
	allocsc21e72ac interface{}
}

// GodotSignal as declared in nativescript/godot_nativescript.h:181
type GodotSignal struct {
	Name           GodotString
	NumArgs        int32
	Args           []GodotSignalArgument
	NumDefaultArgs int32
	DefaultArgs    []GodotVariant
	ref87acf90b    *C.godot_signal
	allocs87acf90b interface{}
}

// GodotPluginscriptInstanceData type as declared in pluginscript/godot_pluginscript.h:40
type GodotPluginscriptInstanceData [0]byte

// GodotPluginscriptScriptData type as declared in pluginscript/godot_pluginscript.h:41
type GodotPluginscriptScriptData [0]byte

// GodotPluginscriptLanguageData type as declared in pluginscript/godot_pluginscript.h:42
type GodotPluginscriptLanguageData [0]byte

// GodotPluginscriptInstanceDesc as declared in pluginscript/godot_pluginscript.h:69
type GodotPluginscriptInstanceDesc struct {
	Init                GodotPluginscriptInstanceData
	Finish              *func(pData *GodotPluginscriptInstanceData)
	SetProp             GodotBool
	GetProp             GodotBool
	CallMethod          GodotVariant
	Notification        *func(pData *GodotPluginscriptInstanceData, pNotification int32)
	GetRpcMode          GodotMethodRpcMode
	GetRsetMode         GodotMethodRpcMode
	RefcountIncremented *func(pData *GodotPluginscriptInstanceData)
	RefcountDecremented *func(pData *GodotPluginscriptInstanceData) bool
	refc0c19139         *C.godot_pluginscript_instance_desc
	allocsc0c19139      interface{}
}

// GodotPluginscriptScriptManifest as declared in pluginscript/godot_pluginscript.h:104
type GodotPluginscriptScriptManifest struct {
	Data           *GodotPluginscriptScriptData
	Name           GodotStringName
	IsTool         GodotBool
	Base           GodotStringName
	MemberLines    GodotDictionary
	Methods        GodotArray
	Signals        GodotArray
	Properties     GodotArray
	reffbf02dfd    *C.godot_pluginscript_script_manifest
	allocsfbf02dfd interface{}
}

// GodotPluginscriptScriptDesc as declared in pluginscript/godot_pluginscript.h:110
type GodotPluginscriptScriptDesc struct {
	Init           GodotPluginscriptScriptManifest
	Finish         *func(pData *GodotPluginscriptScriptData)
	InstanceDesc   GodotPluginscriptInstanceDesc
	ref1aab3210    *C.godot_pluginscript_script_desc
	allocs1aab3210 interface{}
}

// GodotPluginscriptProfilingData as declared in pluginscript/godot_pluginscript.h:119
type GodotPluginscriptProfilingData struct {
	Signature      GodotStringName
	CallCount      GodotInt
	TotalTime      GodotInt
	SelfTime       GodotInt
	ref9c004e5a    *C.godot_pluginscript_profiling_data
	allocs9c004e5a interface{}
}

// GodotPluginscriptLanguageDesc as declared in pluginscript/godot_pluginscript.h:163
type GodotPluginscriptLanguageDesc struct {
	Name                           string
	Type                           string
	Extension                      string
	RecognizedExtensions           []string
	Init                           GodotPluginscriptLanguageData
	Finish                         *func(pData *GodotPluginscriptLanguageData)
	ReservedWords                  []string
	CommentDelimiters              []string
	StringDelimiters               []string
	HasNamedClasses                GodotBool
	SupportsBuiltinMode            GodotBool
	GetTemplateSourceCode          GodotString
	Validate                       GodotBool
	FindFunction                   *func(pData *GodotPluginscriptLanguageData, pFunction []GodotString, pCode []GodotString) int32
	MakeFunction                   GodotString
	CompleteCode                   GodotError
	AutoIndentCode                 *func(pData *GodotPluginscriptLanguageData, pCode []GodotString, pFromLine int32, pToLine int32)
	AddGlobalConstant              *func(pData *GodotPluginscriptLanguageData, pVariable []GodotString, pValue []GodotVariant)
	DebugGetError                  GodotString
	DebugGetStackLevelCount        *func(pData *GodotPluginscriptLanguageData) int32
	DebugGetStackLevelLine         *func(pData *GodotPluginscriptLanguageData, pLevel int32) int32
	DebugGetStackLevelFunction     GodotString
	DebugGetStackLevelSource       GodotString
	DebugGetStackLevelLocals       *func(pData *GodotPluginscriptLanguageData, pLevel int32, pLocals []GodotPoolStringArray, pValues []GodotArray, pMaxSubitems int32, pMaxDepth int32)
	DebugGetStackLevelMembers      *func(pData *GodotPluginscriptLanguageData, pLevel int32, pMembers []GodotPoolStringArray, pValues []GodotArray, pMaxSubitems int32, pMaxDepth int32)
	DebugGetGlobals                *func(pData *GodotPluginscriptLanguageData, pLocals []GodotPoolStringArray, pValues []GodotArray, pMaxSubitems int32, pMaxDepth int32)
	DebugParseStackLevelExpression GodotString
	GetPublicFunctions             *func(pData *GodotPluginscriptLanguageData, rFunctions []GodotArray)
	GetPublicConstants             *func(pData *GodotPluginscriptLanguageData, rConstants []GodotDictionary)
	ProfilingStart                 *func(pData *GodotPluginscriptLanguageData)
	ProfilingStop                  *func(pData *GodotPluginscriptLanguageData)
	ProfilingGetAccumulatedData    *func(pData *GodotPluginscriptLanguageData, rInfo []GodotPluginscriptProfilingData, pInfoMax int32) int32
	ProfilingGetFrameData          *func(pData *GodotPluginscriptLanguageData, rInfo []GodotPluginscriptProfilingData, pInfoMax int32) int32
	ProfilingFrame                 *func(pData *GodotPluginscriptLanguageData)
	ScriptDesc                     GodotPluginscriptScriptDesc
	refdac22bbe                    *C.godot_pluginscript_language_desc
	allocsdac22bbe                 interface{}
}
