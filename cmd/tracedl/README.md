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

`./tracedl get --type nativelog --compress gz --height 3897964`

Download eth logs and store as gzip.  

`./tracedl get --type ethlog --compress gz --height 3897964`

Download tipset and store as gzip.  

`./tracedl get --type tipset --compress gz --height 3897964`