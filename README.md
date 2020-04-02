# VRHR Server

This is the server for VRHR, a piece of software that allows you to use your Wear OS device's heart rate sensor while streaming (or for other things).

## Setup

1. Download the appropriate binary from the releases page.
2. Save it somewhere (preferably in its own directory).
3. If you want to run it as-is with no configuration options, simply double-click the binary (if you are on Windows). If you want to configure it, check out the section below.
4. Download the VRHR app if you haven't already. (Make sure you download the phone app AND watch app. Starting with Wear OS 2.0, you need to download them separately on both devices.)
5. In the phone app, enter the URL to send the heart rate data to. In this case, it would be: `http://your.computer.ip:8000/api/state` if you haven't changed any options.
6. Open the VRHR watch app. You are done!

### OBS Setup

![](https://user-images.githubusercontent.com/2646487/78212484-dd414e00-747d-11ea-8ba9-418835140cdd.png)

The server hosts a web UI you can embed into anything you want. I imagine the most common use case would be OBS for VR streaming.

1. Create a "Browser" source in your scene.
2. Set the width to 800 and the height to 400.
3. Set the URL to `http://localhost:8000` (if you have not changed the port).
4. Modify other parameters as you wish.

Once you're done, you can treat it like you would any other source.

## Configuration

```
Usage of vrhr:
  --csv
        Write data to a csv file
```

The `--csv` option will output a CSV file adjacent to your working directory. The output location is not customizable, but will probably be in the future. The CSV file will contain every heart rate change as well as the date it was logged, and the accuracy at time of capture.

## License

This software is licensed under the GNU GPL Version 3.0. This license does not allow the use of this software for commercial purposes, however **I am granting an exception for those who wish to use the software as a heart rate monitor for Twitch/YouTube/etc streams and want to monetize such streams**. That is 100% alright. All I ask in this case is that you mention & link to the software in the description of your stream or video. Thank you!

The other terms still apply; for example:

- You may not redistribute this software commercially. It is free.
- You may not modify the software and redistribute it without the accompanying source code.

This list is not exhaustive; please see the license for the full terms.

```
VRHR Server
Copyright (C) 2020 TJ Horner

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```