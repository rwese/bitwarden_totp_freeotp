Extract Bitwarden TOTP
======================

This is a tool that allows you to extract Time-based One-Time Passwords (TOTPs) from a Bitwarden JSON file and convert them into [FreeOTP](https://github.com/freeotp/freeotp-android) compatible URLs.

Requirements
------------

-   Go 1.15 or later

Installation
------------

You can install the binary with the following command:

go

`go get -u github.com/rwese/extract_bitwarden_totp`

Usage
-----

The binary can be used from the command line with the following options:

lua

`Usage of extract_bitwarden_totp:
  -input string
        input JSON file (mandatory)
  -output string
        output text file (mandatory)`

### Example

lua

`extract_bitwarden_totp -input bitwarden.json -output totp.txt`

Development
-----------

You can build the binary by running `make` in the project root.

Contributing
------------

Feel free to submit a pull request or open an issue if you have any suggestions or find any bugs.

License
-------

This project is licensed under the [MIT License](https://chat.openai.com/chat/LICENSE).
