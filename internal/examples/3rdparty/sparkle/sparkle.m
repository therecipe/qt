#import <Headers/SUUpdater.h>

static SUUpdater* updater = nil;

void sparkle_checkUpdates()
{
	if (!updater)
		updater = [[SUUpdater sharedUpdater] retain];

	[updater setUpdateCheckInterval:3600];
	[updater checkForUpdatesInBackground];
}
