# s3shot

Take a screenshot and upload it to an S3 bucket. Currently only supported on Linux-based systems.

## Installation

```text
$ git clone https://github.com/Mrtenz/s3shot
$ cd s3shot
$ make
$ sudo make install
```

### Dependencies

* `maim` to capture the screenshot
* `xclip` to copy the URL to the clipboard
* (optional) `notify-send` to get a notification when an image has been uploaded
* (optional) `pngcrush` to compress images before uploading

## Usage

AWS credentials are automatically taken from `~/.aws/credentials`. You can use the [official AWS CLI](https://github.com/aws/aws-cli) and run the command `aws configure` to set it up.

### Commands

* `all` (Alias: `a`)
  
  Capture the whole screen.
  
* `window` (Alias: `w`)

  Capture the current active window.
  
* `selection` (Alias: `s`)

  Capture a selection.
  
### Options

* `--region REGION` (Alias: `-r REGION`)

  Use the region `REGION`.
  
* `--bucket BUCKET` (Alias: `-b BUCKET`)

  Upload the files to `BUCKET`.
  
* `--compress` (Alias: `-c`)

  Compress the image before uploading
  
* `--notify` (Alias: `-n`)

  Get a notification when uploading is finished
