#!/bin/bash
# Get Sentinel RMS lsmon stats for zabbix LLD

JSON="{\"data\":"
lsmonToJson=/usr/lib/zabbix/externalscripts/lsmon-to-json


# Features discovery
# Key: featdiscovery
if [[ $2 = "featdiscovery" ]]
then
get=`zabbix_get -s $1 -k system.run["C:\AVEVA\AVEVA Licensing System\RMS\tools\lsmon.exe"] | $lsmonToJson | jq --compact-output '[{"{#FEATURE}": .Features[].Feature."Feature name"}]'`
JSON=$JSON"$get}"
echo $JSON
fi

# Get all data in JSON
# Key: getlics
if [[ $2 = "getlics" ]]
    then
    zabbix_get -s $1 -k system.run["C:\AVEVA\AVEVA Licensing System\RMS\tools\lsmon.exe"] | $lsmonToJson
fi