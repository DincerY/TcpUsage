package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Bir TCP dinleme soketi oluşturulur.
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	go client()

	// Yeni bağlantılar için sonsuz bir döngü başlatılır.
	for {
		// Bir bağlantı kabul edilir.
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Bağlantıdan gelen veri okunur.
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Okunan veri işlenir.
		fmt.Println("Alınan veri:", string(buf[:n]))

		// Bağlantıya bir yanıt gönderilir.
		_, err = conn.Write([]byte("Merhaba, Dünya!"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Bağlantı kapatılır.
		conn.Close()
	}
}

func client() {
	for {
		// Sunucuya bir bağlantı kurulur.
		time.Sleep(time.Second * 1)
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println(err)
			return
		}

		// "Merhaba, Dünya!" metni bağlantıya gönderilir.
		_, err = conn.Write([]byte("Merhaba, Dünya!"))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Bağlantıdan gelen veri okunur.
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Okunan veri ekrana yazdırılır.
		fmt.Println("Alınan veri:", string(buf[:n]))

		// Bağlantı kapatılır.
		conn.Close()

	}
}
