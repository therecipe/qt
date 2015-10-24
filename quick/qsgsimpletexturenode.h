#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSGSimpleTextureNode_NewQSGSimpleTextureNode();
int QSGSimpleTextureNode_Filtering(QtObjectPtr ptr);
int QSGSimpleTextureNode_OwnsTexture(QtObjectPtr ptr);
void QSGSimpleTextureNode_SetFiltering(QtObjectPtr ptr, int filtering);
void QSGSimpleTextureNode_SetOwnsTexture(QtObjectPtr ptr, int owns);
void QSGSimpleTextureNode_SetRect(QtObjectPtr ptr, QtObjectPtr r);
void QSGSimpleTextureNode_SetSourceRect(QtObjectPtr ptr, QtObjectPtr r);
void QSGSimpleTextureNode_SetTexture(QtObjectPtr ptr, QtObjectPtr texture);
void QSGSimpleTextureNode_SetTextureCoordinatesTransform(QtObjectPtr ptr, int mode);
QtObjectPtr QSGSimpleTextureNode_Texture(QtObjectPtr ptr);
int QSGSimpleTextureNode_TextureCoordinatesTransform(QtObjectPtr ptr);
void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif