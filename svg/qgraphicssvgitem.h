#ifdef __cplusplus
extern "C" {
#endif

char* QGraphicsSvgItem_ElementId(void* ptr);
void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget);
void* QGraphicsSvgItem_Renderer(void* ptr);
void QGraphicsSvgItem_SetElementId(void* ptr, char* id);
void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size);
void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer);
int QGraphicsSvgItem_Type(void* ptr);

#ifdef __cplusplus
}
#endif