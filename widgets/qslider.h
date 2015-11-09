#ifdef __cplusplus
extern "C" {
#endif

void QSlider_SetTickInterval(void* ptr, int ti);
void QSlider_SetTickPosition(void* ptr, int position);
int QSlider_TickInterval(void* ptr);
int QSlider_TickPosition(void* ptr);
void* QSlider_NewQSlider(void* parent);
void* QSlider_NewQSlider2(int orientation, void* parent);
int QSlider_Event(void* ptr, void* event);
void QSlider_DestroyQSlider(void* ptr);

#ifdef __cplusplus
}
#endif