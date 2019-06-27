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
  gui.clickButton(buttons.NextButton, 30000);
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
  var cBoxes = ["Archive", "LTS", "Latest releases", "Preview"];
  for (var i in cBoxes)
  {
    var box = gui.currentPageWidget().findChild(cBoxes[i]);
    if (box)
    {
      box.checked = true;
    }
  }

  var refButton = gui.currentPageWidget().findChild("FetchCategoryButton");
  if (refButton)
  {
    refButton.click();
  }

  var version = "qt5.5130";
  if (installer.value("VERSION") != "")
  {
    version = installer.value("VERSION");
  }

  if (installer.value("DARWIN") == "true")
  {
    gui.currentPageWidget().selectComponent("qt."+version+".clang_64");
  }

  if (installer.value("IOS") == "true")
  {
    gui.currentPageWidget().selectComponent("qt."+version+".ios");
  }

  if (installer.value("WINDOWS") == "true")
  {
    gui.currentPageWidget().selectComponent("qt."+version+".win32_mingw49");
    gui.currentPageWidget().selectComponent("qt."+version+".win32_mingw53");
    if (installer.value("ARCH") == "32")
    {
      gui.currentPageWidget().selectComponent("qt."+version+".win32_mingw73");
    }
    else
    {
      gui.currentPageWidget().selectComponent("qt."+version+".win64_mingw73");
    }
  }

  if (installer.value("LINUX") == "true")
  {
    gui.currentPageWidget().selectComponent("qt."+version+".gcc_64");
  }

  gui.currentPageWidget().selectComponent("qt."+version+".qt3d");
  gui.currentPageWidget().selectComponent("qt."+version+".qtcanvas3d");
  gui.currentPageWidget().selectComponent("qt."+version+".qtcharts");
  gui.currentPageWidget().selectComponent("qt."+version+".qtdatavis3d");
  gui.currentPageWidget().selectComponent("qt."+version+".qtlocation");
  gui.currentPageWidget().selectComponent("qt."+version+".qtnetworkauth");
  gui.currentPageWidget().selectComponent("qt."+version+".qtpurchasing");
  gui.currentPageWidget().selectComponent("qt."+version+".qtquickcontrols");
  gui.currentPageWidget().selectComponent("qt."+version+".qtquickcontrols2");
  gui.currentPageWidget().selectComponent("qt."+version+".qtremoteobjects");
  gui.currentPageWidget().selectComponent("qt."+version+".qtscript");
  gui.currentPageWidget().selectComponent("qt."+version+".qtserialbus");
  gui.currentPageWidget().selectComponent("qt."+version+".qtspeech");
  gui.currentPageWidget().selectComponent("qt."+version+".qtvirtualkeyboard");
  gui.currentPageWidget().selectComponent("qt."+version+".qtwebengine");
  gui.currentPageWidget().selectComponent("qt."+version+".qtwebglplugin");
  gui.currentPageWidget().selectComponent("qt."+version+".qtwebview");

  gui.currentPageWidget().selectComponent("qt."+version+".android_armv7");
  gui.currentPageWidget().selectComponent("qt."+version+".android_x86");
  gui.currentPageWidget().selectComponent("qt."+version+".android_arm64_v8a");

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
  var checkBox = gui.currentPageWidget().LaunchQtCreatorCheckBoxForm;
  if (checkBox && checkBox.launchQtCreatorCheckBox)
  {
    checkBox.launchQtCreatorCheckBox.checked = false;
  }
  gui.clickButton(buttons.FinishButton);
}
