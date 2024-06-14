## tracedl

Helper tool to download raw logs from a node for testing purposes.

### Usage

Replace values in `config.yaml` with desired values.  

```
go build
./tracedl --help
./tracedl get -h
```

 Download traces and store as gzip.  

`./tracedl get --type traces --compress gz --height 3897964 --outPath ../../data/heights`

Download native logs and store as gzip.  

`./tracedl get --type nativelog --compress gz --height 3897964 --outPath ../../data/heights`

Download eth logs and store as gzip.  

`./tracedl get --type ethlog --compress gz --height 3897964 --outPath ../../data/heights`

Download tipset and store as gzip.  

`./tracedl get --type tipset --compress gz --height 3897964 --outPath ../../data/heights`

---
You can use the `script.sh` to automate the download of traces, native logs, eth logs, and tipsets for specified heights.

1. Open `script.sh` and modify the heights array with the desired heights.
2. Run the script:

`./script.sh`
