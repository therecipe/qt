#ifdef __cplusplus
extern "C" {
#endif

int QSGOpaqueTextureMaterial_Filtering(void* ptr);
int QSGOpaqueTextureMaterial_HorizontalWrapMode(void* ptr);
int QSGOpaqueTextureMaterial_MipmapFiltering(void* ptr);
void QSGOpaqueTextureMaterial_SetFiltering(void* ptr, int filtering);
void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(void* ptr, int mode);
void QSGOpaqueTextureMaterial_SetMipmapFiltering(void* ptr, int filtering);
void QSGOpaqueTextureMaterial_SetTexture(void* ptr, void* texture);
void QSGOpaqueTextureMaterial_SetVerticalWrapMode(void* ptr, int mode);
void* QSGOpaqueTextureMaterial_Texture(void* ptr);
int QSGOpaqueTextureMaterial_VerticalWrapMode(void* ptr);

#ifdef __cplusplus
}
#endif