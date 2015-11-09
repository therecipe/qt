#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsSimpleTextItem_Contains(void* ptr, void* point);
int QGraphicsSimpleTextItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsSimpleTextItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsSimpleTextItem_SetFont(void* ptr, void* font);
void QGraphicsSimpleTextItem_SetText(void* ptr, char* text);
char* QGraphicsSimpleTextItem_Text(void* ptr);
int QGraphicsSimpleTextItem_Type(void* ptr);
void QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(void* ptr);

#ifdef __cplusplus
}
#endif