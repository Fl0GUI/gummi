# Gummi ^^
**A friendly integration layer between web services and SAMMI Core**

### [DOWNLOAD](https://github.com/Fl0GUI/gummi/releases/latest)

## Support

* [x] Gumroad+ [![GitHub issues-open](https://badgen.net/github/label-issues/Fl0GUI/gummi/gumroad/open)](https://github.com/Fl0GUI/gummi/issues?q=is%3Aopen%20label%3Agumroad) 
* [x] FourthWall+ [![GitHub issues-open](https://badgen.net/github/label-issues/Fl0GUI/gummi/fourthwall/open)](https://github.com/Fl0GUI/gummi/issues?q=is%3Aopen%20label%3Afourthwall) 
* [ ] StreamElements+* [![GitHub issues-open](https://badgen.net/github/label-issues/Fl0GUI/gummi/streamelements/open)](https://github.com/Fl0GUI/gummi/issues?q=is%3Aopen%20label%3Astreamelements) 
* [x] Throne+ [![GitHub issues-open](https://badgen.net/github/label-issues/Fl0GUI/gummi/throne/open)](https://github.com/Fl0GUI/gummi/issues?q=is%3Aopen%20label%3Athrone)

\+ *this integration is implemented with security*

\* *there is currently a [SAMMI integration](https://github.com/Fl0GUI/gummi/blob/master/streamelements/streamelements.sef) provided instead*

[Create an issue](https://github.com/Fl0GUI/gummi/issues/new?assignees=&labels=integration%2C+request&projects=&template=feature_request.md&title=Integration+request%3A+) to request an integration.

## How it works

### Events

When set up, Gummi will forward any event of a store to a SAMMI webhook trigger.
The trigger name is `store(:type)`, eg. the name of the store, and a type (when applicable) separated by a colon.

For gumroad this is just `gumroad`.

For FourthWall this is `fourthwall:<type>`, where type is `ORDER_PLACED` for example.
The FourthWall types can be found [on the official documentation](https://docs.fourthwall.dev/webhook-event-types/).

There is also a heartbeat that triggers `gummi:heartbeat`.
If you want to configure this, edit the config file (see log output at start).

### Data

The data Gummi forwards is exactly as it's received.

For gumroad see [the official documentation](https://app.gumroad.com/ping).

For FourthWall see [the official documentation](https://docs.fourthwall.dev/webhook-event-types/).

For StreamElements see [the official documentation](https://dev.streamelements.com/docs/api-docs/5a84cc101a9c5-connecting-via-websocket-using-o-auth2#json-schema).

For Throne see TODO

The web services have ways to simulate events, which is picked up by Gummi as well.
Use these to test your SAMMI implementation.

### Requirements

In order for Gummi to work you need to set up a portforward rule on your router.
To learn how please check with your ISP or search the web.
Further details (like ports) will be shown during setup when first running Gummi.
The rest of the setup process is interactive and needs no explanation.

### Security

Some integrations are marked as being secure, but what does this mean?
The actual implications are integration dependend, but it means that in theory this integration cannot be falsly triggered by a third party.
In practice it means this is just very very very very unlikely to happen.
For some integrations an attacker needs to know your public ip adress and a secret value of the store, OR the security is implemented by the store itself.

#### donate

I'm not requiring payment nor do I depend on donations.
Kind words motivate me more than money.
But if that is you way you show your support you can click on this:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/N4N2XG5FJ)
