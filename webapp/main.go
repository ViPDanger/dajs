package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const host = "0.0.0.0:8085"

var certFile = "/etc/letsencrypt/live/dajs.vipdanger.keenetic.pro/fullchain.pem"
var keyFile = "/etc/letsencrypt/live/dajs.vipdanger.keenetic.pro/privkey.pem"
var exeDir string

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir = filepath.Dir(exePath)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить все (в проде использовать конкретные)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.LoadHTMLGlob(exeDir + "/character/*")
	r.Static("/character", exeDir+"/character")
	r.GET("/char", func(c *gin.Context) {
		var tokens struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}
		//if err := c.BindJSON(&tokens); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		//return
		//}
		c.HTML(http.StatusOK, "character.html", gin.H{
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		})
	})
	fmt.Println("webapp started on ", host)
	log.Println(r.RunTLS(host, certFile, keyFile))
	//log.Println(srv.ListenAndServeTLS(exeDir+"/server_cert.pem", exeDir+"/server_key.pem"))
}

func CreateCA() (*x509.Certificate, *ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	serialNumber, _ := rand.Int(rand.Reader, big.NewInt(1<<62))

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My CA Org"},
			Country:      []string{"RU"},
			CommonName:   "My Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 10 лет
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLen:            2,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, nil, err
	}

	// Сохраняем CA сертификат и ключ в файлы (опционально)
	pemEncode("/etc/letsencrypt/live/dajs/ca_cert.pem", "CERTIFICATE", certDER)
	keyBytes, _ := x509.MarshalECPrivateKey(priv)
	pemEncode(exeDir+"/ca_key.pem", "EC PRIVATE KEY", keyBytes)

	return cert, priv, nil
}

// Создаём серверный сертификат, подписанный CA
func CreateServerCert(caCert *x509.Certificate, caKey *ecdsa.PrivateKey) error {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	serialNumber, _ := rand.Int(rand.Reader, big.NewInt(1<<62))

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"My Server Org"},
			Country:      []string{"RU"},
			CommonName:   "localhost",
		},
		DNSNames:    []string{"localhost"},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0), // 1 год
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, caCert, &priv.PublicKey, caKey)
	if err != nil {
		return err
	}

	pemEncode(exeDir+"/server_cert.pem", "CERTIFICATE", certDER)
	keyBytes, _ := x509.MarshalECPrivateKey(priv)
	pemEncode(exeDir+"/server_key.pem", "EC PRIVATE KEY", keyBytes)

	return nil
}

func pemEncode(filename, blockType string, bytes []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return pem.Encode(f, &pem.Block{Type: blockType, Bytes: bytes})
}
