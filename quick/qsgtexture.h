#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSGTexture_Bind(QtObjectPtr ptr);
int QSGTexture_Filtering(QtObjectPtr ptr);
int QSGTexture_HasAlphaChannel(QtObjectPtr ptr);
int QSGTexture_HasMipmaps(QtObjectPtr ptr);
int QSGTexture_HorizontalWrapMode(QtObjectPtr ptr);
int QSGTexture_IsAtlasTexture(QtObjectPtr ptr);
int QSGTexture_MipmapFiltering(QtObjectPtr ptr);
QtObjectPtr QSGTexture_RemovedFromAtlas(QtObjectPtr ptr);
void QSGTexture_SetFiltering(QtObjectPtr ptr, int filter);
void QSGTexture_SetHorizontalWrapMode(QtObjectPtr ptr, int hwrap);
void QSGTexture_SetMipmapFiltering(QtObjectPtr ptr, int filter);
void QSGTexture_SetVerticalWrapMode(QtObjectPtr ptr, int vwrap);
int QSGTexture_TextureId(QtObjectPtr ptr);
void QSGTexture_UpdateBindOptions(QtObjectPtr ptr, int force);
int QSGTexture_VerticalWrapMode(QtObjectPtr ptr);
void QSGTexture_DestroyQSGTexture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif