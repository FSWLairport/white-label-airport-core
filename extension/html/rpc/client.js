const whitelabelairport = require("./whitelabelairport_grpc_web_pb.js");
const extension = require("./extension_grpc_web_pb.js");

const grpcServerAddress = "/";
const extensionClient = new extension.ExtensionHostServicePromiseClient(
  grpcServerAddress,
  null,
  null
);
const whitelabelairportClient = new whitelabelairport.CorePromiseClient(
  grpcServerAddress,
  null,
  null
);

module.exports = { extensionClient, whitelabelairportClient };
