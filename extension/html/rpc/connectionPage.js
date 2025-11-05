const { whitelabelairportClient } = require("./client.js");
const whitelabelairport = require("./whitelabelairport_grpc_web_pb.js");

function openConnectionPage() {
  $("#extension-list-container").show();
  $("#extension-page-container").hide();
  $("#connection-page").show();
  connect();
  $("#connect-button").click(async () => {
    const hsetting_request =
      new whitelabelairport.ChangeWhiteLabelAirportSettingsRequest();
    hsetting_request.setWhiteLabelAirportSettingsJson(
      $("#whitelabelairport-settings").val()
    );
    try {
      const hres =
        await whitelabelairportClient.changeWhiteLabelAirportSettings(
          hsetting_request,
          {}
        );
    } catch (err) {
      $("#whitelabelairport-settings").val("");
      console.log(err);
    }

    const parse_request = new whitelabelairport.ParseRequest();
    parse_request.setContent($("#config-content").val());
    try {
      const pres = await whitelabelairportClient.parse(parse_request, {});
      if (pres.getResponseCode() !== whitelabelairport.ResponseCode.OK) {
        alert(pres.getMessage());
        return;
      }
      $("#config-content").val(pres.getContent());
    } catch (err) {
      console.log(err);
      alert(JSON.stringify(err));
      return;
    }

    const request = new whitelabelairport.StartRequest();

    request.setConfigContent($("#config-content").val());
    request.setEnableRawConfig(false);
    try {
      const res = await whitelabelairportClient.start(request, {});
      console.log(res.getCoreState(), res.getMessage());
      handleCoreStatus(res.getCoreState());
    } catch (err) {
      console.log(err);
      alert(JSON.stringify(err));
      return;
    }
  });

  $("#disconnect-button").click(async () => {
    const request = new whitelabelairport.Empty();
    try {
      const res = await whitelabelairportClient.stop(request, {});
      console.log(res.getCoreState(), res.getMessage());
      handleCoreStatus(res.getCoreState());
    } catch (err) {
      console.log(err);
      alert(JSON.stringify(err));
      return;
    }
  });
}

function connect() {
  const request = new whitelabelairport.Empty();
  const stream = whitelabelairportClient.coreInfoListener(request, {});
  stream.on("data", (response) => {
    console.log("Receving ", response);
    handleCoreStatus(response);
  });

  stream.on("error", (err) => {
    console.error("Error opening extension page:", err);
    // openExtensionPage(extensionId);
  });

  stream.on("end", () => {
    console.log("Stream ended");
    setTimeout(connect, 1000);
  });
}

function handleCoreStatus(status) {
  if (status == whitelabelairport.CoreState.STOPPED) {
    $("#connection-before-connect").show();
    $("#connection-connecting").hide();
  } else {
    $("#connection-before-connect").hide();
    $("#connection-connecting").show();
    if (status == whitelabelairport.CoreState.STARTING) {
      $("#connection-status").text("Starting");
      $("#connection-status").css("color", "yellow");
    } else if (status == whitelabelairport.CoreState.STOPPING) {
      $("#connection-status").text("Stopping");
      $("#connection-status").css("color", "red");
    } else if (status == whitelabelairport.CoreState.STARTED) {
      $("#connection-status").text("Connected");
      $("#connection-status").css("color", "green");
    }
  }
}

module.exports = { openConnectionPage };
