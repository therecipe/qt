#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsPathItem_Contains(void* ptr, void* point);
int QGraphicsPathItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsPathItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsPathItem_SetPath(void* ptr, void* path);
int QGraphicsPathItem_Type(void* ptr);
void QGraphicsPathItem_DestroyQGraphicsPathItem(void* ptr);

#ifdef __cplusplus
}
#endif