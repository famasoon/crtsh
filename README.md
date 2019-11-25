crtsh
===

`crtsh` is **[crt.sh](https://crt.sh)** Golang utility

## Installation

```sh
go get github.com/famasoon/crtsh
```

## Usage
`crtsh` has some option.

### `-q` option
The `-q` option is to query to [https://crt.sh](https://crt.sh)
The result is dictionary items which looks like this:

```sh
$ crtsh -q example.com
{
  Index: 1
  Issuer CA ID: 1191
  Issuer Name: C=US, O=DigiCert Inc, CN=DigiCert SHA2 Secure Server CA
  Name: example.com
  Min Cert ID: 987119772
  Min Entry TimeStamp: 2018-11-29T13:44:14.118
  Not Before: 2018-11-28T00:00:00
  Not After: 2020-12-02T12:00:00
  Donwload Pem file: https://crt.sh/?d=987119772
}
{
  Index: 2
  Issuer CA ID: 1191
  Issuer Name: C=US, O=DigiCert Inc, CN=DigiCert SHA2 Secure Server CA
  Name: example.com
  Min Cert ID: 984858191
  Min Entry TimeStamp: 2018-11-28T21:20:12.606
  Not Before: 2018-11-28T00:00:00
  Not After: 2020-12-02T12:00:00
  Donwload Pem file: https://crt.sh/?d=984858191
}
{
  Index: 3
  Issuer CA ID: 1465
  Issuer Name: C=US, O="thawte, Inc.", CN=thawte SSL CA - G2
  Name: example.com
  Min Cert ID: 24564717
  Min Entry TimeStamp: 2016-07-14T07:55:01.55
  Not Before: 2016-07-14T00:00:00
  Not After: 2017-07-14T23:59:59
  Donwload Pem file: https://crt.sh/?d=24564717
}
{
  Index: 4
  Issuer CA ID: 1465
  Issuer Name: C=US, O="thawte, Inc.", CN=thawte SSL CA - G2
  Name: example.com
  Min Cert ID: 24560643
  Min Entry TimeStamp: 2016-07-14T07:30:08.461
  Not Before: 2016-07-14T00:00:00
  Not After: 2018-07-14T23:59:59
  Donwload Pem file: https://crt.sh/?d=24560643
}
{
  Index: 5
  Issuer CA ID: 1465
  Issuer Name: C=US, O="thawte, Inc.", CN=thawte SSL CA - G2
  Name: example.com
  Min Cert ID: 24560621
  Min Entry TimeStamp: 2016-07-14T07:25:01.93
  Not Before: 2016-07-14T00:00:00
  Not After: 2017-07-14T23:59:59
  Donwload Pem file: https://crt.sh/?d=24560621
}
{
  Index: 6
  Issuer CA ID: 1449
  Issuer Name: C=US, O=Symantec Corporation, OU=Symantec Trust Network, CN=Symantec Class 3 Secure Server CA - G4
  Name: example.com
  Min Cert ID: 24558997
  Min Entry TimeStamp: 2016-07-14T06:40:02.4
  Not Before: 2016-07-14T00:00:00
  Not After: 2018-07-14T23:59:59
  Donwload Pem file: https://crt.sh/?d=24558997
}
{
  Index: 7
  Issuer CA ID: 1397
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 High Assurance Server CA
  Name: example.com
  Min Cert ID: 10557607
  Min Entry TimeStamp: 2015-11-05T14:51:33.941
  Not Before: 2015-11-03T00:00:00
  Not After: 2018-11-28T12:00:00
  Donwload Pem file: https://crt.sh/?d=10557607
}
{
  Index: 8
  Issuer CA ID: 1397
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 High Assurance Server CA
  Name: example.com
  Min Cert ID: 5857507
  Min Entry TimeStamp: 2014-12-11T14:36:57.201
  Not Before: 2014-11-06T00:00:00
  Not After: 2015-11-13T12:00:00
  Donwload Pem file: https://crt.sh/?d=5857507
}
```

This option can query to use wildcard (% = wildcard) and `_`  (_ = completing input)

For Example:

```sh
$ crtsh -q %.example.com -o
www.example.com
www.example.com
www.example.com
*.example.com
*.example.com
m.example.com
www.example.com
dev.example.com
products.example.com
support.example.com
www.example.com
www.example.com
www.example.com
```

```sh
crtsh -q www.sagawa_exp.co.jp
{
  Index: 1
  Issuer CA ID: 1399
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Extended Validation Server CA
  Name: www.sagawa-exp.co.jp
  Min Cert ID: 2114355943
  Min Entry TimeStamp: 2019-11-15T09:15:37.331
  Not Before: 2019-11-12T00:00:00
  Not After: 2020-11-20T12:00:00
  Donwload Pem file: https://crt.sh/?d=2114355943
}
{
  Index: 2
  Issuer CA ID: 1399
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Extended Validation Server CA
  Name: www.sagawa-exp.co.jp
  Min Cert ID: 2101305573
  Min Entry TimeStamp: 2019-11-12T05:39:40.762
  Not Before: 2019-11-12T00:00:00
  Not After: 2020-11-20T12:00:00
  Donwload Pem file: https://crt.sh/?d=2101305573
}
{
  Index: 3
  Issuer CA ID: 1399
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Extended Validation Server CA
  Name: www.sagawa-exp.co.jp
  Min Cert ID: 1470039453
  Min Entry TimeStamp: 2019-05-14T11:02:28.185
  Not Before: 2018-11-20T00:00:00
  Not After: 2019-11-21T12:00:00
  Donwload Pem file: https://crt.sh/?d=1470039453
}
{
  Index: 4
  Issuer CA ID: 1399
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert SHA2 Extended Validation Server CA
  Name: www.sagawa-exp.co.jp
  Min Cert ID: 961131139
  Min Entry TimeStamp: 2018-11-20T05:15:04.525
  Not Before: 2018-11-20T00:00:00
  Not After: 2019-11-21T12:00:00
  Donwload Pem file: https://crt.sh/?d=961131139
}
```

---

And `-q` option can use `-o` option.

The `-o` option only enumerates domains.

```sh
$ crtsh -q example.com -o
example.com
example.com
example.com
example.com
example.com
example.com
example.com
example.com
```

### `-cn` option
THe `-cn` option query CommonName.

For Example: `crtsh -cn <CommonName>

```sh
$ crtsh -cn test
{
  Index: 1
  Issuer CA ID: 6831
  Issuer Name: C=BE, O=GlobalSign nv-sa, CN=GlobalSign PersonalSign 2 CA - G2
  Name: Test
  Min Cert ID: 197744191
  Min Entry TimeStamp: 2017-08-24T18:23:36.43
  Not Before: 2014-07-31T20:44:32
  Not After: 2015-08-01T20:44:32
  Donwload Pem file: https://crt.sh/?d=197744191
}
{
  Index: 2
  Issuer CA ID: 750
  Issuer Name: emailAddress=contacto@procert.net.ve, L=Chacao, ST=Miranda, OU=Proveedor de Certificados PROCERT, O=Sistema Nacional de Certificacion Electronica, C=VE, CN=PSCProcert
  Name: test
  Min Cert ID: 197155020
  Min Entry TimeStamp: 2017-08-23T22:07:22.88
  Not Before: 2017-08-23T13:05:28
  Not After: 2018-08-23T13:05:28
  Donwload Pem file: https://crt.sh/?d=197155020
}
{
  Index: 3
  Issuer CA ID: 750
  Issuer Name: emailAddress=contacto@procert.net.ve, L=Chacao, ST=Miranda, OU=Proveedor de Certificados PROCERT, O=Sistema Nacional de Certificacion Electronica, C=VE, CN=PSCProcert
  Name: test
  Min Cert ID: 197073488
  Min Entry TimeStamp: 2017-08-23T19:42:20.529
  Not Before: 2017-08-23T13:11:13
  Not After: 2018-08-23T13:11:13
  Donwload Pem file: https://crt.sh/?d=197073488
}
{
  Index: 4
  Issuer CA ID: 1715
  Issuer Name: C=CN, O=CNNIC SHA256 SSL, CN=CNNIC SHA256 SSL
  Name: test
  Min Cert ID: 7096879
  Min Entry TimeStamp: 2015-04-08T00:24:19.637
  Not Before: 2014-12-12T06:08:52
  Not After: 2015-12-12T06:08:52
  Donwload Pem file: https://crt.sh/?d=7096879
}
{
  Index: 5
  Issuer CA ID: 1715
  Issuer Name: C=CN, O=CNNIC SHA256 SSL, CN=CNNIC SHA256 SSL
  Name: test
  Min Cert ID: 7096563
  Min Entry TimeStamp: 2015-04-08T00:11:13.016
  Not Before: 2014-12-14T12:00:54
  Not After: 2015-12-14T12:00:54
  Donwload Pem file: https://crt.sh/?d=7096563
}
{
  Index: 6
  Issuer CA ID: 29
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance CA-3
  Name: test
  Min Cert ID: 4202482
  Min Entry TimeStamp: 2014-05-22T23:21:36.633
  Not Before: 2011-07-28T00:00:00
  Not After: 2014-08-01T12:00:00
  Donwload Pem file: https://crt.sh/?d=4202482
}
{
  Index: 7
  Issuer CA ID: 29
  Issuer Name: C=US, O=DigiCert Inc, OU=www.digicert.com, CN=DigiCert High Assurance CA-3
  Name: test
  Min Cert ID: 4202481
  Min Entry TimeStamp: 2014-05-22T23:21:33.786
  Not Before: 2011-07-28T00:00:00
  Not After: 2014-08-01T12:00:00
  Donwload Pem file: https://crt.sh/?d=4202481
}
```


### `-i` option
The `-i` option parse pem file.
If you set this option, you can enumerate DNS records that was implanted pem file.
I will add more features.

For Example: `crtsh -i <Min Cert ID>`

```sh
$ crtsh -i 5857507
CertID: 5857507
Enumrate DNS Names:
www.example.org
example.com
example.edu
example.net
example.org
www.example.com
www.example.edu
www.example.net
```

## Importing
```go
import (
    "github.com/famasoon/crtsh/ctlog"
    "github.com/famasoon/crtsh/parser"
)
```