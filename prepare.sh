#!/bin/bash

set -e

THREADS=${1:-8}

MMD_LIST="GeoIP2-Anonymous-IP-Test
GeoIP2-City-Shield-Test
GeoIP2-City-Test-Broken-Double-Format
GeoIP2-City-Test-Invalid-Node-Count
GeoIP2-City-Test
GeoIP2-Connection-Type-Test
GeoIP2-Country-Shield-Test
GeoIP2-Country-Test
GeoIP2-DensityIncome-Test
GeoIP2-Domain-Test
GeoIP2-Enterprise-Shield-Test
GeoIP2-Enterprise-Test
GeoIP2-IP-Risk-Test
GeoIP2-ISP-Test
GeoIP2-Precision-Enterprise-Shield-Test
GeoIP2-Precision-Enterprise-Test
GeoIP2-Static-IP-Score-Test
GeoIP2-User-Count-Test
GeoIP-Anonymous-Plus
GeoIP-Anonymous-Plus-Test
GeoLite2-ASN-Test
GeoLite2-City-Test
GeoLite2-Country-Test
MaxMind-DB-no-ipv4-search-tree
MaxMind-DB-string-value-entries
MaxMind-DB-test-broken-pointers-24
MaxMind-DB-test-broken-search-tree-24
MaxMind-DB-test-decoder
MaxMind-DB-test-ipv4-24
MaxMind-DB-test-ipv4-28
MaxMind-DB-test-ipv4-32
MaxMind-DB-test-ipv6-24
MaxMind-DB-test-ipv6-28
MaxMind-DB-test-ipv6-32
MaxMind-DB-test-metadata-pointers
MaxMind-DB-test-mixed-24
MaxMind-DB-test-mixed-28
MaxMind-DB-test-mixed-32
MaxMind-DB-test-nested
MaxMind-DB-test-pointer-decoder"

echo "download test mmdb file in $THREADS threads:"

echo "$MMD_LIST" | xargs -P "$THREADS" -I {} bash -c '
    file="$1"
    url="https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/${file}.mmdb"
    output="testdata/${file}.mmdb"

    if command -v curl >/dev/null 2>&1; then
        if curl -L -f -s -o "$output" "$url"; then
            echo "* $file"
            exit 0
        fi
    elif command -v wget >/dev/null 2>&1; then
        if wget -q -O "$output" "$url"; then
            echo "* $file"
            exit 0
        fi
    else
        echo "Error: neither curl nor wget available"
        exit 1
    fi

    echo "failed $file"
    exit 1
' _ {}

echo "done"
