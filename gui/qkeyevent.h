#ifdef __cplusplus
extern "C" {
#endif

int QKeyEvent_Matches(void* ptr, int key);
int QKeyEvent_Count(void* ptr);
int QKeyEvent_IsAutoRepeat(void* ptr);
int QKeyEvent_Key(void* ptr);
int QKeyEvent_Modifiers(void* ptr);
char* QKeyEvent_Text(void* ptr);

#ifdef __cplusplus
}
#endif