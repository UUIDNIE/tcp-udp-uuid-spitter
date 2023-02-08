Handy tool for testing filters / open ports with UUIDs.  

This program listens on TCP and UDP port 8080.  

Fairly simple to expand this to all ports, excluding port 22 by adding a nat prerouting rule to iptables on Linux.  

iptables -t nat -A PREROUTING -p tcp --dport 22 -j ACCEPT
iptables -t nat -A PREROUTING -p udp --dport 22 -j ACCEPT
iptables -t nat -A PREROUTING -p tcp --dport 1:65535 -j REDIRECT --to-port 8080
iptables -t nat -A PREROUTING -p udp --dport 1:65535 -j REDIRECT --to-port 8080

Once you are finished, build and run the program.  Try hitting a port with Telnet and see if you get a UUID :)
