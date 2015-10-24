#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QInputMethod_InputDirection(QtObjectPtr ptr);
int QInputMethod_IsAnimating(QtObjectPtr ptr);
int QInputMethod_IsVisible(QtObjectPtr ptr);
void QInputMethod_ConnectAnimatingChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectAnimatingChanged(QtObjectPtr ptr);
void QInputMethod_Commit(QtObjectPtr ptr);
void QInputMethod_ConnectCursorRectangleChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectCursorRectangleChanged(QtObjectPtr ptr);
void QInputMethod_Hide(QtObjectPtr ptr);
void QInputMethod_ConnectInputDirectionChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectInputDirectionChanged(QtObjectPtr ptr);
void QInputMethod_InvokeAction(QtObjectPtr ptr, int a, int cursorPosition);
void QInputMethod_ConnectKeyboardRectangleChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectKeyboardRectangleChanged(QtObjectPtr ptr);
void QInputMethod_ConnectLocaleChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectLocaleChanged(QtObjectPtr ptr);
char* QInputMethod_QInputMethod_QueryFocusObject(int query, char* argument);
void QInputMethod_Reset(QtObjectPtr ptr);
void QInputMethod_SetInputItemRectangle(QtObjectPtr ptr, QtObjectPtr rect);
void QInputMethod_SetInputItemTransform(QtObjectPtr ptr, QtObjectPtr transform);
void QInputMethod_SetVisible(QtObjectPtr ptr, int visible);
void QInputMethod_Show(QtObjectPtr ptr);
void QInputMethod_Update(QtObjectPtr ptr, int queries);
void QInputMethod_ConnectVisibleChanged(QtObjectPtr ptr);
void QInputMethod_DisconnectVisibleChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif