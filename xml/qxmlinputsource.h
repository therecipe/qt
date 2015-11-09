#ifdef __cplusplus
extern "C" {
#endif

void* QXmlInputSource_NewQXmlInputSource();
void* QXmlInputSource_NewQXmlInputSource2(void* dev);
char* QXmlInputSource_Data(void* ptr);
void QXmlInputSource_FetchData(void* ptr);
void QXmlInputSource_Reset(void* ptr);
void QXmlInputSource_SetData2(void* ptr, void* dat);
void QXmlInputSource_SetData(void* ptr, char* dat);
void QXmlInputSource_DestroyQXmlInputSource(void* ptr);

#ifdef __cplusplus
}
#endif