package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	jwt.StandardClaims
}

var key = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAo+fW9kv2Y5Fjky8TQLK3rYtjqCKIEf0Yjm1lXTnkpWVBjO7x
EzB7HYBkHnIA19HEdzjeKLlLcJevOU6G3p+t8/5vVhciPnhKpwU6ZzrR0P3Q2toC
+KgtrnPHXpm8Py/HVwBxHWZ352NNa5J3dHv6a7B1k7IMAaBT053P99l1NQrBCVJf
kXckQAUOnsYk/PKFFufNhlu3nAd+eUl0Iv1IAWsUb5lHKkIOOWxnlreFN5gdCgL6
YxLwXRBV010zu/Y9zF8zdhSZXWoIvD/JpOvkKh9ym1VE0YzYkSW+m0XnTa4VMdZf
mWcQNKhYZdNg0TqwRco893fVeo756JzOHQ0KcQIDAQABAoIBADnxO4zWTcPlIc5m
VegJReWT4ScPDgtN7eBry+mpDatqoEGyNocSHHPRb5NTufiRr2J2OBMtbf3foZIg
sI0C5fvbdoB4rSJoY/unXX4gzQuUrsCvUV2WFpAVXeWTc3ji6xVWkqZSExE5iaT9
oj0Llvt65hXQW243v2qP33U+6rBVC/2NozGzlq9uHvNJURF0/ZztIo47K6q+giZC
ONvycXVGT8h5PIzfQsnnrsaLYCNe+yJ1x9zPT32xStI8+rF/PFBO1T8umna6+kY2
pY5XQfiYH7gfUIqfjA48Ovnxdmyk9ib6DzcgB+vlY+4cgCGX27sbG99pxGKu8nGg
iJjBGAECgYEA1C4N+RweFi2vzCBLnrSCZ2U6XjFJyg9ugL+X01DM4r3ORVlFV5WZ
2ebKS0Efwh3Jg6fuvf/2/6sFPzH4x8GIKb8Wldsh9mAGUQzOJ+lURqaEnbw+5xSs
y77ORNUes/mdr5YmKpyFsKZXNDKasIeQ7/usp5BwcuOtk5F7+56u+80CgYEAxcGG
g3XNpbqW3mmGf59au+eofYbzliWXFKf3S6A1Mp2EAG02Lb7H2sNQKI2GXMABLeAf
ylNLigKnSBgc2n4Ice91ddkjJhz5zk1P6Rpjz4yChMfC0qqqDbrpwrjuzHRKtWNJ
xh+3qPiZzmpgfK87ddNo/4pTXYBj/rtWocWvjTUCgYBie7nqnV1tp+kIExlmYZyB
h1/PJot8aSs+QS+kWsWunLDoSWZBH+QYWuIcie7Gt/K31DOhJvSreOLnkTdK6I5d
6h7+kYazB6EG762kos2GOXYmjKCZu2P08exl0JH+sWa6gDPY5Wu8MYkaZj6cn1/r
s+JQqF85RGplq0pj4SSRJQKBgQCpFX4QutHZqP9ELc/tIPBwh3Nd2Nw+/eb+p1rf
U50IqPtrbfWMCXpSBvtroQ5IEXcwpVgpIy0MVJZ5IvzQqEBKQqcY206dUNsaVKpF
seWzWP1j0HU4sOlzkeQ7NOog8DHMg5dZilb7Z4wCnJvhH+bkKJ23GKuM37Ef5Uf+
S8t9HQKBgD3bw7YZpnpP05fMPVcpLLOdxOmd8h6kEKqzOu33PHreMl9GCK7J8eOF
oWqAK1ub7vYkke6teZlPjYyrqutTVCGBZXn2ctuDc/o4t17Qk8hXbQBn6W0+XggB
sV0tGy1CL+hkyk3EuUUDPapwXmSoA5P5eJqN16XiqSQCv7Cis0zr
-----END RSA PRIVATE KEY-----`

var pubkey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo+fW9kv2Y5Fjky8TQLK3
rYtjqCKIEf0Yjm1lXTnkpWVBjO7xEzB7HYBkHnIA19HEdzjeKLlLcJevOU6G3p+t
8/5vVhciPnhKpwU6ZzrR0P3Q2toC+KgtrnPHXpm8Py/HVwBxHWZ352NNa5J3dHv6
a7B1k7IMAaBT053P99l1NQrBCVJfkXckQAUOnsYk/PKFFufNhlu3nAd+eUl0Iv1I
AWsUb5lHKkIOOWxnlreFN5gdCgL6YxLwXRBV010zu/Y9zF8zdhSZXWoIvD/JpOvk
Kh9ym1VE0YzYkSW+m0XnTa4VMdZfmWcQNKhYZdNg0TqwRco893fVeo756JzOHQ0K
cQIDAQAB
-----END PUBLIC KEY-----`

func main() {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		log.Println("errorPrivate:", err)
	}
	_, err = jwt.ParseRSAPublicKeyFromPEM([]byte(pubkey))
	if err != nil {
		log.Println("errorPublic:", err)
	}

	fmt.Println("Done")

	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        "uuid",
			Subject:   "userId",
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println("errorSignedString:", err)
	}
	log.Println(signedString)
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"jwt-rsa-golang/token"
//)
//
//func main() {
//	prvKey, err := os.ReadFile("cert/id_rsa")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	pubKey, err := os.ReadFile("cert/id_rsa.pub")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	jwtToken := token.NewJWT(prvKey, pubKey)
//
//	// 1. Create a new JWT token.
//	tok, err := jwtToken.Create(time.Hour, "Can be anything")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println("TOKEN:", tok)
//
//	// 2. Validate an existing JWT token.
//	content, err := jwtToken.Validate(tok)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println("CONTENT:", content)
//}
