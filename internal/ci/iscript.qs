function Controller()
{
  installer.wizardPageInsertionRequested.connect(function(widget, page)
  {
    installer.removeWizardPage(installer.components()[0], "WorkspaceWidget");
  })

  installer.autoRejectMessageBoxes();
  installer.installationFinished.connect(function()
  {
    gui.clickButton(buttons.NextButton);
  })
}

Controller.prototype.WelcomePageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.CredentialsPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.IntroductionPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.TargetDirectoryPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.ComponentSelectionPageCallback = function()
{
  if (installer.value("IOS") != "true")
  {
    gui.currentPageWidget().deselectComponent("qt.591.ios")
  }
  gui.currentPageWidget().selectComponent("qt.591.android_armv7")
  gui.currentPageWidget().selectComponent("qt.591.android_x86")
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.LicenseAgreementPageCallback = function()
{
  gui.currentPageWidget().AcceptLicenseRadioButton.setChecked(true);
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.StartMenuDirectoryPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.ReadyForInstallationPageCallback = function()
{
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.FinishedPageCallback = function()
{
  var checkBox = gui.currentPageWidget().LaunchQtCreatorCheckBoxForm
  if (checkBox && checkBox.launchQtCreatorCheckBox)
  {
    checkBox.launchQtCreatorCheckBox.checked = false;
  }
  gui.clickButton(buttons.FinishButton);
}
