#ifdef __cplusplus
extern "C" {
#endif

void* QScrollBar_NewQScrollBar(void* parent);
void* QScrollBar_NewQScrollBar2(int orientation, void* parent);
int QScrollBar_Event(void* ptr, void* event);
void QScrollBar_DestroyQScrollBar(void* ptr);

#ifdef __cplusplus
}
#endif