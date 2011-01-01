package main

import (
	"dns"
        "dns/dnssec"
        "dns/resolver"
	"fmt"
)

func main() {


	key := new(dns.RR_DNSKEY)
	key.Hdr.Name = "miek.nl"
	key.Hdr.Rrtype = dns.TypeDNSKEY
	key.Hdr.Class = dns.ClassINET
	key.Hdr.Ttl = 3600
	key.Flags = 257
	key.Protocol = 3
	key.Algorithm = dnssec.AlgRSASHA1
	key.PubKey = "AwEAAaHIwpx3w4VHKi6i1LHnTaWeHCL154Jug0Rtc9ji5qwPXpBo6A5sRv7cSsPQKPIwxLpyCrbJ4mr2L0EPOdvP6z6YfljK2ZmTbogU9aSU2fiq/4wjxbdkLyoDVgtO+JsxNN4bjr4WcWhsmk1Hg93FV9ZpkWb0Tbad8DFq NDzr//kZ"

	sig := new(dns.RR_RRSIG)
	sig.Hdr.Name = "miek.nl."
	sig.Hdr.Rrtype = dns.TypeRRSIG
	sig.Hdr.Class = dns.ClassINET
	sig.Hdr.Ttl = 3600
	sig.TypeCovered = dns.TypeDNSKEY
	sig.Algorithm = dnssec.AlgRSASHA1
	sig.OrigTtl = 4000
	sig.Expiration = 1000
	sig.Inception = 800
	sig.KeyTag = 34641
	sig.SignerName = "miek.nl."
	sig.Signature = "AwEAAaHIwpx3w4VHKi6i1LHnTaWeHCL154Jug0Rtc9ji5qwPXpBo6A5sRv7cSsPQKPIwxLpyCrbJ4mr2L0EPOdvP6z6YfljK2ZmTbogU9aSU2fiq/4wjxbdkLyoDVgtO+JsxNN4bjr4WcWhsmk1Hg93FV9ZpkWb0Tbad8DFq NDzr//kZ"

	fmt.Printf("%v", sig)

	res := new(resolver.Resolver)
	ch  := res.NewQuerier()

	// configure the resolver
	res.Servers = []string{"192.168.1.2"}
	res.Timeout = 2
	res.Attempts = 1

	// Setup done, now for some real work
	// Create a new message
	m := new(dns.Msg)
	m.MsgHdr.RecursionDesired = true //only set this bit
	m.Question = make([]dns.Question, 1)

	m.Question[0] = dns.Question{"miek.nl", dns.TypeDS, dns.ClassINET}
	ch <- resolver.DnsMsg{m, nil}
	in := <-ch
	fmt.Printf("%v\n", in.Dns)

	m.Question[0] = dns.Question{"www.nlnetlabs.nl", dns.TypeRRSIG, dns.ClassINET}
	ch <- resolver.DnsMsg{m, nil}
	in = <-ch
	fmt.Printf("%v\n", in.Dns)

	m.Question[0] = dns.Question{"xxxx.nlnetlabs.nl", dns.TypeDNSKEY, dns.ClassINET}
	ch <- resolver.DnsMsg{m, nil}
	in = <-ch
	fmt.Printf("%v\n", in.Dns)

	ch <- resolver.DnsMsg{nil, nil}
        <-ch
}
