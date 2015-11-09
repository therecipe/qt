#ifdef __cplusplus
extern "C" {
#endif

void* QEnterEvent_NewQEnterEvent(void* localPos, void* windowPos, void* screenPos);
int QEnterEvent_GlobalX(void* ptr);
int QEnterEvent_GlobalY(void* ptr);
int QEnterEvent_X(void* ptr);
int QEnterEvent_Y(void* ptr);

#ifdef __cplusplus
}
#endif