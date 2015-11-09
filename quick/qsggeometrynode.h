#ifdef __cplusplus
extern "C" {
#endif

void* QSGGeometryNode_NewQSGGeometryNode();
void* QSGGeometryNode_Material(void* ptr);
void* QSGGeometryNode_OpaqueMaterial(void* ptr);
void QSGGeometryNode_SetMaterial(void* ptr, void* material);
void QSGGeometryNode_SetOpaqueMaterial(void* ptr, void* material);
void QSGGeometryNode_DestroyQSGGeometryNode(void* ptr);

#ifdef __cplusplus
}
#endif