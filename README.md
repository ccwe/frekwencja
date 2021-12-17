# Frekwencja CCWE

This repository contains the code for [frekwencja.ccwe.pl][frekwencja] service that was deployed prior to app's shutdown somewhere near the end of the year 2020. It has been uploaded as a curiosity and to provide some transparency for the app.

### December 2021 update:

As of now, version 1.6 should be live and running (with some major changes to the project's inner workings); however, it was released solely due to a surge of request for bringing back the service. The project is currently undergoing changes to comply with third-party companies' policies, that will be able to let it operate again. When the new project will be finished, this repo will go back into legacy mode.

# Build the executable

> Warning!: the application was only tested on MacOS.

Pre-compiled binaries are included in each release tag, so you can just go get one there. Download links are also available on [frekwencja.ccwe.pl][frekwencja].

To compile the binary yourself, make sure you have installed [Go](https://go.dev) and [Make](https://www.gnu.org/software/make/). Then, go into `cmd` directory and simply run `make <system>`, where `<system>` is one of the following: windows, macos, linux, android, ios. This should match the name of your operating system. The binary will appear in the `cmd/bin` directory. If you can't use Make, just find the appropriate command in the Makefile and run it manually.

# License

The whole project is licensed under GNU General Public License v3 and comes with absolutely no warranty.

All third-party trademarks belong to their respective owners, and are used only to identify the corresponding third-party goods and services. The use of these trademarks doesn't indicate any relationship or endorsement between the third-party and the authors.

Copyright (C) 2019-2022 [Centrum Cyfrowego Wsparcia Edukacji][ccwe]

[frekwencja]: https://frekwencja.ccwe.pl
[ccwe]: https://ccwe.pl
