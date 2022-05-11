# pixelimagedl

## Install as CLI app

Make sure you have Go 1.18 or later installed, then:

`go install github.com/jalavosus/pixelimagedl/cmd/pixelimagedl@latest`

## Usage

To list available OTA or Factory images for your pixel device:

`pixelimagedl list -d [device] -t [image type]`

Example using OTA images:

```
$> pixelimagedl list -d oriole -t ota ### fetch OTA images for Pixel 6 (oriole)

Output:

{
  "version": "12.0.0",
  "build_number": "SD1A.210817.015.A4",
  "build_date": "Oct 2021",
  "build_comment": "",
  "download_uri": "https://dl.google.com/dl/android/aosp/oriole-ota-sd1a.210817.015.a4-19a77b62.zip",
  "sha256_sum": "19a77b62a80732b0e20d3ac369288fc827554c1a8c109c6e9fa2a420998dd826"
}
{
  "version": "12.0.0",
  "build_number": "SD1A.210817.019.B1",
  "build_date": "Oct 2021",
  "build_comment": "AT&T",
  "download_uri": "https://dl.google.com/dl/android/aosp/oriole-ota-sd1a.210817.019.b1-67195ca0.zip",
  "sha256_sum": "67195ca0fa1f9cfc9f331ab3e4f44fdbf819896059dcb1819c1125edef854c25"
}
[...]
{
  "version": "12.1.0",
  "build_number": "SP2A.220405.004",
  "build_date": "Apr 2022",
  "build_comment": "",
  "download_uri": "https://dl.google.com/dl/android/aosp/oriole-ota-sp2a.220405.004-f019343f.zip",
  "sha256_sum": "f019343fc56cf466249580dfbdee30ae341a88d6429ecedda727097f51d41c28"
}
{
  "version": "12.1.0",
  "build_number": "SP2A.220505.002",
  "build_date": "May 2022",
  "build_comment": "",
  "download_uri": "https://dl.google.com/dl/android/aosp/oriole-ota-sp2a.220505.002-513d254d.zip",
  "sha256_sum": "513d254dc29c6a61f26fda316eededf265b4fc4b0b528d761a852d074632ec71"
}
```

To download an OTA or Factory image for a device:

`pixelimagedl download -d [device] -t [image type]`

Example:

```
$> pixelimagedl download -d pixel4a -t factory

Output:

2022/05/11 13:41:33 latest stable factory image for Pixel 4a is 12.1.0 (SP2A.220505.002)
2022/05/11 13:41:33 downloading factory image from https://dl.google.com/dl/android/aosp/sunfish-sp2a.220505.002-factory-2ca902f1.zip
2022/05/11 13:41:33 saving factory image to sunfish-sp2a.220505.002-factory-2ca902f1.zip
 100% |████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| (2.0/2.0 GB, 9.546 MB/s)          
2022/05/11 13:08:36 saved 2.0Gb to sunfish-sp2a.220505.002-factory-2ca902f1.zip
2022/05/11 13:08:41 SHA256 sum 2ca902f10574fa45806c598e8cf10584aa2dc3dd0fb1cf75882e50426187221c of downloaded file matches expected
```

### Device names

Currently supported devices (and their codenames), and their corresponding CLI values:

- Pixel 6 ("Oriole")
  - `pixel6`
  - `oriole`
- Pixel 6 Pro ("Raven")
  - `pixel6pro`
  - `raven`
- Pixel 5 ("Redfin") 
  - `pixel5`
  - `redfin`
- Pixel 5a ("Barbet")
  - `pixel5a`
  - `barbet`
- Pixel 4 ("Flame")
  - `pixel4`
  - `flame`
- Pixel 4 XL ("Coral")
  - `pixel4xl`
  - `coral`  
- Pixel 4a ("Sunfish") 
  - `pixel4a`
  - `sunfish`
- Pixel 4a (5G) ("Bramble") 
  - `pixel4a5g`
  - `bramble`

## TODO

- Implement functionality for listing/downloading available beta versions