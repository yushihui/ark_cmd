## Installation

### Download/clone this repo
```sh
git clone https://github.com/yushihui/ark_cmd
```

## How to use this tool
go to your local ark_cmd directory 

### fetch latest profolio of all ark fund
```sh
ark_cmd fetch
# fetch latest profolio of all ark fund
```

### list all ark fund
```sh
ark_cmd list
#list all ark funds
```

### fund [fund_name] -s start_date -e end_date
```sh
ark_cmd fund arkk -s 2021-01-05 -e 2021-01-06
#get arkk's delta change (trads) between start date and end date
```

### profolio [fund_name]
```sh
ark_cmd profolio arkk
# get the fund arkk's current profolio
```

### index
```sh
ark_cmd index
# build index
```


### ticker [ticker]
```sh
ark_cmd ticker TSLA
# make sure build index first 
# get all activities of Tesla
```

