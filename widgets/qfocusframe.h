#ifdef __cplusplus
extern "C" {
#endif

void* QFocusFrame_NewQFocusFrame(void* parent);
void QFocusFrame_SetWidget(void* ptr, void* widget);
void* QFocusFrame_Widget(void* ptr);
void QFocusFrame_DestroyQFocusFrame(void* ptr);

#ifdef __cplusplus
}
#endif