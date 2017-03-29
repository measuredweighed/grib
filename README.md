GRIB2 Golang experimental parser
================================

Unfinished dirty parser for meteo data

### Usage

Install by typing

    go get -u github.com/nilsmagnus/grib

Usage:

    Usage of grib:
      -category int
          	Filters on Category within discipline. -1 means all categories (default -1)
      -discipline int
          	Filters on Discipline. -1 means all disciplines (default -1)
      -export int
          	Export format. Valid types are 0 (none) 1 (json) 
      -file string
          	Grib filepath
      -maxmsg int
          	Maximum number of messages to parse. (default 2147483647). Does not work in combination with other filters.


### What works?

- basic binary parsing of GRIB2 GFS files from NOAA
- implemented only "Grid point data - complex packing and spatial differencing"

### TODO?

- use a proper logging framework(zap?) instead of fmt.Print*
- unit-tests for Section[2-8]
- unit-test for Data3
- implement and test output values
- support for filtering on geolocations

# Grib Documentation

Grib specification:

http://www.wmo.int/pages/prog/www/WMOCodes/Guides/GRIB/GRIB2_062006.pdf

Documentation from noaa.gov :

http://www.nco.ncep.noaa.gov/pmb/docs/on388/


Examples can be found at

http://www.ftp.ncep.noaa.gov/data/nccf/com/gfs/prod/
