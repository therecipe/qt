#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsTextItem_OpenExternalLinks(void* ptr);
void QGraphicsTextItem_SetOpenExternalLinks(void* ptr, int open);
void QGraphicsTextItem_SetTextCursor(void* ptr, void* cursor);
void QGraphicsTextItem_AdjustSize(void* ptr);
int QGraphicsTextItem_Contains(void* ptr, void* point);
void* QGraphicsTextItem_DefaultTextColor(void* ptr);
void* QGraphicsTextItem_Document(void* ptr);
int QGraphicsTextItem_IsObscuredBy(void* ptr, void* item);
void QGraphicsTextItem_ConnectLinkActivated(void* ptr);
void QGraphicsTextItem_DisconnectLinkActivated(void* ptr);
void QGraphicsTextItem_ConnectLinkHovered(void* ptr);
void QGraphicsTextItem_DisconnectLinkHovered(void* ptr);
void QGraphicsTextItem_Paint(void* ptr, void* painter, void* option, void* widget);
void QGraphicsTextItem_SetDefaultTextColor(void* ptr, void* col);
void QGraphicsTextItem_SetDocument(void* ptr, void* document);
void QGraphicsTextItem_SetFont(void* ptr, void* font);
void QGraphicsTextItem_SetHtml(void* ptr, char* text);
void QGraphicsTextItem_SetPlainText(void* ptr, char* text);
void QGraphicsTextItem_SetTabChangesFocus(void* ptr, int b);
void QGraphicsTextItem_SetTextInteractionFlags(void* ptr, int flags);
void QGraphicsTextItem_SetTextWidth(void* ptr, double width);
int QGraphicsTextItem_TabChangesFocus(void* ptr);
int QGraphicsTextItem_TextInteractionFlags(void* ptr);
double QGraphicsTextItem_TextWidth(void* ptr);
char* QGraphicsTextItem_ToHtml(void* ptr);
char* QGraphicsTextItem_ToPlainText(void* ptr);
int QGraphicsTextItem_Type(void* ptr);
void QGraphicsTextItem_DestroyQGraphicsTextItem(void* ptr);

#ifdef __cplusplus
}
#endif