#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsTextItem_OpenExternalLinks(QtObjectPtr ptr);
void QGraphicsTextItem_SetOpenExternalLinks(QtObjectPtr ptr, int open);
void QGraphicsTextItem_SetTextCursor(QtObjectPtr ptr, QtObjectPtr cursor);
QtObjectPtr QGraphicsTextItem_NewQGraphicsTextItem(QtObjectPtr parent);
QtObjectPtr QGraphicsTextItem_NewQGraphicsTextItem2(char* text, QtObjectPtr parent);
void QGraphicsTextItem_AdjustSize(QtObjectPtr ptr);
int QGraphicsTextItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QGraphicsTextItem_Document(QtObjectPtr ptr);
int QGraphicsTextItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsTextItem_ConnectLinkActivated(QtObjectPtr ptr);
void QGraphicsTextItem_DisconnectLinkActivated(QtObjectPtr ptr);
void QGraphicsTextItem_ConnectLinkHovered(QtObjectPtr ptr);
void QGraphicsTextItem_DisconnectLinkHovered(QtObjectPtr ptr);
void QGraphicsTextItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsTextItem_SetDefaultTextColor(QtObjectPtr ptr, QtObjectPtr col);
void QGraphicsTextItem_SetDocument(QtObjectPtr ptr, QtObjectPtr document);
void QGraphicsTextItem_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QGraphicsTextItem_SetHtml(QtObjectPtr ptr, char* text);
void QGraphicsTextItem_SetPlainText(QtObjectPtr ptr, char* text);
void QGraphicsTextItem_SetTabChangesFocus(QtObjectPtr ptr, int b);
void QGraphicsTextItem_SetTextInteractionFlags(QtObjectPtr ptr, int flags);
int QGraphicsTextItem_TabChangesFocus(QtObjectPtr ptr);
int QGraphicsTextItem_TextInteractionFlags(QtObjectPtr ptr);
char* QGraphicsTextItem_ToHtml(QtObjectPtr ptr);
char* QGraphicsTextItem_ToPlainText(QtObjectPtr ptr);
int QGraphicsTextItem_Type(QtObjectPtr ptr);
void QGraphicsTextItem_DestroyQGraphicsTextItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif