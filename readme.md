# Backup Uploader

Uploads a file to the S3 bucket defined in the config file. Note: it only uploads one file each time it's run, so be sure to zip etc first.

# Usage

Windows:
```
backup-uploader.exe <file-to-upload>
```

Linux:
```
backup-uploader <file-to-upload>
```

# Configuration

A file called `config.json` must be in the same directory as the executable binary. 

An example `config.json.example` is provided as a template. Copy this to `config.json` and replace the underscores with actual values.