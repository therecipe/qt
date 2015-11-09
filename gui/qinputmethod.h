#ifdef __cplusplus
extern "C" {
#endif

int QInputMethod_InputDirection(void* ptr);
int QInputMethod_IsAnimating(void* ptr);
int QInputMethod_IsVisible(void* ptr);
void QInputMethod_ConnectAnimatingChanged(void* ptr);
void QInputMethod_DisconnectAnimatingChanged(void* ptr);
void QInputMethod_Commit(void* ptr);
void QInputMethod_ConnectCursorRectangleChanged(void* ptr);
void QInputMethod_DisconnectCursorRectangleChanged(void* ptr);
void QInputMethod_Hide(void* ptr);
void QInputMethod_ConnectInputDirectionChanged(void* ptr);
void QInputMethod_DisconnectInputDirectionChanged(void* ptr);
void QInputMethod_InvokeAction(void* ptr, int a, int cursorPosition);
void QInputMethod_ConnectKeyboardRectangleChanged(void* ptr);
void QInputMethod_DisconnectKeyboardRectangleChanged(void* ptr);
void QInputMethod_ConnectLocaleChanged(void* ptr);
void QInputMethod_DisconnectLocaleChanged(void* ptr);
void* QInputMethod_QInputMethod_QueryFocusObject(int query, void* argument);
void QInputMethod_Reset(void* ptr);
void QInputMethod_SetInputItemRectangle(void* ptr, void* rect);
void QInputMethod_SetInputItemTransform(void* ptr, void* transform);
void QInputMethod_SetVisible(void* ptr, int visible);
void QInputMethod_Show(void* ptr);
void QInputMethod_Update(void* ptr, int queries);
void QInputMethod_ConnectVisibleChanged(void* ptr);
void QInputMethod_DisconnectVisibleChanged(void* ptr);

#ifdef __cplusplus
}
#endif