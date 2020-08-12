#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    local OP=$(one_line_pem $7)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s/\${ORDERER_PORT}/$6/" \
        -e "s#\${ORDERER_PEM}#$OP#" \
        ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=1
P0PORT=7051
CAPORT=7054
ORDERER_PORT=7050
PEERPEM=organizations/peerOrganizations/org1.develop.com/tlsca/tlsca.org1.develop.com-cert.pem
CAPEM=organizations/peerOrganizations/org1.develop.com/ca/ca.org1.develop.com-cert.pem
ORDERER_PEM=organizations/fabric-ca/ordererOrg/ca-cert.pem

echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORDERER_PORT $ORDERER_PEM)" > organizations/peerOrganizations/org1.develop.com/connection-org1.yaml
