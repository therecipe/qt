#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QOpenGLTexture_SetComparisonFunction(QtObjectPtr ptr, int function);
void QOpenGLTexture_DestroyQOpenGLTexture(QtObjectPtr ptr);
QtObjectPtr QOpenGLTexture_NewQOpenGLTexture(int target);
QtObjectPtr QOpenGLTexture_NewQOpenGLTexture2(QtObjectPtr image, int genMipMaps);
void QOpenGLTexture_AllocateStorage(QtObjectPtr ptr);
void QOpenGLTexture_AllocateStorage2(QtObjectPtr ptr, int pixelFormat, int pixelType);
void QOpenGLTexture_Bind(QtObjectPtr ptr);
void QOpenGLTexture_BorderColor3(QtObjectPtr ptr, int border);
int QOpenGLTexture_ComparisonFunction(QtObjectPtr ptr);
int QOpenGLTexture_ComparisonMode(QtObjectPtr ptr);
int QOpenGLTexture_Create(QtObjectPtr ptr);
QtObjectPtr QOpenGLTexture_CreateTextureView(QtObjectPtr ptr, int target, int viewFormat, int minimumMipmapLevel, int maximumMipmapLevel, int minimumLayer, int maximumLayer);
int QOpenGLTexture_Depth(QtObjectPtr ptr);
int QOpenGLTexture_DepthStencilMode(QtObjectPtr ptr);
void QOpenGLTexture_Destroy(QtObjectPtr ptr);
int QOpenGLTexture_Faces(QtObjectPtr ptr);
int QOpenGLTexture_Format(QtObjectPtr ptr);
void QOpenGLTexture_GenerateMipMaps(QtObjectPtr ptr);
void QOpenGLTexture_GenerateMipMaps2(QtObjectPtr ptr, int baseLevel, int resetBaseLevel);
int QOpenGLTexture_QOpenGLTexture_HasFeature(int feature);
int QOpenGLTexture_Height(QtObjectPtr ptr);
int QOpenGLTexture_IsAutoMipMapGenerationEnabled(QtObjectPtr ptr);
int QOpenGLTexture_IsBound(QtObjectPtr ptr);
int QOpenGLTexture_IsCreated(QtObjectPtr ptr);
int QOpenGLTexture_IsFixedSamplePositions(QtObjectPtr ptr);
int QOpenGLTexture_IsStorageAllocated(QtObjectPtr ptr);
int QOpenGLTexture_IsTextureView(QtObjectPtr ptr);
int QOpenGLTexture_Layers(QtObjectPtr ptr);
int QOpenGLTexture_MagnificationFilter(QtObjectPtr ptr);
int QOpenGLTexture_MaximumMipLevels(QtObjectPtr ptr);
int QOpenGLTexture_MinificationFilter(QtObjectPtr ptr);
int QOpenGLTexture_MipBaseLevel(QtObjectPtr ptr);
int QOpenGLTexture_MipLevels(QtObjectPtr ptr);
int QOpenGLTexture_MipMaxLevel(QtObjectPtr ptr);
void QOpenGLTexture_Release(QtObjectPtr ptr);
int QOpenGLTexture_Samples(QtObjectPtr ptr);
void QOpenGLTexture_SetAutoMipMapGenerationEnabled(QtObjectPtr ptr, int enabled);
void QOpenGLTexture_SetBorderColor(QtObjectPtr ptr, QtObjectPtr color);
void QOpenGLTexture_SetBorderColor3(QtObjectPtr ptr, int r, int g, int b, int a);
void QOpenGLTexture_SetComparisonMode(QtObjectPtr ptr, int mode);
void QOpenGLTexture_SetData9(QtObjectPtr ptr, QtObjectPtr image, int genMipMaps);
void QOpenGLTexture_SetDepthStencilMode(QtObjectPtr ptr, int mode);
void QOpenGLTexture_SetFixedSamplePositions(QtObjectPtr ptr, int fixed);
void QOpenGLTexture_SetFormat(QtObjectPtr ptr, int format);
void QOpenGLTexture_SetLayers(QtObjectPtr ptr, int layers);
void QOpenGLTexture_SetMagnificationFilter(QtObjectPtr ptr, int filter);
void QOpenGLTexture_SetMinMagFilters(QtObjectPtr ptr, int minificationFilter, int magnificationFilter);
void QOpenGLTexture_SetMinificationFilter(QtObjectPtr ptr, int filter);
void QOpenGLTexture_SetMipBaseLevel(QtObjectPtr ptr, int baseLevel);
void QOpenGLTexture_SetMipLevelRange(QtObjectPtr ptr, int baseLevel, int maxLevel);
void QOpenGLTexture_SetMipLevels(QtObjectPtr ptr, int levels);
void QOpenGLTexture_SetMipMaxLevel(QtObjectPtr ptr, int maxLevel);
void QOpenGLTexture_SetSamples(QtObjectPtr ptr, int samples);
void QOpenGLTexture_SetSize(QtObjectPtr ptr, int width, int height, int depth);
void QOpenGLTexture_SetSwizzleMask(QtObjectPtr ptr, int component, int value);
void QOpenGLTexture_SetSwizzleMask2(QtObjectPtr ptr, int r, int g, int b, int a);
void QOpenGLTexture_SetWrapMode2(QtObjectPtr ptr, int direction, int mode);
void QOpenGLTexture_SetWrapMode(QtObjectPtr ptr, int mode);
int QOpenGLTexture_SwizzleMask(QtObjectPtr ptr, int component);
int QOpenGLTexture_Target(QtObjectPtr ptr);
int QOpenGLTexture_Width(QtObjectPtr ptr);
int QOpenGLTexture_WrapMode(QtObjectPtr ptr, int direction);

#ifdef __cplusplus
}
#endif