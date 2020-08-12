
. ./scripts/env.sh

function createOrg1 {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/peerOrganizations/org1.develop.com/

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.develop.com/
#  rm -rf $FABRIC_CA_CLIENT_HOME/${FABRIC_CA_CLIENT}-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  ${FABRIC_CA_CLIENT} enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/config.yaml

  echo
	echo "Register peer0"
  echo
  set -x
	${FABRIC_CA_CLIENT} register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  echo
  echo "Register user"
  echo
  set -x
  ${FABRIC_CA_CLIENT} register --caname ca-org1 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  echo
  echo "Register the org admin"
  echo
  set -x
  ${FABRIC_CA_CLIENT} register --caname ca-org1 --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

	mkdir -p organizations/peerOrganizations/org1.develop.com/peers
  mkdir -p organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com

  echo
  echo "## Generate the peer0 msp"
  echo
  set -x
	${FABRIC_CA_CLIENT} enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/msp --csr.hosts peer0.org1.develop.com --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/msp/config.yaml

  echo
  echo "## Generate the peer0-tls certificates"
  echo
  set -x
  ${FABRIC_CA_CLIENT} enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls --enrollment.profile tls --csr.hosts peer0.org1.develop.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x


  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/ca.crt
  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/signcerts/* ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/server.crt
  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/keystore/* ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/server.key

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/tlscacerts
  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/tlscacerts/ca.crt

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.develop.com/tlsca
  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/tls/tlscacerts/* ${PWD}/organizations/peerOrganizations/org1.develop.com/tlsca/tlsca.org1.develop.com-cert.pem

  mkdir -p ${PWD}/organizations/peerOrganizations/org1.develop.com/ca
  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/peers/peer0.org1.develop.com/msp/cacerts/* ${PWD}/organizations/peerOrganizations/org1.develop.com/ca/ca.org1.develop.com-cert.pem

  mkdir -p organizations/peerOrganizations/org1.develop.com/users
  mkdir -p organizations/peerOrganizations/org1.develop.com/users/User1@org1.develop.com

  echo
  echo "## Generate the user msp"
  echo
  set -x
	${FABRIC_CA_CLIENT} enroll -u https://user1:user1pw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.develop.com/users/User1@org1.develop.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.develop.com/users/User1@org1.develop.com/msp/config.yaml

  mkdir -p organizations/peerOrganizations/org1.develop.com/users/Admin@org1.develop.com

  echo
  echo "## Generate the org admin msp"
  echo
  set -x
	${FABRIC_CA_CLIENT} enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-org1 -M ${PWD}/organizations/peerOrganizations/org1.develop.com/users/Admin@org1.develop.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/org1/tls-cert.pem
  set +x

  cp ${PWD}/organizations/peerOrganizations/org1.develop.com/msp/config.yaml ${PWD}/organizations/peerOrganizations/org1.develop.com/users/Admin@org1.develop.com/msp/config.yaml

}

function createOrderer {

  echo
	echo "Enroll the CA admin"
  echo
	mkdir -p organizations/ordererOrganizations/develop.com

	export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/develop.com
#  rm -rf $FABRIC_CA_CLIENT_HOME/${FABRIC_CA_CLIENT}-config.yaml
#  rm -rf $FABRIC_CA_CLIENT_HOME/msp

  set -x
  ${FABRIC_CA_CLIENT} enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > ${PWD}/organizations/ordererOrganizations/develop.com/msp/config.yaml


  echo
	echo "Register orderer"
  echo
  set -x
	${FABRIC_CA_CLIENT} register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
    set +x

  echo
  echo "Register the orderer admin"
  echo
  set -x
  ${FABRIC_CA_CLIENT} register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

	mkdir -p organizations/ordererOrganizations/develop.com/orderers
  mkdir -p organizations/ordererOrganizations/develop.com/orderers/develop.com

  mkdir -p organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com

  echo
  echo "## Generate the orderer msp"
  echo
  set -x
	${FABRIC_CA_CLIENT} enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/msp --csr.hosts orderer.develop.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/develop.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/msp/config.yaml

  echo
  echo "## Generate the orderer-tls certificates"
  echo
  set -x
  ${FABRIC_CA_CLIENT} enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls --enrollment.profile tls --csr.hosts orderer.develop.com --csr.hosts localhost --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/ca.crt
  cp ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/signcerts/* ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/server.crt
  cp ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/keystore/* ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/server.key

  mkdir -p ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/msp/tlscacerts/tlsca.develop.com-cert.pem

  mkdir -p ${PWD}/organizations/ordererOrganizations/develop.com/msp/tlscacerts
  cp ${PWD}/organizations/ordererOrganizations/develop.com/orderers/orderer.develop.com/tls/tlscacerts/* ${PWD}/organizations/ordererOrganizations/develop.com/msp/tlscacerts/tlsca.develop.com-cert.pem

  mkdir -p organizations/ordererOrganizations/develop.com/users
  mkdir -p organizations/ordererOrganizations/develop.com/users/Admin@develop.com

  echo
  echo "## Generate the admin msp"
  echo
  set -x
	${FABRIC_CA_CLIENT} enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M ${PWD}/organizations/ordererOrganizations/develop.com/users/Admin@develop.com/msp --tls.certfiles ${PWD}/organizations/fabric-ca/ordererOrg/tls-cert.pem
  set +x

  cp ${PWD}/organizations/ordererOrganizations/develop.com/msp/config.yaml ${PWD}/organizations/ordererOrganizations/develop.com/users/Admin@develop.com/msp/config.yaml


}
