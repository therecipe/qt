#ifdef __cplusplus
extern "C" {
#endif

void QSGTexture_Bind(void* ptr);
int QSGTexture_Filtering(void* ptr);
int QSGTexture_HasAlphaChannel(void* ptr);
int QSGTexture_HasMipmaps(void* ptr);
int QSGTexture_HorizontalWrapMode(void* ptr);
int QSGTexture_IsAtlasTexture(void* ptr);
int QSGTexture_MipmapFiltering(void* ptr);
void* QSGTexture_RemovedFromAtlas(void* ptr);
void QSGTexture_SetFiltering(void* ptr, int filter);
void QSGTexture_SetHorizontalWrapMode(void* ptr, int hwrap);
void QSGTexture_SetMipmapFiltering(void* ptr, int filter);
void QSGTexture_SetVerticalWrapMode(void* ptr, int vwrap);
int QSGTexture_TextureId(void* ptr);
void QSGTexture_UpdateBindOptions(void* ptr, int force);
int QSGTexture_VerticalWrapMode(void* ptr);
void QSGTexture_DestroyQSGTexture(void* ptr);

#ifdef __cplusplus
}
#endif