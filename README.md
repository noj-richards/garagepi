A golang application that be used to open and close a garage door using a Telegram BOT, an old Raspberry Pi and a PiFACE.
The PiFACE relay is used to send the signal into the garage door controller.
A reed switch on the door is connected through to one of the inputs on the PiFACE.

The Telegram BOT can be used with a group so everyone is notified if the door is opened.

There are also webhooks to control the door and I recommend servo.net to expose the webhooks to the internet.

You can the create REST shortcuts on a phone to control the app from the homescren or using Siri.
