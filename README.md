# lsmon-to-json
Convert output of Sentinel RMS lsmon to JSON  
#lsmon #zabbix #json

It can help you to get all necessary statistics of Sentinel RMS lsmon to JSON with one Item(features with issued and used lics, active users). Then zabbix can parse it with **JSONPath** to Dependent Items. It's mush more faster then use hundreds of External Items. (read the Warnning on [zabbix docs](ttps://www.zabbix.com/documentation/4.0/manual/config/items/itemtypes/external))

## How to use

copy External Script **lsmon_json.sh** to `/usr/lib/zabbix/externalscripts`. It's default folder for zabbix external scripts. It can be defined by zabbix config file. Wich one can be use to autodiscovery features for LLD and collect all information about licenses  
copy **lsmon-to-json** to `/usr/lib/zabbix/externalscripts` or to the host where FlexLM runs. If you copy **lsmon-to-json** to the host where FlexLM runs you have to rewrite **lsmon_json.sh** for that changes.  

External script which take data to STDIN of **lsmon-to-json** can look like thath examples

Example 1:
<pre>zabbix-get -s 1$ -k system.run["C:\AVEVA\AVEVA Licensing System\RMS\tools\lsmon"] | lsmon-to-json</pre>

Example 2:
<pre>zabbix-get -s 1$ -k system.run["C:\AVEVA\AVEVA Licensing System\RMS\tools\lsmon | lsmon-to-json"]</pre>

On the host where Sentinel RMS lsmon runs you have to set PATH to **lsmon** and to **lsmon-to-json** or use absolute path to these files

Create zabbix Template which will collect stats or import example Template `zbx_export_templates_AVEVA_lics.xml` to zabbix

# Prerequisite
zabbix version > 4.0.11  
[jq](https://stedolan.github.io/jq/)