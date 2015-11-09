#ifdef __cplusplus
extern "C" {
#endif

void* QKeySequence_NewQKeySequence();
void* QKeySequence_NewQKeySequence5(int key);
void* QKeySequence_NewQKeySequence4(void* keysequence);
void* QKeySequence_NewQKeySequence2(char* key, int format);
void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4);
int QKeySequence_Count(void* ptr);
int QKeySequence_IsEmpty(void* ptr);
int QKeySequence_Matches(void* ptr, void* seq);
void QKeySequence_Swap(void* ptr, void* other);
char* QKeySequence_ToString(void* ptr, int format);
void QKeySequence_DestroyQKeySequence(void* ptr);

#ifdef __cplusplus
}
#endif