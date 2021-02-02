# KindleHighlightReader <img width="35px" src="https://www.flaticon.com/svg/static/icons/svg/845/845938.svg">


KindleHighlightsReader is a program that reads your highlights from your Kindle device with options to clean and style the highlights and exports them in the following formats: text, json, csv or pdf.


## Features in 1.0.0
- Trim all highlights for redundant characters before and after like: `"ed. Hello"` > `"Hello"`
- Insert periods on all highlights.
- Insert double quotations on all highlights.
- Remove any quotation from all highlights.
- Capialize first letter on all highlights.
- Export as .txt .json .csv or .pdf files.
- Works for Windows and MacOS

#### JSON

The JSON format is useful for developers who wants to unmarshal the JSON string to objects in any langauge to be used for an app, webpage or database. The JSON output is saved in an `.json` file and has the following output format:

```
[
  {
    "Title": "",
    "Author": "",
    "Text": ""
  }
]
```

## Installation & Usage

**Windows** 

Run the `KindleHighlightsReader.exe`. The program will open in a command prompt window and will automatically find your `My Clippings.txt` file if its already on the Desktop or if your Kindle device is plugged into your Windows computer.

[Download .exe for Windows](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReade.exe)
or clone the project and run it from your IDE

**MacOs**

On Mac you just need to run the binary by.

[Download for MacOS](https://github.com/Muddz/KindleHighlightReader/raw/master/KindleHighlightsReaderMacOS)


## License

    Copyright 2020 Muddi Walid

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
