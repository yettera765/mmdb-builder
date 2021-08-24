# mmdb-builder

build mmdb based on custom resources

for testing and personal use

GeoIP2-Country type mmdb, current geo tags: cn, tg

run: `./mmdb-builder -c /path/to/config.yaml -o /path/to/Country.mmdb`

example [configs](./mmdb.yaml):

```yaml
cn:
  - ip: https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt
  - rule: https://raw.githubusercontent.com/LM-Firefly/Rules/master/Special/TeamViewer-CIDR.list
tg:
  - ip: https://www.cloudflare.com/ips-v4
  - ip: https://www.cloudflare.com/ips-v6
  - ip: https://core.telegram.org/resources/cidr.txt
  - rule: https://raw.githubusercontent.com/DivineEngine/Profiles/master/Surge/Ruleset/Extra/IP-Blackhole.list
gd:
  - ip: https://raw.githubusercontent.com/yettera765/rulesets/dev/block/block.ip.txt
```
