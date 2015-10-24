#include "qopengltexture.h"
#include <QImage>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QOpenGLTexture>
#include "_cgo_export.h"

class MyQOpenGLTexture: public QOpenGLTexture {
public:
};

void QOpenGLTexture_SetComparisonFunction(QtObjectPtr ptr, int function){
	static_cast<QOpenGLTexture*>(ptr)->setComparisonFunction(static_cast<QOpenGLTexture::ComparisonFunction>(function));
}

void QOpenGLTexture_DestroyQOpenGLTexture(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->~QOpenGLTexture();
}

QtObjectPtr QOpenGLTexture_NewQOpenGLTexture(int target){
	return new QOpenGLTexture(static_cast<QOpenGLTexture::Target>(target));
}

QtObjectPtr QOpenGLTexture_NewQOpenGLTexture2(QtObjectPtr image, int genMipMaps){
	return new QOpenGLTexture(*static_cast<QImage*>(image), static_cast<QOpenGLTexture::MipMapGeneration>(genMipMaps));
}

void QOpenGLTexture_AllocateStorage(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->allocateStorage();
}

void QOpenGLTexture_AllocateStorage2(QtObjectPtr ptr, int pixelFormat, int pixelType){
	static_cast<QOpenGLTexture*>(ptr)->allocateStorage(static_cast<QOpenGLTexture::PixelFormat>(pixelFormat), static_cast<QOpenGLTexture::PixelType>(pixelType));
}

void QOpenGLTexture_Bind(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->bind();
}

void QOpenGLTexture_BorderColor3(QtObjectPtr ptr, int border){
	static_cast<QOpenGLTexture*>(ptr)->borderColor(&border);
}

int QOpenGLTexture_ComparisonFunction(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->comparisonFunction();
}

int QOpenGLTexture_ComparisonMode(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->comparisonMode();
}

int QOpenGLTexture_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->create();
}

QtObjectPtr QOpenGLTexture_CreateTextureView(QtObjectPtr ptr, int target, int viewFormat, int minimumMipmapLevel, int maximumMipmapLevel, int minimumLayer, int maximumLayer){
	return static_cast<QOpenGLTexture*>(ptr)->createTextureView(static_cast<QOpenGLTexture::Target>(target), static_cast<QOpenGLTexture::TextureFormat>(viewFormat), minimumMipmapLevel, maximumMipmapLevel, minimumLayer, maximumLayer);
}

int QOpenGLTexture_Depth(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->depth();
}

int QOpenGLTexture_DepthStencilMode(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->depthStencilMode();
}

void QOpenGLTexture_Destroy(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->destroy();
}

int QOpenGLTexture_Faces(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->faces();
}

int QOpenGLTexture_Format(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->format();
}

void QOpenGLTexture_GenerateMipMaps(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->generateMipMaps();
}

void QOpenGLTexture_GenerateMipMaps2(QtObjectPtr ptr, int baseLevel, int resetBaseLevel){
	static_cast<QOpenGLTexture*>(ptr)->generateMipMaps(baseLevel, resetBaseLevel != 0);
}

int QOpenGLTexture_QOpenGLTexture_HasFeature(int feature){
	return QOpenGLTexture::hasFeature(static_cast<QOpenGLTexture::Feature>(feature));
}

int QOpenGLTexture_Height(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->height();
}

int QOpenGLTexture_IsAutoMipMapGenerationEnabled(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isAutoMipMapGenerationEnabled();
}

int QOpenGLTexture_IsBound(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isBound();
}

int QOpenGLTexture_IsCreated(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isCreated();
}

int QOpenGLTexture_IsFixedSamplePositions(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isFixedSamplePositions();
}

int QOpenGLTexture_IsStorageAllocated(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isStorageAllocated();
}

int QOpenGLTexture_IsTextureView(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->isTextureView();
}

int QOpenGLTexture_Layers(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->layers();
}

int QOpenGLTexture_MagnificationFilter(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->magnificationFilter();
}

int QOpenGLTexture_MaximumMipLevels(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->maximumMipLevels();
}

int QOpenGLTexture_MinificationFilter(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->minificationFilter();
}

int QOpenGLTexture_MipBaseLevel(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->mipBaseLevel();
}

int QOpenGLTexture_MipLevels(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->mipLevels();
}

int QOpenGLTexture_MipMaxLevel(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->mipMaxLevel();
}

void QOpenGLTexture_Release(QtObjectPtr ptr){
	static_cast<QOpenGLTexture*>(ptr)->release();
}

int QOpenGLTexture_Samples(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->samples();
}

void QOpenGLTexture_SetAutoMipMapGenerationEnabled(QtObjectPtr ptr, int enabled){
	static_cast<QOpenGLTexture*>(ptr)->setAutoMipMapGenerationEnabled(enabled != 0);
}

void QOpenGLTexture_SetBorderColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QOpenGLTexture*>(ptr)->setBorderColor(*static_cast<QColor*>(color));
}

void QOpenGLTexture_SetBorderColor3(QtObjectPtr ptr, int r, int g, int b, int a){
	static_cast<QOpenGLTexture*>(ptr)->setBorderColor(r, g, b, a);
}

void QOpenGLTexture_SetComparisonMode(QtObjectPtr ptr, int mode){
	static_cast<QOpenGLTexture*>(ptr)->setComparisonMode(static_cast<QOpenGLTexture::ComparisonMode>(mode));
}

void QOpenGLTexture_SetData9(QtObjectPtr ptr, QtObjectPtr image, int genMipMaps){
	static_cast<QOpenGLTexture*>(ptr)->setData(*static_cast<QImage*>(image), static_cast<QOpenGLTexture::MipMapGeneration>(genMipMaps));
}

void QOpenGLTexture_SetDepthStencilMode(QtObjectPtr ptr, int mode){
	static_cast<QOpenGLTexture*>(ptr)->setDepthStencilMode(static_cast<QOpenGLTexture::DepthStencilMode>(mode));
}

void QOpenGLTexture_SetFixedSamplePositions(QtObjectPtr ptr, int fixed){
	static_cast<QOpenGLTexture*>(ptr)->setFixedSamplePositions(fixed != 0);
}

void QOpenGLTexture_SetFormat(QtObjectPtr ptr, int format){
	static_cast<QOpenGLTexture*>(ptr)->setFormat(static_cast<QOpenGLTexture::TextureFormat>(format));
}

void QOpenGLTexture_SetLayers(QtObjectPtr ptr, int layers){
	static_cast<QOpenGLTexture*>(ptr)->setLayers(layers);
}

void QOpenGLTexture_SetMagnificationFilter(QtObjectPtr ptr, int filter){
	static_cast<QOpenGLTexture*>(ptr)->setMagnificationFilter(static_cast<QOpenGLTexture::Filter>(filter));
}

void QOpenGLTexture_SetMinMagFilters(QtObjectPtr ptr, int minificationFilter, int magnificationFilter){
	static_cast<QOpenGLTexture*>(ptr)->setMinMagFilters(static_cast<QOpenGLTexture::Filter>(minificationFilter), static_cast<QOpenGLTexture::Filter>(magnificationFilter));
}

void QOpenGLTexture_SetMinificationFilter(QtObjectPtr ptr, int filter){
	static_cast<QOpenGLTexture*>(ptr)->setMinificationFilter(static_cast<QOpenGLTexture::Filter>(filter));
}

void QOpenGLTexture_SetMipBaseLevel(QtObjectPtr ptr, int baseLevel){
	static_cast<QOpenGLTexture*>(ptr)->setMipBaseLevel(baseLevel);
}

void QOpenGLTexture_SetMipLevelRange(QtObjectPtr ptr, int baseLevel, int maxLevel){
	static_cast<QOpenGLTexture*>(ptr)->setMipLevelRange(baseLevel, maxLevel);
}

void QOpenGLTexture_SetMipLevels(QtObjectPtr ptr, int levels){
	static_cast<QOpenGLTexture*>(ptr)->setMipLevels(levels);
}

void QOpenGLTexture_SetMipMaxLevel(QtObjectPtr ptr, int maxLevel){
	static_cast<QOpenGLTexture*>(ptr)->setMipMaxLevel(maxLevel);
}

void QOpenGLTexture_SetSamples(QtObjectPtr ptr, int samples){
	static_cast<QOpenGLTexture*>(ptr)->setSamples(samples);
}

void QOpenGLTexture_SetSize(QtObjectPtr ptr, int width, int height, int depth){
	static_cast<QOpenGLTexture*>(ptr)->setSize(width, height, depth);
}

void QOpenGLTexture_SetSwizzleMask(QtObjectPtr ptr, int component, int value){
	static_cast<QOpenGLTexture*>(ptr)->setSwizzleMask(static_cast<QOpenGLTexture::SwizzleComponent>(component), static_cast<QOpenGLTexture::SwizzleValue>(value));
}

void QOpenGLTexture_SetSwizzleMask2(QtObjectPtr ptr, int r, int g, int b, int a){
	static_cast<QOpenGLTexture*>(ptr)->setSwizzleMask(static_cast<QOpenGLTexture::SwizzleValue>(r), static_cast<QOpenGLTexture::SwizzleValue>(g), static_cast<QOpenGLTexture::SwizzleValue>(b), static_cast<QOpenGLTexture::SwizzleValue>(a));
}

void QOpenGLTexture_SetWrapMode2(QtObjectPtr ptr, int direction, int mode){
	static_cast<QOpenGLTexture*>(ptr)->setWrapMode(static_cast<QOpenGLTexture::CoordinateDirection>(direction), static_cast<QOpenGLTexture::WrapMode>(mode));
}

void QOpenGLTexture_SetWrapMode(QtObjectPtr ptr, int mode){
	static_cast<QOpenGLTexture*>(ptr)->setWrapMode(static_cast<QOpenGLTexture::WrapMode>(mode));
}

int QOpenGLTexture_SwizzleMask(QtObjectPtr ptr, int component){
	return static_cast<QOpenGLTexture*>(ptr)->swizzleMask(static_cast<QOpenGLTexture::SwizzleComponent>(component));
}

int QOpenGLTexture_Target(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->target();
}

int QOpenGLTexture_Width(QtObjectPtr ptr){
	return static_cast<QOpenGLTexture*>(ptr)->width();
}

int QOpenGLTexture_WrapMode(QtObjectPtr ptr, int direction){
	return static_cast<QOpenGLTexture*>(ptr)->wrapMode(static_cast<QOpenGLTexture::CoordinateDirection>(direction));
}

