#ifdef __cplusplus
extern "C" {
#endif

void* QSGTransformNode_NewQSGTransformNode();
void QSGTransformNode_SetMatrix(void* ptr, void* matrix);
void QSGTransformNode_DestroyQSGTransformNode(void* ptr);

#ifdef __cplusplus
}
#endif