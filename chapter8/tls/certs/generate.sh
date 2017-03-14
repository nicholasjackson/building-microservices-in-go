#!/bin/bash -e

echo 
echo Generate the root key
echo ---
openssl genrsa -aes256 -out 1_root/private/ca.key.pem 4096
chmod 444 1_root/private/ca.key.pem

echo 
echo Generate the root certificate
echo ---
openssl req -config 1_root/openssl.cnf \
      -key 1_root/private/ca.key.pem \
      -new -x509 -days 7300 -sha256 -extensions v3_ca \
      -out 1_root/certs/ca.cert.pem

echo 
echo Verify root key
echo ---
openssl x509 -noout -text -in 1_root/certs/ca.cert.pem

echo 
echo Generate the key for the intermediary certificate
echo ---
openssl genrsa -aes256 \
  -out 2_intermediate/private/intermediate.key.pem 4096

chmod 444 2_intermediate/private/intermediate.key.pem

echo 
echo Generate the signing request for the intermediary certificate
echo ---
openssl req -config 1_root/openssl.cnf -new -sha256 \
      -key 2_intermediate/private/intermediate.key.pem \
      -out 2_intermediate/csr/intermediate.csr.pem

echo 
echo Sign the intermediary
echo ---
openssl ca -config 1_root/openssl.cnf -extensions v3_intermediate_ca \
        -days 3650 -notext -md sha256 \
        -in 2_intermediate/csr/intermediate.csr.pem \
        -out 2_intermediate/certs/intermediate.cert.pem


chmod 444 2_intermediate/certs/intermediate.cert.pem

echo 
echo Verify intermediary
echo ---
openssl x509 -noout -text \
      -in 2_intermediate/certs/intermediate.cert.pem

openssl verify -CAfile 1_root/certs/ca.cert.pem \
      2_intermediate/certs/intermediate.cert.pem

echo 
echo Create the chain file
echo ---
cat 2_intermediate/certs/intermediate.cert.pem \
      1_root/certs/ca.cert.pem > 2_intermediate/certs/ca-chain.cert.pem
chmod 444 2_intermediate/certs/ca-chain.cert.pem

echo 
echo Create the application key
echo ---
openssl genrsa \
    -out 3_application/private/www.example.com.key.pem 2048
chmod 444 3_application/private/www.example.com.key.pem

echo 
echo Create the application signing request
echo ---
openssl req -config 2_intermediate/openssl.cnf \
      -key 3_application/private/www.example.com.key.pem \
      -new -sha256 -out 3_application/csr/www.example.com.csr.pem

echo 
echo Create the application certificate
echo ---
openssl ca -config 2_intermediate/openssl.cnf \
      -extensions server_cert -days 375 -notext -md sha256 \
      -in 3_application/csr/www.example.com.csr.pem \
      -out 3_application/certs/www.example.com.cert.pem
chmod 444 3_application/certs/www.example.com.cert.pem

echo 
echo Validate the certificate
echo ---
openssl x509 -noout -text \
      -in 3_application/certs/www.example.com.cert.pem

echo 
echo Validate the certificate has the correct chain of trust
echo ---
openssl verify -CAfile 2_intermediate/certs/ca-chain.cert.pem \
      3_application/certs/www.example.com.cert.pem
