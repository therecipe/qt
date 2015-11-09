#ifdef __cplusplus
extern "C" {
#endif

void* QSGClipNode_NewQSGClipNode();
int QSGClipNode_IsRectangular(void* ptr);
void QSGClipNode_SetClipRect(void* ptr, void* rect);
void QSGClipNode_SetIsRectangular(void* ptr, int rectHint);
void QSGClipNode_DestroyQSGClipNode(void* ptr);

#ifdef __cplusplus
}
#endif