# gobraviacontrol
An Unofficial GoLang API for Sony Bravia Professional Displays

This library is not affiliated with Sony, not endorsed by Sony, and not supported by Sony.  

Please see the [LICENSE file](LICENSE). 

It implements the Simple IP control protocol, as publicly documented [here](https://pro-bravia.sony.net/develop/integrate/ssip/overview/index.html).

Due to the current pandemic, I do not have access to an _actual_ display to test this with.

## Goals
- [X] Send Commands
- [x] Receive Answers to the Commands
- [X] Correctly match Answers to the Commands that caused them
- [ ] Allows routing of Notification messages from the TV
- [ ] Full implementation of the command set