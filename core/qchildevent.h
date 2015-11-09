#ifdef __cplusplus
extern "C" {
#endif

void* QChildEvent_NewQChildEvent(int ty, void* child);
int QChildEvent_Added(void* ptr);
void* QChildEvent_Child(void* ptr);
int QChildEvent_Polished(void* ptr);
int QChildEvent_Removed(void* ptr);

#ifdef __cplusplus
}
#endif