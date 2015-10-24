package gui

//#include "qopengltexture.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QOpenGLTexture struct {
	ptr unsafe.Pointer
}

type QOpenGLTextureITF interface {
	QOpenGLTexturePTR() *QOpenGLTexture
}

func (p *QOpenGLTexture) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLTexture) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLTexture(ptr QOpenGLTextureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLTexturePTR().Pointer()
	}
	return nil
}

func QOpenGLTextureFromPointer(ptr unsafe.Pointer) *QOpenGLTexture {
	var n = new(QOpenGLTexture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLTexture) QOpenGLTexturePTR() *QOpenGLTexture {
	return ptr
}

//QOpenGLTexture::BindingTarget
type QOpenGLTexture__BindingTarget int

var (
	QOpenGLTexture__BindingTarget1D                 = QOpenGLTexture__BindingTarget(0x8068)
	QOpenGLTexture__BindingTarget1DArray            = QOpenGLTexture__BindingTarget(0x8C1C)
	QOpenGLTexture__BindingTarget2D                 = QOpenGLTexture__BindingTarget(0x8069)
	QOpenGLTexture__BindingTarget2DArray            = QOpenGLTexture__BindingTarget(0x8C1D)
	QOpenGLTexture__BindingTarget3D                 = QOpenGLTexture__BindingTarget(0x806A)
	QOpenGLTexture__BindingTargetCubeMap            = QOpenGLTexture__BindingTarget(0x8514)
	QOpenGLTexture__BindingTargetCubeMapArray       = QOpenGLTexture__BindingTarget(0x900A)
	QOpenGLTexture__BindingTarget2DMultisample      = QOpenGLTexture__BindingTarget(0x9104)
	QOpenGLTexture__BindingTarget2DMultisampleArray = QOpenGLTexture__BindingTarget(0x9105)
	QOpenGLTexture__BindingTargetRectangle          = QOpenGLTexture__BindingTarget(0x84F6)
	QOpenGLTexture__BindingTargetBuffer             = QOpenGLTexture__BindingTarget(0x8C2C)
)

//QOpenGLTexture::ComparisonFunction
type QOpenGLTexture__ComparisonFunction int

var (
	QOpenGLTexture__CompareLessEqual    = QOpenGLTexture__ComparisonFunction(0x0203)
	QOpenGLTexture__CompareGreaterEqual = QOpenGLTexture__ComparisonFunction(0x0206)
	QOpenGLTexture__CompareLess         = QOpenGLTexture__ComparisonFunction(0x0201)
	QOpenGLTexture__CompareGreater      = QOpenGLTexture__ComparisonFunction(0x0204)
	QOpenGLTexture__CompareEqual        = QOpenGLTexture__ComparisonFunction(0x0202)
	QOpenGLTexture__CommpareNotEqual    = QOpenGLTexture__ComparisonFunction(0x0205)
	QOpenGLTexture__CompareAlways       = QOpenGLTexture__ComparisonFunction(0x0207)
	QOpenGLTexture__CompareNever        = QOpenGLTexture__ComparisonFunction(0x0200)
)

//QOpenGLTexture::ComparisonMode
type QOpenGLTexture__ComparisonMode int

var (
	QOpenGLTexture__CompareRefToTexture = QOpenGLTexture__ComparisonMode(0x884E)
	QOpenGLTexture__CompareNone         = QOpenGLTexture__ComparisonMode(0x0000)
)

//QOpenGLTexture::CoordinateDirection
type QOpenGLTexture__CoordinateDirection int

var (
	QOpenGLTexture__DirectionS = QOpenGLTexture__CoordinateDirection(0x2802)
	QOpenGLTexture__DirectionT = QOpenGLTexture__CoordinateDirection(0x2803)
	QOpenGLTexture__DirectionR = QOpenGLTexture__CoordinateDirection(0x8072)
)

//QOpenGLTexture::CubeMapFace
type QOpenGLTexture__CubeMapFace int

var (
	QOpenGLTexture__CubeMapPositiveX = QOpenGLTexture__CubeMapFace(0x8515)
	QOpenGLTexture__CubeMapNegativeX = QOpenGLTexture__CubeMapFace(0x8516)
	QOpenGLTexture__CubeMapPositiveY = QOpenGLTexture__CubeMapFace(0x8517)
	QOpenGLTexture__CubeMapNegativeY = QOpenGLTexture__CubeMapFace(0x8518)
	QOpenGLTexture__CubeMapPositiveZ = QOpenGLTexture__CubeMapFace(0x8519)
	QOpenGLTexture__CubeMapNegativeZ = QOpenGLTexture__CubeMapFace(0x851A)
)

//QOpenGLTexture::DepthStencilMode
type QOpenGLTexture__DepthStencilMode int

var (
	QOpenGLTexture__DepthMode   = QOpenGLTexture__DepthStencilMode(0x1902)
	QOpenGLTexture__StencilMode = QOpenGLTexture__DepthStencilMode(0x1901)
)

//QOpenGLTexture::Feature
type QOpenGLTexture__Feature int

var (
	QOpenGLTexture__ImmutableStorage            = QOpenGLTexture__Feature(0x00000001)
	QOpenGLTexture__ImmutableMultisampleStorage = QOpenGLTexture__Feature(0x00000002)
	QOpenGLTexture__TextureRectangle            = QOpenGLTexture__Feature(0x00000004)
	QOpenGLTexture__TextureArrays               = QOpenGLTexture__Feature(0x00000008)
	QOpenGLTexture__Texture3D                   = QOpenGLTexture__Feature(0x00000010)
	QOpenGLTexture__TextureMultisample          = QOpenGLTexture__Feature(0x00000020)
	QOpenGLTexture__TextureBuffer               = QOpenGLTexture__Feature(0x00000040)
	QOpenGLTexture__TextureCubeMapArrays        = QOpenGLTexture__Feature(0x00000080)
	QOpenGLTexture__Swizzle                     = QOpenGLTexture__Feature(0x00000100)
	QOpenGLTexture__StencilTexturing            = QOpenGLTexture__Feature(0x00000200)
	QOpenGLTexture__AnisotropicFiltering        = QOpenGLTexture__Feature(0x00000400)
	QOpenGLTexture__NPOTTextures                = QOpenGLTexture__Feature(0x00000800)
	QOpenGLTexture__NPOTTextureRepeat           = QOpenGLTexture__Feature(0x00001000)
	QOpenGLTexture__Texture1D                   = QOpenGLTexture__Feature(0x00002000)
	QOpenGLTexture__TextureComparisonOperators  = QOpenGLTexture__Feature(0x00004000)
	QOpenGLTexture__TextureMipMapLevel          = QOpenGLTexture__Feature(0x00008000)
)

//QOpenGLTexture::Filter
type QOpenGLTexture__Filter int

var (
	QOpenGLTexture__Nearest              = QOpenGLTexture__Filter(0x2600)
	QOpenGLTexture__Linear               = QOpenGLTexture__Filter(0x2601)
	QOpenGLTexture__NearestMipMapNearest = QOpenGLTexture__Filter(0x2700)
	QOpenGLTexture__NearestMipMapLinear  = QOpenGLTexture__Filter(0x2702)
	QOpenGLTexture__LinearMipMapNearest  = QOpenGLTexture__Filter(0x2701)
	QOpenGLTexture__LinearMipMapLinear   = QOpenGLTexture__Filter(0x2703)
)

//QOpenGLTexture::MipMapGeneration
type QOpenGLTexture__MipMapGeneration int

var (
	QOpenGLTexture__GenerateMipMaps     = QOpenGLTexture__MipMapGeneration(0)
	QOpenGLTexture__DontGenerateMipMaps = QOpenGLTexture__MipMapGeneration(1)
)

//QOpenGLTexture::PixelFormat
type QOpenGLTexture__PixelFormat int

var (
	QOpenGLTexture__NoSourceFormat = QOpenGLTexture__PixelFormat(0)
	QOpenGLTexture__Red            = QOpenGLTexture__PixelFormat(0x1903)
	QOpenGLTexture__RG             = QOpenGLTexture__PixelFormat(0x8227)
	QOpenGLTexture__RGB            = QOpenGLTexture__PixelFormat(0x1907)
	QOpenGLTexture__BGR            = QOpenGLTexture__PixelFormat(0x80E0)
	QOpenGLTexture__RGBA           = QOpenGLTexture__PixelFormat(0x1908)
	QOpenGLTexture__BGRA           = QOpenGLTexture__PixelFormat(0x80E1)
	QOpenGLTexture__Red_Integer    = QOpenGLTexture__PixelFormat(0x8D94)
	QOpenGLTexture__RG_Integer     = QOpenGLTexture__PixelFormat(0x8228)
	QOpenGLTexture__RGB_Integer    = QOpenGLTexture__PixelFormat(0x8D98)
	QOpenGLTexture__BGR_Integer    = QOpenGLTexture__PixelFormat(0x8D9A)
	QOpenGLTexture__RGBA_Integer   = QOpenGLTexture__PixelFormat(0x8D99)
	QOpenGLTexture__BGRA_Integer   = QOpenGLTexture__PixelFormat(0x8D9B)
	QOpenGLTexture__Stencil        = QOpenGLTexture__PixelFormat(0x1901)
	QOpenGLTexture__Depth          = QOpenGLTexture__PixelFormat(0x1902)
	QOpenGLTexture__DepthStencil   = QOpenGLTexture__PixelFormat(0x84F9)
	QOpenGLTexture__Alpha          = QOpenGLTexture__PixelFormat(0x1906)
	QOpenGLTexture__Luminance      = QOpenGLTexture__PixelFormat(0x1909)
	QOpenGLTexture__LuminanceAlpha = QOpenGLTexture__PixelFormat(0x190A)
)

//QOpenGLTexture::PixelType
type QOpenGLTexture__PixelType int

var (
	QOpenGLTexture__NoPixelType               = QOpenGLTexture__PixelType(0)
	QOpenGLTexture__Int8                      = QOpenGLTexture__PixelType(0x1400)
	QOpenGLTexture__UInt8                     = QOpenGLTexture__PixelType(0x1401)
	QOpenGLTexture__Int16                     = QOpenGLTexture__PixelType(0x1402)
	QOpenGLTexture__UInt16                    = QOpenGLTexture__PixelType(0x1403)
	QOpenGLTexture__Int32                     = QOpenGLTexture__PixelType(0x1404)
	QOpenGLTexture__UInt32                    = QOpenGLTexture__PixelType(0x1405)
	QOpenGLTexture__Float16                   = QOpenGLTexture__PixelType(0x140B)
	QOpenGLTexture__Float16OES                = QOpenGLTexture__PixelType(0x8D61)
	QOpenGLTexture__Float32                   = QOpenGLTexture__PixelType(0x1406)
	QOpenGLTexture__UInt32_RGB9_E5            = QOpenGLTexture__PixelType(0x8C3E)
	QOpenGLTexture__UInt32_RG11B10F           = QOpenGLTexture__PixelType(0x8C3B)
	QOpenGLTexture__UInt8_RG3B2               = QOpenGLTexture__PixelType(0x8032)
	QOpenGLTexture__UInt8_RG3B2_Rev           = QOpenGLTexture__PixelType(0x8362)
	QOpenGLTexture__UInt16_RGB5A1             = QOpenGLTexture__PixelType(0x8034)
	QOpenGLTexture__UInt16_RGB5A1_Rev         = QOpenGLTexture__PixelType(0x8366)
	QOpenGLTexture__UInt16_R5G6B5             = QOpenGLTexture__PixelType(0x8363)
	QOpenGLTexture__UInt16_R5G6B5_Rev         = QOpenGLTexture__PixelType(0x8364)
	QOpenGLTexture__UInt16_RGBA4              = QOpenGLTexture__PixelType(0x8033)
	QOpenGLTexture__UInt16_RGBA4_Rev          = QOpenGLTexture__PixelType(0x8365)
	QOpenGLTexture__UInt32_RGBA8              = QOpenGLTexture__PixelType(0x8035)
	QOpenGLTexture__UInt32_RGBA8_Rev          = QOpenGLTexture__PixelType(0x8367)
	QOpenGLTexture__UInt32_RGB10A2            = QOpenGLTexture__PixelType(0x8036)
	QOpenGLTexture__UInt32_RGB10A2_Rev        = QOpenGLTexture__PixelType(0x8368)
	QOpenGLTexture__UInt32_D24S8              = QOpenGLTexture__PixelType(0x84FA)
	QOpenGLTexture__Float32_D32_UInt32_S8_X24 = QOpenGLTexture__PixelType(0x8DAD)
)

//QOpenGLTexture::SwizzleComponent
type QOpenGLTexture__SwizzleComponent int

var (
	QOpenGLTexture__SwizzleRed   = QOpenGLTexture__SwizzleComponent(0x8E42)
	QOpenGLTexture__SwizzleGreen = QOpenGLTexture__SwizzleComponent(0x8E43)
	QOpenGLTexture__SwizzleBlue  = QOpenGLTexture__SwizzleComponent(0x8E44)
	QOpenGLTexture__SwizzleAlpha = QOpenGLTexture__SwizzleComponent(0x8E45)
)

//QOpenGLTexture::SwizzleValue
type QOpenGLTexture__SwizzleValue int

var (
	QOpenGLTexture__RedValue   = QOpenGLTexture__SwizzleValue(0x1903)
	QOpenGLTexture__GreenValue = QOpenGLTexture__SwizzleValue(0x1904)
	QOpenGLTexture__BlueValue  = QOpenGLTexture__SwizzleValue(0x1905)
	QOpenGLTexture__AlphaValue = QOpenGLTexture__SwizzleValue(0x1906)
	QOpenGLTexture__ZeroValue  = QOpenGLTexture__SwizzleValue(0)
	QOpenGLTexture__OneValue   = QOpenGLTexture__SwizzleValue(1)
)

//QOpenGLTexture::Target
type QOpenGLTexture__Target int

var (
	QOpenGLTexture__Target1D                 = QOpenGLTexture__Target(0x0DE0)
	QOpenGLTexture__Target1DArray            = QOpenGLTexture__Target(0x8C18)
	QOpenGLTexture__Target2D                 = QOpenGLTexture__Target(0x0DE1)
	QOpenGLTexture__Target2DArray            = QOpenGLTexture__Target(0x8C1A)
	QOpenGLTexture__Target3D                 = QOpenGLTexture__Target(0x806F)
	QOpenGLTexture__TargetCubeMap            = QOpenGLTexture__Target(0x8513)
	QOpenGLTexture__TargetCubeMapArray       = QOpenGLTexture__Target(0x9009)
	QOpenGLTexture__Target2DMultisample      = QOpenGLTexture__Target(0x9100)
	QOpenGLTexture__Target2DMultisampleArray = QOpenGLTexture__Target(0x9102)
	QOpenGLTexture__TargetRectangle          = QOpenGLTexture__Target(0x84F5)
	QOpenGLTexture__TargetBuffer             = QOpenGLTexture__Target(0x8C2A)
)

//QOpenGLTexture::TextureFormat
type QOpenGLTexture__TextureFormat int

var (
	QOpenGLTexture__NoFormat                       = QOpenGLTexture__TextureFormat(0)
	QOpenGLTexture__R8_UNorm                       = QOpenGLTexture__TextureFormat(0x8229)
	QOpenGLTexture__RG8_UNorm                      = QOpenGLTexture__TextureFormat(0x822B)
	QOpenGLTexture__RGB8_UNorm                     = QOpenGLTexture__TextureFormat(0x8051)
	QOpenGLTexture__RGBA8_UNorm                    = QOpenGLTexture__TextureFormat(0x8058)
	QOpenGLTexture__R16_UNorm                      = QOpenGLTexture__TextureFormat(0x822A)
	QOpenGLTexture__RG16_UNorm                     = QOpenGLTexture__TextureFormat(0x822C)
	QOpenGLTexture__RGB16_UNorm                    = QOpenGLTexture__TextureFormat(0x8054)
	QOpenGLTexture__RGBA16_UNorm                   = QOpenGLTexture__TextureFormat(0x805B)
	QOpenGLTexture__R8_SNorm                       = QOpenGLTexture__TextureFormat(0x8F94)
	QOpenGLTexture__RG8_SNorm                      = QOpenGLTexture__TextureFormat(0x8F95)
	QOpenGLTexture__RGB8_SNorm                     = QOpenGLTexture__TextureFormat(0x8F96)
	QOpenGLTexture__RGBA8_SNorm                    = QOpenGLTexture__TextureFormat(0x8F97)
	QOpenGLTexture__R16_SNorm                      = QOpenGLTexture__TextureFormat(0x8F98)
	QOpenGLTexture__RG16_SNorm                     = QOpenGLTexture__TextureFormat(0x8F99)
	QOpenGLTexture__RGB16_SNorm                    = QOpenGLTexture__TextureFormat(0x8F9A)
	QOpenGLTexture__RGBA16_SNorm                   = QOpenGLTexture__TextureFormat(0x8F9B)
	QOpenGLTexture__R8U                            = QOpenGLTexture__TextureFormat(0x8232)
	QOpenGLTexture__RG8U                           = QOpenGLTexture__TextureFormat(0x8238)
	QOpenGLTexture__RGB8U                          = QOpenGLTexture__TextureFormat(0x8D7D)
	QOpenGLTexture__RGBA8U                         = QOpenGLTexture__TextureFormat(0x8D7C)
	QOpenGLTexture__R16U                           = QOpenGLTexture__TextureFormat(0x8234)
	QOpenGLTexture__RG16U                          = QOpenGLTexture__TextureFormat(0x823A)
	QOpenGLTexture__RGB16U                         = QOpenGLTexture__TextureFormat(0x8D77)
	QOpenGLTexture__RGBA16U                        = QOpenGLTexture__TextureFormat(0x8D76)
	QOpenGLTexture__R32U                           = QOpenGLTexture__TextureFormat(0x8236)
	QOpenGLTexture__RG32U                          = QOpenGLTexture__TextureFormat(0x823C)
	QOpenGLTexture__RGB32U                         = QOpenGLTexture__TextureFormat(0x8D71)
	QOpenGLTexture__RGBA32U                        = QOpenGLTexture__TextureFormat(0x8D70)
	QOpenGLTexture__R8I                            = QOpenGLTexture__TextureFormat(0x8231)
	QOpenGLTexture__RG8I                           = QOpenGLTexture__TextureFormat(0x8237)
	QOpenGLTexture__RGB8I                          = QOpenGLTexture__TextureFormat(0x8D8F)
	QOpenGLTexture__RGBA8I                         = QOpenGLTexture__TextureFormat(0x8D8E)
	QOpenGLTexture__R16I                           = QOpenGLTexture__TextureFormat(0x8233)
	QOpenGLTexture__RG16I                          = QOpenGLTexture__TextureFormat(0x8239)
	QOpenGLTexture__RGB16I                         = QOpenGLTexture__TextureFormat(0x8D89)
	QOpenGLTexture__RGBA16I                        = QOpenGLTexture__TextureFormat(0x8D88)
	QOpenGLTexture__R32I                           = QOpenGLTexture__TextureFormat(0x8235)
	QOpenGLTexture__RG32I                          = QOpenGLTexture__TextureFormat(0x823B)
	QOpenGLTexture__RGB32I                         = QOpenGLTexture__TextureFormat(0x8D83)
	QOpenGLTexture__RGBA32I                        = QOpenGLTexture__TextureFormat(0x8D82)
	QOpenGLTexture__R16F                           = QOpenGLTexture__TextureFormat(0x822D)
	QOpenGLTexture__RG16F                          = QOpenGLTexture__TextureFormat(0x822F)
	QOpenGLTexture__RGB16F                         = QOpenGLTexture__TextureFormat(0x881B)
	QOpenGLTexture__RGBA16F                        = QOpenGLTexture__TextureFormat(0x881A)
	QOpenGLTexture__R32F                           = QOpenGLTexture__TextureFormat(0x822E)
	QOpenGLTexture__RG32F                          = QOpenGLTexture__TextureFormat(0x8230)
	QOpenGLTexture__RGB32F                         = QOpenGLTexture__TextureFormat(0x8815)
	QOpenGLTexture__RGBA32F                        = QOpenGLTexture__TextureFormat(0x8814)
	QOpenGLTexture__RGB9E5                         = QOpenGLTexture__TextureFormat(0x8C3D)
	QOpenGLTexture__RG11B10F                       = QOpenGLTexture__TextureFormat(0x8C3A)
	QOpenGLTexture__RG3B2                          = QOpenGLTexture__TextureFormat(0x2A10)
	QOpenGLTexture__R5G6B5                         = QOpenGLTexture__TextureFormat(0x8D62)
	QOpenGLTexture__RGB5A1                         = QOpenGLTexture__TextureFormat(0x8057)
	QOpenGLTexture__RGBA4                          = QOpenGLTexture__TextureFormat(0x8056)
	QOpenGLTexture__RGB10A2                        = QOpenGLTexture__TextureFormat(0x906F)
	QOpenGLTexture__D16                            = QOpenGLTexture__TextureFormat(0x81A5)
	QOpenGLTexture__D24                            = QOpenGLTexture__TextureFormat(0x81A6)
	QOpenGLTexture__D24S8                          = QOpenGLTexture__TextureFormat(0x88F0)
	QOpenGLTexture__D32                            = QOpenGLTexture__TextureFormat(0x81A7)
	QOpenGLTexture__D32F                           = QOpenGLTexture__TextureFormat(0x8CAC)
	QOpenGLTexture__D32FS8X24                      = QOpenGLTexture__TextureFormat(0x8CAD)
	QOpenGLTexture__S8                             = QOpenGLTexture__TextureFormat(0x8D48)
	QOpenGLTexture__RGB_DXT1                       = QOpenGLTexture__TextureFormat(0x83F0)
	QOpenGLTexture__RGBA_DXT1                      = QOpenGLTexture__TextureFormat(0x83F1)
	QOpenGLTexture__RGBA_DXT3                      = QOpenGLTexture__TextureFormat(0x83F2)
	QOpenGLTexture__RGBA_DXT5                      = QOpenGLTexture__TextureFormat(0x83F3)
	QOpenGLTexture__R_ATI1N_UNorm                  = QOpenGLTexture__TextureFormat(0x8DBB)
	QOpenGLTexture__R_ATI1N_SNorm                  = QOpenGLTexture__TextureFormat(0x8DBC)
	QOpenGLTexture__RG_ATI2N_UNorm                 = QOpenGLTexture__TextureFormat(0x8DBD)
	QOpenGLTexture__RG_ATI2N_SNorm                 = QOpenGLTexture__TextureFormat(0x8DBE)
	QOpenGLTexture__RGB_BP_UNSIGNED_FLOAT          = QOpenGLTexture__TextureFormat(0x8E8F)
	QOpenGLTexture__RGB_BP_SIGNED_FLOAT            = QOpenGLTexture__TextureFormat(0x8E8E)
	QOpenGLTexture__RGB_BP_UNorm                   = QOpenGLTexture__TextureFormat(0x8E8C)
	QOpenGLTexture__R11_EAC_UNorm                  = QOpenGLTexture__TextureFormat(0x9270)
	QOpenGLTexture__R11_EAC_SNorm                  = QOpenGLTexture__TextureFormat(0x9271)
	QOpenGLTexture__RG11_EAC_UNorm                 = QOpenGLTexture__TextureFormat(0x9272)
	QOpenGLTexture__RG11_EAC_SNorm                 = QOpenGLTexture__TextureFormat(0x9273)
	QOpenGLTexture__RGB8_ETC2                      = QOpenGLTexture__TextureFormat(0x9274)
	QOpenGLTexture__SRGB8_ETC2                     = QOpenGLTexture__TextureFormat(0x9275)
	QOpenGLTexture__RGB8_PunchThrough_Alpha1_ETC2  = QOpenGLTexture__TextureFormat(0x9276)
	QOpenGLTexture__SRGB8_PunchThrough_Alpha1_ETC2 = QOpenGLTexture__TextureFormat(0x9277)
	QOpenGLTexture__RGBA8_ETC2_EAC                 = QOpenGLTexture__TextureFormat(0x9278)
	QOpenGLTexture__SRGB8_Alpha8_ETC2_EAC          = QOpenGLTexture__TextureFormat(0x9279)
	QOpenGLTexture__SRGB8                          = QOpenGLTexture__TextureFormat(0x8C41)
	QOpenGLTexture__SRGB8_Alpha8                   = QOpenGLTexture__TextureFormat(0x8C43)
	QOpenGLTexture__SRGB_DXT1                      = QOpenGLTexture__TextureFormat(0x8C4C)
	QOpenGLTexture__SRGB_Alpha_DXT1                = QOpenGLTexture__TextureFormat(0x8C4D)
	QOpenGLTexture__SRGB_Alpha_DXT3                = QOpenGLTexture__TextureFormat(0x8C4E)
	QOpenGLTexture__SRGB_Alpha_DXT5                = QOpenGLTexture__TextureFormat(0x8C4F)
	QOpenGLTexture__SRGB_BP_UNorm                  = QOpenGLTexture__TextureFormat(0x8E8D)
	QOpenGLTexture__DepthFormat                    = QOpenGLTexture__TextureFormat(0x1902)
	QOpenGLTexture__AlphaFormat                    = QOpenGLTexture__TextureFormat(0x1906)
	QOpenGLTexture__RGBFormat                      = QOpenGLTexture__TextureFormat(0x1907)
	QOpenGLTexture__RGBAFormat                     = QOpenGLTexture__TextureFormat(0x1908)
	QOpenGLTexture__LuminanceFormat                = QOpenGLTexture__TextureFormat(0x1909)
	QOpenGLTexture__LuminanceAlphaFormat           = QOpenGLTexture__TextureFormat(0x190A)
)

//QOpenGLTexture::TextureUnitReset
type QOpenGLTexture__TextureUnitReset int

var (
	QOpenGLTexture__ResetTextureUnit     = QOpenGLTexture__TextureUnitReset(0)
	QOpenGLTexture__DontResetTextureUnit = QOpenGLTexture__TextureUnitReset(1)
)

//QOpenGLTexture::WrapMode
type QOpenGLTexture__WrapMode int

var (
	QOpenGLTexture__Repeat         = QOpenGLTexture__WrapMode(0x2901)
	QOpenGLTexture__MirroredRepeat = QOpenGLTexture__WrapMode(0x8370)
	QOpenGLTexture__ClampToEdge    = QOpenGLTexture__WrapMode(0x812F)
	QOpenGLTexture__ClampToBorder  = QOpenGLTexture__WrapMode(0x812D)
)

func (ptr *QOpenGLTexture) SetComparisonFunction(function QOpenGLTexture__ComparisonFunction) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetComparisonFunction(C.QtObjectPtr(ptr.Pointer()), C.int(function))
	}
}

func (ptr *QOpenGLTexture) DestroyQOpenGLTexture() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_DestroyQOpenGLTexture(C.QtObjectPtr(ptr.Pointer()))
	}
}

func NewQOpenGLTexture(target QOpenGLTexture__Target) *QOpenGLTexture {
	return QOpenGLTextureFromPointer(unsafe.Pointer(C.QOpenGLTexture_NewQOpenGLTexture(C.int(target))))
}

func NewQOpenGLTexture2(image QImageITF, genMipMaps QOpenGLTexture__MipMapGeneration) *QOpenGLTexture {
	return QOpenGLTextureFromPointer(unsafe.Pointer(C.QOpenGLTexture_NewQOpenGLTexture2(C.QtObjectPtr(PointerFromQImage(image)), C.int(genMipMaps))))
}

func (ptr *QOpenGLTexture) AllocateStorage() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_AllocateStorage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTexture) AllocateStorage2(pixelFormat QOpenGLTexture__PixelFormat, pixelType QOpenGLTexture__PixelType) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_AllocateStorage2(C.QtObjectPtr(ptr.Pointer()), C.int(pixelFormat), C.int(pixelType))
	}
}

func (ptr *QOpenGLTexture) Bind() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_Bind(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTexture) BorderColor3(border int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_BorderColor3(C.QtObjectPtr(ptr.Pointer()), C.int(border))
	}
}

func (ptr *QOpenGLTexture) ComparisonFunction() QOpenGLTexture__ComparisonFunction {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__ComparisonFunction(C.QOpenGLTexture_ComparisonFunction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) ComparisonMode() QOpenGLTexture__ComparisonMode {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__ComparisonMode(C.QOpenGLTexture_ComparisonMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) CreateTextureView(target QOpenGLTexture__Target, viewFormat QOpenGLTexture__TextureFormat, minimumMipmapLevel int, maximumMipmapLevel int, minimumLayer int, maximumLayer int) *QOpenGLTexture {
	if ptr.Pointer() != nil {
		return QOpenGLTextureFromPointer(unsafe.Pointer(C.QOpenGLTexture_CreateTextureView(C.QtObjectPtr(ptr.Pointer()), C.int(target), C.int(viewFormat), C.int(minimumMipmapLevel), C.int(maximumMipmapLevel), C.int(minimumLayer), C.int(maximumLayer))))
	}
	return nil
}

func (ptr *QOpenGLTexture) Depth() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Depth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) DepthStencilMode() QOpenGLTexture__DepthStencilMode {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__DepthStencilMode(C.QOpenGLTexture_DepthStencilMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) Destroy() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTexture) Faces() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Faces(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) Format() QOpenGLTexture__TextureFormat {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__TextureFormat(C.QOpenGLTexture_Format(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) GenerateMipMaps() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_GenerateMipMaps(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTexture) GenerateMipMaps2(baseLevel int, resetBaseLevel bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_GenerateMipMaps2(C.QtObjectPtr(ptr.Pointer()), C.int(baseLevel), C.int(qt.GoBoolToInt(resetBaseLevel)))
	}
}

func QOpenGLTexture_HasFeature(feature QOpenGLTexture__Feature) bool {
	return C.QOpenGLTexture_QOpenGLTexture_HasFeature(C.int(feature)) != 0
}

func (ptr *QOpenGLTexture) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) IsAutoMipMapGenerationEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsAutoMipMapGenerationEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) IsBound() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsBound(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) IsCreated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsCreated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) IsFixedSamplePositions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsFixedSamplePositions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) IsStorageAllocated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsStorageAllocated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) IsTextureView() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTexture_IsTextureView(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTexture) Layers() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Layers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MagnificationFilter() QOpenGLTexture__Filter {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__Filter(C.QOpenGLTexture_MagnificationFilter(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MaximumMipLevels() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_MaximumMipLevels(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MinificationFilter() QOpenGLTexture__Filter {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__Filter(C.QOpenGLTexture_MinificationFilter(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MipBaseLevel() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_MipBaseLevel(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MipLevels() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_MipLevels(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) MipMaxLevel() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_MipMaxLevel(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) Release() {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTexture) Samples() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Samples(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) SetAutoMipMapGenerationEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetAutoMipMapGenerationEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QOpenGLTexture) SetBorderColor(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetBorderColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QOpenGLTexture) SetBorderColor3(r int, g int, b int, a int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetBorderColor3(C.QtObjectPtr(ptr.Pointer()), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QOpenGLTexture) SetComparisonMode(mode QOpenGLTexture__ComparisonMode) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetComparisonMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QOpenGLTexture) SetData9(image QImageITF, genMipMaps QOpenGLTexture__MipMapGeneration) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetData9(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQImage(image)), C.int(genMipMaps))
	}
}

func (ptr *QOpenGLTexture) SetDepthStencilMode(mode QOpenGLTexture__DepthStencilMode) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetDepthStencilMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QOpenGLTexture) SetFixedSamplePositions(fixed bool) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetFixedSamplePositions(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(fixed)))
	}
}

func (ptr *QOpenGLTexture) SetFormat(format QOpenGLTexture__TextureFormat) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetFormat(C.QtObjectPtr(ptr.Pointer()), C.int(format))
	}
}

func (ptr *QOpenGLTexture) SetLayers(layers int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetLayers(C.QtObjectPtr(ptr.Pointer()), C.int(layers))
	}
}

func (ptr *QOpenGLTexture) SetMagnificationFilter(filter QOpenGLTexture__Filter) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMagnificationFilter(C.QtObjectPtr(ptr.Pointer()), C.int(filter))
	}
}

func (ptr *QOpenGLTexture) SetMinMagFilters(minificationFilter QOpenGLTexture__Filter, magnificationFilter QOpenGLTexture__Filter) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMinMagFilters(C.QtObjectPtr(ptr.Pointer()), C.int(minificationFilter), C.int(magnificationFilter))
	}
}

func (ptr *QOpenGLTexture) SetMinificationFilter(filter QOpenGLTexture__Filter) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMinificationFilter(C.QtObjectPtr(ptr.Pointer()), C.int(filter))
	}
}

func (ptr *QOpenGLTexture) SetMipBaseLevel(baseLevel int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMipBaseLevel(C.QtObjectPtr(ptr.Pointer()), C.int(baseLevel))
	}
}

func (ptr *QOpenGLTexture) SetMipLevelRange(baseLevel int, maxLevel int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMipLevelRange(C.QtObjectPtr(ptr.Pointer()), C.int(baseLevel), C.int(maxLevel))
	}
}

func (ptr *QOpenGLTexture) SetMipLevels(levels int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMipLevels(C.QtObjectPtr(ptr.Pointer()), C.int(levels))
	}
}

func (ptr *QOpenGLTexture) SetMipMaxLevel(maxLevel int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetMipMaxLevel(C.QtObjectPtr(ptr.Pointer()), C.int(maxLevel))
	}
}

func (ptr *QOpenGLTexture) SetSamples(samples int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetSamples(C.QtObjectPtr(ptr.Pointer()), C.int(samples))
	}
}

func (ptr *QOpenGLTexture) SetSize(width int, height int, depth int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetSize(C.QtObjectPtr(ptr.Pointer()), C.int(width), C.int(height), C.int(depth))
	}
}

func (ptr *QOpenGLTexture) SetSwizzleMask(component QOpenGLTexture__SwizzleComponent, value QOpenGLTexture__SwizzleValue) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetSwizzleMask(C.QtObjectPtr(ptr.Pointer()), C.int(component), C.int(value))
	}
}

func (ptr *QOpenGLTexture) SetSwizzleMask2(r QOpenGLTexture__SwizzleValue, g QOpenGLTexture__SwizzleValue, b QOpenGLTexture__SwizzleValue, a QOpenGLTexture__SwizzleValue) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetSwizzleMask2(C.QtObjectPtr(ptr.Pointer()), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QOpenGLTexture) SetWrapMode2(direction QOpenGLTexture__CoordinateDirection, mode QOpenGLTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetWrapMode2(C.QtObjectPtr(ptr.Pointer()), C.int(direction), C.int(mode))
	}
}

func (ptr *QOpenGLTexture) SetWrapMode(mode QOpenGLTexture__WrapMode) {
	if ptr.Pointer() != nil {
		C.QOpenGLTexture_SetWrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QOpenGLTexture) SwizzleMask(component QOpenGLTexture__SwizzleComponent) QOpenGLTexture__SwizzleValue {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__SwizzleValue(C.QOpenGLTexture_SwizzleMask(C.QtObjectPtr(ptr.Pointer()), C.int(component)))
	}
	return 0
}

func (ptr *QOpenGLTexture) Target() QOpenGLTexture__Target {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__Target(C.QOpenGLTexture_Target(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTexture_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTexture) WrapMode(direction QOpenGLTexture__CoordinateDirection) QOpenGLTexture__WrapMode {
	if ptr.Pointer() != nil {
		return QOpenGLTexture__WrapMode(C.QOpenGLTexture_WrapMode(C.QtObjectPtr(ptr.Pointer()), C.int(direction)))
	}
	return 0
}
