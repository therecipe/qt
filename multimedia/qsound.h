#ifdef __cplusplus
extern "C" {
#endif

void QSound_SetLoops(void* ptr, int number);
void* QSound_NewQSound(char* filename, void* parent);
char* QSound_FileName(void* ptr);
int QSound_IsFinished(void* ptr);
int QSound_Loops(void* ptr);
int QSound_LoopsRemaining(void* ptr);
void QSound_Play2(void* ptr);
void QSound_QSound_Play(char* filename);
void QSound_Stop(void* ptr);
void QSound_DestroyQSound(void* ptr);

#ifdef __cplusplus
}
#endif