#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsItemGroup_NewQGraphicsItemGroup(void* parent);
void QGraphicsItemGroup_AddToGroup(void* ptr, void* item);
int QGraphicsItemGroup_IsObscuredBy(void* ptr, void* item);
void QGraphicsItemGroup_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsItemGroup_RemoveFromGroup(void* ptr, void* item);
int QGraphicsItemGroup_Type(void* ptr);
void QGraphicsItemGroup_DestroyQGraphicsItemGroup(void* ptr);

#ifdef __cplusplus
}
#endif