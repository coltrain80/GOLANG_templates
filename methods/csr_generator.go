// csr_generator.go
/*
CSR Generator in Go
--------------------

This script generates a Certificate Signing Request (CSR) using the crypto package.

Usage:
    Run the script to generate a CSR and private key in PEM format.

Example:
    go run csr_generator.go
*/

package main

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "fmt"
    "os"
)

func main() {
    // Generate a new RSA private key
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        fmt.Printf("Failed to generate private key: %v\n", err)
        return
    }

    // Define the subject of the certificate
    subject := pkix.Name{
        Country:            []string{"US"},
        Province:           []string{"California"},
        Locality:           []string{"San Francisco"},
        Organization:       []string{"Example Corp"},
        OrganizationalUnit: []string{"IT"},
        CommonName:         "example.com",
    }

    // Create a CSR template
    csrTemplate := x509.CertificateRequest{
        Subject:            subject,
        SignatureAlgorithm: x509.SHA256WithRSA,
    }

    // Generate the CSR
    csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
    if err != nil {
        fmt.Printf("Failed to create CSR: %v\n", err)
        return
    }

    // Encode and save the CSR to a file
    csrFile, err := os.Create("example.csr")
    if err != nil {
        fmt.Printf("Failed to create CSR file: %v\n", err)
        return
    }
    defer csrFile.Close()

    pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})
    fmt.Println("CSR written to example.csr")

    // Encode and save the private key to a file
    keyFile, err := os.Create("private_key.pem")
    if err != nil {
        fmt.Printf("Failed to create private key file: %v\n", err)
        return
    }
    defer keyFile.Close()

    pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
    fmt.Println("Private key written to private_key.pem")
}
