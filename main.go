package main

import (
	"fmt"
	"github.com/miekg/dns"
	"os"
	"os/exec"
)

func main() {
	if err := dns.ListenAndServe(":8053", "udp", dns.HandlerFunc(func(w dns.ResponseWriter, r *dns.Msg) {
		name := r.Question[0].Name[:len(r.Question[0].Name)-1]

		resp, err := exec.Command("nslookup", name).CombinedOutput()
		if err != nil {
			return
		}

		ip := handleRemoteIp(resp)
		respMsg := new(dns.Msg)
		respMsg.SetReply(r)
		if ip != nil {
			respMsg.Answer = append(
				respMsg.Answer,
				&dns.A{
					Hdr: dns.RR_Header{
						Name:     r.Question[0].Name,
						Rrtype:   dns.TypeA,
						Class:    dns.ClassINET,
						Ttl:      0,
						Rdlength: 1,
					},
					A: ip,
				},
			)
		} else {
			respMsg.Rcode = dns.RcodeNameError
		}

		w.WriteMsg(respMsg)
	})); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
