#ifdef __cplusplus
extern "C" {
#endif

void* QHelpEvent_NewQHelpEvent(int ty, void* pos, void* globalPos);
int QHelpEvent_GlobalX(void* ptr);
int QHelpEvent_GlobalY(void* ptr);
int QHelpEvent_X(void* ptr);
int QHelpEvent_Y(void* ptr);

#ifdef __cplusplus
}
#endif