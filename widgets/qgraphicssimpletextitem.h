#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGraphicsSimpleTextItem_Contains(QtObjectPtr ptr, QtObjectPtr point);
int QGraphicsSimpleTextItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item);
void QGraphicsSimpleTextItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget);
void QGraphicsSimpleTextItem_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QGraphicsSimpleTextItem_SetText(QtObjectPtr ptr, char* text);
char* QGraphicsSimpleTextItem_Text(QtObjectPtr ptr);
int QGraphicsSimpleTextItem_Type(QtObjectPtr ptr);
void QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif