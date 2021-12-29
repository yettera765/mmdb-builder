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
  - ip: https://core.telegram.org/resources/cidr.txt
```
