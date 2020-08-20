#	decon

A radically simple service for sharing data between devices.

How simple?
- Consists of a REST API and Redis cache
- Store any data (within size limits) in cache associated with any key
  - Allows for specifying retention (TTL in cache)
- Retrieve that data from anywhere

Why is this useful?
- Allows you to share between your devices without copying and pasting anything
  - Simple default URLs, ability to specify your own
  - No need to copy a massive link to your mobile just to access it, can type it out easily
- No restrains on data format, you store whatever you want
  - Can store files, encrypted passwords, whatever you send
  - When you retrieve the data you know what format it's in so you can work with it
    - Unzip the file you uploaded, decrypt the password you stored

Made to be an easier and more user-friendly [PrivateBin](https://privatebin.info/).
