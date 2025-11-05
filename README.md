# whitelabelairport-core

## Docker

To Run our docker image see https://github.com/pppwaw/white-label-airport-core/pkgs/container/whitelabelairport-core

Docker

```
docker pull ghcr.io/pppwaw/whitelabelairport-core:latest
```

Docker Compose

```
git clone https://github.com/pppwaw/white-label-airport-core
cd whitelabelairport-core/docker
docker-compose up
```

## WRT

...

## Extension

An extension is something that can be added to whitelabelairport application by a third party. It will add capability to modify configs, do some extra action, show and receive data from users.

This extension will be shown in all WhiteLabelAirport Platforms such as Android/macOS/Linux/Windows/iOS

[Create an extension](https://github.com/hiddify/hiddify-app-example-extension)

Features and Road map:

- [x] Add Third Party Extension capability
- [x] Test Extension from Browser without any dependency to android/mac/.... `./cmd.sh extension` the open browser `https://127.0.0.1:12346`
- [x] Show Custom UI from Extension `github.com/pppwaw/white-label-airport-core/extension.UpdateUI()`
- [x] Show Custom Dialog from Extension `github.com/pppwaw/white-label-airport-core/extension.ShowDialog()`
- [x] Show Alert Dialog from Extension `github.com/pppwaw/white-label-airport-core/extension.ShowMessage()`
- [x] Get Data from UI `github.com/pppwaw/white-label-airport-core/extension.SubmitData()`
- [x] Save Extension Data from `e.Base.Data`
- [x] Load Extension Data to `e.Base.Data`
- [x] Disable / Enable Extension
- [x] Update user proxies before connecting `github.com/pppwaw/white-label-airport-core/extension.BeforeAppConnect()`
- [x] Run Tiny Independent Instance `github.com/pppwaw/white-label-airport-core/extension/sdk.RunInstance()`
- [x] Parse Any type of configs/url `github.com/pppwaw/white-label-airport-core/extension/sdk.ParseConfig()`
- [ ] ToDo: Add Support for MultiLanguage Interface
- [ ] ToDo: Custom Extension Outbound
- [ ] ToDo: Custom Extension Inbound
- [ ] ToDo: Custom Extension ProxyConfig

Demo Screenshots from HTML:

 <img width="531" alt="image" src="https://github.com/user-attachments/assets/0fbef76f-896f-4c45-a6b8-7a2687c47013">
 <img width="531" alt="image" src="https://github.com/user-attachments/assets/15bccfa0-d03e-4354-9368-241836d82948">
