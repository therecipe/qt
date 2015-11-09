#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsProxyWidget_NewQGraphicsProxyWidget(void* parent, int wFlags);
void* QGraphicsProxyWidget_CreateProxyForChildWidget(void* ptr, void* child);
void QGraphicsProxyWidget_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsProxyWidget_SetGeometry(void* ptr, void* rect);
void QGraphicsProxyWidget_SetWidget(void* ptr, void* widget);
int QGraphicsProxyWidget_Type(void* ptr);
void* QGraphicsProxyWidget_Widget(void* ptr);
void QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(void* ptr);

#ifdef __cplusplus
}
#endif