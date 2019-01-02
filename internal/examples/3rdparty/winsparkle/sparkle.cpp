#include "sparkle.h"

#include "winsparkle.h"
#include <QResource>

void sparkle_checkUpdates()
{
	win_sparkle_set_dsa_pub_pem(reinterpret_cast<const char *>(QResource(":/WinSparkle/dsa_pub.pem").data()));
	win_sparkle_init();
}