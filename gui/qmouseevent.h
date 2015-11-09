#ifdef __cplusplus
extern "C" {
#endif

void* QMouseEvent_NewQMouseEvent(int ty, void* localPos, int button, int buttons, int modifiers);
void* QMouseEvent_NewQMouseEvent2(int ty, void* localPos, void* screenPos, int button, int buttons, int modifiers);
void* QMouseEvent_NewQMouseEvent3(int ty, void* localPos, void* windowPos, void* screenPos, int button, int buttons, int modifiers);
int QMouseEvent_Button(void* ptr);
int QMouseEvent_Buttons(void* ptr);
int QMouseEvent_Flags(void* ptr);
int QMouseEvent_GlobalX(void* ptr);
int QMouseEvent_GlobalY(void* ptr);
int QMouseEvent_Source(void* ptr);
int QMouseEvent_X(void* ptr);
int QMouseEvent_Y(void* ptr);

#ifdef __cplusplus
}
#endif