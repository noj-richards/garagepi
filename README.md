A Go application that can be used to open and close a garage door using web hooks, a Telegram BOT, an old Raspberry Pi and a PiFACE.
The PiFACE relay is used to send the signal into the garage door controller.
A reed switch on the door is connected through to one of the inputs on the PiFACE to see if the door is open or closed.

The Telegram BOT can be used with a group so everyone is notified if the door is opened.

I recommend servo.net to expose the webhooks to the internet. It is then possible to create REST shortcuts on a phone to control the app from the homescreen or using Siri.
