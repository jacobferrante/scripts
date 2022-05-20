#!/bin/bash

## If you need to build a ntp.conf that's on a system with one simple network interface, this could work for you.
## it just pulls the subnet of the system you're running on and builds the file with that, in our case all our NTP
## servers run on .32

mkdir -p /var/log/ntp
subnet="ifconfig -a | grep inet | grep 192 | cut -d"." -f3"

cat <<EOF > /etc/ntp.conf

driftfile /var/log/ntp/ntp.drift
logfile /var/log/ntp/ntp.log
logconfig =all

server 192.168.$subnet.32 iburst
server 127.127.1.0

fudge 127.127.1.0 stratum 14

restrict default

EOF
