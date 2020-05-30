function Controller()
{
  installer.wizardPageInsertionRequested.connect(function(widget, page)
  {
    installer.removeWizardPage(installer.components()[0], "WorkspaceWidget");
  })

  installer.setDefaultPageVisible(4092, false);
  installer.setDefaultPageVisible(4093, false);
  installer.setDefaultPageVisible(4094, false);
  installer.setDefaultPageVisible(4095, false);

  if (!installer.isOfflineOnly())
  {
    var repos = [
        "http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/licenses/",
        "http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/tools_generic/",
        "http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/tools_qtcreator/",

        "http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/"+installer.value("VERSION").replace(".", "_")+"/",
        "http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/"+installer.value("VERSION").replace(".", "_")+"_src_doc_examples/"
    ];

    if (installer.value("PREVIEW") == "true")
    {
      repos.push("http://download.qt.io/online/qtsdkrepository/"+installer.value("os")+"_x64/desktop/"+installer.value("VERSION").replace(".", "_").slice(0, -1)+"_preview/")
    }

    installer.addUserRepositories(repos)
  }

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

Controller.prototype.DynamicTelemetryPluginFormCallback = function()
{
  var page = gui.pageWidgetByObjectName("DynamicTelemetryPluginForm");
  page.statisticGroupBox.disableStatisticRadioButton.setChecked(true);
  gui.clickButton(buttons.NextButton);
}

Controller.prototype.ObligationsPageCallback = function()
{
  var page = gui.pageWidgetByObjectName("ObligationsPage");
  page.obligationsAgreement.setChecked(true);

  //company
  //var nameEdit = gui.findChild(page, "CompanyName")
  //if (nameEdit) { nameEdit.text = "My Company" }

  //individual
  var individualCheckbox = gui.findChild(page, "IndividualPerson")
  if (individualCheckbox) { individualCheckbox.checked = true; }

  page.completeChanged();
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
  var prefix = "";
  if (installer.value("VERSION") != "")
  {
    version = installer.value("VERSION");
  }

  if (installer.value("PREVIEW") == "true")
  {
    version = version.slice(0, -1)
    prefix = "preview."
  }

  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qt3d");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtcanvas3d");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtcharts");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtdatavis3d");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtlocation");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtnetworkauth");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtpurchasing");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtquickcontrols");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtquickcontrols2");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtremoteobjects");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtscript");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtserialbus");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtspeech");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtvirtualkeyboard");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtwebengine");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtwebglplugin");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtwebview");
  gui.currentPageWidget().selectComponent(prefix+"qt."+version+".qtlottie");

  if (installer.value("LINUX") == "true")
  {
    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".gcc_64");

    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".android_armv7");
    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".android_x86");
    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".android_arm64_v8a");
  }

  if (installer.value("DARWIN") == "true")
  {
    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".clang_64");
  }

  if (installer.value("IOS") == "true")
  {
    gui.currentPageWidget().selectComponent(prefix+"qt."+version+".ios");
  }

  if (installer.value("WINDOWS") == "true")
  {
    if (installer.value("MSVC") == "true")
    {
      if (installer.value("ARCH") == "32")
      {
        gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win32_msvc2017");
        gui.currentPageWidget().selectComponent("qt.tools.win32_mingw730")
      }
      else
      {
        gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win64_msvc2017_64");
        gui.currentPageWidget().selectComponent("qt.tools.win64_mingw730")
      }
    }
    else
    {
      gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win32_mingw49");
      gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win32_mingw53");
      if (installer.value("ARCH") == "32")
      {
        gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win32_mingw73");
      }
      else
      {
        gui.currentPageWidget().selectComponent(prefix+"qt."+version+".win64_mingw73");
      }
    }
  }

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
