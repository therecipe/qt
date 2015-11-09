#ifdef __cplusplus
extern "C" {
#endif

void* QSGSimpleTextureNode_NewQSGSimpleTextureNode();
int QSGSimpleTextureNode_Filtering(void* ptr);
int QSGSimpleTextureNode_OwnsTexture(void* ptr);
void QSGSimpleTextureNode_SetFiltering(void* ptr, int filtering);
void QSGSimpleTextureNode_SetOwnsTexture(void* ptr, int owns);
void QSGSimpleTextureNode_SetRect(void* ptr, void* r);
void QSGSimpleTextureNode_SetRect2(void* ptr, double x, double y, double w, double h);
void QSGSimpleTextureNode_SetSourceRect(void* ptr, void* r);
void QSGSimpleTextureNode_SetSourceRect2(void* ptr, double x, double y, double w, double h);
void QSGSimpleTextureNode_SetTexture(void* ptr, void* texture);
void QSGSimpleTextureNode_SetTextureCoordinatesTransform(void* ptr, int mode);
void* QSGSimpleTextureNode_Texture(void* ptr);
int QSGSimpleTextureNode_TextureCoordinatesTransform(void* ptr);
void QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(void* ptr);

#ifdef __cplusplus
}
#endif