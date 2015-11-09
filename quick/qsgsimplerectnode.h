#ifdef __cplusplus
extern "C" {
#endif

void* QSGSimpleRectNode_NewQSGSimpleRectNode2();
void* QSGSimpleRectNode_NewQSGSimpleRectNode(void* rect, void* color);
void* QSGSimpleRectNode_Color(void* ptr);
void QSGSimpleRectNode_SetColor(void* ptr, void* color);
void QSGSimpleRectNode_SetRect(void* ptr, void* rect);
void QSGSimpleRectNode_SetRect2(void* ptr, double x, double y, double w, double h);

#ifdef __cplusplus
}
#endif