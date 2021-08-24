package builder

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/yettera765/mmdb-builder/common"
	"github.com/yettera765/mmdb-builder/config"
	C "github.com/yettera765/mmdb-builder/constant"
)

type Builder interface {
	Process(conf config.Config)
	Build(dbPath string)
}

type builder struct {
	writer *mmdbwriter.Tree
}

const (
	dbtype       = "GeoIP2-Country"
	dbRecordSize = 24
	dbCountry    = "country"
	dbIsoCode    = "iso_code"
)

func New() Builder {
	w, _ := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: dbtype,
			RecordSize:   dbRecordSize,
		},
	)
	return &builder{writer: w}
}

func (b *builder) Process(config config.Config) {
	for geo, ips := range config.GeoNets() {
		b.setNets(geo, mergeIPs(ips))
	}
}

func (b *builder) Build(dbPath string) {
	common.Mkdir(dbPath)
	log.Printf("building to %s ...", dbPath)
	fh, err := os.Create(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	if _, err := b.writer.WriteTo(fh); err != nil {
		log.Fatalln(err)
	}
}

func (b *builder) setNets(geoTag string, ips []C.IP) {
	log.Println(geoTag, "nets count:", len(ips))
	if len(ips) == 0 {
		log.Println("skip ...", geoTag)
		return
	}
	metadata := getMeta(geoTag)
	for _, ip := range ips {
		if err := b.writer.Insert(ip.Net(), metadata); err != nil {
			log.Printf("%s, net:'%s'", err, ip.Net())
		}
	}
}

func getMeta(geoTag string) *mmdbtype.Map {
	code := mmdbtype.String(strings.ToUpper(geoTag))
	return &mmdbtype.Map{
		dbCountry: mmdbtype.Map{
			dbIsoCode: code,
		},
	}
}

func mergeIPs(ips []C.IP) []C.IP {
	var cidrs []*net.IPNet
	for _, p := range ips {
		cidrs = append(cidrs, p.Net())
	}
	coalescedIPV4, coalescedIPV6 := CoalesceCIDRs(cidrs)
	var result []C.IP
	for _, ipNet := range append(coalescedIPV4, coalescedIPV6...) {
		if p, err := C.NewIP(ipNet.String()); err == nil {
			result = append(result, p)
		}
	}
	return result
}
