#ifdef __cplusplus
extern "C" {
#endif

int QPauseAnimation_Duration(void* ptr);
void QPauseAnimation_SetDuration(void* ptr, int msecs);
void* QPauseAnimation_NewQPauseAnimation(void* parent);
void* QPauseAnimation_NewQPauseAnimation2(int msecs, void* parent);
void QPauseAnimation_DestroyQPauseAnimation(void* ptr);

#ifdef __cplusplus
}
#endif