#ifdef __cplusplus
extern "C" {
#endif

void* QSGOpacityNode_NewQSGOpacityNode();
double QSGOpacityNode_Opacity(void* ptr);
void QSGOpacityNode_SetOpacity(void* ptr, double opacity);
void QSGOpacityNode_DestroyQSGOpacityNode(void* ptr);

#ifdef __cplusplus
}
#endif