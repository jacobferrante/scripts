# scripts
some smaller bash/golang/python scripts I use


# ntp_conf.sh
this is just a quick and dirty way to build a NTP.conf file for your systems if you have a one network interface like the embedded systems I have I worked on, it also helped to have all the NTP servers on the same IP with different subnets. 

#file_transfer.go
I built this script to quickly deploy files on local windows systems, where we were doing software/hardware upgrades and were on site, instead of dispering the files and risk putting them in the wrong place/wrong name, etc. this would fix that issue and do it automatically with a single click and let us know if we had any naming errors or missing files.