#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMediaContent_NewQMediaContent();
QtObjectPtr QMediaContent_NewQMediaContent7(QtObjectPtr playlist, char* contentUrl, int takeOwnership);
QtObjectPtr QMediaContent_NewQMediaContent6(QtObjectPtr other);
QtObjectPtr QMediaContent_NewQMediaContent4(QtObjectPtr resource);
QtObjectPtr QMediaContent_NewQMediaContent3(QtObjectPtr request);
QtObjectPtr QMediaContent_NewQMediaContent2(char* url);
char* QMediaContent_CanonicalUrl(QtObjectPtr ptr);
int QMediaContent_IsNull(QtObjectPtr ptr);
QtObjectPtr QMediaContent_Playlist(QtObjectPtr ptr);
void QMediaContent_DestroyQMediaContent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif