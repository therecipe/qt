#ifdef __cplusplus
extern "C" {
#endif

void* QQmlApplicationEngine_NewQQmlApplicationEngine(void* parent);
void* QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, void* parent);
void* QQmlApplicationEngine_NewQQmlApplicationEngine2(void* url, void* parent);
void QQmlApplicationEngine_Load2(void* ptr, char* filePath);
void QQmlApplicationEngine_Load(void* ptr, void* url);
void QQmlApplicationEngine_LoadData(void* ptr, void* data, void* url);
void QQmlApplicationEngine_DestroyQQmlApplicationEngine(void* ptr);

#ifdef __cplusplus
}
#endif