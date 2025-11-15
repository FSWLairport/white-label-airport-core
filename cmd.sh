go mod tidy
TAGS=with_gvisor,with_quic,with_wireguard,with_utls,with_clash_api,with_grpc,badlinkname,tfogo_checklinkname0
# TAGS=with_dhcp,with_low_memory,with_conntrack
go run -ldflags="-checklinkname=0" --tags $TAGS ./cli  $@
