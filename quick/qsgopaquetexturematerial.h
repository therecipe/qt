#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSGOpaqueTextureMaterial_Filtering(QtObjectPtr ptr);
int QSGOpaqueTextureMaterial_HorizontalWrapMode(QtObjectPtr ptr);
int QSGOpaqueTextureMaterial_MipmapFiltering(QtObjectPtr ptr);
void QSGOpaqueTextureMaterial_SetFiltering(QtObjectPtr ptr, int filtering);
void QSGOpaqueTextureMaterial_SetHorizontalWrapMode(QtObjectPtr ptr, int mode);
void QSGOpaqueTextureMaterial_SetMipmapFiltering(QtObjectPtr ptr, int filtering);
void QSGOpaqueTextureMaterial_SetTexture(QtObjectPtr ptr, QtObjectPtr texture);
void QSGOpaqueTextureMaterial_SetVerticalWrapMode(QtObjectPtr ptr, int mode);
QtObjectPtr QSGOpaqueTextureMaterial_Texture(QtObjectPtr ptr);
int QSGOpaqueTextureMaterial_VerticalWrapMode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif