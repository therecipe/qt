#ifdef __cplusplus
extern "C" {
#endif

void* QQmlIncubationController_NewQQmlIncubationController();
void* QQmlIncubationController_Engine(void* ptr);
void QQmlIncubationController_IncubateFor(void* ptr, int msecs);
int QQmlIncubationController_IncubatingObjectCount(void* ptr);

#ifdef __cplusplus
}
#endif