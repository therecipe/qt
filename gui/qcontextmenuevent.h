#ifdef __cplusplus
extern "C" {
#endif

void* QContextMenuEvent_NewQContextMenuEvent3(int reason, void* pos);
void* QContextMenuEvent_NewQContextMenuEvent2(int reason, void* pos, void* globalPos);
void* QContextMenuEvent_NewQContextMenuEvent(int reason, void* pos, void* globalPos, int modifiers);
int QContextMenuEvent_GlobalX(void* ptr);
int QContextMenuEvent_GlobalY(void* ptr);
int QContextMenuEvent_Reason(void* ptr);
int QContextMenuEvent_X(void* ptr);
int QContextMenuEvent_Y(void* ptr);

#ifdef __cplusplus
}
#endif