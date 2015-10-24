#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QXmlInputSource_NewQXmlInputSource();
QtObjectPtr QXmlInputSource_NewQXmlInputSource2(QtObjectPtr dev);
char* QXmlInputSource_Data(QtObjectPtr ptr);
void QXmlInputSource_FetchData(QtObjectPtr ptr);
void QXmlInputSource_Reset(QtObjectPtr ptr);
void QXmlInputSource_SetData2(QtObjectPtr ptr, QtObjectPtr dat);
void QXmlInputSource_SetData(QtObjectPtr ptr, char* dat);
void QXmlInputSource_DestroyQXmlInputSource(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif